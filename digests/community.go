package digests

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func communityDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)
	d := NewDigest(status, UniqueCommittersIndex, "unique committer", "unique committers")
	activityDigest := NewDigest(status, CommittedAtIndex, "days since last commit", "days since last commit")

	if eval != nil && !status.Errored() && eval.RuleID != "d3b66d48-40a1-11eb-b378-0242ac130002" && eval.RuleID != "6db01715-9e9e-4ff9-bd15-1fcd776d81b8" && eval.RuleID != "efcb4ae5-ff36-413a-962b-3f4d2170be2a" {
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

		// only evaluate if our ruleID for community scan single committers is present
		// if eval.RuleID != "2981e1b0-0c8f-0137-8fe7-186590d3c755" {
		// 	d.Evaluated = false
		// }
		// For community scans and digest
		// Only evaluate this rule if one of 1 year, 90, or 30 days rule is not present otherwise we will
		// if eval.RuleID == "d3b66d48-40a1-11eb-b378-0242ac130002" || eval.RuleID == "6db01715-9e9e-4ff9-bd15-1fcd776d81b8" || eval.RuleID == "efcb4ae5-ff36-413a-962b-3f4d2170be2a" {
		// 	activityDigest.Evaluated = false
		// 	activityDigest.Warning = false
		// }
		digests = append(digests, *d)
	}

	if eval != nil && !status.Errored() && eval.RuleID != "2981e1b0-0c8f-0137-8fe7-186590d3c755" {
		b, ok := eval.TranslatedResults.Data.(scans.CommunityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into community")
		}

		activityDigest.MarshalSourceData(b, "committed at")

		var evalDaysSinceLastCommit string
		committedAt := b.CommittedAt
		now := time.Now()
		difference := now.Sub(committedAt)
		daysSinceLastCommmit := int(difference.Hours() / 24)
		evalDaysSinceLastCommit = strconv.Itoa(daysSinceLastCommmit)

		// check for legacy scans (as they won't have this data availible), otherwise display days since last committed at
		if b.CommittedAt.IsZero() {
			evalDaysSinceLastCommit = "N/A"
		}

		err := activityDigest.AppendEval(eval, "chars", evalDaysSinceLastCommit)
		if err != nil {
			return nil, fmt.Errorf("failed to create committed at digest: %v", err.Error())
		}

		// For community scans and digest
		// Only evaluate this rule if one of 1 year, 90, or 30 days rule is present otherwise we won't
		if eval.RuleID != "d3b66d48-40a1-11eb-b378-0242ac130002" && eval.RuleID != "6db01715-9e9e-4ff9-bd15-1fcd776d81b8" && eval.RuleID != "efcb4ae5-ff36-413a-962b-3f4d2170be2a" {
			activityDigest.Evaluated = false
		}
		digests = append(digests, *activityDigest)
	}

	return digests, nil
}
