package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestBuildsystemsDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Buildsystems", func() {
		g.It("should say so when no builds are detected", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"none detected"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce a chars digests when one build is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Compilers: []scans.Compiler{
						scans.Compiler{
							Name: "gcc",
						},
					},
				},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))
			Expect(ds[0].Title).To(Equal("compiler"))
			Expect(string(ds[0].Data)).To(Equal(`{"chars":"gcc"}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should produce a count digests when more than one build is present", func() {
			s := &scanner.ScanStatus{
				Status:  "finished",
				Message: "completed scan",
			}
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "buildsystems",
				Data: scans.BuildsystemResults{
					Compilers: []scans.Compiler{
						scans.Compiler{
							Name: "gcc",
						},
						scans.Compiler{
							Name: "ruby",
						},
						scans.Compiler{
							Name: "python",
						},
						scans.Compiler{
							Name: "rust",
						},
					},
				},
			}

			ds, err := buildsystemsDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))
			Expect(string(ds[0].Data)).To(Equal(`{"count":4}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})
	})
}
