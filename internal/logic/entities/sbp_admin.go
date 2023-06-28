package entities

import uuid "github.com/satori/go.uuid"

// SbpAdmin ...
type SbpAdmin struct {
	ID       uuid.UUID
	Identity string
}
