package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/aliases"
)

// AddAliasOptions struct that allows for adding an alias to a
// project
type AddAliasOptions struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	TeamID    string `json:"team_id"`
	Version   string `json:"version"`
	Org       string `json:"org"`
}

//AddAlias takes a project and adds an alias to it. It returns the
// project stored or an error encountered by the API
func (ic *IonClient) AddAlias(alias AddAliasOptions, token string) (*aliases.Alias, error) {
	params := &url.Values{}
	params.Set("project_id", alias.ProjectID)

	b, err := json.Marshal(alias)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall alias: %v", err.Error())
	}

	b, err = ic.Post(aliases.AddAliasEndpoint, token, params, *bytes.NewBuffer(b), nil)
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
