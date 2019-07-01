package reports

import (
	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/rulesets"
)

// ProjectReport gives the details of a project including past analyses
type ProjectReport struct {
	*projects.Project
	RulesetName       string             `json:"ruleset_name"`
	AnalysisSummaries []analyses.Summary `json:"analysis_summaries"`
}

const (
	ReportGetProjectReportEndpoint = "v1/report/getProject"
)

// NewProjectReport takes a project and analysis summaries to return a
// constructed Project Report
func NewProjectReport(project *projects.Project, summaries []analyses.Summary) *ProjectReport {
	return &ProjectReport{
		Project:           project,
		AnalysisSummaries: summaries,
	}
}

// ProjectReports is used for getting a high level overview, returning a single
// analysis
type ProjectReports struct {
	*projects.Project
	RulesetName     string            `json:"ruleset_name"`
	AnalysisSummary *analyses.Summary `json:"analysis_summary"`
}

// NewProjectReports takes a project, analysis summary, and applied ruleset to
// create a summarized, high level report of a singular project. It returns this
// as a ProjectReports type.
func NewProjectReports(project *projects.Project, summary *analyses.Summary, appliedRuleset *rulesets.AppliedRulesetSummary) *ProjectReports {
	rulesetName := "N/A"
	if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
		rulesetName = appliedRuleset.RuleEvaluationSummary.RulesetName
	}

	if summary != nil {
		summary.AnalysisID = summary.ID
		summary.RulesetName = rulesetName
		summary.Trigger = "source commit"

		risk := "high"
		passed := false

		if appliedRuleset != nil {
			risk, passed = appliedRuleset.SummarizeEvaluation()
		}

		summary.Risk = risk
		summary.Passed = passed
	}

	pr := &ProjectReports{
		Project:         project,
		RulesetName:     rulesetName,
		AnalysisSummary: summary,
	}

	return pr
}
