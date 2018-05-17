package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/tags"
)

const (
	createTagEndpoint = "v1/tag/createTag"
	getTagEndpoint    = "v1/tag/getTag"
)

// CreateTag takes a team ID, name, and description. It returns the details of
// the created tag, or any errors encountered with the API.
func (ic *IonClient) CreateTag(teamID, name, description, token string) (*tags.Tag, error) {
	tag := &tags.Tag{
		TeamID:      teamID,
		Name:        name,
		Description: description,
	}

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tag params to JSON: %v", err.Error())
	}

	b, err = ic.Post(createTagEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create tag: %v", err.Error())
	}

	var t tags.Tag
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from create: %v", err.Error())
	}

	return &t, nil
}

// GetTag takes a tag ID and a team ID. It returns the details of a singular
// tag and any errors encountered with the API.
func (ic *IonClient) GetTag(id, teamID, token string) (*tags.Tag, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(getTagEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %v", err.Error())
	}

	var t tags.Tag
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, fmt.Errorf("cannot parse tag: %v", err.Error())
	}

	return &t, nil
}

// GetRawTag takes a tag ID and a team ID. It returns the details of a singular
// tag and any errors encountered with the API.
func (ic *IonClient) GetRawTag(id, teamID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(getTagEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %v", err.Error())
	}

	return b, nil
}
