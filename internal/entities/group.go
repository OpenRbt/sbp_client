package entities

import uuid "github.com/satori/go.uuid"

type Group struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	Description    string
	IsDefault      bool
	Deleted        bool
	Version        int
}
