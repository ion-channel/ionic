package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func communityDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, uniqueCommittersIndex, "unique committer", "unique committers")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.CommunityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into community")
		}

		err := d.AppendEval(eval, "count", b.Committers)
		if err != nil {
			return nil, fmt.Errorf("failed to create committers digest: %v", err.Error())
		}

		if b.Committers == 1 {
			d.Warning = true
			d.WarningMessage = "single committer repository"
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
