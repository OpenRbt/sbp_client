package entities

import "errors"

var (
	ErrForbidden    = errors.New("access denied")
	ErrNotFound     = errors.New("entity not found")
	ErrUserNotOwner = errors.New("user not owner")
	ErrNotification = errors.New("notification is not correct")
	ErrBadRequest   = errors.New("bad request")

	ErrBadVersion = errors.New("entity version the same or higher")
)
