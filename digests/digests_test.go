package digests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/franela/goblin"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"
	. "github.com/onsi/gomega"
)

func TestDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Digests", func() {
		g.Describe("New Digests", func() {
			g.It("should be sortable by index", func() {
				// ds := []Digest{
				// 	Digest{Index: 1},
				// 	Digest{Index: 3},
				// 	Digest{Index: 2},
				// 	Digest{Index: 0},
				// }

				var analysisStatus scanner.AnalysisStatus
				var appliedRulesetSummary rulesets.AppliedRulesetSummary

				err := json.Unmarshal([]byte(scanSummary), &analysisStatus)
				Expect(err).To(BeNil())
				// s, _ := json.Marshal(analysisStatus)
				// fmt.Printf("%s", s)
				//fmt.Printf("%s", analysisStatus)
				err = json.Unmarshal([]byte(appliedRuleset), &appliedRulesetSummary)
				Expect(err).To(BeNil())
				// s, _ = json.Marshal(appliedRulesetSummary)
				// fmt.Printf("%s", s)

				ds, err := NewDigests(&appliedRulesetSummary, analysisStatus.ScanStatus)

				s, _ := json.Marshal(ds)
				fmt.Printf("\n\n Digests: %s \n", s)

				Expect(ds).ToNot(BeNil())
				Expect(ds[16].RuleID).To(Equal("2981e1b0-0c8f-0137-8fe7-186590d3c755"))
				Expect(ds[17].RuleID).To(Equal("d3b66d48-40a1-11eb-b378-0242ac130002"))
				// Expect(ds[1].Index).To(Equal(3))
				// Expect(ds[2].Index).To(Equal(2))
				// Expect(ds[3].Index).To(Equal(0))

				// sort.Slice(ds, func(i, j int) bool { return ds[i].Index < ds[j].Index })

				// Expect(ds[0].Index).To(Equal(0))
				// Expect(ds[1].Index).To(Equal(1))
				// Expect(ds[2].Index).To(Equal(2))
				// Expect(ds[3].Index).To(Equal(3))
			})
		})
	})
}

const (
	scanSummary    = `{"id":"072bd656-4da2-43d7-b181-a408e4334fa0","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","message":"Completed compliance analysis","branch":"master","status":"finished","unreachable_error":false,"analysis_event_src":"","created_at":"2021-01-05T21:43:46.300312Z","updated_at":"2021-01-05T21:44:45.927714Z","scan_status":[{"id":"29567df8-7fa5-0ac0-bc1f-1c6442113e67","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished vulnerability scan for Apache-Maven, found 21 vulnerabilities.","name":"vulnerability","read":"false","status":"finished","created_at":"2021-01-05T21:43:54.418975Z","updated_at":"2021-01-05T21:44:39.371464Z"},{"id":"35dc6087-c4e0-2f6b-3a61-c4ea2aa54c4b","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished dependency scan for Apache-Maven, found 284 dependencies, 151 with no version and 60 with updates available.","name":"dependency","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.294151Z","updated_at":"2021-01-05T21:44:02.970115Z"},{"id":"abc98be7-2c83-2ec9-bb40-95edc68a5210","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished clamav scan for Apache-Maven, found 0 infected files.","name":"virus","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.217233Z","updated_at":"2021-01-05T21:44:00.56586Z"},{"id":"9e197528-3a89-3b89-629b-bdaeddaab025","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished ecosystems scan for Apache-Maven, Shell, Batchfile, Java, HTML were detected in project.","name":"ecosystems","read":"false","status":"finished","created_at":"2021-01-05T21:43:53.200719Z","updated_at":"2021-01-05T21:43:57.672072Z"},{"id":"fb7a273c-8f4c-7050-0dc4-96f824eba14c","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished license scan for Apache-Maven, found Apache-2.0 license.","name":"license","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.367868Z","updated_at":"2021-01-05T21:43:57.520001Z"},{"id":"195914b8-fad8-9421-c1b0-25079df579d4","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished difference scan for Apache-Maven, a difference was detected.","name":"difference","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.15093Z","updated_at":"2021-01-05T21:43:57.475984Z"},{"id":"308bd861-3282-3e5c-95a1-ced7ba7f36a7","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished community scan for Apache-Maven, community data was detected.","name":"community","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.502173Z","updated_at":"2021-01-05T21:43:57.217945Z"},{"id":"edb1347d-5545-24a3-364b-848a12a4376f","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished buildsystems scan for Apache-Maven, no compilers were detected in project.","name":"buildsystems","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.422098Z","updated_at":"2021-01-05T21:43:54.910371Z"},{"id":"7ec70ef1-46c9-3499-5798-1a400547d6ca","analysis_status_id":"072bd656-4da2-43d7-b181-a408e4334fa0","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","message":"Finished about_yml scan for Apache-Maven, valid .about.yml found.","name":"about_yml","read":"false","status":"finished","created_at":"2021-01-05T21:43:52.07924Z","updated_at":"2021-01-05T21:43:52.347611Z"}],"deliveries":{"git":{"id":"ce8f8b37-0bfa-4788-8320-a45ba2d314b6","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","analysis_id":"072bd656-4da2-43d7-b181-a408e4334fa0","destination":"ion-channel-deliver/","status":"delivery_canceled","filename":"","hash":"","message":"","created_at":"2021-01-05T21:45:34.797761Z","updated_at":"2021-01-05T21:45:34.797761Z","delivered_at":"0001-01-01T00:00:00Z"},"report":{"id":"89423b05-1d75-408e-818a-95ba5d8253e2","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","analysis_id":"072bd656-4da2-43d7-b181-a408e4334fa0","destination":"ion-channel-deliver/analysis-20210105.214444.json","status":"delivery_finished","filename":"analysis-20210105.214444.json","hash":"","message":"","created_at":"2021-01-05T21:45:34.991734Z","updated_at":"2021-01-05T21:45:34.991734Z","delivered_at":"0001-01-01T00:00:00Z"},"seva":{"id":"7aa6102e-b341-4cc0-8c56-8654ef64de3d","team_id":"646fa3e5-e274-4884-aef2-1d47f029c289","project_id":"c5e4672e-85c2-4c35-ac3e-c08449341f12","analysis_id":"072bd656-4da2-43d7-b181-a408e4334fa0","destination":"ion-channel-deliver/seva-20210105.214444.xml","status":"delivery_finished","filename":"seva-20210105.214444.xml","hash":"","message":"","created_at":"2021-01-05T21:45:35.093575Z","updated_at":"2021-01-05T21:45:35.093575Z","delivered_at":"0001-01-01T00:00:00Z"}}}`
	appliedRuleset = `{
			"project_id": "c5e4672e-85c2-4c35-ac3e-c08449341f12",
			"team_id": "646fa3e5-e274-4884-aef2-1d47f029c289",
			"analysis_id": "072bd656-4da2-43d7-b181-a408e4334fa0",
			"rule_evaluation_summary": {
				"ruleset_name": "single comitter, 1 year activity",
				"summary": "fail",
				"risk": "high",
				"passed": false,
				"ruleresults": [{
					"id": "308bd861-3282-3e5c-95a1-ced7ba7f36a7",
					"team_id": "646fa3e5-e274-4884-aef2-1d47f029c289",
					"project_id": "c5e4672e-85c2-4c35-ac3e-c08449341f12",
					"analysis_id": "072bd656-4da2-43d7-b181-a408e4334fa0",
					"rule_id": "2981e1b0-0c8f-0137-8fe7-186590d3c755",
					"ruleset_id": "e41a4cd5-0774-4166-b18f-bf73a0852ba5",
					"summary": "Finished community scan for Apache-Maven, community data was detected.",
					"results": {
						"type": "community",
						"data": {
							"committers": 3,
							"name": "apache/maven",
							"url": "https://github.com/apache/maven",
							"committed_at": "2019-12-15T21:00:00Z"
						}
					},
					"created_at": "2021-01-05T21:43:55.670147Z",
					"updated_at": "2021-01-05T21:43:55.670147Z",
					"duration": 1441.2947529999656,
					"name": "Has more than one committer",
					"description": "The project must have more than 1 committer.",
					"risk": "low",
					"type": "community",
					"passed": true
				}, {
					"id": "308bd861-3282-3e5c-95a1-ced7ba7f36a7",
					"team_id": "646fa3e5-e274-4884-aef2-1d47f029c289",
					"project_id": "c5e4672e-85c2-4c35-ac3e-c08449341f12",
					"analysis_id": "072bd656-4da2-43d7-b181-a408e4334fa0",
					"rule_id": "d3b66d48-40a1-11eb-b378-0242ac130002",
					"ruleset_id": "e41a4cd5-0774-4166-b18f-bf73a0852ba5",
					"summary": "Finished community scan for Apache-Maven, community data was detected.",
					"results": {
						"type": "community",
						"data": {
							"committers": 3,
							"name": "apache/maven",
							"url": "https://github.com/apache/maven",
							"committed_at": "2019-12-15T21:00:00Z"
						}
					},
					"created_at": "2021-01-05T21:43:55.670147Z",
					"updated_at": "2021-01-05T21:43:55.670147Z",
					"duration": 1441.2947529999656,
					"name": "1 year since last commit",
					"description": "The project must be active within the the given timeframe. (Note: this rule is only applicable to .git type projects)",
					"risk": "high",
					"type": "community_activity",
					"passed": false
				},{
					"id": "7ec70ef1-46c9-3499-5798-1a400547d6ca",
					"team_id": "646fa3e5-e274-4884-aef2-1d47f029c289",
					"project_id": "c5e4672e-85c2-4c35-ac3e-c08449341f12",
					"analysis_id": "072bd656-4da2-43d7-b181-a408e4334fa0",
					"rule_id": "n/a",
					"ruleset_id": "e41a4cd5-0774-4166-b18f-bf73a0852ba5",
					"summary": "Finished about_yml scan for Apache-Maven, valid .about.yml found.",
					"results": {
						"type": "about_yml",
						"data": {
							"message": "",
							"valid": true,
							"content": ""
						}
					},
					"created_at": "2021-01-05T21:43:52.178618Z",
					"updated_at": "2021-01-05T21:43:52.178618Z",
					"duration": 1.5864579982007854,
					"name": "About_yml",
					"description": "",
					"risk": "n/a",
					"type": "Not Evaluated",
					"passed": false
				}]
			},
			"created_at": "2021-01-05T21:44:45.059529Z",
			"updated_at": "2021-01-05T21:44:45.059529Z"
		}`
)
