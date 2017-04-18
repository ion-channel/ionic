package ionic

type Event struct {
	Vulns EventVulnerability `json:"vulns"`
}

type EventVulnerability struct {
	Updates  []string       `json:"updates"`
	Projects []EventProject `json:"projects"`
}

type EventProject struct {
	Project  string   `json:"project"`
	Org      string   `json:"org"`
	Versions []string `json:"versions"`
}
