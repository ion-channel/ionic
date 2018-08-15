package ionic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/dependencies"
)

const (
	getLatestVersionForDependencyEndpoint = "v1/dependency/getLatestVersionForDependency"

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

	return &dep, nil
}
