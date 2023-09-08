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
