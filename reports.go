package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/reports"
)

const (
	reportGetAnalysisReportEndpoint = "v1/report/getAnalysis"
	reportGetProjectReportEndpoint  = "v1/report/getProject"
)

func (ic *IonClient) GetAnalysisReport(id, teamID, projectID string) (*reports.AnalysisReport, error) {
	params := &url.Values{}
	params.Set("analysis_id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(reportGetAnalysisReportEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis report: %v", err.Error())
	}

	var r reports.AnalysisReport
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal analysis report: %v", err.Error())
	}

	return &r, nil
}

func (ic *IonClient) GetRawAnalysisReport(id, teamID, projectID string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("analysis_id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(reportGetAnalysisReportEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis report: %v", err.Error())
	}

	return b, nil
}

func (ic *IonClient) GetProjectReport(id, teamID string) (*reports.ProjectReport, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", id)

	b, err := ic.get(reportGetProjectReportEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project report: %v", err.Error())
	}

	var r reports.ProjectReport
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal project report: %v", err.Error())
	}

	return &r, nil
}

func (ic *IonClient) GetRawProjectReport(id, teamID string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", id)

	b, err := ic.get(reportGetProjectReportEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project report: %v", err.Error())
	}

	return b, nil
}
