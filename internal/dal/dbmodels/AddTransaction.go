package dbmodels

type AddTransaction struct {
	ServerID  string `db:"server_id"`
	PostID    string `db:"post_id"`
	Amount    int    `db:"amount"`
	PaymentID string `db:"payment_id_bank"`
}
