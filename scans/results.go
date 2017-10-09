package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Results is an Ion Channel representation of the results from a
// scan summary.  It contains what type of results and the data pertaining to
// the results.
type Results struct {
	Type string      `json:"type" xml:"type"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

type results struct {
	Type    string          `json:"type"`
	RawData json.RawMessage `json:"data"`
}

// UnmarshalJSON is a custom JSON unmarshaller implementation for the standard
// go json package to know how to properly interpret ScanSummaryResults from
// JSON.
func (r *Results) UnmarshalJSON(b []byte) error {
	var tr results
	err := json.Unmarshal(b, &tr)
	if err != nil {
		return err
	}

	r.Type = tr.Type

	switch strings.ToLower(tr.Type) {
	case "about_yml":
		var a AboutYMLResults
		err := json.Unmarshal(tr.RawData, &a)
		if err != nil {
			return err
		}

		r.Data = a
	case "coverage", "external_coverage":
		var c CoverageResults
		err := json.Unmarshal(tr.RawData, &c)
		if err != nil {
			return err
		}

		r.Data = c
	case "dependency":
		var d DependencyResults
		err := json.Unmarshal(tr.RawData, &d)
		if err != nil {
			return err
		}

		r.Data = d
	case "ecosystems":
		var e EcosystemResults
		err := json.Unmarshal(tr.RawData, &e)
		if err != nil {
			return err
		}

		r.Data = e
	case "license":
		var l LicenseResults
		err := json.Unmarshal(tr.RawData, &l)
		if err != nil {
			return err
		}

		r.Data = l
	case "virus", "clamav":
		var v VirusResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return err
		}

		r.Data = v
	case "vulnerability":
		var v VulnerabilityResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return err
		}

		r.Data = v
	default:
		return fmt.Errorf("invalid results type")
	}

	return nil
}

// AboutYMLResults represents the data collected from the AboutYML scan.  It
// includes a message and whether or not the About YML file found was valid or
// not.
type AboutYMLResults struct {
	Message string `json:"message" xml:"message"`
	Valid   bool   `json:"valid" xml:"valid"`
	Content string `json:"content" xml:"content"`
}

// CoverageResults represents the data collected from a code coverage scan.  It
// includes the value of the code coverage seen for the project.
type CoverageResults struct {
	Value float64 `json:"value" xml:"value"`
}

// DependencyResults represents the data collected from a dependency scan.  It
// includes a list of the dependencies seen and meta data counts about those
// dependencies seen.
type DependencyResults struct {
	Dependencies []struct {
		LatestVersion string `json:"latest_version" xml:"latest_version"`
		Org           string `json:"org" xml:"org"`
		Name          string `json:"name" xml:"name"`
		Type          string `json:"type" xml:"type"`
		Package       string `json:"package" xml:"package"`
		Version       string `json:"version" xml:"version"`
		Scope         string `json:"scope" xml:"scope"`
	} `json:"dependencies" xml:"dependencies"`
	Meta struct {
		FirstDegreeCount     int `json:"first_degree_count" xml:"first_degree_count"`
		NoVersionCount       int `json:"no_version_count" xml:"no_version_count"`
		TotalUniqueCount     int `json:"total_unique_count" xml:"total_unique_count"`
		UpdateAvailableCount int `json:"update_available_count" xml:"update_available_count"`
	} `json:"meta" xml:"meta"`
}

// EcosystemResults represents the data collected from an ecosystems scan.  It
// include the name of the ecosystem and the number of lines seen for the given
// ecosystem.
type EcosystemResults struct {
	Ecosystems []struct {
		Ecosystem string `json:"ecosystem" xml:"ecosystem"`
		Lines     int    `json:"lines" xml:"lines"`
	} `json:"ecosystems" xml:"ecosystems"`
}

// LicenseResults represents the data colleced from a license scan.  It
// includes the name and type of each license seen within the project.
type LicenseResults struct {
	License struct {
		Name string `json:"name" xml:"name"`
		Type []struct {
			Name string `json:"name" xml:"name"`
		} `json:"type" xml:"type"`
	} `json:"license" xml:"license"`
}

// VirusResults represents the data colleced from a virus scan.  It includes
// information of the viruses seen and the virus scanner used.
type VirusResults struct {
	KnownViruses       int    `json:"known_viruses" xml:"known_viruses"`
	EngineVersion      string `json:"engine_version" xml:"engine_version"`
	ScannedDirectories int    `json:"scanned_directories" xml:"scanned_directories"`
	ScannedFiles       int    `json:"scanned_files" xml:"scanned_files"`
	InfectedFiles      int    `json:"infected_files" xml:"infected_files"`
	DataScanned        string `json:"data_scanned" xml:"data_scanned"`
	DataRead           string `json:"data_read" xml:"data_read"`
	Time               string `json:"time" xml:"time"`
	FileNotes          string `json:"file_notes" xml:"file_notes"`
	ClamavDetails      struct {
		ClamavVersion   string `json:"clamav_version" xml:"clamav_version"`
		ClamavDbVersion string `json:"clamav_db_version" xml:"clamav_db_version"`
	} `json:"clam_av_details" xml:"clam_av_details"`
}

type VulnerabilityResults struct {
	Vulnerabilities []struct {
		ID              int         `json:"id"`
		Name            string      `json:"name"`
		Org             string      `json:"org"`
		Version         string      `json:"version"`
		Up              interface{} `json:"up"`
		Edition         interface{} `json:"edition"`
		Aliases         interface{} `json:"aliases"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		Title           interface{} `json:"title"`
		References      interface{} `json:"references"`
		Part            interface{} `json:"part"`
		Language        interface{} `json:"language"`
		SourceID        int         `json:"source_id"`
		ExternalID      string      `json:"external_id"`
		Vulnerabilities []struct {
			ID           int    `json:"id"`
			ExternalID   string `json:"external_id"`
			Title        string `json:"title"`
			Summary      string `json:"summary"`
			Score        string `json:"score"`
			ScoreVersion string `json:"score_version"`
			ScoreSystem  string `json:"score_system"`
			ScoreDetails struct {
				Cvssv2 struct {
					VectorString          string  `json:"vectorString"`
					AccessVector          string  `json:"accessVector"`
					AccessComplexity      string  `json:"accessComplexity"`
					Authentication        string  `json:"authentication"`
					ConfidentialityImpact string  `json:"confidentialityImpact"`
					IntegrityImpact       string  `json:"integrityImpact"`
					AvailabilityImpact    string  `json:"availabilityImpact"`
					BaseScore             float64 `json:"baseScore"`
				} `json:"cvssv2"`
				Cvssv3 struct {
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
				} `json:"cvssv3"`
			} `json:"score_details"`
			Vector                      string      `json:"vector"`
			AccessComplexity            string      `json:"access_complexity"`
			VulnerabilityAuthentication string      `json:"vulnerability_authentication"`
			ConfidentialityImpact       string      `json:"confidentiality_impact"`
			IntegrityImpact             string      `json:"integrity_impact"`
			AvailabilityImpact          string      `json:"availability_impact"`
			VulnerabilitySource         interface{} `json:"vulnerability_source"`
			AssessmentCheck             interface{} `json:"assessment_check"`
			Scanner                     interface{} `json:"scanner"`
			Recommendation              string      `json:"recommendation"`
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
			SourceID    int       `json:"source_id"`
		} `json:"vulnerabilities"`
	} `json:"vulnerabilities"`
	Meta struct {
		VulnerabilityCount int `json:"vulnerability_count"`
	} `json:"meta"`
}
