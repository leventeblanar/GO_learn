package decodedencoded

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// HTTP GET lekérés -> Egyszerű API lekérés adott endpointról url alapján
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("HTTP GET Error: ", err)
	}
	defer resp.Body.Close()

	// Státuszkód ellenőrzés
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %d", resp.StatusCode)
	}

	var todo Todo

	// Decoder -> először kéri azt, hogy mit, aztán hogy hova. (honnan -> hova)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		log.Fatal("JSON decode error: ", err)
	}

	// Normál kiírás
	fmt.Printf("Todo #%d: %s | completed: %t\n", todo.ID, todo.Title, todo.Completed)

	// Encoder -> Todo structot kiírjuk JSON-ként az os.Stdout-ra
	fmt.Println("Encoded back to JSON:")

	// Először megmondjuk hova írjunk ki és mit (hova -> honnan)
	err = json.NewEncoder(os.Stdout).Encode(todo)
	if err != nil {
		log.Fatal("JSON encode error: ", err)
	}
}