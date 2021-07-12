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
	// UsersGetUserEndpoint is a string representation of the current endpoint for getting user
	UsersGetUserEndpoint = "v1/users/getUser"
	// UsersGetUsers is a string representation of the current endpoint for getting users
	UsersGetUsers = "v1/users/getUsers"
	// UsersGetUserNames is a string representation of the current endpoint for getting users
	UsersGetUserNames = "v1/users/getUserNames"
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
	Organizations     map[string]string `json:"organizations"`
	Teams             map[string]string `json:"teams"`
}

// NameAndID represents the data object for user name and its ID
type NameAndID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// String returns a JSON formatted string of the user object
func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return fmt.Sprintf("failed to format user: %v", err.Error())
	}
	return string(b)
}

// IsMemberOfOrganization takes a team id and returns true if user is a member of that team.
func (u User) IsMemberOfOrganization(id string) bool {
	_, ok := u.Organizations[id]
	return ok
}

// IsAdminOfOrganization takes a team id and returns true if user is an admin of that team.
func (u User) IsAdminOfOrganization(id string) bool {
	return u.Organizations[id] == "admin"
}

// IsMemberOfTeam takes a team id and returns true if user is a member of that team.
func (u User) IsMemberOfTeam(id string) bool {
	_, ok := u.Teams[id]
	return ok
}

// IsAdminOfTeam takes a team id and returns true if user is an admin of that team.
func (u User) IsAdminOfTeam(id string) bool {
	return u.Teams[id] == "admin"
}
