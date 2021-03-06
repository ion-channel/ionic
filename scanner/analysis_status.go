package scanner

import (
	"time"
)

const (
	// AnalysisStatusQueued denotes a request for analysis has been
	// accepted and queued
	AnalysisStatusQueued = "queued"
	// AnalysisStatusErrored denotes a request for analysis has errored during
	// the run, the message field will have more details
	AnalysisStatusErrored = "errored"
	// AnalysisStatusFinished denotes a request for analysis has been
	// completed, view the passed field from an Analysis and the scan details for
	// more information
	AnalysisStatusFinished = "finished"
	// AnalysisStatusPassed denotes a request for analysis has failed to
	// run, the message field will have more details
	AnalysisStatusPassed = "passed"
	// AnalysisStatusFailed denotes a request for analysis has been
	// accepted and has failed
	AnalysisStatusFailed = "failed"
	// AnalysisStatusAnalyzing denotes a request for analysis has been
	// accepted and has begun
	AnalysisStatusAnalyzing = "analyzing"
)

const (
	// ScannerAnalyzeProjectEndpoint is a string representation of the current endpoint for analyzing project
	ScannerAnalyzeProjectEndpoint = "v1/scanner/analyzeProject"
	// ScannerGetAnalysisStatusEndpoint is a string representation of the current endpoint for getting analysis status
	ScannerGetAnalysisStatusEndpoint = "v1/scanner/getAnalysisStatus"
	// ScannerGetLatestAnalysisStatusEndpoint is a string representation of the current endpoint for getting latest analysis status
	ScannerGetLatestAnalysisStatusEndpoint = "v1/scanner/getLatestAnalysisStatus"
	// ScannerGetLatestAnalysisStatusesEndpoint is a string representation of the current endpoint for getting latest analysis statuses
	ScannerGetLatestAnalysisStatusesEndpoint = "v1/scanner/getLatestAnalysisStatuses"
)

// AnalysisStatus is a representation of an Ion Channel Analysis Status within the system
type AnalysisStatus struct {
	ID                  string              `json:"id"`
	TeamID              string              `json:"team_id"`
	ProjectID           string              `json:"project_id"`
	Message             string              `json:"message"`
	Branch              string              `json:"branch"`
	Status              string              `json:"status"`
	UnreachableError    bool                `json:"unreachable_error"`
	AnalysisEventSource string              `json:"analysis_event_src"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	ScanStatus          []ScanStatus        `json:"scan_status"`
	Deliveries          map[string]Delivery `json:"deliveries"`
}

// Done indicates an analyse has stopped processing
func (a *AnalysisStatus) Done() bool {
	return a.Status == AnalysisStatusErrored ||
		a.Status == AnalysisStatusFailed ||
		a.Status == AnalysisStatusFinished
}

// Navigation represents a navigational meta data reference to given analysis
type Navigation struct {
	Analysis       *AnalysisStatus `json:"analysis"`
	LatestAnalysis *AnalysisStatus `json:"latest_analysis"`
}
