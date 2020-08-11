package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/portfolios"
)

// GetVulnerabilityStats takes slice of project ids and token and returns vulnerability stats and any errors
func (ic *IonClient) GetVulnerabilityStats(ids []string, token string) (*portfolios.VulnerabilityStat, error) {
	p := struct {
		Ids []string `json:"ids"`
	}{
		ids,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.VulnerabilityStatsEndpoint, token, nil, *bytes.NewBuffer(b), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to request vulnerability list: %v", err.Error())
	}

	var vs portfolios.VulnerabilityStat
	err = json.Unmarshal(r, &vs)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal vunlerability stats response: %v", err.Error())
	}

	return &vs, nil
}

// GetRawVulnerabilityList gets a raw response from the API
func (ic *IonClient) GetRawVulnerabilityList(ids []string, listType, limit, token string) ([]byte, error) {
	p := portfolios.PortfolioListParams{
		ListType: listType,
		Ids:      ids,
		Limit:    limit,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	resp, err := ic.Post(portfolios.VulnerabilityListEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request vulnerability list: %v", err.Error())
	}

	return resp, nil
}

// GetRawVulnerabilityMetrics takes slice of strings (project ids), metric, and token
// and returns raw response from the API
func (ic *IonClient) GetRawVulnerabilityMetrics(ids []string, metric, token string) ([]byte, error) {
	mb := portfolios.MetricsBody{
		Metric:     metric,
		ProjectIDs: ids,
	}

	b, err := json.Marshal(mb)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	resp, err := ic.Post(portfolios.VulnerabilityMetricsEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request vulnerability metrics: %v", err.Error())
	}

	return resp, nil
}

// GetPortfolioPassFailSummary takes project ids (slice of strings) and a token (string) and returns a status summary
func (ic *IonClient) GetPortfolioPassFailSummary(ids []string, token string) (*portfolios.PortfolioPassingFailingSummary, error) {
	ri := portfolios.PortfolioRequestedIds{
		IDs: ids,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.PortfolioPassFailSummaryEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request portfolio status summary: %v", err.Error())
	}

	var ps portfolios.PortfolioPassingFailingSummary
	err = json.Unmarshal(r, &ps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return &ps, nil
}

// GetPortfolioStartedErroredSummary takes project ids (slice of strings) and a token (string) and returns PortfolioStartedErroredSummary
func (ic *IonClient) GetPortfolioStartedErroredSummary(ids []string, token string) (*portfolios.PortfolioStartedErroredSummary, error) {
	ri := portfolios.PortfolioRequestedIds{
		IDs: ids,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.PortfolioStartedErroredSummaryEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request portfolio status summary: %v", err.Error())
	}

	var ps portfolios.PortfolioStartedErroredSummary
	err = json.Unmarshal(r, &ps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return &ps, nil
}

// GetPortfolioAffectedProjects takes team id, external id, and a token (string) and returns a slice of affected projects
func (ic *IonClient) GetPortfolioAffectedProjects(teamID, externalID, token string) ([]portfolios.AffectedProject, error) {
	params := &url.Values{}
	params.Set("id", teamID)
	params.Set("external_id", externalID)

	r, _, err := ic.Get(portfolios.PortfolioGetAffectedProjectIdsEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request portfolio affected projects: %v", err.Error())
	}

	var aps []portfolios.AffectedProject
	err = json.Unmarshal(r, &aps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return aps, nil
}

// GetPortfolioAffectedProjectsInfo takes team id, external id, and a token (string) and returns a slice of affected projects
func (ic *IonClient) GetPortfolioAffectedProjectsInfo(ids []string, token string) ([]portfolios.AffectedProject, error) {
	ri := portfolios.PortfolioRequestedIds{
		IDs: ids,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.PortfolioGetAffectedProjectsInfoEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request portfolio affected projects info: %v", err.Error())
	}

	var aps []portfolios.AffectedProject
	err = json.Unmarshal(r, &aps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return aps, nil
}

// GetDependencyStats takes slice of project ids and token and returns dependency stat and any errors
func (ic *IonClient) GetDependencyStats(ids []string, token string) (*portfolios.DependencyStat, error) {
	p := struct {
		Ids []string `json:"ids"`
	}{
		ids,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.DependencyStatsEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request dependency list: %v", err.Error())
	}

	var ds portfolios.DependencyStat
	err = json.Unmarshal(r, &ds)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal dependency stats response: %v", err.Error())
	}

	return &ds, nil
}

// GetRawDependencyList gets a raw response from the API
func (ic *IonClient) GetRawDependencyList(ids []string, listType, limit, token string) ([]byte, error) {
	p := portfolios.PortfolioListParams{
		ListType: listType,
		Ids:      ids,
		Limit:    limit,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	resp, err := ic.Post(portfolios.DependencyListEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to request dependency list: %v", err.Error())
	}

	return resp, nil
}

// GetProjectsStatusHistory takes slice of project ids and token and returns list of status history for projects
func (ic *IonClient) GetProjectsStatusHistory(ids []string, token string) ([]portfolios.StatusesHistory, error) {
	p := struct {
		Ids []string `json:"ids"`
	}{
		ids,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(portfolios.RulesetsGetStatusesHistoryEndpoint, token, nil, *bytes.NewBuffer(b), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to request status history: %v", err.Error())
	}

	var sh []portfolios.StatusesHistory
	err = json.Unmarshal(r, &sh)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal status history response: %v", err.Error())
	}

	return sh, nil
}

// GetMttr takes team id and optional project ID and returns the mttr for project
// If project id is not given, it will return mttr of all active projects on the team
func (ic *IonClient) GetMttr(teamID, projectID string, token string) (*portfolios.Mttr, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	r, _, err := ic.Get(portfolios.ReportsGetMttrEndpoint, token, params, nil, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to request mttr: %v", err.Error())
	}

	var mttr portfolios.Mttr
	err = json.Unmarshal(r, &mttr)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal mttr response: %v", err.Error())
	}

	return &mttr, nil
}
