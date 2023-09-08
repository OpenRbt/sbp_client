package repentities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID        uuid.UUID `db:"id"`
	WashID    string    `db:"wash_id"`
	PostID    string    `db:"post_id"`
	Amount    int64     `db:"amount"`
	PaymentID string    `db:"payment_id_bank"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TransactionCreate struct {
	ID        uuid.UUID `db:"id"`
	WashID    string    `db:"wash_id"`
	PostID    string    `db:"post_id"`
	Amount    int64     `db:"amount"`
	PaymentID string    `db:"payment_id_bank"`
	Status    string    `db:"status"`
}
