package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("== Stateless DB Connector Demo ===\n")

	db, err := ConnectDb()
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer db.Close()

	fmt.Println(" --- Get All Cikk Tipus ---")
	items, err := GetAllCikkTipus(db)
	if err != nil {
		log.Println("Error: ", err)
	} else {
		for _, item := range items {
			fmt.Printf("ID: %d, Version: %d, Név: %s, Kód: %s\n", item.ID, item.VersionNum, item.Nev, item.Kod)
		}
		fmt.Printf("Total: %d items \n\n", len(items))
	}

	fmt.Println("--- Get Cikk Tipus by ID ---")
	item, err := GetCikkTipusById(db, 1000369)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Printf("Found: %+v\n\n", item)
	}

}