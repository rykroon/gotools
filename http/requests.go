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

// SetRequestBody sets the request body to the given content and content type.
func SetRequestBody(req *http.Request, contentType string, content []byte) {
	req.Body = io.NopCloser(bytes.NewReader(content))
	req.ContentLength = int64(len(content))
	req.Header.Set("Content-Type", contentType)
}

func SetRequestJson(req *http.Request, value any) error {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return err
	}
	SetRequestBody(req, "application/json", jsonBody)
	return nil
}
