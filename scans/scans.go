package scans

import (
	"encoding/json"
	"strings"
	"time"
)

// ScanSummary is an Ion Channel representation of the summary produced by an
// individual scan on a project.  It contains all the details the Ion Channel
// platform discovers for that scan.
type ScanSummary struct {
	ID          string    `json:"id" xml:"id"`
	TeamID      string    `json:"team_id" xml:"team_id"`
	ProjectID   string    `json:"project_id" xml:"project_id"`
	AnalysisID  string    `json:"analysis_id" xml:"analysis_id"`
	Summary     string    `json:"summary" xml:"summary"`
	Results     Results   `json:"results" xml:"results"`
	UpdatedAt   time.Time `json:"updated_at" xml:"updated_at"`
	CreatedAt   time.Time `json:"created_at" xml:"created_at"`
	Duration    float64   `json:"duration" xml:"duration"`
	Passed      bool      `json:"passed" xml:"passed"`
	Risk        string    `json:"risk" xml:"risk"`
	Name        string    `json:"name" xml:"name"`
	Description string    `json:"description" xml:"description"`
	Type        string    `json:"type" xml:"type"`
}

// Results is an Ion Channel representation of the results from a
// scan summary.  It contains what type of results and the data pertaining to
// the results.
type Results struct {
	Type string      `json:"type" xml:"type"`
	Data interface{} `json:"data" xml:"data"`
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
	case "ecosystems":
		var e EcosystemResults
		err := json.Unmarshal(tr.RawData, &e)
		if err != nil {
			return err
		}

		r.Data = e
	case "virus":
		var v VirusResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return err
		}

		r.Data = v
	default:
	}

	return nil
}

type AboutYMLResults struct {
	Message string `json:"message" xml:"message"`
	Valid   bool   `json:"valid" xml:"valid"`
	Content string `json:"content" xml:"content"`
}

type EcosystemResults struct {
	Ecosystems []struct {
		Ecosystem string `json:"ecosystem" xml:"ecosystem"`
		Lines     string `json:"lines" xml:"lines"`
	} `json:"ecosystems" xml:"ecosystems"`
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
