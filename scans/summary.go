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
	Passed      bool            `json:"passed" xml:"passed"`
	Risk        string          `json:"risk" xml:"risk"`
	Name        string          `json:"name" xml:"name"`
	Description string          `json:"description" xml:"description"`
	Type        string          `json:"type" xml:"type"`
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
