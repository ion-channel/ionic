package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/requests"
)

// GetAnalysis takes an analysis ID, team ID, project ID, and token.  It returns the
// analysis found.  If the analysis is not found it will return an error, and
// will return an error for any other API issues it encounters.
func (ic *IonClient) GetAnalysis(id, teamID, projectID, token string) (*analyses.Analysis, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetAnalysisEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a analyses.Analysis
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal analysis: %v", err.Error())
	}

	return &a, nil
}

// GetLatestAnalysis takes a team ID, project ID, and token.  It returns the
// latest analysis found.  If the analysis is not found it will return an error, and
// will return an error for any other API issues it encounters.
func (ic *IonClient) GetLatestAnalysis(teamID, projectID, token string) (*analyses.Analysis, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetLatestAnalysisEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a analyses.Analysis
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal analysis: %v", err.Error())
	}

	return &a, nil
}

// GetAnalyses takes a team ID, project ID, and token. It returns a slice of
// analyses for the project or an error for any API issues it encounters.
func (ic *IonClient) GetAnalyses(teamID, projectID, token string, page *pagination.Pagination) ([]analyses.Analysis, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetAnalysesEndpoint, token, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get analyses: %v", err.Error())
	}

	var as []analyses.Analysis
	err = json.Unmarshal(b, &as)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal analyses: %v", err.Error())
	}

	return as, nil
}

// GetLatestPublicAnalysis takes a project ID and branch.  It returns the
// analysis found.  If the analysis is not found it will return an error, and
// will return an error for any other API issues it encounters.
func (ic *IonClient) GetLatestPublicAnalysis(projectID, branch string) (*analyses.Analysis, error) {
	params := &url.Values{}
	params.Set("project_id", projectID)
	params.Set("branch", branch)

	b, _, err := ic.Get(analyses.AnalysisGetLatestPublicAnalysisEndpoint, "", params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a analyses.Analysis
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal analysis: %v", err.Error())
	}

	return &a, nil
}

// GetPublicAnalysis takes an analysis ID.  It returns the
// analysis found.  If the analysis is not found it will return an error, and
// will return an error for any other API issues it encounters.
func (ic *IonClient) GetPublicAnalysis(id string) (*analyses.Analysis, error) {
	params := &url.Values{}
	params.Set("id", id)

	b, _, err := ic.Get(analyses.AnalysisGetPublicAnalysisEndpoint, "", params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a analyses.Analysis
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal analysis: %v", err.Error())
	}

	return &a, nil
}

// GetRawAnalysis takes an analysis ID, team ID, project ID, and token.  It returns the
// raw JSON from the API.  It returns an error for any API issues it encounters.
func (ic *IonClient) GetRawAnalysis(id, teamID, projectID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetAnalysisEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return b, nil
}

// GetRawAnalyses takes a team ID, project ID, and token. It returns the raw
// JSON from the API. It returns an error for any API issue it encounters.
func (ic *IonClient) GetRawAnalyses(teamID, projectID, token string, page *pagination.Pagination) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetAnalysesEndpoint, token, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return b, nil
}

// GetLatestAnalysisIDs takes a team ID, project ID(s), and token. It returns the
// latest analysis IDs for the project as a map in the form map[project_id] = latest_analysis_id
// It returns an error for any API issues it encounters.
func (ic *IonClient) GetLatestAnalysisIDs(teamID string, projectIDs []string, token string) (*map[string]string, error) {
	ri := requests.ByIDsAndTeamID{
		IDs:    projectIDs,
		TeamID: teamID,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(analyses.AnalysisGetLatestAnalysisIDsEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis IDs: %v", err.Error())
	}

	a := make(map[string]string)
	err = json.Unmarshal(r, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis IDs: %v", err.Error())
	}

	return &a, nil
}

// GetLatestAnalysisSummary takes a team ID, project ID, and token. It returns the
// latest analysis summary for the project. It returns an error for any API
// issues it encounters.
func (ic *IonClient) GetLatestAnalysisSummary(teamID, projectID, token string) (*analyses.Summary, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetLatestAnalysisSummaryEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	var a analyses.Summary
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	return &a, nil
}

// GetRawLatestAnalysisSummary takes a team ID, project ID, and token. It returns the
// raw JSON from the API.  It returns an error for any API issues it encounters.
func (ic *IonClient) GetRawLatestAnalysisSummary(teamID, projectID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(analyses.AnalysisGetLatestAnalysisSummaryEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	return b, nil
}

// GetAnalysesExportData takes team id and a slice of analysis ids
// returns a slice of analyses exported data
func (ic *IonClient) GetAnalysesExportData(teamID string, ids []string, token string) ([]analyses.ExportData, error) {
	ri := requests.ByIDsAndTeamID{
		IDs:    ids,
		TeamID: teamID,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(analyses.AnalysisGetAnalysesExportData, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project states: %v", err.Error())
	}

	var ps []analyses.ExportData
	err = json.Unmarshal(r, &ps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return ps, nil
}

// GetAnalysesVulnerabilityExportData takes team id and a slice of analysis ids
// returns a slice of vulnerability analyses exported data
func (ic *IonClient) GetAnalysesVulnerabilityExportData(teamID string, ids []string, token string) ([]analyses.VulnerabilityExportData, error) {
	ri := requests.ByIDsAndTeamID{
		IDs:    ids,
		TeamID: teamID,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(analyses.AnalysisGetAnalysesVulnerabilityExportData, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project states: %v", err.Error())
	}

	var vulnerabilities []analyses.VulnerabilityExportData
	err = json.Unmarshal(r, &vulnerabilities)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return vulnerabilities, nil
}
