package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/teams"
)

const (
	teamsGetTeamEndpoint = "v1/teams/getTeam"
)

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
