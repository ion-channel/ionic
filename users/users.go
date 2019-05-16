package users

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ctxKey int

const (
	userCtxKey ctxKey = 0
)

var (
	// ErrUserNotFoundInContext is the error returned when FromContext is unable
	// to find a user in the given context
	ErrUserNotFoundInContext = fmt.Errorf("user not found in context")
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
	Teams             map[string]string `json:"teams"`
}

// FromContext takes a given context in and tries to find the user stored in
// the context. It returns the user if it is present, or returns an error if it
// is not able to find it.
func FromContext(ctx context.Context) (*User, error) {
	u, ok := ctx.Value(userCtxKey).(*User)
	if !ok {
		return nil, ErrUserNotFoundInContext
	}

	return u, nil
}

// WithContext takes a context and returns a context containing the user object
func (u *User) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, userCtxKey, u)
}
