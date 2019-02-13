package digests

import (
	"fmt"
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestEcosystemsDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

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

			ds, err := ecosystemsDigests(s, e)
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
}
