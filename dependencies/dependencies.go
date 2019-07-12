package dependencies

// Dependency represents all the known information for a dependency object
// within the Ion Channel API
type Dependency struct {
	Name          string       `json:"name,omitempty"`
	Version       string       `json:"version"`
	LatestVersion string       `json:"latest_version"`
	Org           string       `json:"org"`
	Type          string       `json:"type"`
	Package       string       `json:"package"`
	Scope         string       `json:"scope"`
	Requirement   string       `json:"requirement"`
	Dependencies  []Dependency `json:"dependencies"`
}

// Meta represents all the known meta information for a dependency set
// within the Ion Channel API
type Meta struct {
	// {"first_degree_count":13,"no_version_count":0,"total_unique_count":62,"update_available_count":12}
	FirstDegreeCount     int `json:"first_degree_count"`
	NoVersionCount       int `json:"no_version_count"`
	TotalUniqueCount     int `json:"total_unique_count"`
	UpdateAvailableCount int `json:"update_available_count"`
}

// DependencyResolutionResponse represents all the known information
// for a dependency object within the Ion Channel API
type DependencyResolutionResponse struct {
	Dependencies []Dependency `json:"dependencies,omitempty"`
	Meta         Meta         `json:"meta"`
}

// DependencyResolutionRequest options for creating a resolution request
// for a dependency file of a ecosystem type
type DependencyResolutionRequest struct {
	Ecosystem string
	File      string
	Flatten   bool
}
