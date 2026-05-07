package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"os"
)

type Describer interface {
	Describe() string
}

type Product struct {
	Name		string		`json:"name"`
	Price		string		`json:"price"`
	Quantity	string		`json:"quantity"`
}

func (p Product) Describe() string {
	// Itt azért használunk Sprintf-et, mert ez stringet ad vissza de nem printel
	return fmt.Sprintf("%s | %.1f | qty: %d", p.Name, p.Price, p.Quantity)
}

func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	// szimulált API response
	jsonData := `[
		{"name":"Keyboard","price":12990,"quantity":3},
		{"name":"Mouse","price":7990,"quantity":5},
		{"name":"Monitor","price":64990,"quantity":2}
	]`

	// stringből csinálnuk egy readert
	// A decoder nem []byte-ból dolgozik hanem io.Readerból
	reader := strings.NewReader(jsonData)

	var products []Product

	// Decoder -> readerből olvas - JSON-ként értelmez - beletölti a products sliceba
	err := json.NewDecoder(reader).Decode(&products)
	if err != nil {
		log.Fatal("JSON decode error: ", err)
	}

	fmt.Println("Descriptions:")

	for _, product := range products {
		PrintDescription(product)
	}

	data, err := json.Marshal(products)
	if err != nil {
		log.Fatal("JSON marshal error: ", err)
	}

	fmt.Println()
	fmt.Println("Marshal result:")
	fmt.Println(string(data))

	var decoded []Product

	err = json.Unmarshal(data, &decoded)
	if err != nil {
		log.Fatal("JSON unmarshall error: ", err)
	}

	fmt.Println()
	fmt.Println("Encoder result:")

	err = json.NewEncoder(os.Stdout).Encode(decoded)
	if err != nil {
		log.Fatal("JSON encode error: ", err)
	}
}
