package digests

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func communityDigests(status *scanner.ScanStatus, evals []*scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	for _, eval := range evals {
		if len(evals) > 1 {
			// we have both types of community evals present for unique committers and days since last commit
			// evaluate them both
			if eval.RuleID == "2981e1b0-0c8f-0137-8fe7-186590d3c755" {
				d, err := communityCommittersDigest(status, eval)
				if err != nil {
					return nil, err
				}
				digests = append(digests, *d)

			}
			// if rule(s) for 30, 90, 180 days, 1 year since last activity create and evaluate for days since last commit
			if eval.RuleID == "d3b66d48-40a1-11eb-b378-0242ac130002" || eval.RuleID == "6db01715-9e9e-4ff9-bd15-1fcd776d81b8" || eval.RuleID == "21353ecd-b3a9-411d-ab93-c868cdf2c69c" || eval.RuleID == "efcb4ae5-ff36-413a-962b-3f4d2170be2a" {
				d, err := communityCommittedAtDigest(status, eval)
				if err != nil {
					return nil, err
				}
				digests = append(digests, *d)
			}
		} else {
			// we have one OR the other, or no evaluation but we still want both digests created.
			d, err := communityCommittersDigest(status, eval)
			if err != nil {
				return nil, err
			}
			digests = append(digests, *d)

			d, err = communityCommittedAtDigest(status, eval)
			if err != nil {
				return nil, err
			}
			digests = append(digests, *d)

		}

	}
	return digests, nil
}

func communityCommittersDigest(status *scanner.ScanStatus, eval *scans.Evaluation) (*Digest, error) {
	d := NewDigest(status, UniqueCommittersIndex, "unique committer", "unique committers")
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

		if eval.RuleID != "2981e1b0-0c8f-0137-8fe7-186590d3c755" {
			d.Evaluated = false
		}
	}
	return d, nil
}

func communityCommittedAtDigest(status *scanner.ScanStatus, eval *scans.Evaluation) (*Digest, error) {
	activityDigest := NewDigest(status, CommittedAtIndex, "days since last commit", "days since last commit")
	if eval != nil && !status.Errored() {
		b, ok := eval.TranslatedResults.Data.(scans.CommunityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into community")
		}

		activityDigest.MarshalSourceData(b, "committed at")

		var evalDaysSinceLastCommit string
		committedAt := b.CommittedAt
		now := time.Now()
		difference := now.Sub(committedAt)
		daysSinceLastCommit := int(difference.Hours() / 24)
		evalDaysSinceLastCommit = strconv.Itoa(daysSinceLastCommit)

		// check for legacy scans (as they won't have this data available), otherwise display days since last committed at
		if b.CommittedAt.IsZero() {
			evalDaysSinceLastCommit = "N/A"
		}

		err := activityDigest.AppendEval(eval, "chars", evalDaysSinceLastCommit)
		if err != nil {
			return nil, fmt.Errorf("failed to create committed at digest: %v", err.Error())
		}

		// evaluate if our Rule for days since last commit is present
		if eval.RuleID != "d3b66d48-40a1-11eb-b378-0242ac130002" && eval.RuleID != "6db01715-9e9e-4ff9-bd15-1fcd776d81b8" && eval.RuleID != "21353ecd-b3a9-411d-ab93-c868cdf2c69c" && eval.RuleID != "efcb4ae5-ff36-413a-962b-3f4d2170be2a" {
			activityDigest.Evaluated = false
		}
	}
	return activityDigest, nil

}
