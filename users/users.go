package users

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// UsersCreateUserEndpoint is a string representation of the current endpoint for creating users
	UsersCreateUserEndpoint = "v1/users/createUser"
	// UsersGetSelfEndpoint is a string representation of the current endpoint for get user self
	UsersGetSelfEndpoint = "v1/users/getSelf"
	// UsersSubscribedForEventEndpoint is a string representation of the current endpoint for users subscribed for event
	UsersSubscribedForEventEndpoint = "v1/users/subscribedForEvent"
	// UsersGetUserEndpoint is a string representation of the current endpoint for getting user
	UsersGetUserEndpoint = "v1/users/getUser"
	// UsersGetUsers is a string representation of the current endpoint for getting users
	UsersGetUsers = "v1/users/getUsers"

	// TODO: the following endpoints need functions attached to them

	// UsersSignupEndpoint is a string representation of the current endpoint for user signup
	UsersSignupEndpoint = "v1/users/signup"
	// UsersSubscribeMeToEndpoint is a string representation of the current endpoint for adding my subscription
	UsersSubscribeMeToEndpoint = "v1/users/subscribeMeTo"
	//UsersAmISubscribedToEndpoint is a string representation of the current endpoint for my subscription
	UsersAmISubscribedToEndpoint = "v1/users/amISubscribedTo"
	//UsersUnsubscribeMeFromEndpoint is a string representation of the current endpoint for unsubscribe me
	UsersUnsubscribeMeFromEndpoint = "v1/users/unsubscribeMeFrom"
	//UsersMySubscriptionsEndpoint is a string representation of the current endpoint for getting my subscription
	UsersMySubscriptionsEndpoint = "v1/users/mySubscriptions"
	// UsersResetPasswordEndpoint is a string representation of the current endpoint for resetting user password
	UsersResetPasswordEndpoint = "v1/users/resetPassword"
	// UsersUpdateUserEndpoint is a string representation of the current endpoint for updating user
	UsersUpdateUserEndpoint = "v1/users/updateUser"
	// UsersDeleteUserEndpoint is a string representation of the current endpoint for deleting user
	UsersDeleteUserEndpoint = "v1/users/deleteUser"
)

// User is a representation of an Ion Channel User within the system
type User struct {
	ID                string            `json:"id"`
	Email             string            `json:"email"`
	Username          string            `json:"username"`
	ChatHandle        string            `json:"chat_handle"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
	LastActive        time.Time         `json:"last_active_at"`
	ExternallyManaged bool              `json:"externally_managed"`
	Metadata          json.RawMessage   `json:"metadata"`
	SysAdmin          bool              `json:"sys_admin"`
	System            bool              `json:"system"`
	Teams             map[string]string `json:"teams"`
}

// String returns a JSON formatted string of the user object
func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return fmt.Sprintf("failed to format user: %v", err.Error())
	}
	return string(b)
}
