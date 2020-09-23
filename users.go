package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/errors"
	"github.com/ion-channel/ionic/requests"
	"github.com/ion-channel/ionic/users"
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

	b, err = ic.Post(users.UsersCreateUserEndpoint, token, nil, *buff, nil)
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

// GetSelf returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetSelf(token string) (*users.User, error) {
	b, _, err := ic.Get(users.UsersGetSelfEndpoint, token, nil, nil, nil)
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

	b, _, err := ic.Get(users.UsersGetUserEndpoint, token, params, nil, nil)
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
	b, _, err := ic.Get(users.UsersGetUsers, token, nil, nil, nil)
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

//GetUserNames takes slice of ids and returns user names with their ids
func (ic *IonClient) GetUserNames(ids []string, token string) ([]users.NameAndID, error) {
	byIDs := requests.ByIDs{
		IDs: ids,
	}

	b, err := json.Marshal(byIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	r, err := ic.Post(users.UsersGetUserNames, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user names: %v", err.Error())
	}

	var s []users.NameAndID
	err = json.Unmarshal(r, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user names: %v", err.Error())
	}

	return s, nil
}
