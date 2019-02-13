package digests

import (
	"testing"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestCommunityDigests(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

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
}
