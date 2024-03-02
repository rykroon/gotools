package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func NewDataRequest(method, url string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("NewDataRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	return req, nil
}

func NewTextRequest(method, url, content string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("NewTextRequest: %w", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	return req, nil
}

func NewJsonRequest(method, url string, value any) (*http.Request, error) {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("NewJsonRequest: %w", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("NewJsonRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func SetAuth(req *http.Request, schema, creds string) {
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", schema, creds))
}

func SetBearerToken(req *http.Request, token string) {
	SetAuth(req, "Bearer", token)
}
