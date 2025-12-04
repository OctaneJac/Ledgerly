package api
package ledger

import "time"

type Account struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func (a *Account) Validate() error {
	// Add validation rules later
	return nil
}
