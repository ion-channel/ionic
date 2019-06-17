package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func licenseDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
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

		switch len(licenseList) {
		case 1:
			err := d.AppendEval(eval, "chars", licenseList[0])
			if err != nil {
				return nil, fmt.Errorf("failed to create license list digest: %v", err.Error())
			}

			d.UseSingularTitle()
		case 0:
			d.Warning = true
			d.WarningMessage = "no licenses found"
			fallthrough
		default:
			err := d.AppendEval(eval, "count", len(licenseList))
			if err != nil {
				return nil, fmt.Errorf("failed to create license list digest: %v", err.Error())
			}
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
