package entities

import uuid "github.com/satori/go.uuid"

type User struct {
	ID             string
	Email          string
	Name           string
	Role           UserRole
	OrganizationID *uuid.UUID
	Version        int
	Deleted        bool
}

type UserRole string

const (
	SystemManagerRole UserRole = "system_manager"
	AdminRole         UserRole = "admin"
	NoAccessRole      UserRole = "no_access"
)
