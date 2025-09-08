package management

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var compiledWhitelistedPathRegexes = compileWhitelistedPathPatterns()

func compileWhitelistedPathPatterns() []*regexp.Regexp {
	patterns := []string{
		`^/api/v2/jobs/verification-email$`,
		`^/api/v2/tickets/email-verification$`,
		`^/api/v2/tickets/password-change$`,
		`^/api/v2/organizations/[^/]+/invitations$`,
		`^/api/v2/users$`,
		`^/api/v2/users/[^/]+$`,
		`^/api/v2/guardian/enrollments/ticket$`,
	}

	compiled := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		compiled[i] = regexp.MustCompile(pattern)
	}

	return compiled
}

// isCustomDomainPathWhitelisted checks if the given path is in the whitelist
// for custom domain header application by matching it against predefined regex patterns.
// Returns true if the path matches any of the whitelisted patterns, false otherwise.
func isCustomDomainPathWhitelisted(path string) bool {
	for _, r := range compiledWhitelistedPathRegexes {
		if r.MatchString(path) {
			return true
		}
	}

	return false
}

// URI returns the absolute URL of the Management API with any path segments
// appended to the end.
func (m *Management) URI(path ...string) string {
	baseURL := &url.URL{
		Scheme: m.url.Scheme,
		Host:   m.url.Host,
		Path:   m.basePath + "/",
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

// methodAllowsBody checks if the HTTP method is one that does not require a
// request body.
//
// This can be used to decide whether to remove an empty object from the
// request body to fix issues with CDNs like CloudFront.
//
// For example:
// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#RequestCustom-get-body
func methodAllowsBody(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return false
	default:
		return true
	}
}

// NewRequest returns a new HTTP request. If the payload is not nil it will be
// encoded as JSON.
func (m *Management) NewRequest(
	ctx context.Context,
	method,
	uri string,
	payload interface{},
	options ...RequestOption,
) (*http.Request, error) {
	var body bytes.Buffer

	setContentType := false

	if payload != nil && methodAllowsBody(method) {
		if err := json.NewEncoder(&body).Encode(payload); err != nil {
			return nil, fmt.Errorf("encoding request payload failed: %w", err)
		}

		encoded := bytes.TrimSpace(body.Bytes())
		if !bytes.Equal(encoded, []byte("null")) {
			setContentType = true
		} else {
			body.Reset()
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, uri, &body)
	if err != nil {
		return nil, err
	}

	if setContentType {
		request.Header.Add("Content-Type", "application/json")
	}

	m.applyCustomDomainHeader(request, uri)

	for _, option := range options {
		if option == nil {
			continue
		}

		option.apply(request)
	}

	return request, nil
}

// Do triggers an HTTP request and returns an HTTP response,
// handling any context cancellations or timeouts.
func (m *Management) Do(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	response, err := m.http.Do(req)
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

// applyCustomDomainHeader adds the Auth0-Custom-Domain header to the given HTTP request
// when the request URI path is in the whitelist for custom domains.
//
// If the customDomainHeader is not set or the URI cannot be parsed, this function does nothing.
func (m *Management) applyCustomDomainHeader(request *http.Request, uri string) {
	if m.customDomainHeader == "" {
		return
	}

	parsedURL, err := url.Parse(uri)
	if err != nil {
		return
	}

	if isCustomDomainPathWhitelisted(parsedURL.Path) {
		request.Header.Set("Auth0-Custom-Domain", m.customDomainHeader)
	}
}

// Request combines NewRequest and Do, while also handling decoding of response payload.
func (m *Management) Request(ctx context.Context, method, uri string, payload interface{}, options ...RequestOption) error {
	request, err := m.NewRequest(ctx, method, uri, payload, options...)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	response, err := m.Do(request)
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
		if err = json.Unmarshal(responseBody, &payload); err != nil {
			return fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return nil
}

// List is an envelope which is typically used when calling List() or Search()
// methods.
//
// It holds metadata such as the total result count, starting offset and limit.
//
// Specific implementations embed this struct, therefore its direct use is not
// useful. Rather it has been made public in order to aid documentation.
type List struct {
	Start  int    `json:"start"`
	Limit  int    `json:"limit"`
	Length int    `json:"length"`
	Total  int    `json:"total"`
	Next   string `json:"next"`
}

// HasNext returns true if the list has more results.
func (l List) HasNext() bool {
	if l.Next != "" {
		return true
	}

	return l.Total > l.Start+l.Limit
}

// RequestOption configures a call (typically to retrieve a resource) to Auth0 with
// query parameters.
type RequestOption interface {
	apply(*http.Request)
}

func newRequestOption(fn func(r *http.Request)) *requestOption {
	return &requestOption{applyFn: fn}
}

type requestOption struct {
	applyFn func(r *http.Request)
}

func (o *requestOption) apply(r *http.Request) {
	o.applyFn(r)
}

func applyListDefaults(options []RequestOption) RequestOption {
	return newRequestOption(func(r *http.Request) {
		PerPage(50).apply(r)
		IncludeTotals(true).apply(r)

		for _, option := range options {
			if option == nil {
				continue
			}

			option.apply(r)
		}
	})
}

func applyListCheckpointDefaults(options []RequestOption) RequestOption {
	return newRequestOption(func(r *http.Request) {
		Take(50).apply(r)

		for _, option := range options {
			if option == nil {
				continue
			}

			option.apply(r)
		}
	})
}

// IncludeFields configures a request to include the desired fields.
func IncludeFields(fields ...string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("fields", strings.Join(fields, ","))
		q.Set("include_fields", "true")
		r.URL.RawQuery = q.Encode()
	})
}

// ExcludeFields configures a request to exclude the desired fields.
func ExcludeFields(fields ...string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("fields", strings.Join(fields, ","))
		q.Set("include_fields", "false")
		r.URL.RawQuery = q.Encode()
	})
}

// Page configures a request to receive a specific page, if the results where
// concatenated.
func Page(page int) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("page", strconv.FormatInt(int64(page), 10))
		r.URL.RawQuery = q.Encode()
	})
}

// PerPage configures a request to limit the amount of items in the result.
func PerPage(items int) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("per_page", strconv.FormatInt(int64(items), 10))
		r.URL.RawQuery = q.Encode()
	})
}

// IncludeTotals configures a request to include totals.
func IncludeTotals(include bool) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("include_totals", strconv.FormatBool(include))
		r.URL.RawQuery = q.Encode()
	})
}

// From configures a request to start from the specified checkpoint.
func From(checkpoint string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("from", checkpoint)
		r.URL.RawQuery = q.Encode()
	})
}

// Take configures a request to limit the amount of items in the result for a checkpoint based request.
func Take(items int) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("take", strconv.FormatInt(int64(items), 10))
		r.URL.RawQuery = q.Encode()
	})
}

// Query configures a request to search on specific query parameters.
//
// For example:
//
//	List(Query(`email:"alice@example.com"`))
//	List(Query(`name:"jane smith"`))
//	List(Query(`logins_count:[100 TO 200}`))
//	List(Query(`logins_count:{100 TO *]`))
//
// See: https://auth0.com/docs/users/search/v3/query-syntax
func Query(s string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("search_engine", "v3")
		q.Set("q", s)
		r.URL.RawQuery = q.Encode()
	})
}

// Parameter configures a request to add arbitrary query parameters to requests
// made to Auth0.
func Parameter(key, value string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set(key, value)
		r.URL.RawQuery = q.Encode()
	})
}

// Header configures a request to add HTTP headers to requests made to Auth0.
func Header(key, value string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		r.Header.Set(key, value)
	})
}

// Body configures a requests body.
func Body(b []byte) RequestOption {
	return newRequestOption(func(r *http.Request) {
		r.Body = io.NopCloser(bytes.NewReader(b))
		if len(b) > 0 {
			r.Header.Set("Content-Type", "application/json")
		}
	})
}

// Stringify returns a string representation of the value passed as an argument.
func Stringify(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Sort configures a request to sort data by the selected field.
// Use 1 to sort ascending and -1 to sort descending.
func Sort(sort string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		q := r.URL.Query()
		q.Set("sort", sort)
		r.URL.RawQuery = q.Encode()
	})
}

// CustomDomainHeader sets the 'Auth0-Custom-Domain' header for this request only,
// overriding any global setting.
func CustomDomainHeader(domain string) RequestOption {
	return newRequestOption(func(r *http.Request) {
		r.Header.Set("Auth0-Custom-Domain", domain)
	})
}
