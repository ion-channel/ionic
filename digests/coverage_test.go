package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestCoverageDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

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

			ds, err := coveragDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			Expect(ds[0].Title).To(Equal("code coverage"))
			Expect(string(ds[0].Data)).To(Equal(`{"percent":93.88}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})
	})
}
