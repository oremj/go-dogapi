package dogapi

type Metric struct {
	Host   string   `json:"host"`
	Metric string   `json:"metric"`
	Points [][]int  `json:"points"`
	Tags   []string `json:"tags"`
	Type   string   `json:"type"`
}

type SeriesRequest struct {
	Series []Metric `json:"series"`
}

type SeriesResponse struct {
	Status string `json:"status"`
}

func (c *Client) PostSeries(req *SeriesRequest) (*SeriesResponse, error) {
	resp := new(SeriesResponse)
	err := c.makeJSONRequest("POST", c.BaseURL+"/v1/series", req, resp)
	return resp, err
}
