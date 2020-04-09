package digests

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/ion-channel/ionic/scanner"
	"github.com/ion-channel/ionic/scans"
)

var (
	// ErrFailedValueAssertion is the error returned when the value for
	// constructing a new digest does not convert to the digest data type
	// specified
	ErrFailedValueAssertion = fmt.Errorf("failed to assert value for given data type")

	// ErrUnsupportedType is the error returned when the data type for
	// constructing the digest is not a valid type
	ErrUnsupportedType = fmt.Errorf("unsupported digest data type")
)

// Digest is only the important parts for display
type Digest struct {
	Index          int             `json:"index"`
	Title          string          `json:"title"`
	Data           json.RawMessage `json:"data"`
	SourceData     json.RawMessage `json:"source_data"`
	ScanID         string          `json:"scan_id"`
	RuleID         string          `json:"rule_id"`
	RulesetID      string          `json:"ruleset_id"`
	Evaluated      bool            `json:"evaluated"`
	Pending        bool            `json:"pending"`
	Passed         bool            `json:"passed"`
	PassedMessage  string          `json:"passed_message"`
	Warning        bool            `json:"warning"`
	WarningMessage string          `json:"warning_message,omitempty"`
	Errored        bool            `json:"errored"`
	ErroredMessage string          `json:"errored_message,omitempty"`

	singularTitle string
	pluralTitle   string
}

type boolean struct {
	Bool bool `json:"bool"`
}

type chars struct {
	Chars string `json:"chars"`
}

type count struct {
	Count int `json:"count"`
}

type list struct {
	List []string `json:"list"`
}

type percent struct {
	Percent float64 `json:"percent"`
}

// String returns a JSON formatted string of the digest object
func (d Digest) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("failed to format user: %v", err.Error())
	}
	return string(b)
}

// NewDigest takes a scan status, ordering index, and a singular and plural
// version of the digest title. The status is leveraged to determine the first
// levels of a digest status. The index provides an ordered location if a list
// of Digests are sorted. The singular and plural titles are used to display as
// singular title if data is added that warrants a singular title, and otherwise
// will favor a plural title. A newly constructed Digest with the appropriate
// settings is returned.
func NewDigest(status *scanner.ScanStatus, index int, singular, plural string) *Digest {
	// We have other states now
	errored := false
	pending := true
	erroredMsg := "evaluation not received"

	if status != nil {
		switch status.Status {
		case scanner.AnalysisStatusErrored:
			errored = true
			erroredMsg = status.Message
			pending = false
		case scanner.AnalysisStatusQueued, scanner.AnalysisStatusAnalyzing:
			pending = true
		case scanner.AnalysisStatusFinished, scanner.AnalysisStatusFailed:
			pending = false
		default:
			errored = true
			pending = true
			erroredMsg = "Unknown status"
		}
	}

	d := &Digest{
		Index:          index,
		Pending:        pending,
		Errored:        errored,
		ErroredMessage: erroredMsg,
		Title:          plural,
		singularTitle:  singular,
		pluralTitle:    plural,
	}

	if status != nil {
		d.ScanID = status.ID
	}

	return d
}

// AppendEval takes an evaluation, data type, and value to interleave into the
// digest. It will also use the data to determine whether or not to show it's
// singular or plural title. It returns an error if the data type does not match
// the value provided.
func (d *Digest) AppendEval(eval *scans.Evaluation, dataType string, value interface{}) error {
	var data []byte
	var err error
	title := d.pluralTitle

	switch strings.ToLower(dataType) {
	case "bool", "boolean":
		b, ok := value.(bool)
		if !ok {
			return ErrFailedValueAssertion
		}

		data, err = json.Marshal(boolean{b})
	case "chars":
		c, ok := value.(string)
		if !ok {
			return ErrFailedValueAssertion
		}

		data, err = json.Marshal(chars{c})
	case "count":
		c, ok := value.(int)
		if !ok {
			return ErrFailedValueAssertion
		}

		data, err = json.Marshal(count{c})
		if c == 1 {
			title = d.singularTitle
		}

		// Counts less than 0 are an indicator of error state. Record where the
		// information came from, but do not overwrite errors and other states.
		if c < 0 {
			d.ScanID = eval.ID
			d.RuleID = eval.RuleID
			d.RulesetID = eval.RulesetID

			return nil
		}
	case "list":
		l, ok := value.([]string)
		if !ok {
			return ErrFailedValueAssertion
		}

		data, err = json.Marshal(&list{l})
		if len(l) == 1 {
			title = d.singularTitle
		}
	case "percent":
		p, ok := value.(float64)
		if !ok {
			return ErrFailedValueAssertion
		}

		data, err = json.Marshal(percent{math.Round(p*100) / 100})
	default:
		return ErrUnsupportedType
	}

	if err != nil {
		return fmt.Errorf("failed to marshal digest data: %v", err.Error())
	}

	d.Data = data
	d.Title = title
	d.Errored = false
	d.ErroredMessage = ""
	d.ScanID = eval.ID
	d.RuleID = eval.RuleID
	d.RulesetID = eval.RulesetID
	d.Evaluated = (strings.ToLower(eval.Type) != "not evaluated")
	d.Passed = eval.Passed
	d.PassedMessage = eval.Description

	return nil
}

// MarshalSourceData takes a data struct, generates json
// and sets the source data on the digest
func (d *Digest) MarshalSourceData(data interface{}, t string) error {
	r := scans.TranslatedResults{
		Type: t,
		Data: data,
	}
	b, err := json.Marshal(&r)
	if err != nil {
		return fmt.Errorf("failed to marshal digest data: %v", err.Error())
	}
	d.SourceData = b
	return nil
}

// UseSingularTitle forcibly sets the title to be singular
func (d *Digest) UseSingularTitle() {
	d.Title = d.singularTitle
}

// UsePluralTitle forcibly sets the title to be plural
func (d *Digest) UsePluralTitle() {
	d.Title = d.pluralTitle
}
