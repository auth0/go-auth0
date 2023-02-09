package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"encoding/base64"
	"encoding/json"

	"github.com/auth0/go-auth0"
)

// UserAgent is the default user agent string.
var UserAgent = fmt.Sprintf("Go-Auth0-SDK/%s", auth0.Version)

// Telemetry is the structure used to send telemetry data in the "Auth0-Client" header.
type Telemetry struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Env     map[string]string `json:"env,omitempty"`
}

// IsEmpty checks whether the provide Telemetry data is nil or has no data to allow
// short-circuiting the "Auth0-Client" header configuration.
func (td *Telemetry) IsEmpty() bool {
	if td == nil {
		return true
	}
	return td.Name == "" && td.Version == "" && len(td.Env) == 0
}

// DefaultTelemetryData is the default telemetry data sent by the go-auth0 SDK.
var DefaultTelemetryData = &Telemetry{Name: "go-auth0", Version: auth0.Version}

// RoundTripFunc is an adapter to allow the use of ordinary functions as HTTP
// round trips.
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
func (rf RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rf(req)
}

// RateLimitTransport wraps base transport with rate limiting functionality.
//
// When a 429 status code is returned by the remote server, the
// "X-RateLimit-Reset" header is used to determine how long the transport will
// wait until re-issuing the failed request.
func RateLimitTransport(base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return rehttp.NewTransport(base, retry, delay)
}

func retry(attempt rehttp.Attempt) bool {
	if attempt.Response == nil {
		return false
	}
	return attempt.Response.StatusCode == http.StatusTooManyRequests
}

func delay(attempt rehttp.Attempt) time.Duration {
	resetAt := attempt.Response.Header.Get("X-RateLimit-Reset")
	resetAtUnix, err := strconv.ParseInt(resetAt, 10, 64)
	if err != nil {
		resetAtUnix = time.Now().Add(5 * time.Second).Unix()
	}
	return time.Duration(resetAtUnix-time.Now().Unix()) * time.Second
}

// UserAgentTransport wraps base transport with a customized "User-Agent" header.
func UserAgentTransport(base http.RoundTripper, userAgent string) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("User-Agent", userAgent)
		return base.RoundTrip(req)
	})
}

// TelemetryTransport wraps base transport with a customized "Auth0-Client" header.
func TelemetryTransport(base http.RoundTripper, telemetryData *Telemetry) (http.RoundTripper, error) {
	if base == nil {
		base = http.DefaultTransport
	}

	telemetryDataJson, err := json.Marshal(telemetryData)
	if err != nil {
		return nil, err
	}

	encodedTelemetry := base64.StdEncoding.EncodeToString(telemetryDataJson)

	return RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("Auth0-Client", encodedTelemetry)
		return base.RoundTrip(req)
	}), nil
}

func dumpRequest(r *http.Request) {
	b, _ := httputil.DumpRequestOut(r, true)
	log.Printf("\n%s\n", b)
}

func dumpResponse(r *http.Response) {
	b, _ := httputil.DumpResponse(r, true)
	log.Printf("\n%s\n\n", b)
}

// DebugTransport wraps base transport with the ability to log the contents
// of requests and responses.
func DebugTransport(base http.RoundTripper, debug bool) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	if !debug {
		return base
	}
	return RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		dumpRequest(req)
		res, err := base.RoundTrip(req)
		if err != nil {
			return res, err
		}
		dumpResponse(res)
		return res, nil
	})
}

// Option is the type used to configure a client.
type Option func(*http.Client)

// WithDebug configures the client to enable debug.
func WithDebug(debug bool) Option {
	return func(c *http.Client) {
		c.Transport = DebugTransport(c.Transport, debug)
	}
}

// WithRateLimit configures the client to enable rate limiting.
func WithRateLimit() Option {
	return func(c *http.Client) {
		c.Transport = RateLimitTransport(c.Transport)
	}
}

// WithUserAgent configures the client to overwrite the user agent header.
func WithUserAgent(userAgent string) Option {
	return func(c *http.Client) {
		c.Transport = UserAgentTransport(c.Transport, userAgent)
	}
}

// WithTelemetry configures the client to overwrite the "Auth0-Client" header.
func WithTelemetry(telemetryData *Telemetry) Option {
	return func(c *http.Client) {
		if telemetryData.IsEmpty() {
			return
		}
		transport, err := TelemetryTransport(c.Transport, telemetryData)
		if err != nil {
			return
		}
		c.Transport = transport
	}
}

// Wrap the base client with transports that enable OAuth2 authentication.
func Wrap(base *http.Client, tokenSource oauth2.TokenSource, options ...Option) *http.Client {
	if base == nil {
		base = http.DefaultClient
	}
	client := &http.Client{
		Timeout: base.Timeout,
		Transport: &oauth2.Transport{
			Base:   base.Transport,
			Source: tokenSource,
		},
	}
	for _, option := range options {
		option(client)
	}
	return client
}

// OAuth2ClientCredentials sets the oauth2 client credentials.
func OAuth2ClientCredentials(ctx context.Context, uri, clientID, clientSecret string) oauth2.TokenSource {
	audience := uri + "/api/v2/"
	return OAuth2ClientCredentialsAndAudience(ctx, uri, clientID, clientSecret, audience)
}

// OAuth2ClientCredentialsAndAudience sets the oauth2
// client credentials with a custom audience.
func OAuth2ClientCredentialsAndAudience(
	ctx context.Context,
	uri,
	clientID,
	clientSecret,
	audience string,
) oauth2.TokenSource {
	cfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     uri + "/oauth/token",
		EndpointParams: url.Values{
			"audience": []string{audience},
		},
	}

	return cfg.TokenSource(ctx)
}

// StaticToken sets a static token to be used for oauth2.
func StaticToken(token string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
}
