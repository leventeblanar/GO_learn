package api

import (
	"net/http"
)

type Client struct {
	baseURL 		string
	httpClient		*http.Client
}


func NewClient(cfg Config) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}