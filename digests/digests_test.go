package digests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Digests", func() {
		g.Describe("Ecosystems", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{
					Status:  "finished",
					Message: "completed scan",
				}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "ecosystems",
					Data: scans.EcosystemResults{
						Ecosystems: map[string]int{
							"Makefile": 771,
							"C#":       430056,
							"Shell":    328,
						},
					},
				}

				ds, err := ecosystemsDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))
				Expect(string(ds[0].Data)).To(Equal(`{"list":["C#"]}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})

			g.It("should return a single dominant language", func() {
				languages := map[string]int{
					"Makefile": 100,
					"Go":       300000,
					"Ruby":     10000,
				}

				dom := getDominantLanguages(languages)
				Expect(len(dom)).To(Equal(1))
				Expect(dom[0]).To(Equal("Go"))
			})

			g.It("should return the top two languages if there is no majority", func() {
				languages := map[string]int{
					"Makefile": 100,
					"Go":       300,
					"Ruby":     300,
				}

				dom := getDominantLanguages(languages)
				Expect(len(dom)).To(Equal(2))
				Expect(fmt.Sprintf("%v", dom)).To(ContainSubstring("Go"))
				Expect(fmt.Sprintf("%v", dom)).To(ContainSubstring("Ruby"))
			})
		})

		g.Describe("Dependencies", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "dependency",
					Data: scans.DependencyResults{
						Dependencies: nil,
						Meta: scans.DependencyMeta{
							FirstDegreeCount:     2,
							NoVersionCount:       1,
							TotalUniqueCount:     115,
							UpdateAvailableCount: 2,
						},
					},
				}

				ds, err := dependencyDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(4))
				Expect(ds[0].Title).To(Equal("dependencies outdated"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())

				Expect(ds[1].Title).To(Equal("dependency no version specified"))
				Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
				Expect(ds[1].Pending).To(BeFalse())
				Expect(ds[1].Errored).To(BeFalse())

				Expect(ds[2].Title).To(Equal("direct dependencies"))
				Expect(string(ds[2].Data)).To(Equal(`{"count":2}`))
				Expect(ds[2].Pending).To(BeFalse())
				Expect(ds[2].Errored).To(BeFalse())

				Expect(ds[3].Title).To(Equal("transitive dependencies"))
				Expect(string(ds[3].Data)).To(Equal(`{"count":113}`))
				Expect(ds[3].Pending).To(BeFalse())
				Expect(ds[3].Errored).To(BeFalse())
			})
		})

		g.Describe("Vulnerabilities", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()

				var r scans.VulnerabilityResults
				b := []byte(`{"vulnerabilities":[{"id":910244496,"name":"eslint-scope","org":"eslint","version":"3.7.2","up":"","edition":"","aliases":null,"created_at":"2018-08-15T22:40:27.281Z","updated_at":"2018-08-15T22:41:03.309Z","title":"Eslint Eslint-scope 3.7.2","references":[],"part":"/a","language":"","external_id":"cpe:/a:eslint:eslint-scope:3.7.2","cpe23":null,"target_hw":null,"target_sw":null,"sw_edition":null,"other":null,"vulnerabilities":[{"id":267967698,"external_id":"NPM-673","title":"Malicious package","summary":"Version 3.7.2 of eslint-scope was published without authorization and was found to contain malicious code. This code would read the users .npmrc file and send any found authentication tokens to 2 remote servers.","score":"10.0","score_version":"","score_system":"CVSS","score_details":{"cvssv2":null,"cvssv3":{"vectorString":"","accessVector":"","accessComplexity":"","privilegesRequired":"","userInteraction":"","scope":"","confidentialityImpact":"","integrityImpact":"","availabilityImpact":"","baseScore":10.0,"baseSeverity":""}},"vector":"","access_complexity":"","vulnerability_authentication":null,"confidentiality_impact":"","integrity_impact":"","availability_impact":"","vulnerability_source":null,"assessment_check":null,"scanner":null,"recommendation":"","references":[{"type":"UNKNOWN","source":"NPM","url":"https://www.npmjs.com/advisories/673","text":"https://www.npmjs.com/advisories/673"},{"id":910244496}],"modified_at":"0001-01-01T00:00:00.000Z","published_at":"0001-01-01T00:00:00.000Z","created_at":"2018-08-15T22:56:54.294Z","updated_at":"2018-08-15T22:56:54.294Z","source":[{"id":3,"name":"NPM","description":"NPM Security advisories","created_at":"2018-08-15T22:03:09.285Z","updated_at":"2018-08-15T22:03:09.285Z","attribution":"NPM Term and Licenses https://www.npmjs.com/policies/terms","license":"https://www.npmjs.com/policies/open-source-terms","copyright_url":"https://www.npmjs.com/policies/dmca"}]}]}, {"id":910244496,"name":"eslint-scope","org":"eslint","version":"3.7.2","up":"","edition":"","aliases":null,"created_at":"2018-08-15T22:40:27.281Z","updated_at":"2018-08-15T22:41:03.309Z","title":"Eslint Eslint-scope 3.7.2","references":[],"part":"/a","language":"","external_id":"cpe:/a:eslint:eslint-scope:3.7.2","cpe23":null,"target_hw":null,"target_sw":null,"sw_edition":null,"other":null,"vulnerabilities":[{"id":267967698,"external_id":"NPM-673","title":"Malicious package","summary":"Version 3.7.2 of eslint-scope was published without authorization and was found to contain malicious code. This code would read the users .npmrc file and send any found authentication tokens to 2 remote servers.","score":"10.0","score_version":"","score_system":"CVSS","score_details":{"cvssv2":null,"cvssv3":{"vectorString":"","accessVector":"","accessComplexity":"","privilegesRequired":"","userInteraction":"","scope":"","confidentialityImpact":"","integrityImpact":"","availabilityImpact":"","baseScore":10.0,"baseSeverity":""}},"vector":"","access_complexity":"","vulnerability_authentication":null,"confidentiality_impact":"","integrity_impact":"","availability_impact":"","vulnerability_source":null,"assessment_check":null,"scanner":null,"recommendation":"","references":[{"type":"UNKNOWN","source":"NPM","url":"https://www.npmjs.com/advisories/673","text":"https://www.npmjs.com/advisories/673"},{"id":910244496}],"modified_at":"0001-01-01T00:00:00.000Z","published_at":"0001-01-01T00:00:00.000Z","created_at":"2018-08-15T22:56:54.294Z","updated_at":"2018-08-15T22:56:54.294Z","source":[{"id":3,"name":"NPM","description":"NPM Security advisories","created_at":"2018-08-15T22:03:09.285Z","updated_at":"2018-08-15T22:03:09.285Z","attribution":"NPM Term and Licenses https://www.npmjs.com/policies/terms","license":"https://www.npmjs.com/policies/open-source-terms","copyright_url":"https://www.npmjs.com/policies/dmca"}]}]}],"meta":{"vulnerability_count":2}}`)
				err := json.Unmarshal(b, &r)
				Expect(err).To(BeNil())

				e.TranslatedResults = &scans.TranslatedResults{
					Type: "vulnerability",
					Data: r,
				}

				ds, err := vulnerabilityDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(2))

				Expect(ds[0].Title).To(Equal("total vulnerabilities"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())

				Expect(ds[1].Title).To(Equal("unique vulnerability"))
				Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
				Expect(ds[1].Pending).To(BeFalse())
				Expect(ds[1].Errored).To(BeFalse())
			})
		})

		g.Describe("Viruses", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "virus",
					Data: scans.VirusResults{
						ScannedFiles:  622,
						InfectedFiles: 0,
					},
				}

				ds, err := virusDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(2))

				Expect(ds[0].Title).To(Equal("total files scanned"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":622}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())

				Expect(ds[1].Title).To(Equal("viruses found"))
				Expect(string(ds[1].Data)).To(Equal(`{"count":0}`))
				Expect(ds[1].Pending).To(BeFalse())
				Expect(ds[1].Errored).To(BeFalse())
			})

			g.It("should warn when no files are seen", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "virus",
					Data: scans.VirusResults{
						ScannedFiles:  0,
						InfectedFiles: 0,
					},
				}

				ds, err := virusDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(2))

				Expect(ds[0].Title).To(Equal("total files scanned"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":0}`))
				Expect(ds[0].Warning).To(BeTrue())
				Expect(ds[0].WarningMessage).To(Equal("no files were seen"))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})

			g.It("should warn when viruses are seen", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "virus",
					Data: scans.VirusResults{
						ScannedFiles:  0,
						InfectedFiles: 1,
					},
				}

				ds, err := virusDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(2))

				Expect(ds[1].Title).To(Equal("virus found"))
				Expect(string(ds[1].Data)).To(Equal(`{"count":1}`))
				Expect(ds[1].Warning).To(BeTrue())
				Expect(ds[1].WarningMessage).To(Equal("infected files were seen"))
				Expect(ds[1].Pending).To(BeFalse())
				Expect(ds[1].Errored).To(BeFalse())
			})
		})

		g.Describe("Community", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "community",
					Data: scans.CommunityResults{
						Committers: 123321,
					},
				}

				ds, err := communityDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("unique committers"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":123321}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})

			g.It("should warn about single committer repos", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "community",
					Data: scans.CommunityResults{
						Committers: 1,
					},
				}

				ds, err := communityDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("unique committer"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":1}`))
				Expect(ds[0].Warning).To(BeTrue())
				Expect(ds[0].WarningMessage).To(Equal("single committer repository"))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})
		})

		g.Describe("Licenses", func() {
			g.It("should produce digests with count when more than one", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "license",
					Data: scans.LicenseResults{
						License: &scans.License{
							Type: []scans.LicenseType{
								scans.LicenseType{Name: "apache-2.0"},
								scans.LicenseType{Name: "mit"},
							},
						},
					},
				}

				ds, err := licenseDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("licenses"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})

			g.It("should produce digests with license name when single", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "license",
					Data: scans.LicenseResults{
						License: &scans.License{
							Type: []scans.LicenseType{
								scans.LicenseType{Name: "mit"},
							},
						},
					},
				}

				ds, err := licenseDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("license"))
				Expect(string(ds[0].Data)).To(Equal(`{"chars":"mit"}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})

			g.It("should warn when no licenses are found", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "license",
					Data: scans.LicenseResults{
						License: &scans.License{
							Type: []scans.LicenseType{},
						},
					},
				}

				ds, err := licenseDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("licenses"))
				Expect(string(ds[0].Data)).To(Equal(`{"count":0}`))
				Expect(ds[0].Warning).To(BeTrue())
				Expect(ds[0].WarningMessage).To(Equal("no licenses found"))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})
		})

		g.Describe("Coverage", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "coverage",
					Data: scans.CoverageResults{
						Value: 93.88185664008439,
					},
				}

				ds, err := coveragDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("code coverage"))
				Expect(string(ds[0].Data)).To(Equal(`{"percent":93.88}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})
		})

		g.Describe("About YML", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "about_yml",
					Data: scans.AboutYMLResults{
						Valid: true,
					},
				}

				ds, err := aboutYMLDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("valid about yaml"))
				Expect(string(ds[0].Data)).To(Equal(`{"bool":true}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})
		})

		g.Describe("Difference", func() {
			g.It("should produce digests", func() {
				s := &scanner.ScanStatus{}
				e := scans.NewEval()
				e.TranslatedResults = &scans.TranslatedResults{
					Type: "difference",
					Data: scans.DifferenceResults{
						Difference: true,
					},
				}

				ds, err := differenceDigests(e, s)
				Expect(err).To(BeNil())
				Expect(len(ds)).To(Equal(1))

				Expect(ds[0].Title).To(Equal("difference detected"))
				Expect(string(ds[0].Data)).To(Equal(`{"bool":true}`))
				Expect(ds[0].Pending).To(BeFalse())
				Expect(ds[0].Errored).To(BeFalse())
			})
		})
	})
}
