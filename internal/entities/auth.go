package entities

import uuid "github.com/satori/go.uuid"

type Auth struct {
	User         User
	UserMetadata *UserAuthMetadata
	Disabled     bool
}

type UserAuthMetadata struct {
	CreationTimestamp    int64
	LastLogInTimestamp   int64
	LastRefreshTimestamp int64
}

func (auth *Auth) IsSystemManager() bool {
	return auth.User.Role == SystemManagerRole
}

func (auth *Auth) IsAdmin() bool {
	return auth.User.Role == AdminRole
}

func (auth *Auth) IsAdminManageOrganization(organizationID uuid.UUID) bool {
	return auth.IsAdmin() && auth.User.OrganizationID != nil && *auth.User.OrganizationID == organizationID
}
