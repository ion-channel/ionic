package teamusers

import (
	"time"
)

const (
	// TeamsCreateTeamUserEndpoint is a string representation of the current endpoint for creating team user
	TeamsCreateTeamUserEndpoint = "v1/teamUsers/createTeamUser"
	// TeamsGetTeamUserEndpoint is a string representation of the current endpoint for getting team user
	TeamsGetTeamUserEndpoint = "v1/teamUsers/getTeamUser"
	// TeamsUpdateTeamUserEndpoint is a string representation of the current endpoint for updating team user
	TeamsUpdateTeamUserEndpoint = "v1/teamUsers/updateTeamUser"
	// TeamsDeleteTeamUserEndpoint is a string representation of the current endpoint for deleting team user
	TeamsDeleteTeamUserEndpoint = "v1/teamUsers/deleteTeamUser"

	// TODO: The following endpoints will need to have functions attached to them

	// TeamsInviteTeamUserEndpoint is a string representation of the current endpoint for Inviting a team user
	TeamsInviteTeamUserEndpoint = "v1/teamUsers/inviteTeamUser"
	// TeamsAcceptInviteTeamUserEndpoint is a string representation of the current endpoint for accepting a team user invite
	TeamsAcceptInviteTeamUserEndpoint = "v1/teamUsers/acceptInvite"
	// TeamsGetInviteTeamUserEndpoint is a string representation of the current endpoint for inviting a team user
	TeamsGetInviteTeamUserEndpoint = "v1/teamUsers/getInvite"
	// TeamsResendInviteTeamUserEndpoint is a string representation of the current endpoint for resending invite to team user
	TeamsResendInviteTeamUserEndpoint = "v1/teamUsers/resendInvite"
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
