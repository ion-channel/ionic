package analysis

import (
	"encoding/json"
	"time"
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

// AnalysisSummary is a representation of a summarized Ion Channel Analysis
// within the system
type AnalysisSummary struct {
	ID            string    `json:"id"`
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
