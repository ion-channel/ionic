package rulesets

import (
	"encoding/json"
	"time"

	"github.com/ion-channel/ionic/rules"
)

type RuleSet struct {
	ID          string       `json:"id"`
	TeamID      string       `json:"team_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	RuleIDs     []string     `json:"rule_ids"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Rules       []rules.Rule `json:"rules"`
}

type AppliedRulesetSummary struct {
	ProjectID             string `json:"project_id"`
	TeamID                string `json:"team_id"`
	AnalysisID            string `json:"analysis_id"`
	RuleEvaluationSummary struct {
		RulesetName           string `json:"ruleset_name"`
		Summary     string `json:"summary"`
		Ruleresults []struct {
			ID          string          `json:"id"`
			AnalysisID  string          `json:"analysis_id"`
			TeamID      string          `json:"team_id"`
			ProjectID   string          `json:"project_id"`
			Description string          `json:"description"`
			Name        string          `json:"name"`
			Summary     string          `json:"summary"`
			CreatedAt   time.Time       `json:"created_at"`
			UpdatedAt   time.Time       `json:"updated_at"`
			Results     json.RawMessage `json:"results"`
			Duration    float64         `json:"duration"`
			Passed      bool            `json:"passed"`
			Risk        string          `json:"risk"`
			Type        string          `json:"type"`
		} `json:"ruleresults"`
	} `json:"rule_evaluation_summary"`
	RuleEvalCreatedAt time.Time `json:"rule_eval_created_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
