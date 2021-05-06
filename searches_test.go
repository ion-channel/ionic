package ionic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestSearches(t *testing.T) {
	g := goblin.Goblin(t)
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
		g.It("should perform a search", func() {
			server.AddPath("/v1/search").
				SetMethods("GET").
				SetPayload([]byte(sampleValidProdIDSearch)).
				SetStatus(http.StatusOK)
			searchResults, meta, err := client.GetSearch("query", "products", "sometoken")
			Expect(err).NotTo(HaveOccurred())
			Expect(searchResults).To(HaveLen(2))
			Expect(meta.TotalCount).To(Equal(2))
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Query.Get("q")).To(Equal("query"))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer sometoken"))

			Expect(searchResults[0].Name).To(Equal("jdk"))
			Expect(searchResults[0].Confidence).To(Equal(float32(1.0)))
			Expect(searchResults[0].Org).To(Equal("ibm"))
		})
	})
}

const (
	sampleValidProdIDSearch = `{"data": [{"name": "jdk", "language": "", "title": "IBM Jdk 1.4.2", "created_at": "2017-02-13T20:02:38.33Z", "up": "", "updated_at": "2017-02-13T20:02:38.33Z", "edition": "", "version": "1.4.2", "references": [], "source": null, "org": "ibm", "part": "/a", "external_id": "cpe:/a:ibm:jdk:1.4.2", "id": 52994, "aliases": null, "confidence": 1.0}, {"name": "jdk", "language": "", "title": "IBM Jdk 5.0", "created_at": "2017-02-13T20:02:38.33Z", "up": "", "updated_at": "2017-02-13T20:02:38.33Z", "edition": "", "version": "5.0", "references": [], "source": null, "org": "ibm", "part": "/a", "external_id": "cpe:/a:ibm:jdk:5.0", "id": 52995, "aliases": null}], "meta": { "authors": [ "Ion Channel Dev Team" ], "copyright": "Copyright 2018 Selection Pressure LLC www.selectpress.net", "offset": 0, "total_count": 2, "version": "v1"}}`
)
