package reports

import (
	"time"

	"github.com/ion-channel/ionic/analysis"
	"github.com/ion-channel/ionic/scans"
)

// AnalysisReport is a Ion Channel representation of a report output from a given analysis
type AnalysisReport struct {
	ID            string              `json:"id" xml:"id"`
	TeamID        string              `json:"team_id" xml:"team_id"`
	ProjectID     string              `json:"project_id" xml:"project_id"`
	BuildNumber   string              `json:"build_number" xml:"build_number"`
	Name          string              `json:"name" xml:"name"`
	Text          string              `json:"text" xml:"text"`
	Type          string              `json:"type" xml:"type"`
	Source        string              `json:"source" xml:"source"`
	Branch        string              `json:"branch" xml:"branch"`
	Description   string              `json:"description" xml:"description"`
	Status        string              `json:"status" xml:"status"`
	RulesetID     string              `json:"ruleset_id" xml:"ruleset_id"`
	RulesetName   string              `json:"ruleset_name" xml:"ruleset_name"`
	Passed        bool                `json:"passed" xml:"passed"`
	CreatedAt     time.Time           `json:"created_at" xml:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at" xml:"updated_at"`
	Duration      float64             `json:"duration" xml:"duration"`
	Trigger       string              `json:"trigger" xml:"trigger"`
	TriggerHash   string              `json:"trigger_hash" xml:"trigger_hash"`
	TriggerText   string              `json:"trigger_text" xml:"trigger_text"`
	TriggerAuthor string              `json:"trigger_author" xml:"trigger_author"`
	Risk          string              `json:"risk" xml:"risk"`
	Summary       string              `json:"summary" xml:"summary"`
	ScanSummaries []scans.ScanSummary `json:"scan_summaries" xml:"scan_summaries"`
}

// ProjectReport gives the details of a project including past analyses
type ProjectReport struct {
	ID                string                     `json:"id"`
	TeamID            string                     `json:"team_id"`
	RulesetID         string                     `json:"ruleset_id"`
	Name              string                     `json:"name"`
	Type              string                     `json:"type"`
	Source            string                     `json:"source"`
	Branch            string                     `json:"branch"`
	Description       string                     `json:"description"`
	Active            bool                       `json:"active"`
	ChatChannel       string                     `json:"chat_channel"`
	CreatedAt         time.Time                  `json:"created_at"`
	UpdatedAt         time.Time                  `json:"updated_at"`
	DeployKey         string                     `json:"deploy_key"`
	Monitor           bool                       `json:"should_monitor"`
	PocName           string                     `json:"poc_name"`
	PocEmail          string                     `json:"poc_email"`
	Username          string                     `json:"username"`
	Password          string                     `json:"password"`
	KeyFingerprint    string                     `json:"key_fingerprint"`
	PocNameHash       string                     `json:"poc_name_hash"`
	PocEmailHash      string                     `json:"poc_email_hash"`
	Aliases           []string                   `json:"aliases"`
	Tags              []string                   `json:"tags"`
	RulesetName       string                     `json:"ruleset_name"`
	AnalysisSummaries []analysis.AnalysisSummary `json:"analysis_summaries"`
}

// ProjectReports is used for getting a high level overview, returning a single analysis
type ProjectReports struct {
	ID              string                   `json:"id"`
	TeamID          string                   `json:"team_id"`
	RulesetID       string                   `json:"ruleset_id"`
	Name            string                   `json:"name"`
	Type            string                   `json:"type"`
	Source          string                   `json:"source"`
	Branch          string                   `json:"branch"`
	Description     string                   `json:"description"`
	Active          bool                     `json:"active"`
	ChatChannel     string                   `json:"chat_channel"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	DeployKey       string                   `json:"deploy_key"`
	Monitor         bool                     `json:"should_monitor"`
	PocName         string                   `json:"poc_name"`
	PocEmail        string                   `json:"poc_email"`
	Username        string                   `json:"username"`
	Password        string                   `json:"password"`
	KeyFingerprint  string                   `json:"key_fingerprint"`
	PocNameHash     string                   `json:"poc_name_hash"`
	PocEmailHash    string                   `json:"poc_email_hash"`
	Aliases         []string                 `json:"aliases"`
	Tags            []string                 `json:"tags"`
	RulesetName     string                   `json:"ruleset_name"`
	AnalysisSummary analysis.AnalysisSummary `json:"analysis_summary"`
}
