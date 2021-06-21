package organizations

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// OrganizationsCreateOrganizationEndpoint is the endpoint for creating an organization
	OrganizationsCreateOrganizationEndpoint = "v1/organizations/createOrganizations"
	// OrganizationsGetOrganizationEndpoint is the endpoint for getting an organization
	OrganizationsGetOrganizationEndpoint = "v1/teams/getTeam"
	// OrganizationsGetOrganizationsEndpoint is the endpoint for getting organizations
	OrganizationsGetOrganizationsEndpoint = "v1/teams/getTeams"
)

// Organization is a logical collection of teams.
type Organization struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
	Name             string    `json:"name"`
}

// String returns a JSON formatted string of the team object
func (o Organization) String() string {
	b, err := json.Marshal(o)
	if err != nil {
		return fmt.Sprintf("failed to format organization: %v", err.Error())
	}

	return string(b)
}
