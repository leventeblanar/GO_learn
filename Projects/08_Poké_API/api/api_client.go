package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	BASEURL_POKE = "https://pokeapi.co/api/v2/pokemon/pikachu"
)

type Client struct {
	BaseURL			string
	apiKey			string
	HTTPClient		*http.Client
}

type errorResponse struct {
	Error		bool	`json:"error"`
	Reason		string 	`json:"reason"`
}

type PokeResponse struct {
	Name		string		`json:"name"`
	Height		int			`json:"height"`
	Weight		int			`json:"weight"`
	Types		[]struct {
		Type struct {
			Name string		`json:"name"`
		}	`json:"type"`
	} `json:"types"`
	Sprites	struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}


func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BASEURL_POKE,
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}


func (c *Client) GetPokeData(ctx context.Context) (*PokeResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var apiErr errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("weather api error: status%d", resp.StatusCode)
		}
		return nil, fmt.Errorf("weather api error: %s", apiErr.Reason)
	}

	var pokemon PokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &pokemon, nil
}