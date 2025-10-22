package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "itunes_api/models"
)

const (
	BASEURL = "https://itunes.apple.com/search?term=%s&media=music&limit=5"
)

type Client struct {
	BaseURL			string
	HTTPClient		*http.Client
}

type errorResponse struct {
	Error 			bool		`json:"error"`
	Reason			string		`json:"reason"`
}

func NewClient() *Client {
	return &Client {
		BaseURL: BASEURL,
		HTTPClient: &http.Client {
			Timeout: time.Second,
		},
	}
}


func (c *Client) Search(ctx context.Context, query string) (*models.ITunesResponse, error) {

	url := fmt.Sprintf(c.BaseURL, query)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var requestError errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&requestError); err != nil {
			return nil, fmt.Errorf("request api error: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("request api error: %s", requestError.Reason)
	}
	
	var ittunesItems models.ITunesResponse
	if err := json.NewDecoder(resp.Body).Decode(&ittunesItems); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &ittunesItems, nil
}