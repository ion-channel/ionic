package analysis

import (
	"time"

	"github.com/ion-channel/ionic/scans"
)

// Analysis is a representation of an Ion Channel Analysis within the system
type Analysis struct {
	ID            string          `json:"id" xml:"id"`
	TeamID        string          `json:"team_id" xml:"team_id"`
	ProjectID     string          `json:"project_id" xml:"project_id"`
	BuildNumber   string          `json:"build_number" xml:"build_number"`
	Name          string          `json:"name" xml:"name"`
	Text          *string         `json:"text" xml:"text"`
	Type          string          `json:"type" xml:"type"`
	Source        string          `json:"source" xml:"source"`
	Branch        string          `json:"branch" xml:"branch"`
	Description   string          `json:"description" xml:"description"`
	Status        string          `json:"status" xml:"status"`
	RulesetID     string          `json:"ruleset_id" xml:"ruleset_id"`
	CreatedAt     time.Time       `json:"created_at" xml:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" xml:"updated_at"`
	Duration      float64         `json:"duration" xml:"duration"`
	TriggerHash   string          `json:"trigger_hash" xml:"trigger_hash"`
	TriggerText   string          `json:"trigger_text" xml:"trigger_text"`
	TriggerAuthor string          `json:"trigger_author" xml:"trigger_author"`
	ScanSummaries []scans.Summary `json:"scan_summaries" xml:"scan_summaries"`
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
