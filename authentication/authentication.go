package authentication

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/auth0/go-auth0/internal/client"
)

// Authentication is the auth client.
type Authentication struct {
	Database *Database
	OAuth    *OAuth

	url             *url.URL
	basePath        string
	debug           bool
	http            *http.Client
	auth0ClientInfo *client.Auth0ClientInfo
	common          manager
	clientID        string
	clientSecret    string
}

type manager struct {
	authentication *Authentication
}

// New creates an auth client.
func New(domain string, options ...Option) (*Authentication, error) {
	// Ignore the scheme if it was defined in the domain variable, then prefix
	// with https as it's the only scheme supported by the Auth0 API.
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}
	domain = "https://" + domain

	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}

	a := &Authentication{
		url:             u,
		basePath:        "",
		debug:           false,
		http:            http.DefaultClient,
		auth0ClientInfo: client.DefaultAuth0ClientInfo,
	}

	for _, option := range options {
		option(a)
	}

	a.common.authentication = a
	a.Database = (*Database)(&a.common)
	a.OAuth = (*OAuth)(&a.common)

	return a, nil
}
