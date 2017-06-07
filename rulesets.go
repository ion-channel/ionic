package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/rulesets"
)

const (
	getAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	getRuleSetEndpoint        = "v1/ruleset/getRuleset"
	getRuleSetsEndpoint       = "v1/ruleset/getRulesets"
)

func (ic *IonClient) GetRuleSet(id, teamID string) (*rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.get(getRuleSetEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	var rs rulesets.RuleSet
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	return &rs, nil
}

func (ic *IonClient) GetRuleSets(teamID string) ([]rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := ic.get(getRuleSetsEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	var rs []rulesets.RuleSet
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, fmt.Errorf("failed to get rulesets: %v", err.Error())
	}

	return rs, nil
}
