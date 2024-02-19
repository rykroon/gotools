package httpx

import (
	"fmt"
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
