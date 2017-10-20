package vulnerabilities

import (
	"encoding/json"
	"time"

	"github.com/ion-channel/ionic/products"
)

// Vulnerability represents a singular vulnerability record in the Ion Channel
// Platform
type Vulnerability struct {
	ID           int    `json:"id" xml:"id"`
	ExternalID   string `json:"external_id" xml:"exteral_id"`
	SourceID     int    `json:"source_id" xml:"source_id"`
	Title        string `json:"title" xml:"title"`
	Summary      string `json:"summary" xml:"summary"`
	Score        string `json:"score" xml:"score"`
	ScoreVersion string `json:"score_version" xml:"score_version"`
	ScoreSystem  string `json:"score_system" xml:"score_system"`
	ScoreDetails struct {
		CVSSv2 CVSSv2 `json:"cvssv2" xml:"cvssv2"`
		CVSSv3 CVSSv3 `json:"cvssv3" xml:"cvssv3"`
	} `json:"score_details" xml:"score_details"`
	Vector                      string             `json:"vector" xml:"vector"`
	AccessComplexity            string             `json:"access_complexity" xml:"access_complexity"`
	VulnerabilityAuthentication string             `json:"vulnerability_authentication" xml:"vulnerability_authentication"`
	ConfidentialityImpact       string             `json:"confidentiality_impact" xml:"confidentiality_impact"`
	IntegrityImpact             string             `json:"integrity_impact" xml:"integrity_impact"`
	AvailabilityImpact          string             `json:"availability_impact" xml:"availability_impact"`
	VulnerabilitySource         string             `json:"vulnerabilty_source" xml:"vulnerability_source"`
	AssessmentCheck             json.RawMessage    `json:"assessment_check" xml:"assessment_check"`
	Scanner                     json.RawMessage    `json:"scanner" xml:"scanner"`
	Recommendation              string             `json:"recommendation" xml:"recommendation"`
	Dependencies                []products.Product `json:"dependencies" xml:"dependencies"`
	References                  []struct {
		Type   string `json:"type" xml:"type"`
		Source string `json:"source" xml:"source"`
		URL    string `json:"url" xml:"url"`
		Text   string `json:"text" xml:"text"`
	} `json:"references" xml:"references"`
	ModifiedAt  time.Time `json:"modified_at" xml:"modified_at"`
	PublishedAt time.Time `json:"published_at" xml:"published_at"`
	CreatedAt   time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" xml:"updated_at"`
}

// CVSSv2 represents the variables that go into determining the CVSS v2 score
// for a given vulnerability
type CVSSv2 struct {
	VectorString          string  `json:"vectorString" xml:"vectorString"`
	AccessVector          string  `json:"accessVector" xml:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity" xml:"accessComplexity"`
	Authentication        string  `json:"authentication" xml:"authentication"`
	ConfidentialityImpact string  `json:"confidentialityImpact" xml:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact" xml:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact" xml:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore" xml:"baseScore"`
}

// CVSSv3 represents the variables that go into determining the CVSS v3 score
// for a given vulnerability
type CVSSv3 struct {
	VectorString          string  `json:"vectorString" xml:"vectorString"`
	AccessVector          string  `json:"accessVector" xml:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity" xml:"accessComplexity"`
	PrivilegesRequired    string  `json:"privilegesRequired" xml:"privilegesRequired"`
	UserInteraction       string  `json:"userInteraction" xml:"userInteraction"`
	Scope                 string  `json:"scope" xml:"scope"`
	ConfidentialityImpact string  `json:"confidentialityImpact" xml:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact" xml:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact" xml:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore" xml:"baseScore"`
	BaseSeverity          string  `json:"baseSeverity" xml:"baseSeverity"`
}
