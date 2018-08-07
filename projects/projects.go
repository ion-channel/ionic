package projects

import (
	"fmt"
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/tags"
)

var (
	// ErrInvalidProject is returned when a given project does not pass the
	// standards for a project
	ErrInvalidProject = fmt.Errorf("project has invalid fields")
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
	ChatChannel    string          `json:"chat_channel"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeployKey      string          `json:"deploy_key"`
	Monitor        bool            `json:"should_monitor"`
	POCName        string          `json:"poc_name"`
	POCEmail       string          `json:"poc_email"`
	Username       string          `json:"username"`
	Password       string          `json:"password"`
	KeyFingerprint string          `json:"key_fingerprint"`
	Aliases        []aliases.Alias `json:"aliases"`
	Tags           []tags.Tag      `json:"tags"`
}

// Validate returns a slice of fields as a string and an error. The fields will
// be a list of fields that did not pass the validation. An error will only be
// returned if any of the fields fail their validation.
func (p *Project) Validate() ([]string, error) {
	var fields []string
	var err error

	if p.ID == nil {
		fields = append(fields, "id")
		err = ErrInvalidProject
	}

	if p.TeamID == nil {
		fields = append(fields, "team_id")
		err = ErrInvalidProject
	}

	if p.RulesetID == nil {
		fields = append(fields, "ruleset_id")
		err = ErrInvalidProject
	}

	if p.Name == nil {
		fields = append(fields, "name")
		err = ErrInvalidProject
	}

	if p.Type == nil {
		fields = append(fields, "type")
		err = ErrInvalidProject
	}

	if p.Source == nil {
		fields = append(fields, "source")
		err = ErrInvalidProject
	}

	if p.Branch == nil {
		fields = append(fields, "branch")
		err = ErrInvalidProject
	}

	if p.Description == nil {
		fields = append(fields, "description")
		err = ErrInvalidProject
	}

	return fields, err
}
