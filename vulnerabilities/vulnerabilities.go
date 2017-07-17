package vulnerabilities

import (
	"encoding/json"
	"time"
)

// Vulnerability represents a singular vulnerability record in the Ion Channel
// Platform
type Vulnerability struct {
	ID                          int             `json:"id"`
	ExternalID                  string          `json:"external_id"`
	Title                       string          `json:"title"`
	Summary                     string          `json:"summary"`
	Score                       string          `json:"score"`
	Vector                      string          `json:"vector"`
	AccessComplexity            string          `json:"access_complexity"`
	VulnerabilityAuthentication string          `json:"vulnerability_authentication"`
	ConfidentialityImpact       string          `json:"confidentiality_impact"`
	IntegrityImpact             string          `json:"integrity_impact"`
	AvailabilityImpact          string          `json:"availability_impact"`
	VulnerabilitySource         string          `json:"vulnerabilty_source"`
	AssessmentCheck             json.RawMessage `json:"assessment_check"`
	Scanner                     json.RawMessage `json:"scanner"`
	Recommendation              string          `json:"recommendation"`
	ModifiedAt                  time.Time       `json:"modified_at"`
	PublishedAt                 time.Time       `json:"published_at"`
	CreatedAt                   time.Time       `json:"created_at"`
	UpdatedAt                   time.Time       `json:"updated_at"`
	SourceID                    int             `json:"source_id"`
}
