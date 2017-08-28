package reports

import (
	"encoding/json"
	"time"
)

type Report struct {
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
	RulesetName   string            `json:"ruleset_name"`
	Passed        bool              `json:"passed"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Duration      float64           `json:"duration"`
	Trigger       string            `json:"trigger"`
	TriggerHash   string            `json:"trigger_hash"`
	TriggerText   string            `json:"trigger_text"`
	TriggerAuthor string            `json:"trigger_author"`
	Risk          string            `json:"risk"`
	Summary       string            `json:"summary"`
	ScanSummaries []json.RawMessage `json:"scan_summaries"`
}
