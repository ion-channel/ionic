package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	usersSubscribedForEventEndpoint = "v1/users/subscribedForEvent"
	usersGetSelfEndpoint            = "v1/users/getSelf"
)

// User is a representation of an Ion Channel User within the system
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	ChatHandle string `json:"chat_handle"`
	SysAdmin   bool   `json:"sys_admin"`
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
	b, err = ic.post(usersSubscribedForEventEndpoint, nil, *buff, nil)
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

func (ic *IonClient) GetSelf() (*User, error) {
	b, err := ic.get(usersGetSelfEndpoint, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get self: %v", err.Error())
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, fmt.Errorf("cannot parse user: %v", err.Error())
	}

	return &user, nil
}
