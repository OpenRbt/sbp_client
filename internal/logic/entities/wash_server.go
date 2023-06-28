package entities

import (
	uuid "github.com/satori/go.uuid"
)

// WashServer ...
type WashServer struct {
	ID               uuid.UUID
	Title            string
	Description      string
	ServiceKey       string
	Owner            uuid.UUID
	TerminalKey      string
	TerminalPassword string
}

// RegisterWashServer ...
type RegisterWashServer struct {
	Title            string
	Description      string
	TerminalKey      string
	TerminalPassword string
}

// UpdateWashServer ...
type UpdateWashServer struct {
	ID               uuid.UUID
	Title            *string
	Description      *string
	TerminalKey      *string
	TerminalPassword *string
}
