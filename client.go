package httpx

import (
	"fmt"
	"net/http"
	neturl "net/url"
	"time"
)

type Client struct {
	client     *http.Client
	baseUrl    string
	decorators []RequestDecorator
}

func NewClient(timeout time.Duration, baseUrl string, decorators ...RequestDecorator) *Client {
	return &Client{
		client:     &http.Client{Timeout: timeout},
		baseUrl:    baseUrl,
		decorators: decorators,
	}
}

func (c *Client) Request(method, url string, decorators ...RequestDecorator) (*Response, error) {
	if c.baseUrl != "" {
		// If the URL is relative, prepend the base URL
		newUrl, err := neturl.JoinPath(c.baseUrl, url)
		if err != nil {
			return nil, fmt.Errorf("Request: %w", err)
		}
		url = newUrl
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Request: %w", err)
	}
	// Apply the client's decorators first, then the request's decorators
	for _, transformer := range c.decorators {
		transformer(req)
	}
	for _, transformer := range decorators {
		transformer(req)
	}
	resp, err := c.client.Do(req)
	return &Response{resp}, err
}

func (c *Client) Get(url string) (*Response, error) {
	return c.Request("GET", url)
}

func (c *Client) Post(url string, payload any) (*Response, error) {
	return nil, nil
}

// Creates a JSON HTTP request with the given method, url, and payload
// func NewJsonRequest(method, url string, payload any) (*http.Request, error) {
// 	jsonBody, err := json.Marshal(payload)
// 	if err != nil {
// 		return nil, fmt.Errorf("NewJsonRequest: %w", err)
// 	}
// 	req, err := NewRequest(method, url, jsonBody)
// 	if err != nil {
// 		return nil, fmt.Errorf("NewJsonRequest: %w", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	return req, nil
// }

// // Creates an HTTP request with the given method, url, and body
// func NewRequest(method, url string, body []byte) (*http.Request, error) {
// 	var bodyReader io.Reader = nil
// 	if body != nil {
// 		bodyReader = bytes.NewReader(body)
// 	}
// 	req, err := http.NewRequest(method, url, bodyReader)
// 	if err != nil {
// 		return nil, fmt.Errorf("NewRequest: %w", err)
// 	}
// 	return req, nil
// }

// func UnmarshalResponse(resp *http.Response, target any) error {
// 	defer resp.Body.Close()
// 	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
// 		return fmt.Errorf("UnmarshalResponse: %w", err)
// 	}
// 	return nil
// }
