package reports

import (
	"encoding/json"
	"time"
)

// Report is a Ion Channel representation of a report output from a given analysis
type Report struct {
	ID            string            `json:"id" xml:"id"`
	TeamID        string            `json:"team_id" xml:"team_id"`
	ProjectID     string            `json:"project_id" xml:"project_id"`
	BuildNumber   string            `json:"build_number" xml:"build_number"`
	Name          string            `json:"name" xml:"name"`
	Text          string            `json:"text" xml:"text"`
	Type          string            `json:"type" xml:"type"`
	Source        string            `json:"source" xml:"source"`
	Branch        string            `json:"branch" xml:"branch"`
	Description   string            `json:"description" xml:"description"`
	Status        string            `json:"status" xml:"status"`
	RulesetID     string            `json:"ruleset_id" xml:"ruleset_id"`
	RulesetName   string            `json:"ruleset_name" xml:"ruleset_name"`
	Passed        bool              `json:"passed" xml:"passed"`
	CreatedAt     time.Time         `json:"created_at" xml:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at" xml:"updated_at"`
	Duration      float64           `json:"duration" xml:"duration"`
	Trigger       string            `json:"trigger" xml:"trigger"`
	TriggerHash   string            `json:"trigger_hash" xml:"trigger_hash"`
	TriggerText   string            `json:"trigger_text" xml:"trigger_text"`
	TriggerAuthor string            `json:"trigger_author" xml:"trigger_author"`
	Risk          string            `json:"risk" xml:"risk"`
	Summary       string            `json:"summary" xml:"summary"`
	ScanSummaries []json.RawMessage `json:"scan_summaries" xml:"scan_summaries"`
}
