package ionic

import (
	"encoding/json"
	"fmt"
)

// Event represents a singular occurance of a change within the Ion Channel
// system that can be emmitted to trigger a notification
type Event struct {
	Vulnerability *VulnerabilityEvent `json:"vulnerability,omitempty"`
	User          *UserEvent          `json:"user,omitempty"`
}

// Append takes an event to join and leaves the union of the two events
func (e *Event) Append(toAppend Event) {
	for _, v := range toAppend.Vulnerability.Updates {
		if !e.contains(v) {
			e.Vulnerability.Updates = append(e.Vulnerability.Updates, v)
		}
	}
}

func (e *Event) contains(vuln string) bool {
	if e.Vulnerability.Updates == nil {
		return false
	}

	for _, existing := range e.Vulnerability.Updates {
		if existing == vuln {
			return true
		}
	}

	return false
}

// VulnerabilityEvent represents the vulnerability releated segement of an Event
// within Ion Channel
type VulnerabilityEvent struct {
	Updates  []string       `json:"updates,omitempty"`
	Projects []ProjectEvent `json:"projects,omitempty"`
}

// ProjectEvent represents a project within an Event within Ion Channel
type ProjectEvent struct {
	Project  string   `json:"project"`
	Org      string   `json:"org"`
	Versions []string `json:"versions"`
}

var validUserEventActions = map[string]string{
	"account_created":  "account_created",
	"forgot_password":  "forgot_password",
	"password_changed": "password_changed",
}

// UserEventAction represents possible actions related to a user event
type UserEventAction string

// UnmarshalJSON is a custom unmarshaller for enforcing a user event action is
// a valid value and returns an error if the value is invalid
func (a *UserEventAction) UnmarshalJSON(b []byte) error {
	var aStr string
	err := json.Unmarshal(b, &aStr)
	if err != nil {
		return err
	}

	_, ok := validUserEventActions[aStr]
	if !ok {
		return fmt.Errorf("invalid user event action")
	}

	*a = UserEventAction(validUserEventActions[aStr])
	return nil
}

// UserEvent represents the user releated segement of an Event within Ion Channel
type UserEvent struct {
	Action UserEventAction `json:"action"`
	User   User            `json:"user"`
}
