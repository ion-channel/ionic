package scanner

import (
	"time"
)

const (
	// AnalysisStatusAccepted denotes a request for analysis has been
	// accepted and queued
	AnalysisStatusAccepted = "accepted"
	// AnalysisStatusFinished denotes a request for analysis has been
	// completed, view the passed field from an Analysis and the scan details for
	// more information
	AnalysisStatusFinished = "finished"
	// AnalysisStatusFailed denotes a request for analysis has failed to
	// run, the message field will have more details
	AnalysisStatusFailed = "failed"
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
