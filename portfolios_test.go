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

	g.Describe("Portfolios", func() {
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

		g.It("should get raw vulnerability list", func() {
			server.AddPath("/v1/animal/getVulnerabilityList").
				SetMethods("POST").
				SetPayload([]byte(SampleVulnList)).
				SetStatus(http.StatusOK)

			vl, err := client.GetRawVulnerabilityList([]string{"1", "2"}, "somelist", "5", "sometoken")
			Expect(err).To(BeNil())
			Expect(string(vl)).To(Equal("{\"cve_list\":[{\"title\":\"cve1\",\"projects_affected\":3,\"product\":\"someproduct2\",\"rating\":8.8,\"system\":\"cvssv3\"}]}"))
		})

		g.It("should get raw vulnerability metrics", func() {
			server.AddPath("/v1/animal/getScanMetrics").
				SetMethods("POST").
				SetPayload([]byte(SampleVulnMetrics)).
				SetStatus(http.StatusOK)

			vm, err := client.GetRawVulnerabilityMetrics([]string{"1", "2"}, "somemetric", "sometoken")
			Expect(err).To(BeNil())
			Expect(string(vm)).To(Equal("{\"line_graph\":{\"title\":\"vulnerabilities over time\",\"lines\":[{\"domain\":\"date\",\"range\":\"count\",\"legend\":\"vulnerabilities\",\"points\":{\"2019-10-08\":9}},{\"domain\":\"date\",\"range\":\"count\",\"legend\":\"projects\",\"points\":{\"2019-10-08\":3}}]}}"))
		})
	})
}

const (
	SampleVulnStats   = `{"data":{"total_vulnerabilities":4,"unique_vulnerabilities":2,"most_frequent_vulnerability":"somecve"}}`
	SampleVulnList    = `{"data":{"cve_list":[{"title":"cve1","projects_affected":3,"product":"someproduct2","rating":8.8,"system":"cvssv3"}]}}`
	SampleVulnMetrics = `{"data":{"line_graph":{"title":"vulnerabilities over time","lines":[{"domain":"date","range":"count","legend":"vulnerabilities","points":{"2019-10-08":9}},{"domain":"date","range":"count","legend":"projects","points":{"2019-10-08":3}}]}}}`
)
