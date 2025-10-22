package main

import (
	"context"
	"fmt"
	"log"
	"time"

	apiclient "itunes_api/api"
)

func main() {
	client := apiclient.NewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	itunes_json, err := client.Search(ctx, "queen")
	if err != nil {
		log.Fatalf("unable to fetch itunes response: %s", err)
	}

	fmt.Println("Itunes response:")
	for i, item := range itunes_json.Results {
		fmt.Printf("%d. %s - %s (%s)\n", i, item.ArtistName, item.TrackName, item.CollectionName)
	}

}