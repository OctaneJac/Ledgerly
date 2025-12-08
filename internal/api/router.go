package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/octanejac/Ledgerly/internal/store"
)

func NewRouter(db *store.PostgresStore) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Post("/accounts", CreateAccountHandler(db))
	r.Post("/transactions", CreateTransactionHandler(db))

	return r
}
