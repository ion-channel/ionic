package reports

import (
	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"
)

const (
	// ReportGetAnalysisReportEndpoint is a string representation of the current endpoint for getting report analysis
	ReportGetAnalysisReportEndpoint = "v1/report/getAnalysis"
	// ReportGetAnalysisNavigationEndpoint is a string representation of the current endpoint for getting report analysis navigation
	ReportGetAnalysisNavigationEndpoint = "v1/report/getAnalysisNav"
)

// AnalysisReport is a Ion Channel representation of a report output from a
// given analysis
type AnalysisReport struct {
	Analysis *analyses.Analysis `json:"analysis" xml:"analysis"`
}

// NewAnalysisReport takes an Analysis and returns an initialized AnalysisReport
func NewAnalysisReport(status *scanner.AnalysisStatus, analysis *analyses.Analysis, appliedRuleset *rulesets.AppliedRulesetSummary, experimental bool) (*AnalysisReport, error) {
	if analysis == nil {
		analysis = &analyses.Analysis{
			ID:        status.ID,
			ProjectID: status.ProjectID,
			TeamID:    status.TeamID,
			Status:    status.Status,
		}
	}

	if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
		for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
			appliedRuleset.RuleEvaluationSummary.Ruleresults[i].Translate()
		}
	}

	ar := AnalysisReport{
		Analysis: analysis,
	}

	switch status.Status {
	case scanner.AnalysisStatusErrored:
		analysis.Status = scanner.AnalysisStatusErrored
	case scanner.AnalysisStatusFinished:
		analysis.Status = scanner.AnalysisStatusFailed
		if appliedRuleset.RuleEvaluationSummary.Passed {
			analysis.Status = scanner.AnalysisStatusPassed
		}
	case scanner.AnalysisStatusQueued:
		analysis.Status = scanner.AnalysisStatusQueued
	case scanner.AnalysisStatusAnalyzing:
		analysis.Status = scanner.AnalysisStatusAnalyzing
	}

	return &ar, nil
}
