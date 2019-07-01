package community

// Repo is a representation of a github repo and corresponding metrics about
// that repo pulled from github
type Repo struct {
	Name       string   `json:"name" xml:"name"`
	URL        string   `json:"url" xml:"url"`
	Committers int      `json:"committers" xml:"committers"`
	Confidence float64  `json:"confidence" xml:"confidence"`
	OldNames   []string `json:"old_names" xml:"old_names"`
	Stars      int      `json:"stars" xml:"stars"`
}

const (
	GetRepoEndpoint    = `v1/repo/getRepo`
	SearchRepoEndpoint = `v1/repo/search`
)
