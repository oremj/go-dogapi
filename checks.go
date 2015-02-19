package dogapi

// PostCheckRunRequest
type PostCheckRunRequest struct {
	Check     string `json:"check"`
	HostName  string `json:"host_name"`
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message,omitempty"`
	Tags      string `json:"tags,omitempty"`
}

// PostCheckRunResponse
type PostCheckRunResponse struct {
	Status string `json:"status"`
}

// PostCheckRun posts check statuses for monitors
func (c *Client) PostCheckRun(req *PostCheckRunRequest) (*PostCheckRunResponse, error) {
	resp := new(PostCheckRunResponse)
	err := c.makeJSONRequest("POST", c.BaseURL+"/v1/check_run", req, resp)
	return resp, err
}
