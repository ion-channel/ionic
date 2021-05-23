package events

import (
	"encoding/json"
	"fmt"

	"github.com/ion-channel/ionic/users"
)

var validUserEventActions = map[string]string{
	AccountCreated:    AccountCreated,
	ForgotPassword:    ForgotPassword,
	PasswordChanged:   PasswordChanged,
	UserSignup:        UserSignup,
	UserSignupStarted: UserSignupStarted,
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

// UserEvent represents the user related segment of an Event within Ion Channel
type UserEvent struct {
	Action UserEventAction `json:"action"`
	User   users.User      `json:"user"`
	URL    string          `json:"url"`
	Team   string          `json:"team"`
}
