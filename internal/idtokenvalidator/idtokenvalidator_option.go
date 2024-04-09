package idtokenvalidator

import (
	"net/http"
	"strings"
	"time"
)

// Option is used for passing options to an `IDTokenValidator`.
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

// WithInsecure configures the issuer to use HTTP instead of HTTPS.
//
// This option is available for testing purposes and should not be used in
// production.
func WithInsecure() Option {
	return func(iv *IDTokenValidator) {
		if i := strings.Index(iv.issuer, "//"); i != -1 {
			iv.issuer = "http://" + iv.issuer[i+2:]
		}
	}
}
