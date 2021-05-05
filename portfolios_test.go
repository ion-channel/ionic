package ionic

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestPortfolios(t *testing.T) {
	g := goblin.Goblin(t)
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

		g.It("should return a passing/failing status summary", func() {
			server.AddPath("/v1/ruleset/getStatuses").
				SetMethods("POST").
				SetPayload([]byte(SamplePassFailSummary)).
				SetStatus(http.StatusOK)

			vss, err := client.GetPortfolioPassFailSummary([]string{"1", "2"}, "sometoken")
			Expect(err).To(BeNil())
			Expect(vss.PassingProjects).To(Equal(0))
			Expect(vss.FailingProjects).To(Equal(4))
		})

		g.It("should return a started and errored summary", func() {
			server.AddPath("/v1/scanner/getStatuses").
				SetMethods("POST").
				SetPayload([]byte(SampleStartedEndedSummary)).
				SetStatus(http.StatusOK)

			s, err := client.GetPortfolioStartedErroredSummary([]string{"1", "2"}, "sometoken")
			Expect(err).To(BeNil())
			Expect(s.AnalyzingProjects).To(Equal(2))
			Expect(s.ErroredProjects).To(Equal(6))
			Expect(s.FinishedProjects).To(Equal(3))
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

		g.It("should return dependency stats", func() {
			server.AddPath("/v1/animal/getDependencyStats").
				SetMethods("POST").
				SetPayload([]byte(SampleDependencyStats)).
				SetStatus(http.StatusOK)

			r, err := client.GetDependencyStats([]string{"aprojectid"}, "sometoken")
			Expect(err).To(BeNil())
			Expect(r.DirectDependencies).To(Equal(44))
			Expect(r.TransitiveDependencies).To(Equal(33))
			Expect(r.OutdatedDependencies).To(Equal(22))
			Expect(r.NoVersionSpecified).To(Equal(11))
		})

		g.It("should get raw dependency list", func() {
			server.AddPath("/v1/animal/getDependencyList").
				SetMethods("POST").
				SetPayload([]byte(SampleDependencyList)).
				SetStatus(http.StatusOK)

			vl, err := client.GetRawDependencyList([]string{"1", "2"}, "somelist", "5", "sometoken")
			Expect(err).To(BeNil())
			Expect(string(vl)).To(Equal("{\"dependency_list\":[{\"name\":\"name1\",\"org\":\"org1\",\"version\":\"someversion1\",\"package\":\"package1\",\"type\":\"type1\",\"latest_version\":\"latestversion1\",\"scope\":\"scope1\",\"requirement\":\"requirement1\",\"file\":\"file1\",\"projects_count\":2}]}"))
		})

		g.It("should get status history", func() {
			server.AddPath("/v1/ruleset/getStatusesHistory").
				SetMethods("POST").
				SetPayload([]byte(SampleStatusHistory)).
				SetStatus(http.StatusOK)

			sh, err := client.GetProjectsStatusHistory([]string{"1", "2"}, "sometoken")
			Expect(err).To(BeNil())
			t, _ := time.Parse(time.RFC3339Nano, "2020-07-31T10:54:51.725241-07:00")
			Expect(sh[0].Status).To(Equal("pass"))
			Expect(sh[0].Count).To(Equal(4))
			Expect(sh[0].FirstCreatedAt).To(Equal(t))
			t, _ = time.Parse(time.RFC3339Nano, "2020-07-31T10:58:51.725241-07:00")
			Expect(sh[1].Status).To(Equal("fail"))
			Expect(sh[1].Count).To(Equal(1))
			Expect(sh[1].FirstCreatedAt).To(Equal(t))
			t, _ = time.Parse(time.RFC3339Nano, "2020-07-31T10:59:51.725241-07:00")
			Expect(sh[2].Status).To(Equal("pass"))
			Expect(sh[2].Count).To(Equal(1))
			Expect(sh[2].FirstCreatedAt).To(Equal(t))
		})

		g.It("should return a mttr of a project", func() {
			server.AddPath("/v1/report/getMttr").
				SetMethods("GET").
				SetPayload([]byte(SampleGetMttr)).
				SetStatus(http.StatusOK)

			r, err := client.GetMttr("team_id", "project_id", "sometoken")
			Expect(err).To(BeNil())
			Expect(r).NotTo(BeNil())
			Expect(r.Mttr).To(Equal("1 day"))
			Expect(r.UnresolvedIncident).To(Equal(false))
			Expect(r.FailedMttrIncidents).To(Equal(1))
			Expect(r.ProjectCount).To(Equal(1))
		})

		g.It("should return a list of projects by dependency", func() {
			server.AddPath("/v1/animal/getProjectIdsByDependency").
				SetMethods("GET").
				SetPayload([]byte(SampleGetProjectsByDep)).
				SetStatus(http.StatusOK)

			r, err := client.GetProjectIdsByDependency("someteam", "activesupport", "rubyonrails", "5.2.4.3", "sometoken")
			Expect(err).To(BeNil())
			Expect(r.TeamID).To(Equal("someteam"))
			Expect(r.Name).To(Equal("activesupport"))
			Expect(r.Org).To(Equal("rubyonrails"))
			Expect(r.Version).To(Equal("5.2.4.3"))
			Expect(r.ProjectIDs[0]).To(Equal("335e6e38-1c35-48e0-ac05-7e54a0950acb"))
			Expect(r.ProjectIDs[1]).To(Equal("bc169c32-5d3c-4685-ae7e-8efe3a47c4fa"))
		})
	})
}

const (
	SampleVulnStats           = `{"data":{"total_vulnerabilities":4,"unique_vulnerabilities":2,"most_frequent_vulnerability":"somecve"}}`
	SampleVulnList            = `{"data":{"cve_list":[{"title":"cve1","projects_affected":3,"product":"someproduct2","rating":8.8,"system":"cvssv3"}]}}`
	SampleVulnMetrics         = `{"data":{"line_graph":{"title":"vulnerabilities over time","lines":[{"domain":"date","range":"count","legend":"vulnerabilities","points":{"2019-10-08":9}},{"domain":"date","range":"count","legend":"projects","points":{"2019-10-08":3}}]}}}`
	SamplePassFailSummary     = `{"data":{"passing_projects":0,"failing_projects":4}}`
	SampleStartedEndedSummary = `{"data":{"analyzing_projects":2,"errored_projects":6,"finished_projects":3}}`
	SampleAffectedProjectIds  = `{"data":[{"id":"1984b037-71f5-4bc2-84f0-5baf37a25fa5","name":"","version":"","vulnerabilities":15},{"id":"bc169c32-5d3c-4685-ae7e-8efe3a47c4fa","name":"","version":"","vulnerabilities":1}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","total_count":0,"offset":0}}`
	SampleAffectedProjectInfo = `{"data":[{"id":"1984b037-71f5-4bc2-84f0-5baf37a25fa5","name":"someName1","version":"someVersion1","vulnerabilities":0},{"id":"bc169c32-5d3c-4685-ae7e-8efe3a47c4fa","name":"someName2","version":"someVersion2","vulnerabilities":0}],"meta":{"copyright":"Copyright 2018 Selection Pressure LLC www.selectpress.net","authors":["Ion Channel Dev Team"],"version":"v1","total_count":0,"offset":0}}`
	SampleDependencyStats     = `{"data":{"direct_dependencies":44,"transitive_dependencies":33,"outdated_dependencies":22,"no_vesion_dependencies":11}}`
	SampleDependencyList      = `{"data":{"dependency_list":[{"name":"name1","org":"org1","version":"someversion1","package":"package1","type":"type1","latest_version":"latestversion1","scope":"scope1","requirement":"requirement1","file":"file1","projects_count":2}]}}`
	SampleStatusHistory       = `{"data":[{"status":"pass","count":4,"first_created_at":"2020-07-31T10:54:51.725241-07:00"},{"status":"fail","count":1,"first_created_at":"2020-07-31T10:58:51.725241-07:00"},{"status":"pass","count":1,"first_created_at":"2020-07-31T10:59:51.725241-07:00"}],"meta":{"total_count":3,"offset":0,"last_update":"0001-01-01T00:00:00Z"}}`
	SampleGetMttr             = `{"data":{"mttr":"1 day","unresolved_incident":false,"time_in_current_status":"1.83 days","failed_mttr_incidents":1,"project_count":1,"data":[{"status":"pass","count":2,"first_created_at":"2020-08-01T22:54:34.543703Z"},{"status":"fail","count":1,"first_created_at":"2020-08-02T22:54:34.579672Z"},{"status":"pass","count":1,"first_created_at":"2020-08-03T22:54:34.608181Z"}]},"meta":{"total_count":1,"offset":0,"last_update":"2020-08-05T18:48:10.6778402Z"}}`
	SampleGetProjectsByDep    = `{"data":{"team_id":"someteam","name":"activesupport","org":"rubyonrails","version":"5.2.4.3","project_ids":["335e6e38-1c35-48e0-ac05-7e54a0950acb", "bc169c32-5d3c-4685-ae7e-8efe3a47c4fa"]},"meta":{"total_count":0,"offset":0}}`
)
