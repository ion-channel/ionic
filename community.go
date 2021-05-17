package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/community"
)

// GetRepo takes in a repository string and calls the Ion API to get
// a pointer to the Ionic community.Repo
func (ic *IonClient) GetRepo(repo, token string) (*community.Repo, error) {
	params := &url.Values{}
	params.Set("repo", repo)

	b, _, err := ic.Get(community.GetRepoEndpoint, token, params, nil, nil)
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

// GetReposInCommonOptions encapsulates params for repos in common requests
type GetReposInCommonOptions struct {
	Subject    string   `json:"subject"`
	Comparands []string `json:"comparands"`
	ByActor    bool     `json:"by_actor"`
}

// GetReposInCommonOutput encapsulates params for repos in common requests
type GetReposInCommonOutput struct {
	community.Repo
	CommonCommitters []string `json:"common_committers,omitempty" xml:"common_committers,omitempty"`
	CommonActors     []string `json:"common_actors,omitempty" xml:"common_actors,omitempty"`
}

// GetReposInCommon takes in an subject repo, a slice of string camparands and bool option for actors
//  and calls the Ion API to get matches with the count of committers shared
func (ic *IonClient) GetReposInCommon(options GetReposInCommonOptions, token string) ([]GetReposInCommonOutput, error) {
	body, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal options for repos in common (%s) : %v", options.Subject, err.Error())
	}

	b, err := ic.Post(community.GetReposInCommonEndpoint, token, nil, *bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get repos in common (%s) : %v", options.Subject, err.Error())
	}
	var resultRepos []GetReposInCommonOutput
	err = json.Unmarshal(b, &resultRepos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal repos in common results: %v (%v)", err.Error(), string(b))
	}
	return resultRepos, nil
}

// GetReposForActor takes in an user, committer or actor string and calls the Ion API to get
// a slice of Ionic community.Repo
func (ic *IonClient) GetReposForActor(name, token string) ([]community.Repo, error) {
	params := &url.Values{}
	params.Set("name", name)

	b, _, err := ic.Get(community.GetReposForActorEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get repos for actor (%s) : %v", name, err.Error())
	}
	var resultRepos []community.Repo
	err = json.Unmarshal(b, &resultRepos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal getRepos results: %v (%v)", err.Error(), string(b))
	}
	return resultRepos, nil
}

// SearchRepo takes a query `org AND name` and
// calls the Ion API to retrieve the information, then forms a slice of
// Ionic community.Repo objects
func (ic *IonClient) SearchRepo(q, token string) ([]community.Repo, error) {
	params := &url.Values{}
	params.Set("q", q)

	b, _, err := ic.Get(community.SearchRepoEndpoint, token, params, nil, nil)
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
