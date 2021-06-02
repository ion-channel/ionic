package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/analyses"
	"net/url"

	"github.com/ion-channel/ionic/reports"
	"github.com/ion-channel/ionic/requests"
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

// GetExportedProjectsData takes slice of project ids, team id, and token
// returns slice of exported data for the requested projects
func (ic *IonClient) GetExportedProjectsData(ids []string, teamID, token string) (*reports.ExportedData, error) {
	p := requests.ByIDsAndTeamID{
		TeamID: teamID,
		IDs:    ids,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(reports.ReportGetExportedDataEndpoint, token, nil, *bytes.NewBuffer(b), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to request exported data: %v", err.Error())
	}

	var ed reports.ExportedData
	err = json.Unmarshal(r, &ed)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal exported projects data response: %v", err.Error())
	}

	return &ed, nil
}

// GetExportedVulnerabilityData takes slice of project ids, team id, and token
// returns slice of exported vulnerability data for the requested projects
func (ic *IonClient) GetExportedVulnerabilityData(ids []string, teamID, token string) (*[]analyses.VulnerabilityExportData, error) {
	p := requests.ByIDsAndTeamID{
		TeamID: teamID,
		IDs:    ids,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(reports.ReportGetExportedVulnerabilityDataEndpoint, token, nil, *bytes.NewBuffer(b), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to request exported data: %v", err.Error())
	}

	var ed []analyses.VulnerabilityExportData
	err = json.Unmarshal(r, &ed)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal exported projects data response: %v", err.Error())
	}

	return &ed, nil
}

// GetSBOM takes slice of project ids, team id, SBOM format, and token.
// Returns one or more SBOMs for the requested project(s).
func (ic *IonClient) GetSBOM(ids []string, teamID string, options reports.SBOMExportOptions, token string) (string, error) {
	body := requests.ByIDsAndTeamID{
		TeamID: teamID,
		IDs:    ids,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	params := options.Params()

	r, err := ic.Post(reports.ReportGetSBOMEndpoint, token, params, *bytes.NewBuffer(b), nil)

	if err != nil {
		return "", fmt.Errorf("failed to request SBOM: %v", err.Error())
	}

	return string(r), nil
}
