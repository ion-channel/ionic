package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func virusDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
	digests := make([]Digest, 0)

	var scannedFiles, infectedFiles int
	if eval != nil {
		b, ok := eval.TranslatedResults.Data.(scans.VirusResults)
		if !ok {
			return nil, fmt.Errorf("error coercing evaluation translated results into virus")
		}

		scannedFiles = b.ScannedFiles
		infectedFiles = b.InfectedFiles
	}

	d := NewDigest(status, filesScannedIndex, "total file scanned", "total files scanned")

	if eval != nil && !status.Errored() {
		err := d.AppendEval(eval, "count", scannedFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to create total files scanned digest: %v", err.Error())
		}

		d.Evaluated = false // As of now there's no rule to evaluate this against so it's set to not evaluated.

		if scannedFiles < 1 {
			d.Warning = true
			d.WarningMessage = "no files were seen"
		}
	}

	digests = append(digests, *d)

	d = NewDigest(status, virusFoundIndex, "virus found", "viruses found")

	if eval != nil && !status.Errored() {
		err := d.AppendEval(eval, "count", infectedFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to create total files scanned digest: %v", err.Error())
		}

		if infectedFiles > 0 {
			d.Warning = true
			d.WarningMessage = "infected files were seen"
		}
	}

	digests = append(digests, *d)

	return digests, nil
}
