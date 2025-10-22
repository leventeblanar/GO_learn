package main

import (
	"context"
	"fmt"
	"log"
	"time"

	apiclient "jsonresponse/api"
)

func main() {
	client := apiclient.NewClient("")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todos, err := client.GetResponsefromApi(ctx)
	if err != nil {
		log.Fatalf("unable to fetch jsonresponse: %s", err)
	}

	const maxEntries = 12

	fmt.Println("Get request results:")
	for i, todo := range todos {
		if i > maxEntries {
			break
		}
		fmt.Printf("%d: %s (completed %t)\n", todo.TodoId, todo.Title, todo.Completed)
	}
}