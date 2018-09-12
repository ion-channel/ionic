package teams

import (
	"time"
)

// Team is a representation of an Ion Channel Team within the system
type Team struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	Name       string    `json:"name"`
	Delivering bool      `json:"delivering"`
	SysAdmin   bool      `json:"sys_admin"`
	POCName    string    `json:"poc_name"`
	POCEmail   string    `json:"poc_email"`
}
