package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"

	"github.com/ion-channel/ionic/pagination"
)

func TestRuleSets(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("RuleSets", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get a ruleset", func() {
			server.AddPath("/v1/ruleset/getRuleset").
				SetMethods("GET").
				SetPayload([]byte(SampleValidRuleSet)).
				SetStatus(http.StatusOK)

			ruleset, err := client.GetRuleSet("c0210380-3d44-495d-9d10-c7d436a63870", "a2d2a3e5-e274-bb88-aef2-1d47f029c289")
			Expect(err).To(BeNil())
			Expect(ruleset.ID).To(Equal("c0210380-3d44-495d-9d10-c7d436a63870"))
			Expect(ruleset.Name).To(Equal("all things"))
		})

		g.It("should get rulesets for a team id", func() {
			server.AddPath("/v1/ruleset/getRulesets").
				SetMethods("GET").
				SetPayload([]byte(SampleValidRuleSets)).
				SetStatus(http.StatusOK)

			rulesets, err := client.GetRuleSets("a2d2a3e5-e274-bb88-aef2-1d47f029c289", pagination.AllItems)
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

			applied, err := client.GetAppliedRuleSet("0AA84783-C2D7-440F-B78B-7E315288398A", "48501083-16EA-4254-9C73-60419A7A4ECB", "")
			Expect(err).To(BeNil())
			Expect(applied.Status).To(Equal("finished"))
			Expect(len(applied.ScanSummaries)).To(Equal(2))
		})

		g.It("should get a raw applied ruleset", func() {
			server.AddPath("/v1/ruleset/getAppliedRulesetForProject").
				SetMethods("GET").
				SetPayload([]byte(SampleAppliedRuleset)).
				SetStatus(http.StatusOK)

			applied, err := client.GetRawAppliedRuleSet("0AA84783-C2D7-440F-B78B-7E315288398A", "48501083-16EA-4254-9C73-60419A7A4ECB", "", pagination.AllItems)
			Expect(err).To(BeNil())
			Expect(string(applied)).To(ContainSubstring("\"team_id\":\"48501083-16EA-4254-9C73-60419A7A4ECB\""))
			Expect(string(applied)).To(ContainSubstring("\"project_id\":\"0AA84783-C2D7-440F-B78B-7E315288398A\""))
		})
	})
}

const (
	SampleValidRuleSet   = `{"data":{"id":"c0210380-3d44-495d-9d10-c7d436a63870","team_id":"a2d2a3e5-e274-bb88-aef2-1d47f029c289","name":"all things","description":"about.yml dependencies vulnerabilities code coverage","rule_ids":["d928de6b-9aa0-2b98-4663-17c23d68efc3","c30b9179-56c3-040d-aa2c-571ef31dbe3a","276bbec3-cc77-44b9-a46d-c7760947ec9d","00be1862-959c-45d8-8fb5-2b748fe854d6"],"created_at":"2016-10-04T16:51:59.966Z","updated_at":"2016-10-04T16:51:59.966Z","rules":[{"id":"d928de6b-9aa0-2b98-4663-17c23d68efc3","scan_type":"coverage","name":"Code Coverage \u003e 70%","description":"A longer description of the rule: Code Coverage \u003e 70%","category":"Code Coverage","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:26.257Z","updated_at":"2016-09-19T21:38:26.257Z"},{"id":"c30b9179-56c3-040d-aa2c-571ef31dbe3a","scan_type":"about_yml","name":"Has a valid .about.yml file","description":"The project source is required to include a valid .about.yml file.","category":"About Dot Yaml","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:27.112Z","updated_at":"2016-09-19T21:38:27.112Z"},{"id":"276bbec3-cc77-44b9-a46d-c7760947ec9d","scan_type":"dependencies","name":"Dependencies Version Exist","description":"A longer description of the rule: Dependencies Exist","category":"Dependencies","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:48:30.725Z","updated_at":"2016-09-19T21:48:30.725Z"},{"id":"00be1862-959c-45d8-8fb5-2b748fe854d6","scan_type":"vulnerabilities","name":"Critical Vulnerabilities \u003c 1","description":"A longer description of the rule: Critical Vulnerabilities \u003c 1","category":"Vulnerabilities","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:48:30.731Z","updated_at":"2016-09-19T21:48:30.731Z"}]}}`
	SampleValidRuleSets  = `{"data":[{"id":"c0210380-3d44-495d-9d10-c7d436a63870","team_id":"a2d2a3e5-e274-bb88-aef2-1d47f029c289","name":"all things","description":"about.yml dependencies vulnerabilities code coverage","rule_ids":["d928de6b-9aa0-2b98-4663-17c23d68efc3","c30b9179-56c3-040d-aa2c-571ef31dbe3a"],"created_at":"2016-10-04T16:51:59.966Z","updated_at":"2016-10-04T16:51:59.966Z","rules":[{"id":"d928de6b-9aa0-2b98-4663-17c23d68efc3","scan_type":"coverage","name":"Code Coverage > 70%","description":"A longer description of the rule: Code Coverage > 70%","category":"Code Coverage","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:26.257Z","updated_at":"2016-09-19T21:38:26.257Z"},{"id":"c30b9179-56c3-040d-aa2c-571ef31dbe3a","scan_type":"about_yml","name":"Has a valid .about.yml file","description":"The project source is required to include a valid .about.yml file.","category":"About Dot Yaml","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:27.112Z","updated_at":"2016-09-19T21:38:27.112Z"}]},{"id":"ec4b43e6-ecfc-42c8-b58c-8a47eab0cc68","team_id":"a2d2a3e5-e274-bb88-aef2-1d47f029c289","name":"Code Coverage > 70%","description":"Code Coverage > 70%","rule_ids":["d928de6b-9aa0-2b98-4663-17c23d68efc3"],"created_at":"2016-10-26T19:30:56.726Z","updated_at":"2016-10-26T19:30:56.726Z","rules":[{"id":"d928de6b-9aa0-2b98-4663-17c23d68efc3","scan_type":"coverage","name":"Code Coverage > 70%","description":"A longer description of the rule: Code Coverage > 70%","category":"Code Coverage","policy_url":"url","remediation_url":"url","created_at":"2016-09-19T21:38:26.257Z","updated_at":"2016-09-19T21:38:26.257Z"}]}]}`
	SampleAppliedRuleset = `{"data":{"id":"89188D9D-A125-435A-BAB1-CA872B3A7F2F","team_id":"48501083-16EA-4254-9C73-60419A7A4ECB","project_id":"0AA84783-C2D7-440F-B78B-7E315288398A","build_number":"160","name":"a-ionmock","text":null,"type":"git","source":"https://github.com/matthewkmayer/ionmockjavaapp","branch":"master","description":"","status":"finished","ruleset_id":"ec4b43e6-ecfc-42c8-b58c-8a47eab0cc68","created_at":"2017-09-27T12:55:50.000Z","updated_at":"2017-09-27T12:55:50.622Z","duration":32287.0713759994,"trigger_hash":"6a2b494c22e3aeed1fc22fbc549b243b57a7d304","trigger_text":"Merge pull request #1 from ion-channel/Adding-Unit-Tests\n\nAdding jacoco plugin, adding new methods and tests to up code coverage","trigger_author":"Kit Plummer","scan_summaries":[{"id":"2138FBCA-6804-4689-A90D-2FFC0737A308","analysis_id":"89188D9D-A125-435A-BAB1-CA872B3A7F2F","team_id":"48501083-16EA-4254-9C73-60419A7A4ECB","project_id":"0AA84783-C2D7-440F-B78B-7E315288398A","description":"some description","name":"License","summary":"Finished license scan for a-ionmock, failed to detect license.","created_at":"2017-09-27T12:55:34.480Z","updated_at":"2017-09-27T12:55:34.480Z","results":{"type":"license","data":{"license":{"name":"Not found","type":[]}}},"duration":1.10026499987725,"passed":false,"risk":"n/a","type":"Not Evaluated"},{"id":"2469F81B-CA83-4F92-BCFC-5A6508F654DF","analysis_id":"89188D9D-A125-435A-BAB1-CA872B3A7F2F","team_id":"48501083-16EA-4254-9C73-60419A7A4ECB","project_id":"0AA84783-C2D7-440F-B78B-7E315288398A","description":"some description","name":"Ecosystems","summary":"Finished ecosystems scan for a-ionmock, found {\"Java\"=>2582} ecosystems in project.","created_at":"2017-09-27T12:55:34.503Z","updated_at":"2017-09-27T12:55:34.503Z","results":{"type":"ecosystems","data":{"Java":2582}},"duration":16.95143500001,"passed":false,"risk":"n/a","type":"Not Evaluated"}],"ruleset_name":"Code Coverage > 70%","risk":"high","summary":"","trigger":"source commit","passed":false}}`
)
