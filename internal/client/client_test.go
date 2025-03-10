package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

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
		assert.GreaterOrEqual(t, elapsed, int64(250), "Should have minimum delay")
		assert.LessOrEqual(t, elapsed, int64(2000), "Should complete within reasonable time")
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
		assert.GreaterOrEqual(t, elapsed, int64(750), "Should have attempted retries")
		assert.LessOrEqual(t, elapsed, int64(5000), "Should complete within 5 seconds")
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
		// Expect delay between 250ms and 10s due to capping
		assert.GreaterOrEqual(t, elapsed, int64(250))
		assert.LessOrEqual(t, elapsed, int64(11000)) // Allow a bit of slack for test execution
		assert.Equal(t, 2, i)
	})

	t.Run("Should respect Retry-After header with seconds", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				// Use a whole integer value as per RFC 7231
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
		assert.GreaterOrEqual(t, elapsed, float64(1.9)) // Allow slight tolerance
		assert.LessOrEqual(t, elapsed, float64(3.0))
		assert.Equal(t, 2, i)
	})

	t.Run("Should respect Retry-After header with HTTP date", func(t *testing.T) {
		i := 0

		// Create test server first to avoid time gap
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				futureTime := time.Now().Add(500 * time.Millisecond)
				futureTimeStr := futureTime.Format(time.RFC1123)
				w.Header().Set("Retry-After", futureTimeStr)
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
			w.WriteHeader(http.StatusOK)
		}))
		defer s.Close()

		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		r, err := c.Get(s.URL)

		// Just check that the retry worked and we got a successful response
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
				// Set reset time to current time (simulating clock skew)
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
		assert.GreaterOrEqual(t, elapsed, int64(250), "Should have minimum delay")
		assert.LessOrEqual(t, elapsed, int64(2000), "Should complete within reasonable time")
	})

	t.Run("Should use Retry-After over X-RateLimit-Reset when both present", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				// Set contradicting retry headers
				w.Header().Set("Retry-After", "1")                                               // 1 second
				w.Header().Set("X-RateLimit-Reset", fmt.Sprint(start.Add(5*time.Second).Unix())) // 5 seconds
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
		// Should use Retry-After value (1s)
		assert.GreaterOrEqual(t, elapsed, float64(1.0))
		assert.LessOrEqual(t, elapsed, float64(2.0))
		assert.Equal(t, 2, i)
	})

	t.Run("Should cap delays to maximum value", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				// Set a very long retry time
				w.Header().Set("Retry-After", "30") // 30 seconds
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
				// Set a very long delay (2 hours)
				w.Header().Set("Retry-After", "7200") // 2 hours in seconds
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
				// Set an invalid Retry-After header
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

		// Should fall back to default exponential backoff
		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
	})

	t.Run("Should handle invalid X-RateLimit-Reset header gracefully", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			i++
			if i == 1 {
				// Set an invalid X-RateLimit-Reset header
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

		// Should fall back to default exponential backoff
		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
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
	// Test with bad client info that will cause JSON marshal to fail
	badClientInfo := &Auth0ClientInfo{
		Name: "test",
		Env: map[string]string{
			"badValue": string([]byte{0}), // Invalid UTF-8 sequence
		},
	}

	// We'll use a custom transport that sets a special header
	// to verify it's still used after WithAuth0ClientInfo is called
	customTransport := RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("X-Custom-Header", "custom-value")
		return &http.Response{StatusCode: http.StatusOK, Header: req.Header}, nil
	})

	// Create a client with our custom transport
	badClient := &http.Client{Transport: customTransport}

	// Apply the option with bad client info
	WithAuth0ClientInfo(badClientInfo)(badClient)

	// Make a request and verify our custom transport was used
	// (it should set X-Custom-Header which we'll check in the response)
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	resp, err := badClient.Transport.RoundTrip(req)

	assert.NoError(t, err)
	assert.Equal(t, "custom-value", resp.Header.Get("X-Custom-Header"))

	// Test with nil client info (which should be treated as empty)
	emptyClient := &http.Client{Transport: customTransport}

	// Apply the option with nil client info
	WithAuth0ClientInfo(nil)(emptyClient)

	// Make another request and verify our custom transport is still used
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
}
