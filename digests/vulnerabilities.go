package digests

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"

	"github.com/ion-channel/ionic/products"
	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
	"github.com/ion-channel/ionic/vulnerabilities"
)

const (
	scoreVersion2 = "2.0"
	scoreHigh     = 7.0
	scoreCritical = 9.0
)

// ByScore sort interface impl for sorting vulnerabilities by score
type ByScore []vulnerabilities.Vulnerability

func (v ByScore) Len() int           { return len(v) }
func (v ByScore) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByScore) Less(i, j int) bool { return v[i].Score > v[j].Score }

type filter func(*vulnerabilities.Vulnerability) bool

func all(v *vulnerabilities.Vulnerability) bool {
	return true
}

func high(v *vulnerabilities.Vulnerability) bool {
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

func critical(v *vulnerabilities.Vulnerability) bool {
	score, err := strconv.ParseFloat(v.Score, 32)
	if err != nil {
		return false
	}
	if v.ScoreVersion != scoreVersion2 && score >= scoreCritical {
		return true
	}
	return false
}

func pivotToVulnerabilities(data interface{}, unique bool, f filter) ([]vulnerabilities.Vulnerability, error) {
	b, ok := data.(scans.VulnerabilityResults)
	if !ok {
		return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
	}

	uu := map[string]*vulnerabilities.Vulnerability{}
	for _, p := range b.Vulnerabilities {
		for _, v := range p.Vulnerabilities {
			vv := v

			key := vv.ExternalID

			if !unique {
				key = fmt.Sprintf("%d", rand.Int())
			}
			if uu[key] == nil {
				uu[key] = &vv
			}
			d := products.Product{
				ExternalID: p.ExternalID,
				Name:       p.Name,
				Org:        p.Org,
				Version:    p.Version,
			}
			uu[key].Dependencies = append(uu[key].Dependencies, d)
		}
	}
	values := []vulnerabilities.Vulnerability{}
	for _, v := range uu {
		if f(v) {
			values = append(values, *v)
		}
	}
	sort.Sort(ByScore(values))
	return values, nil
}

func vulnerabilityDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	var vulnCount, uniqVulnCount int
	var highs int
	var crits int
	var data interface{}
	if eval != nil {
		data = eval.TranslatedResults.Data
		b, ok := data.(scans.VulnerabilityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
		}

		vulnCount = b.Meta.VulnerabilityCount

		ids := make(map[int]bool, 0)

		for i := range b.Vulnerabilities {
			for j := range b.Vulnerabilities[i].Vulnerabilities {
				v := b.Vulnerabilities[i].Vulnerabilities[j]
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
					switch v.ScoreVersion {
					case "3.0":
						if v.ScoreDetails.CVSSv3 != nil && v.ScoreDetails.CVSSv3.BaseScore >= 9.0 {
							crits++
						} else if v.ScoreDetails.CVSSv3 != nil && v.ScoreDetails.CVSSv3.BaseScore >= 7.0 {
							highs++
						}
					case "2.0":
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

	// total vulns
	d := NewDigest(status, totalVulnerabilitiesIndex, "total vulnerability", "total vulnerabilities")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, false, all)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		eval.TranslatedResults.Data = pivoted
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
	d = NewDigest(status, uniqueVulnerabilitiesIndex, "unique vulnerability", "unique vulnerabilities")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, true, all)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		eval.TranslatedResults.Data = pivoted
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
	d = NewDigest(status, highVulnerabilitiesIndex, "high vulnerability", "high vulnerabilities")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, true, high)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		eval.TranslatedResults.Data = pivoted

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
	d = NewDigest(status, criticalVulnerabilitiesIndex, "critical vulnerability", "critical vulnerabilities")

	if eval != nil && !status.Errored() {
		pivoted, err := pivotToVulnerabilities(data, true, critical)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
		eval.TranslatedResults.Data = pivoted

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
