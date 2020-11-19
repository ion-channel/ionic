package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	goblin "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/products"
	. "github.com/onsi/gomega"
)

func TestProducts(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Products", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient
		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})
		g.AfterEach(func() {
			server.Close()
		})

		g.It("should get a product", func() {
			server.AddPath("/v1/vulnerability/getProducts").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProduct)).
				SetStatus(http.StatusOK)

			product, err := client.GetProducts("cpe:/a:oracle:jdk:1.6.0:update_71", "someapikey")
			Expect(err).To(BeNil())
			Expect(product[0].Sources[0].Name).To(Equal("NVD"))
			Expect(product[0].ID).To(Equal(84647))
			Expect(product[0].Name).To(Equal("jdk"))
		})

		g.It("should get a product's versions", func() {
			server.AddPath("/v1/product/getProductVersions").
				SetMethods("GET").
				SetPayload([]byte(sameProductVersions)).
				SetStatus(http.StatusOK)

			product, err := client.GetProductVersions("jdk", "11.0", "someapikey")
			Expect(err).To(BeNil())
			Expect(product[0].Version).To(Equal("13.0.0"))

		})

		g.It("should get a raw product", func() {
			server.AddPath("/v1/vulnerability/getProducts").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProduct)).
				SetStatus(http.StatusOK)

			raw, err := client.GetRawProducts("cpe:/a:oracle:jdk:1.6.0:update_71", "fookey")
			Expect(err).To(BeNil())
			Expect(raw).To(Equal(json.RawMessage(SampleValidRawProduct)))
		})

		g.It("should search for a product", func() {
			server.AddPath("/v1/product/search").
				SetMethods("GET").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			products, _, err := client.GetProductSearch("less+mahVersion", nil, "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer someapikey"))
			Expect(hitRecord.Query.Get("q")).To(Equal("less+mahVersion"))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
		g.It("should search for a product by pages", func() {
			server.AddPath("/v1/product/search").
				SetMethods("GET").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			page := pagination.Pagination{
				Offset: 0,
				Limit:  100,
			}
			products, _, err := client.GetProductSearch("less+mahVersion", &page, "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer someapikey"))
			Expect(hitRecord.Query.Get("q")).To(Equal("less+mahVersion"))
			Expect(hitRecord.Query.Get("offset")).To(Equal("0"))
			Expect(hitRecord.Query.Get("limit")).To(Equal("100"))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
		g.It("should omit version from product search when it is not given", func() {
			server.AddPath("/v1/product/search").
				SetMethods("GET").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			products, _, err := client.GetProductSearch("less", nil, "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer someapikey"))
			Expect(hitRecord.Query.Get("q")).To(Equal("less"))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
		g.It("should unmarshal search results with scores", func() {
			searchResultJSON := `{"product":{"name":"django","language":"","source":null,"created_at":"2017-02-13T20:02:35.667Z","title":"Django Project Django 1.0-alpha-1","up":"alpha1","updated_at":"2017-02-13T20:02:35.667Z","edition":"","part":"/a","references":[],"version":"1.0","org":"djangoproject","external_id":"cpe:/a:djangoproject:django:1.0:alpha1","id":30955,"aliases":null},"github":{"committer_count":2,"uri":"https://github.com/monsooncommerce/gstats"},"confidence":0.534,"scores":[{"term":"django","score":0.393},{"term":"1.0","score":0.842}]}`
			var searchResult products.SoftwareEntity
			err := json.Unmarshal([]byte(searchResultJSON), &searchResult)
			Expect(err).NotTo(HaveOccurred())
			Expect(searchResult.Product.Up).To(Equal("alpha1"))
			Expect(searchResult.Scores).To(HaveLen(2))
			Expect(searchResult.Scores[0].Term).To(Equal("django"))
			Expect(searchResult.Scores[1].Term).To(Equal("1.0"))
			Expect(fmt.Sprintf("%.3f", searchResult.Scores[0].Score)).To(Equal("0.393"))
			Expect(fmt.Sprintf("%.3f", searchResult.Scores[1].Score)).To(Equal("0.842"))
			Expect(fmt.Sprintf("%.3f", searchResult.Confidence)).To(Equal("0.534"))
		})
		g.It("should marshal search results with scores", func() {
			product := products.Product{
				ID:         1234,
				Name:       "product name",
				Org:        "some org",
				Version:    "3.1.2",
				Title:      "mah title",
				ExternalID: "cpe:/a:djangoproject:django:1.0:alpha1",
			}
			github := products.Github{
				URI:            "https://github.com/some/repo",
				CommitterCount: 5,
			}
			scores := []products.ProductSearchScore{
				{Term: "foo", Score: 3.2},
			}
			searchResult := products.SoftwareEntity{
				Product:    &product,
				Github:     &github,
				Confidence: 3.6,
				Scores:     scores,
			}
			b, err := json.Marshal(searchResult)
			Expect(err).NotTo(HaveOccurred())
			s := string(b)
			Expect(s).To(MatchRegexp(`"name":\s*"product name"`))
			Expect(s).To(MatchRegexp(`"confidence":\s*3.6`))
			Expect(s).To(MatchRegexp(`"score":\s*3.2`))
			Expect(s).To(MatchRegexp(`"term":\s*"foo"`))
		})
		g.It("should unmarshal search results with packages", func() {
			searchResultJSON := `{"product":{"id":1234,"name":"product name","org":"some org","version":"3.1.2","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"mah title","references":null,"part":"","language":"","external_id":"cpe:/a:djangoproject:django:1.0:alpha1","source":null},"github":{"uri":"https://github.com/some/repo","committer_count":5},"package":{"name":"mahProject","version":"2.3.6","type":"pypi"},"confidence":3.6,"scores":[{"term":"foo","score":3.2}]}`
			var searchResult products.SoftwareEntity
			err := json.Unmarshal([]byte(searchResultJSON), &searchResult)
			Expect(err).NotTo(HaveOccurred())
			Expect(searchResult.Package).NotTo(BeNil())
			Expect(searchResult.Package.Name).To(Equal("mahProject"))
			Expect(searchResult.Package.Version).To(Equal("2.3.6"))
			Expect(searchResult.Package.Type).To(Equal("pypi"))
		})
		g.It("should marshal search results with packages", func() {
			product := products.Product{
				ID:         1234,
				Name:       "product name",
				Org:        "some org",
				Version:    "3.1.2",
				Title:      "mah title",
				ExternalID: "cpe:/a:djangoproject:django:1.0:alpha1",
			}
			github := products.Github{
				URI:            "https://github.com/some/repo",
				CommitterCount: 5,
			}
			pkg := products.Package{
				Name:    "mahProject",
				Version: "2.3.6",
				Type:    "pypi",
			}
			scores := []products.ProductSearchScore{
				{Term: "foo", Score: 3.2},
			}
			searchResult := products.SoftwareEntity{
				Product:    &product,
				Github:     &github,
				Package:    &pkg,
				Confidence: 3.6,
				Scores:     scores,
			}
			b, err := json.Marshal(searchResult)
			Expect(err).NotTo(HaveOccurred())
			s := string(b)
			Expect(s).To(MatchRegexp(`"name":\s*"product name"`))
			Expect(s).To(MatchRegexp(`"confidence":\s*3.6`))
			Expect(s).To(MatchRegexp(`"score":\s*3.2`))
			Expect(s).To(MatchRegexp(`"term":\s*"foo"`))
			Expect(s).To(MatchRegexp(`"package":\s*{`))
			Expect(s).To(MatchRegexp(`"name":\s*"mahProject"`))
		})

		g.It("should validate a good request", func() {
			search := products.ProductSearchQuery{
				SearchType:        "concatenated",
				SearchStrategy:    "strat",
				ProductIdentifier: "productIdentifier",
				Version:           "1.0.0",
				Terms:             []string{"foo"},
			}
			Expect(search.IsValid()).To(BeTrue())
			search.SearchType = "deconcatenated"
			Expect(search.IsValid()).To(BeTrue())
		})

		g.It("should validate a bad request", func() {
			search := products.ProductSearchQuery{
				SearchType:        "type",
				SearchStrategy:    "strat",
				ProductIdentifier: "productIdentifier",
				Version:           "1.0.0",
				Terms:             []string{"foo"},
			}
			Expect(search.IsValid()).To(BeFalse())

			search.SearchType = "concatenated"
			search.SearchStrategy = ""
			Expect(search.IsValid()).To(BeFalse())
		})
	})
}

const (
	SampleValidProduct         = `{"data":[{"id":84647,"name":"jdk","org":"oracle","version":"1.6.0","up":"update_71","edition":"","aliases":null,"created_at":"2017-02-13T20:02:42.600Z","updated_at":"2017-02-13T20:02:42.600Z","title":"Oracle JDK 1.6.0 Update 71","references":[{"April 2014 CPU":"http://www.oracle.com/technetwork/topics/security/cpuapr2014-1972952.html"}],"part":"/a","language":"","external_id":"cpe:/a:oracle:jdk:1.6.0:update_71","source":[{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-09T20:18:35.385Z","updated_at":"2017-02-13T20:12:05.342Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®)","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}]}]}`
	SampleValidRawProduct      = `[{"id":84647,"name":"jdk","org":"oracle","version":"1.6.0","up":"update_71","edition":"","aliases":null,"created_at":"2017-02-13T20:02:42.600Z","updated_at":"2017-02-13T20:02:42.600Z","title":"Oracle JDK 1.6.0 Update 71","references":[{"April 2014 CPU":"http://www.oracle.com/technetwork/topics/security/cpuapr2014-1972952.html"}],"part":"/a","language":"","external_id":"cpe:/a:oracle:jdk:1.6.0:update_71","source":[{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-09T20:18:35.385Z","updated_at":"2017-02-13T20:12:05.342Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®)","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}]}]`
	sampleBunsenSearchResponse = `{"data":[{"id":39862,"name":"less","org":"gnu","version":"-","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:-"},{"id":39863,"name":"less","org":"gnu","version":"358","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 358","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:358"},{"id":39864,"name":"less","org":"gnu","version":"381","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 381","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:381"},{"id":39865,"name":"less","org":"gnu","version":"382","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 382","references":[],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:382"},{"id":39866,"name":"less","org":"gnu","version":"471","up":"","edition":"","aliases":null,"created_at":"2017-02-13T20:02:36.794Z","updated_at":"2017-02-13T20:02:36.794Z","title":"GNU less 471","references":[{"Vendor Website":"http://www.gnu.org/software/less/"}],"part":"/a","language":"","external_id":"cpe:/a:gnu:less:471"}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","last_update":"2018-05-03T16:27:42.409Z","total_count":5,"limit":10,"offset":0},"links":{"self":"https://api.ionchannel.io/v1/product/search?user_query=less"}}`
	sameProductVersions        = `{"data":[{"id":0,"name":"","org":"","version":"13.0.0","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"12.0.1","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"12.0.0","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"12","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"11.0.4","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"11.0.3","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"11.0.2","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"11.0.1","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0},{"id":0,"name":"","org":"","version":"11.0.0","up":"","edition":"","aliases":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","title":"","references":null,"part":"","language":"","external_id":"","source":null,"confidence":0}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","total_count":9,"offset":0}}`
)
