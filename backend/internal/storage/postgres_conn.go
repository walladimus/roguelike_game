package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// creates a *sql.DB using the ROGUE_DB_DSN env var.
func NewPostgresDBFromEnv() (*sql.DB, error) {
	dsn := os.Getenv("ROGUE_DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("ROGUE_DB_DSN environment variable is not set")
	}

	//pgx stdlib driver
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	//immediate connection check
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
