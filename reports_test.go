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
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a report", func() {
			server.AddPath("/v1/report/getAnalysis").
				SetMethods("GET").
				SetPayload([]byte(SampleValidReport)).
				SetStatus(http.StatusOK)

			report, err := client.GetReport("d0fbcdaa-4559-4441-1fcc-d43574004088", "ateamid", "aprojectid")
			Expect(err).To(BeNil())
			Expect(report.ID).To(Equal("d0fbcdaa-4559-4441-1fcc-d43574004088"))
			Expect(report.Name).To(Equal("Kermit"))
		})
	})
}

const (
	SampleValidReport = `{"data":{"branch":"release","build_number":"315","created_at":"2017-08-25T20:32:09.000Z","description":"","duration":49295.8599179983,"id":"d0fbcdaa-4559-4441-1fcc-d43574004088","name":"Kermit","passed":true,"project_id":"bccae6d2-3c3c-44d4-8172-a3959e3bbb3b","risk":"low","ruleset_id":"fcd09ba9-c939-4b51-a865-394cc8ddcffa","ruleset_name":"Production + License","scan_summaries":[{"analysis_id":"d0fbcdaa-4559-4441-1fcc-d43574004088","created_at":"2017-08-25T20:31:09.058Z","description":"The project source is required to include a valid .about.yml file.","duration":998.804689967074,"id":"74b52cc5-38d7-6661-da8c-d4b467a3b6af","name":"Has a valid .about.yml file","passed":true,"project_id":"bccae6d2-3c3c-44d4-8172-a3959e3bbb3b","results":{"about_yml":{"content":"---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: Kermit\n\n# This is the display name of your project. (required)\nfull_name: Kermit API\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: This is the api Kermit\n\ntype: app\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: alpha\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  kermit:\n    name: CC0\n    url: https://github.com/ion-channel/kermit/blob/master/LICENSE.md\n\nteam:\n- github: kellyp\n  role: lead\n\ncontact:\n- url: mailto:info@ionchannel.io\n  text: Ion Channel Info Line\n\n","message":"","valid":true}},"risk":"low","summary":"Finished about_yml scan for Kermit, valid .about.yml found.","team_id":"89cda62c-1926-4990-8ef8-f326be59d6fc","type":"about_yml","updated_at":"2017-08-25T20:31:09.058Z"},{"analysis_id":"d0fbcdaa-4559-4441-1fcc-d43574004088","created_at":"2017-08-25T20:31:09.071Z","description":"The project source is required to include a valid license file.","duration":935.975396016147,"id":"d49a0c05-79c5-840c-0b46-b798e1179efb","name":"Has a valid license file","passed":true,"project_id":"bccae6d2-3c3c-44d4-8172-a3959e3bbb3b","results":{"license":{"license":{"name":"LICENSE.md","type":[{"name":"apache-2.0"}]}}},"risk":"low","summary":"Finished license scan for Kermit, found apache-2.0 license.","team_id":"89cda62c-1926-4990-8ef8-f326be59d6fc","type":"license","updated_at":"2017-08-25T20:31:09.071Z"}],"source":"git@github.com:ion-channel/kermit.git","status":"finished","summary":"","team_id":"89cda62c-1926-4990-8ef8-f326be59d6fc","text":null,"trigger":"source commit","trigger_author":"Kelly Plummer","trigger_hash":"4d1bdadb305e9204b7898b3bbc38fae0a5c8241c","trigger_text":"Merge pull request #215 from ion-channel/master\n\nautomated release partay","type":"git","updated_at":"2017-08-25T20:32:09.356Z"}}`
)
