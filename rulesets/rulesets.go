package rulesets

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ion-channel/ionic/requests"
	"github.com/ion-channel/ionic/rules"
)

const (
	// CreateRuleSetEndpoint is a string representation of the current endpoint for creating ruleset
	CreateRuleSetEndpoint = "v1/ruleset/createRuleset"
	// GetAppliedRuleSetEndpoint is a string representation of the current endpoint for getting applied ruleset
	GetAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	// GetRuleSetEndpoint is a string representation of the current endpoint for getting ruleset
	GetRuleSetEndpoint = "v1/ruleset/getRuleset"
	// GetRuleSetsEndpoint is a string representation of the current endpoint for getting rulesets (plural)
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

// RuleSetExists takes a ruleSetID, teamId and token string and checks against api to see if ruleset exists.
// It returns whether or not ruleset exists and any errors it encounters with the API.
// This is used internally in the SDK
func RuleSetExists(client *http.Client, baseURL *url.URL, ruleSetID, teamID, token string) (bool, error) {
	params := &url.Values{}
	params.Set("id", ruleSetID)
	params.Set("team_id", teamID)

	err := requests.Head(client, baseURL, GetRuleSetEndpoint, token, params, nil, nil)
	if err != nil {
		return false, fmt.Errorf("failed to find ruleset: %v", err.Error())
	}

	return true, nil
}
