package scanner

import (
	"time"
)

const (
	// DeliveryStatusErrored denotes a request for delivery has errored during
	// the run, the message field will have more details
	DeliveryStatusErrored = "errored"
	// DeliveryStatusFinished denotes a request for delivery has been
	// completed, view the passed field from an Delivery and the scan details for
	// more information
	DeliveryStatusFinished = "finished"
	// DeliveryStatusFailed denotes a request for delivery has failed to
	// run, the message field will have more details
	DeliveryStatusFailed = "failed"
	// DeliveryStatusNotConfigured denotes a request for delivery has been
	// rejected due to no delivery details
	DeliveryStatusNotConfigured = "not_configured"
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
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
