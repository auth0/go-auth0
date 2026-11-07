package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	internal "github.com/auth0/go-auth0/v3/internal/client"
	core "github.com/auth0/go-auth0/v3/management/core"
	option "github.com/auth0/go-auth0/v3/management/option"
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

	// Clear any path that url.Parse may have extracted from the domain
	// (e.g. trailing slash in "example.com/" ends up in u.Path, not u.Host).
	u.Path = ""

	retryOptions := internal.RetryOptions{
		MaxRetries: 3,
	}

	// Extract NoAuth0ClientInfo, Auth0ClientEnv and any caller supplied HTTP
	// client in a single loop
	var noAuth0ClientInfo bool
	var envOnlyOptions []option.RequestOption
	var callerHTTPClient core.HTTPClient

	passthroughOptions := make([]option.RequestOption, 0, len(options))

	for _, opt := range options {
		switch opt := opt.(type) {
		case *core.NoAuth0ClientInfoOption:
			noAuth0ClientInfo = true
		case *core.Auth0ClientEnvEntryOption:
			envOnlyOptions = append(envOnlyOptions, opt)
		case *core.HTTPClientOption:
			// Keep the client so the transports below can be layered onto it,
			// and drop the option itself. Options are applied in order, so
			// passing it through would replace the wrapped client built here
			// and with it the User-Agent, Auth0-Client and custom domain
			// headers.
			callerHTTPClient = opt.HTTPClient

			continue
		}

		passthroughOptions = append(passthroughOptions, opt)
	}

	// Build the base options that will be passed to NewWithOptions
	baseOptions := []option.RequestOption{
		option.WithBaseURL(u.String() + "/api/v2"),
		option.WithMaxAttempts(uint(retryOptions.MaxRetries)),
	}

	// Start from the caller's client when one was given, so its timeout, proxy
	// and transport settings are kept, and fall back to a clone of
	// DefaultClient otherwise
	httpClient := baseHTTPClient(callerHTTPClient)

	// Apply transports in order: UserAgent, Auth0-Client (if needed), then CustomDomain
	httpClient.Transport = internal.UserAgentTransport(httpClient.Transport, internal.UserAgent)

	// Only add Auth0-Client header transport if NoAuth0ClientInfo is not set
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

		// Apply Auth0ClientInfo transport
		transport, err := internal.Auth0ClientInfoTransport(httpClient.Transport, auth0ClientInfo)
		if err != nil {
			return nil, err
		}
		httpClient.Transport = transport
	}

	// Apply CustomDomainHeaderTransport last (pass empty string for client-level domain; request-level will use the hint header)
	httpClient.Transport = internal.CustomDomainHeaderTransport(httpClient.Transport, "")

	baseOptions = append(baseOptions, option.WithHTTPClient(httpClient))

	m := NewWithOptions(append(baseOptions, passthroughOptions...)...)

	return m, nil
}

// baseHTTPClient returns the client that the SDK transports are layered onto.
//
// Callers can supply any core.HTTPClient, so a client that is not an
// *http.Client is delegated to through a round tripper, which keeps the SDK
// headers on the requests it issues. Clients are cloned rather than used
// directly, so that wrapping their transport does not modify the caller's copy.
func baseHTTPClient(callerHTTPClient core.HTTPClient) *http.Client {
	switch client := callerHTTPClient.(type) {
	case nil:
	case *http.Client:
		if client != nil {
			clone := *client

			return &clone
		}
	default:
		return &http.Client{
			Transport: internal.RoundTripFunc(client.Do),
		}
	}

	clone := *http.DefaultClient

	return &clone
}
