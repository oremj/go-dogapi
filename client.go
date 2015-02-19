package dogapi

import "net/http"

// APIEndpoint is the default datadog api endpoint
var APIEndpoint = "https://app.datadoghq.com/api"

// A Client is a Datadog client
type Client struct {
	APIKey         string
	ApplicationKey string

	BaseURL string

	Client *http.Client
}

// NewClient returns new datadog client given an apiKey and appKey
//
// appKey will be ignored if set to ""
func NewClient(apiKey, appKey string) *Client {
	return &Client{
		APIKey:         apiKey,
		ApplicationKey: appKey,
		Client:         &http.Client{},
		BaseURL:        APIEndpoint,
	}
}
