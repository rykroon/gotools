package http

import (
	"bytes"
	"io"
	"net/http"
)

// SetRequestBody sets the request body to the given content and content type.
func SetRequestBody(req *http.Request, contentType string, content []byte) {
	req.Body = io.NopCloser(bytes.NewReader(content))
	req.ContentLength = int64(len(content))
	req.Header.Set("Content-Type", contentType)
}

// ReadResponseBody reads the response body and returns it as a byte slice.
func ReadResponseBody(res *http.Response) ([]byte, error) {
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return result, nil
}
