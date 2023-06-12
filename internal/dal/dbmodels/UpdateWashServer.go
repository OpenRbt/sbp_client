package dbmodels

import uuid "github.com/satori/go.uuid"

type UpdateWashServer struct {
	ID               uuid.NullUUID `db:"id"`
	Name             *string       `db:"name"`
	Description      *string       `db:"description"`
	TerminalKey      *string       `db:"terminal_key"`
	TerminalPassword *string       `db:"terminal_password"`
}
