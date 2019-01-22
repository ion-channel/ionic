package digests

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ion-channel/ionic/rulesets"
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
func NewDigests(appliedRuleset *rulesets.AppliedRulesetSummary) ([]Digest, error) {
	ds := make([]Digest, 0)
	errs := make([]string, 0, 0)

	for i := range appliedRuleset.RuleEvaluationSummary.Ruleresults {
		e := appliedRuleset.RuleEvaluationSummary.Ruleresults[i]

		d, err := _newDigests(&e)
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

func _newDigests(eval *scans.Evaluation) ([]Digest, error) {
	err := eval.Translate()
	if err != nil {
		return nil, fmt.Errorf("evaluation translate error: %v", err.Error())
	}

	switch strings.ToLower(eval.TranslatedResults.Type) {
	case "ecosystems":
		return ecosystemsDigests(eval)

	case "dependency":
		return dependencyDigests(eval)

	case "vulnerability":
		return vulnerabilityDigests(eval)

	case "virus":
		return virusDigests(eval)

	case "community":
		return communityDigests(eval)

	case "license":
		return licenseDigests(eval)

	case "coverage":
		return coveragDigests(eval)

	case "about_yml":
		return aboutYMLDigests(eval)

	case "difference":
		return differenceDigests(eval)

	default:
		return nil, fmt.Errorf("Couldn't figure out how to map '%v' to a digest", eval.TranslatedResults.Type)
	}
}

func ecosystemsDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.EcosystemResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into languages bytes")
	}

	dom := getDominantLanguages(b.Ecosystems)

	t := "dominant languages"
	if len(dom) == 1 {
		t = "dominant language"
	}

	d, err := NewFromEval(dominantLanguagesIndex, t, "list", dom, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create dominant language digest: %v", err.Error())
	}

	d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

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

func dependencyDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.DependencyResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into dependency bytes")
	}

	t := "dependencies outdated"
	if b.Meta.UpdateAvailableCount == 1 {
		t = "dependency outdated"
	}

	d, err := NewFromEval(dependencyOutdatedIndex, t, "count", b.Meta.UpdateAvailableCount, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create dependencies outdated digest: %v", err.Error())
	}

	d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

	digests = append(digests, *d)

	t = "dependencies no version specified"
	if b.Meta.NoVersionCount == 1 {
		t = "dependency no version specified"
	}

	d, err = NewFromEval(noVersionIndex, t, "count", b.Meta.NoVersionCount, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create dependencies no version digest: %v", err.Error())
	}
	digests = append(digests, *d)

	t = "direct dependencies"
	if b.Meta.FirstDegreeCount == 1 {
		t = "direct dependency"
	}

	d, err = NewFromEval(directDependencyIndex, t, "count", b.Meta.FirstDegreeCount, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create direct dependencies digeest: %v", err.Error())
	}

	d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

	digests = append(digests, *d)

	transCount := b.Meta.TotalUniqueCount - b.Meta.FirstDegreeCount

	t = "transitive dependencies"
	if transCount == 1 {
		t = "transitive dependency"
	}

	d, err = NewFromEval(transitiveDependencyIndex, t, "count", transCount, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create transitive dependencies digeest: %v", err.Error())
	}

	d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

	digests = append(digests, *d)

	return digests, nil
}

func vulnerabilityDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.VulnerabilityResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
	}

	t := "total vulnerabilities"
	if b.Meta.VulnerabilityCount == 1 {
		t = "total vulnerability"
	}

	d, err := NewFromEval(totalVulnerabilitiesIndex, t, "count", b.Meta.VulnerabilityCount, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create total vulnerabilities digest: %v", err.Error())
	}

	if b.Meta.VulnerabilityCount > 0 {
		d.Warning = true
		d.WarningMessage = "vulnerabilities found"

		if b.Meta.VulnerabilityCount == 1 {
			d.WarningMessage = "vulnerability found"
		}
	}

	digests = append(digests, *d)

	// de-dupe the vulns by `id` field
	ids := make(map[int]bool, 0)

	for i := range b.Vulnerabilities {
		ids[b.Vulnerabilities[i].ID] = true
	}

	t = "unique vulnerabilities"
	if len(ids) == 1 {
		t = "unique vulnerability"
	}

	d, err = NewFromEval(uniqueVulnerabilitiesIndex, t, "count", len(ids), eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create unique vulnerabilities digest: %v", err.Error())
	}
	digests = append(digests, *d)

	return digests, nil
}

func virusDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.VirusResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into virus")
	}

	t := "total files scanned"
	if b.ScannedFiles == 1 {
		t = "total file scanned"
	}

	d, err := NewFromEval(filesScannedIndex, t, "count", b.ScannedFiles, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create total files scanned digest: %v", err.Error())
	}

	d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

	if b.ScannedFiles < 1 {
		d.Warning = true
		d.WarningMessage = "no files were seen"
	}

	digests = append(digests, *d)

	t = "viruses found"
	if b.InfectedFiles == 1 {
		t = "virus found"
	}

	d, err = NewFromEval(virusFoundIndex, t, "count", b.InfectedFiles, eval)
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

	digests = append(digests, *d)

	return digests, nil
}

func communityDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.CommunityResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into community")
	}

	t := "unique committers"
	if b.Committers == 1 {
		t = "unique committer"
	}

	d, err := NewFromEval(uniqueCommittersIndex, t, "count", b.Committers, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create committers digest: %v", err.Error())
	}

	if b.Committers == 1 {
		d.Warning = true
		d.WarningMessage = "single committer repository"
	}

	digests = append(digests, *d)

	return digests, nil
}

func licenseDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.LicenseResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into license")
	}

	licenseList := make([]string, 0)
	for i := range b.Type {
		licenseList = append(licenseList, b.Type[i].Name)
	}

	var d *Digest
	var err error
	if len(licenseList) == 1 {
		d, err = NewFromEval(licensesIndex, "license", "chars", licenseList[0], eval)
		if err != nil {
			return nil, fmt.Errorf("failed to create license list digest: %v", err.Error())
		}
	} else {
		d, err = NewFromEval(licensesIndex, "licenses", "count", len(licenseList), eval)
		if err != nil {
			return nil, fmt.Errorf("failed to create license list digest: %v", err.Error())
		}
	}

	if len(licenseList) < 1 {
		d.Warning = true
		d.WarningMessage = "no licenses found"
	}

	digests = append(digests, *d)

	return digests, nil
}

func coveragDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.CoverageResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into coverage")
	}

	d, err := NewFromEval(codeCoverageIndex, "code coverage", "percent", b.Value, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create code coverage digest: %v", err.Error())
	}

	digests = append(digests, *d)

	return digests, nil
}

func aboutYMLDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.AboutYMLResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into aboutyaml")
	}

	d, err := NewFromEval(aboutYMLIndex, "valid about yaml", "bool", b.Valid, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create valid about yaml digest: %v", err.Error())
	}
	digests = append(digests, *d)

	return digests, nil
}

func differenceDigests(eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	b, ok := eval.TranslatedResults.Data.(scans.DifferenceResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into difference")
	}

	d, err := NewFromEval(differenceIndex, "difference detected", "bool", b.Difference, eval)
	if err != nil {
		return nil, fmt.Errorf("failed to create difference digest: %v", err.Error())
	}

	d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

	digests = append(digests, *d)

	return digests, nil
}
