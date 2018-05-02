package reports

import (
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

	summaries := make([]scans.Summary, 0, len(a.ScanSummaries))
	for i := range a.ScanSummaries {
		scan := a.ScanSummaries[i]
		scan.Translate()
		a.ScanSummaries[i] = scan

		summary := scans.NewSummary(&a.ScanSummaries[i])
		summaries = append(summaries, *summary)
	}

	ar.ScanSummaries = summaries

	return &ar, nil
}
