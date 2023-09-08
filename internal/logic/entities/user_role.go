package entities

// userRole ...
type userRole string

const (
	UserRoleUser    userRole = "user"
	UserRoleAdmin   userRole = "admin"
	UserRoleDefault userRole = UserRoleUser
)

// validUserRoles ...
var validUserRoles map[userRole]bool = map[userRole]bool{
	UserRoleUser:  true,
	UserRoleAdmin: true,
}

// IsUserRoleValid ...
func IsUserRoleValid(r userRole) bool {
	return validUserRoles[r]
}

// String ...
func (u userRole) String() string {
	return string(u)
}

// UserRoleFromString ...
func UserRoleFromString(s string) userRole {
	role := userRole(s)
	ok := validUserRoles[role]
	if !ok {
		role = UserRoleDefault
	}
	return role
}
