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
		var server *bogus.Bogus
		var h, p string
		var client *IonClient

		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})

		g.It("should get an analysis", func() {
			server.AddPath("/v1/animal/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysis)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetAnalysis("f9bca953-80ac-46c4-b195-d37f3bc4f498", "ateamid", "aprojectid", "sometoken")
			Expect(err).To(BeNil())
			Expect(analysis.ID).To(Equal("f9bca953-80ac-46c4-b195-d37f3bc4f498"))
			Expect(analysis.Status).To(Equal("finished"))
			Expect(analysis.Type).To(Equal("git"))
			Expect(analysis.TriggerAuthor).To(Equal("Daniel Hess"))
			Expect(len(analysis.ScanSummaries)).To(Equal(2))
		})

		g.It("should return an error when backing service is unavailable", func() {
			server.AddPath("/v1/animal/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte("Fake 503 error as if from ALB")).
				SetStatus(http.StatusServiceUnavailable)
			_, err := client.GetAnalysis("f9bca953-80ac-46c4-b195-d37f3bc4f498", "ateamid", "aprojectid", "sometoken")
			Expect(err).To(HaveOccurred())
		})

		g.It("should get the latest public analysis", func() {
			server.AddPath("/v1/animal/getLatestPublicAnalysisSummary").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysis)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetLatestPublicAnalysis("33ef183d-4d37-4515-84c4-099ed0fb8db0", "master")
			Expect(err).To(BeNil())
			Expect(analysis.ProjectID).To(Equal("33ef183d-4d37-4515-84c4-099ed0fb8db0"))
			Expect(analysis.Status).To(Equal("finished"))
			Expect(analysis.Type).To(Equal("git"))
			Expect(analysis.TriggerAuthor).To(Equal("Daniel Hess"))
			Expect(analysis.Branch).To(Equal("master"))
			Expect(len(analysis.ScanSummaries)).To(Equal(2))
			Expect(server.Hits()).To(Equal(1))
		})

		g.It("should get a public analysis", func() {
			server.AddPath("/v1/animal/getPublicAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysis)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetPublicAnalysis("f9bca953-80ac-46c4-b195-d37f3bc4f498")
			Expect(err).To(BeNil())
			Expect(analysis.ID).To(Equal("f9bca953-80ac-46c4-b195-d37f3bc4f498"))
			Expect(analysis.Status).To(Equal("finished"))
			Expect(analysis.Type).To(Equal("git"))
			Expect(analysis.TriggerAuthor).To(Equal("Daniel Hess"))
			Expect(len(analysis.ScanSummaries)).To(Equal(2))
			Expect(server.Hits()).To(Equal(1))
		})

		g.It("should get a raw analysis", func() {
			server.AddPath("/v1/animal/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysis)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetRawAnalysis("f9bca953-80ac-46c4-b195-d37f3bc4f498", "ateamid", "aprojectid", "sometoken")
			Expect(err).To(BeNil())
			Expect(string(analysis)).To(ContainSubstring("f9bca953-80ac-46c4-b195-d37f3bc4f498"))
			Expect(string(analysis)).To(ContainSubstring("finished"))
			Expect(string(analysis)).To(ContainSubstring("git"))
			Expect(string(analysis)).To(ContainSubstring("Daniel Hess"))
		})

		g.It("should error when unable to contact sister-service", func() {
			server.AddPath("/v1/animal/getLatestAnalysisSummary").
				SetMethods("GET").
				SetPayload([]byte("503 error as if from ALB")).
				SetStatus(http.StatusServiceUnavailable)

			_, err := client.GetLatestAnalysisSummary("ateamid", "aprojectid", "sometoken")
			Expect(err).To(HaveOccurred())
		})

		g.It("should get the latest analysis summary", func() {
			server.AddPath("/v1/animal/getLatestAnalysisSummary").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisSummary)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetLatestAnalysisSummary("ateamid", "aprojectid", "sometoken")
			Expect(err).To(BeNil())
			Expect(analysis.TeamID).To(Equal("cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc"))
			Expect(analysis.Passed).To(Equal(false))
		})

		g.It("should get the raw latest analysis summary", func() {
			server.AddPath("/v1/animal/getLatestAnalysisSummary").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisSummary)).
				SetStatus(http.StatusOK)

			analysis, err := client.GetRawLatestAnalysisSummary("ateamid", "aprojectid", "sometoken")
			Expect(err).To(BeNil())
			Expect(string(analysis)).To(ContainSubstring("cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc"))
			Expect(string(analysis)).To(ContainSubstring("failed"))
		})
	})

	g.Describe("Analyses", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get analyses", func() {
			server.AddPath("/v1/animal/getAnalyses").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalyses)).
				SetStatus(http.StatusOK)

			analyses, err := client.GetAnalyses("ateamid", "aprojectid", "sometoken", nil)
			Expect(err).To(BeNil())
			Expect(len(analyses)).To(Equal(2))
			Expect(analyses[0].ID).To(Equal("a5f7409b-d976-403b-af86-a9976fea901a"))
			Expect(analyses[0].Status).To(Equal("finished"))
			Expect(analyses[0].Type).To(Equal("git"))
			Expect(analyses[0].TriggerAuthor).To(Equal("Dan"))
		})

		g.It("should get a raw analysis", func() {
			server.AddPath("/v1/animal/getAnalyses").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalyses)).
				SetStatus(http.StatusOK)

			analyses, err := client.GetRawAnalyses("ateamid", "aprojectid", "sometoken", nil)
			Expect(err).To(BeNil())
			Expect(string(analyses)).To(ContainSubstring("a5f7409b-d976-403b-af86-a9976fea901a"))
			Expect(string(analyses)).To(ContainSubstring("finished"))
			Expect(string(analyses)).To(ContainSubstring("git"))
			Expect(string(analyses)).To(ContainSubstring("Pepe"))
		})
	})
}

const (
	SampleValidAnalysis        = `{"data":{"id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"33ef183d-4d37-4515-84c4-099ed0fb8db0","build_number":293,"name":"compliance analysis","text":null,"type":"git","source":"git@github.com:ion-channel/ionic.git","branch":"master","description":"","status":"finished","ruleset_id":"fcd09ba9-c939-4b51-a865-394cc8ddcffa","created_at":"2017-07-18T20:21:30.000Z","updated_at":"2017-07-18T20:21:30.316Z","duration":40209.8742020007,"trigger_hash":"aa8a66375adef765fe9eed5920cfa8352e4c4b70","trigger_text":"Merge pull request #220 from ion-channel/foobranch\n\nadding new coverage format","trigger_author":"Daniel Hess","scan_summaries":[{"id":"c2430d2a-5063-f360-e222-d0dd96b90e22","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"33ef183d-4d37-4515-84c4-099ed0fb8db0","description":"some description","name":"license","summary":"Finished license scan for Statler, found apache-2.0 license.","created_at":"2017-07-18T20:20:53.249Z","updated_at":"2017-07-18T20:20:53.249Z","results":{"data":{"license":{"license":{"name":"LICENSE.md","type":[{"name":"apache-2.0"}]}}},"type":"license"},"duration":674.671019000016},{"id":"d8214cb1-65fa-2d29-226f-d76728b4fb88","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"33ef183d-4d37-4515-84c4-099ed0fb8db0","description":"some description","name":"about_yml","summary":"Finished about_yml scan for Statler, valid .about.yml found.","created_at":"2017-07-18T20:20:53.566Z","updated_at":"2017-07-18T20:20:53.566Z","results":{"type":"about_yml","data":{"about_yml":{"message":"","valid":true,"content":"---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: Statler\n\n# This is the display name of your project. (required)\nfull_name: Statler API\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: This is the api Statler\n\ntype: app\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: alpha\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  statler:\n    name: CC0\n    url: https://github.com/ion-channel/statler/blob/master/LICENSE.md\n\nteam:\n- github: kellyp\n  role: lead\n\ncontact:\n- url: mailto:info@ionchannel.io\n  text: Ion Channel Info Line\n"}}},"duration":1130.52755299941}]}}`
	SampleValidAnalyses        = `{"data":[{"id":"a5f7409b-d976-403b-af86-a9976fea901a","team_id":"dfa7f4d1-1926-4bc0-8ef8-f896be59d6fc","project_id":"57b1478e-fb1a-09a0-b5df-f9ace409f2a5","build_number":263,"name":"Pepe","text":null,"type":"git","source":"git@github.com:ion-channel/pepe.git","branch":"release","description":"","status":"finished","ruleset_id":"39363ed6-697b-4195-b68d-6fea8d2e5fb8","created_at":"2018-03-27T23:37:36.000Z","updated_at":"2018-03-27T23:37:37.029Z","duration":96230.7563549998,"trigger_hash":"805063c8ec9d0daddd95772204b6618c3b8673b6","trigger_text":"Merge pull request #199 from ion-channel/master\n\nrelease","trigger_author":"Dan","suspect":null,"scan_summaries":[{"id":"56530462-8af5-8703-fefd-d4454c48dde8","analysis_id":"a5f7409b-d976-403b-af86-a9976fea901a","team_id":"dfa7f4d1-1926-4bc0-8ef8-f896be59d6fc","project_id":"57b1478e-fb1a-09a0-b5df-f9ace409f2a5","description":"This scan data has not been evaluated against a rule.","name":"virus","summary":"Finished clamav scan for Pepe, found 0 infected files.","created_at":"2018-03-27T23:37:36.429Z","updated_at":"2018-03-27T23:37:36.429Z","results":{"type":"clamav","data":{"clam_av_details":{"clamav_version":"ClamAV 0.99.2","clamav_db_version":"Wed Mar 21 12:24:41 2018\n"},"clamav":{"known_viruses":6444402,"engine_version":"0.99.2","scanned_directories":1703,"scanned_files":9358,"infected_files":0,"data_scanned":"357.46 MB","data_read":"201.43 MB (ratio 1.77:1)","time":"85.213 sec (1 m 25 s)","file_notes":{"empty_file":["/workspace/a5f7409b-d976-403b-af86-a9976fea901a/pepe/ext/lambdas/send-notifications/functions/send-notifications/requests/packages/urllib3/packages/backports/__init__.py","/workspace/a5f7409b-d976-403b-af86-a9976fea901a/pepe/ext/lambdas/send-notifications/functions/send-notifications/requests/packages/urllib3/contrib/__init__.py"]}}}},"duration":85461.409004},{"id":"cfe486c5-4f6b-4e25-ca78-7af5a28f8bfd","analysis_id":"a5f7409b-d976-403b-af86-a9976fea901a","team_id":"dfa7f4d1-1926-4bc0-8ef8-f896be59d6fc","project_id":"57b1478e-fb1a-09a0-b5df-f9ace409f2a5","description":"This scan data has not been evaluated against a rule.","name":"external_coverage","summary":"Code coverage scan run with 51.8% coverage.","created_at":"2018-03-27T23:36:04.779Z","updated_at":"2018-03-27T23:36:04.779Z","results":{"type":"coverage","data":{"coverage":{"value":51.8,"previous_values":[51.8,51.8,51.8,51.8,51.8]},"source":{"name":"","url":""},"notes":""}},"duration":346.718486999634}]},{"id":"5bc4d97e-6399-471c-bf77-1bb89f9951ea","team_id":"dfa7f4d1-1926-4bc0-8ef8-f896be59d6fc","project_id":"57b1478e-fb1a-09a0-b5df-f9ace409f2a5","build_number":262,"name":"Pepe","text":null,"type":"git","source":"git@github.com:ion-channel/pepe.git","branch":"master","description":"","status":"finished","ruleset_id":"39363ed6-697b-4195-b68d-6fea8d2e5fb8","created_at":"2018-03-27T23:28:47.000Z","updated_at":"2018-03-27T23:28:47.501Z","duration":97442.3360035289,"trigger_hash":"c2242f19589b5f854cb4fa9f6be22d57bfb55b32","trigger_text":"Merge pull request #198 from ion-channel/hotrax\n\nHotrax","trigger_author":"Dan","suspect":null,"scan_summaries":[{"id":"30a1b587-adc6-ebc3-d602-1869bdb19265","analysis_id":"5bc4d97e-6399-471c-bf77-1bb89f9951ea","team_id":"dfa7f4d1-1926-4bc0-8ef8-f896be59d6fc","project_id":"57b1478e-fb1a-09a0-b5df-f9ace409f2a5","description":"This scan data has not been evaluated against a rule.","name":"virus","summary":"Finished clamav scan for Pepe, found 0 infected files.","created_at":"2018-03-27T23:28:46.848Z","updated_at":"2018-03-27T23:28:46.848Z","results":{"type":"clamav","data":{"clam_av_details":{"clamav_version":"ClamAV 0.99.2","clamav_db_version":"Wed Mar 21 12:24:41 2018\n"},"clamav":{"known_viruses":6444402,"engine_version":"0.99.2","scanned_directories":1703,"scanned_files":9356,"infected_files":0,"data_scanned":"357.46 MB","data_read":"201.43 MB (ratio 1.77:1)","time":"86.556 sec (1 m 26 s)","file_notes":{"empty_file":["/workspace/5bc4d97e-6399-471c-bf77-1bb89f9951ea/pepe/ext/lambdas/send-notifications/functions/send-notifications/requests/packages/urllib3/packages/backports/__init__.py","/workspace/5bc4d97e-6399-471c-bf77-1bb89f9951ea/pepe/ext/lambdas/send-notifications/functions/send-notifications/requests/packages/urllib3/contrib/__init__.py"]}}}},"duration":86819.9439779855},{"id":"dcbd62aa-29bb-de79-f4ae-4a331131272b","analysis_id":"5bc4d97e-6399-471c-bf77-1bb89f9951ea","team_id":"dfa7f4d1-1926-4bc0-8ef8-f896be59d6fc","project_id":"57b1478e-fb1a-09a0-b5df-f9ace409f2a5","description":"This scan data has not been evaluated against a rule.","name":"difference","summary":"Finished difference scan for Pepe, a difference was detected.","created_at":"2018-03-27T23:27:24.698Z","updated_at":"2018-03-27T23:27:24.698Z","results":{"data":{"difference":true,"checksum":"a5762da03d5aacee21638ff931272d9a41606cd87d5ba7f8ae7993046a5acde1"},"type":"difference"},"duration":5129.1365949437}]}]}`
	SampleValidAnalysisSummary = `{"data":{"id":"a07b82b0-9742-447c-a277-cee0e56abf7f","team_id":"cf47e4d1-bcf8-4990-8ef8-f325ae59d6fc","project_id":"fc42d773-3764-4a78-b3a3-f09bf7241a9c","build_number":189,"name":"a-ionmock","text":null,"type":"git","source":"https://github.com/matthewkmayer/ionmockjavaapp","branch":"master","description":"","status":"failed","ruleset_id":"ec4b43e6-ecfc-42c8-b58c-8a47eab0cc68","created_at":"2017-11-02T12:56:06.000Z","updated_at":"2017-11-02T12:56:06.523Z","duration":50076.0678370007,"trigger_hash":"6a2b494c22e3aeed1fc22fbc549b243b57a7d304","trigger_text":"Merge pull request #1 from ion-channel/Adding-Unit-Tests\n\nAdding jacoco plugin, adding new methods and tests to up code coverage","trigger_author":"Kit Plummer"}}`
)
