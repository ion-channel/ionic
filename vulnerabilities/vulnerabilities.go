package vulnerabilities

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ion-channel/ionic/products"
)

const (
	// GetVulnerabilitiesEndpoint is a string representation of the current endpoint for getting projects
	GetVulnerabilitiesEndpoint = "v1/vulnerability/getVulnerabilities"
	// GetVulnerabilitiesInFileEndpoint is a string representation of the current endpoint for getting projects
	GetVulnerabilitiesInFileEndpoint = "v1/vulnerability/getVulnerabilitiesInFile"
	// GetVulnerabilityEndpoint is a string representation of the current endpoint for getting projects
	GetVulnerabilityEndpoint = "v1/vulnerability/getVulnerability"
	// PostVulnerabilityEndpoint is a string representation of the current endpoint for getting projects
	PostVulnerabilityEndpoint = "v1/internal/vulnerability/addVulnerability"
)

// Vulnerability represents a singular vulnerability record in the Ion Channel
// Platform
type Vulnerability struct {
	ID                          int                `json:"id" xml:"id"`
	ExternalID                  string             `json:"external_id" xml:"exteral_id"`
	Source                      []Source           `json:"source" xml:"source"`
	Title                       string             `json:"title" xml:"title"`
	Summary                     string             `json:"summary" xml:"summary"`
	Score                       string             `json:"score,omitempty" xml:"score"`
	ScoreVersion                string             `json:"score_version,omitempty" xml:"score_version"`
	ScoreSystem                 string             `json:"score_system" xml:"score_system"`
	ScoreDetails                ScoreDetails       `json:"score_details" xml:"score_details"`
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
	References                  []Reference        `json:"references" xml:"references"`
	ModifiedAt                  time.Time          `json:"modified_at" xml:"modified_at"`
	PublishedAt                 time.Time          `json:"published_at" xml:"published_at"`
	CreatedAt                   time.Time          `json:"created_at" xml:"created_at"`
	UpdatedAt                   time.Time          `json:"updated_at" xml:"updated_at"`
}

// Source represents information about where the vulnerability data came from
type Source struct {
	ID           int       `json:"id" xml:"id"`
	Name         string    `json:"name" xml:"name"`
	Description  string    `json:"description" xml:"description"`
	CreatedAt    time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" xml:"updated_at"`
	Attribution  string    `json:"attribution" xml:"attribution"`
	License      string    `json:"license" xml:"license"`
	CopyrightURL string    `json:"copyright_url" xml:"copyright_url"`
}

// ScoreDetails represents the possible details for each version of scoring
type ScoreDetails struct {
	CVSSv2 *CVSSv2 `json:"cvssv2,omitempty" xml:"cvssv2"`
	CVSSv3 *CVSSv3 `json:"cvssv3,omitempty" xml:"cvssv3"`
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

// Reference represents a location where a CVE may have been referenced
type Reference struct {
	Type   string `json:"type" xml:"type"`
	Source string `json:"source" xml:"source"`
	URL    string `json:"url" xml:"url"`
	Text   string `json:"text" xml:"text"`
}

// NewV3FromShorthand takes a shorthand representation of a CVSSv3 and returns
// an expanded struct representation
func NewV3FromShorthand(shorthand string) *CVSSv3 {
	shorthand = strings.ToUpper(shorthand)
	shorthand = strings.TrimPrefix(shorthand, "CVSS:3.0/")
	metrics := strings.Split(shorthand, "/")

	sv := &CVSSv3{}

	for _, metric := range metrics {
		parts := strings.Split(metric, ":")
		switch parts[0] {
		case "AV":
			switch parts[1] {
			case "N":
				sv.AccessVector = "network"
			case "A":
				sv.AccessVector = "adjacent"
			case "L":
				sv.AccessVector = "local"
			case "P":
				sv.AccessVector = "physical"
			}
		case "AC":
			sv.AccessComplexity = parseLowHighNone(parts[1])
		case "PR":
			sv.PrivilegesRequired = parseLowHighNone(parts[1])
		case "UI":
			switch parts[1] {
			case "N":
				sv.UserInteraction = "none"
			case "R":
				sv.UserInteraction = "required"
			}
		case "S":
			switch parts[1] {
			case "U":
				sv.Scope = "unchanged"
			case "C":
				sv.Scope = "changed"
			}
		case "C":
			sv.ConfidentialityImpact = parseLowHighNone(parts[1])
		case "I":
			sv.IntegrityImpact = parseLowHighNone(parts[1])
		case "A":
			sv.AvailabilityImpact = parseLowHighNone(parts[1])
		}
	}

	return sv
}

func parseLowHighNone(lhn string) string {
	switch lhn {
	case "N":
		return "none"
	case "L":
		return "low"
	case "H":
		return "high"
	default:
		return lhn
	}
}
