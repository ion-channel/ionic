package rules

import (
	"net/url"
	"time"
)

//Rule identifies an Ion system predicate a project is held against
type Rule struct {
	ID             string
	ScanType       string
	Name           string
	Description    string
	Category       string
	PolicyURL      *url.URL
	RemediationURL *url.URL
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
