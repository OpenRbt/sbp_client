package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID        uuid.UUID
	Identity  string
	Role      userRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

// IsAdmin ...
func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}

// IsUser...
func (u *User) IsUser() bool {
	return u.Role == UserRoleUser
}
