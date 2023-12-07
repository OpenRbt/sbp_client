package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Wash struct {
	ID               uuid.UUID
	Owner            string
	Password         string
	Title            string
	Description      string
	TerminalKey      string
	TerminalPassword string
	OrganizationID   uuid.NullUUID
	GroupID          uuid.NullUUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type WashCreation struct {
	OwnerID          string
	Password         string
	Title            string
	Description      string
	TerminalKey      string
	TerminalPassword string
	GroupID          uuid.UUID
}

type WashUpdate struct {
	Title            *string
	Description      *string
	TerminalKey      *string
	TerminalPassword *string
}

type WashFilter struct {
	Pagination
	GroupID        *uuid.UUID
	OrganizationID *uuid.UUID
}
