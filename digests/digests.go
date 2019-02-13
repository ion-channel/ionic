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

func ecosystemsDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, dominantLanguagesIndex, "dominant language", "dominant languages")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.EcosystemResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into languages bytes")
		}

		dom := getDominantLanguages(b.Ecosystems)

		err := d.AppendEval(eval, "list", dom)
		if err != nil {
			return nil, fmt.Errorf("failed to create dominant language digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}

func getDominantLanguages(languages map[string]int) []string {
	dom := []string{}

	total := 0
	for _, lines := range languages {
		total += lines
	}

	majority := float64(len(languages)-1) / float64(len(languages)) * 100.0

	h := 0.0
	h2 := 0.0
	top := ""
	top2 := ""
	dominant := ""
	for lang, lines := range languages {
		p := float64(lines) / float64(total) * 100.0

		if p > h {
			h = p
			top = lang
		} else {
			if p > h2 {
				h2 = p
				top2 = lang
			}
		}

		if p >= majority {
			dominant = lang
		}
	}

	if dominant != "" {
		dom = append(dom, dominant)
	} else {
		dom = append(dom, top, top2)
	}

	return dom
}

func dependencyDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, dependencyOutdatedIndex, "dependency outdated", "dependencies outdated")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.DependencyResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
		}

		err := d.AppendEval(eval, "count", b.Meta.UpdateAvailableCount)
		if err != nil {
			return nil, fmt.Errorf("failed to create dependencies outdated digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	d = NewDigest(status, noVersionIndex, "dependency no version specified", "dependencies no version specified")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.DependencyResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
		}

		err := d.AppendEval(eval, "count", b.Meta.NoVersionCount)
		if err != nil {
			return nil, fmt.Errorf("failed to create dependencies no version digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, directDependencyIndex, "direct dependency", "direct dependencies")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.DependencyResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
		}

		err := d.AppendEval(eval, "count", b.Meta.FirstDegreeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to create direct dependencies digeest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	d = NewDigest(status, transitiveDependencyIndex, "transitive dependency", "transitive dependencies")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.DependencyResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
		}

		transCount := b.Meta.TotalUniqueCount - b.Meta.FirstDegreeCount

		err := d.AppendEval(eval, "count", transCount)
		if err != nil {
			return nil, fmt.Errorf("failed to create transitive dependencies digeest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}

func vulnerabilityDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, totalVulnerabilitiesIndex, "total vulnerability", "total vulnerabilities")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VulnerabilityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
		}

		err := d.AppendEval(eval, "count", b.Meta.VulnerabilityCount)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to total vulnerabilities digest: %v", err.Error())
		}

		if b.Meta.VulnerabilityCount > 0 {
			d.Warning = true
			d.WarningMessage = "vulnerabilities found"

			if b.Meta.VulnerabilityCount == 1 {
				d.WarningMessage = "vulnerability found"
			}
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, uniqueVulnerabilitiesIndex, "unique vulnerability", "unique vulnerabilities")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VulnerabilityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
		}

		// de-dupe the vulns by `id` field
		ids := make(map[int]bool, 0)
		for i := range b.Vulnerabilities {
			ids[b.Vulnerabilities[i].ID] = true
		}

		err := d.AppendEval(eval, "count", len(ids))
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	return digests, nil
}

func virusDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, filesScannedIndex, "total file scanned", "total files scanned")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VirusResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into virus")
		}

		err := d.AppendEval(eval, "count", b.ScannedFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to create total files scanned digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

		if b.ScannedFiles < 1 {
			d.Warning = true
			d.WarningMessage = "no files were seen"
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, virusFoundIndex, "virus found", "viruses found")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VirusResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into virus")
		}

		err := d.AppendEval(eval, "count", b.InfectedFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to create total files scanned digest: %v", err.Error())
		}

		if b.InfectedFiles > 0 {
			d.Warning = true
			d.WarningMessage = "infected files were seen"

			if b.InfectedFiles == 1 {
				d.WarningMessage = "infected files were seen"
			}
		}
	}

	digests = append(digests, *d)

	return digests, nil
}

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

func licenseDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, licensesIndex, "license found", "licenses found")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.LicenseResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into license")
		}

		licenseList := make([]string, 0)
		for i := range b.Type {
			licenseList = append(licenseList, b.Type[i].Name)
		}

		err := d.AppendEval(eval, "count", len(licenseList))
		if err != nil {
			return nil, fmt.Errorf("failed to create license list digest: %v", err.Error())
		}

		if len(licenseList) < 1 {
			d.Warning = true
			d.WarningMessage = "no licenses found"
		}
	}

	digests = append(digests, *d)

	return digests, nil
}

func coveragDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, codeCoverageIndex, "code coverage", "code coverage")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.CoverageResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into coverage")
		}

		err := d.AppendEval(eval, "percent", b.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to create code coverage digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	return digests, nil
}

func aboutYMLDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, aboutYMLIndex, "valid about yaml", "valid about yaml")

	if eval != nil {
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

func differenceDigests(eval *scans.Evaluation, status *scanner.ScanStatus) ([]Digest, error) {
	digests := make([]Digest, 0)

	d := NewDigest(status, differenceIndex, "difference detected", "difference detected")

	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.DifferenceResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into difference")
		}

		err := d.AppendEval(eval, "bool", b.Difference)
		if err != nil {
			return nil, fmt.Errorf("failed to create difference digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	return digests, nil
}
