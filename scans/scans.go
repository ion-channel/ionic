package scans

import (
	"time"
)

// Scan represents the data collected from an individual scan on a project.
type Scan struct {
	ID          string             `json:"id"`
	TeamID      string             `json:"team_id"`
	ProjectID   string             `json:"project_id"`
	AnalysisID  string             `json:"analysis_id"`
	Summary     string             `json:"summary"`
	Results     *TranslatedResults `json:"results"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Duration    float64            `json:"duration"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}
