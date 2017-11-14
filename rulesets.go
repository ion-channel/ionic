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

//GetAppliedRuleSet takes a projectID, teamID, and analysisID and returns the corresponding applied ruleset summary or an error encountered by the API
func (ic *IonClient) GetAppliedRuleSet(projectID, teamID, analysisID string) (*rulesets.AppliedRulesetSummary, error) {
	params := &url.Values{}
	params.Set("project_id", projectID)
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

//GetRawAppliedRuleSet takes a projectID, teamID, analysisID, and page definition and returns the corresponding applied ruleset summary json or an error encountered by the API
func (ic *IonClient) GetRawAppliedRuleSet(projectID, teamID, analysisID string, page *pagination.Pagination) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("project_id", projectID)
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

//GetRuleSet takes a rule set ID and a teamID returns the corresponding rule set or an error encountered by the API
func (ic *IonClient) GetRuleSet(ruleSetID, teamID string) (*rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("id", ruleSetID)
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

//GetRuleSets takes a teamID and page definition and returns a collection of rule sets or an error encountered by the API
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
