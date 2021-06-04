package analyses

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scans"
)

const (
	// AnalysisGetAnalysisEndpoint returns a single raw analysis. Requires team id, project id and analysis id.
	AnalysisGetAnalysisEndpoint = "v1/animal/getAnalysis"
	// AnalysisGetAnalysesEndpoint returns multiple raw analyses. Requires team id and project id.
	AnalysisGetAnalysesEndpoint = "v1/animal/getAnalyses"
	// AnalysisGetLatestAnalysisIDsEndpoint returns the latest analysis summary. Requires team id, project ID(s) optional
	AnalysisGetLatestAnalysisIDsEndpoint = "v1/animal/getLatestAnalysisIDs"
	// AnalysisGetLatestAnalysisSummaryEndpoint returns the latest analysis summary. Requires team id and project id.
	AnalysisGetLatestAnalysisSummaryEndpoint = "v1/animal/getLatestAnalysisSummary"
	// AnalysisGetLatestAnalysisSummariesEndpoint returns the latest analysis summaries for multiple projects. Requires team id and project id.
	AnalysisGetLatestAnalysisSummariesEndpoint = "v1/animal/getLatestAnalysisSummaries"
	// AnalysisGetPublicAnalysisEndpoint returns a public analysis.  Requires an analysis id.
	AnalysisGetPublicAnalysisEndpoint = "v1/animal/getPublicAnalysis"
	// AnalysisGetLatestPublicAnalysisEndpoint returns a public analysis.  Requires an analysis id.
	AnalysisGetLatestPublicAnalysisEndpoint = "v1/animal/getLatestPublicAnalysisSummary"
	// AnalysisGetLatestAnalysisEndpoint returns the latest analysis. Requires team id and project id.
	AnalysisGetLatestAnalysisEndpoint = "v1/animal/getLatestAnalysis"
	// AnalysisGetScanEndpoint returns a scan. Requires a team id, project id, analysis id and scan id.
	AnalysisGetScanEndpoint = "v1/animal/getAnalysisExportData"
	// AnalysisGetAnalysesExportData returns exported data for a list of analyses. Requires a team id and list of analyses ids.
	AnalysisGetAnalysesExportData = "v1/animal/getAnalysesExportData"
	// AnalysisGetAnalysesVulnerabilityExportData returns exported vulnerability data for a list of analyses.
	// Requires a team id and list of analyses ids.
	AnalysisGetAnalysesVulnerabilityExportData = "v1/animal/getAnalysesVulnerabilityExportData"
)

// Analysis is a representation of an Ion Channel Analysis within the system
type Analysis struct {
	ID            string       `json:"id" xml:"id"`
	TeamID        string       `json:"team_id" xml:"team_id"`
	ProjectID     string       `json:"project_id" xml:"project_id"`
	Name          string       `json:"name" xml:"name"`
	Text          *string      `json:"text" xml:"text"`
	Type          string       `json:"type" xml:"type"`
	Source        string       `json:"source" xml:"source"`
	Branch        string       `json:"branch" xml:"branch"`
	Description   string       `json:"description" xml:"description"`
	Status        string       `json:"status" xml:"status"`
	RulesetID     string       `json:"ruleset_id" xml:"ruleset_id"`
	CreatedAt     time.Time    `json:"created_at" xml:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" xml:"updated_at"`
	Duration      float64      `json:"duration" xml:"duration"`
	TriggerHash   string       `json:"trigger_hash" xml:"trigger_hash"`
	TriggerText   string       `json:"trigger_text" xml:"trigger_text"`
	TriggerAuthor string       `json:"trigger_author" xml:"trigger_author"`
	ScanSummaries []scans.Scan `json:"scan_summaries" xml:"scan_summaries"`
	Public        bool         `json:"public" xml:"public"`
}

// Summary is a representation of a summarized Ion Channel Analysis
// within the system
type Summary struct {
	ID            string    `json:"id"`
	AnalysisID    string    `json:"analysis_id"`
	TeamID        string    `json:"team_id"`
	Branch        string    `json:"branch"`
	Description   string    `json:"description"`
	Risk          string    `json:"risk"`
	Summary       string    `json:"summary"`
	Passed        bool      `json:"passed"`
	RulesetID     string    `json:"ruleset_id"`
	RulesetName   string    `json:"ruleset_name"`
	Duration      float64   `json:"duration"`
	CreatedAt     time.Time `json:"created_at"`
	TriggerHash   string    `json:"trigger_hash"`
	TriggerText   string    `json:"trigger_text"`
	TriggerAuthor string    `json:"trigger_author"`
	Trigger       string    `json:"trigger"`
}

// ExportData is the data representation of a scan's exported data
type ExportData struct {
	AnalysisID          string `json:"analysis_id"`
	ProjectID           string `json:"project_id"`
	CPE                 string `json:"cpe"`
	Status              string `json:"status"`
	Source              string `json:"source"`
	CommitterCount      int    `json:"committer_count"`
	DaysSinceLastCommit int    `json:"days_since_last_commit"`
	VirusCount          int    `json:"virus_count"`
	VulnerabilityCount  int    `json:"vulnerability_count"`
	HighVulnCount       int    `json:"high_vulnerability_count"`
	CritVulnCount       int    `json:"critical_vulnerability_count"`
}

// VulnerabilityScore is just a normal float64, but it is rounded down to one decimal place when marshaled to JSON
type VulnerabilityScore float64

// MarshalJSON rounds the VulnerabilityScore down to one decimal place to avoid weird values like 7.80000002
func (s VulnerabilityScore) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.1f", s)), nil
}

// VulnerabilityExportData summarizes key details about a vulnerability to be included in a project's export data.
type VulnerabilityExportData struct {
	ProjectID         string             `json:"project_id"`
	ProjectName       string             `json:"project_name"`
	AnalysisID        string             `json:"analysis_id"`
	Title             string             `json:"title"`
	ExternalID        string             `json:"external_id"`
	Severity          string             `json:"severity"`
	Score             VulnerabilityScore `json:"score"`
	Dependency        string             `json:"dependency"`
	DependencyVersion string             `json:"dependency_version"`
}

// String returns a JSON formatted string of the analysis object
func (a Analysis) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("failed to format user: %v", err.Error())
	}
	return string(b)
}

// NewSummary takes an Analysis and AppliedRulesetSummary to calculate and
// return a Summary of the Analysis
func NewSummary(a *Analysis, appliedRuleset *rulesets.AppliedRulesetSummary) *Summary {
	if a != nil {
		rulesetName := "N/A"
		risk := "high"
		passed := false

		if appliedRuleset != nil {
			risk, passed = appliedRuleset.SummarizeEvaluation()

			if appliedRuleset.RuleEvaluationSummary != nil && appliedRuleset.RuleEvaluationSummary.RulesetName != "" {
				rulesetName = appliedRuleset.RuleEvaluationSummary.RulesetName
			}
		}

		return &Summary{
			ID:            a.ID,
			AnalysisID:    a.ID,
			TeamID:        a.TeamID,
			Branch:        a.Branch,
			Description:   a.Description,
			Risk:          risk,
			Summary:       "",
			Passed:        passed,
			RulesetID:     a.RulesetID,
			RulesetName:   rulesetName,
			Duration:      a.Duration,
			CreatedAt:     a.CreatedAt,
			TriggerHash:   a.TriggerHash,
			TriggerText:   a.TriggerText,
			TriggerAuthor: a.TriggerAuthor,
			Trigger:       "source commit",
		}
	}

	return &Summary{}
}
