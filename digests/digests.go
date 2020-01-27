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
	criticalVulnerabilitiesIndex
	highVulnerabilitiesIndex
	totalVulnerabilitiesIndex
	uniqueVulnerabilitiesIndex
	licensesIndex
	filesScannedIndex
	directDependencyIndex
	transitiveDependencyIndex
	dependencyOutdatedIndex
	noVersionIndex
	languagesIndex
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
		if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
			for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
				if appliedRuleset.RuleEvaluationSummary.Ruleresults[i].ID == s.ID {
					e = &appliedRuleset.RuleEvaluationSummary.Ruleresults[i]
					break
				}
			}
		}

		d, err := _newDigests(&s, e)
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to make digest(s) from scan: %v", err.Error()))
			continue
		}

		if d != nil {
			ds = append(ds, d...)
		}
	}

	sort.Slice(ds, func(i, j int) bool { return ds[i].Index < ds[j].Index })

	if len(errs) > 0 {
		return ds, fmt.Errorf("failed to make some digests: %v", strings.Join(errs, "; "))
	}

	return ds, nil
}

func _newDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	if eval != nil {
		err := eval.Translate()
		if err != nil {
			return nil, fmt.Errorf("evaluation translate error: %v", err.Error())
		}
	}

	switch strings.ToLower(status.Name) {
	case "ecosystems":
		return ecosystemsDigests(status, eval)

	case "dependency":
		return dependencyDigests(status, eval)

	case "vulnerability":
		return vulnerabilityDigests(status, eval)

	case "virus":
		return virusDigests(status, eval)

	case "community":
		return communityDigests(status, eval)

	case "license":
		return licenseDigests(status, eval)

	case "external_coverage", "code_coverage", "coverage":
		return coveragDigests(status, eval)

	case "difference":
		return differenceDigests(status, eval)

	case "about_yml", "file_type":
		return nil, nil

	default:
		return nil, fmt.Errorf("Couldn't figure out how to map '%v' to a digest", status.Name)
	}
}
