package repentities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID        uuid.NullUUID `db:"id"`
	Identity  string        `db:"identity"`
	Role      string        `db:"role"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
}
