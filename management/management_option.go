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

// WithContext configures the management client to use the provided context
// instead of the provided one.
func WithContext(ctx context.Context) Option {
	return func(m *Management) {
		m.ctx = ctx
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
func WithClientCredentials(clientID, clientSecret string) Option {
	return func(m *Management) {
		m.tokenSource = client.OAuth2ClientCredentials(m.ctx, m.url.String(), clientID, clientSecret)
	}
}

// WithClientCredentialsAndAudience configures management to authenticate using the client
// credentials authentication flow and a custom audience.
func WithClientCredentialsAndAudience(clientID, clientSecret, audience string) Option {
	return func(m *Management) {
		m.tokenSource = client.OAuth2ClientCredentialsAndAudience(
			m.ctx,
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
func WithClient(client *http.Client) Option {
	return func(m *Management) {
		m.http = client
	}
}

// WithTelemetry configures the management client to use the provided telemetry data
// instead of the default one.
func WithTelemetry(telemetry client.Telemetry) Option {
	return func(m *Management) {
		if !telemetry.IsEmpty() {
			m.telemetry = &telemetry
		}
	}
}

// WithNoTelemetry configures the management client to not send the "Auth0-Client" header
// at all.
func WithNoTelemetry() Option {
	return func(m *Management) {
		m.telemetry = nil
	}
}
