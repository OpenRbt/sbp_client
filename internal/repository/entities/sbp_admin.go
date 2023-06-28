package repentities

import uuid "github.com/satori/go.uuid"

type SbpAdmin struct {
	ID       uuid.NullUUID `db:"id"`
	Identity string        `db:"identity"`
}
