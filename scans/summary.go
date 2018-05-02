package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Summary is an Ion Channel representation of the summary produced by an
// individual scan on a project.  It contains all the details the Ion Channel
// platform discovers for that scan.
type Summary struct {
	*summary
	TranslatedResults   *TranslatedResults   `json:"-"`
	UntranslatedResults *UntranslatedResults `json:"-"`
}

type summary struct {
	ID          string          `json:"id" xml:"id"`
	TeamID      string          `json:"team_id" xml:"team_id"`
	ProjectID   string          `json:"project_id" xml:"project_id"`
	AnalysisID  string          `json:"analysis_id" xml:"analysis_id"`
	Summary     string          `json:"summary" xml:"summary"`
	Results     json.RawMessage `json:"results"`
	CreatedAt   time.Time       `json:"created_at" xml:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" xml:"updated_at"`
	Duration    float64         `json:"duration" xml:"duration"`
	Name        string          `json:"name" xml:"name"`
	Description string          `json:"description" xml:"description"`
	Risk        string          `json:"risk" xml:"risk"`
	Type        string          `json:"type" xml:"type"`
	Passed      bool            `json:"passed" xml:"passed"`
}

// NewSummary takes a scan and returns the appropriate fields as part of a
// summary
func NewSummary(s *Scan) *Summary {
	if s.scan == nil {
		s.scan = &scan{}
	}

	return &Summary{
		summary: &summary{
			ID:          s.ID,
			TeamID:      s.TeamID,
			ProjectID:   s.ProjectID,
			Summary:     s.Summary,
			Results:     s.Results,
			CreatedAt:   s.CreatedAt,
			UpdatedAt:   s.UpdatedAt,
			Duration:    s.Duration,
			Name:        s.Name,
			Description: s.Description,
		},
		UntranslatedResults: s.UntranslatedResults,
		TranslatedResults:   s.TranslatedResults,
	}
}

// Translate performs a one way translation on a summary by translating the
// UntranslatedResults if they are not nil, putting the output into Translated
// results, and setting UntranslatedResults to be nil. It returns an error if
// it encounters a JSON error when translating the results.
func (s *Summary) Translate() error {
	if s.UntranslatedResults != nil {
		translated := s.UntranslatedResults.Translate()
		s.TranslatedResults = translated
		s.UntranslatedResults = nil

		b, err := json.Marshal(s.TranslatedResults)
		if err != nil {
			return fmt.Errorf("failed to translate scan: %v", err.Error())
		}

		if s.summary == nil {
			s.summary = &summary{}
		}

		s.Results = b
	}

	return nil
}

// MarshalJSON meets the marshaller interface to custom wrangle translated or
// untranslated results into the same results key for the JSON
func (s *Summary) MarshalJSON() ([]byte, error) {
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

	return json.Marshal(s.summary)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle results from
// a singular results key into the correct translated or untranslated results
// field
func (s *Summary) UnmarshalJSON(b []byte) error {
	var ss summary
	err := json.Unmarshal(b, &ss)
	if err != nil {
		return fmt.Errorf("failed to unmarshal scans summary: %v", err.Error())
	}

	s.summary = &ss

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
