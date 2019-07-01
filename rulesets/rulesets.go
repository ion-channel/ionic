package rulesets

import (
	"time"

	"github.com/ion-channel/ionic/rules"
)

const (
	// CreateRuleSetEndpoint is a string representation of the current endpoint for getting projects
	CreateRuleSetEndpoint = "v1/ruleset/createRuleset"
	// GetAppliedRuleSetEndpoint is a string representation of the current endpoint for getting projects
	GetAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	// GetRuleSetEndpoint is a string representation of the current endpoint for getting projects
	GetRuleSetEndpoint = "v1/ruleset/getRuleset"
	// GetRuleSetsEndpoint is a string representation of the current endpoint for getting projects
	GetRuleSetsEndpoint = "v1/ruleset/getRulesets"
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
