package main

import (
	"log"
	"net/http"

	"github.com/octanejac/Ledgerly/internal/api"
	"github.com/octanejac/Ledgerly/internal/config"
	"github.com/octanejac/Ledgerly/internal/store"
)

func main() {
	// Load environment/config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to DB
	db, err := store.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Setup router (pass dependencies if needed)
	router := api.NewRouter(db)

	log.Printf("API running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
