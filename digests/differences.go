package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func differenceDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, differenceIndex, "difference detected", "difference detected")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.DifferenceResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into difference")
		}

		err := d.AppendEval(eval, "bool", b.Difference)
		if err != nil {
			return nil, fmt.Errorf("failed to create difference digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}
