package reports

import (
	"github.com/ion-channel/ionic/analysis"
	"github.com/ion-channel/ionic/projects"
)

// ProjectReport gives the details of a project including past analyses
type ProjectReport struct {
	*projects.Project
	RulesetName       string             `json:"ruleset_name"`
	AnalysisSummaries []analysis.Summary `json:"analysis_summaries"`
}

// ProjectReports is used for getting a high level overview, returning a single
// analysis
type ProjectReports struct {
	*projects.Project
	RulesetName     string            `json:"ruleset_name"`
	AnalysisSummary *analysis.Summary `json:"analysis_summary"`
}
