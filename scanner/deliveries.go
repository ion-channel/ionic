package scanner

import (
	"time"
)

const (
	// DeliveryFinished is the status returned when a delivery has been
	// successfully handed off to the delivery location
	DeliveryFinished = "delivery_finished"
	// DeliveryFailed is the status returned when a delivery encountered issues
	// while trying to deliver
	DeliveryFailed = "delivery_failed"
	// DeliveryCanceled is the status returned when a delivery is purposefully
	// halted from going through
	DeliveryCanceled = "delivery_canceled"
)

// Delivery represents the delivery information of a singular artifact
// associated with an analysis status
type DeliveryStatus struct {
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
