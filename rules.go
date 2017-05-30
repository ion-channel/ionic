package ionic

import (
	"net/url"
	"time"
)

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
