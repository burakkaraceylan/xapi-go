package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Request represents a request to any LRS
type Request struct {
	Method      string
	URL         string
	Headers     *map[string]string
	QueryParams *map[string]string
	Content     *string
}

// This function is used to generate an http.Request from a Request struct
func (r *Request) Init() (*http.Request, error) {
	if len(r.Method) == 0 {
		return nil, errors.New("method can't be empty")
	}

	if len(r.URL) == 0 {
		return nil, errors.New("url can't be empty")
	}

	req, err := http.NewRequest(r.Method, r.URL, nil)

	if r.Content != nil && len(*r.Content) > 0 {
		req.Body = io.NopCloser(strings.NewReader(*r.Content))
	}

	if r.QueryParams != nil && len(*r.QueryParams) > 0 {
		q := req.URL.Query()

		for param, value := range *r.QueryParams {
			q.Add(param, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	if r.Headers != nil && len(*r.Headers) > 0 {
		for header, value := range *r.Headers {
			req.Header.Set(header, value)
		}
	}

	if err != nil {
		return nil, err
	}

	return req, nil
}

// Response represents a Response fron any LRS
type Response struct {
	Status   int
	Request  *http.Request
	Response *http.Response
}

func (r *Response) Bind(object any) error {

	bodyBytes, _ := io.ReadAll(r.Response.Body)

	if err := json.Unmarshal(bodyBytes, object); err != nil {
		return err
	}

	return nil
}

func (r *Response) String() string {
	str := fmt.Sprintf("Code: %d\n", r.Status)
	str += fmt.Sprintf("URL: %s\n", r.Request.URL)
	return str
}
