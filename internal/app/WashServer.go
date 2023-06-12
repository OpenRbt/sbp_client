package app

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID               uuid.UUID
	Title            string
	Description      string
	ServiceKey       string
	Owner            uuid.UUID
	TerminalKey      string
	TerminalPassword string
}

type RegisterWashServer struct {
	Title            string
	Description      string
	TerminalKey      string
	TerminalPassword string
}

type UpdateWashServer struct {
	ID               uuid.UUID
	Title            *string
	Description      *string
	TerminalKey      *string
	TerminalPassword *string
}
