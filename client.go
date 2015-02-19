package dogapi

import "net/http"

var ApiEndpoint = "https://app.datadoghq.com/api"

type Client struct {
	ApiKey         string
	ApplicationKey string

	BaseURL string

	Client *http.Client
}

func NewClient(apiKey, appKey string) *Client {
	return &Client{
		ApiKey:         apiKey,
		ApplicationKey: appKey,
		Client:         &http.Client{},
		BaseURL:        ApiEndpoint,
	}
}
