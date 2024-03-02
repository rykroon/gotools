package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SetAuth(req *http.Request, schema, creds string) {
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", schema, creds))
}

func SetBearerToken(req *http.Request, token string) {
	SetAuth(req, "Bearer", token)
}

// SetBody sets the request body to the given content and content type.
func SetBody(req *http.Request, contentType string, content []byte) {
	// Set the body and content length.
	reader := bytes.NewReader(content)
	req.ContentLength = int64(reader.Len())
	req.Body = io.NopCloser(reader)

	// Set the GetBody function.
	snapshot := *reader
	req.GetBody = func() (io.ReadCloser, error) {
		r := snapshot
		return io.NopCloser(&r), nil
	}
	// Set the content type.
	req.Header.Set("Content-Type", contentType)
}

func SetJsonBody(req *http.Request, value any) error {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("SetJsonBody: %w", err)
	}
	SetBody(req, "application/json", jsonBody)
	return nil
}
