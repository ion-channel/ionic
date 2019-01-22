package digests

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

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

// NewFromEval takes a title, data type and value to attempt constructing a digest. It
// returns a digest and any error that it encounters while trying to construct
// the digest.
func NewFromEval(index int, title string, dataType string, value interface{}, eval *scans.Evaluation) (*Digest, error) {
	var data []byte
	var err error

	switch strings.ToLower(dataType) {
	case "bool", "boolean":
		b, ok := value.(bool)
		if !ok {
			return nil, ErrFailedValueAssertion
		}

		data, err = json.Marshal(boolean{b})
	case "chars":
		c, ok := value.(string)
		if !ok {
			return nil, ErrFailedValueAssertion
		}

		data, err = json.Marshal(chars{c})
	case "count":
		c, ok := value.(int)
		if !ok {
			return nil, ErrFailedValueAssertion
		}

		data, err = json.Marshal(count{c})
	case "list":
		l, ok := value.([]string)
		if !ok {
			return nil, ErrFailedValueAssertion
		}

		data, err = json.Marshal(&list{l})
	case "percent":
		p, ok := value.(float64)
		if !ok {
			return nil, ErrFailedValueAssertion
		}

		data, err = json.Marshal(percent{math.Round(p*100) / 100})
	default:
		return nil, ErrUnsupportedType
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal digest data: %v", err.Error())
	}

	d := &Digest{
		Index: index,
		Title: title,
		Data:  data,

		ScanID:    eval.ID,
		RuleID:    eval.RuleID,
		RulesetID: eval.RulesetID,

		Evaluated:     (strings.ToLower(eval.Type) != "not evaluated"),
		Passed:        eval.Passed,
		PassedMessage: eval.Description,
	}

	return d, nil
}
