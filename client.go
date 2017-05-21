package nexmo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 5,
}

// Client represents a Nexmo client
type Client struct {
	Client    *http.Client
	APIKey    string
	APISecret string
}

// New creates a new client with a default HTTP client
func New(apiKey string, apiSecret string) *Client {
	return &Client{
		Client:    defaultHTTPClient,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

// request builds an http request loaded with API credentials and loads the response into resp
func (c *Client) request(method string, path string, params url.Values, resp interface{}) error {
	params.Set("api_key", c.APIKey)
	params.Set("api_secret", c.APISecret)

	var body io.Reader
	switch method {
	case http.MethodPost:
		body = strings.NewReader(params.Encode())
	case http.MethodGet:
		path += "?" + params.Encode()
	}

	// different APIs use different bases apparently...
	switch {
	case strings.HasPrefix(path, "/sms") || strings.HasPrefix(path, "/account"):
		path = "https://rest.nexmo.com" + path
	case strings.HasPrefix(path, "/ni"):
		path = "https://api.nexmo.com" + path
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return errors.Wrap(err, "failed to make request")
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer rsp.Body.Close()

	if err := json.NewDecoder(rsp.Body).Decode(resp); err != nil {
		return errors.Wrap(err, "failed to decode response")
	}

	return nil
}
