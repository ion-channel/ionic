package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/dependencies"
)

const (
	getLatestVersionForDependencyEndpoint  = "v1/dependency/getLatestVersionForDependency"
	getLatestVersionsForDependencyEndpoint = "v1/dependency/getLatestVersionsForDependency"

	// RubyEcosystem represents the ruby ecosystem for resolving dependencies
	RubyEcosystem = "ruby"
)

// GetLatestVersionForDependency takes a package name, an ecosystem to find the
// package in, and a token for accessing the API. It returns a dependency
// representation of the latest version and any errors it encounters with the
// API.
func (ic *IonClient) GetLatestVersionForDependency(packageName, ecosystem, token string) (*dependencies.Dependency, error) {
	params := &url.Values{}
	params.Set("name", packageName)
	params.Set("type", ecosystem)

	b, err := ic.Get(getLatestVersionForDependencyEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest version for dependency: %v", err.Error())
	}

	var dep dependencies.Dependency
	err = json.Unmarshal(b, &dep)
	if err != nil {
		return nil, fmt.Errorf("cannot parse dependency: %v", err.Error())
	}

	dep.Name = packageName
	return &dep, nil
}

// GetLatestVersionsForDependency takes a package name, an ecosystem to find the
// package in, and a token for accessing the API. It returns a dependency
// representation of the latest versions and any errors it encounters with the
// API.
func (ic *IonClient) GetLatestVersionsForDependency(packageName, ecosystem, token string) ([]dependencies.Dependency, error) {
	params := &url.Values{}
	params.Set("name", packageName)
	params.Set("type", ecosystem)

	b, err := ic.Get(getLatestVersionsForDependencyEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest version for dependency: %v", err.Error())
	}

	var vs []string
	err = json.Unmarshal(b, &vs)
	if err != nil {
		return nil, fmt.Errorf("cannot parse dependency: %v", err.Error())
	}

	deps := []dependencies.Dependency{}
	for index := range vs {
		dep := dependencies.Dependency{
			Name:    packageName,
			Version: vs[index],
		}
		deps = append(deps, dep)
	}

	return deps, nil
}
