package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	BASE_URL = "https://jsonplaceholder.typicode.com"
)

type Client struct {
	BaseURL		string
	apiKey		string
	HTTPClient	*http.Client
}

type errorResponse struct {
	Error	bool	`json:"error"`
	Reason	string	`json:"reason"`
}


// got this by "curl "https://jsonplaceholder.typicode.com/todo""
//   {
//     "userId": 6,
//     "id": 120,
//     "title": "dolorem laboriosam vel voluptas et aliquam quasi",
//     "completed": false
//   },

type JSONResponse struct {
	UserId		int		`json:"id"`
	TodoId		int		`json:"todoid"`
	Title		string	`json:"title"`
	Completed	bool	`json:"completed"`
}


func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BASE_URL,
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetResponsefromApi(ctx context.Context) ([]JSONResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/todos", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer"+c.apiKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var apiErr errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("jsonresponse api error, status: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("jsonresponse api error: %w", apiErr.Reason)
	}

	var todos []JSONResponse
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return todos, nil
}

