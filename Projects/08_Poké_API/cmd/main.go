package main

import (
	"context"
	"fmt"
	"log"
	"time"

	apiclient "poke_api/api"
)

func main() {

	client := apiclient.NewClient("")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pokemon, err := client.GetPokeData(ctx) 
	if err != nil {
		log.Fatalf("unable to fetch pokemon data: %v", err)
	}

	fmt.Println("Pokemon data:")
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nType: %s\nFront sprite: %s\n",
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
		pokemon.Types[0].Type.Name,
		pokemon.Sprites.FrontDefault,
	)
}