package scansummary

import (
	"encoding/json"
	"time"
)

// ScanSummary is an Ion Channel representation of the summary produced by an
// individual scan on a project.  It contains all the details the Ion Channel
// platform discovers for that scan.
type ScanSummary struct {
	ID          string          `json:"id"`
	TeamID      string          `json:"team_id"`
	ProjectID   string          `json:"project_id"`
	AnalysisID  string          `json:"analysis_id"`
	Summary     string          `json:"summary"`
	Results     json.RawMessage `json:"results"`
	UpdatedAt   time.Time       `json:"updated_at"`
	CreatedAt   time.Time       `json:"created_at"`
	Duration    float64         `json:"duration"`
	Passed      bool            `json:"passed"`
	Risk        string          `json:"risk"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
}
