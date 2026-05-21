package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)


type WeatherResponse struct {
	Latitude		float64					`json:"latitude"`
	Longitude		float64					`json:"longitude"`
	Timezone		string					`json:"timezone"`
	CurrentUnits	CurrentWeatherUnits		`json:"current_units"`
	Current			CurrentWeather			`json:"current"`
}

type CurrentWeather struct {
	Time			string					`json:"time"`
	Interval		int						`json:"interval"`
	Temperature2M	float64					`json:"temperature_2m"`
	WindSpeed10M	float64					`json:"wind_speed_10m"`
	WeatherCode 	int						`json:"weather_code"`
}

type CurrentWeatherUnits struct {
	Time			string					`json:"time"`
	Interval		string					`json:"interval"`
	Temperature2M	string					`json:"temperature_2m"`
	WindSpeed10M	string					`json:"wind_speed_10m"`
	WeatherCode		string					`json:"weather_code"`
}


func (c *Client) GetCurrentWeather(ctx context.Context, lat float64, lon float64) (WeatherResponse, error) {
	endpoint := strings.TrimRight(c.baseURL, "/") + "/v1/forecast"

	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("parse endpoint url: %w", err)
	}

	query := parsedURL.Query()

	query.Set("latitude", strconv.FormatFloat(lat, 'f', -1, 64))
	query.Set("longitude", strconv.FormatFloat(lon, 'f', -1, 64))
	query.Set("current", "temperature_2m,wind_speed_10m,weather_code")
	query.Set("timezone", "auto")

	parsedURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var weather WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("decode weather response: %w", err)
	}

	return weather, nil
}