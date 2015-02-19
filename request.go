package dogapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) makeJSONRequest(method, url string, body, res interface{}) error {
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}

	// Add api_key to URL
	qs := req.URL.Query()
	qs.Add("api_key", c.ApiKey)
	if c.ApplicationKey != "" {
		qs.Add("application_key", c.ApplicationKey)
	}
	req.URL.RawQuery = qs.Encode()

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBodyBuf := new(bytes.Buffer)
	_, err = io.Copy(respBodyBuf, resp.Body)
	if err != nil {
		return err
	}

	if res == nil {
		println(respBodyBuf.String())
		return nil
	}

	respBody := respBodyBuf.String()
	err = json.NewDecoder(respBodyBuf).Decode(res)
	if err != nil {
		return fmt.Errorf("JSON decode error: %s, Body: %s", err, respBody)
	}

	return nil
}
