package repentities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Wash ...
type Wash struct {
	ID               uuid.UUID `db:"id"`
	OwnerID          uuid.UUID `db:"owner_id"`
	Password         string    `db:"password"`
	Title            string    `db:"title"`
	Description      string    `db:"description"`
	TerminalKey      string    `db:"terminal_key"`
	TerminalPassword string    `db:"terminal_password"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

// RegisterWash ...
type RegisterWash struct {
	OwnerID          uuid.UUID `db:"owner_id"`
	Password         string    `db:"password"`
	Title            string    `db:"title"`
	Description      string    `db:"description"`
	TerminalKey      string    `db:"terminal_key"`
	TerminalPassword string    `db:"terminal_password"`
}

// UpdateWash ...
type UpdateWash struct {
	ID               uuid.NullUUID `db:"id"`
	Name             *string       `db:"name"`
	Description      *string       `db:"description"`
	TerminalKey      *string       `db:"terminal_key"`
	TerminalPassword *string       `db:"terminal_password"`
}

// DeleteWash ...
type DeleteWash struct {
	ID uuid.NullUUID `db:"id"`
}
