package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/aliases"
)

const (
	addAliasEndpoint = "v1/project/addAlias"
)

type createAliasOptions struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	TeamID    string `json:"team_id"`
	Version   string `json:"version"`
}

//AddAlias takes a project and adds an alias to it. It returns the
// project stored or an error encountered by the API
func (ic *IonClient) AddAlias(projectID, teamID, name, version, token string) (*aliases.Alias, error) {
	params := &url.Values{}

	alias := createAliasOptions{
		Name:      name,
		Version:   version,
		TeamID:    teamID,
		ProjectID: projectID,
	}

	b, err := json.Marshal(alias)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall alias: %v", err.Error())
	}

	b, err = ic.Post(addAliasEndpoint, token, params, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create alias: %v", err.Error())
	}

	var a aliases.Alias
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from create: %v", err.Error())
	}

	return &a, nil
}
