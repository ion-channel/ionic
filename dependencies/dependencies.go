package dependencies

// Dependency represents all the known information for a dependency object
// within the Ion Channel API
type Dependency struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version"`
}

const (
	GetLatestVersionForDependencyEndpoint = "v1/dependency/getLatestVersionForDependency"
	GetVersionsForDependencyEndpoint      = "v1/dependency/getVersionsForDependency"
)
