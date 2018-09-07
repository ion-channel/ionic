package events

import (
	"encoding/json"
	"fmt"
)

var validAnalysisEventActions = map[string]string{
	AnalysisFailed:   AnalysisFailed,
	AnalysisFinished: AnalysisFinished,
	AnalysisPassed:   AnalysisPassed,
}

// AnalysisEventAction represents possible actions related to a analysis event
type AnalysisEventAction string

// UnmarshalJSON is a custom unmarshaller for validating an AnalysisEventAction
// or it returns an error
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

//AnalysisEvent identifies the result of an analysis of a project
type AnalysisEvent struct {
	Action    AnalysisEventAction `json:"action"`
	Analysis  string              `json:"analysis"`
	URL       string              `json:"url"`
	TeamID    string              `json:"team_id"`
	ProjectID string              `json:"project_id"`
}
