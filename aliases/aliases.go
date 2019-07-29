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
	// AddAliasEndpoint allows the user to attach an alias to a project.  Consists of org, name, version. Requires team id and project id.
	AddAliasEndpoint = "v1/project/addAlias"
)
