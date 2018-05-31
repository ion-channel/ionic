package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestSearches(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Searches", func() {
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
		g.It("should perform a productidentifiers search", func() {
			server.AddPath("/v1/search/productidentifiers").
				SetMethods("GET").
				SetPayload([]byte(sampleValidProdIDSearch)).
				SetStatus(http.StatusOK)
			searchResults, err := client.GetSearchProductIdentifiers("prodIdent", "version", "vendor", "sometoken")
			Expect(err).NotTo(HaveOccurred())
			Expect(searchResults).To(HaveLen(2))
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Query.Get("productidentifiers")).To(Equal("prodIdent"))
			Expect(hitRecords[0].Query.Get("version")).To(Equal("version"))
			Expect(hitRecords[0].Query.Get("vendor")).To(Equal("vendor"))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer sometoken"))

			Expect(searchResults[0].Product.Name).To(Equal("jdk"))
			Expect(searchResults[1].Github.CommitterCount).To(Equal(uint(31337)))
		})
		g.It("should perform a repositories search", func() {
			server.AddPath("/v1/search/repositories").
				SetMethods("GET").
				SetPayload([]byte(sampleValidRositaRepoSearch)).
				SetStatus(http.StatusOK)
			searchResults, err := client.GetSearchRepositories("https://github.com/apache/ode", "sometoken")
			Expect(err).NotTo(HaveOccurred())
			Expect(searchResults).To(HaveLen(1))
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Query.Get("github")).To(Equal("https://github.com/apache/ode"))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer sometoken"))

			Expect(searchResults[0].Product.Name).To(Equal("ode"))
			Expect(searchResults[0].Product.Title).To(Equal("Apache Ode 1.1.1"))
			Expect(searchResults[0].Github.CommitterCount).To(Equal(uint(11)))
		})
	})
}

const (
	sampleValidProdIDSearch     = `{"data": [{"product": {"name": "jdk", "language": "", "title": "IBM Jdk 1.4.2", "created_at": "2017-02-13T20:02:38.33Z", "up": "", "updated_at": "2017-02-13T20:02:38.33Z", "edition": "", "version": "1.4.2", "references": [], "source": null, "org": "ibm", "part": "/a", "external_id": "cpe:/a:ibm:jdk:1.4.2", "id": 52994, "aliases": null}, "github": {"committer_count": 31337, "uri": "https://github.com/ibm/jdk"}}, {"product": {"name": "jdk", "language": "", "title": "IBM Jdk 5.0", "created_at": "2017-02-13T20:02:38.33Z", "up": "", "updated_at": "2017-02-13T20:02:38.33Z", "edition": "", "version": "5.0", "references": [], "source": null, "org": "ibm", "part": "/a", "external_id": "cpe:/a:ibm:jdk:5.0", "id": 52995, "aliases": null}, "github": {"committer_count": 31337, "uri": "https://github.com/ibm/jdk"}}], "meta": { "authors": [ "Ion Channel Dev Team" ], "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net", "offset": 0, "total_count": 2, "version": "v1"}}`
	sampleValidRositaRepoSearch = `{"meta": {"total_count": 2, "offset": 0, "version": "v1", "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net", "authors": ["Ion Channel Dev Team"]}, "data": [{"product": {"name": "ode", "language": "", "source": null, "created_at": "2018-03-27T22:26:14.991Z", "title": "Apache Ode 1.1.1", "up": "rc1", "updated_at": "2018-03-27T22:26:14.991Z", "edition": "", "part": "/a", "references": null, "version": "1.1.1", "org": "apache", "external_id": "cpe:/a:apache:ode:1.1.1:rc1", "id": 759066869, "aliases": null}, "github": {"committer_count": 11, "uri": "https://github.com/apache/ode"}}]}`
)
