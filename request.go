package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
 * Notes: I like this approach much better than the decoraotors.
 * I like the that the request has setter methods and the response has
 * getter methods.
 */

type Request struct {
	*http.Request
}

func NewRequest(method, urlString string) (*Request, error) {
	req, err := http.NewRequest(method, urlString, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %w", err)
	}
	return &Request{req}, nil
}

func (r *Request) SetHeader(key, value string) {
	r.Header.Set(key, value)
}

func (r *Request) SetAuth(scheme, credentials string) {
	r.Header.Set("Authorization", fmt.Sprint(scheme, " ", credentials))
}

func (r *Request) SetBearerToken(token string) {
	r.SetAuth("Bearer", token)
}

func (r *Request) SetContentType(contentType string) {
	r.Header.Set("Content-Type", contentType)
}

func (r *Request) SetPath(path string) {
	r.URL = r.URL.JoinPath(path)
}

func (r *Request) SetBody(content []byte) {
	r.Body = io.NopCloser(bytes.NewReader(content))
}

func (r *Request) SetJson(value any) error {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("SetJson: %w", err)
	}
	r.SetBody(jsonBody)
	r.SetContentType("application/json")
	return nil
}
