package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func dependencyDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	var updateAvailable, noVersions, directDeps, transDeps int
	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.DependencyResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
		}

		updateAvailable = b.Meta.UpdateAvailableCount
		noVersions = b.Meta.NoVersionCount
		directDeps = b.Meta.FirstDegreeCount
		transDeps = b.Meta.TotalUniqueCount - b.Meta.FirstDegreeCount
	}

	d := NewDigest(status, dependencyOutdatedIndex, "dependency outdated", "dependencies outdated")

	if eval != nil {
		err := d.AppendEval(eval, "count", updateAvailable)
		if err != nil {
			return nil, fmt.Errorf("failed to create dependencies outdated digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	d = NewDigest(status, noVersionIndex, "dependency no version specified", "dependencies no version specified")

	if eval != nil {
		err := d.AppendEval(eval, "count", noVersions)
		if err != nil {
			return nil, fmt.Errorf("failed to create dependencies no version digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, directDependencyIndex, "direct dependency", "direct dependencies")

	if eval != nil {
		err := d.AppendEval(eval, "count", directDeps)
		if err != nil {
			return nil, fmt.Errorf("failed to create direct dependencies digeest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	d = NewDigest(status, transitiveDependencyIndex, "transitive dependency", "transitive dependencies")

	if eval != nil {
		err := d.AppendEval(eval, "count", transDeps)
		if err != nil {
			return nil, fmt.Errorf("failed to create transitive dependencies digeest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}
