package http

import (
	"fmt"
	"net/http"
	neturl "net/url"
	"time"
)

type Client struct {
	client  *http.Client
	baseUrl *neturl.URL
}

func NewClient(timeout time.Duration, baseUrlString string) (*Client, error) {
	// TODO: Handle when the base URL is not provided
	url, err := neturl.ParseRequestURI(baseUrlString)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %w", err)
	}

	return &Client{
		client:  &http.Client{Timeout: timeout},
		baseUrl: url,
	}, nil
}

func (c *Client) Send(req *Request) (*Response, error) {
	resp, err := c.client.Do(req.Request)
	if err != nil {
		return nil, fmt.Errorf("Send: %w", err)
	}
	return &Response{Response: resp}, nil
}
