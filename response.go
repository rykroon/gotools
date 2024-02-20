package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	*http.Response
	cachedBody []byte
}

func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

func (r *Response) IsRedirect() bool {
	return r.StatusCode >= 300 && r.StatusCode < 400
}

func (r *Response) IsClientError() bool {
	return r.StatusCode >= 400 && r.StatusCode < 500
}

func (r *Response) IsServerError() bool {
	return r.StatusCode >= 500
}

func (r *Response) IsError() bool {
	return r.IsClientError() || r.IsServerError()
}

func (r *Response) GetContentType() string {
	return r.Header.Get("Content-Type")
}

func (r *Response) GetBody() ([]byte, error) {
	if r.cachedBody != nil {
		return r.cachedBody, nil
	}
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	r.cachedBody = content
	return content, nil
}

func (r *Response) IsJson() bool {
	return r.GetContentType() == "application/json"
}

func (r *Response) GetJson() (map[string]any, error) {
	var target map[string]any
	err := r.UnmarshalJson(target)
	return target, err
}

func (r *Response) UnmarshalJson(target any) error {
	if !r.IsJson() {
		return fmt.Errorf("UnmarshalJson: response is not JSON")
	}
	content, err := r.GetBody()
	if err != nil {
		return fmt.Errorf("UnmarshalJson: %w", err)
	}
	return json.Unmarshal(content, target)

}
