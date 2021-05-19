package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/ion-channel/ionic/scans"

	"github.com/ion-channel/ionic/dependencies"
)

const (
	// RubyEcosystem represents the ruby ecosystem for resolving dependencies
	RubyEcosystem = "ruby"
)

// ResolveDependenciesInFile takes a dependency file location and token to send
// the specified file to the API. All dependencies that are able to be resolved will
// be with their info returned, and a list of any errors encountered during the
// process.
func (ic *IonClient) ResolveDependenciesInFile(o dependencies.DependencyResolutionRequest, token string) (*dependencies.DependencyResolutionResponse, error) {
	params := &url.Values{}
	params.Set("type", o.Ecosystem)
	if o.Flatten {
		params.Set("flatten", "true")
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	fw, err := w.CreateFormFile("file", o.File)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err.Error())
	}

	fh, err := os.Open(o.File)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err.Error())
	}

	_, err = io.Copy(fw, fh)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file contents: %v", err.Error())
	}

	w.Close()

	var endpoint string
	switch {
	case strings.Contains(path.Base(o.File), "Gemfile.lock") || strings.Contains(path.Base(o.File), "go.mod"):
		endpoint = dependencies.ResolveFromFileEndpoint
	default:
		endpoint = dependencies.ResolveDependenciesInFileEndpoint
	}

	h := http.Header{}
	h.Set("Content-Type", w.FormDataContentType())

	b, err := ic.Post(endpoint, token, params, buf, h)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve dependencies: %v", err.Error())
	}

	var resp dependencies.DependencyResolutionResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err.Error())
	}

	return &resp, nil
}

// GetLatestVersionForDependency takes a package name, an ecosystem to find the
// package in, and a token for accessing the API. It returns a dependency
// representation of the latest version and any errors it encounters with the
// API.
func (ic *IonClient) GetLatestVersionForDependency(packageName, ecosystem, token string) (*dependencies.Dependency, error) {
	params := &url.Values{}
	params.Set("name", packageName)
	params.Set("type", ecosystem)

	b, _, err := ic.Get(dependencies.GetLatestVersionForDependencyEndpoint, token, params, nil, nil)
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

// GetVersionsForDependency takes a package name, an ecosystem to find the
// package in, and a token for accessing the API. It returns a dependency
// representation of the latest versions and any errors it encounters with the
// API.
func (ic *IonClient) GetVersionsForDependency(packageName, ecosystem, token string) ([]dependencies.Dependency, error) {
	params := &url.Values{}
	params.Set("name", packageName)
	params.Set("type", ecosystem)

	b, _, err := ic.Get(dependencies.GetVersionsForDependencyEndpoint, token, params, nil, nil)
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

// SearchDependencies takes a query `org AND name` and
// calls the Ion API to retrieve the information, then forms a slice of
// Ionic dependencies.Dependency objects
func (ic *IonClient) SearchDependencies(q, token string) ([]dependencies.Dependency, error) {
	params := &url.Values{}
	params.Set("q", q)

	b, _, err := ic.Get(dependencies.ResolveDependencySearchEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get dependencies: %v", err.Error())
	}
	var results []dependencies.Dependency
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal search results: %v (%v)", err.Error(), string(b))
	}
	return results, nil
}

// GetDependencyVersions takes a package name, an ecosystem to find the
// package in, optional version, and a token for accessing the API.
// If version is supplied, it will return all known versions greater than what was given
// It returns a slice of Ionic dependencies.Dependency objects
func (ic *IonClient) GetDependencyVersions(packageName, ecosystem, version, token string) ([]dependencies.Dependency, error) {
	params := &url.Values{}
	params.Set("name", packageName)
	params.Set("type", ecosystem)
	if version != "" {
		params.Set("version", version)
	}

	b, _, err := ic.Get(dependencies.GetDependencyVersions, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get dependency versions: %v", err.Error())
	}

	var deps []dependencies.Dependency
	err = json.Unmarshal(b, &deps)
	if err != nil {
		return nil, fmt.Errorf("cannot parse dependency: %v", err.Error())
	}

	return deps, nil
}

// GetDifferenceBetweenVersions calculates the difference between two version strings, returning an OutdatedMeta
// object, or an error.
func GetDifferenceBetweenVersions(newerVersion, olderVersion string) (outdatedMeta scans.OutdatedMeta, err error) {
	ver, err := version.NewVersion(olderVersion)
	if err != nil {
		return outdatedMeta, err
	}

	latestVersion, err := version.NewVersion(newerVersion)
	if err != nil {
		return outdatedMeta, err
	}

	// for major.minor.patch
	var versionsBehind [3]int
	// check if our version is out of date, and calculate its versions behind if so
	if ver.LessThan(latestVersion) {
		versions := ver.Segments()
		latestVer := latestVersion.Segments()
		for i := 0; i < len(versionsBehind); i++ {
			if versions[i] <= latestVer[i] {
				versionsBehind[i] = latestVer[i] - versions[i]
			} else {
				versionsBehind[i] = 0
			}
		}

		outdatedMeta = scans.OutdatedMeta{
			MajorBehind: versionsBehind[0],
			MinorBehind: versionsBehind[1],
			PatchBehind: versionsBehind[2],
		}
	}

	return outdatedMeta, err
}
