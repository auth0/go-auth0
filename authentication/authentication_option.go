package authentication

import (
	"net/http"
	"time"
)

// Option is used for passing options to the Authentication client.
type Option func(*Authentication)

// WithClientID configures the Client ID to be provided with requests if one is not provided.
func WithClientID(clientID string) Option {
	return func(a *Authentication) {
		a.clientID = clientID
	}
}

// WithClientSecret configures the Client secret to be provided with requests if one is not provided.
func WithClientSecret(clientSecret string) Option {
	return func(a *Authentication) {
		a.clientSecret = clientSecret
	}
}

// WithIDTokenSigningAlg configures the signing algorithm used for the ID token.
func WithIDTokenSigningAlg(alg string) Option {
	return func(a *Authentication) {
		a.idTokenSigningAlg = alg
	}
}

// WithIDTokenClockTolerance configures the allowed clock tolerance when validating time based claims.
func WithIDTokenClockTolerance(clockTolerance time.Duration) Option {
	return func(a *Authentication) {
		a.idTokenClockTolerance = clockTolerance
	}
}

// WithClient configures to use the provided client for authentication and JWKS calls.
func WithClient(client *http.Client) Option {
	return func(a *Authentication) {
		a.http = client
	}
}
