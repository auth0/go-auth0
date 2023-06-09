package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// URI returns the absolute URL of the Management API with any path segments
// appended to the end.
func (a *Authentication) URI(path ...string) string {
	baseURL := &url.URL{
		Scheme: a.url.Scheme,
		Host:   a.url.Host,
		Path:   a.basePath + "/",
	}

	const escapedForwardSlash = "%2F"
	var escapedPath []string
	for _, unescapedPath := range path {
		// Go's url.PathEscape will not escape "/", but some user IDs do have a valid "/" in them.
		// See https://github.com/golang/go/blob/b55a2fb3b0d67b346bac871737b862f16e5a6447/src/net/url/url.go#L141.
		defaultPathEscaped := url.PathEscape(unescapedPath)
		escapedPath = append(
			escapedPath,
			strings.ReplaceAll(defaultPathEscaped, "/", escapedForwardSlash),
		)
	}

	return baseURL.String() + strings.Join(escapedPath, "/")
}

// NewRequest returns a new HTTP request. If the payload is not nil it will be
// encoded as JSON.
func (a *Authentication) NewRequest(
	ctx context.Context,
	method,
	uri string,
	payload interface{},
	opts ...RequestOption,
) (*http.Request, error) {
	var body bytes.Buffer
	if payload != nil {
		if err := json.NewEncoder(&body).Encode(payload); err != nil {
			return nil, fmt.Errorf("encoding request payload failed: %w", err)
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, uri, &body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	for _, option := range opts {
		option.apply(request, nil)
	}

	return request, nil
}

// Do triggers an HTTP request and returns an HTTP response,
// handling any context cancellations or timeouts.
func (a *Authentication) Do(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	response, err := a.http.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	return response, nil
}

// Request combines NewRequest and Do, while also handling decoding of response payload.
func (a *Authentication) Request(ctx context.Context, method, uri string, payload interface{}, resp interface{}) error {
	request, err := a.NewRequest(ctx, method, uri, &payload)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	response, err := a.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}
	defer response.Body.Close()

	// If the response contains a client or a server error then return the error.
	if response.StatusCode >= http.StatusBadRequest {
		return newError(response)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	if len(responseBody) > 0 && string(responseBody) != "{}" {
		if err = json.Unmarshal(responseBody, &resp); err != nil {
			return fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return nil
}

// NewFormRequest returns a new HTTP request. If the payload is not nil it will be
// encoded as JSON.
func (a *Authentication) NewFormRequest(
	ctx context.Context,
	method,
	uri string,
	payload url.Values,
	opts ...RequestOption,
) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, uri, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for _, option := range opts {
		option.apply(request, payload)
	}

	body := strings.NewReader(payload.Encode())

	request.Body = io.NopCloser(body)
	request.ContentLength = int64(body.Len())

	return request, nil
}

// FormRequest combines FormRequest and Do, while also handling decoding of response payload.
func (a *Authentication) FormRequest(ctx context.Context, method, uri string, payload url.Values, resp interface{}, opts ...RequestOption) error {
	request, err := a.NewFormRequest(ctx, method, uri, payload, opts...)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	response, err := a.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}
	defer response.Body.Close()

	// If the response contains a client or a server error then return the error.
	if response.StatusCode >= http.StatusBadRequest {
		return newError(response)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	if len(responseBody) > 0 && string(responseBody) != "{}" {
		if err = json.Unmarshal(responseBody, &resp); err != nil {
			return fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return nil
}

// RequestOption configures a call (typically to retrieve a resource) to Auth0 with
// query parameters.
type RequestOption interface {
	apply(*http.Request, url.Values)
}

func newRequestOption(fn func(r *http.Request, b url.Values)) *requestOption {
	return &requestOption{applyFn: fn}
}

type requestOption struct {
	applyFn func(r *http.Request, b url.Values)
}

func (o *requestOption) apply(r *http.Request, b url.Values) {
	o.applyFn(r, b)
}

// Header configures a request to add HTTP headers to requests made to Auth0.
func Header(key, value string) RequestOption {
	return newRequestOption(func(r *http.Request, b url.Values) {
		r.Header.Set(key, value)
	})
}
