package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"api_lek/api"
)

func main() {
	lat := flag.Float64("lat", 47.4979, "latitude")
	lon := flag.Float64("lon", 19.0402, "latitude")

	flag.Parse()

	cfg, err := api.LoadConfig()
	if err != nil {
		log.Fatal("config error: ", err)
	}

	client := api.NewClient(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	weather, err := client.GetCurrentWeather(ctx, *lat, *lon)
	if err != nil {
		log.Fatal("get current weather error: ", err)
	}

	fmt.Println("Current weather")
	fmt.Println("----------------")
	fmt.Printf("Location: %.4f, %.4f\n", weather.Latitude, weather.Longitude)
	fmt.Printf("Timezone: %s\n", weather.Timezone)
	fmt.Printf("Time: %s\n", weather.Current.Time)
	fmt.Printf(
		"Temperature: %.1f %s\n",
		weather.Current.Temperature2M,
		weather.CurrentUnits.Temperature2M,
	)
	fmt.Printf(
		"Wind speed: %.1f %s\n",
		weather.Current.WindSpeed10M,
		weather.CurrentUnits.WindSpeed10M,
	)
	fmt.Printf(
		"Weather code: %d %s\n",
		weather.Current.WeatherCode,
		weather.CurrentUnits.WeatherCode,
	)
}