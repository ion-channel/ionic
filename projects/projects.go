package projects

import (
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/tags"
)

//Project is a representation of a project within the Ion Channel system
type Project struct {
	ID             *string          `json:"id"`
	TeamID         *string          `json:"team_id"`
	RulesetID      *string          `json:"ruleset_id"`
	Name           *string          `json:"name"`
	Type           *string          `json:"type"`
	Source         *string          `json:"source"`
	Branch         *string          `json:"branch"`
	Description    *string          `json:"description"`
	Active         bool            `json:"active"`
	ChatChannel    *string          `json:"chat_channel"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeployKey      *string          `json:"deploy_key"`
	Monitor        bool            `json:"should_monitor"`
	POCName        *string          `json:"poc_name"`
	POCEmail       *string          `json:"poc_email"`
	Username       *string          `json:"username"`
	Password       *string          `json:"password"`
	KeyFingerprint *string          `json:"key_fingerprint"`
	Aliases        []aliases.Alias `json:"aliases"`
	Tags           []tags.Tag      `json:"tags"`
}
