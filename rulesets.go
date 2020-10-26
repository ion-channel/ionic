package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/requests"
	"github.com/ion-channel/ionic/rulesets"
)

// CreateRuleSet Creates a project attached to the team id supplied
func (ic *IonClient) CreateRuleSet(opts rulesets.CreateRuleSetOptions, token string) (*rulesets.RuleSet, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal project: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(rulesets.CreateRuleSetEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ruleset: %v", err.Error())
	}

	var p rulesets.RuleSet
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to create ruleset: %v", err.Error())
	}

	return &p, nil
}

//GetAppliedRuleSet takes a projectID, teamID, and analysisID and returns the corresponding applied ruleset summary or an error encountered by the API
func (ic *IonClient) GetAppliedRuleSet(projectID, teamID, analysisID, token string) (*rulesets.AppliedRulesetSummary, error) {
	params := &url.Values{}
	params.Set("project_id", projectID)
	params.Set("team_id", teamID)
	if analysisID != "" {
		params.Set("analysis_id", analysisID)
	}

	b, _, err := ic.Get(rulesets.GetAppliedRuleSetEndpoint, token, params, nil, nil)
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

//GetAppliedRuleSets takes a slice of AppliedRulesetRequest and returns their applied ruleset results, omitting any not found
func (ic *IonClient) GetAppliedRuleSets(appliedRequestBatch []*rulesets.AppliedRulesetRequest, token string) (*[]rulesets.AppliedRulesetSummary, error) {
	b, err := json.Marshal(appliedRequestBatch)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal project: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	r, err := ic.Post(rulesets.GetBatchAppliedRulesetEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get applied ruleset summary: %v", err.Error())
	}

	var s []rulesets.AppliedRulesetSummary
	err = json.Unmarshal(r, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal applied ruleset summary: %v", err.Error())
	}

	return &s, nil
}

//GetRawAppliedRuleSet takes a projectID, teamID, analysisID, and page definition and returns the corresponding applied ruleset summary json or an error encountered by the API
func (ic *IonClient) GetRawAppliedRuleSet(projectID, teamID, analysisID, token string, page *pagination.Pagination) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("project_id", projectID)
	params.Set("team_id", teamID)
	if analysisID != "" {
		params.Set("analysis_id", analysisID)
	}

	b, _, err := ic.Get(rulesets.GetAppliedRuleSetEndpoint, token, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get applied rulesets: %v", err.Error())
	}

	return b, nil
}

//GetRuleSet takes a rule set ID and a teamID returns the corresponding rule set or an error encountered by the API
func (ic *IonClient) GetRuleSet(ruleSetID, teamID, token string) (*rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("id", ruleSetID)
	params.Set("team_id", teamID)

	b, _, err := ic.Get(rulesets.GetRuleSetEndpoint, token, params, nil, nil)
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
func (ic *IonClient) GetRuleSets(teamID, token string, page *pagination.Pagination) ([]rulesets.RuleSet, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, _, err := ic.Get(rulesets.GetRuleSetsEndpoint, token, params, nil, page)
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

// RuleSetExists takes a ruleSetID, teamId and token string and checks against api to see if ruleset exists.
// It returns whether or not ruleset exists and any errors it encounters with the API.
func (ic *IonClient) RuleSetExists(ruleSetID, teamID, token string) (bool, error) {
	return rulesets.RuleSetExists(ic.client, ic.baseURL, ruleSetID, teamID, token)
}

//GetProjectPassFailHistory takes a project id and returns a daily history of pass/fail statuses
func (ic *IonClient) GetProjectPassFailHistory(projectID, token string) ([]rulesets.ProjectPassFailHistory, error) {
	params := &url.Values{}
	params.Set("project_id", projectID)

	b, _, err := ic.Get(rulesets.GetProjectHistoryEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project history: %v", err.Error())
	}

	var ph []rulesets.ProjectPassFailHistory
	err = json.Unmarshal(b, &ph)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal ruleset: %v", err.Error())
	}

	return ph, nil
}

//GetRulesetNames takes slice of ids and returns the ruleset names with the ids
func (ic *IonClient) GetRulesetNames(ids []string, token string) ([]rulesets.NameForID, error) {
	byIDs := requests.ByIDs{
		IDs: ids,
	}

	b, err := json.Marshal(byIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	r, err := ic.Post(rulesets.RulesetsGetRulesetNames, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get applied ruleset summary: %v", err.Error())
	}

	var s []rulesets.NameForID
	err = json.Unmarshal(r, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal applied ruleset summary: %v", err.Error())
	}

	return s, nil
}

// GetAnalysesStatuses takes a team id, slice of project ids and token
// returns a slice of project ids, analysis ids, and status
func (ic *IonClient) GetAnalysesStatuses(teamID string, ids []string, token string) ([]rulesets.Status, error) {
	body := requests.ByIDsAndTeamID{
		TeamID: teamID,
		IDs:    ids,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(rulesets.GetRulesetAnalysesStatuses, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analyses statuses: %v", err.Error())
	}

	var statuses []rulesets.Status
	err = json.Unmarshal(r, &statuses)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return statuses, nil
}
