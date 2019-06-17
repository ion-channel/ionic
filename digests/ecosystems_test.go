package digests

import (
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
		g.It("should say so when no languages are detected", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "ecosystems",
				Data: scans.EcosystemResults{
					Ecosystems: map[string]int{},
				},
			}

			ds, err := ecosystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"none detected"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce a chars digests when one language is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "ecosystems",
				Data: scans.EcosystemResults{
					Ecosystems: map[string]int{
						"C#": 430056,
					},
				},
			}

			ds, err := ecosystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))
			Expect(ds[0].Title).To(Equal("language"))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"C#"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce a count digests when more than one language is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "ecosystems",
				Data: scans.EcosystemResults{
					Ecosystems: map[string]int{
						"Makefile": 4400,
						"Golang":   43000,
						"Shell":    43000,
						"C#":       43000,
					},
				},
			}

			ds, err := ecosystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))
			Expect(string(ds[0].Data)).To(Equal(`{"count":4}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})
	})
}
