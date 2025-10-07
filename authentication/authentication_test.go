package authentication

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/v2/management"

	"github.com/auth0/go-auth0/v2"
	"github.com/auth0/go-auth0/v2/authentication/database"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
	"github.com/auth0/go-auth0/v2/internal/client"
	managementClient "github.com/auth0/go-auth0/v2/management/client"
	"github.com/auth0/go-auth0/v2/management/option"
)

var (
	domain                string
	clientID              string
	clientSecret          string
	mgmtClientID          string
	mgmtClientSecret      string
	httpRecordings        string
	httpRecordingsEnabled = false
	authAPI               = &Authentication{}
	mgmtAPI               = &managementClient.Management{}
	jwtPublicKey          = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8foXPIpkeLKAVfVg/W0X
steFas2XwrxAGG0lnLS3mc/cYc/pD/plsR779O8It/2YmHFWIDmCIcW57boDae/K
AhVBLHUa3ato7h5agbY2mKSDUEjqjWilAbdyUZDz8US8ocAmehyVWMuVqeGxunPH
opm4JQ2OGcE31MbtcJN07zCa/R/LUi8KMeuujQ6cceIGupCdOsK6JkoUB2wkvFpU
CiOwqTG51Eq4DSTukDr7tDfe0s5e1MBxVUxLkrw7zBrlDxPgZ+M260FUlRqOKKsk
3IUke/vcQac6t+js1zOs0mapqLybkszGwl2wY0JaOLOtcL5zi4U9w/GeOHOVROdn
6wIDAQAB
-----END PUBLIC KEY-----`
	jwtPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDx+hc8imR4soBV
9WD9bRey14VqzZfCvEAYbSWctLeZz9xhz+kP+mWxHvv07wi3/ZiYcVYgOYIhxbnt
ugNp78oCFUEsdRrdq2juHlqBtjaYpINQSOqNaKUBt3JRkPPxRLyhwCZ6HJVYy5Wp
4bG6c8eimbglDY4ZwTfUxu1wk3TvMJr9H8tSLwox666NDpxx4ga6kJ06wromShQH
bCS8WlQKI7CpMbnUSrgNJO6QOvu0N97Szl7UwHFVTEuSvDvMGuUPE+Bn4zbrQVSV
Go4oqyTchSR7+9xBpzq36OzXM6zSZqmovJuSzMbCXbBjQlo4s61wvnOLhT3D8Z44
c5VE52frAgMBAAECggEALCx3qXmqNc6AVzDgb+NGfEOT+5dkqQwst0jVoPHswouL
s998sIoJnngFjwVEFjKZdNrb2i4lb3zlIFzg2qoHurGeoDsQmH7+PNoVs7BL7zm5
LyLgjsgXt2SB3hoULmtZ9D1byNcG/JrNy6GEDIGuZCSj1T/QPStkwdc+6VpB8pgW
E8D7jCt40Tik2neYQkDnY775kGAHGWEqpdPCwm+KOnuE1fHx/jk38lmUgYNjKq0h
JK6Ncjen1X+ZsYfGx4dALWG4cqo3lE0YXXuHuvjJV3aVfzH8t7W4fuZ4+8xvdhhV
F4br5FimWLbTe2qT4lSpadkbLm3aBlSUR7eAP0BlwQKBgQD5ayZpP5OMp1zfa4hA
fM8nVUEaVLkRwFK5NChfjHGiaye2RjrnIorXMsFxXjEscgTn2Ux9CgcBhp1fTBhy
6cmhkp1talAIqLBivNQJT0YTfA+uHrHTTyMfEUgsMzPiiAg7FV7BCG6xd/nsk3yg
ZUfoXefrhq9LIHsJx7cK12VViQKBgQD4XKvwYmX5t7fZFBPd7dv5ZrcMHQnBMHd7
is3QhgyKuEgVDzKQ9SA004I9iSvcI3dE/npj31P39N5bbuvYTh4WR/SR4VvXavNG
AqUR7wm8jTlbiWEPgF9MxC24zaa07Kbxs+P8XT/7wWuijf6+baSFgxQMb80fUArv
7guKikCo0wKBgCUn3DIDoZRrfj9eQo7wyN9gKPGmO2e0kd47MeSCBI+gjOrvbWjv
UWWbjwu3b3Xiim6LhYR/EOoeRqViraW4xCvIrqEVHFUd5CDhZmj4oUTXz3It6mnD
OUUwiuLiwdD2WNuMZHA3NF5FtDqVAhTW4a5xBtKkXsq/TPT5BoCb8+GZAoGAUWAD
0gpbgTuJ2G10qPWDaq8V8Lke9haMP4VWNCmHuHfy3juRhN9cAxL+DG2CWmmgbZG3
xjtpRsgLhwfL7J6DyyceYiHltqpLNTgun7ajiQz4qx5TGAImt39bv75aDdOwS2d2
nrxq93EDdEp0Gi7QhhJRolWLbuQKAV0MmQL9dpMCgYEA5+ug3CDI/jyTHG4ZEVoG
qmIg7QoHrVEmZrvCMiFw8bbuBvoMnvu1o1zfvAkNrDfibZyxYKHzSqgeVPQShvLa
P6JCu67ieCGP8C8CMFiQhJ9n4sYGnkzkz67NpkHSzDPA6DfvG4pYuvBQRIefnhYh
IDGpghhKHMV2DAyzeM4cDU8=
-----END PRIVATE KEY-----`
)

func envVarEnabled(envVar string) bool {
	return envVar == "true" || envVar == "1" || envVar == "on"
}

func TestMain(m *testing.M) {
	_ = godotenv.Load("./../.env", ".env")

	domain = os.Getenv("AUTH0_DOMAIN")
	clientID = os.Getenv("AUTH0_AUTH_CLIENT_ID")
	clientSecret = os.Getenv("AUTH0_AUTH_CLIENT_SECRET")
	mgmtClientID = os.Getenv("AUTH0_CLIENT_ID")
	mgmtClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	httpRecordings = os.Getenv("AUTH0_HTTP_RECORDINGS")

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
		// WithIDTokenSigningAlg("HS256"),
	)
	if err != nil {
		log.Fatal("failed to initialize the auth api client")
	}

	mgmtAPI, err = managementClient.New(
		domain,
		option.WithClientCredentials(context.Background(), mgmtClientID, mgmtClientSecret),
	)

	if err != nil {
		log.Fatal("failed to initialize the management api client")
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
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	user, err := authAPI.UserInfo(context.Background(), "test-access-token")

	assert.NoError(t, err)
	assert.Equal(t, "test-sub", user.Sub)
	assert.Equal(t, "test-value", user.AdditionalClaims["unknown-param"])
}

func TestAuth0Client(t *testing.T) {
	t.Run("Defaults to the default data", func(t *testing.T) {
		h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			auth0ClientDecoded, err := base64.StdEncoding.DecodeString(header)
			assert.NoError(t, err)

			var auth0Client client.Auth0ClientInfo
			err = json.Unmarshal(auth0ClientDecoded, &auth0Client)

			assert.NoError(t, err)
			assert.Equal(t, "go-auth0", auth0Client.Name)
			assert.Equal(t, auth0.Version, auth0Client.Version)
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
		h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
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
		h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			auth0ClientDecoded, err := base64.StdEncoding.DecodeString(header)
			assert.NoError(t, err)

			var auth0Client client.Auth0ClientInfo
			err = json.Unmarshal(auth0ClientDecoded, &auth0Client)

			assert.NoError(t, err)
			assert.Equal(t, "go-auth0", auth0Client.Name)
			assert.Equal(t, auth0.Version, auth0Client.Version)
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
		h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
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

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
		assert.Equal(t, http.StatusBadGateway, err.(*Error).StatusCode)
		assert.Equal(t, 1, i)
	})

	t.Run("Retry and context", func(t *testing.T) {
		i := 0
		ctx, cancel := context.WithCancel(context.Background())

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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

	t.Run("Retry strategy", func(t *testing.T) {
		i := 0
		ctx, cancel := context.WithCancel(context.Background())

		h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
			WithRetryStrategy(RetryStrategy{
				MaxRetries:        3,
				Statuses:          []int{http.StatusBadGateway},
				PerAttemptTimeout: time.Second,
			}),
		)
		assert.NoError(t, err)

		_, err = a.UserInfo(ctx, "123")
		assert.ErrorIs(t, err, context.Canceled)
		assert.Equal(t, 1, i) // 1 request should have been made before the context times out
	})

	t.Run("Retry per request timeout", func(t *testing.T) {
		var i atomic.Int64

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			c := i.Add(1)
			t.Log(c)

			if c == 2 {
				cancel()
			}

			timer := time.NewTimer(10 * time.Second)
			select {
			case <-timer.C:
				t.Log("completed")
				w.WriteHeader(http.StatusOK)

				return
			case <-ctx.Done():
				t.Log("cancelled")
				w.WriteHeader(499)

				return
			}
		})

		s := httptest.NewTLSServer(h)
		defer s.Close()

		m, err := New(
			context.Background(),
			s.URL,
			WithIDTokenSigningAlg("HS256"),
			WithClient(s.Client()),
			WithRetryStrategy(RetryStrategy{
				MaxRetries: 10,
				Statuses: []int{
					http.StatusInternalServerError,
					http.StatusBadGateway,
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout,
				},
				PerAttemptTimeout: 5 * time.Millisecond,
			}),
		)
		assert.NoError(t, err)

		_, err = m.UserInfo(ctx, "123")
		assert.ErrorIs(t, err, context.Canceled)
		assert.Equal(t, int64(2), i.Load()) // 1 request should have been made before the context times out
	})
}

func TestWithClockTolerance(t *testing.T) {
	idTokenClientSecret := "test-client-secret"
	idTokenClientID := "test-client-id"

	var idToken string

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		tokenSet := &oauth.TokenSet{
			AccessToken: "test-access-token",
			ExpiresIn:   86400,
			IDToken:     idToken,
			TokenType:   "Bearer",
		}

		b, err := json.Marshal(tokenSet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if _, err := fmt.Fprint(w, string(b)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	s := httptest.NewTLSServer(h)

	t.Cleanup(func() {
		s.Close()
	})

	URL, err := url.Parse(s.URL)
	require.NoError(t, err)

	builder := jwt.NewBuilder().
		Issuer(s.URL + "/").
		Subject("me").
		Audience([]string{idTokenClientID}).
		Expiration(time.Now().Add(time.Hour)).
		IssuedAt(time.Now().Add(5 * time.Second))

	token, err := builder.Build()
	require.NoError(t, err)

	b, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(idTokenClientSecret)))
	require.NoError(t, err)

	idToken = string(b)

	api, err := New(
		context.Background(),
		URL.Host,
		WithClient(s.Client()),
		WithClientID(idTokenClientID),
		WithClientSecret(idTokenClientSecret),
		WithIDTokenSigningAlg("HS256"),
		WithIDTokenClockTolerance(1*time.Second), // Set a low clock tolerance to cause a failure
	)
	assert.NoError(t, err)

	_, err = api.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
		Code: "my-code",
	}, oauth.IDTokenValidationOptions{})
	assert.ErrorContains(t, err, "\"iat\" not satisfied")
}

func skipE2E(t *testing.T) {
	t.Helper()

	if !httpRecordingsEnabled {
		t.Skip("Skipped as cannot be test in E2E scenario")
	}
}

func usingRecordingResponses(t *testing.T) bool {
	t.Helper()

	return httpRecordingsEnabled && domain == "go-auth0-dev.eu.auth0.com"
}

func givenAUser(t *testing.T) userDetails {
	t.Helper()

	if !usingRecordingResponses(t) {
		user := &management.CreateUserRequestContent{
			Connection:    "Username-Password-Authentication",
			Email:         auth0.String("chuck@example.com"),
			Password:      auth0.String("Testpassword123!"),
			Username:      auth0.String("test-user"),
			EmailVerified: auth0.Bool(true),
			VerifyEmail:   auth0.Bool(false),
		}

		userCreated, err := mgmtAPI.Users.Create(context.Background(), user)
		require.NoError(t, err)

		t.Cleanup(func() {
			err := mgmtAPI.Users.Delete(context.Background(), userCreated.GetUserID())
			require.NoError(t, err)
		})
	}

	return userDetails{
		connection: "Username-Password-Authentication",
		email:      "chuck@example.com",
		password:   "Testpassword123!",
		username:   "test-user",
	}
}
