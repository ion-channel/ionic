package vulnerabilities

import (
	"encoding/json"
	"time"
)

// Vulnerability represents a singular vulnerability record in the Ion Channel
// Platform
type Vulnerability struct {
	ID           int    `json:"id"`
	ExternalID   string `json:"external_id"`
	SourceID     int    `json:"source_id"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Score        string `json:"score"`
	ScoreVersion string `json:"score_version"`
	ScoreSystem  string `json:"score_system"`
	ScoreDetails struct {
		CVSSv2 CVSSv2 `json:"cvssv2"`
		CVSSv3 CVSSv3 `json:"cvssv3"`
	} `json:"score_details"`
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
	References                  []struct {
		Type   string `json:"type"`
		Source string `json:"source"`
		URL    string `json:"url"`
		Text   string `json:"text"`
	} `json:"references"`
	ModifiedAt  time.Time `json:"modified_at"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CVSSv2 represents the variables that go into determining the CVSS v2 score
// for a given vulnerability
type CVSSv2 struct {
	VectorString          string  `json:"vectorString"`
	AccessVector          string  `json:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity"`
	Authentication        string  `json:"authentication"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
}

// CVSSv3 represents the variables that go into determining the CVSS v3 score
// for a given vulnerability
type CVSSv3 struct {
	VectorString          string  `json:"vectorString"`
	AttackVector          string  `json:"attackVector"`
	AttackComplexity      string  `json:"attackComplexity"`
	PrivilegesRequired    string  `json:"privilegesRequired"`
	UserInteraction       string  `json:"userInteraction"`
	Scope                 string  `json:"scope"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
	BaseSeverity          string  `json:"baseSeverity"`
}
