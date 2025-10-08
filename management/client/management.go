package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	internal "github.com/auth0/go-auth0/v2/internal/client"
	core "github.com/auth0/go-auth0/v2/management/core"
	option "github.com/auth0/go-auth0/v2/management/option"
)

// New creates a new Auth0 Management client by authenticating using the
// supplied client id and secret.
func New(domain string, options ...option.RequestOption) (*Management, error) {
	// Strip any existing scheme from the domain (e.g., "https://", "http://")
	// This ensures consistent handling whether user provides "example.com" or "https://example.com"
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}

	// Always prepend https:// (WithInsecure() will change this to http:// if needed)
	u, err := url.Parse("https://" + domain)
	if err != nil {
		return nil, fmt.Errorf("invalid Auth0 domain %q: %w", domain, err)
	}

	// Remove any trailing slash in the host (e.g. "example.com/")
	u.Host = strings.TrimSuffix(u.Host, "/")

	retryOptions := internal.RetryOptions{
		MaxRetries: 3,
	}

	// Extract NoAuth0ClientInfo and Auth0ClientEnv options in a single loop
	var noAuth0ClientInfo bool
	var envOnlyOptions []option.RequestOption
	for _, opt := range options {
		switch opt.(type) {
		case *core.NoAuth0ClientInfoOption:
			noAuth0ClientInfo = true
		case *core.Auth0ClientEnvEntryOption:
			envOnlyOptions = append(envOnlyOptions, opt)
		}
	}

	// Build the base options that will be passed to NewWithOptions
	baseOptions := []option.RequestOption{
		option.WithBaseURL(u.String() + "/api/v2"),
		option.WithHTTPClient(http.DefaultClient),
		option.WithHTTPHeader(http.Header{
			"User-Agent": []string{internal.UserAgent},
		}),
		option.WithMaxAttempts(uint(retryOptions.MaxRetries)),
	}

	// Apply CustomDomainHeaderTransport to support request-level custom domain headers
	// Pass empty string for client-level domain; request-level will use the hint header
	httpClient := *http.DefaultClient // Clone to avoid modifying the global client
	httpClient.Transport = internal.CustomDomainHeaderTransport(http.DefaultClient.Transport, "")
	baseOptions = append(baseOptions, option.WithHTTPClient(&httpClient))

	// Only add Auth0-Client header if NoAuth0ClientInfo is not set
	if !noAuth0ClientInfo {
		// Process env options if any were found
		envOpts := core.NewRequestOptions(envOnlyOptions...)

		// Build Auth0ClientInfo with default values plus any custom env entries
		auth0ClientInfo := &internal.Auth0ClientInfo{
			Name:    internal.DefaultAuth0ClientInfo.Name,
			Version: internal.DefaultAuth0ClientInfo.Version,
			Env:     make(map[string]string),
		}

		// Copy default env entries
		for k, v := range internal.DefaultAuth0ClientInfo.Env {
			auth0ClientInfo.Env[k] = v
		}

		// Merge custom env entries
		for k, v := range envOpts.Auth0ClientEnv {
			auth0ClientInfo.Env[k] = v
		}

		auth0ClientJSON, err := json.Marshal(auth0ClientInfo)
		if err != nil {
			return nil, err
		}

		auth0ClientEncoded := base64.StdEncoding.EncodeToString(auth0ClientJSON)

		// Add the Auth0-Client header
		baseOptions = append(baseOptions, option.WithHTTPHeader(http.Header{
			"Auth0-Client": []string{auth0ClientEncoded},
		}))
	}

	m := NewWithOptions(append(baseOptions, options...)...)

	return m, nil
}
