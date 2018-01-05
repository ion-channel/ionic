package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/events"
	"github.com/ion-channel/ionic/users"
)

const (
	usersCreateUserEndpoint         = "v1/users/createUser"
	usersGetSelfEndpoint            = "v1/users/getSelf"
	usersSubscribedForEventEndpoint = "v1/users/subscribedForEvent"
)

// CreateUser takes an email, username, and password.  The username and password
// are not required, and can be left blank if so chosen.  It will return the
// instantiated user object from the API or an error if it encounters one with
// the API.
func (ic *IonClient) CreateUser(email, username, password string) (*users.User, error) {
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}

	params := &url.Values{}
	params.Set("email", email)
	if username != "" {
		params.Set("username", username)
	}
	if password != "" {
		params.Set("password", password)
		params.Set("password_confirmation", password)
	}

	b, err := ic.Post(usersCreateUserEndpoint, params, bytes.Buffer{}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err.Error())
	}

	var u users.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response from api: %v", err.Error())
	}

	return &u, nil
}

// GetUsersSubscribedForEvent takes an event and returns a list of users
// subscribed to that event and returns an error if there are JSON marshalling
// or unmarshalling issues or issues with the request
func (ic *IonClient) GetUsersSubscribedForEvent(event events.Event) ([]users.User, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event to JSON: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(usersSubscribedForEventEndpoint, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err.Error())
	}

	var users struct {
		Users []users.User `json:"users"`
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, fmt.Errorf("cannot parse users: %v", err.Error())
	}

	return users.Users, nil
}

// GetSelf returns the user object associated with the bearer token in use by
// the Ion Client.  An error is returned if the client cannot talk to the API
// or the returned user object is nil or blank
func (ic *IonClient) GetSelf() (*users.User, error) {
	b, err := ic.Get(usersGetSelfEndpoint, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get self: %v", err.Error())
	}

	var user users.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, fmt.Errorf("cannot parse user: %v", err.Error())
	}

	return &user, nil
}
