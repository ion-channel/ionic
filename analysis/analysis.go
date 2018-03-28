package analysis

import (
	"encoding/json"
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/tags"
)

// Analysis is a representation of an Ion Channel Analysis within the system
type Analysis struct {
	ID            string            `json:"id"`
	TeamID        string            `json:"team_id"`
	ProjectID     string            `json:"project_id"`
	BuildNumber   string            `json:"build_number"`
	Name          string            `json:"name"`
	Text          string            `json:"text"`
	Type          string            `json:"type"`
	Source        string            `json:"source"`
	Branch        string            `json:"branch"`
	Description   string            `json:"description"`
	Status        string            `json:"status"`
	RulesetID     string            `json:"ruleset_id"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Duration      float64           `json:"duration"`
	TriggerHash   string            `json:"trigger_hash"`
	TriggerText   string            `json:"trigger_text"`
	TriggerAuthor string            `json:"trigger_author"`
	ScanSummaries []json.RawMessage `json:"scan_summaries"`
}
type AugmentedAnalysis struct {
	Analysis
	// things that appear to bleed in from AnalysisSummary in the Statler ruby code
	RulesetName string          `json:"ruleset_name,omitempty"`
	Risk        string          `json:"risk,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Trigger     string          `json:"trigger,omitempty"`
	Passed      bool            `json:"passed,omitempty"`
	Aliases     []aliases.Alias `json:"aliases,omitempty"`
	Tags        []tags.Tag      `json:"tags,omitempty"`

}

func (a *AugmentedAnalysis) ValuesFrom(b *Analysis) {
	a.ID = b.ID
	a.TeamID = b.TeamID
	a.ProjectID = b.ProjectID
	a.BuildNumber = b.BuildNumber
	a.Name = b.Name
	a.Text = b.Text
	a.Type = b.Type
	a.Source = b.Source
	a.Branch = b.Branch
	a.Description = b.Description
	a.Status = b.Status
	a.RulesetID = b.RulesetID
	a.CreatedAt = b.CreatedAt
	a.UpdatedAt = b.UpdatedAt
	a.Duration = b.Duration
	a.TriggerHash = b.TriggerHash
	a.TriggerText = b.TriggerText
	a.TriggerAuthor = b.TriggerAuthor
	a.ScanSummaries = b.ScanSummaries
}

// Summary is a representation of a summarized Ion Channel Analysis
// within the system
type Summary struct {
	ID            string    `json:"id"`
	AnalysisID    string    `json:"analysis_id"`
	TeamID        string    `json:"team_id"`
	BuildNumber   string    `json:"build_number"`
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
