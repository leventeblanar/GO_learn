package main

import (
	"log"

	"github.com/joho/godotenv"
	database "db_connection/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars.")
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Db Conencted!")
}