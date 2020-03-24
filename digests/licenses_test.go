package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestLicensesDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Licenses", func() {
		g.It("should produce digests with license name when only one", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.AnalysisStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "license",
				Data: scans.LicenseResults{
					License: &scans.License{
						Type: []scans.LicenseType{
							scans.LicenseType{Name: "apache-2.0"},
						},
					},
				},
			}

			ds, err := licenseDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			Expect(ds[0].Title).To(Equal("license found"))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"apache-2.0"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce digests with count when more than one", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.AnalysisStatusFinished
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

			ds, err := licenseDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			Expect(ds[0].Title).To(Equal("licenses found"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":2}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should warn when no licenses are found", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.AnalysisStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "license",
				Data: scans.LicenseResults{
					License: &scans.License{
						Type: []scans.LicenseType{},
					},
				},
			}

			ds, err := licenseDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			Expect(ds[0].Title).To(Equal("licenses found"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":0}`))
			Expect(ds[0].Warning).To(BeTrue())
			Expect(ds[0].WarningMessage).To(Equal("no licenses found"))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})
	})
}
