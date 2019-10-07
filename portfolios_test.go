package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestPortfolios(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Reports", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get vulnerability statistics", func() {
			server.AddPath("/v1/animal/getVulnerabilityStats").
				SetMethods("POST").
				SetPayload([]byte(SampleVulnStats)).
				SetStatus(http.StatusOK)

			vs, err := client.GetVulnerabilityStats([]string{"1", "2"}, "sometoken")
			Expect(err).To(BeNil())

			Expect(vs.TotalVulnerabilities).NotTo(BeNil())
			Expect(vs.UniqueVulnerabilities).NotTo(BeNil())
			Expect(vs.MostFrequentVulnerability).NotTo(BeNil())

			Expect(vs.TotalVulnerabilities).To(Equal(4))
			Expect(vs.UniqueVulnerabilities).To(Equal(2))
			Expect(vs.MostFrequentVulnerability).To(Equal("somecve"))
		})
	})
}

const (
	SampleVulnStats = `{"data":{"total_vulnerabilities":4,"unique_vulnerabilities":2,"most_frequent_vulnerability":"somecve"}}`
)
