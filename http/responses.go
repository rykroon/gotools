package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	*http.Response
}

// IsSuccess checks if the response is a success
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// IsRedirect checks if the response is a redirect
func (r *Response) IsRedirect() bool {
	return r.StatusCode >= 300 && r.StatusCode < 400
}

// IsClientError checks if the response is a client error
func (r *Response) IsClientError() bool {
	return r.StatusCode >= 400 && r.StatusCode < 500
}

// IsServerError checks if the response is a server error
func (r *Response) IsServerError() bool {
	return r.StatusCode >= 500
}

// IsError checks if the response is an error
func (r *Response) IsError() bool {
	return r.IsClientError() || r.IsServerError()
}

// GetContentType returns the response content type
func (r *Response) GetContentType() string {
	return r.Header.Get("Content-Type")
}

// GetBody returns the response body as a byte slice
func (r *Response) GetBody() ([]byte, error) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("GetBody: %w", err)
	}
	defer r.Body.Close()
	return content, nil
}

func (r *Response) IsForm() bool {
	return r.GetContentType() == "application/x-www-form-urlencoded"
}

func (r *Response) GetForm() (url.Values, error) {
	if !r.IsForm() {
		return nil, fmt.Errorf("GetForm: response is not a form")
	}
	content, err := r.GetBody()
	if err != nil {
		return nil, fmt.Errorf("GetForm: %w", err)
	}
	return url.ParseQuery(string(content))
}

// IsJson checks if the response is JSON
func (r *Response) IsJson() bool {
	return r.GetContentType() == "application/json"
}

// GetJson returns the JSON response as a map
func (r *Response) GetJson() (map[string]any, error) {
	var target map[string]any
	err := r.UnmarshalJson(target)
	return target, err
}

// UnmarshalJson unmarshals the JSON response into the target
func (r *Response) UnmarshalJson(target any) error {
	if !r.IsJson() {
		return fmt.Errorf("UnmarshalJson: response is not JSON")
	}
	content, err := r.GetBody()
	if err != nil {
		return fmt.Errorf("UnmarshalJson: %w", err)
	}
	return json.Unmarshal(content, target)
}
