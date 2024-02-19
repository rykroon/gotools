package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type RequestBodyMaker interface {
	ContentType() string
	Content() (io.Reader, error)
}

type JsonRequest struct {
	payload any
}

func NewJsonRequest(payload any) *JsonRequest {
	return &JsonRequest{payload}
}

func (r *JsonRequest) ContentType() string {
	return "application/json"
}

func (r *JsonRequest) Content() (io.Reader, error) {
	jsonBody, err := json.Marshal(r.payload)
	if err != nil {
		return nil, fmt.Errorf("JsonRequest.Content: %w", err)
	}
	return bytes.NewReader(jsonBody), nil
}
