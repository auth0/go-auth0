package client

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/auth0/go-auth0"
)

// UserAgent is the default user agent string.
var UserAgent = fmt.Sprintf("Go-Auth0/%s", auth0.Version)

// Auth0ClientInfo is the structure used to send client information in the "Auth0-Client" header.
type Auth0ClientInfo struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Env     map[string]string `json:"env,omitempty"`
}

// IsEmpty checks whether the provided Auth0ClientInfo data is nil or has no data to allow
// short-circuiting the "Auth0-Client" header configuration.
func (td *Auth0ClientInfo) IsEmpty() bool {
	if td == nil {
		return true
	}
	return td.Name == "" && td.Version == "" && len(td.Env) == 0
}

// DefaultAuth0ClientInfo is the default client information sent by the go-auth0 SDK.
var DefaultAuth0ClientInfo = &Auth0ClientInfo{
	Name:    "go-auth0",
	Version: auth0.Version,
	Env: map[string]string{
		"go": runtime.Version(),
	},
}

// RetryOptions defines the retry rules that should be followed by the SDK when making requests.
type RetryOptions struct {
	MaxRetries int
	Statuses   []int
}

// IsEmpty checks whether the provided Auth0ClientInfo data is nil or has no data to allow
// short-circuiting the "Auth0-Client" header configuration.
func (r *RetryOptions) IsEmpty() bool {
	if r == nil {
		return true
	}
	return r.MaxRetries == 0 && len(r.Statuses) == 0
}

// DefaultRetryOptions is the default retry configuration used by the SDK.
// It will only retry on 429 errors and will retry a request twice.
var DefaultRetryOptions = RetryOptions{
	MaxRetries: 2,
	Statuses:   []int{http.StatusTooManyRequests},
}

// RoundTripFunc is an adapter to allow the use of ordinary functions as HTTP
// round trips.
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
func (rf RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rf(req)
}

// RetriesTransport wraps base transport with retry functionality.
//
// This transport will retry in the following circumstances:
// Total retries is less than the configured amount
// AND
// The configuration specifies to retry on the status OR the error.
func RetriesTransport(base http.RoundTripper, r RetryOptions) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}

	// Configure a retry transport that will retry if:
	// Total retries is less than the configured amount
	// AND
	// The configuration specifies to retry on the status OR the error
	tr := rehttp.NewTransport(
		base,
		rehttp.RetryAll(
			rehttp.RetryMaxRetries(r.MaxRetries),
			rehttp.RetryAny(
				rehttp.RetryStatuses(r.Statuses...),
				rehttp.RetryIsErr(retryErrors),
			),
		),
		backoffDelay(),
	)

	return tr
}

// Matches the error returned by net/http when the configured number of redirects
// is exhausted.
var redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

// Matches the error returned by net/http when the TLS certificate is not trusted.
// When version 1.20 is the minimum supported this can be replaced by tls.CertificateVerificationError.
var certVerificationErrorRe = regexp.MustCompile(`certificate is not trusted`)

// retryErrors checks whether the error returned is a potentially recoverable
// error and returns true if it is.
// By default the errors retried are too many redirects and unknown cert.
func retryErrors(err error) bool {
	if err == nil {
		return false
	}

	// Too many redirects.
	if redirectsErrorRe.MatchString(err.Error()) {
		return false
	}

	// These two checks handle a bad certificate error across our supported versions.
	if certVerificationErrorRe.MatchString(err.Error()) {
		return false
	}
	if ok := errors.As(err, &x509.UnknownAuthorityError{}); ok {
		return false
	}

	// Retry other errors as they are most likely recoverable.
	return true
}

// backoffDelay implements an exponential backoff with jitter and handles rate limiting.
func backoffDelay() rehttp.DelayFn {
	prng := rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404: Random generator
	const (
		minDelay  = 1 * time.Second
		maxDelay  = 10 * time.Second
		baseDelay = 1 * time.Second
	)

	return func(attempt rehttp.Attempt) time.Duration {
		// Calculate exponential backoff with jitter
		expBackoff := time.Duration(float64(baseDelay) * math.Pow(2, float64(attempt.Index)))
		jitter := time.Duration(prng.Float64() * float64(baseDelay))
		wait := expBackoff + jitter

		// Clamp the delay within min and max bounds
		if wait < minDelay {
			wait = minDelay
		} else if wait > maxDelay {
			wait = maxDelay
		}

		// If response is nil or not a 429 status, return computed delay
		if attempt.Response == nil || attempt.Response.StatusCode != http.StatusTooManyRequests {
			return wait
		}

		// Check Retry-After header (RFC 7231)
		if retryAfter := attempt.Response.Header.Get("Retry-After"); retryAfter != "" {
			if seconds, err := strconv.Atoi(retryAfter); err == nil {
				retryAfterDuration := time.Duration(seconds) * time.Second
				// Add 25% Padding to beat caching
				retryAfterDuration = time.Duration(float64(retryAfterDuration) * 1.25)

				if retryAfterDuration > maxDelay {
					return maxDelay
				}
				if retryAfterDuration < minDelay {
					return minDelay
				}
				return retryAfterDuration
			}
			if date, err := http.ParseTime(retryAfter); err == nil {
				retryAfterDuration := time.Until(date)
				// Add 25% Padding to beat caching
				retryAfterDuration = time.Duration(float64(retryAfterDuration) * 1.25)

				if retryAfterDuration > maxDelay {
					return maxDelay
				}
				if retryAfterDuration < minDelay {
					return minDelay
				}
				return retryAfterDuration
			}
		}

		// Fallback to X-RateLimit-Reset if Retry-After is unavailable
		if resetAt, err := strconv.ParseInt(attempt.Response.Header.Get("X-RateLimit-Reset"), 10, 64); err == nil {
			delay := time.Duration(resetAt-time.Now().Unix()) * time.Second
			// Add 25% Padding to beat caching
			delay = time.Duration(float64(delay) * 1.25)

			if delay > maxDelay {
				return maxDelay
			}
			if delay < minDelay {
				return minDelay
			}
			return delay
		}

		return wait
	}
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

// Auth0ClientInfoTransport wraps base transport with a customized "Auth0-Client" header.
func Auth0ClientInfoTransport(base http.RoundTripper, auth0ClientInfo *Auth0ClientInfo) (http.RoundTripper, error) {
	if base == nil {
		base = http.DefaultTransport
	}

	auth0ClientJSON, err := json.Marshal(auth0ClientInfo)
	if err != nil {
		return nil, err
	}

	auth0ClientEncoded := base64.StdEncoding.EncodeToString(auth0ClientJSON)

	return RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("Auth0-Client", auth0ClientEncoded)
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

// WithRetries configures the retry transports on the http client used.
func WithRetries(r RetryOptions) Option {
	return func(c *http.Client) {
		if !r.IsEmpty() {
			c.Transport = RetriesTransport(c.Transport, r)
		}
	}
}

// WithUserAgent configures the client to overwrite the user agent header.
func WithUserAgent(userAgent string) Option {
	return func(c *http.Client) {
		c.Transport = UserAgentTransport(c.Transport, userAgent)
	}
}

// WithAuth0ClientInfo configures the client to overwrite the "Auth0-Client" header.
func WithAuth0ClientInfo(auth0ClientInfo *Auth0ClientInfo) Option {
	return func(c *http.Client) {
		if auth0ClientInfo.IsEmpty() {
			return
		}
		transport, err := Auth0ClientInfoTransport(c.Transport, auth0ClientInfo)
		if err != nil {
			return
		}
		c.Transport = transport
	}
}

// WrapWithTokenSource wraps the base client with transports that enable OAuth2 authentication.
func WrapWithTokenSource(base *http.Client, tokenSource oauth2.TokenSource, options ...Option) *http.Client {
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

	return Wrap(client, options...)
}

// Wrap the base client with just the internal transports.
func Wrap(base *http.Client, options ...Option) *http.Client {
	if base == nil {
		base = http.DefaultClient
	}

	for _, option := range options {
		option(base)
	}
	return base
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

// OAuth2ClientCredentialsPrivateKeyJwt sets the oauth2
// client credentials with Private Key JWT authentication.
func OAuth2ClientCredentialsPrivateKeyJwt(ctx context.Context, uri, clientAssertionSigningAlg, clientAssertionSigningKey, clientID string) oauth2.TokenSource {
	audience := uri + "/api/v2/"
	return OAuth2ClientCredentialsPrivateKeyJwtAndAudience(ctx, uri, clientAssertionSigningAlg, clientAssertionSigningKey, clientID, audience)
}

// OAuth2ClientCredentialsPrivateKeyJwtAndAudience sets the oauth2
// client credentials with Private Key JWT authentication
// with a custom audience.
func OAuth2ClientCredentialsPrivateKeyJwtAndAudience(
	ctx context.Context,
	uri,
	clientAssertionSigningAlg,
	clientAssertionSigningKey,
	clientID,
	audience string,
) oauth2.TokenSource {
	return newPrivateKeyJwtTokenSource(ctx, uri, clientAssertionSigningAlg, clientAssertionSigningKey, clientID, audience)
}

// StaticToken sets a static token to be used for oauth2.
func StaticToken(token string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
}

func determineAlg(alg string) (jwa.SignatureAlgorithm, error) {
	switch alg {
	case "RS256":
		return jwa.RS256, nil
	case "RS384":
		return jwa.RS384, nil
	case "PS256":
		return jwa.PS256, nil
	default:
		return "", fmt.Errorf("Unsupported client assertion algorithm \"%s\" provided", alg)
	}
}

// CreateClientAssertion creates a client assertion for the provided client ID and domain.
func CreateClientAssertion(alg jwa.SignatureAlgorithm, signingKey, clientID, domain string) (string, error) {
	key, err := jwk.ParseKey([]byte(signingKey), jwk.WithPEM(true))
	if err != nil {
		return "", err
	}

	token, err := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Subject(clientID).
		JwtID(uuid.New().String()).
		Issuer(clientID).
		Claim("aud", domain).
		Expiration(time.Now().Add(2 * time.Minute)).
		Build()
	if err != nil {
		return "", err
	}

	b, err := jwt.Sign(token, jwt.WithKey(alg, key))
	if err != nil {
		return "", err
	}

	return string(b), nil
}
