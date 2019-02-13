package digests

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

const (
	differenceIndex = iota
	virusFoundIndex
	totalVulnerabilitiesIndex
	uniqueVulnerabilitiesIndex
	licensesIndex
	filesScannedIndex
	directDependencyIndex
	transitiveDependencyIndex
	dependencyOutdatedIndex
	noVersionIndex
	aboutYMLIndex
	dominantLanguagesIndex
	uniqueCommittersIndex
	codeCoverageIndex
)

// NewDigests takes an applied ruleset and returns the relevant digests derived
// from all the evaluations in it, and any errors it encounters.
func NewDigests(appliedRuleset *rulesets.AppliedRulesetSummary, statuses []scanner.ScanStatus) ([]Digest, error) {
	ds := make([]Digest, 0)
	errs := make([]string, 0, 0)

	for i := range statuses {
		s := statuses[i]

		var e *scans.Evaluation
		for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
			if appliedRuleset.RuleEvaluationSummary.Ruleresults[i].ID == s.ID {
				e = &appliedRuleset.RuleEvaluationSummary.Ruleresults[i]
				break
			}
		}

		d, err := _newDigests(e, &s)
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to make digest(s) from scan: %v", err.Error()))
			continue
		}

		ds = append(ds, d...)
	}

	sort.Slice(ds, func(i, j int) bool { return ds[i].Index < ds[j].Index })

	if len(errs) > 0 {
		return ds, fmt.Errorf("failed to make some digests: %v", strings.Join(errs, "; "))
	}

	return ds, nil
}

func _newDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	err := eval.Translate()
	if err != nil {
		return nil, fmt.Errorf("evaluation translate error: %v", err.Error())
	}

	switch strings.ToLower(eval.TranslatedResults.Type) {
	case "ecosystems":
		return ecosystemsDigests(eval, status)

	case "dependency":
		return dependencyDigests(eval, status)

	case "vulnerability":
		return vulnerabilityDigests(eval, status)

	case "virus":
		return virusDigests(eval, status)

	case "community":
		return communityDigests(eval, status)

	case "license":
		return licenseDigests(eval, status)

	case "coverage":
		return coveragDigests(eval, status)

	case "about_yml":
		return aboutYMLDigests(eval, status)

	case "difference":
		return differenceDigests(eval, status)

	default:
		return nil, fmt.Errorf("Couldn't figure out how to map '%v' to a digest", eval.TranslatedResults.Type)
	}
}
