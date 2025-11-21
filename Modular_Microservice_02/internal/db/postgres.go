package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf(".env is not found in project root: %w", err)
	}

	connString := os.Getenv("CHINOOK_DB")
	if connString == "" {
		return nil, fmt.Errorf("")
	}

	if !strings.Contains(connString, "?sslmode=disable") {
		connString+="?sslmode=disable"
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error db connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}
