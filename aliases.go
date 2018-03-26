package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/projects"
)

const (
	addAliasEndpoint = "v1/project/addAlias"
)

//AddAlias takes a project and adds an alias to it. It returns the
// project stored or an error encountered by the API
func (ic *IonClient) AddAlias(projectID, teamID, name, version, token string) (*projects.Project, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("project_id", projectID)

	alias := aliases.Alias{
		Name:    name,
		Version: version,
	}

	b, err := json.Marshal(alias)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall alias: %v", err.Error())
	}

	b, err = ic.Post(addAliasEndpoint, token, params, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create alias: %v", err.Error())
	}

	var p projects.Project
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from create: %v", err.Error())
	}

	return &p, nil
}
