package rulesets

import (
	"strings"
	"time"

	"github.com/ion-channel/ionic/scans"
)

const (
	// GetProjectHistoryEndpoint is a string representation of the current endpoint for getting a project's history
	GetProjectHistoryEndpoint = "v1/ruleset/getProjectHistory"
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

// SummarizeEvaluation returns the calculated risk and passing values for the
// AppliedRulsetSummary. Only if the RuleEvalutionSummary has passed, will it
// return low risk and passing.
func (ar *AppliedRulesetSummary) SummarizeEvaluation() (string, bool) {
	if ar.RuleEvaluationSummary != nil && strings.ToLower(ar.RuleEvaluationSummary.Summary) == "pass" {
		return "low", true
	}

	return "high", false
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

// ProjectPassFailHistory represents a pass failing history of a project's evaluation
type ProjectPassFailHistory struct {
	TeamID        string    `json:"team_id"`
	ProjectID     string    `json:"project_id"`
	AnalysisID    string    `json:"analysis_id"`
	Status        bool      `json:"pass"`
	FailCount     int       `json:"fail_count"`
	CreatedAt     time.Time `json:"created_at"`
	StatusFlipped bool      `json:"status_flipped"`
}

// ProjectRulesetHistory represents history of a project's ruleset changing
type ProjectRulesetHistory struct {
	OldRulesetID   string    `json:"old_ruleset_id"`
	OldRulesetName string    `json:"old_ruleset_name"`
	NewRulesetID   string    `json:"new_ruleset_id"`
	NewRulesetName string    `json:"new_ruleset_name"`
	UserID         string    `json:"user_id"`
	UserName       string    `json:"user_name"`
	CreatedAt      time.Time `json:"created_at"`
}

// ProjectAudit represents a projects history agregated by date
type ProjectAudit struct {
	Date           time.Time               `json:"date"`
	PassFail       *ProjectPassFailHistory `json:"project_pass_fail,omitempty"`
	RulesetHistory []ProjectRulesetHistory `json:"ruleset_history,omitempty"`
}
