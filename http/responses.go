package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ReadResponseBody reads the response body and returns it as a byte slice.
func GetResponseBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetResponseJson(res *http.Response, target any) error {
	body, err := GetResponseBody(res)
	if err != nil {
		return err
	}
	if res.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("response content type is not application/json")
	}
	err = json.Unmarshal(body, target)
	return err
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
