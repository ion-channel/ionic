package rulesets

import (
	"time"

	"github.com/ion-channel/ionic/rules"
)

// CreateRuleSetOptions struct for creating a ruleset
type CreateRuleSetOptions struct {
	Name        string   `json:"name"`
	Description string   `json:"description" default:" "`
	TeamID      string   `json:"team_id"`
	RuleIDs     []string `json:"rule_ids"`
}

// RuleSet is a collection of rules
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
