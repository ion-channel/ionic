package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/scanner"
)

const (
	scannerAnalyzeProjectEndpoint    = "v1/scanner/analyzeProject"
	scannerGetAnalysisStatusEndpoint = "v1/scanner/getAnalysisStatus"
	scannerAddScanEndpoint           = "v1/scanner/addScanResult"
)

type analyzeRequest struct {
	TeamID    string `json:"team_id"`
	ProjectID string `json:"project_id"`
	Branch    string `json:"branch,omitempty"`
}

type addScanRequest struct {
	TeamID    string               `json:"team_id"`
	ProjectID string               `json:"project_id"`
	ID        string               `json:"analysis_id"`
	Status    string               `json:"status"`
	Results   scanner.ExternalScan `json:"results"`
	Type      string               `json:"scan_type"`
}

//AnalyzeProject takes a projectID, teamID, and project branch, performs an analysis, and returns the result status or an error encountered by the API
func (ic *IonClient) AnalyzeProject(projectID, teamID, branch, token string) (*scanner.AnalysisStatus, error) {
	request := &analyzeRequest{}
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
	b, err = ic.Post(scannerAnalyzeProjectEndpoint, token, nil, *buff, nil)
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

//GetAnalysisStatus takes an analysisID, teamID, and projectID and returns the analysis status or an error encountered by the API
func (ic *IonClient) GetAnalysisStatus(analysisID, teamID, projectID, token string) (*scanner.AnalysisStatus, error) {
	params := &url.Values{}
	params.Set("id", analysisID)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.Get(scannerGetAnalysisStatusEndpoint, token, params, nil, nil)
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
	b, err = ic.Post(scannerAddScanEndpoint, token, nil, *buff, nil)
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
