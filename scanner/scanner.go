package scanner

import (
	"time"
)

//ScanStatus needs a comment
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

//AnalysisStatus is a representation of an Ion Channel Analysis Status within the system
type AnalysisStatus struct {
	ID          string       `json:"id"`
	TeamID      string       `json:"team_id"`
	ProjectID   string       `json:"project_id"`
	BuildNumber string       `json:"build_number"`
	Message     string       `json:"message"`
	Branch      string       `json:"branch"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ScanStatus  []ScanStatus `json:"scan_status"`
}
