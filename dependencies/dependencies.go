package dependencies

const (
	// GetLatestVersionForDependencyEndpoint is a string representation of the current endpoint for getting latest versions for dependency
	GetLatestVersionForDependencyEndpoint = "v1/dependency/getLatestVersionForDependency"
	// GetVersionsForDependencyEndpoint is a string representation of the current endpoint for getting versions for dependency
	GetVersionsForDependencyEndpoint = "v1/dependency/getVersionsForDependency"
)

// Dependency represents all the known information for a dependency object
// within the Ion Channel API
type Dependency struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version"`
}
