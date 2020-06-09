package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func buildsystemsDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, compilersIndex, "compiler", "compilers")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.BuildsystemResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into buildss bytes")
		}

		d.MarshalSourceData(b, "buildsystems")

		switch len(b.Compilers) {
		case 0:
			err := d.AppendEval(eval, "chars", "none detected")
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}
		case 1:
			n := ""
			for _, c := range b.Compilers {
				n = c.Name
			}

			err := d.AppendEval(eval, "chars", n)
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}

			d.UseSingularTitle()
		default:
			err := d.AppendEval(eval, "count", len(b.Compilers))
			if err != nil {
				return nil, fmt.Errorf("failed to create builds digest: %v", err.Error())
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}
