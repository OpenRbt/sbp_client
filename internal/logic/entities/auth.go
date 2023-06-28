package entities

// Auth ...
type Auth struct {
	UID          string
	Disabled     bool
	UserMetadata *AuthUserMeta
}

// AuthUserMeta ...
type AuthUserMeta struct {
	CreationTimestamp    int64
	LastLogInTimestamp   int64
	LastRefreshTimestamp int64
}

// Auth ... for swagger
type AuthFunc func(token string) (*Auth, error)

// Token
type Token struct {
	Value string
}
