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

		g.It("should return a status summary", func() {
			server.AddPath("/v1/ruleset/getPortfolioSummary").
				SetMethods("POST").
				SetPayload([]byte(SampleStatusSummary)).
				SetStatus(http.StatusOK)

			vss, err := client.GetPortfolioStatusSummary([]string{"1", "2"}, "sometoken")
			Expect(err).To(BeNil())
			Expect(vss.PassingProjects).To(Equal(0))
			Expect(vss.FailingProjects).To(Equal(4))
			Expect(vss.ErroredProjects).To(Equal(1))
			Expect(vss.PendingProjects).To(Equal(14))
		})

		g.It("should return a started and errored summary", func() {
			server.AddPath("/v1/scanner/getStatuses").
				SetMethods("POST").
				SetPayload([]byte(SampleStartedEndedSummary)).
				SetStatus(http.StatusOK)

			s, err := client.GetProjectStatusForStartedAndErrored([]string{"1", "2"}, "sometoken")
			Expect(err).To(BeNil())
			Expect(s.StartedProjects).To(Equal(2))
			Expect(s.ErroredProjects).To(Equal(6))
		})

		g.It("should return a list of affected projects", func() {
			server.AddPath("/v1/animal/getAffectedProjectIds").
				SetMethods("GET").
				SetPayload([]byte(SampleAffectedProjectIds)).
				SetStatus(http.StatusOK)

			r, err := client.GetPortfolioAffectedProjects("team_id", "external_id", "sometoken")
			Expect(err).To(BeNil())
			Expect(len(r)).To(Equal(2))
			Expect(r[0].ID).To(Equal("1984b037-71f5-4bc2-84f0-5baf37a25fa5"))
			Expect(r[1].Vulnerabilities).To(Equal(1))
		})

		g.It("should return a list of affected projects info", func() {
			server.AddPath("/v1/project/getAffectedProjectsInfo").
				SetMethods("POST").
				SetPayload([]byte(SampleAffectedProjectInfo)).
				SetStatus(http.StatusOK)

			r, err := client.GetPortfolioAffectedProjectsInfo([]string{"aprojectid"}, "sometoken")
			Expect(err).To(BeNil())
			Expect(len(r)).To(Equal(2))
			Expect(r[0].ID).To(Equal("1984b037-71f5-4bc2-84f0-5baf37a25fa5"))
			Expect(r[0].Name).To(Equal("someName1"))
			Expect(r[1].Version).To(Equal("someVersion2"))
		})
	})
}

const (
	SampleVulnStats           = `{"data":{"total_vulnerabilities":4,"unique_vulnerabilities":2,"most_frequent_vulnerability":"somecve"}}`
	SampleVulnList            = `{"data":{"cve_list":[{"title":"cve1","projects_affected":3,"product":"someproduct2","rating":8.8,"system":"cvssv3"}]}}`
	SampleVulnMetrics         = `{"data":{"line_graph":{"title":"vulnerabilities over time","lines":[{"domain":"date","range":"count","legend":"vulnerabilities","points":{"2019-10-08":9}},{"domain":"date","range":"count","legend":"projects","points":{"2019-10-08":3}}]}}}`
	SampleStatusSummary       = `{"data":{"passing_projects":0,"failing_projects":4,"errored_projects":1,"pending_projects":14}}`
	SampleStartedEndedSummary = `{"data":{"started_projects":2,"errored_projects":6}}`
	SampleAffectedProjectIds  = `{"data":[{"id":"1984b037-71f5-4bc2-84f0-5baf37a25fa5","name":"","version":"","vulnerabilities":15},{"id":"bc169c32-5d3c-4685-ae7e-8efe3a47c4fa","name":"","version":"","vulnerabilities":1}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","total_count":0,"offset":0}}`
	SampleAffectedProjectInfo = `{"data":[{"id":"1984b037-71f5-4bc2-84f0-5baf37a25fa5","name":"someName1","version":"someVersion1","vulnerabilities":0},{"id":"bc169c32-5d3c-4685-ae7e-8efe3a47c4fa","name":"someName2","version":"someVersion2","vulnerabilities":0}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","total_count":0,"offset":0}}`
)
