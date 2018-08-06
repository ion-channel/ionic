package projects

import (
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/tags"
)

//Project is a representation of a project within the Ion Channel system
type Project struct {
	ID             *string         `json:"id,omitempty"`
	TeamID         *string         `json:"team_id,omitempty"`
	RulesetID      *string         `json:"ruleset_id,omitempty"`
	Name           *string         `json:"name,omitempty"`
	Type           *string         `json:"type,omitempty"`
	Source         *string         `json:"source,omitempty"`
	Branch         *string         `json:"branch,omitempty"`
	Description    *string         `json:"description,omitempty"`
	Active         bool            `json:"active"`
	ChatChannel    *string         `json:"chat_channel,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeployKey      *string         `json:"deploy_key,omitempty"`
	Monitor        bool            `json:"should_monitor"`
	POCName        *string         `json:"poc_name,omitempty"`
	POCEmail       *string         `json:"poc_email,omitempty"`
	Username       *string         `json:"username,omitempty"`
	Password       *string         `json:"password,omitempty"`
	KeyFingerprint *string         `json:"key_fingerprint,omitempty"`
	Aliases        []aliases.Alias `json:"aliases"`
	Tags           []tags.Tag      `json:"tags"`
}
