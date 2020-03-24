package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDifferencesDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Difference", func() {
		g.It("should produce digests", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "difference",
				Data: scans.DifferenceResults{
					Difference: true,
				},
			}

			ds, err := differenceDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			Expect(ds[0].Title).To(Equal("difference detected"))
			Expect(string(ds[0].Data)).To(Equal(`{"bool":true}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})
	})
}
