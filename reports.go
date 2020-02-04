package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/reports"
	"github.com/ion-channel/ionic/scanner"
)

//GetAnalysisReport takes an analysisID, teamID, projectID, and token. It
// returns the corresponding analysis report or an error encountered by the API
func (ic *IonClient) GetAnalysisReport(analysisID, teamID, projectID, token string) (*reports.AnalysisReport, error) {
	params := &url.Values{}
	params.Set("analysis_id", analysisID)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(reports.ReportGetAnalysisReportEndpoint, token, params, nil, nil)
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

//GetRawAnalysisReport takes an analysisID, teamID, projectID, and token. It
// returns the corresponding analysis report json or an error encountered by the
// API
func (ic *IonClient) GetRawAnalysisReport(analysisID, teamID, projectID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("analysis_id", analysisID)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(reports.ReportGetAnalysisReportEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis report: %v", err.Error())
	}

	return b, nil
}

//GetProjectReport takes a projectID, a teamID, and token. It returns the
// corresponding project report or an error encountered by the API
func (ic *IonClient) GetProjectReport(projectID, teamID, token string) (*reports.ProjectReport, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(reports.ReportGetProjectReportEndpoint, token, params, nil, nil)
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

//GetRawProjectReport takes a projectID, a teamID, and token. It returns the
// corresponding project report json or an error encountered by the API
func (ic *IonClient) GetRawProjectReport(projectID, teamID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(reports.ReportGetProjectReportEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project report: %v", err.Error())
	}

	return b, nil
}

// GetAnalysisNavigation takes an analysisID, teamID, projectID, and a token. It
// returns the related/tangential analyses to the analysis provided or returns
// any errors encountered with the API.
func (ic *IonClient) GetAnalysisNavigation(analysisID, teamID, projectID, token string) (*scanner.Navigation, error) {
	params := &url.Values{}
	params.Set("analysis_id", analysisID)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(reports.ReportGetAnalysisNavigationEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis navigation: %v", err.Error())
	}

	var n scanner.Navigation
	err = json.Unmarshal(b, &n)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return &n, nil
}
