package httpx

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	*http.Response
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

func (r *Response) GetContent() ([]byte, error) {
	defer r.Body.Close()
	// I believe the body will need to be cached in order to be read again
	return io.ReadAll(r.Body)
}

func (r *Response) GetJson() error {
	if !r.IsJson() {
		return fmt.Errorf("Response is not JSON")
	}
	return nil
}

func (r *Response) IsJson() bool {
	return r.GetContentType() == "application/json"
}
