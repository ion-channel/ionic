package scanner

import (
	"time"
)

// Delivery represents the delivery information of a singular artifact
// associated with an analysis status
type Delivery struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"team_id"`
	ProjectID   string    `json:"project_id"`
	AnalysisID  string    `json:"analysis_id"`
	Destination string    `json:"destination"`
	Status      string    `json:"status"`
	Filename    string    `json:"filename"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
