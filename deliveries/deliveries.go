package deliveries

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// DeliveriesGetDestinationsEndpoint returns all destinations for a team. Requires team id.
	DeliveriesGetDestinationsEndpoint = "/v1/teams/getDeliveryDestinations"
	// DeliveriesDeleteDestinationEndpoint markes a delivery destination as deleted. It requires a delivery destination id.
	DeliveriesDeleteDestinationEndpoint = "/v1/teams/deleteDeliveryDestination"
	// CreateDeleteDestinationEndpoint creates a destination. Requires team id, location, region, name, destination type, access key (empty string allowed), secret key (empty string allowed) and token.
	CreateDeleteDestinationEndpoint = "/v1/teams/createDeliveryDestination"
)

// Destination is a representation of a single location that a team can deliver results to.
type Destination struct {
	ID        string     `json:"id"`
	TeamID    string     `json:"team_id"`
	Location  string     `json:"location"`
	Region    string     `json:"region"`
	Name      string     `json:"name"`
	DestType  string     `json:"type"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// CreateDestination is an input representation of a single location that a team can deliver results to.
type CreateDestination struct {
	Destination
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// String returns a JSON formatted string of the delivery object
func (p Destination) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("failed to format delivery: %v", err.Error())
	}
	return string(b)
}
