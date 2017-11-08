package tags

import "time"

type Tag struct {
	ID          string      `json:"id"`
	TeamID      string      `json:"team_id"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
