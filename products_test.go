package ionic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/products"
	. "github.com/onsi/gomega"
)

func TestProducts(t *testing.T) {
	g := Goblin(t)
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
			products, err := client.GetProductSearch("less", "mahVersion", "mahVendor", "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer someapikey"))
			Expect(hitRecord.Query.Get("product_identifier")).To(Equal("less"))
			Expect(hitRecord.Query.Get("version")).To(Equal("mahVersion"))
			Expect(hitRecord.Query.Get("vendor")).To(Equal("mahVendor"))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
		g.It("should omit version and vendor from product search when it is not given", func() {
			server.AddPath("/v1/product/search").
				SetMethods("GET").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			products, err := client.GetProductSearch("less", "", "", "someapikey")
			Expect(err).To(BeNil())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			hitRecord := hitRecords[0]
			Expect(hitRecord.Header.Get("Authorization")).To(Equal("Bearer someapikey"))
			Expect(hitRecord.Query.Get("product_identifier")).To(Equal("less"))
			Expect(hitRecord.Query.Get("version")).To(Equal(""))
			Expect(hitRecord.Query.Get("vendor")).To(Equal(""))
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
		})
		g.It("should unmarshal search results with scores", func() {
			searchResultJSON := `{"product":{"name":"django","language":"","source":null,"created_at":"2017-02-13T20:02:35.667Z","title":"Django Project Django 1.0-alpha-1","up":"alpha1","updated_at":"2017-02-13T20:02:35.667Z","edition":"","part":"/a","references":[],"version":"1.0","org":"djangoproject","external_id":"cpe:/a:djangoproject:django:1.0:alpha1","id":30955,"aliases":null},"github":{"committer_count":2,"uri":"https://github.com/monsooncommerce/gstats"},"mean_score":0.534,"scores":[{"term":"django","score":0.393},{"term":"1.0","score":0.842}]}`
			var searchResult products.ProductSearchResult
			err := json.Unmarshal([]byte(searchResultJSON), &searchResult)
			Expect(err).NotTo(HaveOccurred())
			Expect(searchResult.Product.Up).To(Equal("alpha1"))
			Expect(searchResult.Scores).To(HaveLen(2))
			Expect(searchResult.Scores[0].Term).To(Equal("django"))
			Expect(searchResult.Scores[1].Term).To(Equal("1.0"))
			Expect(fmt.Sprintf("%.3f", searchResult.Scores[0].Score)).To(Equal("0.393"))
			Expect(fmt.Sprintf("%.3f", searchResult.Scores[1].Score)).To(Equal("0.842"))
			Expect(fmt.Sprintf("%.3f", searchResult.MeanScore)).To(Equal("0.534"))
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
			searchResult := products.ProductSearchResult{
				Product:   product,
				Github:    github,
				MeanScore: 3.6,
				Scores:    scores,
			}
			b, err := json.Marshal(searchResult)
			Expect(err).NotTo(HaveOccurred())
			s := string(b)
			Expect(s).To(MatchRegexp(`"name":\s*"product name"`))
			Expect(s).To(MatchRegexp(`"mean_score":\s*3.6`))
			Expect(s).To(MatchRegexp(`"score":\s*3.2`))
			Expect(s).To(MatchRegexp(`"term":\s*"foo"`))
		})
		g.It("should post a product request", func() {
			server.AddPath("/v1/product/search").
				SetMethods("POST").
				SetPayload([]byte(sampleBunsenSearchResponse)).
				SetStatus(http.StatusOK)
			searchInput := products.ProductSearchQuery{
				SearchType:        "concatenated",
				SearchStrategy:    "searchStrategy",
				ProductIdentifier: "productIdentifier",
				Version:           "version",
				Vendor:            "vendor",
				Terms:             []string{"term01", "term02"},
			}
			products, err := client.ProductSearch(searchInput, "token")
			Expect(err).NotTo(HaveOccurred())
			Expect(products).To(HaveLen(5))
			Expect(products[0].ID).To(Equal(39862))
			Expect(server.HitRecords()).To(HaveLen(1))
			hr := server.HitRecords()[0]
			Expect(hr.Query).To(HaveLen(0))
			Expect(hr.Header.Get("Authorization")).To(Equal("Bearer token"))
			Expect(string(hr.Body)).To(Equal(`{"search_type":"concatenated","search_strategy":"searchStrategy","product_identifier":"productIdentifier","version":"version","vendor":"vendor","terms":["term01","term02"]}`))
		})
		g.It("should post a maven request", func() {
			server.AddPath("/v1/dependency/search").
				SetMethods("POST").
				SetPayload([]byte(sampleMulchResponse)).
				SetStatus(http.StatusOK)
			results, err := client.MavenSearch("testartifact001", "testgroup001", "token")
			Expect(err).NotTo(HaveOccurred())
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(len(hitRecords[0].Body)).To(BeNumerically(">", 0))
			Expect(hitRecords[0].Body).To(ContainSubstring("testartifact001"))
			Expect(hitRecords[0].Body).To(ContainSubstring("testgroup001"))
			Expect(results).NotTo(BeNil())
			Expect(results).To(HaveLen(1))
			Expect(results[0]).NotTo(BeNil())
			Expect(results[0].ArtifactID).To(Equal("testartifact001"))
			Expect(results[0].GroupID).To(Equal("testgroup001"))
			Expect(results[0].Metadata).NotTo(BeNil())
			Expect(results[0].Metadata.GroupID).To(Equal("testgroup001"))
			Expect(results[0].Metadata.ArtifactID).To(Equal("testartifact001"))
			Expect(results[0].Metadata.Version).To(Equal("1.0"))
			Expect(results[0].Metadata.Versions).To(HaveLen(20))
			Expect(results[0].Metadata.Versions[12]).To(Equal("1.2-rc1"))
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
	sampleMulchResponse        = `{"meta":{"copyright":"Copyright 2017 - Ion Channel Corp (ionchannel.io)","authors":["tlpinney","Matthew Mayer"],"version":"v1","total_count":1},"links":{"self":"http://localhost:9000/v1/dependency/search"},"timestamps":{"created":"2018-08-29T19:23:16.792+00:00","updated":"2018-08-29T19:23:16.792+00:00"},"data":[{"id":1,"group_id":"testgroup001","artifact_id":"testartifact001","metadata":"\u003cmetadata\u003e\n  \u003cgroupId\u003etestgroup001\u003c/groupId\u003e\n  \u003cartifactId\u003etestartifact001\u003c/artifactId\u003e\n  \u003cversion\u003e1.0\u003c/version\u003e\n  \u003cversioning\u003e\n    \u003cversions\u003e\n      \u003cversion\u003e1.0\u003c/version\u003e\n      \u003cversion\u003e1.0-m4\u003c/version\u003e\n      \u003cversion\u003e1.0-rc1\u003c/version\u003e\n      \u003cversion\u003e1.1\u003c/version\u003e\n      \u003cversion\u003e1.1-rc1\u003c/version\u003e\n      \u003cversion\u003e1.1-rc2\u003c/version\u003e\n      \u003cversion\u003e1.1.1\u003c/version\u003e\n      \u003cversion\u003e1.1.2\u003c/version\u003e\n      \u003cversion\u003e1.1.3\u003c/version\u003e\n      \u003cversion\u003e1.1.4\u003c/version\u003e\n      \u003cversion\u003e1.1.5\u003c/version\u003e\n      \u003cversion\u003e1.2\u003c/version\u003e\n      \u003cversion\u003e1.2-rc1\u003c/version\u003e\n      \u003cversion\u003e1.2-rc2\u003c/version\u003e\n      \u003cversion\u003e1.2.1\u003c/version\u003e\n      \u003cversion\u003e1.2.2\u003c/version\u003e\n      \u003cversion\u003e1.2.3\u003c/version\u003e\n      \u003cversion\u003e1.2.4\u003c/version\u003e\n      \u003cversion\u003e1.2.5\u003c/version\u003e\n      \u003cversion\u003e1.2.6\u003c/version\u003e\n    \u003c/versions\u003e\n  \u003c/versioning\u003e\n\u003c/metadata\u003e","created_at":"2018-08-29T19:15:33.045Z","updated_at":"2018-08-29T19:15:33.045Z"}]}`
)
