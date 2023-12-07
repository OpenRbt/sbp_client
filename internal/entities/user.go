package entities

import uuid "github.com/satori/go.uuid"

type User struct {
	ID             string
	Email          string
	Name           string
	Role           Role
	OrganizationID *uuid.UUID
	Version        int
	Deleted        bool
}

type Role string

const (
	SystemManagerRole Role = "system_manager"
	AdminRole         Role = "admin"
	NoAccessRole      Role = "no_access"
)
