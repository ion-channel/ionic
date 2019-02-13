package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAboutYMLDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

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
}
