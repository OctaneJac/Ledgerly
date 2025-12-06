package store

import (
	"database/sql"
	"fmt"

	"your/module/internal/config"

	_ "github.com/lib/pq"
)

func NewPostgres(c config.Config) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName,
	)

	return sql.Open("postgres", conn)
}
