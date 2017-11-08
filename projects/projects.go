package projects

import (
	"time"
)

//Project is a representation of a project within the Ion Channel system
type Project struct {
	ID      string `json:"id"`
	Active  bool   `json:"active"`
	Aliases []struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Version   string    `json:"version"`
	} `json:"aliases"`
	Branch         string    `json:"branch"`
	ChatChannel    string    `json:"chat_channel"`
	CreatedAt      time.Time `json:"created_at"`
	DeployKey      string    `json:"deploy_key"`
	Description    string    `json:"description"`
	KeyFingerprint string    `json:"key_fingerprint"`
	Name           string    `json:"name"`
	RulesetID      string    `json:"ruleset_id"`
	Monitor        bool      `json:"should_monitor"`
	Source         string    `json:"source"`
	Tags           []struct {
		ID          string      `json:"id"`
		TeamID      string      `json:"team_id"`
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
	} `json:"tags"`
	TeamID    string    `json:"team_id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
