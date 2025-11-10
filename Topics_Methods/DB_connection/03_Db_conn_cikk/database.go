package main


import (
	"fmt"
	"database/sql"
	"os"
	"log"
	"strings"

	_"github.com/lib/pq"
	"github.com/joho/godotenv"
)

func ConnectDb()(*sql.DB, error){
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("there is no .env file on the given location: %w", err)
	}

	connStr := os.Getenv("ATLAS_DB")
	if connStr == "" {
		return nil, fmt.Errorf("missing db config string")
	}

	if !strings.Contains(connStr, "?sslmode=disable") {
		connStr += "?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("database connection unsuccessful: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("cannot ping database, %w", err)
	}

	log.Println("Database connected")
	return db, nil
}

func GetAllCikk(db *sql.DB) ([]Cikk, error) {
	query := `
	SELECT Id, Nev, Cikkszam FROM atlas.cikk
	LIMIT 15;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var cikks []Cikk
	for rows.Next() {
		var item Cikk
		err := rows.Scan(&item.ID, &item.Nev, &item.Cikkszam)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		cikks = append(cikks, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return cikks, nil
}