package users

import (
	"encoding/json"
	"time"
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
