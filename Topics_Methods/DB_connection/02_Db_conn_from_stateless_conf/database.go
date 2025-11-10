package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_"github.com/lib/pq"
	"github.com/joho/godotenv"
)

// DB CONNECTION - stateless függvény

func ConnectDb() (*sql.DB, error) {
	// .env betöltés
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env: %w", err)
	}

	// connection string
	connStr := os.Getenv("ATLAS_DB")
	if connStr == "" {
		return nil, fmt.Errorf("ATLAS_DB environment variable is not set")
	}

	if !strings.Contains(connStr, "?sslmode=disabled") {
		connStr+="?sslmode=disable"
	}

	// DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	log.Println("Database connected")
	return db, nil
}


func GetAllCikkTipus(db *sql.DB) ([]CikkTipus, error) {
	query := `
	SELECT id, version_num, nev, kod
	FROM atlas.cikk_tipus
	ORDER BY id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var items []CikkTipus
	for rows.Next(){
		var item CikkTipus
		err := rows.Scan(&item.ID, &item.VersionNum, &item.Nev, &item.Kod)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return items, nil
}

// GetCikkTipusByID - egy cikk típus lekérése ID alapján
func GetCikkTipusById(db *sql.DB, id int) (*CikkTipus, error) {
	query := `
	SELECT id, version_num, nev, kod
	FROM atlas.cikk_tipus
	WHERE id = $1
	`

	var item CikkTipus
	err := db.QueryRow(query, id).Scan(&item.ID, &item.VersionNum, &item.Nev, &item.Kod)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("cikk tipus not found with ID: %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &item, nil
}