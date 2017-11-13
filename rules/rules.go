package rules

import (
	"net/url"
	"time"
)

//Rule needs a comment
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
