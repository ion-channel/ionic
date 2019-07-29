package teams

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// TeamsCreateTeamEndpoint is a string representation of the current endpoint for creating team
	TeamsCreateTeamEndpoint = "v1/teams/createTeam"
	// TeamsGetTeamEndpoint is a string representation of the current endpoint for getting team
	TeamsGetTeamEndpoint = "v1/teams/getTeam"
	// TeamsGetTeamsEndpoint is a string representation of the current endpoint for getting teams
	TeamsGetTeamsEndpoint = "v1/teams/getTeams"

	// TODO: the following endpoints need functions attached to them

	// TeamsEstablishTeamEndpoint is a string representation of the current endpoint for establishing a team
	TeamsEstablishTeamEndpoint = "v1/teams/establishTeam"
	// TeamsUpdateTeamEndpoint is a string representation of the current endpoint for updating a team
	TeamsUpdateTeamEndpoint = "v1/teams/updateTeam"
	// TeamsDeleteTeamEndpoint is a string representation of the current endpoint for deleting a team
	TeamsDeleteTeamEndpoint = "v1/teams/deleteTeam"
)

// Team is a representation of an Ion Channel Team within the system
type Team struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	Name       string    `json:"name"`
	Delivering bool      `json:"delivering"`
	SysAdmin   bool      `json:"sys_admin"`
	POCName    string    `json:"poc_name"`
	POCEmail   string    `json:"poc_email"`
}

// String returns a JSON formatted string of the team object
func (t Team) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("failed to format team: %v", err.Error())
	}
	return string(b)
}
