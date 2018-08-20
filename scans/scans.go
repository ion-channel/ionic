package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Scan represents the data collected from an individual scan on a project.
type Scan struct {
	*scan
	TranslatedResults   *TranslatedResults   `json:"-"`
	UntranslatedResults *UntranslatedResults `json:"-"`
}

type scan struct {
	ID          string          `json:"id"`
	TeamID      string          `json:"team_id"`
	ProjectID   string          `json:"project_id"`
	AnalysisID  string          `json:"analysis_id"`
	Summary     string          `json:"summary"`
	Results     json.RawMessage `json:"results"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Duration    float64         `json:"duration"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
}

// NewScan creates and returns a new Scan struct
func NewScan(
	id, teamID, projectID, analysisID, summary, name, description string,
	results json.RawMessage,
	createdAt, updatedAt time.Time,
	duration float64,
) (*Scan, error) {
	s := scan{
		ID:          id,
		TeamID:      teamID,
		ProjectID:   projectID,
		AnalysisID:  analysisID,
		Summary:     summary,
		Results:     results,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Duration:    duration,
		Name:        name,
		Description: description,
	}
	var tr TranslatedResults

	var ur UntranslatedResults
	err := json.Unmarshal(results, &ur)
	if err != nil {
		err = json.Unmarshal(results, &tr)
		if err != nil {
			return nil, fmt.Errorf("failed to get valid results: %v", err.Error())
		}
	} else {
		tr = *ur.Translate()
	}
	bigS := &Scan{
		scan:              &s,
		TranslatedResults: &tr,
	}

	return bigS, nil
}

// Translate performs a one way translation on a scan by translating the
// UntranslatedResults if they are not nil, putting the output into Translated
// results, and setting UntranslatedResults to be nil. It returns an error if
// it encounters a JSON error when translating the results.
func (s *Scan) Translate() error {
	if s.UntranslatedResults != nil {
		translated := s.UntranslatedResults.Translate()
		s.TranslatedResults = translated
		s.UntranslatedResults = nil

		b, err := json.Marshal(s.TranslatedResults)
		if err != nil {
			return fmt.Errorf("failed to translate scan: %v", err.Error())
		}

		if s.scan == nil {
			s.scan = &scan{}
		}

		s.Results = b
	}

	return nil
}

// MarshalJSON meets the marshaller interface to custom wrangle translated or
// untranslated results into the same results key for the JSON
func (s *Scan) MarshalJSON() ([]byte, error) {
	if s.TranslatedResults != nil {
		b, err := json.Marshal(s.TranslatedResults)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal translated results: %v", err.Error())
		}

		s.Results = b
	}

	if s.UntranslatedResults != nil {
		b, err := json.Marshal(s.UntranslatedResults)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal untranslated results: %v", err.Error())
		}

		s.Results = b
	}

	return json.Marshal(s.scan)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle results from
// a singular results key into the correct translated or untranslated results
// field
func (s *Scan) UnmarshalJSON(b []byte) error {
	var ss scan
	err := json.Unmarshal(b, &ss)
	if err != nil {
		return fmt.Errorf("failed to unmarshal scans: %v", err.Error())
	}

	s.scan = &ss

	var tr TranslatedResults
	err = json.Unmarshal(s.Results, &tr)
	if err != nil {
		if strings.Contains(err.Error(), "unsupported results type found") {
			var un UntranslatedResults
			err := json.Unmarshal(s.Results, &un)
			if err != nil {
				return fmt.Errorf("failed to unmarshal untranslated results: %v", err.Error())
			}

			s.UntranslatedResults = &un

			return nil
		}

		return fmt.Errorf("failed to unmarshal translated results: %v", err.Error())
	}

	s.TranslatedResults = &tr

	return nil
}
