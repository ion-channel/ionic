package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func coveragDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, codeCoverageIndex, "code coverage", "code coverage")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.CoverageResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into coverage")
		}

		err := d.AppendEval(eval, "percent", b.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to create code coverage digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
