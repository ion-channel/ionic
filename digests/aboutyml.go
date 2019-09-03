package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func aboutYMLDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, aboutYMLIndex, "valid about yaml", "valid about yaml")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.AboutYMLResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into aboutyaml")
		}

		err := d.AppendEval(eval, "bool", b.Valid)
		if err != nil {
			return nil, fmt.Errorf("failed to create valid about yaml digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
