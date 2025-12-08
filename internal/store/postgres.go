package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/octanejac/Ledgerly/internal/config"
)

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgres(cfg config.Config) (*PostgresStore, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{DB: db}, nil
}
