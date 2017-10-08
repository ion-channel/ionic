package scans

import (
	"encoding/json"
	"fmt"
	"strings"
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

type AboutYMLResults struct {
	Message string `json:"message" xml:"message"`
	Valid   bool   `json:"valid" xml:"valid"`
	Content string `json:"content" xml:"content"`
}

type CoverageResults struct {
	Value float64 `json:"value"`
}

type DependencyResults struct {
	Dependencies []struct {
		LatestVersion string `json:"latest_version"`
		Org           string `json:"org"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		Package       string `json:"package"`
		Version       string `json:"version"`
		Scope         string `json:"scope"`
	} `json:"dependencies"`
	Meta struct {
		FirstDegreeCount     int `json:"first_degree_count"`
		NoVersionCount       int `json:"no_version_count"`
		TotalUniqueCount     int `json:"total_unique_count"`
		UpdateAvailableCount int `json:"update_available_count"`
	} `json:"meta"`
}

type EcosystemResults struct {
	Ecosystems []struct {
		Ecosystem string `json:"ecosystem" xml:"ecosystem"`
		Lines     int    `json:"lines" xml:"lines"`
	} `json:"ecosystems" xml:"ecosystems"`
}

type LicenseResults struct {
	License struct {
		Name string `json:"name"`
		Type []struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"license"`
}

type VirusResults struct {
	KnownViruses       int      `json:"known_viruses" xml:"known_viruses"`
	EngineVersion      string   `json:"engine_version" xml:"engine_version"`
	ScannedDirectories int      `json:"scanned_directories" xml:"scanned_directories"`
	ScannedFiles       int      `json:"scanned_files" xml:"scanned_files"`
	InfectedFiles      int      `json:"infected_files" xml:"infected_files"`
	DataScanned        string   `json:"data_scanned" xml:"data_scanned"`
	DataRead           string   `json:"data_read" xml:"data_read"`
	Time               string   `json:"time" xml:"time"`
	FileNotes          struct{} `json:"file_notes" xml:"file_notes"`
	ClamavDetails      struct {
		ClamavVersion   string `json:"clamav_version" xml:"clamav_version"`
		ClamavDbVersion string `json:"clamav_db_version" xml:"clamav_db_version"`
	} `json:"clam_av_details" xml:"clam_av_details"`
}

type VulnerabilityResults struct {
	Vulnerabilities []interface{} `json:"vulnerabilities"`
	Meta            struct {
		VulnerabilityCount int `json:"vulnerability_count"`
	} `json:"meta"`
}
