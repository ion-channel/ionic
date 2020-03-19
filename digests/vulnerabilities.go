package digests

import (
	"fmt"
	"strconv"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func vulnerabilityDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	var vulnCount, uniqVulnCount int
	var highs int
	var crits int
	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VulnerabilityResults)
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

	// total vulns
	d := NewDigest(status, totalVulnerabilitiesIndex, "total vulnerability", "total vulnerabilities")

	if eval != nil && !status.Errored() {
		err := d.AppendEval(eval, "count", vulnCount)
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
		err := d.AppendEval(eval, "count", uniqVulnCount)
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
		err := d.AppendEval(eval, "count", highs)
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
		err := d.AppendEval(eval, "count", crits)
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
