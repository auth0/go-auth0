package authentication

import (
	"net/http"
	"time"

	"github.com/auth0/go-auth0/internal/client"
)

// Option is used for passing options to the Authentication client.
type Option func(*Authentication)

// WithClientID configures the Client ID to be provided with requests if one is not provided.
func WithClientID(clientID string) Option {
	return func(a *Authentication) {
		a.clientID = clientID
	}
}

// WithClientSecret configures the Client Secret to be provided with requests if one is not provided.
func WithClientSecret(clientSecret string) Option {
	return func(a *Authentication) {
		a.clientSecret = clientSecret
	}
}

// WithClientAssertion configures the signing key to be used when performing Private Key JWT Auth.
func WithClientAssertion(signingKey string, signingAlg string) Option {
	return func(a *Authentication) {
		a.clientAssertionSigningKey = signingKey
		a.clientAssertionSigningAlg = signingAlg
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

// WithInsecure configures management to use HTTP instead of HTTPS.
//
// This option is available for testing purposes and should not be used in production.
func WithInsecure() Option {
	return func(a *Authentication) {
		a.url.Scheme = "http"
	}
}

// WithClient configures to use the provided client for authentication and JWKS calls.
func WithClient(client *http.Client) Option {
	return func(a *Authentication) {
		a.http = client
	}
}

// WithNoAuth0ClientInfo configures the management client to not send the "Auth0-Client" header
// at all.
func WithNoAuth0ClientInfo() Option {
	return func(a *Authentication) {
		a.auth0ClientInfo = nil
	}
}

// WithAuth0ClientEnvEntry allows adding extra environment keys to the client information.
func WithAuth0ClientEnvEntry(key string, value string) Option {
	return func(a *Authentication) {
		if !a.auth0ClientInfo.IsEmpty() {
			if a.auth0ClientInfo.Env == nil {
				a.auth0ClientInfo.Env = map[string]string{}
			}
			a.auth0ClientInfo.Env[key] = value
		}
	}
}

// WithRetries configures the management client to only retry under the conditions provided.
func WithRetries(maxRetries int, statuses []int) Option {
	return func(a *Authentication) {
		a.retryStrategy = client.RetryOptions{
			MaxRetries: maxRetries,
			Statuses:   statuses,
		}
	}
}

// WithNoRetries configures the management client to only retry under the conditions provided.
func WithNoRetries() Option {
	return func(a *Authentication) {
		a.retryStrategy = client.RetryOptions{}
	}
}
