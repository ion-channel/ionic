package reports

import (
	"testing"

	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestProjectReport(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Project Report", func() {
		g.It("should return a new project report", func() {
			expectedProjectID := "thisproject"
			expectedSummaryID := "thisanalysis"

			p := &projects.Project{
				ID: &expectedProjectID,
			}

			ss := []analyses.Summary{
				analyses.Summary{
					ID: expectedSummaryID,
				},
			}

			r := NewProjectReport(p, ss)
			Expect(*r.Project.ID).To(Equal(expectedProjectID))
			Expect(len(r.AnalysisSummaries)).To(Equal(1))
			Expect(r.AnalysisSummaries[0].ID).To(Equal(expectedSummaryID))
		})
	})

	g.Describe("Project Reports", func() {
		g.It("should return a new project reports", func() {
			expectedProjectID := "thisproject"
			expectedAnalysisID := "badanalysis"
			expectedRulesetName := "super-secure-ruleset"

			p := &projects.Project{
				ID: &expectedProjectID,
			}
			s := &analyses.Summary{
				ID: expectedAnalysisID,
			}
			ar := &rulesets.AppliedRulesetSummary{
				RuleEvaluationSummary: &rulesets.RuleEvaluationSummary{
					RulesetName: expectedRulesetName,
					Summary:     "pass",
				},
			}

			as := &scanner.AnalysisStatus{
				Status: scanner.AnalysisStatusFinished,
			}

			input := NewProjectReportsInput{
				p, s, ar, as,
			}
			pr := NewProjectReports(input)
			Expect(pr).NotTo(BeNil())
			Expect(pr.Status).To(Equal(ProjectStatusPassing))
			Expect(*pr.Project.ID).To(Equal(expectedProjectID))
			Expect(*pr.Project.ID).To(Equal(expectedProjectID))
			Expect(pr.RulesetName).To(Equal(expectedRulesetName))

			Expect(pr.AnalysisSummary.ID).To(Equal(expectedAnalysisID))
			Expect(pr.AnalysisSummary.AnalysisID).To(Equal(expectedAnalysisID))
			Expect(pr.AnalysisSummary.Trigger).To(Equal("source commit"))
			Expect(pr.AnalysisSummary.RulesetName).To(Equal(expectedRulesetName))
			Expect(pr.AnalysisSummary.Risk).To(Equal("low"))
			Expect(pr.AnalysisSummary.Passed).To(Equal(true))
		})
	})
}
