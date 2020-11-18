package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

type dfilter func(*scans.Dependency) *scans.Dependency

var eVs = [...]string{"", ">= 0", "\u003e= 0", ">=0.0", "\u003e=0.0"}

func nv(d *scans.Dependency) *scans.Dependency {
	for _, e := range eVs {
		if d.Requirement == e {
			return d
		}
	}
	if d.Dependencies != nil {
		for _, dep := range d.Dependencies {
			f := nv(&dep)
			if f != nil {
				return d
			}
		}
	}
	return nil
}

func od(d *scans.Dependency) *scans.Dependency {
	// this should be changed to also handle things like `>= 0`
	if d.Version < d.LatestVersion && d.Version != "" {
		return d
	}
	if d.Dependencies != nil {
		for _, dep := range d.Dependencies {
			f := od(&dep)
			if f != nil {
				return d
			}
		}
	}
	return nil
}

// if a dep has outdated deps, count them
func outdatedWithMeta(d *scans.Dependency) *scans.Dependency {
	scanDeps := *d

	filtered := filterDependenciesOnList(d.Dependencies, false, nv)
	noVersCount := len(filtered)

	// for each direct dep, figure out how many of its deps are outdated
	filtered = filterDependenciesOnList(d.Dependencies, false, od)
	// we're only interested in the outdated deps
	scanDeps.Dependencies = filtered

	outdatedCount := len(filtered)
	m := scans.DependencyMeta{
		NoVersionCount:       noVersCount,
		UpdateAvailableCount: outdatedCount,
		VulnerableCount:      0, // how do we calculate this? We need vulnerability results in here to tie them together
	}
	scanDeps.DepMeta = &m

	return &scanDeps
}

func giveem(d *scans.Dependency) *scans.Dependency {
	return d
}

func directWithMeta(d *scans.Dependency) *scans.Dependency {
	scanDeps := *d
	topLevelDependencies := make([]scans.Dependency, 0)
	noVersCount := 0
	outdatedCount := 0
	for x := range scanDeps.Dependencies {
		scanDeps.Dependencies[x].DepMeta = nil
		scanDeps.Dependencies[x].Dependencies = nil
		topLevelDependencies = append(topLevelDependencies, scanDeps.Dependencies[x])
	}

	// We want to remove all transitive deps but keep them around for calculating values with
	scanDeps.Dependencies = nil

	filtered := filterDependenciesOnList(topLevelDependencies, false, nv)
	noVersCount = len(filtered)

	filtered = filterDependenciesOnList(topLevelDependencies, false, od)
	outdatedCount = len(filtered)

	m := scans.DependencyMeta{
		NoVersionCount:       noVersCount,
		UpdateAvailableCount: outdatedCount,
		VulnerableCount:      0, // how do we calculate this? We need vulnerability results in here to tie them together
	}
	scanDeps.DepMeta = &m

	return &scanDeps
}

func direct(d *scans.Dependency) *scans.Dependency {
	ff := d
	ff.Dependencies = nil
	return ff
}

func filterDependenciesOnList(deps []scans.Dependency, unique bool, f dfilter) []scans.Dependency {
	ds := []scans.Dependency{}
	for _, dr := range deps {

		filtered := f(&dr)
		if filtered != nil {
			ds = append(ds, *filtered)
		}
	}
	return ds
}

func filterDependencies(data interface{}, unique bool, f dfilter) ([]scans.Dependency, error) {
	b, ok := data.(scans.DependencyResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into dep")
	}

	return filterDependenciesOnList(b.Dependencies, unique, f), nil
}

func dependencyDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)
	var data interface{}

	var updateAvailable, noVersions, directDeps, transDeps int
	if eval != nil && !status.Errored() {
		data = eval.TranslatedResults.Data
		b, ok := data.(scans.DependencyResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
		}
		updateAvailable = b.Meta.UpdateAvailableCount
		noVersions = b.Meta.NoVersionCount
		directDeps = b.Meta.FirstDegreeCount
		transDeps = b.Meta.TotalUniqueCount - b.Meta.FirstDegreeCount
	}

	d := NewDigest(status, DependencyOutdatedIndex, "dependency outdated", "dependencies outdated")

	if eval != nil && !status.Errored() {
		// return all outdated deps as well as direct deps and the count of their outdated dependencies
		filtered, err := filterDependencies(data, false, outdatedWithMeta)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to no version dependency digest: %v", err.Error())
		}

		err = d.AppendEval(eval, "count", updateAvailable)
		if err != nil {
			return nil, fmt.Errorf("failed to create dependencies outdated digest: %v", err.Error())
		}

		d.MarshalSourceData(filtered, "dependency")
		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	// No version specified
	d = NewDigest(status, NoVersionIndex, "dependency no version specified", "dependencies no version specified")

	if eval != nil && !status.Errored() {
		filtered, err := filterDependencies(data, false, nv)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to no version dependency digest: %v", err.Error())
		}
		d.MarshalSourceData(filtered, "dependency")
		err = d.AppendEval(eval, "count", noVersions)
		if err != nil {
			return nil, fmt.Errorf("failed to create dependencies no version digest: %v", err.Error())
		}

		if noVersions > 0 {
			d.Warning = true
			d.WarningMessage = "dependencies with no version specified"

			if noVersions == 1 {
				d.WarningMessage = "dependency with no version specified"
			}
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, DirectDependencyIndex, "direct dependency", "direct dependencies")

	if eval != nil && !status.Errored() {
		filtered, err := filterDependencies(data, false, direct)

		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to direct dependency digest: %v", err.Error())
		}
		d.MarshalSourceData(filtered, "dependency")
		err = d.AppendEval(eval, "count", directDeps)
		if err != nil {
			return nil, fmt.Errorf("failed to create direct dependencies digeest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

		if directDeps < 1 {
			d.Warning = true
			d.WarningMessage = "no direct dependencies found"
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, TransitiveDependencyIndex, "transitive dependency", "transitive dependencies")

	if eval != nil && !status.Errored() {
		filtered, err := filterDependencies(data, false, giveem)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to transitive dependency digest: %v", err.Error())
		}
		d.MarshalSourceData(filtered, "dependency")
		err = d.AppendEval(eval, "count", transDeps)
		if err != nil {
			return nil, fmt.Errorf("failed to create transitive dependencies digeest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

		if transDeps < 1 {
			d.Warning = true
			d.WarningMessage = "no transitive dependencies found"
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
