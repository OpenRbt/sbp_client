package models

import (
	"sbp/internal/entities"
	"time"

	uuid "github.com/satori/go.uuid"
)

type SimpleGroup struct {
	ID      uuid.UUID `db:"group_id"`
	Name    string    `db:"group_name"`
	Deleted bool      `db:"group_deleted"`
}

type SimpleWash struct {
	ID      uuid.UUID `db:"wash_id"`
	Title   string    `db:"wash_title"`
	Deleted bool      `db:"wash_deleted"`
}

type SimpleOrganization struct {
	ID      uuid.UUID `db:"org_id"`
	Name    string    `db:"org_name"`
	Deleted bool      `db:"org_deleted"`
}

type Transaction struct {
	ID           uuid.UUID                  `db:"id"`
	PostID       int64                      `db:"post_id"`
	Amount       int64                      `db:"amount"`
	Status       entities.TransactionStatus `db:"status"`
	CreatedAt    time.Time                  `db:"created_at"`
	Wash         SimpleWash
	Group        SimpleGroup
	Organization SimpleOrganization
}
