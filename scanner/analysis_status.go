package scanner

import (
	"time"
)

const (
	// AnalysisStatusAccepted denotes a request for analysis has been
	// accepted and queued
	AnalysisStatusAccepted = "accepted"
	// AnalysisStatusErrored denotes a request for analysis has errored during
	// the run, the message field will have more details
	AnalysisStatusErrored = "errored"
	// AnalysisStatusFinished denotes a request for analysis has been
	// completed, view the passed field from an Analysis and the scan details for
	// more information
	AnalysisStatusFinished = "finished"
	// AnalysisStatusFailed denotes a request for analysis has failed to
	// run, the message field will have more details
	AnalysisStatusFailed = "failed"
	// AnalysisStatusStarted denotes a request for analysis has been
	// accepted and has begun
	AnalysisStatusStarted = "started"
)

// AnalysisStatus is a representation of an Ion Channel Analysis Status within the system
type AnalysisStatus struct {
	ID         string       `json:"id"`
	TeamID     string       `json:"team_id"`
	ProjectID  string       `json:"project_id"`
	Message    string       `json:"message"`
	Branch     string       `json:"branch"`
	Status     string       `json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	ScanStatus []ScanStatus `json:"scan_status"`
	Deliveries []Delivery   `json:"deliveries"`
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
