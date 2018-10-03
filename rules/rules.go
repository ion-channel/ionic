package rules

import (
	"time"
)

//Rule identifies an Ion system predicate a project is held against
type Rule struct {
	ID             string    `json:"id"`
	ScanType       string    `json:"scan_type"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	PolicyURL      string    `json:"policy_url"`
	RemediationURL string    `json:"remediation_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
