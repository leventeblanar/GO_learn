package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func main() {
	// initialize product using Product struct
	product := Product {
		Name:	"Keyboard",
		Price:	12990.0,
		Quantity: 3,
	}

	// =================================
	// Marshal: GO struct -> JSON []byte
	// =================================
	data, err := json.Marshal(product) // -> Marshal paraméterként csak azt várja, hogy mit fog szétbontani, acél formátumot tudja
	if err != nil {
		log.Fatal("JSON marshal error: ", err)
	}

	fmt.Println("OG form -> []byte:")
	fmt.Println(data)

	fmt.Println(("JSON:"))
	fmt.Println(string(data))


	// =================================
	// Unmarshal: JSON []byte -> GO Struct
	// =================================
	var decoded Product

	err = json.Unmarshal(data, &decoded) // Az Unmarshal paraméterként elsőre várja a json-t, majd a struct pointerét ahova rakni kell
	if err != nil {
		log.Fatal("JSON unmarshall error ", err)
	}

	fmt.Printf(
		"Decoded product: %s | %.1f Ft | qty: %d\n",
		decoded.Name,
		decoded.Price,
		decoded.Quantity,
	)
}