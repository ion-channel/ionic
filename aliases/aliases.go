package aliases

import "time"

//Alias map user defined project names to a Common Platform Enumeration (CPE).
type Alias struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Org       string    `json:"org"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   string    `json:"version"`
}

const (
	// AddAliasEndpoint is a string representation of the current endpoint for getting aliases
	AddAliasEndpoint = "v1/project/addAlias"

	// TODO: The following will need to have functions attached to them:

	// DeleteAliasEndpoint is a string representation of the current endpoint for deleting an alias
	DeleteAliasEndpoint = "v1/project/deleteAlias"
	// UpdateAliasEndpoint is a string representation of the current endpoint for updating an alias
	UpdateAliasEndpoint = "v1/project/updateAlias"
)
