package reports

import (
	"encoding/json"
	"testing"

	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAnalysisReport(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Analysis Report", func() {
		g.Describe("New Analysis Report", func() {
			g.It("should return a new analysis report", func() {
				s := &scanner.AnalysisStatus{
					Status:  "finished",
					Message: "analysis finished",
				}

				// Note: Once the scans and other objects no longer have the private
				// anonymous fields, this can be changed to use a struct literal
				// initialization of the analysis instead of from json
				var a analyses.Analysis
				json.Unmarshal([]byte(sampleAnalysisPayload), &a)
				Expect(a.ID).To(Equal("f9bca953-80ac-46c4-b195-d37f3bc4f498"))

				var eval scans.Evaluation
				json.Unmarshal([]byte(sampleUntranslatedResults), &eval)

				app := &rulesets.AppliedRulesetSummary{
					RuleEvaluationSummary: &rulesets.RuleEvaluationSummary{
						RulesetName: "super cool ruleset",
						Summary:     "pass",
						Risk:        "low",
						Passed:      true,
						Ruleresults: []scans.Evaluation{eval},
					},
				}

				ar, err := NewAnalysisReport(s, &a, app, true)
				Expect(err).To(BeNil())
				Expect(ar).NotTo(BeNil())

				Expect(ar.Analysis.TriggerText).To(Equal("Merge pull request #220 from ion-channel/foobranch\n\nadding new coverage format"))
			})

			g.It("should return a new analysis report even if the analysis contains translated results", func() {
				s := &scanner.AnalysisStatus{
					Status:  "finished",
					Message: "analysis finished",
				}

				// Note: Once the scans and other objects no longer have the private
				// anonymous fields, this can be changed to use a struct literal
				// initialization of the analysis instead of from json
				var a analyses.Analysis
				json.Unmarshal([]byte(sampleAnalysisPayload), &a)
				Expect(a.ID).To(Equal("f9bca953-80ac-46c4-b195-d37f3bc4f498"))

				var eval scans.Evaluation
				json.Unmarshal([]byte(sampleTranslatedResults), &eval)

				app := &rulesets.AppliedRulesetSummary{
					RuleEvaluationSummary: &rulesets.RuleEvaluationSummary{
						RulesetName: "super cool ruleset",
						Summary:     "pass",
						Risk:        "low",
						Passed:      true,
						Ruleresults: []scans.Evaluation{eval},
					},
				}

				ar, err := NewAnalysisReport(s, &a, app, true)
				Expect(err).To(BeNil())
				Expect(ar).NotTo(BeNil())

				Expect(ar.Analysis.TriggerText).To(Equal("Merge pull request #220 from ion-channel/foobranch\n\nadding new coverage format"))
			})
		})
	})
}

const (
	sampleAnalysisPayload           = `{"id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"e3f62fde-2fd3-4ecd-890f-761281082398","project_id":"5aafe974-b198-434e-9388-1edece09b390","build_number":"293","name":"compliance analysis","text":null,"type":"git","source":"git@github.com:ion-channel/ionic.git","branch":"master","description":"","status":"finished","ruleset_id":"fcd09ba9-c939-4b51-a865-394cc8ddcffa","created_at":"2017-07-18T20:21:30Z","updated_at":"2017-07-18T20:21:30.316Z","duration":40209.8742020007,"trigger_hash":"aa8a66375adef765fe9eed5920cfa8352e4c4b70","trigger_text":"Merge pull request #220 from ion-channel/foobranch\n\nadding new coverage format","trigger_author":"Daniel Hess","scan_summaries":[{"id":"c2430d2a-5063-f360-e222-d0dd96b90e22","team_id":"e3f62fde-2fd3-4ecd-890f-761281082398","project_id":"5aafe974-b198-434e-9388-1edece09b390","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","summary":"Finished license scan for Statler, found apache-2.0 license.","results":{"license":{"license":{"name":"LICENSE.md","type":[{"name":"apache-2.0"}]}}},"created_at":"2017-07-18T20:20:53.249Z","updated_at":"2017-07-18T20:20:53.249Z","duration":674.671019000016,"passed":false,"risk":"","name":"license","description":"some description","type":""},{"id":"d8214cb1-65fa-2d29-226f-d76728b4fb88","team_id":"e3f62fde-2fd3-4ecd-890f-761281082398","project_id":"5aafe974-b198-434e-9388-1edece09b390","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","summary":"Finished about_yml scan for Statler, valid .about.yml found.","results":{"about_yml":{"message":"","valid":true,"content":"---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: Statler\n\n# This is the display name of your project. (required)\nfull_name: Statler API\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: This is the api Statler\n\ntype: app\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: alpha\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  statler:\n    name: CC0\n    url: https://github.com/ion-channel/statler/blob/master/LICENSE.md\n\nteam:\n- github: kellyp\n  role: lead\n\ncontact:\n- url: mailto:info@ionchannel.io\n  text: Ion Channel Info Line\n"}},"created_at":"2017-07-18T20:20:53.566Z","updated_at":"2017-07-18T20:20:53.566Z","duration":1130.52755299941,"passed":false,"risk":"","name":"about_yml","description":"some description","type":""}]}`
	sampleTranslatedAnalysisPayload = `{"id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","team_id":"e3f62fde-2fd3-4ecd-890f-761281082398","project_id":"5aafe974-b198-434e-9388-1edece09b390","build_number":"293","name":"compliance analysis","text":null,"type":"git","source":"git@github.com:ion-channel/ionic.git","branch":"master","description":"","status":"finished","ruleset_id":"fcd09ba9-c939-4b51-a865-394cc8ddcffa","created_at":"2017-07-18T20:21:30Z","updated_at":"2017-07-18T20:21:30.316Z","duration":40209.8742020007,"trigger_hash":"aa8a66375adef765fe9eed5920cfa8352e4c4b70","trigger_text":"Merge pull request #220 from ion-channel/foobranch\n\nadding new coverage format","trigger_author":"Daniel Hess","scan_summaries":[{"id":"d8214cb1-65fa-2d29-226f-d76728b4fb88","team_id":"e3f62fde-2fd3-4ecd-890f-761281082398","project_id":"5aafe974-b198-434e-9388-1edece09b390","analysis_id":"f9bca953-80ac-46c4-b195-d37f3bc4f498","summary":"Finished about_yml scan for Statler, valid .about.yml found.","results":{"type":"about_yml", "data":{"message": "foo message", "valid": true, "content": "some content"}},"created_at":"2017-07-18T20:20:53.566Z","updated_at":"2017-07-18T20:20:53.566Z","duration":1130.52755299941,"passed":false,"risk":"","name":"about_yml","description":"some description","type":""}]}`
	sampleUntranslatedResults       = `{
  "id": "41e6905a-a16d-45a7-9d2c-2794840ca03e",
  "team_id": "cuketest",
  "project_id": "35b06118-da91-4ac8-a3d0-db25a3e554c5",
  "analysis_id": "f9bca953-80ac-46c4-b195-d37f3bc4f498",
  "summary": "oh hi",
  "results": {
    "license": {
      "license": {
        "name": "LICENSE.md",
        "type": [
          {
            "name": "apache-2.0"
          }
        ]
      }
    }
  },
  "created_at": "2018-03-29T13:33:45.924135248-07:00",
  "updated_at": "2018-03-29T13:33:45.924135258-07:00",
  "duration": 1000.1,
  "passed": false,
  "risk": "",
  "name": "license",
  "description": "This scan data has not been evaluated against a rule.",
  "type": ""
}`
	sampleTranslatedResults = `{
  "id": "41e6905a-a16d-45a7-9d2c-2794840ca03e",
  "team_id": "cuketest",
  "project_id": "35b06118-da91-4ac8-a3d0-db25a3e554c5",
  "analysis_id": "f9bca953-80ac-46c4-b195-d37f3bc4f498",
  "summary": "oh hi",
  "results": {
    "type": "community",
    "data": {
      "committers": 5,
      "name": "reponame",
      "url": "http://github.com/reponame"
    }
  },
  "created_at": "2018-03-29T13:50:18.273379563-07:00",
  "updated_at": "2018-03-29T13:50:18.273379579-07:00",
  "duration": 1000.1,
  "passed": false,
  "risk": "",
  "name": "license",
  "description": "This scan data has not been evaluated against a rule.",
  "type": ""
}`
)
