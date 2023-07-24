package authentication

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"runtime"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0/authentication/database"
	"github.com/auth0/go-auth0/internal/client"
)

var (
	domain                = os.Getenv("AUTH0_DOMAIN")
	clientID              = os.Getenv("AUTH0_AUTH_CLIENT_ID")
	clientSecret          = os.Getenv("AUTH0_AUTH_CLIENT_SECRET")
	httpRecordings        = os.Getenv("AUTH0_HTTP_RECORDINGS")
	httpRecordingsEnabled = false
	authAPI               = &Authentication{}
)

func envVarEnabled(envVar string) bool {
	return envVar == "true" || envVar == "1" || envVar == "on"
}

func TestMain(m *testing.M) {
	httpRecordingsEnabled = envVarEnabled(httpRecordings)

	// If we're running in standard `make test` then set a specific clientID and clientSecret
	// to ensure that we exactly match the body parameters
	if httpRecordingsEnabled && domain == "go-auth0-dev.eu.auth0.com" {
		clientID = "test-client_id"
		clientSecret = "test-client_secret"
	}

	initializeTestClient()

	code := m.Run()
	os.Exit(code)
}

func initializeTestClient() {
	var err error

	authAPI, err = New(
		context.Background(),
		domain,
		WithClientID(clientID),
		WithIDTokenSigningAlg("HS256"),
	)
	if err != nil {
		log.Fatal("failed to initialize the auth api client")
	}
}

func TestAuthenticationNew(t *testing.T) {
	for _, domain := range []string{
		"example.com ",
		" example.com",
		" example.com ",
		"%2Fexample.com",
		" a.b.c.example.com",
	} {
		_, err := New(context.Background(), domain)
		assert.Errorf(t, err, "expected New to fail with domain %q", domain)
	}
}

func TestAuthenticatioNOptionHeader(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	Header("foo", "bar").apply(r, url.Values{})

	h := r.Header.Get("foo")
	assert.Equal(t, h, "bar")
}

func TestAuthenticationRequestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the request.

	err := authAPI.Request(ctx, "GET", "/", nil, nil)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestAuthenticationRequestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(50 * time.Millisecond) // Delay until the deadline is exceeded.

	err := authAPI.Request(ctx, "GET", "/", nil, nil)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestAuthenticationApiCallContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the request.

	a, err := New(
		context.Background(),
		"http://localhost:8080",
		WithIDTokenSigningAlg("HS256"),
	)

	assert.NoError(t, err)

	_, err = a.Database.Signup(ctx, database.SignupRequest{
		Username: "test",
		Password: "test",
	})
	assert.ErrorIs(t, err, context.Canceled)
}

func TestAuthenticationApiCallContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(50 * time.Millisecond) // Delay until the deadline is exceeded.

	a, err := New(
		context.Background(),
		"http://localhost:8080",
		WithIDTokenSigningAlg("HS256"),
	)

	assert.NoError(t, err)

	_, err = a.Database.Signup(ctx, database.SignupRequest{
		Username: "test",
		Password: "test",
	})
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestUserInfo(t *testing.T) {
	configureHTTPTestRecordings(t)

	user, err := authAPI.UserInfo(context.Background(), "test-access-token")

	assert.NoError(t, err)
	assert.Equal(t, "test-sub", user.Sub)
	assert.Equal(t, "test-value", user.AdditionalClaims["unknown-param"])
}

func TestAuth0Client(t *testing.T) {
	t.Run("Defaults to the default data", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			auth0ClientDecoded, err := base64.StdEncoding.DecodeString(header)
			assert.NoError(t, err)

			var auth0Client client.Auth0ClientInfo
			err = json.Unmarshal(auth0ClientDecoded, &auth0Client)

			assert.NoError(t, err)
			assert.Equal(t, "go-auth0", auth0Client.Name)
			assert.Equal(t, "latest", auth0Client.Version)
			assert.Equal(t, runtime.Version(), auth0Client.Env["go"])
		})
		s := httptest.NewTLSServer(h)
		t.Cleanup(func() {
			s.Close()
		})

		a, err := New(
			context.Background(),
			s.URL,
			WithClient(s.Client()),
			WithIDTokenSigningAlg("HS256"),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(context.Background(), "123")

		assert.NoError(t, err)
	})

	t.Run("Allows disabling Auth0ClientInfo", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawHeader := r.Header.Get("Auth0-Client")
			assert.Empty(t, rawHeader)
		})
		s := httptest.NewTLSServer(h)
		t.Cleanup(func() {
			s.Close()
		})

		a, err := New(
			context.Background(),
			s.URL,
			WithClient(s.Client()),
			WithIDTokenSigningAlg("HS256"),
			WithNoAuth0ClientInfo(),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(context.Background(), "123")
		assert.NoError(t, err)
	})

	t.Run("Allows passing extra env info", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			auth0ClientDecoded, err := base64.StdEncoding.DecodeString(header)
			assert.NoError(t, err)

			var auth0Client client.Auth0ClientInfo
			err = json.Unmarshal(auth0ClientDecoded, &auth0Client)

			assert.NoError(t, err)
			assert.Equal(t, "go-auth0", auth0Client.Name)
			assert.Equal(t, "latest", auth0Client.Version)
			assert.Equal(t, runtime.Version(), auth0Client.Env["go"])
			assert.Equal(t, "bar", auth0Client.Env["foo"])
		})
		s := httptest.NewTLSServer(h)
		t.Cleanup(func() {
			s.Close()
		})

		a, err := New(
			context.Background(),
			s.URL,
			WithClient(s.Client()),
			WithAuth0ClientEnvEntry("foo", "bar"),
			WithIDTokenSigningAlg("HS256"),
		)
		assert.NoError(t, err)
		_, err = a.UserInfo(context.Background(), "123")
		assert.NoError(t, err)
	})

	t.Run("Handles when client info has been disabled", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			assert.Equal(t, "", header)
		})
		s := httptest.NewTLSServer(h)
		t.Cleanup(func() {
			s.Close()
		})

		a, err := New(
			context.Background(),
			s.URL,
			WithClient(s.Client()),
			WithNoAuth0ClientInfo(),
			WithAuth0ClientEnvEntry("foo", "bar"),
			WithIDTokenSigningAlg("HS256"),
		)
		assert.NoError(t, err)
		_, err = a.UserInfo(context.Background(), "123")
		assert.NoError(t, err)
	})
}

func TestRetries(t *testing.T) {
	t.Run("Default retry logic", func(t *testing.T) {
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

		s := httptest.NewTLSServer(h)
		defer s.Close()

		a, err := New(
			context.Background(),
			s.URL,
			WithIDTokenSigningAlg("HS256"),
			WithClient(s.Client()),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(context.Background(), "123")
		assert.NoError(t, err)

		elapsed := time.Since(start).Milliseconds()
		assert.Greater(t, elapsed, int64(250))
		assert.Equal(t, 2, i)
	})

	t.Run("Custom retry logic", func(t *testing.T) {
		start := time.Now()
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			i++
			if i < 2 {
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			w.WriteHeader(http.StatusOK)
		})

		s := httptest.NewTLSServer(h)
		defer s.Close()

		a, err := New(
			context.Background(),
			s.URL,
			WithIDTokenSigningAlg("HS256"),
			WithClient(s.Client()),
			WithRetries(1, []int{http.StatusBadGateway}),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(context.Background(), "123")
		assert.NoError(t, err)
		assert.Equal(t, 2, i)

		elapsed := time.Since(start)
		assert.Greater(t, elapsed, 250*time.Millisecond)
		assert.NoError(t, err)
	})

	t.Run("Disabling retries", func(t *testing.T) {
		i := 0

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			i++
			w.WriteHeader(http.StatusBadGateway)
		})

		s := httptest.NewTLSServer(h)
		defer s.Close()

		a, err := New(
			context.Background(),
			s.URL,
			WithIDTokenSigningAlg("HS256"),
			WithClient(s.Client()),
			WithNoRetries(),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(context.Background(), "123")
		assert.Equal(t, http.StatusBadGateway, err.(*authenticationError).StatusCode)
		assert.Equal(t, 1, i)
	})

	t.Run("Retry and context", func(t *testing.T) {
		i := 0
		ctx, cancel := context.WithCancel(context.Background())

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			i++
			cancel()
			w.WriteHeader(http.StatusBadGateway)
		})

		s := httptest.NewTLSServer(h)
		defer s.Close()

		a, err := New(
			context.Background(),
			s.URL,
			WithIDTokenSigningAlg("HS256"),
			WithClient(s.Client()),
			WithRetries(3, []int{http.StatusBadGateway}),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(ctx, "123")
		assert.ErrorIs(t, err, context.Canceled)
		assert.Equal(t, 1, i) // 1 request should have been made before the context times out
	})
}
