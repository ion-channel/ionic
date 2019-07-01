package tags

import "time"

const (
	CreateTagEndpoint = "v1/tag/createTag"
	GetTagEndpoint    = "v1/tag/getTag"
	GetTagsEndpoint   = "v1/tag/getTags"
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
