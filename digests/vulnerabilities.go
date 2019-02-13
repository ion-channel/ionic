package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

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
