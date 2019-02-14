package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func vulnerabilityDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	var vulnCount, uniqVulnCount int
	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VulnerabilityResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into vuln")
		}

		vulnCount = b.Meta.VulnerabilityCount

		ids := make(map[int]bool, 0)
		for i := range b.Vulnerabilities {
			ids[b.Vulnerabilities[i].ID] = true
		}

		uniqVulnCount = len(ids)
	}

	d := NewDigest(status, totalVulnerabilitiesIndex, "total vulnerability", "total vulnerabilities")

	if eval != nil {
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
	}

	digests = append(digests, *d)

	d = NewDigest(status, uniqueVulnerabilitiesIndex, "unique vulnerability", "unique vulnerabilities")

	if eval != nil {
		err := d.AppendEval(eval, "count", uniqVulnCount)
		if err != nil {
			return nil, fmt.Errorf("failed to add evaluation data to unique vulnerabilities digest: %v", err.Error())
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
