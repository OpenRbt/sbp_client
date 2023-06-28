package repentities

import uuid "github.com/satori/go.uuid"

// WashServer ...
type WashServer struct {
	ID               uuid.NullUUID `db:"id"`
	Title            string        `db:"title"`
	Description      string        `db:"description"`
	Owner            uuid.NullUUID `db:"owner"`
	ServiceKey       string        `db:"service_key"`
	TerminalKey      string        `db:"terminal_key"`
	TerminalPassword string        `db:"terminal_password"`
}

// AddWashServer ...
type AddWashServer struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}

// RegisterWashServer ...
type RegisterWashServer struct {
	Title            string        `db:"title"`
	Description      string        `db:"description"`
	Owner            uuid.NullUUID `db:"owner"`
	ServiceKey       string        `db:"service_key"`
	TerminalKey      string        `db:"terminal_key"`
	TerminalPassword string        `db:"terminal_password"`
}

// UpdateWashServer ...
type UpdateWashServer struct {
	ID               uuid.NullUUID `db:"id"`
	Name             *string       `db:"name"`
	Description      *string       `db:"description"`
	TerminalKey      *string       `db:"terminal_key"`
	TerminalPassword *string       `db:"terminal_password"`
}

// DeleteWashServer ...
type DeleteWashServer struct {
	ID uuid.NullUUID `db:"id"`
}
