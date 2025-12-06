package main

import (
	"log"

	"your/module/internal/api"
	"your/module/internal/config"
	"your/module/internal/store"
)

func main() {
	cfg := config.Load()

	db, err := store.NewPostgres(cfg)
	if err != nil {
		log.Fatal("DB error:", err)
	}

	r := api.NewRouter(db)

	r.Run(":8080")
}
