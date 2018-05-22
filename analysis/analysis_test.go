package analysis

import (
	"testing"

	"github.com/ion-channel/ionic/rulesets"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAnalysis(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("New Analysis Summary", func() {
		g.It("should return a new analysis summary", func() {
			expectedAnalysisID := "analysizing"

			a := &Analysis{
				ID: expectedAnalysisID,
			}
			ar := &rulesets.AppliedRulesetSummary{}

			s := NewSummary(a, ar)
			Expect(s.ID).To(Equal(expectedAnalysisID))
			Expect(s.AnalysisID).To(Equal(expectedAnalysisID))
			Expect(s.RulesetName).To(Equal("N/A"))
			Expect(s.Passed).To(Equal(false))
			Expect(s.Risk).To(Equal("high"))
		})
	})
}
