package scans

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestScanSummary(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Scan Summary", func() {
		g.Describe("Unmarshalling", func() {
			g.It("should populate results with untranslated result", func() {
				var ss ScanSummary
				err := json.Unmarshal([]byte(sampleUntranslatedResults), &ss)

				Expect(err).To(BeNil())
				Expect(ss.TeamID).To(Equal("cuketest"))
				Expect(ss.TranslatedResults).To(BeNil())
				Expect(ss.UntranslatedResults).NotTo(BeNil())
				Expect(ss.UntranslatedResults.License).NotTo(BeNil())
				Expect(ss.UntranslatedResults.License.Name).To(Equal("some license"))
			})

			g.It("should populate results with translated result", func() {
				var ss ScanSummary
				err := json.Unmarshal([]byte(sampleTranslatedResults), &ss)

				Expect(err).To(BeNil())
				Expect(ss.TeamID).To(Equal("cuketest"))
				Expect(ss.UntranslatedResults).To(BeNil())
				Expect(ss.TranslatedResults).NotTo(BeNil())
				Expect(ss.TranslatedResults.Type).To(Equal("community"))
			})
		})

		g.Describe("Marshalling", func() {
			g.It("should output results with untranslated result", func() {
				s := &ScanSummary{
					scanSummary: &scanSummary{
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
				s := &ScanSummary{
					scanSummary: &scanSummary{
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
