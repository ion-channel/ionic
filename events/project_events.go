package events

// ProjectEvent represents a project within an Event within Ion Channel
type ProjectEvent struct {
	Project  string   `json:"project"`
	Org      string   `json:"org"`
	Versions []string `json:"versions"`
}
