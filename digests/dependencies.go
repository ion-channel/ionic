package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

type dfilter func(*scans.Dependency) *scans.Dependency

func nv(d *scans.Dependency) *scans.Dependency {
	if d.Requirement == "" || d.Requirement == ">= 0" || d.Requirement == "\u003e= 0" {
		return d
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
	if d.Version < d.LatestVersion {
		return d
	}
	return nil
}

func giveem(d *scans.Dependency) *scans.Dependency {
	return d
}

func direct(d *scans.Dependency) *scans.Dependency {
	ff := d
	ff.Dependencies = nil
	return ff
}

func filterDependencies(data interface{}, unique bool, f dfilter) ([]scans.Dependency, error) {
	b, ok := data.(scans.DependencyResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into dep")
	}
	ds := []scans.Dependency{}
	for _, dr := range b.Dependencies {

		filtered := f(&dr)
		if filtered != nil {

			ds = append(ds, *filtered)
		}
	}
	return ds, nil
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

	d := NewDigest(status, dependencyOutdatedIndex, "dependency outdated", "dependencies outdated")

	if eval != nil && !status.Errored() {
		filtered, err := filterDependencies(data, false, od)
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
	d = NewDigest(status, noVersionIndex, "dependency no version specified", "dependencies no version specified")

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

	d = NewDigest(status, directDependencyIndex, "direct dependency", "direct dependencies")

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

	d = NewDigest(status, transitiveDependencyIndex, "transitive dependency", "transitive dependencies")

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
