package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func communityDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)
	d := NewDigest(status, UniqueCommittersIndex, "unique committer", "unique committers")
	activityDigest := NewDigest(status, CommittedAtIndex, "committed at", "committed at")

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.CommunityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into community")
		}
		d.MarshalSourceData(b, "committers")
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

	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.CommunityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into community")
		}

		activityDigest.MarshalSourceData(b, "committed at")

		var evalTime string
		evalTime = b.CommittedAt.String()
		// check for legacy scans (as they won't have this data availible), otherwise display last time committed at
		if b.CommittedAt.IsZero() {
			evalTime = "N/A"
		}

		err := activityDigest.AppendEval(eval, "chars", evalTime)
		if err != nil {
			return nil, fmt.Errorf("failed to create committed at digest: %v", err.Error())
		}
	}

	digests = append(digests, *activityDigest)
	return digests, nil
}
