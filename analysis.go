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

// GetAnalysis returns the specified analysis
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

// GetLatestAnalysisSumary returns the latest analysis for the specified project
func (ic *IonClient) GetLatestAnalysisSumary(teamID, projectID string) (*analysis.AnalysisSummary, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(analysisGetLatestAnalysisSummaryEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	var a analysis.AnalysisSummary
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest analysis: %v", err.Error())
	}

	return &a, nil
}
