package management

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"

	"github.com/auth0/go-auth0/internal/client"
)

// Management is an Auth0 management client used to interact with the Auth0 Management API v2.
type Management struct {
	url             *url.URL
	basePath        string
	userAgent       string
	debug           bool
	tokenSource     oauth2.TokenSource
	http            *http.Client
	auth0ClientInfo *client.Auth0ClientInfo
	retryStrategy   client.RetryOptions
}

// New creates a new Auth0 Management client by authenticating using the supplied client ID and secret.
func New(domain string, options ...Option) (*Management, error) {
	// Sanitize the domain by removing any existing scheme and prefixing it with "https://".
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}
	domain = "https://" + domain

	// Parse the sanitized domain into a URL structure.
	u, err := url.Parse(domain)
	if err != nil {
		return nil, fmt.Errorf("failed to parse domain: %w", err)
	}

	// Initialize the Management client with default settings.
	m := &Management{
		url:             u,
		basePath:        "api/v2",
		userAgent:       client.UserAgent,
		debug:           false,
		http:            http.DefaultClient,
		auth0ClientInfo: client.DefaultAuth0ClientInfo,
		retryStrategy:   client.DefaultRetryOptions,
	}

	// Apply functional options to customize the client.
	for _, option := range options {
		option(m)
	}

	// Wrap the HTTP client with token source and other middleware settings.
	m.http = client.WrapWithTokenSource(
		m.http,
		m.tokenSource,
		client.WithDebug(m.debug),
		client.WithUserAgent(m.userAgent),
		client.WithAuth0ClientInfo(m.auth0ClientInfo),
		client.WithRetries(m.retryStrategy),
	)

	return m, nil
}
