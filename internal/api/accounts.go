package api

import (
	"encoding/json"
	"net/http"

	"github.com/octanejac/Ledgerly/internal/store"
)

type CreateAccountRequest struct {
	Name string `json:"name"`
}

func CreateAccountHandler(db *store.PostgresStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateAccountRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		_, err := db.DB.Exec(`INSERT INTO accounts (name) VALUES ($1)`, req.Name)
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("account created"))
	}
}
