package users

// User is a representation of an Ion Channel User within the system
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	ChatHandle string `json:"chat_handle"`
	SysAdmin   bool   `json:"sys_admin"`
}
