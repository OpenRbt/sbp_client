package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Wash ...
type Wash struct {
	ID               uuid.UUID
	OwnerID          uuid.UUID
	Password         string
	Title            string
	Description      string
	TerminalKey      string
	TerminalPassword string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// RegisterWash ...
type RegisterWash struct {
	ID               uuid.UUID
	OwnerID          uuid.UUID
	Password         string
	Title            string
	Description      string
	TerminalKey      string
	TerminalPassword string
}

// UpdateWash ...
type UpdateWash struct {
	ID               uuid.UUID
	OwnerID          uuid.UUID
	Title            *string
	Description      *string
	TerminalKey      *string
	TerminalPassword *string
}
