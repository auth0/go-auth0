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
)

func TestRetries(t *testing.T) {
	t.Run("Default config", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(250))
		assert.LessOrEqual(t, elapsed, int64(500))
		assert.Equal(t, 2, i)
	})

	t.Run("Max retries", func(t *testing.T) {
		i := 0
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

		s := httptest.NewUnstartedServer(h)
		c := WrapWithTokenSource(s.Client(), StaticToken(""), WithRetries(DefaultRetryOptions))
		_, err := c.Get(s.URL)

		assert.Error(t, err)
		elapsed := time.Since(start).Milliseconds()
		assert.GreaterOrEqual(t, elapsed, int64(750))
		assert.LessOrEqual(t, elapsed, int64(2000))
	})

	t.Run("Should not retry some errors", func(t *testing.T) {
		i := 0
		start := time.Now()
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			i++
			if i == 1 {
				t.Log(start.Unix())
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

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			i++
			if i == 1 {
				t.Log(start.Unix())
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
		assert.LessOrEqual(t, elapsed, int64(500))
		assert.Equal(t, 2, i)
	})
}

func TestWrapUserAgent(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			auth0ClientDecoded, err := base64.StdEncoding.DecodeString(header)
			assert.NoError(t, err)

			var auth0Client Auth0ClientInfo
			err = json.Unmarshal(auth0ClientDecoded, &auth0Client)
			assert.NoError(t, err)

			assert.Equal(t, "go-auth0", auth0Client.Name)
			assert.Equal(t, "latest", auth0Client.Version)
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
			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
