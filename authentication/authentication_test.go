package authentication

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0/authentication/database"
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
	initializeTestClient()

	// If we're running in standard `make test` then set a clientSecret to ensure tests work
	if httpRecordingsEnabled && clientSecret == "" && domain == "go-auth0-dev.eu.auth0.com" {
		clientSecret = "test-client-secret"
	}

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
