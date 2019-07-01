package scanner

import (
	"time"
)

const (
	ScannerAddScanEndpoint = "v1/scanner/addScanResult"
)

//ScanStatus identifies the state of a scan performed by the Ion system
type ScanStatus struct {
	ID               string    `json:"id"`
	AnalysisStatusID string    `json:"analysis_status_id"`
	ProjectID        string    `json:"project_id"`
	TeamID           string    `json:"team_id"`
	Message          string    `json:"message"`
	Name             string    `json:"name"`
	Read             string    `json:"read"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
