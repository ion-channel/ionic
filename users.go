package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/errors"
	"github.com/ion-channel/ionic/events"
	"github.com/ion-channel/ionic/users"
)

const (
	usersCreateUserEndpoint         = "v1/users/createUser"
	usersGetSelfEndpoint            = "v1/users/getSelf"
	usersSubscribedForEventEndpoint = "v1/users/subscribedForEvent"
	usersGetUserEndpoint            = "v1/users/getUser"
	usersGetUsers                   = "v1/users/getUsers"
)

type createUserOptions struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// CreateUser takes an email, username, and password.  The username and password
// are not required, and can be left blank if so chosen.  It will return the
// instantiated user object from the API or an error if it encounters one with
// the API.
func (ic *IonClient) CreateUser(email, username, password, token string) (*users.User, error) {
	if email == "" {
		return nil, fmt.Errorf("create user: email is required")
	}

	opts := createUserOptions{
		Email:                email,
		Username:             username,
		Password:             password,
		PasswordConfirmation: password,
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, errors.Prepend("create user: failed to marshal request", err)
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(usersCreateUserEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, errors.Prepend("create user", err)
	}

	var u users.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, errors.Prepend("create user: failed to unmarshal user", err)
	}

	return &u, nil
}

// GetUsersSubscribedForEvent takes an event and token, and returns a list of users
// subscribed to that event and returns an error if there are JSON marshaling
// or unmarshaling issues or issues with the request
func (ic *IonClient) GetUsersSubscribedForEvent(event events.Event, token string) ([]users.User, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return nil, errors.Prepend("get users subscribed for event: failed marshaling event", err)
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(usersSubscribedForEventEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, errors.Prepend("get users subscribed for event", err)
	}

	var users struct {
		Users []users.User `json:"users"`
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, errors.Prepend("get users subscribed for event: failed unmarshaling users", err)
	}

	return users.Users, nil
}

// GetSelf returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetSelf(token string) (*users.User, error) {
	b, err := ic.Get(usersGetSelfEndpoint, token, nil, nil, nil)
	if err != nil {
		return nil, errors.Prepend("get self", err)
	}

	var user users.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, errors.Prepend("get self: failed unmarshaling user", err)
	}

	return &user, nil
}

// GetUser returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetUser(id, token string) (*users.User, error) {
	params := &url.Values{}
	params.Set("id", id)

	b, err := ic.Get(usersGetUserEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, errors.Prepend("get user", err)
	}

	var user users.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, errors.Prepend("get user: failed unmarshaling user", err)
	}

	return &user, nil
}

// GetUsers requests and returns all users for a given installation
func (ic *IonClient) GetUsers(token string) ([]users.User, error) {
	b, err := ic.Get(usersGetUsers, token, nil, nil, nil)
	if err != nil {
		return nil, errors.Prepend("get users", err)
	}

	var us []users.User
	err = json.Unmarshal(b, &us)
	if err != nil {
		return nil, errors.Prepend("get users: failed unmarshaling users", err)
	}

	return us, nil
}
