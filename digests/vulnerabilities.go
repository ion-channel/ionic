package digests

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

const (
	scoreVersion2 = "2.0"
	scoreHigh     = 7.0
	scoreCritical = 9.0
)

// ByScore sort interface impl for sorting vulnerabilities by score
type ByScore []scans.VulnerabilityResultsVulnerability

func (v ByScore) Len() int      { return len(v) }
func (v ByScore) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v ByScore) Less(i, j int) bool {
	iscore, err := strconv.ParseFloat(v[i].Score, 32)
	if err != nil {
		return false
	}

	jscore, err := strconv.ParseFloat(v[j].Score, 32)
	if err != nil {
		return true
	}
	return iscore > jscore
}

type vfilter func(*scans.VulnerabilityResultsVulnerability) bool

func all(v *scans.VulnerabilityResultsVulnerability) bool {
	return true
}

func high(v *scans.VulnerabilityResultsVulnerability) bool {
	score, err := strconv.ParseFloat(v.Score, 32)
	if err != nil {
		return false
	}
	if v.ScoreVersion == scoreVersion2 && score >= scoreHigh {
		return true
	}
	if v.ScoreVersion != scoreVersion2 && score >= scoreHigh && score < scoreCritical {
		return true
	}
	return false
}

func critical(v *scans.VulnerabilityResultsVulnerability) bool {
	score, err := strconv.ParseFloat(v.Score, 32)
	if err != nil {
		return false
	}
	if v.ScoreVersion != scoreVersion2 && score >= scoreCritical {
		return true
	}
	return false
}

func pivotToVulnerabilities(data interface{}, unique bool, f vfilter) ([]scans.VulnerabilityResultsVulnerability, error) {
	b, ok := data.(scans.VulnerabilityResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
	}

	var values []scans.VulnerabilityResultsVulnerability
	uu := map[string]*scans.VulnerabilityResultsVulnerability{}
	count := 0
	values = []scans.VulnerabilityResultsVulnerability{}
	for _, p := range b.Vulnerabilities {
		pp := p
		pp.Vulnerabilities = []scans.VulnerabilityResultsVulnerability{}
		for _, v := range p.Vulnerabilities {
			if f(&v) {
				key := v.ExternalID
				if !unique {
					key = fmt.Sprintf("%v:%d", key, rand.Int())
				}
				if uu[key] == nil {
					v.Dependencies = append(v.Dependencies, pp)
					values = append(values, v)
					uu[key] = &v
				}
				count++
			}
		}
	}

	sort.Sort(ByScore(values))
	return values, nil
}

func vulnerabilityDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := []Digest{}

	var vulnCount int
	var uniqVulnCount int
	var highs int
	var crits int
	var data interface{}
	var results scans.VulnerabilityResults
	if eval != nil {
		eval.Translate()
		data = eval.TranslatedResults.Data
		b, ok := data.(scans.VulnerabilityResults)
		results = b
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
		}

		vulnCount = results.Meta.VulnerabilityCount

		ids := make(map[int]bool, 0)

		for i := range results.Vulnerabilities {
			for j := range results.Vulnerabilities[i].Vulnerabilities {
				v := results.Vulnerabilities[i].Vulnerabilities[j]
				ids[v.ID] = true

				if v.ScoreSystem == "NPM" {
					if npmScore, err := strconv.ParseFloat(v.Score, 32); err == nil {
						if npmScore > 7 { // 10, 9, 8
							crits++
						} else if npmScore > 5 { // 7, 6
							highs++
						}
					}
				} else {
					ver, _ := strconv.ParseFloat(v.ScoreVersion, 32)
					switch int(ver) {
					case 3:
						if v.ScoreDetails.CVSSv3 != nil && v.ScoreDetails.CVSSv3.BaseScore >= 9.0 {
							crits++
						} else if v.ScoreDetails.CVSSv3 != nil && v.ScoreDetails.CVSSv3.BaseScore >= 7.0 {
							highs++
						}
					case 2:
						if v.ScoreDetails.CVSSv2 != nil && v.ScoreDetails.CVSSv2.BaseScore >= 7.0 {
							highs++
						}
					default:
					}
				}

			}
		}

		uniqVulnCount = len(ids)
	}

	d := NewDigest(status, TotalVulnerabilitiesIndex, "total vulnerability", "total vulnerabilities")
	d.MarshalSourceData(nil, "vulnerability")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, false, all)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		d.MarshalSourceData(pivoted, "vulnerability")
		err = d.AppendEval(eval, "count", vulnCount)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to total vulnerabilities digest: %v", err.Error())
		}

		if vulnCount > 0 {
			d.Warning = true
			d.WarningMessage = "vulnerabilities found"

			if vulnCount == 1 {
				d.WarningMessage = "vulnerability found"
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	// unique vulns
	d = NewDigest(status, UniqueVulnerabilitiesIndex, "unique vulnerability", "unique vulnerabilities")
	d.MarshalSourceData(nil, "vulnerability")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, true, all)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		d.MarshalSourceData(pivoted, "vulnerability")
		err = d.AppendEval(eval, "count", uniqVulnCount)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}

		if uniqVulnCount > 0 {
			d.Warning = true
			d.WarningMessage = "vulnerabilities found"

			if uniqVulnCount == 1 {
				d.WarningMessage = "vulnerability found"
			}
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.
	}

	digests = append(digests, *d)

	// high vulns
	d = NewDigest(status, HighVulnerabilitiesIndex, "high vulnerability", "high vulnerabilities")
	d.MarshalSourceData(nil, "vulnerability")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, false, high)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		d.MarshalSourceData(pivoted, "vulnerability")

		err = d.AppendEval(eval, "count", highs)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}

		if highs == 0 {
			d.Passed = true
		}
	}

	digests = append(digests, *d)

	// critical vulns
	d = NewDigest(status, CriticalVulnerabilitiesIndex, "critical vulnerability", "critical vulnerabilities")
	d.MarshalSourceData(nil, "vulnerability")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, false, critical)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		d.MarshalSourceData(pivoted, "vulnerability")

		err = d.AppendEval(eval, "count", crits)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}

		if crits == 0 {
			d.Passed = true
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
