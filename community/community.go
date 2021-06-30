package community

import "time"

const (
	// GetRepoEndpoint is a string representation of the current endpoint for getting repo
	GetRepoEndpoint = `v1/repo/getRepo`
	// GetReposInCommonEndpoint is a string representation of the current endpoint for getting repos
	GetReposInCommonEndpoint = `/v1/repo/getReposInCommon`
	// GetReposForActorEndpoint is a string representation of the current endpoint for getting repos
	GetReposForActorEndpoint = `v1/repo/getReposForActor`
	// SearchRepoEndpoint is a string representation of the current endpoint for searching repo
	SearchRepoEndpoint = `v1/repo/search`
)

// Repo is a representation of a github repo and corresponding metrics about
// that repo pulled from github
type Repo struct {
	ID            string    `json:"id" xml:"id"`
	Name          string    `json:"name" xml:"name"`
	URL           string    `json:"url" xml:"url"`
	Committers    int       `json:"committers" xml:"committers"`
	TotalActors   int       `json:"total_actors,omitempty" xml:"total_actors,omitempty"`
	Confidence    float64   `json:"confidence" xml:"confidence"`
	OldNames      []string  `json:"old_names" xml:"old_names"`
	DefaultBranch string    `json:"default_branch,omitempty" xml:"default_branch,omitempty"`
	MasterBranch  string    `json:"master_branch,omitempty" xml:"master_branch,omitempty"`
	Stars         int       `json:"stars" xml:"stars"`
	CommittedAt   time.Time `json:"committed_at" xml:"committed_at"`
	UpdatedAt     time.Time `json:"updated_at" xml:"updated_at"`
}
