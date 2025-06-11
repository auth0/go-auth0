package client

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestRetries(t *testing.T) {
	t.Run("Default config", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
		assert.LessOrEqual(t, elapsed, int64(2000))
	})

	t.Run("Max retries", func(t *testing.T) {
		i := 0
		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i <= 5 {
				w.WriteHeader(http.StatusGatewayTimeout)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		rc := RetryOptions{MaxRetries: 5, Statuses: []int{http.StatusGatewayTimeout}}
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(rc))
		r, err := c.Get(s.URL)

		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 6, i)
		assert.NoError(t, err)
	})

	t.Run("Pass empty struct to disable", func(t *testing.T) {
		i := 0
		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i <= 2 {
				w.WriteHeader(http.StatusGatewayTimeout)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(RetryOptions{}))
		r, err := c.Get(s.URL)

		assert.Equal(t, http.StatusGatewayTimeout, r.StatusCode)
		assert.Equal(t, 1, i)
		assert.NoError(t, err)
	})

	t.Run("Should retry errors", func(t *testing.T) {
		start := time.Now()

		h := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})

		s := httptest.NewUnstartedServer(h)
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		_, err := c.Get(s.URL)

		assert.Error(t, err)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(750))
		assert.LessOrEqual(t, elapsed, int64(5000))
	})

	t.Run("Should not retry some errors", func(t *testing.T) {
		i := 0
		start := time.Now()
		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewTLSServer(h)
		c := WrapWithTokenSource(http.DefaultClient, StaticToken(""), WithRetries(DefaultRetryOptions))
		_, err := c.Get(s.URL)

		elapsed := time.Since(start).Milliseconds()

		assert.Error(t, err)
		assert.Equal(t, 0, i)
		assert.Less(t, elapsed, int64(250))
	})

	t.Run("Should ensure time is after rate limit reset when retrying", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("X-RateLimit-Limit", "1")
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("X-RateLimit-Reset", fmt.Sprint(start.Add(2*time.Second).Unix()))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, r.StatusCode)

		elapsed := time.Since(start).Seconds()
		assert.GreaterOrEqual(t, elapsed, float64(2))
		assert.LessOrEqual(t, elapsed, float64(3))
		assert.Equal(t, 2, i)
	})

	t.Run("Should not use rate limit if it is exceptionally long", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("X-RateLimit-Limit", "1")
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("X-RateLimit-Reset", fmt.Sprint(start.Add(2*time.Hour).Unix()))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, r.StatusCode)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
		assert.LessOrEqual(t, elapsed, int64(11000))
		assert.Equal(t, 2, i)
	})

	t.Run("Should respect Retry-After header with seconds", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "2")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, r.StatusCode)

		elapsed := time.Since(start).Seconds()
		assert.GreaterOrEqual(t, elapsed, float64(1.9))
		assert.LessOrEqual(t, elapsed, float64(3.0))
		assert.Equal(t, 2, i)
	})

	t.Run("Should respect Retry-After header with HTTP date", func(t *testing.T) {
		i := 0

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				futureTime := time.Now().Add(500 * time.Millisecond).In(time.FixedZone("GMT", 0))
				w.Header().Set("Retry-After", futureTime.Format(time.RFC1123))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		}))
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)
	})

	t.Run("Should handle clock skew with X-RateLimit-Reset", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("X-RateLimit-Limit", "1")
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("X-RateLimit-Reset", fmt.Sprint(start.Unix()))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
		assert.LessOrEqual(t, elapsed, int64(2000))
	})

	t.Run("Should use Retry-After over X-RateLimit-Reset when both present", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "1")
				w.Header().Set("X-RateLimit-Reset", fmt.Sprint(start.Add(5*time.Second).Unix()))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, r.StatusCode)

		elapsed := time.Since(start).Seconds()
		assert.GreaterOrEqual(t, elapsed, float64(1.0))
		assert.LessOrEqual(t, elapsed, float64(2.0))
		assert.Equal(t, 2, i)
	})

	t.Run("Should cap delays to maximum value", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "30")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)
	})

	t.Run("Should handle rate limit headers correctly", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)
	})

	t.Run("Should cap very long delays", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "7200")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)
	})

	t.Run("Should handle invalid Retry-After header gracefully", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "not-a-valid-value")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
	})

	t.Run("Should handle invalid X-RateLimit-Reset header gracefully", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("X-RateLimit-Reset", "not-a-timestamp")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
	})

	t.Run("Should enforce minimum delay with Retry-After header", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				w.Header().Set("Retry-After", "0")
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		start := time.Now()
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(1000))
		assert.LessOrEqual(t, elapsed, int64(2000))
	})

	t.Run("Should handle Retry-After with HTTP date format", func(t *testing.T) {
		i := 0

		var requestTimes []time.Time

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			requestTimes = append(requestTimes, time.Now())

			i++
			if i == 1 {
				futureTime := time.Now().Add(1 * time.Second).In(time.FixedZone("GMT", 0))
				w.Header().Set("Retry-After", futureTime.Format(time.RFC1123))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		}))
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		assert.Len(t, requestTimes, 2)

		actualWaitTime := requestTimes[1].Sub(requestTimes[0]).Milliseconds()

		minExpectedDelay := int64(1000)
		maxExpectedDelay := int64(3000)

		t.Logf("Wait time between requests: %dms", actualWaitTime)

		assert.GreaterOrEqual(t, actualWaitTime, minExpectedDelay,
			"Wait between requests should be at least %dms, got %dms",
			minExpectedDelay, actualWaitTime)

		assert.LessOrEqual(t, actualWaitTime, maxExpectedDelay,
			"Wait time shouldn't be excessive")
	})

	t.Run("Should enforce minimum delay with HTTP date Retry-After header", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				futureTime := time.Now().Add(200 * time.Millisecond).In(time.FixedZone("GMT", 0))
				w.Header().Set("Retry-After", futureTime.Format(time.RFC1123))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		start := time.Now()
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(1000))
		assert.LessOrEqual(t, elapsed, int64(2000))
	})

	t.Run("Should cap HTTP date Retry-After to maximum delay", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				futureTime := time.Now().Add(30 * time.Second).In(time.FixedZone("GMT", 0))
				w.Header().Set("Retry-After", futureTime.Format(time.RFC1123))
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		start := time.Now()
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		assert.LessOrEqual(t, elapsed, int64(11000))
	})
	t.Run("Should return calculated delay when within min and max limits", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				futureTime := time.Now().Add(5 * time.Second).UTC()
				retryAfter := futureTime.Format(time.RFC1123)
				retryAfter = strings.Replace(retryAfter, "UTC", "GMT", 1) // Ensure correct format
				t.Logf("Setting Retry-After header to: %s", retryAfter)

				w.Header().Set("Retry-After", retryAfter)
				w.WriteHeader(http.StatusTooManyRequests)

				return
			}

			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewServer(h)
		defer s.Close()

		start := time.Now()
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start).Milliseconds()
		t.Logf("Actual wait time: %dms", elapsed)

		assert.GreaterOrEqual(t, elapsed, int64(5000), "Expected wait >= 6000ms, got %dms", elapsed)
		assert.LessOrEqual(t, elapsed, int64(8000), "Expected wait <= 8000ms, got %dms", elapsed)
	})
}

func TestWrapUserAgent(t *testing.T) {
	testHandler := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		if ua != UserAgent {
			t.Errorf("Expected User-Agent header to match %q but got %q", UserAgent, ua)
		}
	})

	testServer := httptest.NewServer(testHandler)
	defer testServer.Close()

	httpClient := WrapWithTokenSource(testServer.Client(), StaticToken(""), WithUserAgent(UserAgent))
	_, err := httpClient.Get(testServer.URL)
	assert.NoError(t, err)
}

func TestOAuth2ClientCredentials(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth/token":
			// Make sure we read the body before trying to parse the form
			bodyBytes, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(`{"access_token":"defaultToken","token_type":"Bearer"}`))
			assert.NoError(t, err)
		default:
			http.NotFound(w, r)
		}
	})

	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	tokenSource := OAuth2ClientCredentials(
		context.Background(),
		testServer.URL,
		"clientID",
		"clientSecret",
	)

	token, err := tokenSource.Token()
	assert.NoError(t, err)
	assert.Equal(t, "defaultToken", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
}

func TestOAuth2ClientCredentialsAndAudience(t *testing.T) {
	const expectedAudience = "myAudience"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth/token":
			requestBody, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.Contains(t, string(requestBody), expectedAudience)

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(`{"access_token":"someToken","token_type":"Bearer"}`))
			assert.NoError(t, err)
		default:
			http.NotFound(w, r)
		}
	})
	testServer := httptest.NewServer(handler)

	tokenSource := OAuth2ClientCredentialsAndAudience(
		context.Background(),
		testServer.URL,
		"clientID",
		"clientSecret",
		expectedAudience,
	)

	token, err := tokenSource.Token()
	assert.NoError(t, err)
	assert.Equal(t, "someToken", token.AccessToken)
}

func TestOAuth2ClientCredentialsPrivateKeyJwt(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth/token":
			// Capture the raw body for debugging
			rawBody, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			r.Body = io.NopCloser(bytes.NewBuffer(rawBody))
			err = r.ParseForm()
			assert.NoError(t, err)

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(`{"access_token":"jwtToken","token_type":"Bearer"}`))
			assert.NoError(t, err)
		default:
			http.NotFound(w, r)
		}
	})

	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Generate test private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	tokenSource := OAuth2ClientCredentialsPrivateKeyJwt(
		context.Background(),
		testServer.URL,
		"test-client",
		string(privateKeyPEM),
		"RS256",
	)

	token, err := tokenSource.Token()
	assert.NoError(t, err)
	assert.Equal(t, "jwtToken", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
}

func TestOAuth2ClientCredentialsPrivateKeyJwtAndAudience(t *testing.T) {
	const customAudience = "https://custom-api.example.com"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth/token":
			// Capture the raw body for debugging
			rawBody, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			// Use the form parsing instead for x-www-form-urlencoded
			r.Body = io.NopCloser(bytes.NewBuffer(rawBody))
			err = r.ParseForm()
			assert.NoError(t, err)

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(`{"access_token":"customJwtToken","token_type":"Bearer"}`))
			assert.NoError(t, err)
		default:
			http.NotFound(w, r)
		}
	})

	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Generate test private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	tokenSource := OAuth2ClientCredentialsPrivateKeyJwtAndAudience(
		context.Background(),
		testServer.URL,
		"test-client",
		string(privateKeyPEM),
		"RS256",
		customAudience,
	)

	token, err := tokenSource.Token()
	assert.NoError(t, err)
	assert.Equal(t, "customJwtToken", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
}

func TestWrapAuth0ClientInfo(t *testing.T) {
	t.Run("Default client", func(t *testing.T) {
		testHandler := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			auth0ClientDecoded, err := base64.StdEncoding.DecodeString(header)
			assert.NoError(t, err)

			var auth0Client Auth0ClientInfo
			err = json.Unmarshal(auth0ClientDecoded, &auth0Client)
			assert.NoError(t, err)

			assert.Equal(t, "go-auth0", auth0Client.Name)
			assert.Equal(t, auth0.Version, auth0Client.Version)
			assert.Equal(t, runtime.Version(), auth0Client.Env["go"])
		})

		testServer := httptest.NewServer(testHandler)

		t.Cleanup(func() {
			testServer.Close()
		})

		httpClient := WrapWithTokenSource(testServer.Client(), StaticToken(""), WithAuth0ClientInfo(DefaultAuth0ClientInfo))
		_, err := httpClient.Get(testServer.URL)
		assert.NoError(t, err)
	})

	var testCases = []struct {
		name     string
		given    Auth0ClientInfo
		expected string
	}{
		{
			name:     "Custom client",
			given:    Auth0ClientInfo{"foo", "1.0.0", map[string]string{"os": "windows"}},
			expected: "eyJuYW1lIjoiZm9vIiwidmVyc2lvbiI6IjEuMC4wIiwiZW52Ijp7Im9zIjoid2luZG93cyJ9fQ==",
		},
		{
			name:     "Missing information",
			given:    Auth0ClientInfo{Name: "foo"},
			expected: "eyJuYW1lIjoiZm9vIiwidmVyc2lvbiI6IiJ9",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testHandler := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
				header := r.Header.Get("Auth0-Client")
				assert.Equal(t, testCase.expected, header)
			})

			testServer := httptest.NewServer(testHandler)

			t.Cleanup(func() {
				testServer.Close()
			})

			httpClient := WrapWithTokenSource(testServer.Client(), StaticToken(""), WithAuth0ClientInfo(&testCase.given))
			_, err := httpClient.Get(testServer.URL)
			assert.NoError(t, err)
		})
	}
}

func TestWrap(t *testing.T) {
	t.Run("Should handle nil client", func(t *testing.T) {
		c := Wrap(nil)
		assert.NotNil(t, c)
	})

	t.Run("Should apply multiple options", func(t *testing.T) {
		var appliedOptions int

		option1 := func(_ *http.Client) {
			appliedOptions++
		}

		option2 := func(_ *http.Client) {
			appliedOptions++
		}

		Wrap(http.DefaultClient, option1, option2)
		assert.Equal(t, 2, appliedOptions)
	})
}

func TestWithAuth0ClientInfoErrors(t *testing.T) {
	badClientInfo := &Auth0ClientInfo{
		Name: "test",
		Env: map[string]string{
			"badValue": string([]byte{0}),
		},
	}

	customTransport := RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("X-Custom-Header", "custom-value")
		return &http.Response{StatusCode: http.StatusOK, Header: req.Header}, nil
	})

	badClient := &http.Client{Transport: customTransport}

	WithAuth0ClientInfo(badClientInfo)(badClient)

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	resp, err := badClient.Transport.RoundTrip(req)

	assert.NoError(t, err)
	assert.Equal(t, "custom-value", resp.Header.Get("X-Custom-Header"))

	emptyClient := &http.Client{Transport: customTransport}

	WithAuth0ClientInfo(nil)(emptyClient)

	req2, _ := http.NewRequest("GET", "http://example.com", nil)
	resp2, err := emptyClient.Transport.RoundTrip(req2)

	assert.NoError(t, err)
	assert.Equal(t, "custom-value", resp2.Header.Get("X-Custom-Header"))
}

func TestEmptyChecks(t *testing.T) {
	t.Run("Auth0ClientInfo.IsEmpty", func(t *testing.T) {
		var nilInfo *Auth0ClientInfo

		assert.True(t, nilInfo.IsEmpty())

		emptyInfo := &Auth0ClientInfo{}
		assert.True(t, emptyInfo.IsEmpty())

		populatedInfo := &Auth0ClientInfo{Name: "test"}
		assert.False(t, populatedInfo.IsEmpty())
	})

	t.Run("RetryOptions.IsEmpty", func(t *testing.T) {
		var nilOptions *RetryOptions

		assert.True(t, nilOptions.IsEmpty())

		emptyOptions := &RetryOptions{}
		assert.True(t, emptyOptions.IsEmpty())

		populatedOptions := &RetryOptions{MaxRetries: 1}
		assert.False(t, populatedOptions.IsEmpty())
	})
}

func TestRoundTripFunc(t *testing.T) {
	called := false
	f := RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
		called = true
		return &http.Response{StatusCode: http.StatusOK}, nil
	})

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	resp, err := f.RoundTrip(req)

	assert.True(t, called)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDebugTransport(t *testing.T) {
	t.Run("With debugging disabled", func(t *testing.T) {
		base := RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusOK}, nil
		})

		transport := DebugTransport(base, false)
		req, _ := http.NewRequest("GET", "http://example.com", nil)
		resp, _ := transport.RoundTrip(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("With debugging enabled", func(t *testing.T) {
		var buf bytes.Buffer

		log.SetOutput(&buf)

		defer log.SetOutput(os.Stderr)

		base := RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("test response body")),
				Header:     http.Header{"Content-Type": []string{"text/plain"}},
			}, nil
		})

		transport := DebugTransport(base, true)
		req, _ := http.NewRequest("GET", "http://example.com", nil)
		req.Header.Set("User-Agent", "test-agent")

		resp, err := transport.RoundTrip(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		logOutput := buf.String()
		assert.Contains(t, logOutput, "GET /")
		assert.Contains(t, logOutput, "Host: example.com")
		assert.Contains(t, logOutput, "User-Agent: test-agent")
		assert.Contains(t, logOutput, "HTTP/0.0 200 OK")
		assert.Contains(t, logOutput, "Content-Type: text/plain")
		assert.Contains(t, logOutput, "test response body")
	})

	t.Run("With error response", func(t *testing.T) {
		var buf bytes.Buffer

		log.SetOutput(&buf)

		defer log.SetOutput(os.Stderr)

		expectedErr := fmt.Errorf("network error")
		base := RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return nil, expectedErr
		})

		transport := DebugTransport(base, true)
		req, _ := http.NewRequest("GET", "http://example.com", nil)

		resp, err := transport.RoundTrip(req)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, resp)

		logOutput := buf.String()
		assert.Contains(t, logOutput, "GET /")
		assert.Contains(t, logOutput, "Host: example.com")
		assert.Contains(t, logOutput, "GET")
		assert.NotContains(t, logOutput, "HTTP/0.0 200 OK")
		assert.Contains(t, logOutput, "GET")
		assert.NotContains(t, logOutput, "<")
	})
}

func TestWithDebug(t *testing.T) {
	t.Run("Enables debug transport when true", func(t *testing.T) {
		var requestReceived bool

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			requestReceived = true

			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		var buf bytes.Buffer

		log.SetOutput(&buf)

		defer log.SetOutput(os.Stderr)

		client := Wrap(http.DefaultClient, WithDebug(true))

		resp, err := client.Get(ts.URL)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.True(t, requestReceived)

		logOutput := buf.String()
		assert.Contains(t, logOutput, "GET /")
		assert.Contains(t, logOutput, "Host: 127.0.0.1:")
		assert.Contains(t, logOutput, "HTTP/1.1 200 OK")
	})

	t.Run("Keeps original transport when false", func(t *testing.T) {
		customTransport := RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusTeapot,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		})

		client := &http.Client{Transport: customTransport}

		WithDebug(false)(client)

		req, _ := http.NewRequest("GET", "http://example.com", nil)
		resp, err := client.Transport.RoundTrip(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusTeapot, resp.StatusCode)
	})
}

func TestDumpRequest(t *testing.T) {
	var buf bytes.Buffer

	log.SetOutput(&buf)

	defer log.SetOutput(os.Stderr)

	body := strings.NewReader(`{"test":"value"}`)
	req, err := http.NewRequest("POST", "https://api.example.com/test", body)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test-token")

	dumpRequest(req)

	logOutput := buf.String()
	assert.Contains(t, logOutput, "POST /test")
	assert.Contains(t, logOutput, "Host: api.example.com")
	assert.Contains(t, logOutput, "Content-Type: application/json")
	assert.Contains(t, logOutput, "Authorization: Bearer test-token")
	assert.Contains(t, logOutput, `{"test":"value"}`)
}

func TestDumpResponse(t *testing.T) {
	var buf bytes.Buffer

	log.SetOutput(&buf)

	defer log.SetOutput(os.Stderr)

	body := io.NopCloser(strings.NewReader(`{"result":"success"}`))
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       body,
	}

	dumpResponse(resp)

	logOutput := buf.String()
	assert.Contains(t, logOutput, "HTTP/0.0 200 OK")
	assert.Contains(t, logOutput, "Content-Type: application/json")
	assert.Contains(t, logOutput, `{"result":"success"}`)
}
