package analyses

import (
	"time"

	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/scans"
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

// Navigation represents a navigational meta data reference to given analysis
type Navigation struct {
	Analysis       *Analysis `json:"analysis"`
	LatestAnalysis *Analysis `json:"latest_analysis"`
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
