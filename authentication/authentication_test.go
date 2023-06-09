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

	code := m.Run()
	os.Exit(code)
}

func initializeTestClient() {
	var err error

	authAPI, err = New(
		domain,
		WithClientID(clientID),
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
		_, err := New(domain)
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
		"http://localhost:8080",
	)

	assert.NoError(t, err)

	_, err = a.Database.SignUp(ctx, database.SignUpRequest{
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
		"http://localhost:8080",
	)

	assert.NoError(t, err)

	_, err = a.Database.SignUp(ctx, database.SignUpRequest{
		Username: "test",
		Password: "test",
	})
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}
