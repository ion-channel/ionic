package events

import (
	"encoding/json"
	"fmt"
)

var validAnalysisEventActions = map[string]string{
	"analysis_failed":   "analysis_failed",
	"analysis_finished": "analysis_finished",
	"analysis_passed":   "analysis_passed",
}

// AnalysisEventAction represents possible actions related to a analysis event
type AnalysisEventAction string

// UnmarshalJSON is a custom unmarshaller for enforcing a analysis event action is
// a valid value and returns an error if the value is invalid
func (a *AnalysisEventAction) UnmarshalJSON(b []byte) error {
	var aStr string
	err := json.Unmarshal(b, &aStr)
	if err != nil {
		return err
	}

	_, ok := validAnalysisEventActions[aStr]
	if !ok {
		return fmt.Errorf("invalid analysis event action")
	}

	*a = AnalysisEventAction(validAnalysisEventActions[aStr])
	return nil
}

type AnalysisEvent struct {
	Action   AnalysisEventAction `json:"action"`
	Analysis string              `json:"analysis"`
	URL      string              `json:"url"`
}
