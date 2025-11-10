package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type CikkTipus struct {
	Id			int
	Version_num	int
	Nev			string
	Kod			string
}

func main() {
	// .env betöltés
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connection string kiolvasás
	connStr := os.Getenv("ATLAS_DB")
	fmt.Println("Connection string:", connStr)

	// Ha nincs benne az sllmode, adjuk hozzá
	if !strings.Contains(connStr, "sslmode") {
		connStr += "?sslmode=disable"
	}

	// connection létrehozás
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot open database:", err)
	}
	defer db.Close()

	// teszt
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to database!")


	// =========== SQL Query és lekérés =====================

	query := `
	SELECT id, version_num, nev, kod
	FROM atlas.cikk_tipus
	ORDER BY id
	`
	// query futtatás
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

	// eredmény beolvasás
	fmt.Println("\n=== Cikk tipusok ===")
	var cikkTipusok []CikkTipus

	for rows.Next() {
		var ct CikkTipus
		err := rows.Scan(&ct.Id, &ct.Version_num, &ct.Nev, &ct.Kod)
		if err != nil {
			log.Fatal("Scan failed:", err)
		}
		cikkTipusok = append(cikkTipusok, ct)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Rows error: ", err)
	}

	//Kiírás
	for _, ct := range cikkTipusok {
		fmt.Printf("ID: %d, Version: %d, Név: %s, Kód: %s\n",
	ct.Id, ct.Version_num, ct.Nev, ct.Kod)
	}

	fmt.Printf("\nÖsszesen %d cikk típus\n", len(cikkTipusok))
}

