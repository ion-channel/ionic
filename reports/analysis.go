package reports

import (
	"encoding/json"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/analysis"
	"github.com/ion-channel/ionic/scans"
	"github.com/ion-channel/ionic/tags"
)

// AnalysisReport is a Ion Channel representation of a report output from a
// given analysis
type AnalysisReport struct {
	*analysis.Analysis
	RulesetName   string          `json:"ruleset_name" xml:"ruleset_name"`
	Passed        bool            `json:"passed" xml:"passed"`
	Aliases       []aliases.Alias `json:"aliases"`
	Tags          []tags.Tag      `json:"tags"`
	Trigger       string          `json:"trigger" xml:"trigger"`
	Risk          string          `json:"risk" xml:"risk"`
	Summary       string          `json:"summary" xml:"summary"`
	ScanSummaries []scans.Summary `json:"scan_summaries" xml:"scan_summaries"`
}

// NewAnalysisReport takes an Analysis and returns an initialized AnalysisReport
func NewAnalysisReport(a *analysis.Analysis) (*AnalysisReport, error) {
	ar := AnalysisReport{Analysis: a}

	for _, oldScan := range a.ScanSummaries {
		if oldScan.TranslatedResults == nil && oldScan.UntranslatedResults != nil {
			translatedResults := oldScan.UntranslatedResults.Translate()

			newScan := oldScan
			newScan.UntranslatedResults = nil
			newScan.TranslatedResults = translatedResults
			b, err := json.Marshal(translatedResults)
			if err != nil {
				return nil, err
			}

			newScan.Results = b

			summary := scans.NewSummary(&newScan)
			ar.ScanSummaries = append(ar.ScanSummaries, *summary)
		} else {
			// already translated
			summary := scans.NewSummary(&oldScan)
			ar.ScanSummaries = append(ar.ScanSummaries, *summary)
		}
	}

	return &ar, nil
}
