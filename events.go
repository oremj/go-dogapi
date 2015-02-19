package dogapi

import "net/url"

// PostEventsRequest contains args for PostEvents
type PostEventsRequest struct {
	AlertType string   `json:"alert_type"`
	Priority  string   `json:"priority"`
	Tags      []string `json:"tags"`
	Text      string   `json:"text"`
	Title     string   `json:"title"`
}

// PostEventsResponse is the result of PostEvents
type PostEventsResponse struct {
	Event struct {
		DateHappened   int         `json:"date_happened"`
		Handle         interface{} `json:"handle"`
		ID             int         `json:"id"`
		Priority       string      `json:"priority"`
		RelatedEventID interface{} `json:"related_event_id"`
		Tags           []string    `json:"tags"`
		Text           interface{} `json:"text"`
		Title          string      `json:"title"`
		URL            string      `json:"url"`
	} `json:"event"`
	Status string `json:"status"`
}

// PostEvents posts events to the event stream
func (c *Client) PostEvents(req *PostEventsRequest) (*PostEventsResponse, error) {
	resp := new(PostEventsResponse)
	err := c.makeJSONRequest("POST", c.BaseURL+"/v1/events", req, resp)
	return resp, err
}

// Event represents a datastream event
type Event struct {
	AlertType    string   `json:"alert_type"`
	DateHappened int      `json:"date_happened"`
	DeviceName   string   `json:"device_name"`
	Host         string   `json:"host"`
	ID           int      `json:"id"`
	Payload      string   `json:"payload"`
	Priority     string   `json:"priority"`
	Resource     string   `json:"resource"`
	Tags         []string `json:"tags"`
	Text         string   `json:"text"`
	Title        string   `json:"title"`
	URL          string   `json:"url"`
}

// GetEventResponse is the response of GetEvent
type GetEventResponse struct {
	Event Event `json:"event"`
}

// GetEvent queries an event for details
func (c *Client) GetEvent(eventId string) (*GetEventResponse, error) {
	resp := new(GetEventResponse)
	err := c.makeJSONRequest("GET", c.BaseURL+"/v1/events/"+eventId, nil, resp)
	return resp, err
}

// GetEventsResponse contains a list of GetEventResponse
type GetEventsResponse struct {
	Events []Event `json:"events"`
}

// GetEvents queries the event stream
//
// priority sources and tags are optional
func (c *Client) GetEvents(start, end, priority, sources, tags string) (*GetEventsResponse, error) {
	resp := new(GetEventsResponse)

	qs := url.Values{}
	qs.Add("start", start)
	qs.Add("end", end)
	if priority != "" {
		qs.Add("priority", priority)
	}
	if sources != "" {
		qs.Add("sources", sources)
	}
	if tags != "" {
		qs.Add("tags", sources)
	}

	err := c.makeJSONRequest("GET", c.BaseURL+"/v1/events?"+qs.Encode(), nil, resp)
	return resp, err
}
