package community

import "time"

const (
	// GetRepoEndpoint is a string representation of the current endpoint for getting repo
	GetRepoEndpoint = `v1/repo/getRepo`
	// SearchRepoEndpoint is a string representation of the current endpoint for searching repo
	SearchRepoEndpoint = `v1/repo/search`
)

// Repo is a representation of a github repo and corresponding metrics about
// that repo pulled from github
type Repo struct {
	Name        string    `json:"name" xml:"name"`
	URL         string    `json:"url" xml:"url"`
	Committers  int       `json:"committers" xml:"committers"`
	Confidence  float64   `json:"confidence" xml:"confidence"`
	OldNames    []string  `json:"old_names" xml:"old_names"`
	Stars       int       `json:"stars" xml:"stars"`
	CommittedAt time.Time `json:"committed_at" xml:"committed_at"`
	UpdatedAt   time.Time `json:"updated_at" xml:"updated_at"`
}
