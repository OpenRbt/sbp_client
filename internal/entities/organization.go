package entities

import uuid "github.com/satori/go.uuid"

type Organization struct {
	ID          uuid.UUID
	Name        string
	Description string
	DisplayName string
	IsDefault   bool
	Deleted     bool
	Version     int
}

type SimpleOrganization struct {
	ID      uuid.UUID
	Name    string
	Deleted bool
}
