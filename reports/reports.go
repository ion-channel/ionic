package reports

import (
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/analysis"
	"github.com/ion-channel/ionic/projects"
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

// ProjectReport gives the details of a project including past analyses
type ProjectReport struct {
	*projects.Project
	RulesetName       string             `json:"ruleset_name"`
	AnalysisSummaries []analysis.Summary `json:"analysis_summaries"`
}

// ProjectReports is used for getting a high level overview, returning a single
// analysis
type ProjectReports struct {
	ID              string            `json:"id"`
	TeamID          string            `json:"team_id"`
	RulesetID       string            `json:"ruleset_id"`
	Name            string            `json:"name"`
	Type            string            `json:"type"`
	Source          string            `json:"source"`
	Branch          string            `json:"branch"`
	Description     string            `json:"description"`
	Active          bool              `json:"active"`
	ChatChannel     string            `json:"chat_channel"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeployKey       string            `json:"deploy_key"`
	Monitor         bool              `json:"should_monitor"`
	PocName         string            `json:"poc_name"`
	PocEmail        string            `json:"poc_email"`
	Username        string            `json:"username"`
	Password        string            `json:"password"`
	KeyFingerprint  string            `json:"key_fingerprint"`
	Aliases         []aliases.Alias   `json:"aliases"`
	Tags            []tags.Tag        `json:"tags"`
	RulesetName     string            `json:"ruleset_name"`
	AnalysisSummary *analysis.Summary `json:"analysis_summary"`
}
