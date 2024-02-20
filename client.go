package httpx

import (
	"fmt"
	"net/http"
	neturl "net/url"
	"time"
)

type Client struct {
	client     *http.Client
	baseUrl    *neturl.URL
	decorators []RequestDecorator
}

func NewClient(timeout time.Duration, baseUrlString string, decorators ...RequestDecorator) (*Client, error) {
	// TODO: Handle when the base URL is not provided
	url, err := neturl.ParseRequestURI(baseUrlString)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %w", err)
	}

	return &Client{
		client:     &http.Client{Timeout: timeout},
		baseUrl:    url,
		decorators: decorators,
	}, nil
}

func (c *Client) Request(method, urlString string, decorators ...RequestDecorator) (*Response, error) {
	// TODO: Handle when the client url is not provided

	url := c.baseUrl.JoinPath(urlString)

	req, err := http.NewRequest(method, url.String(), nil)
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

func (c *Client) Get(url string, decorators ...RequestDecorator) (*Response, error) {
	return c.Request("GET", url, decorators...)
}

func (c *Client) Post(url string, decorators ...RequestDecorator) (*Response, error) {
	return c.Request("POST", url, decorators...)
}
