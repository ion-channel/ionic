package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/rulesets"
)

func TestRuleSets(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("RuleSets", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should create a ruleset", func() {
			server.AddPath("/v1/ruleset/createRuleset").
				SetMethods("POST").
				SetPayload([]byte(SampleValidRuleSet)).
				SetStatus(http.StatusOK)

			create := rulesets.CreateRuleSetOptions{
				Description: "all things",
				Name:        "all things",
			}

			ruleset, err := client.CreateRuleSet(create, "sometoken")
			Expect(err).To(BeNil())
			Expect(ruleset.ID).To(Equal("c0210380-3d44-495d-9d10-c7d436a63870"))
			Expect(ruleset.Name).To(Equal("all things"))
			rec := server.HitRecords()[len(server.HitRecords())-1]
			Expect(rec.Body).To(Equal([]byte("{\"name\":\"all things\",\"description\":\"all things\",\"team_id\":\"\",\"rule_ids\":null}")))
			Expect(rec.Path).To(Equal("/v1/ruleset/createRuleset"))
			Expect(rec.Header.Get("Authorization")).To(Equal("Bearer sometoken"))
		})

		g.It("should get a ruleset", func() {
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("GET").
				SetPayload([]byte(SampleValidRuleSet)).
				SetStatus(http.StatusOK)

			ruleset, err := client.GetRuleSet("c0210380-3d44-495d-9d10-c7d436a63870", "a2d2a3e5-e274-bb88-aef2-1d47f029c289", "sometoken")
			Expect(err).To(BeNil())
			Expect(ruleset.ID).To(Equal("c0210380-3d44-495d-9d10-c7d436a63870"))
			Expect(ruleset.Name).To(Equal("all things"))
		})

		g.It("should get rulesets for a team id", func() {
			server.AddPath("/v1/ruleset/getRulesets").
				SetMethods("GET").
				SetPayload([]byte(SampleValidRuleSets)).
				SetStatus(http.StatusOK)

			rulesets, err := client.GetRuleSets("a2d2a3e5-e274-bb88-aef2-1d47f029c289", "sometoken", pagination.AllItems)
			Expect(err).To(BeNil())
			Expect(len(rulesets)).To(Equal(2))
			Expect(rulesets[0].ID).To(Equal("c0210380-3d44-495d-9d10-c7d436a63870"))
			Expect(rulesets[0].Name).To(Equal("all things"))
			Expect(len(rulesets[0].Rules)).To(Equal(2))
		})

		g.It("should get an applied ruleset", func() {
			server.AddPath("/v1/ruleset/getAppliedRulesetForProject").
				SetMethods("GET").
				SetPayload([]byte(SampleAppliedRuleset)).
				SetStatus(http.StatusOK)

			applied, err := client.GetAppliedRuleSet("0AA84783-C2D7-440F-B78B-7E315288398A", "48501083-16EA-4254-9C73-60419A7A4ECB", "", "sometoken")
			Expect(err).To(BeNil())
			Expect(len(applied.RuleEvaluationSummary.Ruleresults)).To(Equal(7))
		})

		g.It("should get a raw applied ruleset", func() {
			server.AddPath("/v1/ruleset/getAppliedRulesetForProject").
				SetMethods("GET").
				SetPayload([]byte(SampleAppliedRuleset)).
				SetStatus(http.StatusOK)

			applied, err := client.GetRawAppliedRuleSet("0AA84783-C2D7-440F-B78B-7E315288398A", "48501083-16EA-4254-9C73-60419A7A4ECB", "", "sometoken", pagination.AllItems)
			Expect(err).To(BeNil())
			Expect(string(applied)).To(ContainSubstring("\"team_id\":\"800E898B-CCD8-4394-A559-F17D08030413\""))
			Expect(string(applied)).To(ContainSubstring("\"project_id\":\"32D701E1-E173-43EF-9CC8-E4CB27417FD8\""))
		})
	})
}

const (
	SampleValidRuleSet   = `{"data":{"id":"c0210380-3d44-495d-9d10-c7d436a63870","team_id":"a2d2a3e5-e274-bb88-aef2-1d47f029c289","name":"all things","description":"about.yml dependencies vulnerabilities code coverage","rule_ids":["d928de6b-9aa0-2b98-4663-17c23d68efc3","c30b9179-56c3-040d-aa2c-571ef31dbe3a","276bbec3-cc77-44b9-a46d-c7760947ec9d","00be1862-959c-45d8-8fb5-2b748fe854d6"],"created_at":"2016-10-04T16:51:59.966Z","updated_at":"2016-10-04T16:51:59.966Z","rules":[{"id":"d928de6b-9aa0-2b98-4663-17c23d68efc3","scan_type":"coverage","name":"Code Coverage \u003e 70%","description":"A longer description of the rule: Code Coverage \u003e 70%","category":"Code Coverage","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:26.257Z","updated_at":"2016-09-19T21:38:26.257Z"},{"id":"c30b9179-56c3-040d-aa2c-571ef31dbe3a","scan_type":"about_yml","name":"Has a valid .about.yml file","description":"The project source is required to include a valid .about.yml file.","category":"About Dot Yaml","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:27.112Z","updated_at":"2016-09-19T21:38:27.112Z"},{"id":"276bbec3-cc77-44b9-a46d-c7760947ec9d","scan_type":"dependencies","name":"Dependencies Version Exist","description":"A longer description of the rule: Dependencies Exist","category":"Dependencies","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:48:30.725Z","updated_at":"2016-09-19T21:48:30.725Z"},{"id":"00be1862-959c-45d8-8fb5-2b748fe854d6","scan_type":"vulnerabilities","name":"Critical Vulnerabilities \u003c 1","description":"A longer description of the rule: Critical Vulnerabilities \u003c 1","category":"Vulnerabilities","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:48:30.731Z","updated_at":"2016-09-19T21:48:30.731Z"}]}}`
	SampleValidRuleSets  = `{"data":[{"id":"c0210380-3d44-495d-9d10-c7d436a63870","team_id":"a2d2a3e5-e274-bb88-aef2-1d47f029c289","name":"all things","description":"about.yml dependencies vulnerabilities code coverage","rule_ids":["d928de6b-9aa0-2b98-4663-17c23d68efc3","c30b9179-56c3-040d-aa2c-571ef31dbe3a"],"created_at":"2016-10-04T16:51:59.966Z","updated_at":"2016-10-04T16:51:59.966Z","rules":[{"id":"d928de6b-9aa0-2b98-4663-17c23d68efc3","scan_type":"coverage","name":"Code Coverage > 70%","description":"A longer description of the rule: Code Coverage > 70%","category":"Code Coverage","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:26.257Z","updated_at":"2016-09-19T21:38:26.257Z"},{"id":"c30b9179-56c3-040d-aa2c-571ef31dbe3a","scan_type":"about_yml","name":"Has a valid .about.yml file","description":"The project source is required to include a valid .about.yml file.","category":"About Dot Yaml","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:27.112Z","updated_at":"2016-09-19T21:38:27.112Z"}]},{"id":"ec4b43e6-ecfc-42c8-b58c-8a47eab0cc68","team_id":"a2d2a3e5-e274-bb88-aef2-1d47f029c289","name":"Code Coverage > 70%","description":"Code Coverage > 70%","rule_ids":["d928de6b-9aa0-2b98-4663-17c23d68efc3"],"created_at":"2016-10-26T19:30:56.726Z","updated_at":"2016-10-26T19:30:56.726Z","rules":[{"id":"d928de6b-9aa0-2b98-4663-17c23d68efc3","scan_type":"coverage","name":"Code Coverage > 70%","description":"A longer description of the rule: Code Coverage > 70%","category":"Code Coverage","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:26.257Z","updated_at":"2016-09-19T21:38:26.257Z"}]}]}`
	SampleAppliedRuleset = `{"data":{"project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","team_id":"800E898B-CCD8-4394-A559-F17D08030413","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","rule_evaluation_summary":{"summary":"fail","ruleresults":[{"id":"f9eec625-88d9-fca1-02db-d5062957ced5","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"License","summary":"Finished license scan for a-ionmock, failed to detect license.","created_at":"2017-09-27T12:55:34.480Z","updated_at":"2017-09-27T12:55:34.480Z","results":{"license":{"license":{"name":"Not found","type":[]}}},"duration":1.10026499987725,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"e0eb6936-9074-6f03-e861-ae65290fa3c3","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"Ecosystems","summary":"Finished ecosystems scan for a-ionmock, found {\"Java\"=>2582} ecosystems in project.","created_at":"2017-09-27T12:55:34.503Z","updated_at":"2017-09-27T12:55:34.503Z","results":{"ecosystems":{"Java":2582}},"duration":16.95143500001,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"d1035d70-6516-aa94-faa4-bf77b06bfa82","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"Difference","summary":"Finished difference scan for a-ionmock, a difference was detected.","created_at":"2017-09-27T12:55:34.783Z","updated_at":"2017-09-27T12:55:34.783Z","results":{"difference":{"difference":true,"checksum":"d63371f4cea3a8b80fc7838764e448955cc8ff32bdb41d06ba6055b98883380b"}},"duration":542.002373999821,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"02ce4f55-6038-05f0-0303-e5e43b36beed","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"About_yml","summary":"Finished about_yml scan for a-ionmock, valid .about.yml found.","created_at":"2017-09-27T12:55:35.354Z","updated_at":"2017-09-27T12:55:35.354Z","results":{"about_yml":{"message":"","valid":true,"content":"---\n# .about.yml project metadata\n#\n# Copy this template into your project repository's root directory as\n# .about.yml and fill in the fields as described below.\n\n# This is a short name of your project that can be used as a URL slug.\n# (required)\nname: ionmockjavaapp\n\n# This is the display name of your project. (required)\nfull_name: ionmockjavaapp\n\n# What is the problem your project solves? What is the solution? Use the\n# format shown below. The #dashboard team will gladly help you put this\n# together for your project. (required)\ndescription: Provides a test harness for java (maven) projects\n\n# What is the measurable impact of your project? Use the format shown below.\n# The #dashboard team will gladly help you put this together for your project.\n# (required)\nimpact: high\n\n# What kind of team owns the repository? (required)\n# values: guild, working-group, project\nowner_type: project\n\n# What is your project's current status? (required)\n# values: discovery, alpha, beta, live\nstage: live\n\n# Should this repo have automated tests? If so, set to true. (required)\n# values: true, false\ntestable: true\n\nlicenses:\n  doozer:\n    name: GPLV2\n    url: https://github.com/ion-channel/java-lew/blob/master/license.txt\n\nteam:\n- github: kitplummer\n  role: lead\n"}},"duration":1168.83430600001,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"504ea20a-366e-ef90-0723-6febb6f350a1","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"Dependency","summary":"Finished dependency scan for a-ionmock, found 0 with no version and 2 with updates available.","created_at":"2017-09-27T12:55:48.484Z","updated_at":"2017-09-27T12:55:48.484Z","results":{"dependency":{"dependencies":[{"latest_version":"2.0","org":"net.sourceforge.javacsv","name":"javacsv","type":"maven","package":"jar","version":"2.0","scope":"compile"},{"latest_version":"4.12","org":"junit","name":"junit","type":"maven","package":"jar","version":"4.11","scope":"test"},{"latest_version":"1.4-atlassian-1","org":"org.hamcrest","name":"hamcrest-core","type":"maven","package":"jar","version":"1.3","scope":"test"},{"latest_version":"4.5.2","org":"org.apache.httpcomponents","name":"httpclient","type":"maven","package":"jar","version":"4.3.4","scope":"compile"},{"latest_version":"4.4.5","org":"org.apache.httpcomponents","name":"httpcore","type":"maven","package":"jar","version":"4.3.2","scope":"compile"},{"latest_version":"99.0-does-not-exist","org":"commons-logging","name":"commons-logging","type":"maven","package":"jar","version":"1.1.3","scope":"compile"},{"latest_version":"20041127.091804","org":"commons-codec","name":"commons-codec","type":"maven","package":"jar","version":"1.6","scope":"compile"}],"meta":{"first_degree_count":3,"no_version_count":0,"total_unique_count":7,"update_available_count":2}}},"duration":14287.5910550001,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"3abfd693-5f61-e4ac-ec72-763d42dfb4fb","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"Vulnerability","summary":"Finished vulnerability scan for a-ionmock, found 0 vulnerabilities.","created_at":"2017-09-27T12:55:48.719Z","updated_at":"2017-09-27T12:55:48.719Z","results":{"vulnerabilities":{"vulnerabilities":[],"meta":{"vulnerability_count":0}}},"duration":89.9386969999796,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"ce49954f-02d3-9675-380a-eb974ab8b68d","analysis_id":"B061D58B-FDFD-46BF-A766-2D38DE3B1D7B","team_id":"800E898B-CCD8-4394-A559-F17D08030413","project_id":"32D701E1-E173-43EF-9CC8-E4CB27417FD8","description":"some description","name":"Virus","summary":"Finished clamav scan for a-ionmock, found 0 infected files.","created_at":"2017-09-27T12:55:50.525Z","updated_at":"2017-09-27T12:55:50.525Z","results":{"clam_av_details":{"clamav_version":"ClamAV 0.99.2","clamav_db_version":"Wed Sep 27 04:44:38 2017\n"},"clamav":{"known_viruses":6303819,"engine_version":"0.99.2","scanned_directories":29,"scanned_files":30,"infected_files":0,"data_scanned":"0.04 MB","data_read":"0.02 MB (ratio 1.83:1)","time":"16.144 sec (0 m 16 s)","file_notes":{}}},"duration":16180.6532439996,"passed":false,"risk":"n/a","type":"Not Evaluated"}]},"rule_eval_created_at":"2017-09-27T12:55:50+00:00","created_at":"2017-09-27T12:55:50.814Z","updated_at":"2017-09-27T12:55:50.814Z"}}`
)
