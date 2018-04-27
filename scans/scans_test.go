package scans

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestSummary(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Scan Summary", func() {
		g.Describe("Unmarshalling", func() {
			g.It("should populate results with untranslated result", func() {
				var ss Summary
				err := json.Unmarshal([]byte(sampleUntranslatedResults), &ss)

				Expect(err).To(BeNil())
				Expect(ss.TeamID).To(Equal("cuketest"))
				Expect(ss.TranslatedResults).To(BeNil())
				Expect(ss.UntranslatedResults).NotTo(BeNil())
				Expect(ss.UntranslatedResults.License).NotTo(BeNil())
				Expect(ss.UntranslatedResults.License.Name).To(Equal("some license"))
			})

			g.It("should populate results with translated result", func() {
				var ss Summary
				err := json.Unmarshal([]byte(sampleTranslatedResults), &ss)

				Expect(err).To(BeNil())
				Expect(ss.TeamID).To(Equal("cuketest"))
				Expect(ss.UntranslatedResults).To(BeNil())
				Expect(ss.TranslatedResults).NotTo(BeNil())
				Expect(ss.TranslatedResults.Type).To(Equal("community"))
			})
			g.It("should unmarshal a scan summary with a bad, list-ified community results member", func() {
				var ss Summary
				err := json.Unmarshal([]byte(badCommunityResults), &ss)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		g.Describe("Marshalling", func() {
			g.It("should output results with untranslated result", func() {
				s := &Summary{
					summary: &summary{
						ID:          "41e6905a-a16d-45a7-9d2c-2794840ca03e",
						TeamID:      "cuketest",
						AnalysisID:  "c2a79402-e3bf-4069-89c8-7a4ecb10d33f",
						CreatedAt:   time.Now(),
						Description: "This scan data has not been evaluated against a rule.",
						Duration:    1000.1,
						Name:        "license",
						ProjectID:   "35b06118-da91-4ac8-a3d0-db25a3e554c5",
						Summary:     "oh hi",
						UpdatedAt:   time.Now(),
					},
					UntranslatedResults: &UntranslatedResults{
						License: &LicenseResults{
							&License{
								Name: "some license",
								Type: []LicenseType{
									LicenseType{Name: "a license"},
								},
							},
						},
					},
				}

				b, err := json.MarshalIndent(s, "", "  ")
				Expect(err).To(BeNil())

				body := string(b)
				Expect(body).To(ContainSubstring("team_id\": \"cuketest\""))
				Expect(body).NotTo(ContainSubstring("about_yml"))
				Expect(body).NotTo(ContainSubstring("community"))
				Expect(body).NotTo(ContainSubstring("coverage"))
				Expect(body).NotTo(ContainSubstring("dependency"))
				Expect(body).NotTo(ContainSubstring("difference"))
				Expect(body).NotTo(ContainSubstring("ecosystem"))
				Expect(body).NotTo(ContainSubstring("external_vulnerability"))
				Expect(body).NotTo(ContainSubstring("virus"))
				Expect(body).NotTo(ContainSubstring("vulnerability"))
			})

			g.It("should output results with translated result", func() {
				s := &Summary{
					summary: &summary{
						ID:          "41e6905a-a16d-45a7-9d2c-2794840ca03e",
						TeamID:      "cuketest",
						AnalysisID:  "c2a79402-e3bf-4069-89c8-7a4ecb10d33f",
						CreatedAt:   time.Now(),
						Description: "This scan data has not been evaluated against a rule.",
						Duration:    1000.1,
						Name:        "community",
						ProjectID:   "35b06118-da91-4ac8-a3d0-db25a3e554c5",
						Summary:     "oh hi",
						UpdatedAt:   time.Now(),
					},
					TranslatedResults: &TranslatedResults{
						Type: "community",
						Data: &CommunityResults{
							Committers: 5,
							Name:       "reponame",
							URL:        "http://github.com/reponame",
						},
					},
				}

				b, err := json.MarshalIndent(s, "", "  ")
				Expect(err).To(BeNil())

				body := string(b)
				Expect(body).To(ContainSubstring("team_id\": \"cuketest\""))
				Expect(body).To(ContainSubstring("committers\": 5"))
				Expect(body).NotTo(ContainSubstring("about_yml"))
				Expect(body).NotTo(ContainSubstring("coverage"))
				Expect(body).NotTo(ContainSubstring("dependency"))
				Expect(body).NotTo(ContainSubstring("difference"))
				Expect(body).NotTo(ContainSubstring("ecosystem"))
				Expect(body).NotTo(ContainSubstring("external_vulnerability"))
				Expect(body).NotTo(ContainSubstring("license"))
				Expect(body).NotTo(ContainSubstring("virus"))
				Expect(body).NotTo(ContainSubstring("vulnerability"))
			})
		})
	})
}

const (
	badCommunityResults = `{
                "description": "This scan data has not been evaluated against a rule.",
                "created_at": "2018-01-11T05:38:49.478Z",
                "results": {
                    "data": [
                        {
                            "url": "https://github.com/simon/putty",
                            "committers": 0,
                            "name": "simon/putty"
                        }
                    ],
                    "type": "community"
                },
                "updated_at": "2018-01-11T05:38:49.478Z",
                "summary": "Finished community scan for Putty-Source, community data was detected.",
                "team_id": "78536455-a13c-4661-890a-785197f6d9d4",
                "analysis_id": "7f19b79d-5c5f-4c4a-99b3-86c2b0483785",
                "duration": 493.197211995721,
                "project_id": "474b67d8-4817-4209-b614-09c74f3d6c12",
                "id": "6ff786fc-8374-cf33-3a7b-eb5168dc9105",
                "name": "community"
            }`
	sampleUntranslatedResults = `{
  "id": "41e6905a-a16d-45a7-9d2c-2794840ca03e",
  "team_id": "cuketest",
  "project_id": "35b06118-da91-4ac8-a3d0-db25a3e554c5",
  "analysis_id": "c2a79402-e3bf-4069-89c8-7a4ecb10d33f",
  "summary": "oh hi",
  "results": {
    "license": {
      "license": {
        "name": "some license",
        "type": [
          {
            "name": "a license"
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
  "analysis_id": "c2a79402-e3bf-4069-89c8-7a4ecb10d33f",
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
