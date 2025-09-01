package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	internal "github.com/auth0/go-auth0/v2/internal/client"
	option "github.com/auth0/go-auth0/v2/management/option"
)

// New creates a new Auth0 Management client by authenticating using the
// supplied client id and secret.
func New(domain string, options ...option.RequestOption) (*Management, error) {
	// Ensure the domain is normalized and always uses https://,
	// since Auth0 only supports HTTPS.

	// Parse the input (with or without scheme)
	u, err := url.Parse(domain)
	if err != nil {
		return nil, fmt.Errorf("invalid Auth0 domain %q: %w", domain, err)
	}

	// If no scheme was provided, Parse puts the entire string in Path.
	// Example: "example.auth0.com" -> Path="example.auth0.com"
	if u.Host == "" {
		u, err = url.Parse("https://" + domain)
		if err != nil {
			return nil, fmt.Errorf("invalid Auth0 domain %q: %w", domain, err)
		}
	}

	// Always force https
	u.Scheme = "https"

	// Remove any trailing slash in the host (e.g. "example.com/")
	u.Host = strings.TrimSuffix(u.Host, "/")

	retryOptions := internal.RetryOptions{
		MaxRetries: 3,
	}

	auth0ClientInfo := internal.DefaultAuth0ClientInfo

	auth0ClientJSON, err := json.Marshal(auth0ClientInfo)
	if err != nil {
		return nil, err
	}

	auth0ClientEncoded := base64.StdEncoding.EncodeToString(auth0ClientJSON)

	m := NewWithOptions(
		append([]option.RequestOption{
			option.WithBaseURL(u.String() + "/api/v2"),
			option.WithHTTPClient(http.DefaultClient),
			option.WithHTTPHeader(http.Header{
				"User-Agent": []string{internal.UserAgent},
			}),
			option.WithHTTPHeader(http.Header{
				"Auth0-Client": []string{auth0ClientEncoded},
			}),
			option.WithMaxAttempts(uint(retryOptions.MaxRetries)),
		}, options...)...,
	)

	return m, nil
}
