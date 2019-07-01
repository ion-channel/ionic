package tags

import "time"

const (
	// CreateTagEndpoint is a string representation of the current endpoint for getting projects
	CreateTagEndpoint = "v1/tag/createTag"
	// GetTagEndpoint is a string representation of the current endpoint for getting projects
	GetTagEndpoint = "v1/tag/getTag"
	// GetTagsEndpoint is a string representation of the current endpoint for getting projects
	GetTagsEndpoint = "v1/tag/getTags"
	// UpdateTagEndpoint is a string representation of the current endpoint for getting projects
	UpdateTagEndpoint = "v1/tag/updateTag"
)

//Tag is a client provided identifier to group projects
type Tag struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"team_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
