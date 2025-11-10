package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf(".env is not found in project root: %w", err)
	}

	connStr := os.Getenv("ATLAS_DB")
	if connStr == "" {
		return nil, fmt.Errorf("missing ATLAS_DB connection string in .env")
	}

	if !strings.Contains(connStr, "?sslmode=disable") {
		connStr += "?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error db connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}
