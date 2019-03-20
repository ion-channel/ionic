package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/community"
)

const (
	getRepoEndpoint    = `v1/repo/getRepo`
	searchRepoEndpoint = `v1/repo/search`
)

// GetRepo takes in a repository string and calls the Ion API to get
// a pointer to the Ionic community.Repo
func (ic *IonClient) GetRepo(repo, token string) (*community.Repo, error) {
	params := &url.Values{}
	params.Set("repo", repo)

	b, err := ic.Get(getRepoEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get repo: %v", err.Error())
	}
	var resultRepo community.Repo
	err = json.Unmarshal(b, &resultRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal getRepo results: %v (%v)", err.Error(), string(b))
	}
	return &resultRepo, nil
}

// SearchRepo takes an organization name (org) and a project name (project) and
// calls the Ion API to retrieve the information, then forms a slice of
// Ionic community.Repo objects
func (ic *IonClient) SearchRepo(q, token string) ([]community.Repo, error) {
	params := &url.Values{}
	params.Set("q", q)

	b, err := ic.Get(searchRepoEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get repo: %v", err.Error())
	}
	var results []community.Repo
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal getRepo results: %v (%v)", err.Error(), string(b))
	}
	return results, nil
}
