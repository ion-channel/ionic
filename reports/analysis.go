package reports

import (
	"fmt"
	"strings"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/digests"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
	"github.com/ion-channel/ionic/tags"
)

// AnalysisReport is a Ion Channel representation of a report output from a
// given analysis
type AnalysisReport struct {
	*analyses.Analysis
	Trigger  string                  `json:"trigger" xml:"trigger"`
	Statuses *scanner.AnalysisStatus `json:"statuses" xml:"statuses"`
	Summary  string                  `json:"summary" xml:"summary"`
	Digests  []digests.Digest        `json:"digests" xml:"digests"`

	// Evaluation Details
	RulesetName   string             `json:"ruleset_name" xml:"ruleset_name"`
	Passed        bool               `json:"passed" xml:"passed"`
	Risk          string             `json:"risk" xml:"risk"`
	ScanSummaries []scans.Evaluation `json:"scan_summaries" xml:"scan_summaries"`
	Evaluations   []scans.Evaluation `json:"evaluations" xml:"evaluations"`

	// Project Details
	Active   bool            `json:"active"`
	Monitor  bool            `json:"should_monitor"`
	Private  bool            `json:"private"`
	POCName  string          `json:"poc_name"`
	POCEmail string          `json:"poc_email"`
	Aliases  []aliases.Alias `json:"aliases"`
	Tags     []tags.Tag      `json:"tags"`
}

// NewAnalysisReport takes an Analysis and returns an initialized AnalysisReport
func NewAnalysisReport(status *scanner.AnalysisStatus, analysis *analyses.Analysis, project *projects.Project, appliedRuleset *rulesets.AppliedRulesetSummary) (*AnalysisReport, error) {
	if analysis == nil {
		analysis = &analyses.Analysis{
			ID:        status.ID,
			ProjectID: status.ProjectID,
			TeamID:    status.TeamID,
			Status:    status.Status,
		}
	}

	ar := AnalysisReport{
		Analysis: analysis,
		Trigger:  "source commit",
		Risk:     "high",
		Statuses: status,
	}

	// Project Details
	ar.Active = project.Active
	ar.Monitor = project.Monitor
	ar.Private = project.Private
	ar.POCName = project.POCName
	ar.POCEmail = project.POCEmail
	ar.Aliases = project.Aliases
	ar.Tags = project.Tags

	// RulesetEval Details
	if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
		ar.RulesetName = appliedRuleset.RuleEvaluationSummary.RulesetName

		if strings.ToLower(appliedRuleset.RuleEvaluationSummary.Summary) == "pass" {
			ar.Risk = "low"
			ar.Passed = true
		}

		for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
			appliedRuleset.RuleEvaluationSummary.Ruleresults[i].Translate()
		}

		// TODO: Remove ScanSummaries field
		ar.ScanSummaries = appliedRuleset.RuleEvaluationSummary.Ruleresults
		ar.Evaluations = appliedRuleset.RuleEvaluationSummary.Ruleresults
	}

	ds, err := digests.NewDigests(appliedRuleset, status.ScanStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get digests: %v", err.Error())
	}

	ar.Digests = ds

	return &ar, nil
}
