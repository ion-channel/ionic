package rulesets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ion-channel/ionic/requests"
	"github.com/ion-channel/ionic/rules"
)

const (
	// CreateRuleSetEndpoint is a string representation of the current endpoint for creating ruleset
	CreateRuleSetEndpoint = "v1/ruleset/createRuleset"
	// GetAppliedRuleSetEndpoint is a string representation of the current endpoint for getting applied ruleset
	GetAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	// GetBatchAppliedRulesetEndpoint is a string representation of the current endpoint for getting batched applied rulesets
	GetBatchAppliedRulesetEndpoint = "v1/ruleset/getAppliedRulesets"
	// GetRuleSetEndpoint is a string representation of the current endpoint for getting ruleset
	GetRuleSetEndpoint = "v1/ruleset/getRuleset"
	// GetRuleSetsEndpoint is a string representation of the current endpoint for getting rulesets (plural)
	GetRuleSetsEndpoint = "v1/ruleset/getRulesets"
	//RulesetsGetRulesEndpoint is a string representation of the current endpoint for getting rules.
	RulesetsGetRulesEndpoint = "v1/ruleset/getRules"
	//RulesetsGetRulesetNames is a string representation of the current endpoint for getting ruleset names.
	RulesetsGetRulesetNames = "v1/ruleset/getRulesetNames"
)

// AppliedRulesetRequest represents a request for an applied ruleset result
type AppliedRulesetRequest struct {
	ProjectID  string `json:"project_id"`
	TeamID     string `json:"team_id"`
	SummaryID  string `json:"summary_id"`
	AnalysisID string `json:"analysis_id"`
}

// CreateRuleSetOptions struct for creating a ruleset
type CreateRuleSetOptions struct {
	Name        string   `json:"name"`
	Description string   `json:"description" default:" "`
	TeamID      string   `json:"team_id"`
	RuleIDs     []string `json:"rule_ids"`
}

// RuleSet is a collection of rules
type RuleSet struct {
	ID          string        `json:"id"`
	TeamID      string        `json:"team_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	RuleIDs     []string      `json:"rule_ids"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Rules       []rules.Rule  `json:"rules"`
	Deprecated  bool          `json:"has_deprecated_rules"`
	IsUsed      bool          `json:"has_projects_assigned"`
	DeletedAt   *sql.NullTime `json:"deleted_at,omitempty"`
	DeletedBy   string        `json:"deleted_by,omitempty"`
}

// NameForID represents the data object for ruleset name and its ID
type NameForID struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}

// String returns a JSON formatted string of the ruleset object
func (r RuleSet) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("failed to format ruleset: %v", err.Error())
	}
	return string(b)
}

// RuleSetExists takes a client, baseURL, ruleSetID, teamId and token string and checks against api to see if ruleset exists.
// It returns whether or not ruleset exists and any errors it encounters with the API.
// This is used internally in the SDK
func RuleSetExists(client *http.Client, baseURL *url.URL, ruleSetID, teamID, token string) (bool, error) {
	params := &url.Values{}
	params.Set("id", ruleSetID)
	params.Set("team_id", teamID)

	err := requests.Head(client, baseURL, GetRuleSetEndpoint, token, params, nil, nil)

	if err != nil {
		if strings.Contains(err.Error(), "(404)") {
			return false, nil
		}

		return false, fmt.Errorf("failed to request ruleset: %v", err.Error())
	}

	return true, nil
}
