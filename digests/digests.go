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
	//DifferenceIndex represents the index number for this digest type
	DifferenceIndex = iota
	//VirusFoundIndex represents the index number for this digest type
	VirusFoundIndex
	//CriticalVulnerabilitiesIndex represents the index number for this digest type
	CriticalVulnerabilitiesIndex
	//HighVulnerabilitiesIndex represents the index number for this digest type
	HighVulnerabilitiesIndex
	//TotalVulnerabilitiesIndex represents the index number for this digest type
	TotalVulnerabilitiesIndex
	//UniqueVulnerabilitiesIndex represents the index number for this digest type
	UniqueVulnerabilitiesIndex
	//LicensesIndex represents the index number for this digest type
	LicensesIndex
	//FilesScannedIndex represents the index number for this digest type
	FilesScannedIndex
	//DirectDependencyIndex represents the index number for this digest type
	DirectDependencyIndex
	//TransitiveDependencyIndex represents the index number for this digest type
	TransitiveDependencyIndex
	//DependencyOutdatedIndex represents the index number for this digest type
	DependencyOutdatedIndex
	//NoVersionIndex represents the index number for this digest type
	NoVersionIndex
	//CompilersIndex represents the index number for this digest type
	CompilersIndex
	//ContainerImagesIndex represents the index number for this digest type
	ContainerImagesIndex
	//ContainerDependenciesIndex represents the index number for this digest type
	ContainerDependenciesIndex
	//LanguagesIndex represents the index number for this digest type
	LanguagesIndex
	//UniqueCommittersIndex represents the index number for this digest type
	UniqueCommittersIndex
	//CodeCoverageIndex represents the index number for this digest type
	CodeCoverageIndex
	//CommittedAtIndex represents the index number for this digest type
	CommittedAtIndex
)

// GroupedDigests represents an organized grouped report of digests
type GroupedDigests struct {
	Name    string   `json:"name"`
	Digests []Digest `json:"digests"`
}

// DigestReport is the report of organized grouped and ungrouped digests
type DigestReport struct {
	SingleDigests  *[]Digest         `json:"single_digests,omitempty,omitnull"`
	GroupedDigests *[]GroupedDigests `json:"grouped_digests,omitempty,omitnull"`
}

// NewDigests takes an applied ruleset and returns the relevant digests derived
// from all the evaluations in it, and any errors it encounters.
func NewDigests(appliedRuleset *rulesets.AppliedRulesetSummary, statuses []scanner.ScanStatus) ([]Digest, error) {
	ds := make([]Digest, 0)
	errs := make([]string, 0, 0)

	for i := range statuses {
		s := statuses[i]

		var e *scans.Evaluation
		var evals []*scans.Evaluation

		if appliedRuleset != nil && appliedRuleset.RuleEvaluationSummary != nil {
			for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
				// a scan type may have multiple rule evaluations
				if appliedRuleset.RuleEvaluationSummary.Ruleresults[i].ID == s.ID {
					e = &appliedRuleset.RuleEvaluationSummary.Ruleresults[i]
					evals = append(evals, e)
				}
			}
		}

		// if we didn't find an matching ruleset evals, we still want the digests for scanstatus even if a ruleset doesnt match
		if len(evals) == 0 {
			evals = append(evals, e)
		}

		for i := range evals {
			ev := evals[i]
			d, err := _newDigests(&s, ev)

			if err != nil {
				errs = append(errs, fmt.Sprintf("failed to make digest(s) from scan: %v", err.Error()))
				continue
			}

			if d != nil {
				ds = append(ds, d...)
			}
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

	case "buildsystems":
		return buildsystemsDigests(status, eval)

	case "about_yml", "file_type":
		return nil, nil

	default:
		return nil, fmt.Errorf("Couldn't figure out how to map '%v' to a digest", status.Name)
	}
}
