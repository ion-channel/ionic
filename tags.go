package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/tags"
)

const (
	getTagEndpoint = "v1/tag/getTag"
)

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
