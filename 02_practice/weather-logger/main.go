package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/leventeblanar/GO_learn/weather-logger/api"
	"github.com/leventeblanar/GO_learn/weather-logger/db"
	"github.com/leventeblanar/GO_learn/weather-logger/server"
)

func main() {
	// API FETCH
	weather, err := api.FetchWeather()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Weather currently fetched from API:\n")
	fmt.Printf("Temp: %.1f°C, Wind: %.1f km/h, Code: %d\n",
		weather.Temperature, weather.Windspeed, weather.Weathercode)


	// DB CONN INIT
	database, err := db.InitDB("postgresql://levente_blanar:brutal.shred01@localhost:5432/hermes_sync_local?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	
	// INSERT
	err = db.InsertReading(database, weather)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved to DB!")

	// GET CURRENT
	readings, err := db.GetReadings(database, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(readings)

	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			weather, err := api.FetchWeather()
			if err != nil {
				log.Println("Fetch error:", err)
				continue
			}
			err = db.InsertReading(database, weather)
			if err != nil {
				log.Println("Insert error", err)
			}
			fmt.Println("Auto-saved reading")
		}
	}()


	// HTTP SERVER
	http.HandleFunc("/readings", server.ReadingsHandler(database))
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}