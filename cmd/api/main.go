package main

import (
	"log"
	"net/http"

	"github.com/octanejac/Ledgerly/internal/api"
	"github.com/octanejac/Ledgerly/internal/config"
	"github.com/octanejac/Ledgerly/internal/store"
)

func main() {
	cfg := config.Load()

	db, err := store.NewPostgres(cfg)
	if err != nil {
		log.Fatal("DB error:", err)
	}
	r := api.NewRouter(db)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error:", err)
	}
}

