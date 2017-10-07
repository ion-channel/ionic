package scans

import (
	"time"
)

// ScanSummary is an Ion Channel representation of the summary produced by an
// individual scan on a project.  It contains all the details the Ion Channel
// platform discovers for that scan.
type ScanSummary struct {
	ID          string    `json:"id" xml:"id"`
	TeamID      string    `json:"team_id" xml:"team_id"`
	ProjectID   string    `json:"project_id" xml:"project_id"`
	AnalysisID  string    `json:"analysis_id" xml:"analysis_id"`
	Summary     string    `json:"summary" xml:"summary"`
	Results     *Results  `json:"results,omitempty" xml:"results,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" xml:"updated_at"`
	CreatedAt   time.Time `json:"created_at" xml:"created_at"`
	Duration    float64   `json:"duration" xml:"duration"`
	Passed      bool      `json:"passed" xml:"passed"`
	Risk        string    `json:"risk" xml:"risk"`
	Name        string    `json:"name" xml:"name"`
	Description string    `json:"description" xml:"description"`
	Type        string    `json:"type" xml:"type"`
}
