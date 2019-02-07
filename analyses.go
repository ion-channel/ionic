package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/analyses"
	"github.com/ion-channel/ionic/pagination"
)

const (
	analysisGetAnalysisEndpoint              = "v1/animal/getAnalysis"
	analysisGetAnalysesEndpoint              = "v1/animal/getAnalyses"
	analysisGetLatestAnalysisSummaryEndpoint = "v1/animal/getLatestAnalysisSummary"
	analysisGetPublicAnalysisEndpoint        = "v1/animal/getPublicAnalysis"
	analysisGetLatestPublicAnalysisEndpoint  = "v1/animal/getLatestPublicAnalysisSummary"
)

// GetAnalysis takes an analysis ID, team ID, project ID, and token.  It returns the
// analysis found.  If the analysis is not found it will return an error, and
// will return an error for any other API issues it encounters.
func (ic *IonClient) GetAnalysis(id, teamID, projectID, token string) (*analyses.Analysis, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.Get(analysisGetAnalysisEndpoint, token, params, nil, nil)
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

	b, err := ic.Get(analysisGetAnalysesEndpoint, token, params, nil, page)
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

	b, err := ic.Get(analysisGetLatestPublicAnalysisEndpoint, "", params, nil, nil)
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

	b, err := ic.Get(analysisGetPublicAnalysisEndpoint, "", params, nil, nil)
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

	b, err := ic.Get(analysisGetAnalysisEndpoint, token, params, nil, nil)
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

	b, err := ic.Get(analysisGetAnalysesEndpoint, token, params, nil, page)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return b, nil
}

// GetLatestAnalysisSummary takes a team ID, project ID, and token. It returns the
// latest analysis summary for the project. It returns an error for any API
// issues it encounters.
func (ic *IonClient) GetLatestAnalysisSummary(teamID, projectID, token string) (*analyses.Summary, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.Get(analysisGetLatestAnalysisSummaryEndpoint, token, params, nil, nil)
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

	b, err := ic.Get(analysisGetLatestAnalysisSummaryEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	return b, nil
}
