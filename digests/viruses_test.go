package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestVirusesDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

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

			ds, err := virusDigests(s, e)
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

			ds, err := virusDigests(s, e)
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

			ds, err := virusDigests(s, e)
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
}
