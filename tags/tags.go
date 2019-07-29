package tags

import "time"

const (
	// CreateTagEndpoint is a string representation of the current endpoint for creating tag
	CreateTagEndpoint = "v1/tag/createTag"
	// GetTagEndpoint is a string representation of the current endpoint for getting tag
	GetTagEndpoint = "v1/tag/getTag"
	// GetTagsEndpoint is a string representation of the current endpoint for getting tags
	GetTagsEndpoint = "v1/tag/getTags"
	// UpdateTagEndpoint is a string representation of the current endpoint for update tag
	UpdateTagEndpoint = "v1/tag/updateTag"

	// TODO: The following endpoints will need to have functions attached to them

	// DeleteTagsEndpoint is a string representation of the current endpoint for deleting tags
	DeleteTagsEndpoint = "/v1/tag/deleteTags"
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
