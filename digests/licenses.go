package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

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
