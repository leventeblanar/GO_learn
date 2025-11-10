package main

import (
	"fmt"
	"log"

	"modular_microservice/internal/db"
	"modular_microservice/internal/repository"
)

func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
	defer conn.Close()

	fmt.Println("Database connected")

	cikkRepo := repository.NewCikkRepository(conn)

	cikkek, err := cikkRepo.GetAllCikk()
	if err != nil {
		log.Fatal("Failed to get cikk:", err)
	}

	fmt.Printf("\n=== Found %d cikk ===\n", len(cikkek))
	for i, c := range cikkek {
		fmt.Printf("%d. ID: %d, Név: %s, Cikkszám: %s, CikkTípusID: %d\n", i+1, c.ID, c.Nev, c.Cikkszam, c.CikkTipusId)
	}
}
