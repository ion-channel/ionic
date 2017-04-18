package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	getUsersSubscribedForEventEndpoint = "v1/users/subscribedForEvent"
)

// User is a representation of an Ion Channel User within the system
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetUsersSubscribedForEvent takes an event and returns a list of users
// subscribed to that event and returns an error if there are JSON marshalling
// or unmarshalling issues or issues with the request
func (ic *IonClient) GetUsersSubscribedForEvent(event Event) ([]User, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.post(getUsersSubscribedForEventEndpoint, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err.Error())
	}

	var users struct {
		Users []User `json:"users"`
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, fmt.Errorf("cannot parse users: %v", err.Error())
	}

	return users.Users, nil
}
