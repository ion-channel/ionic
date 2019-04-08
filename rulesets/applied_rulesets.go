package rulesets

import (
	"time"

	"github.com/ion-channel/ionic/scans"
)

// AppliedRulesetSummary identifies the rule set applied to an analysis of a
// project and the result of their evaluation
type AppliedRulesetSummary struct {
	ProjectID             string                 `json:"project_id"`
	TeamID                string                 `json:"team_id"`
	AnalysisID            string                 `json:"analysis_id"`
	RuleEvaluationSummary *RuleEvaluationSummary `json:"rule_evaluation_summary"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// RuleEvaluationSummary represents the ruleset and the scans that were
// evaluated with the ruleset
type RuleEvaluationSummary struct {
	RulesetName string             `json:"ruleset_name"`
	Summary     string             `json:"summary"`
	Risk        string             `json:"risk"`
	Passed      bool               `json:"passed"`
	Ruleresults []scans.Evaluation `json:"ruleresults"`
}
