package community

// Repo is a representation of a github repo and corresponding metrics about
// that repo pulled from github
type Repo struct {
	Name       string `json:"name" xml:"name"`
	URL        string `json:"url" xml:"url"`
	Committers int    `json:"committers" xml:"committers"`
}
