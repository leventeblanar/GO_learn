package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	BASEURL_WF = "https://api.open-meteo.com/v1/forecast?latitude=47.4979&longitude=19.0402&hourly=temperature_2m"
)

// struct for the client containing BASEURL, apiKEY(in this case it is not needed), and the HTTP modul's client
type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

// the error response should have the Error (if there is or not, true vs false), and a Reason
type errorResponse struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}

// Successfull response has the data: latitude, longitude, hourly strict has time and temperature 
type ForecastResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}


func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BASEURL_WF,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// GetHourlyForecast fetches hourly temperature data for the configured coordinates
// this function uses the client struct - receives the context module context and returns the forecastresponse
func (c *Client) GetHourlyForecast(ctx context.Context) (*ForecastResponse, error) {
	//  determine the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// set authorization
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	//  run the request, store the response in resp
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()  // close Body

	// API response -> error, decode the error
	if resp.StatusCode >= http.StatusBadRequest {
		var apiErr errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("weather api error: status %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("weather api error: %s", apiErr.Reason)
	}

	// if successful, decode for ForcastResponse
	var forecast ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &forecast, nil
}
