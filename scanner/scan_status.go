package scanner

import (
	"strings"
	"time"
)

const (
	// ScannerAddScanEndpoint is a string representation of the current endpoint for scanner add scan
	ScannerAddScanEndpoint = "v1/scanner/addScanResult"
	// ScannerGetProjectsStates is a string representation of the current endpoint getting requested projects scan states
	ScannerGetProjectsStates = "v1/scanner/getProjectsStates"

	// ScanStatusErrored denotes a request for scan has errored during
	// the run, the message field will have more details
	ScanStatusErrored = "errored"
	// ScanStatusFinished denotes a request for scan has been
	// completed, view the passed field from an Scan and the scan details for
	// more information
	ScanStatusFinished = "finished"
	// ScanStatusStarted denotes a request for scan has been
	// accepted and has begun
	ScanStatusStarted = "started"
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

// ProjectsStates identfies the state of a scan for requested projects
type ProjectsStates struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	AnalysisID string `json:"analysis_id"`
}

// Errored encapsulates the concerns of whether a ScanStatus is in an errored
// state or not. It returns true or false based on whether the ScanStatus is
// errored.
func (s *ScanStatus) Errored() bool {
	if strings.ToLower(s.Status) == ScanStatusErrored {
		return true
	}

	return false
}
