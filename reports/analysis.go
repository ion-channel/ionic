package reports

import (
	"strings"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/analysis"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scans"
	"github.com/ion-channel/ionic/tags"
)

// AnalysisReport is a Ion Channel representation of a report output from a
// given analysis
type AnalysisReport struct {
	*analysis.Analysis
	RulesetName   string             `json:"ruleset_name" xml:"ruleset_name"`
	Passed        bool               `json:"passed" xml:"passed"`
	Aliases       []aliases.Alias    `json:"aliases"`
	Tags          []tags.Tag         `json:"tags"`
	Trigger       string             `json:"trigger" xml:"trigger"`
	Risk          string             `json:"risk" xml:"risk"`
	Summary       string             `json:"summary" xml:"summary"`
	ScanSummaries []scans.Evaluation `json:"scan_summaries" xml:"scan_summaries"`
}

// NewAnalysisReport takes an Analysis and returns an initialized AnalysisReport
func NewAnalysisReport(analysis *analysis.Analysis, project *projects.Project, appliedRuleset *rulesets.AppliedRulesetSummary) (*AnalysisReport, error) {
	ar := AnalysisReport{
		Analysis: analysis,
		Trigger:  "source commit",
		Risk:     "high",
	}

	// Project Details
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

		ar.ScanSummaries = appliedRuleset.RuleEvaluationSummary.Ruleresults
	}

	return &ar, nil
}
