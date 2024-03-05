package http

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
)

// ReadResponseBody reads the response body and returns it as a byte slice.
func ReadBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("GetBody: %w", err)
	}
	return result, nil
}

func ReadJson(res *http.Response, target any) error {
	mediaType, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
	if err != nil {
		return fmt.Errorf("ReadJson: %w", err)
	}
	if mediaType != "application/json" {
		return fmt.Errorf("ReadJson: content type is not application/json")
	}
	body, err := ReadBody(res)
	if err != nil {
		return fmt.Errorf("ReadJson: %w", err)
	}
	if err = json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("ReadJson: %w", err)
	}
	return nil
}

func IsSuccess(res *http.Response) bool {
	return res.StatusCode >= 200 && res.StatusCode < 300
}

func IsRedirect(res *http.Response) bool {
	return res.StatusCode >= 300 && res.StatusCode < 400
}

func IsClientError(res *http.Response) bool {
	return res.StatusCode >= 400 && res.StatusCode < 500
}

func IsServerError(res *http.Response) bool {
	return res.StatusCode >= 500 && res.StatusCode < 600
}

func IsError(res *http.Response) bool {
	return IsClientError(res) || IsServerError(res)
}
