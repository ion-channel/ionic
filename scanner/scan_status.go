package scanner

import (
	"strings"
	"time"
)

const (
	// ScannerAddScanEndpoint is a string representation of the current endpoint for scanner add scan
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

// Errored encapsulates the concerns of whether a ScanStatus is in an errored
// state or not. It returns true or false based on whether the ScanStatus is
// errored.
func (s *ScanStatus) Errored() bool {
	if strings.ToLower(s.Status) == "errored" {
		return true
	}

	return false
}
