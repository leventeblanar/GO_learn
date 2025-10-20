package main

import (
	"context"
	"fmt"
	"log"
	"time"

	apiclient "weather_api/api_client"
)

func main() {
	// call client
	client := apiclient.NewClient("")

	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	forecast, err := client.GetHourlyForecast(ctx)
	if err != nil {
		log.Fatalf("unable to fetch weather data: %v", err)
	}

	const maxEntries = 12
	const inputLayout = "2006-01-02T15:04"

	fmt.Println("Budapest hourly temperature (°C)")
	for i := 0; i < len(forecast.Hourly.Time) && i < len(forecast.Hourly.Temperature2m) && i < maxEntries; i++ {
		timestamp := forecast.Hourly.Time[i]
		parsedTime, err := time.Parse(inputLayout, timestamp)
		if err != nil {
			fmt.Printf("%s -> %.1f°C\n", timestamp, forecast.Hourly.Temperature2m[i])
			continue
		}
		fmt.Printf("%s -> %.1f°C\n", parsedTime.Format("Jan 02 15:04"), forecast.Hourly.Temperature2m[i])
	}
}
