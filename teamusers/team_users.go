package teamusers

import (
	"time"
)

const (
	// TeamsCreateTeamUserEndpoint is a string representation of the current endpoint for getting projects
	TeamsCreateTeamUserEndpoint = "v1/teamUsers/createTeamUser"
	// TeamsGetTeamUserEndpoint is a string representation of the current endpoint for getting projects
	TeamsGetTeamUserEndpoint = "v1/teamUsers/getTeamUser"
	// TeamsUpdateTeamUserEndpoint is a string representation of the current endpoint for getting projects
	TeamsUpdateTeamUserEndpoint = "v1/teamUsers/updateTeamUser"
	// TeamsDeleteTeamUserEndpoint is a string representation of the current endpoint for getting projects
	TeamsDeleteTeamUserEndpoint = "v1/teamUsers/deleteTeamUser"
)

// TeamUser is a representation of an Ion Channel Team User relationship within the system
type TeamUser struct {
	ID        string    `json:"id"`
	TeamID    string    `json:"team_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
}
