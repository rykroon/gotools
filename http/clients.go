package http

import (
	"fmt"
	"net/http"
	neturl "net/url"
)

type Client struct {
	*http.Client
}

// Do sends an HTTP request and returns an HTTP response
func (c *Client) Do(req *Request) (*Response, error) {
	resp, err := c.Client.Do(req.Request)
	if err != nil {
		return nil, fmt.Errorf("Do: %w", err)
	}
	return &Response{Response: resp}, nil
}

func (c *Client) Get(url string, params neturl.Values) (*Response, error) {
	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Get: %w", err)
	}
	req.SetQueryParams(params)
	return c.Do(req)
}

func (c *Client) Post(url string, contentType string, body []byte) (*Response, error) {
	req, err := NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("Post: %w", err)
	}
	req.SetContentType(contentType)
	return c.Do(req)
}

func (c *Client) PostForm(url string, data neturl.Values) (*Response, error) {
	req, err := NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("PostForm: %w", err)
	}
	req.SetForm(data)
	return c.Do(req)
}

func (c *Client) PostJson(url string, body any) (*Response, error) {
	req, err := NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("PostJson: %w", err)
	}
	req.SetJson(body)
	return c.Do(req)
}
