package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const (
	getAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	getRuleSetEndpoint        = "v1/ruleset/getRuleset"
	getRuleSetsEndpoint       = "v1/ruleset/getRulesets"
)

type RuleSet struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"team_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RuleIDs     []string  `json:"rule_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rules       []Rule    `json:"rules"`
}

func (ic *IonClient) GetRuleSet(id, teamID string) (*RuleSet, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.get(getRuleSetEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	var rs RuleSet
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	return &rs, nil
}

func (ic *IonClient) GetRuleSets(teamID string) ([]RuleSet, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := ic.get(getRuleSetsEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	var rs []RuleSet
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	return rs, nil
}
