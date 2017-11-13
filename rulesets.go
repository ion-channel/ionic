package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/rulesets"
)

const (
	getAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	getRuleSetEndpoint        = "v1/ruleset/getRuleset"
	getRuleSetsEndpoint       = "v1/ruleset/getRulesets"
)

//GetAppliedRuleSet needs a comment
func (ic *IonClient) GetAppliedRuleSet(id, teamID, analysisID string) (*rulesets.AppliedRulesetSummary, error) {
	params := &url.Values{}
	params.Set("project_id", id)
	params.Set("team_id", teamID)
	if analysisID != "" {
		params.Set("analysis_id", analysisID)
	}

	b, err := ic.get(getAppliedRuleSetEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get applied ruleset summary: %v", err.Error())
	}

	var s rulesets.AppliedRulesetSummary
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal applied ruleset summary: %v", err.Error())
	}

	return &s, nil
}

//GetRawAppliedRuleSet needs a comment
func (ic *IonClient) GetRawAppliedRuleSet(id, teamID, analysisID string, page *pagination.Pagination) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("project_id", id)
	params.Set("team_id", teamID)
	if analysisID != "" {
		params.Set("analysis_id", analysisID)
	}

	b, err := ic.get(getAppliedRuleSetEndpoint, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get applied rulesets: %v", err.Error())
	}

	return b, nil
}

//GetRuleSet needs a comment
func (ic *IonClient) GetRuleSet(id, teamID string) (*rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.get(getRuleSetEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get ruleset: %v", err.Error())
	}

	var rs rulesets.RuleSet
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal ruleset: %v", err.Error())
	}

	return &rs, nil
}

//GetRuleSets needs a comment
func (ic *IonClient) GetRuleSets(teamID string, page *pagination.Pagination) ([]rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := ic.get(getRuleSetsEndpoint, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	var rs []rulesets.RuleSet
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal rulesets: %v", err.Error())
	}

	return rs, nil
}
