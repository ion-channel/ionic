package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/analysis"
)

const (
	analysisGetAnalysisEndpoint              = "v1/animal/getAnalysis"
	analysisGetLatestAnalysisSummaryEndpoint = "v1/animal/getLatestAnalysisSummary"
)

// GetAnalysis takes an analysis ID, team ID, and project ID.  It returns the
// analysis found.  If the analysis is not found it will return an error, and
// will return an error for any other API issues it encounters.
func (ic *IonClient) GetAnalysis(id, teamID, projectID string) (*analysis.Analysis, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(analysisGetAnalysisEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a analysis.Analysis
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return &a, nil
}

// GetRawAnalysis takes an analysis ID, team ID, and project ID, and returns the
// raw JSON from the API.  It returns an error for any API issues it encounters.
func (ic *IonClient) GetRawAnalysis(id, teamID, projectID string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(analysisGetAnalysisEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return b, nil
}

// GetLatestAnalysisSummary takes a team ID and project ID, and returns the
// latest analysis summary for the project. It returns an error for any API
// issues it encounters.
func (ic *IonClient) GetLatestAnalysisSummary(teamID, projectID string) (*analysis.Summary, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(analysisGetLatestAnalysisSummaryEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	var a analysis.Summary
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	return &a, nil
}

// GetRawLatestAnalysisSummary takes a team ID and project ID, and returns the
// raw JSON from the API.  It returns an error for any API issues it encounters.
func (ic *IonClient) GetRawLatestAnalysisSummary(teamID, projectID string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(analysisGetLatestAnalysisSummaryEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	return b, nil
}
