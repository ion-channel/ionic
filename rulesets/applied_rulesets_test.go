package rulesets

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAppliedRulesets(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Applied Ruleset Summary", func() {
		g.It("should return low risk and passed if the evaluation summary passed", func() {
			ar := AppliedRulesetSummary{
				RuleEvaluationSummary: &RuleEvaluationSummary{
					Summary: "pass",
				},
			}

			r, p := ar.SummarizeEvaluation()
			Expect(r).To(Equal("low"))
			Expect(p).To(Equal(true))
		})

		g.It("should return high risk and failed if the evaluation summary failed", func() {
			ar := AppliedRulesetSummary{
				RuleEvaluationSummary: &RuleEvaluationSummary{},
			}

			r, p := ar.SummarizeEvaluation()
			Expect(r).To(Equal("high"))
			Expect(p).To(Equal(false))
		})

		g.It("should return high risk and failed if the evaluation summary doesn't exist", func() {
			ar := AppliedRulesetSummary{}

			r, p := ar.SummarizeEvaluation()
			Expect(r).To(Equal("high"))
			Expect(p).To(Equal(false))
		})
	})
}
