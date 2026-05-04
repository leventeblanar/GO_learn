package api

import (
	"encoding/json"
	"net/http"
)

// ===== STRUCTS =====

type WeatherResponse struct {
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	Windspeed 	float64 `json:"windspeed"`
	Weathercode	int		`json:"weathercode"`
}

// ===== FUNCTIONS =====


func FetchWeather() (CurrentWeather, error) {

	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=47.4979&longitude=19.0402&current_weather=true", )
	if err != nil {
		return CurrentWeather{}, err
	}
	defer resp.Body.Close()

	var result WeatherResponse
	// decoder object -> 
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return CurrentWeather{}, err
	}

	return result.CurrentWeather, nil
}