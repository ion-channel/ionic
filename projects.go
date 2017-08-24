package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/projects"
)

const (
	getProjectEndpoint = "v1/project/getProject"
)

// GetProject takes a project ID and team ID and returns the project.  It
// returns an error if it receives a bad response from the API or fails to
// unmarshal the JSON response from the API.
func (ic *IonClient) GetProject(id, teamID string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.get(getProjectEndpoint, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %v", err.Error())
	}

	return &p, nil
}
