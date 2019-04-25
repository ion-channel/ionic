package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestReports(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Reports", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get an analysis report", func() {
			server.AddPath("/v1/report/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisReport)).
				SetStatus(http.StatusOK)

			report, err := client.GetAnalysisReport("d0fbcdaa-4559-4441-1fcc-d43574004088", "ateamid", "aprojectid", "atoken")
			Expect(err).To(BeNil())
			Expect(report.ID).To(Equal("d0fbcdaa-4559-4441-1fcc-d43574004088"))
			Expect(report.Name).To(Equal("Kermit"))
		})

		g.It("should get an raw analysis report", func() {
			server.AddPath("/v1/report/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidAnalysisReport)).
				SetStatus(http.StatusOK)

			report, err := client.GetRawAnalysisReport("d0fbcdaa-4559-4441-1fcc-d43574004088", "ateamid", "aprojectid", "atoken")
			Expect(err).To(BeNil())
			Expect(string(report)).To(ContainSubstring("d0fbcdaa-4559-4441-1fcc-d43574004088"))
			Expect(string(report)).To(ContainSubstring("Kermit"))
		})

		g.It("should get a project report", func() {
			server.AddPath("/v1/report/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProjectReport)).
				SetStatus(http.StatusOK)

			report, err := client.GetProjectReport("aprojectid", "ateamid", "atoken")
			Expect(err).To(BeNil())
			Expect(*report.ID).To(Equal("AB3DC2C7-4BB8-4211-8F42-158C8AD4BAE3"))
			Expect(*report.Name).To(Equal("Pepe"))
		})

		g.It("should get a raw project report", func() {
			server.AddPath("/v1/report/getProject").
				SetMethods("GET").
				SetPayload([]byte(SampleValidProjectReport)).
				SetStatus(http.StatusOK)

			report, err := client.GetRawProjectReport("aprojectid", "ateamid", "atoken")
			Expect(err).To(BeNil())
			Expect(string(report)).To(ContainSubstring("AB3DC2C7-4BB8-4211-8F42-158C8AD4BAE3"))
			Expect(string(report)).To(ContainSubstring("Pepe"))
		})
	})
}

const (
	SampleValidAnalysisReport = `{"data":{"branch":"release","created_at":"2017-08-25T20:32:09.000Z","description":"","duration":49295.8599179983,"id":"d0fbcdaa-4559-4441-1fcc-d43574004088","name":"Kermit","passed":true,"project_id":"bccae6d2-3c3c-44d4-8172-a3959e3bbb3b","risk":"low","ruleset_id":"fcd09ba9-c939-4b51-a865-394cc8ddcffa","ruleset_name":"Production + License","scan_summaries":[{"analysis_id":"d0fbcdaa-4559-4441-1fcc-d43574004088","created_at":"2017-08-25T20:31:09.058Z","description":"The project source is required to include a valid .about.yml file.","duration":998.804689967074,"id":"74b52cc5-38d7-6661-da8c-d4b467a3b6af","name":"Has a valid .about.yml file","passed":true,"project_id":"bccae6d2-3c3c-44d4-8172-a3959e3bbb3b","results":{"type":"about_yml","data":{"content":"---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: Kermit\n\n# This is the display name of your project. (required)\nfull_name: Kermit API\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: This is the api Kermit\n\ntype: app\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: alpha\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  kermit:\n    name: CC0\n    url: https://github.com/ion-channel/kermit/blob/master/LICENSE.md\n\nteam:\n- github: kellyp\n  role: lead\n\ncontact:\n- url: mailto:info@ionchannel.io\n  text: Ion Channel Info Line\n\n","message":"","valid":true}},"risk":"low","summary":"Finished about_yml scan for Kermit, valid .about.yml found.","team_id":"89cda62c-1926-4990-8ef8-f326be59d6fc","type":"about_yml","updated_at":"2017-08-25T20:31:09.058Z"},{"analysis_id":"d0fbcdaa-4559-4441-1fcc-d43574004088","created_at":"2017-08-25T20:31:09.071Z","description":"The project source is required to include a valid license file.","duration":935.975396016147,"id":"d49a0c05-79c5-840c-0b46-b798e1179efb","name":"Has a valid license file","passed":true,"project_id":"bccae6d2-3c3c-44d4-8172-a3959e3bbb3b","results":{"type":"license","data":{"license":{"name":"LICENSE.md","type":[{"name":"apache-2.0"}]}}},"risk":"low","summary":"Finished license scan for Kermit, found apache-2.0 license.","team_id":"89cda62c-1926-4990-8ef8-f326be59d6fc","type":"license","updated_at":"2017-08-25T20:31:09.071Z"}],"source":"git@github.com:ion-channel/kermit.git","status":"finished","statuses":{"status":"finished","message":"analysis has successfully finished","scan_status":[],"deliveries":{"seva":{"filename":"foo"}}},"summary":"","team_id":"89cda62c-1926-4990-8ef8-f326be59d6fc","text":null,"trigger":"source commit","trigger_author":"Kelly Plummer","trigger_hash":"4d1bdadb305e9204b7898b3bbc38fae0a5c8241c","trigger_text":"Merge pull request #215 from ion-channel/master\n\nautomated release partay","type":"git","updated_at":"2017-08-25T20:32:09.356Z"}}`
	SampleValidProjectReport  = `{"data":{"id":"AB3DC2C7-4BB8-4211-8F42-158C8AD4BAE3","team_id":"28FB6CD7-2F18-444F-9925-BAB75CFD4A04","ruleset_id":"25174480-5C8F-4C12-8E8D-3E9F125660BE","name":"Pepe","type":"git","source":"git@github.com:ion-channel/pepe.git","branch":"master","description":"","active":true,"chat_channel":"","created_at":"2017-05-26T21:18:28.667Z","updated_at":"2017-07-19T20:02:07.010Z","deploy_key":null,"should_monitor":false,"poc_name":"Daniel","poc_email":"","username":null,"password":null,"key_fingerprint":"","poc_name_hash":"","poc_email_hash":"","aliases":[],"tags":[],"ruleset_name":"Go Project Ruleset","analysis_summaries":[{"analysis_id":"F9D328A5-53E2-4D17-B0E3-09ED60CB1CA2","description":"","branch":"master","risk":"high","summary":"","passed":false,"ruleset_id":"25174480-5C8F-4C12-8E8D-3E9F125660BE","ruleset_name":"Go Project Ruleset","duration":34011.3134330059,"created_at":"2017-09-26T18:23:46.000Z","trigger_hash":"798627047292aa4342cb706c0a5507cd7340a39e","trigger_text":"Merge pull request #151 from ion-channel/test-for-new-account-emails\n\nAdd a test for ensuring we send the right link via email.","trigger_author":"Daniel Hess","trigger":"source commit"},{"analysis_id":"B1454451-C3F0-4226-B2A6-5427E3213116","description":"","branch":"master","risk":"high","summary":"","passed":false,"ruleset_id":"25174480-5C8F-4C12-8E8D-3E9F125660BE","ruleset_name":"Go Project Ruleset","duration":24763.7515759998,"created_at":"2017-07-19T23:19:51.000Z","trigger_hash":"bbd96df8639568106e5ef2fe4a2a7954a587ceb8","trigger_text":"Merge pull request #137 from ion-channel/testing-notice\n\nupdating to notify of running against testing","trigger_author":"Matthew Mayer","trigger":"source commit"}]}}`
)
