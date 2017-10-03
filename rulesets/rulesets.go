package rulesets

import (
	"time"

	"github.com/ion-channel/ionic/rules"
	"github.com/ion-channel/ionic/scans"
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
	ID            string              `json:"id"`
	TeamID        string              `json:"team_id"`
	ProjectID     string              `json:"project_id"`
	BuildNumber   string              `json:"build_number"`
	Name          string              `json:"name"`
	Text          string              `json:"text"`
	Type          string              `json:"type"`
	Source        string              `json:"source"`
	Branch        string              `json:"branch"`
	Description   string              `json:"description"`
	Status        string              `json:"status"`
	RulesetID     string              `json:"ruleset_id"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	Duration      float64             `json:"duration"`
	TriggerHash   string              `json:"trigger_hash"`
	TriggerText   string              `json:"trigger_text"`
	TriggerAuthor string              `json:"trigger_author"`
	ScanSummaries []scans.ScanSummary `json:"scan_summaries"`
	RulesetName   string              `json:"ruleset_name"`
	Risk          string              `json:"risk"`
	Summary       string              `json:"summary"`
	Trigger       string              `json:"trigger"`
	Passed        bool                `json:"passed"`
}
