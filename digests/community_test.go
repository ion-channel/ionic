package digests

import (
	"fmt"
	"testing"
	"time"

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
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			committedAt := time.Now()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "community",
				Data: scans.CommunityResults{
					Committers:  123321,
					CommittedAt: committedAt,
				},
			}

			ds, err := communityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))

			Expect(ds[0].Title).To(Equal("unique committers"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":123321}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("days since last commit"))
			res := fmt.Sprintf("{\"chars\":\"%s\"}", "0")
			Expect(string(ds[1].Data)).To(Equal(res))
		})

		g.It("should warn about single committer repos", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "community",
				Data: scans.CommunityResults{
					Committers: 1,
				},
			}

			ds, err := communityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))

			Expect(ds[0].Title).To(Equal("unique committer"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":1}`))
			Expect(ds[0].Warning).To(BeTrue())
			Expect(ds[0].WarningMessage).To(Equal("single committer repository"))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
		})

		g.It("should respond properly a zeroed time", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "community",
				Data: scans.CommunityResults{
					Committers:  123321,
					CommittedAt: time.Time{},
				},
			}

			ds, err := communityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))

			Expect(ds[0].Title).To(Equal("unique committers"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":123321}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("days since last commit"))
			res := fmt.Sprintf("{\"chars\":\"%s\"}", "N/A")
			Expect(string(ds[1].Data)).To(Equal(res))
		})

		g.It("should respond properly to a committed_at time in the past", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			now := time.Now()
			committedAt := now.AddDate(0, 0, -15)
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "community",
				Data: scans.CommunityResults{
					Committers:  123321,
					CommittedAt: committedAt,
				},
			}

			ds, err := communityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(2))

			Expect(ds[0].Title).To(Equal("unique committers"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":123321}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())

			Expect(ds[1].Title).To(Equal("days since last commit"))
			fmt.Printf(fmt.Sprintf("%s", ds[1]))
			res := fmt.Sprintf("{\"chars\":\"%s\"}", "15")
			Expect(string(ds[1].Data)).To(Equal(res))
		})

		g.It("should properly evaluate a community scan evaluation without days since last activity rule", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			now := time.Now()
			committedAt := now.AddDate(0, 0, -15)
			e.RuleID = "2981e1b0-0c8f-0137-8fe7-186590d3c755"
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "community",
				Data: scans.CommunityResults{
					Committers:  123321,
					CommittedAt: committedAt,
				},
			}

			ds, err := communityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			Expect(ds[0].Title).To(Equal("unique committers"))
			Expect(string(ds[0].Data)).To(Equal(`{"count":123321}`))
			Expect(ds[0].Pending).To(BeFalse())
			Expect(ds[0].Errored).To(BeFalse())
			Expect(ds[0].Evaluated).To(BeTrue())

			// Expect(ds[1].Title).To(Equal("days since last commit"))
			// fmt.Printf(fmt.Sprintf("%s", ds[1]))
			// res := fmt.Sprintf("{\"chars\":\"%s\"}", "15")
			// Expect(string(ds[1].Data)).To(Equal(res))
			// Expect(ds[1].Evaluated).To(BeFalse())
		})

		g.It("should properly evaluate a community scan evaluation with days since last activity rule", func() {
			s := &scanner.ScanStatus{}
			s.Status = scanner.ScanStatusFinished
			e := scans.NewEval()
			now := time.Now()
			e.RuleID = "efcb4ae5-ff36-413a-962b-3f4d2170be2a"
			committedAt := now.AddDate(0, 0, -15)
			e.TranslatedResults = &scans.TranslatedResults{
				Type: "community",
				Data: scans.CommunityResults{
					Committers:  123321,
					CommittedAt: committedAt,
				},
			}

			ds, err := communityDigests(s, e)
			Expect(err).To(BeNil())
			Expect(len(ds)).To(Equal(1))

			// Expect(ds[0].Title).To(Equal("unique committers"))
			// Expect(string(ds[0].Data)).To(Equal(`{"count":123321}`))
			// Expect(ds[0].Pending).To(BeFalse())
			// Expect(ds[0].Errored).To(BeFalse())
			// Expect(ds[0].Evaluated).To(BeFalse())

			Expect(ds[0].Title).To(Equal("days since last commit"))
			fmt.Printf(fmt.Sprintf("%s", ds[0]))
			res := fmt.Sprintf("{\"chars\":\"%s\"}", "15")
			Expect(string(ds[0].Data)).To(Equal(res))
			Expect(ds[0].Evaluated).To(BeTrue())
		})
	})
}
