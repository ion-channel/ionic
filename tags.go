package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/tags"
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

	b, err = ic.Post(tags.CreateTagEndpoint, token, nil, *bytes.NewBuffer(b), nil)
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

// UpdateTag takes an ID, team ID, name, and description. It returns the details of
// the updated tag, or any errors encountered with the API.
func (ic *IonClient) UpdateTag(id, teamID, name, description, token string) (*tags.Tag, error) {
	tag := &tags.Tag{
		ID:          id,
		TeamID:      teamID,
		Name:        name,
		Description: description,
	}

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tag params to JSON: %v", err.Error())
	}

	b, err = ic.Put(tags.UpdateTagEndpoint, token, nil, *bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update tag: %v", err.Error())
	}

	var t tags.Tag
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from update: %v", err.Error())
	}

	return &t, nil
}

// GetTag takes a tag ID and a team ID. It returns the details of a singular
// tag and any errors encountered with the API.
func (ic *IonClient) GetTag(id, teamID, token string) (*tags.Tag, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(tags.GetTagEndpoint, token, params, nil, nil)
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

// GetTags takes a team ID. It returns the details of a singular tag and any
// errors encountered with the API.
func (ic *IonClient) GetTags(teamID, token string) ([]tags.Tag, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := ic.Get(tags.GetTagsEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %v", err.Error())
	}

	var ts []tags.Tag
	err = json.Unmarshal(b, &ts)
	if err != nil {
		return nil, fmt.Errorf("cannot parse tag: %v", err.Error())
	}

	return ts, nil
}

// GetRawTags takes a team ID. It returns the details of a singular tag and any
// errors encountered with the API.
func (ic *IonClient) GetRawTags(teamID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)

	b, err := ic.Get(tags.GetTagsEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %v", err.Error())
	}

	return b, nil
}

// GetRawTag takes a tag ID and a team ID. It returns the details of a singular
// tag and any errors encountered with the API.
func (ic *IonClient) GetRawTag(id, teamID, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("id", id)
	params.Set("team_id", teamID)

	b, err := ic.Get(tags.GetTagEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %v", err.Error())
	}

	return b, nil
}
