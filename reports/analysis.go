package reports

import (
	"fmt"

	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/digests"
	"github.com/ion-channel/ionic/projects"
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
	Report   *analysisReport    `json:"report" xml:"report"`
}

type analysisReport struct {
	Project           *projects.Project               `json:"project" xml:"project"`
	ProjectRuleset    *rulesets.RuleSet               `json:"project_ruleset" xml:"project_ruleset"`
	Statuses          *scanner.AnalysisStatus         `json:"statuses" xml:"statuses"`
	Digests           []digests.Digest                `json:"digests" xml:"digests"`
	RulesetEvaluation *rulesets.AppliedRulesetSummary `json:"ruleset_evaluation" xml:"ruleset_evaluation"`
}

// NewAnalysisReport takes an Analysis and returns an initialized AnalysisReport
func NewAnalysisReport(status *scanner.AnalysisStatus, analysis *analyses.Analysis, project *projects.Project, projectRuleset *rulesets.RuleSet, appliedRuleset *rulesets.AppliedRulesetSummary) (*AnalysisReport, error) {
	if analysis == nil {
		analysis = &analyses.Analysis{
			ID:        status.ID,
			ProjectID: status.ProjectID,
			TeamID:    status.TeamID,
			Status:    status.Status,
		}
	}

	ds, err := digests.NewDigests(appliedRuleset, status.ScanStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get digests: %v", err.Error())
	}

	if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
		for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
			appliedRuleset.RuleEvaluationSummary.Ruleresults[i].Translate()
		}
	}

	ar := AnalysisReport{
		Analysis: analysis,
		Report: &analysisReport{
			Project:           project,
			ProjectRuleset:    projectRuleset,
			Statuses:          status,
			Digests:           ds,
			RulesetEvaluation: appliedRuleset,
		},
	}

	if ar.Analysis.Status == "finished" {
		ar.Analysis.Status = status.Status
	}

	return &ar, nil
}
