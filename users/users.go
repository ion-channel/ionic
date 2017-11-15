package users

import (
	"time"
)

// User is a representation of an Ion Channel User within the system
type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	ChatHandle string    `json:"chat_handle"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastActive time.Time `json:"last_active_at"`
	Metadata   string    `json:"metadata"`
	SysAdmin   bool      `json:"sys_admin"`
}
