package digests

import (
	"fmt"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

func ecosystemsDigests(status *scanner.ScanStatus, eval *scans.Evaluation) ([]Digest, error) {
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
