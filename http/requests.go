package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	*http.Request
}

func NewRequest(method, url string, body []byte) (*Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %w", err)
	}
	request := &Request{req}
	if body != nil {
		request.SetBody(body)
	}
	return request, nil
}

func (r *Request) SetHeader(key, value string) {
	r.Header.Set(key, value)
}

func (r *Request) SetAuth(scheme, credentials string) {
	r.SetHeader("Authorization", fmt.Sprint(scheme, " ", credentials))
}

func (r *Request) SetBearerToken(token string) {
	r.SetAuth("Bearer", token)
}

func (r *Request) SetContentType(contentType string) {
	r.SetHeader("Content-Type", contentType)
}

func (r *Request) AppendPath(path string) {
	r.URL = r.URL.JoinPath(path)
}

func (r *Request) AddQueryParam(key, value string) {
	query := r.URL.Query()
	query.Add(key, value)
	r.URL.RawQuery = query.Encode()
}

func (r *Request) SetQueryParam(key, value string) {
	query := r.URL.Query()
	query.Set(key, value)
	r.URL.RawQuery = query.Encode()
}

func (r *Request) SetBody(content []byte) {
	r.Body = io.NopCloser(bytes.NewReader(content))
}

func (r *Request) SetForm(form url.Values) {
	r.SetBody([]byte(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
