package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/teams"
)

const (
	teamsCreateTeamEndpoint = "v1/teams/createTeam"
	teamsGetTeamEndpoint    = "v1/teams/getTeam"
)

// CreateTeamOptions represents all the values that can be provided for a team
// at the time of creation
type CreateTeamOptions struct {
	Name             string `json:"name"`
	POCName          string `json:"poc_name"`
	POCEmail         string `json:"poc_email"`
	DeliveryLocation string `json:"delivery_location"`
	AccessKey        string `json:"access_key"`
	SecretKey        string `json:"secret_key"`
	DeliveryRegion   string `json:"delivery_region"`
}

// CreateTeam takes a create team options, validates the minimum info is
// present, and makes the calls to create the team. It returns the team created
// and any errors it encounters with the API.
func (ic *IonClient) CreateTeam(opts CreateTeamOptions) (*teams.Team, error) {
	if opts.Name == "" {
		return nil, fmt.Errorf("name missing from options")
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(teamsCreateTeamEndpoint, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %v", err.Error())
	}

	var t teams.Team
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to parse team from response: %v", err.Error())
	}

	return &t, nil
}

// GetTeam takes a team id and returns the Ion Channel representation of that
// team.  An error is returned for client communications and unmarshalling
// errors.
func (ic *IonClient) GetTeam(id string) (*teams.Team, error) {
	params := &url.Values{}
	params.Set("someid", id)

	b, err := ic.Get(teamsGetTeamEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get team: %v", err.Error())
	}

	var team teams.Team
	err = json.Unmarshal(b, &team)
	if err != nil {
		return nil, fmt.Errorf("cannot parse team: %v", err.Error())
	}

	return &team, nil
}
