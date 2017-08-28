package teams

import (
	"time"
)

// Team is a representation of an Ion Channel Team within the system
type Team struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	SysAdmin     bool      `json:"sys_admin"`
	POCName      string    `json:"poc_name"`
	POCEmail     string    `json:"poc_email"`
	POCNameHash  string    `json:"poc_name_hash"`
	POCEmailHash string    `json:"poc_email_hash"`
}
