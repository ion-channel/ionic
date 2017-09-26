package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/scanner"
)

const (
	scannerAnalyzeProjectEndpoint     = "v1/scanner/analyzeProject"
	scannerGetAnalysisStatusEndpoint = "v1/scanner/getAnalysisStatus"
	scannerAddScanEndpoint = "v1/scanner/addScanResult"
)

type analyzeRequest struct {
	TeamID    string `json:"team_id"`
	ProjectID string `json:"project_id"`
}

type addScanRequest struct {
	TeamID    string               `json:"team_id"`
	ProjectID string               `json:"project_id"`
	ID        string               `json:"analysis_id"`
	Status    string               `json:"status"`
	Results   scanner.ExternalScan `json:"results"`
	Type      string               `json:"scan_type"`
}

func (ic *IonClient) AnalyzeProject(teamID, projectID string) (*scanner.AnalysisStatus, error) {
	request := &analyzeRequest{}
	request.TeamID = teamID
	request.ProjectID = projectID

	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.post(scannerAnalyzeProjectEndpoint, nil, *buff, nil)
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

func (ic *IonClient) GetAnalysisStatus(id, teamID, projectID string) (*scanner.AnalysisStatus, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	b, err := ic.get(scannerGetAnalysisStatusEndpoint, params, nil, nil)
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

func (ic *IonClient) AddScanResult(id, teamID, projectID, status, scanType string, scanResults scanner.ExternalScan) (*scanner.AnalysisStatus, error) {
	request := &addScanRequest{}
	request.ID = id
	request.TeamID = teamID
	request.ProjectID = projectID
	request.Results = scanResults
	request.Type = scanType

	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.post(scannerAddScanEndpoint, nil, *buff, nil)
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
