package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("== Cikk table query machine == ")

	db, err := ConnectDb()
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer db.Close()

	fmt.Println(" --- Get all Cikk ---")
	cikks, err := GetAllCikk(db)
	if err != nil {
		log.Println("Error: ", err)
	} else {
		for _, cikk := range cikks {
			fmt.Printf("CikkId: %d, Név: %s, Cikkszám: %s\n", cikk.ID, cikk.Nev, cikk.Cikkszam)
		}
	}
	fmt.Printf("Total: %d items \n\n", len(cikks))
}