package dogapi

import "net/http"

var APIEndpoint = "https://app.datadoghq.com/api"

type Client struct {
	APIKey         string
	ApplicationKey string

	BaseURL string

	Client *http.Client
}

func NewClient(apiKey, appKey string) *Client {
	return &Client{
		APIKey:         apiKey,
		ApplicationKey: appKey,
		Client:         &http.Client{},
		BaseURL:        APIEndpoint,
	}
}
