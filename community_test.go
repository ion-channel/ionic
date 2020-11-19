package ionic

import (
	"fmt"
	"net/http"
	"testing"

	goblin "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestCommunity(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Community", func() {
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
		g.It("should get a repo", func() {
			server.AddPath("/v1/repo/getRepo").
				SetMethods("GET").
				SetPayload([]byte(sampleValidGetRepoResponse)).
				SetStatus(http.StatusOK)
			searchResults, err := client.GetRepo("monsooncommerce/gstats", "someToken")
			Expect(err).NotTo(HaveOccurred())

			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer someToken"))
			Expect(hitRecords[0].Query.Get("repo")).To(Equal("monsooncommerce/gstats"))

			Expect(searchResults).NotTo(BeNil())
			Expect(searchResults.Name).To(Equal("monsooncommerce/gstats"))
			Expect(searchResults.URL).To(Equal("https://github.com/monsooncommerce/gstats"))
			Expect(searchResults.Committers).To(Equal(2))
		})
		g.It("should search repos", func() {
			server.AddPath("/v1/repo/search").
				SetMethods("GET").
				SetPayload([]byte(sampleValidSearchRepoResponse)).
				SetStatus(http.StatusOK)
			searchResults, err := client.SearchRepo("monsooncommerce", "blaToken")
			Expect(err).NotTo(HaveOccurred())

			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer blaToken"))
			Expect(hitRecords[0].Query.Get("q")).To(Equal("monsooncommerce"))

			Expect(searchResults).NotTo(BeNil())
			Expect(searchResults).To(HaveLen(1))
			Expect(searchResults[0].Name).To(Equal("monsooncommerce/gstats"))
			Expect(searchResults[0].URL).To(Equal("https://github.com/monsooncommerce/gstats"))
			Expect(searchResults[0].Committers).To(Equal(2))
			Expect(searchResults[0].Confidence).To(Equal(0.99999))

		})
	})
}

const (
	sampleValidGetRepoResponse    = `{"data":{"name":"monsooncommerce/gstats","url":"https://github.com/monsooncommerce/gstats","committers":2},"meta":{"copyright":"","authors":null,"version":"","last_update":"0001-01-01T00:00:00Z","total_count":1}}`
	sampleValidSearchRepoResponse = `{"meta": {"total_count": 1, "version": "", "last_update": "0001-01-01T00:00:00Z", "copyright": "", "authors": null}, "data": [{"confidence":0.99999, "url": "https://github.com/monsooncommerce/gstats", "committers": 2, "name": "monsooncommerce/gstats"}]}`
)
