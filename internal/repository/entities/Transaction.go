package repentities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID         uuid.NullUUID `db:"id"`
	ServerID   string        `db:"server_id"`
	PostID     string        `db:"post_id"`
	Amount     int64         `db:"amount"`
	PaymentID  string        `db:"payment_id_bank"`
	Status     string        `db:"status"`
	DataCreate time.Time     `db:"data_create"`
}

type TransactionCreate struct {
	ID        uuid.UUID `db:"id"`
	ServerID  string    `db:"server_id"`
	PostID    string    `db:"post_id"`
	Amount    int64     `db:"amount"`
	PaymentID string    `db:"payment_id_bank"`
	Status    string    `db:"status"`
}
