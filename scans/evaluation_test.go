package scans

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestEvaluation(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Evaluation", func() {
		g.Describe("Translating", func() {
			g.It("should return virus results from a new evaluation", func() {
				j := json.RawMessage(`{"results":{"clamav": {"time": "0.133 sec (0 m 0 s)", "file_notes": {"win.test.eicar_hdb-1_found": ["/workspace/ae0013af-9dfc-41be-9ae5-99c66d0e43c0/eicar/bin/eicar.com"]}, "scanned_files": 37, "infected_files": 1}, "clam_av_details": {"clamav_version": "ClamAV 0.101.5", "clamav_db_version": "Mon Apr 20 12:00:23 2020\n"}}}`)
				r := json.RawMessage(`{"type":"virus","data":{"known_viruses":0,"engine_version":"","scanned_directories":0,"scanned_files":37,"infected_files":1,"data_scanned":"","data_read":"","time":"0.133 sec (0 m 0 s)","file_notes":{"win.test.eicar_hdb-1_found":["/workspace/ae0013af-9dfc-41be-9ae5-99c66d0e43c0/eicar/bin/eicar.com"]},"clam_av_details":{"clamav_version":"ClamAV 0.101.5","clamav_db_version":"Mon Apr 20 12:00:23 2020\n"}}}`)
				e := NewEval()
				e.UnmarshalJSON(j)
				Expect(e.Results).To(Equal(r))
				j, err := e.MarshalJSON()
				Expect(err).To(BeNil())
				err = json.Unmarshal(j, &e)
				Expect(err).To(BeNil())
				v := e.TranslatedResults.Data.(VirusResults)
				Expect(v.ClamavDetails.ClamavVersion).To(Equal("ClamAV 0.101.5"))
			})
			g.It("should translate an untranslated evaluation", func() {
				ee := &Evaluation{
					UntranslatedResults: &UntranslatedResults{
						License: &LicenseResults{},
					},
				}
				Expect(ee.UntranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults).To(BeNil())

				err := ee.Translate()
				Expect(err).To(BeNil())
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("license"))
				Expect(ee.Results).NotTo(BeNil())
				Expect(len(ee.Results)).NotTo(Equal(0))
			})

			g.It("should not translate an already translated summary", func() {
				ee := &Evaluation{
					UntranslatedResults: &UntranslatedResults{
						License: &LicenseResults{},
					},
				}
				Expect(ee.UntranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults).To(BeNil())

				err := ee.Translate()
				Expect(err).To(BeNil())
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("license"))
				Expect(ee.Results).NotTo(BeNil())
				Expect(len(ee.Results)).NotTo(Equal(0))

				err = ee.Translate()
				Expect(err).To(BeNil())
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("license"))
				Expect(ee.Results).NotTo(BeNil())
				Expect(len(ee.Results)).NotTo(Equal(0))
			})
		})

		g.Describe("Unmarshalling", func() {
			g.It("should populate results with untranslated result", func() {
				var ee Evaluation
				err := json.Unmarshal([]byte(sampleUntranslatedResults), &ee)

				Expect(err).To(BeNil())
				Expect(ee.TeamID).To(Equal("cuketest"))
				Expect(ee.UntranslatedResults).To(BeNil())
				tr := ee.TranslatedResults.Data.(LicenseResults)
				Expect(tr.Name).To(Equal("some license"))
			})

			g.It("should populate results with translated result", func() {
				var ee Evaluation
				err := json.Unmarshal([]byte(sampleTranslatedResults), &ee)

				Expect(err).To(BeNil())
				Expect(ee.TeamID).To(Equal("cuketest"))
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("community"))
			})
			g.It("should unmarshal a evaluation with a bad, list-ified community results member", func() {
				var ee Evaluation
				err := json.Unmarshal([]byte(badCommunityResults), &ee)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		g.Describe("Marshalling", func() {
			g.It("should output results with untranslated result", func() {
				s := &Evaluation{
					evaluation: &evaluation{
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
				s := &Evaluation{
					evaluation: &evaluation{
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
