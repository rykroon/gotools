package http

import (
	"fmt"
	"net/http"
)

/*
 * Future ideas
 * - Add 'params' parameter to Get
 * - Add 'headers' parameter
 */

func Get(url string) (*Response, error) {
	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Get: %w", err)
	}
	resp, err := http.DefaultClient.Do(req.Request)
	if err != nil {
		return nil, fmt.Errorf("Get: %w", err)
	}
	return &Response{Response: resp}, nil
}

func Post(url string, contentType string, body []byte) (*Response, error) {
	req, err := NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("Post: %w", err)
	}
	req.SetContentType(contentType)
	resp, err := http.DefaultClient.Do(req.Request)
	if err != nil {
		return nil, fmt.Errorf("Post: %w", err)
	}
	return &Response{Response: resp}, nil
}

func PostJson(url string, value any) (*Response, error) {
	req, err := NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("PostJson: %w", err)
	}
	err = req.SetJson(value)
	if err != nil {
		return nil, fmt.Errorf("PostJson: %w", err)
	}
	resp, err := http.DefaultClient.Do(req.Request)
	if err != nil {
		return nil, fmt.Errorf("PostJson: %w", err)
	}
	return &Response{Response: resp}, nil
}
