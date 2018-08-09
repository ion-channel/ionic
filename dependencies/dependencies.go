package dependencies

type Dependency struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version"`
}
