package dogapi

// A Metric represents the metrics portion on a series POST
type Metric struct {
	Metric string  `json:"metric"`
	Points [][]int `json:"points"`

	Host string   `json:"host,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Type string   `json:"type,omitempty"`
}

// A SeriesRequest is a time series request
type SeriesRequest struct {
	Series []Metric `json:"series"`
}

// A SeriesResponse is a time series response
type SeriesResponse struct {
	Status string `json:"status"`
}

// PostSeries posts time series data
func (c *Client) PostSeries(req *SeriesRequest) (*SeriesResponse, error) {
	resp := new(SeriesResponse)
	err := c.makeJSONRequest("POST", c.BaseURL+"/v1/series", req, resp)
	return resp, err
}
