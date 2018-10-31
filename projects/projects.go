package projects

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/tags"
)

const (
	validEmailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
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

// Validate takes an http client and returns a slice of fields as a string and
// an error. The fields will be a list of fields that did not pass the
// validation. An error will only be returned if any of the fields fail their
// validation.
func (p *Project) Validate(client *http.Client) (map[string]string, error) {
	invalidFields := make(map[string]string)
	var projErr error

	if p.ID == nil {
		invalidFields["id"] = "missing id"
		projErr = ErrInvalidProject
	}

	if p.TeamID == nil {
		invalidFields["team_id"] = "missing team id"
		projErr = ErrInvalidProject
	}

	if p.RulesetID == nil {
		invalidFields["ruleset_id"] = "missing ruleset id"
		projErr = ErrInvalidProject
	}

	if p.Name == nil {
		invalidFields["name"] = "missing name"
		projErr = ErrInvalidProject
	}

	if p.Type == nil {
		invalidFields["type"] = "missing type"
		projErr = ErrInvalidProject
	}

	if p.Source == nil {
		invalidFields["source"] = "missing source"
		projErr = ErrInvalidProject
	}

	if p.Branch == nil {
		invalidFields["branch"] = "missing branch"
		projErr = ErrInvalidProject
	}

	if p.Description == nil {
		invalidFields["description"] = "missing description"
		projErr = ErrInvalidProject
	}

	r := regexp.MustCompile(validEmailRegex)
	if p.POCEmail != "" && !r.MatchString(p.POCEmail) {
		invalidFields["poc_email"] = "invalid email supplied"
		projErr = ErrInvalidProject
	}

	block, rest := pem.Decode([]byte(p.DeployKey))
	if block != nil {
		pkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			invalidFields["deploy_key"] = "must be a valid ssh key"
			projErr = ErrInvalidProject
		} else {
			err = pkey.Validate()
			if err != nil {
				invalidFields["deploy_key"] = "must be a valid ssh key"
				projErr = ErrInvalidProject
			}
		}
	}

	if block == nil && rest != nil && string(rest) != "" {
		invalidFields["deploy_key"] = "must be a valid ssh key"
		projErr = ErrInvalidProject
	}

	u, err := url.Parse(*p.Source)
	if err != nil {
		invalidFields["source"] = fmt.Sprintf("source must be a valid url: %v", err.Error())
	}

	if u != nil {
		res, err := client.Head(u.String())
		if err != nil {
			invalidFields["source"] = "source failed to return a response"
		}

		if res != nil && res.StatusCode == http.StatusNotFound {
			invalidFields["source"] = "source returned a not found"
		}
	}

	return invalidFields, projErr
}
