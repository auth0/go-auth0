package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWrapRateLimit(t *testing.T) {
	start := time.Now()
	first := true

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if first {
			t.Log(start.Unix())
			w.Header().Set("X-RateLimit-Limit", "1")
			w.Header().Set("X-RateLimit-Remaining", "0")
			w.Header().Set("X-RateLimit-Reset", fmt.Sprint(start.Add(time.Second).Unix()))
			w.WriteHeader(http.StatusTooManyRequests)
			first = !first
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	s := httptest.NewServer(h)
	defer s.Close()

	c := Wrap(s.Client(), StaticToken(""), WithRateLimit(), WithDebug(true))
	r, err := c.Get(s.URL)
	if err != nil {
		t.Error(err)
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("Expected status code to be %d but got %d", http.StatusOK, r.StatusCode)
	}

	elapsed := time.Since(start)
	if elapsed < time.Second {
		t.Errorf("Time since start is sooner than expected. Expected >= 1s but got %s", elapsed)
	}
}

func TestWrapUserAgent(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		if ua != UserAgent {
			t.Errorf("Expected User-Agent header to match %q but got %q", UserAgent, ua)
		}
	})

	s := httptest.NewServer(h)
	defer s.Close()

	c := Wrap(s.Client(), StaticToken(""), WithUserAgent(UserAgent))
	c.Get(s.URL)
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
			w.Write([]byte(`{"access_token":"someToken","token_type":"Bearer"}`))
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
