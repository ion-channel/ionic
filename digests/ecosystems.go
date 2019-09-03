package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func ecosystemsDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, languagesIndex, "language", "languages")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.EcosystemResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into languages bytes")
		}

		switch len(b.Ecosystems) {
		case 0:
			err := d.AppendEval(eval, "chars", "none detected")
			if err != nil {
				return nil, fmt.Errorf("failed to create language digest: %v", err.Error())
			}
		case 1:
			lang := ""
			for lang = range b.Ecosystems {
			}

			err := d.AppendEval(eval, "chars", lang)
			if err != nil {
				return nil, fmt.Errorf("failed to create language digest: %v", err.Error())
			}

			d.UseSingularTitle()
		default:
			err := d.AppendEval(eval, "count", len(b.Ecosystems))
			if err != nil {
				return nil, fmt.Errorf("failed to create language digest: %v", err.Error())
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}
