package reports

import (
	"testing"

	"github.com/ion-channel/ionic/analysis"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/rulesets"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestProjectReport(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Project Reports", func() {
		g.It("should return a new project reports", func() {
			expectedProjectID := "thisproject"
			expectedAnalysisID := "badanalysis"
			expectedRulesetName := "super-secure-ruleset"

			p := &projects.Project{
				ID: expectedProjectID,
			}
			s := &analysis.Summary{
				ID: expectedAnalysisID,
			}
			ar := &rulesets.AppliedRulesetSummary{
				RuleEvaluationSummary: &rulesets.RuleEvaluationSummary{
					RulesetName: expectedRulesetName,
				},
			}

			pr := NewProjectReports(p, s, ar)
			Expect(pr).NotTo(BeNil())

			Expect(pr.Project.ID).To(Equal(expectedProjectID))
			Expect(pr.RulesetName).To(Equal(expectedRulesetName))

			Expect(pr.AnalysisSummary.ID).To(Equal(expectedAnalysisID))
			Expect(pr.AnalysisSummary.AnalysisID).To(Equal(expectedAnalysisID))
			Expect(pr.AnalysisSummary.Trigger).To(Equal("source commit"))
			Expect(pr.AnalysisSummary.RulesetName).To(Equal(expectedRulesetName))
			Expect(pr.AnalysisSummary.Risk).To(Equal("high"))
			Expect(pr.AnalysisSummary.Passed).To(Equal(false))
		})
	})
}
