package management

import (
	"context"
	"net/http"

	"github.com/auth0/go-auth0/internal/client"
)

// Option is used for passing options to the Management client.
type Option func(*Management)

// WithDebug configures the management client to dump http requests and
// responses to stdout.
func WithDebug(d bool) Option {
	return func(m *Management) {
		m.debug = d
	}
}

// WithUserAgent configures the management client to use the provided user agent
// string instead of the default one.
func WithUserAgent(userAgent string) Option {
	return func(m *Management) {
		m.userAgent = userAgent
	}
}

// WithClientCredentials configures management to authenticate using the client
// credentials authentication flow.
func WithClientCredentials(ctx context.Context, clientID, clientSecret string) Option {
	return func(m *Management) {
		m.tokenSource = client.OAuth2ClientCredentials(ctx, m.url.String(), clientID, clientSecret)
	}
}

// WithClientCredentialsAndAudience configures management to authenticate using the client
// credentials authentication flow and a custom audience.
func WithClientCredentialsAndAudience(ctx context.Context, clientID, clientSecret, audience string) Option {
	return func(m *Management) {
		m.tokenSource = client.OAuth2ClientCredentialsAndAudience(
			ctx,
			m.url.String(),
			clientID,
			clientSecret,
			audience,
		)
	}
}

// WithStaticToken configures management to authenticate using a static
// authentication token.
func WithStaticToken(token string) Option {
	return func(m *Management) {
		m.tokenSource = client.StaticToken(token)
	}
}

// WithInsecure configures management to not use an authentication token and
// use HTTP instead of HTTPS.
//
// This option is available for testing purposes and should not be used in
// production.
func WithInsecure() Option {
	return func(m *Management) {
		m.tokenSource = client.StaticToken("insecure")
		m.url.Scheme = "http"
	}
}

// WithClient configures management to use the provided client.
// Note: If you are providing a client with a retry strategy, then you should
// also pass `WithNoRetries` so that the default retry strategy is not added.
func WithClient(client *http.Client) Option {
	return func(m *Management) {
		m.http = client
	}
}

// WithAuth0ClientEnvEntry allows adding extra environment keys to the client information.
func WithAuth0ClientEnvEntry(key string, value string) Option {
	return func(m *Management) {
		if !m.auth0ClientInfo.IsEmpty() {
			if m.auth0ClientInfo.Env == nil {
				m.auth0ClientInfo.Env = map[string]string{}
			}
			m.auth0ClientInfo.Env[key] = value
		}
	}
}

// WithNoAuth0ClientInfo configures the management client to not send the "Auth0-Client" header
// at all.
func WithNoAuth0ClientInfo() Option {
	return func(m *Management) {
		m.auth0ClientInfo = nil
	}
}

// WithRetries configures the management client to only retry under the conditions provided.
func WithRetries(maxRetries int, statuses []int) Option {
	return func(m *Management) {
		m.retryStrategy = client.RetryOptions{
			MaxRetries: maxRetries,
			Statuses:   statuses,
		}
	}
}

// WithNoRetries configures the management client to only retry under the conditions provided.
func WithNoRetries() Option {
	return func(m *Management) {
		m.retryStrategy = client.RetryOptions{}
	}
}
