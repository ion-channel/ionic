package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/teamusers"
)

const (
	teamsCreateTeamUserEndpoint = "v1/teamUsers/createTeamUser"
	teamsGetTeamUserEndpoint    = "v1/teamUsers/getTeamUser"
	teamsUpdateTeamUserEndpoint = "v1/teamUsers/updateTeamUser"
)

// CreateTeamUserOptions represents all the values that can be provided for a team
// user at the time of creation
type CreateTeamUserOptions struct {
	Status    string `json:"status"`
	Role      string `json:"role"`
	TeamID    string `json:"team_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UpdateTeamUserOptions represents all the values that can be provided to update a team user
type UpdateTeamUserOptions struct {
	Status    string `json:"status"`
	Role      string `json:"role"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateTeamUser takes a create team options, validates the minimum info is
// present, and makes the calls to create the team. It returns the team created
// and any errors it encounters with the API.
func (ic *IonClient) CreateTeamUser(opts CreateTeamUserOptions, token string) (*teamusers.TeamUser, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(teamsCreateTeamUserEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create team user: %v", err.Error())
	}

	var tu teamusers.TeamUser
	err = json.Unmarshal(b, &tu)
	if err != nil {
		return nil, fmt.Errorf("failed to parse team user from response: %v", err.Error())
	}

	return &tu, nil
}

// GetTeamUser takes a team id and returns the Ion Channel representation of that
// team.  An error is returned for client communications and unmarshalling
// errors.
func (ic *IonClient) GetTeamUser(teamID, userID, token string) (*teamusers.TeamUser, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("user_id", userID)

	b, err := ic.Get(teamsGetTeamUserEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get team: %v", err.Error())
	}

	var teamU teamusers.TeamUser
	err = json.Unmarshal(b, &teamU)
	if err != nil {
		return nil, fmt.Errorf("cannot parse team: %v", err.Error())
	}

	return &teamU, nil
}

// UpdateTeamUser takes a teamusers options, validates the minimum info is
// present, and makes the calls to update the teamuser. It returns the team created
// and any errors it encounters with the API.
func (ic *IonClient) UpdateTeamUser(opts UpdateTeamUserOptions, token string) (*teamusers.TeamUser, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(teamsUpdateTeamUserEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update team user: %v", err.Error())
	}

	var tu teamusers.TeamUser
	err = json.Unmarshal(b, &tu)
	if err != nil {
		return nil, fmt.Errorf("failed to parse team user from response: %v", err.Error())
	}

	return &tu, nil
}
