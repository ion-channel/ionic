package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestAnalysis(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Analysis", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get an analysis", func() {
			server.AddPath("/v1/animal/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysis)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetAnalysis("f9bca953-80ac-46c4-b195-d37f3bc4f498", "ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(analysis.ID).To(Equal("f9bca953-80ac-46c4-b195-d37f3bc4f498"))
			Expect(analysis.Status).To(Equal("finished"))
			Expect(analysis.Type).To(Equal("git"))
			Expect(analysis.TriggerAuthor).To(Equal("Daniel Hess"))
			Expect(len(analysis.ScanSummaries)).To(Equal(2))
		})

		g.It("should get a raw analysis", func() {
			server.AddPath("/v1/animal/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysis)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetRawAnalysis("f9bca953-80ac-46c4-b195-d37f3bc4f498", "ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(string(analysis)).To(ContainSubstring("f9bca953-80ac-46c4-b195-d37f3bc4f498"))
			Expect(string(analysis)).To(ContainSubstring("finished"))
			Expect(string(analysis)).To(ContainSubstring("git"))
			Expect(string(analysis)).To(ContainSubstring("Daniel Hess"))
		})

		g.It("should get the latest analysis summary", func() {
			server.AddPath("/v1/animal/getLatestAnalysisSummary").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisSummary)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetLatestAnalysisSummary("ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(analysis.TeamID).To(Equal("cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc"))
			Expect(analysis.Passed).To(Equal(false))
		})

		g.It("should get the raw latest analysis summary", func() {
			server.AddPath("/v1/animal/getLatestAnalysisSummary").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisSummary)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetRawLatestAnalysisSummary("ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(string(analysis)).To(ContainSubstring("cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc"))
			Expect(string(analysis)).To(ContainSubstring("failed"))
		})
	})
}

const (
	SampleValidAnalysis        = `{"data":{"id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"33ef183d-4d37-4515-84c4-099ed0fb8db0","build_number":"293","name":"compliance analysis","text":null,"type":"git","source":"git@github.com:ion-channel/ionic.git","branch":"master","description":"","status":"finished","ruleset_id":"fcd09ba9-c939-4b51-a865-394cc8ddcffa","created_at":"2017-07-18T20:21:30.000Z","updated_at":"2017-07-18T20:21:30.316Z","duration":40209.8742020007,"trigger_hash":"aa8a66375adef765fe9eed5920cfa8352e4c4b70","trigger_text":"Merge pull request #220 from ion-channel/foobranch\n\nadding new coverage format","trigger_author":"Daniel Hess","scan_summaries":[{"id":"c2430d2a-5063-f360-e222-d0dd96b90e22","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"33ef183d-4d37-4515-84c4-099ed0fb8db0","description":"some description","name":"license","summary":"Finished license scan for Statler, found apache-2.0 license.","created_at":"2017-07-18T20:20:53.249Z","updated_at":"2017-07-18T20:20:53.249Z","results":{"license":{"license":{"name":"LICENSE.md","type":[{"name":"apache-2.0"}]}}},"duration":674.671019000016},{"id":"d8214cb1-65fa-2d29-226f-d76728b4fb88","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"33ef183d-4d37-4515-84c4-099ed0fb8db0","description":"some description","name":"about_yml","summary":"Finished about_yml scan for Statler, valid .about.yml found.","created_at":"2017-07-18T20:20:53.566Z","updated_at":"2017-07-18T20:20:53.566Z","results":{"about_yml":{"message":"","valid":true,"content":"---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: Statler\n\n# This is the display name of your project. (required)\nfull_name: Statler API\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: This is the api Statler\n\ntype: app\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: alpha\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  statler:\n    name: CC0\n    url: https://github.com/ion-channel/statler/blob/master/LICENSE.md\n\nteam:\n- github: kellyp\n  role: lead\n\ncontact:\n- url: mailto:info@ionchannel.io\n  text: Ion Channel Info Line\n"}},"duration":1130.52755299941}]}}`
	SampleValidAnalysisSummary = `{"data":{"id":"a07b82b0-9742-447c-a277-cee0e56abf7f","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"fc42d773-3764-4a78-b3a3-f09bf7241a9c","build_number":"189","name":"a-ionmock","text":null,"type":"git","source":"https://github.com/matthewkmayer/ionmockjavaapp","branch":"master","description":"","status":"failed","ruleset_id":"ec4b43e6-ecfc-42c8-b58c-8a47eab0cc68","created_at":"2017-11-02T12:56:06.000Z","updated_at":"2017-11-02T12:56:06.523Z","duration":50076.0678370007,"trigger_hash":"6a2b494c22e3aeed1fc22fbc549b243b57a7d304","trigger_text":"Merge pull request #1 from ion-channel/Adding-Unit-Tests\n\nAdding jacoco plugin, adding new methods and tests to up code coverage","trigger_author":"Kit Plummer"}}`
)
