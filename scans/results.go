package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ion-channel/ionic/vulnerabilities"
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
			return fmt.Errorf("failed to unmarshall about yml results: %v", err)
		}

		r.Data = a
	case "community":
		var c []CommunityResults
		err := json.Unmarshal(tr.RawData, &c)
		if err != nil {
			return fmt.Errorf("failed to unmarshall community results: %v", err)
		}

		r.Data = c
	case "coverage", "external_coverage":
		var c CoverageResults
		err := json.Unmarshal(tr.RawData, &c)
		if err != nil {
			return fmt.Errorf("failed to unmarshall coverage results: %v", err)
		}

		r.Data = c
	case "dependency":
		var d DependencyResults
		err := json.Unmarshal(tr.RawData, &d)
		if err != nil {
			return fmt.Errorf("failed to unmarshall dependency results: %v", err)
		}

		r.Data = d
	case "ecosystems":
		var e EcosystemResults
		err := json.Unmarshal(tr.RawData, &e)
		if err != nil {
			return fmt.Errorf("failed to unmarshall ecosystems results: %v", err)
		}

		r.Data = e
	case "license":
		var l LicenseResults
		err := json.Unmarshal(tr.RawData, &l)
		if err != nil {
			return fmt.Errorf("failed to unmarshall license results: %v", err)
		}

		r.Data = l
	case "virus", "clamav":
		var v VirusResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall virus results: %v", err)
		}

		r.Data = v
	case "vulnerability":
		var v VulnerabilityResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall vulnerability results: %v", err)
		}

		r.Data = v
	case "external_vulnerability":
		var v ExternalVulnerabilitiesResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall external vulnerabilities results: %v", err)
		}

		r.Data = v
	case "difference":
		var v DifferenceResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall difference results: %v", err)
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

// CommunityResults represents the data collected from a community scan.  It
// represents all known data regarding the open community of a software project
type CommunityResults struct {
	Committers int    `json:"committers" xml:"committers"`
	Name       string `json:"name" xml:"name"`
	URL        string `json:"url" xml:"url"`
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

// DifferenceResults represents the checksum of a project.  It includes a checksum
// and flag indicating if there was a difference detected within that last 5 scans
type DifferenceResults struct {
	Checksum   string `json:"checksum" xml:"checksum"`
	Difference bool   `json:"difference" xml:"difference"`
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

// ExternalVulnerabilitiesResults represents the data colleced from an external
// vulnerability scan.  It includes the number of each vulnerability criticality
// seen within the project.
type ExternalVulnerabilitiesResults struct {
	Critical int `json:"critical" xml:"critical"`
	High     int `json:"high" xml:"high"`
	Medium   int `json:"medium" xml:"medium"`
	Low      int `json:"low" xml:"low"`
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
	FileNotes          struct {
		EmptyFiles []string `json:"empty_file" xml:"empty_file"`
	} `json:"file_notes" xml:"file_notes"`
	ClamavDetails struct {
		ClamavVersion   string `json:"clamav_version" xml:"clamav_version"`
		ClamavDbVersion string `json:"clamav_db_version" xml:"clamav_db_version"`
	} `json:"clam_av_details" xml:"clam_av_details"`
}

type VulnerabilityResults struct {
	Vulnerabilities []struct {
		ID              int                             `json:"id" xml:"id"`
		ExternalID      string                          `json:"external_id" xml:"external_id"`
		SourceID        int                             `json:"source_id" xml:"source_id"`
		Title           string                          `json:"title" xml:"title"`
		Name            string                          `json:"name" xml:"name"`
		Org             string                          `json:"org" xml:"org"`
		Version         string                          `json:"version" xml:"version"`
		Up              interface{}                     `json:"up" xml:"up"`
		Edition         interface{}                     `json:edition"" xml:"edition"`
		Aliases         []string                        `json:"aliases" xml:"aliases"`
		CreatedAt       time.Time                       `json:"created_at" xml:"created_at"`
		UpdatedAt       time.Time                       `json:"updated_at" xml:"updated_at"`
		References      interface{}                     `json:"references" xml:"references"`
		Part            interface{}                     `json:"part" xml:"part"`
		Language        interface{}                     `json:"language" xml:"language"`
		Vulnerabilities []vulnerabilities.Vulnerability `json:"vulnerabilities" xml:"vulnerabilities"`
	} `json:"vulnerabilities" xml:"vulnerabilities"`
	Meta struct {
		VulnerabilityCount int `json:"vulnerability_count" xml:"vulnerability_count"`
	} `json:"meta" xml:"meta"`
}
