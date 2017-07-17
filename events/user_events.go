package events

import (
	"encoding/json"
	"fmt"

	"github.com/ion-channel/ionic/users"
)

var validUserEventActions = map[string]string{
	"account_created":  "account_created",
	"forgot_password":  "forgot_password",
	"password_changed": "password_changed",
	"user_signup":      "user_signup",
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
	User   users.User      `json:"user"`
	Link   string          `json:"link"`
	Team   string          `json:"team"`
}
