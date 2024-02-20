package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestDecorator func(req *http.Request)

func BasicAuth(username, password string) RequestDecorator {
	return func(req *http.Request) {
		req.SetBasicAuth(username, password)
	}
}

func Header(key, value string) RequestDecorator {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

func BearerAuth(token string) RequestDecorator {
	return Header("Authorization", fmt.Sprintf("Bearer %s", token))
}

func ContentType(contentType string) RequestDecorator {
	return Header("Content-Type", contentType)
}

func JoinPath(paths ...string) RequestDecorator {
	return func(req *http.Request) {
		req.URL = req.URL.JoinPath(paths...)
	}
}

func Body(contentType string, body []byte) RequestDecorator {
	return func(req *http.Request) {
		req.Body = io.NopCloser(bytes.NewReader(body))
		ContentType(contentType)(req)
	}
}

func Json(value any) (RequestDecorator, error) {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("Json: %w", err)
	}
	return Body("application/json", jsonBody), nil
}
