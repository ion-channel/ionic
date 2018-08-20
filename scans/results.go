package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ion-channel/ionic/vulnerabilities"
)

// UntranslatedResults represents a result of a specific type that has not been
// translated for use in reports
type UntranslatedResults struct {
	AboutYML                *AboutYMLResults                `json:"about_yml,omitempty"`
	Community               *CommunityResults               `json:"community,omitempty"`
	Coverage                *CoverageResults                `json:"coverage,omitempty"`
	Dependency              *DependencyResults              `json:"dependency,omitempty"`
	Difference              *DifferenceResults              `json:"difference,omitempty"`
	Ecosystem               *EcosystemResults               `json:"ecosystems,omitempty"`
	ExternalVulnerabilities *ExternalVulnerabilitiesResults `json:"external_vulnerability,omitempty"`
	License                 *LicenseResults                 `json:"license,omitempty"`
	Virus                   *VirusResults                   `json:"clamav,omitempty"`
	Vulnerability           *VulnerabilityResults           `json:"vulnerabilities,omitempty"`
}

// Translate moves information from the particular sub-struct, IE
// AboutYMLResults or LicenseResults into a generic, Data struct
func (u *UntranslatedResults) Translate() *TranslatedResults {
	var tr TranslatedResults
	// There is an argument to be made that the following "if" clauses
	// could be simplified with introspection since they all do
	// basically the same thing. I've (dmiles) chosen to writ it all
	// out in the name of explicit, easily-readable code.
	if u.AboutYML != nil {
		tr.Type = "about_yml"
		tr.Data = u.AboutYML
	}
	if u.Community != nil {
		tr.Type = "community"
		tr.Data = u.Community
	}
	if u.Coverage != nil {
		tr.Type = "coverage"
		tr.Data = u.Coverage
	}
	if u.Dependency != nil {
		tr.Type = "dependency"
		tr.Data = u.Dependency
	}
	if u.Difference != nil {
		tr.Type = "difference"
		tr.Data = u.Difference
	}
	if u.Ecosystem != nil {
		tr.Type = "ecosystems"
		tr.Data = u.Ecosystem
	}
	if u.ExternalVulnerabilities != nil {
		tr.Type = "external_vulnerability"
		tr.Data = u.ExternalVulnerabilities
	}
	if u.License != nil {
		tr.Type = "license"
		tr.Data = u.License
	}
	if u.Virus != nil {
		tr.Type = "virus"
		tr.Data = u.Virus
	}
	if u.Vulnerability != nil {
		tr.Type = "vulnerability"
		tr.Data = u.Vulnerability
	}
	return &tr
}

// TranslatedResults represents a result of a specific type that has been
// translated for use in reports
type TranslatedResults struct {
	Type string      `json:"type" xml:"type"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

type translatedResults struct {
	Type    string          `json:"type"`
	RawData json.RawMessage `json:"data"`
}

// UnmarshalJSON is a custom JSON unmarshaller implementation for the standard
// go json package to know how to properly interpret ScanSummaryResults from
// JSON.
func (r *TranslatedResults) UnmarshalJSON(b []byte) error {
	var tr translatedResults
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
		var c CommunityResults
		err := json.Unmarshal(tr.RawData, &c)
		if err != nil {
			// Note: Could be a slice, needs to be fixed
			if strings.Contains(err.Error(), "cannot unmarshal array") {
				var sliceOfCommunityResults []CommunityResults
				err := json.Unmarshal(tr.RawData, &sliceOfCommunityResults)
				if err == nil {
					c = sliceOfCommunityResults[0]
					break
				}
			}
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
		return fmt.Errorf("unsupported results type found: %v", tr.Type)
	}

	return nil
}

// UnmarshalJSON is a custom JSON unmarshaller implementation for the standard
// go json package to know how to properly interpret ScanSummaryResults from
// JSON.
func (u *UntranslatedResults) UnmarshalJSON(b []byte) error {
	// first look for results in the proper translated format
	// e.g. CommunityResults
	tr := &translatedResults{}
	err := json.Unmarshal(b, tr)
	if err != nil {
		// we have received invalid stringified json
		return fmt.Errorf("unable to unmarshal json")
	}

	// if there is a type and it is `community`
	// parse the data out
	if tr.Type == "community" {
		c := &CommunityResults{}
		err = json.Unmarshal(tr.RawData, c)
		if err != nil {
			return err
		}
		u.Community = c
		return nil
	}

	// it is not translated and not community
	// ur2 is required to keep the parser from
	// recursing here
	type ur2 UntranslatedResults
	err = json.Unmarshal(b, (*ur2)(u))
	if err != nil {
		// we have received invalid stringified json
		return fmt.Errorf("unable to unmarshal json")
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

// Dependency represents data for an individual requirement resolution
type Dependency struct {
	LatestVersion string `json:"latest_version" xml:"latest_version"`
	Org           string `json:"org" xml:"org"`
	Name          string `json:"name" xml:"name"`
	Type          string `json:"type" xml:"type"`
	Package       string `json:"package" xml:"package"`
	Version       string `json:"version" xml:"version"`
	Scope         string `json:"scope" xml:"scope"`
}

// DependencyMeta represents data for a summary of all dependencies resolved
type DependencyMeta struct {
	FirstDegreeCount     int `json:"first_degree_count" xml:"first_degree_count"`
	NoVersionCount       int `json:"no_version_count" xml:"no_version_count"`
	TotalUniqueCount     int `json:"total_unique_count" xml:"total_unique_count"`
	UpdateAvailableCount int `json:"update_available_count" xml:"update_available_count"`
}

// DependencyResults represents the data collected from a dependency scan.  It
// includes a list of the dependencies seen and meta data counts about those
// dependencies seen.
type DependencyResults struct {
	Dependencies []Dependency   `json:"dependencies" xml:"dependencies"`
	Meta         DependencyMeta `json:"meta" xml:"meta"`
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
	Ecosystems map[string]int `json:"ecosystems" xml:"ecosystems"`
}

// MarshalJSON meets the marshaller interface to custom wrangle an ecosystem
// result into the json shape
func (e EcosystemResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Ecosystems)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle the
// ecosystem scan into an ecosystem result
func (e *EcosystemResults) UnmarshalJSON(b []byte) error {
	var m map[string]int
	err := json.Unmarshal(b, &m)
	if err != nil {
		return fmt.Errorf("failed to unmarshal ecosystem result: %v", err.Error())
	}

	e.Ecosystems = m
	return nil
}

// ExternalVulnerabilitiesResults represents the data collected from an external
// vulnerability scan.  It includes the number of each vulnerability criticality
// seen within the project.
type ExternalVulnerabilitiesResults struct {
	Critical int `json:"critical" xml:"critical"`
	High     int `json:"high" xml:"high"`
	Medium   int `json:"medium" xml:"medium"`
	Low      int `json:"low" xml:"low"`
}

// LicenseResults represents the data collected from a license scan.  It
// includes the name and type of each license seen within the project.
type LicenseResults struct {
	*License `json:"license" xml:"license"`
}

// License represents a name and slice of types of licenses seen in a given file
type License struct {
	Name string        `json:"name" xml:"name"`
	Type []LicenseType `json:"type" xml:"type"`
}

// LicenseType represents a type of license such as MIT, Apache 2.0, etc
type LicenseType struct {
	Name string `json:"name" xml:"name"`
}

// VirusResults represents the data collected from a virus scan.  It includes
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

//VulnerabilityResults represents the data collected from a vulnerability scan.  It includes
// information of the vulnerabilities seen.
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
		Edition         interface{}                     `json:"edition" xml:"edition"`
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
