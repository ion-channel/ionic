package events

import (
	"encoding/json"
	"fmt"
)

var validProjectEventActions = map[string]string{
	ProjectAdded:       ProjectAdded,
	VersionAdded:       VersionAdded,
	VulnerabilityAdded: VulnerabilityAdded,
}

// ProjectEventAction represents possible actions related to a project event
type ProjectEventAction string

// UnmarshalJSON is a custom unmarshaller for enforcing a project event action is
// a valid value and returns an error if the value is invalid
func (a *ProjectEventAction) UnmarshalJSON(b []byte) error {
	var aStr string
	err := json.Unmarshal(b, &aStr)
	if err != nil {
		return err
	}

	_, ok := validProjectEventActions[aStr]
	if !ok {
		return fmt.Errorf("invalid project event action")
	}

	*a = ProjectEventAction(validProjectEventActions[aStr])
	return nil
}

// ProjectEvent represents a project within an Event within Ion Channel
type ProjectEvent struct {
	Action   ProjectEventAction `json:"action"`
	Project  string             `json:"project"`
	Org      string             `json:"org"`
	Versions []string           `json:"versions"`
	URL      string             `json:"url"`
}
