package api

import (
	"encoding/json"
	"net/http"

	"github.com/octanejac/Ledgerly/internal/store"
)

type CreateTransactionRequest struct {
	AccountID int     `json:"account_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"` // credit/debit
}

func CreateTransactionHandler(db *store.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateTransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		_, err := db.DB.Exec(`
			INSERT INTO transactions (account_id, amount, type)
			VALUES ($1, $2, $3)`,
			req.AccountID, req.Amount, req.Type,
		)
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("transaction created"))
	}
}
