package rulesets

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestAppliedRulesets(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Applied Ruleset Summary", func() {
		g.It("should return low risk and passed if the evaluation summary passed", func() {
			ar := AppliedRulesetSummary{
				RuleEvaluationSummary: &RuleEvaluationSummary{
					Summary: "passed",
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

	g.Describe("Rule Evaluation Summary", func() {
		g.Describe("Passing", func() {
			g.It("should return true if passing", func() {
				res := &RuleEvaluationSummary{
					Summary: "pass",
				}

				Expect(res.Passed()).To(Equal(true))
			})

			g.It("should not care about case", func() {
				res := &RuleEvaluationSummary{
					Summary: "Pass",
				}

				Expect(res.Passed()).To(Equal(true))

				res.Summary = "pASS"
				Expect(res.Passed()).To(Equal(true))
			})

			g.It("should handle multiple cases of the phrase passing", func() {
				res := &RuleEvaluationSummary{
					Summary: "pass",
				}

				Expect(res.Passed()).To(Equal(true))

				res.Summary = "passed"
				Expect(res.Passed()).To(Equal(true))

				res.Summary = "passing"
				Expect(res.Passed()).To(Equal(true))
			})
		})

		g.Describe("Failing", func() {
			g.It("should return false if the rule evaluation summary failed", func() {
				res := &RuleEvaluationSummary{}
				Expect(res.Passed()).To(Equal(false))

				res.Summary = "foo"
				Expect(res.Passed()).To(Equal(false))
			})
		})
	})
}
