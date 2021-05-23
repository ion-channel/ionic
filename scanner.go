package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/scanner"
)

type addScanRequest struct {
	TeamID    string               `json:"team_id"`
	ProjectID string               `json:"project_id"`
	ID        string               `json:"analysis_id"`
	Status    string               `json:"status"`
	Results   scanner.ExternalScan `json:"results"`
	Type      string               `json:"scan_type"`
}

// AnalyzeProject takes a projectID, teamID, and project branch, performs an
// analysis, and returns the result status or an error encountered by the API
func (ic *IonClient) AnalyzeProject(projectID, teamID, branch, token string) (*scanner.AnalysisStatus, error) {
	request := &scanner.AnalyzeRequest{}
	request.TeamID = teamID
	request.ProjectID = projectID

	if branch != "" {
		request.Branch = branch
	}

	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(scanner.ScannerAnalyzeProjectEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start analysis: %v", err.Error())
	}

	var a scanner.AnalysisStatus
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis status: %v", err.Error())
	}

	return &a, nil
}

// AnalyzeProjects takes a team ID to perform analyses on. An optional set of
// params can be provided to go along with the request, predominantly for
// internal purposes. It will return the IDs of the analyses created and any
// errors it encounters with the request.
func (ic *IonClient) AnalyzeProjects(teamID, token string, params *url.Values) ([]string, error) {
	request := &scanner.AnalyzeRequest{
		TeamID: teamID,
	}

	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(scanner.ScannerAnalyzeProjectEndpoint, token, params, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start analysis: %v", err.Error())
	}

	// We got a "team recently analyzed" message back
	if string(b) == "" {
		return nil, nil
	}

	var ids []string
	err = json.Unmarshal(b, &ids)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis status: %v", err.Error())
	}

	return ids, nil
}

//GetAnalysisStatus takes an analysisID, teamID, and projectID and returns the analysis status or an error encountered by the API
func (ic *IonClient) GetAnalysisStatus(analysisID, teamID, projectID, token string) (*scanner.AnalysisStatus, error) {
	params := &url.Values{}
	params.Set("id", analysisID)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(scanner.ScannerGetAnalysisStatusEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis status: %v", err.Error())
	}

	var a scanner.AnalysisStatus
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis status: %v", err.Error())
	}

	return &a, nil
}

//GetLatestAnalysisStatus takes a teamID, and projectID and returns the latest analysis status or an error encountered by the API
func (ic *IonClient) GetLatestAnalysisStatus(teamID, projectID, token string) (*scanner.AnalysisStatus, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, _, err := ic.Get(scanner.ScannerGetLatestAnalysisStatusEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a scanner.AnalysisStatus
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return &a, nil
}

//GetLatestAnalysisStatuses takes a teamID and returns the latest analysis statuses or an error encountered by the API
func (ic *IonClient) GetLatestAnalysisStatuses(teamID, token string) ([]scanner.AnalysisStatus, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, _, err := ic.Get(scanner.ScannerGetLatestAnalysisStatusesEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	var a []scanner.AnalysisStatus
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis: %v", err.Error())
	}

	return a, nil
}

//AddScanResult takes a scanResultID, teamID, projectID, status, scanType, and
//client provided scan results, and adds them to the returned project analysis
//or an error encountered by the API
func (ic *IonClient) AddScanResult(scanResultID, teamID, projectID, status, scanType, token string, scanResults scanner.ExternalScan) (*scanner.AnalysisStatus, error) {
	request := &addScanRequest{}
	request.ID = scanResultID
	request.TeamID = teamID
	request.ProjectID = projectID
	request.Results = scanResults
	request.Type = scanType

	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(scanner.ScannerAddScanEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start analysis: %v", err.Error())
	}

	var a scanner.AnalysisStatus
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to get analysis status: %v", err.Error())
	}

	return &a, nil
}

type projectStates struct {
	Filter string   `json:"filter"`
	IDs    []string `json:"ids"`
}

// GetProjectsStates takes a slice of project ids and an optional filter
// returns a slice of id's with each respected state
func (ic *IonClient) GetProjectsStates(ids []string, filter string, token string) ([]scanner.ProjectsStates, error) {
	ri := projectStates{
		IDs:    ids,
		Filter: filter,
	}

	b, err := json.Marshal(ri)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	r, err := ic.Post(scanner.ScannerGetProjectsStates, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project states: %v", err.Error())
	}

	var ps []scanner.ProjectsStates
	err = json.Unmarshal(r, &ps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}

	return ps, nil
}
