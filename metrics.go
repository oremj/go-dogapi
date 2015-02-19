package dogapi

// A Metric represents the metrics portion on a series POST
type Metric struct {
	Host   string   `json:"host"`
	Metric string   `json:"metric"`
	Points [][]int  `json:"points"`
	Tags   []string `json:"tags"`
	Type   string   `json:"type"`
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
