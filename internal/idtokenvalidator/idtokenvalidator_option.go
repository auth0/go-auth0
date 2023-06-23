package idtokenvalidator

import (
	"net/http"
	"time"
)

// Option is used for passings options to a IDTokenValidator.
type Option func(*IDTokenValidator)

// WithClockTolerance configures the allowed clock tolerance when validating time based claims.
func WithClockTolerance(clockTolerance time.Duration) Option {
	return func(iv *IDTokenValidator) {
		iv.clockTolerance = clockTolerance
	}
}

// WithHTTPClient configures the HTTP Client used by the JWKS fetcher.
func WithHTTPClient(client *http.Client) Option {
	return func(iv *IDTokenValidator) {
		iv.httpClient = client
	}
}
