package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Evaluation is an Ion Channel representation of an evaluation produced by an
// individual scan on a project. It contains all the details of the evaluation
// of a scan by a rule.
type Evaluation struct {
	*evaluation
	TranslatedResults   *TranslatedResults   `json:"-"`
	UntranslatedResults *UntranslatedResults `json:"-"`
}

type evaluation struct {
	ID          string          `json:"id"`
	TeamID      string          `json:"team_id"`
	ProjectID   string          `json:"project_id"`
	AnalysisID  string          `json:"analysis_id"`
	RuleID      string          `json:"rule_id"`
	RulesetID   string          `json:"ruleset_id"`
	Summary     string          `json:"summary"`
	Results     json.RawMessage `json:"results"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Duration    float64         `json:"duration"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Risk        string          `json:"risk"`
	Type        string          `json:"type"`
	Passed      bool            `json:"passed"`
}

// NewEval returns an Evaluation that wont throw nil pointer exceptions
func NewEval() *Evaluation {
	return &Evaluation{evaluation: &evaluation{}}
}

// Translate performs a one way translation on an evaluation by translating the
// UntranslatedResults if they are not nil, putting the output into Translated
// results, and setting UntranslatedResults to be nil. It returns an error if
// it encounters a JSON error when translating the results.
func (e *Evaluation) Translate() error {
	if e.UntranslatedResults != nil {
		translated := e.UntranslatedResults.Translate()
		e.TranslatedResults = translated
		e.UntranslatedResults = nil

		b, err := json.Marshal(e.TranslatedResults)
		if err != nil {
			return fmt.Errorf("failed to translate scan: %v", err.Error())
		}

		if e.evaluation == nil {
			e.evaluation = &evaluation{}
		}

		e.Results = b
	}

	return nil
}

// MarshalJSON meets the marshaller interface to custom wrangle translated or
// untranslated results into the same results key for the JSON
func (e *Evaluation) MarshalJSON() ([]byte, error) {
	if e.TranslatedResults != nil {
		b, err := json.Marshal(e.TranslatedResults)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal translated results: %v", err.Error())
		}

		e.Results = b
	}

	if e.UntranslatedResults != nil {
		b, err := json.Marshal(e.UntranslatedResults)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal untranslated results: %v", err.Error())
		}

		e.Results = b
	}

	return json.Marshal(e.evaluation)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle results from
// a singular results key into the correct translated or untranslated results
// field
func (e *Evaluation) UnmarshalJSON(b []byte) error {
	var ee evaluation
	err := json.Unmarshal(b, &ee)
	if err != nil {
		return fmt.Errorf("failed to unmarshal evaluation: %v", err.Error())
	}

	e.evaluation = &ee

	var tr TranslatedResults
	err = json.Unmarshal(e.Results, &tr)
	if err != nil {
		if strings.Contains(err.Error(), "unsupported results type found") {
			var un UntranslatedResults
			err := json.Unmarshal(e.Results, &un)
			if err != nil {
				return fmt.Errorf("failed to unmarshal untranslated results: %v", err.Error())
			}

			e.UntranslatedResults = &un
			return e.Translate()
		}

		return fmt.Errorf("failed to unmarshal translated results: %v", err.Error())
	}

	e.TranslatedResults = &tr

	return nil
}
