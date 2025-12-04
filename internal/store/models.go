package store

import "time"

type AccountRow struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type TransactionRow struct {
	ID          string    `db:"id"`
	AccountID   string    `db:"account_id"`
	Amount      int64     `db:"amount"`
	Type        string    `db:"type"` // "debit" or "credit"
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}
