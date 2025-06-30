package management

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/internal/client"
)

var (
	domain                string
	clientID              string
	clientSecret          string
	debug                 string
	httpRecordings        string
	httpRecordingsEnabled = false
	api                   = &Management{}
)

func envVarEnabled(envVar string) bool {
	return envVar == "true" || envVar == "1" || envVar == "on"
}

func TestMain(m *testing.M) {
	_ = godotenv.Load("./../.env", ".env")

	domain = os.Getenv("AUTH0_DOMAIN")
	clientID = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	debug = os.Getenv("AUTH0_DEBUG")
	httpRecordings = os.Getenv("AUTH0_HTTP_RECORDINGS")

	httpRecordingsEnabled = envVarEnabled(httpRecordings)

	initializeTestClient()

	code := m.Run()
	os.Exit(code)
}

func initializeTestClient() {
	var err error

	api, err = New(
		domain,
		WithClientCredentials(context.Background(), clientID, clientSecret),
		WithDebug(envVarEnabled(debug)),
	)
	if err != nil {
		log.Fatal("failed to initialize the api client")
	}
}

func TestNew(t *testing.T) {
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

func TestOptionFields(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	IncludeFields("foo", "bar").apply(r)

	v := r.URL.Query()

	fields := v.Get("fields")
	assert.Equal(t, "foo,bar", fields)

	includeFields := v.Get("include_fields")
	assert.Equal(t, "true", includeFields)

	ExcludeFields("foo", "bar").apply(r)

	includeFields = v.Get("include_fields")
	assert.Equal(t, "true", includeFields)
}

func TestOptionPage(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	Page(3).apply(r)
	PerPage(10).apply(r)

	v := r.URL.Query()

	page := v.Get("page")
	assert.Equal(t, "3", page)

	perPage := v.Get("per_page")
	assert.Equal(t, "10", perPage)
}

func TestOptionTotals(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	IncludeTotals(true).apply(r)

	v := r.URL.Query()

	includeTotals := v.Get("include_totals")
	assert.Equal(t, "true", includeTotals)
}

func TestOptionFrom(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	From("abc").apply(r)
	Take(10).apply(r)

	v := r.URL.Query()

	from := v.Get("from")
	assert.Equal(t, "abc", from)

	take := v.Get("take")
	assert.Equal(t, "10", take)
}

func TestOptionParameter(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	Parameter("foo", "123").apply(r)
	Parameter("bar", "xyz").apply(r)

	v := r.URL.Query()

	foo := v.Get("foo")
	assert.Equal(t, "123", foo)

	bar := v.Get("bar")
	assert.Equal(t, "xyz", bar)
}

func TestOptionDefaults(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	applyListDefaults([]RequestOption{
		PerPage(20),          // This should be persisted (default is 50).
		IncludeTotals(false), // This should be persisted (default is true).
	}).apply(r)

	v := r.URL.Query()

	perPage := v.Get("per_page")
	assert.Equal(t, "20", perPage)

	includeTotals := v.Get("include_totals")
	assert.Equal(t, "false", includeTotals)
}

func TestOptionCheckpointDefaults(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	applyListCheckpointDefaults([]RequestOption{
		Take(20),    // This should be persisted (default is 50).
		From("abc"), // This should be persisted (default is nil).
	}).apply(r)

	v := r.URL.Query()

	take := v.Get("take")
	assert.Equal(t, "20", take)

	from := v.Get("from")
	assert.Equal(t, "abc", from)
}

func TestOptionSort(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	Sort("name:-1").apply(r)

	v := r.URL.Query()

	sort := v.Get("sort")
	assert.Equal(t, "name:-1", sort)
}

func TestOptionHeader(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	Header("foo", "bar").apply(r)

	v := r.Header.Get("foo")
	assert.Equal(t, "bar", v)
}

func TestOptionBody(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	body := []byte("fooo")
	Body(body).apply(r)

	v, err := io.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, v)
}

func TestStringify(t *testing.T) {
	expected := `{
  "foo": "bar"
}`

	v := struct {
		Foo string `json:"foo"`
	}{
		"bar",
	}

	s := Stringify(v)
	assert.Equal(t, expected, s)
}

func TestRequestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the request.

	err := api.Request(ctx, "GET", "/", nil)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestRequestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(50 * time.Millisecond) // Delay until the deadline is exceeded.

	err := api.Request(ctx, "GET", "/", nil)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestNew_WithInsecure(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v2/users/123":
			w.Write([]byte(`{"user_id":"123"}`))
		default:
			http.NotFound(w, r)
		}
	})
	s := httptest.NewServer(h)

	m, err := New(s.URL, WithInsecure())
	assert.NoError(t, err)

	u, err := m.User.Read(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, "123", u.GetID())
}

func TestManagement_URI(t *testing.T) {
	var testCases = []struct {
		name     string
		given    []string
		expected string
	}{
		{
			name:     "encodes regular user_id",
			given:    []string{"users", "1234"},
			expected: "https://" + domain + "/api/v2/users/1234",
		},
		{
			name:     "encodes a user_id with a space",
			given:    []string{"users", "123 4"},
			expected: "https://" + domain + "/api/v2/users/123%204",
		},
		{
			name:     "encodes a user_id with a |",
			given:    []string{"users", "auth0|12345678"},
			expected: "https://" + domain + "/api/v2/users/auth0%7C12345678",
		},
		{
			name:     "encodes a user_id with a | and /",
			given:    []string{"users", "auth0|1234/5678"},
			expected: "https://" + domain + "/api/v2/users/auth0%7C1234%2F5678",
		},
		{
			name:     "encodes a user_id with a /",
			given:    []string{"users", "anotherUserId/secret"},
			expected: "https://" + domain + "/api/v2/users/anotherUserId%2Fsecret",
		},
		{
			name:     "encodes a user_id with a percentage",
			given:    []string{"users", "anotherUserId/secret%23"},
			expected: "https://" + domain + "/api/v2/users/anotherUserId%2Fsecret%2523",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := api.URI(testCase.given...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
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
		s := httptest.NewServer(h)

		m, err := New(
			s.URL,
			WithInsecure(),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(context.Background(), "123")

		assert.NoError(t, err)
	})

	t.Run("Allows disabling Auth0ClientInfo", func(t *testing.T) {
		h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			rawHeader := r.Header.Get("Auth0-Client")
			assert.Empty(t, rawHeader)
		})
		s := httptest.NewServer(h)

		m, err := New(
			s.URL,
			WithInsecure(),
			WithNoAuth0ClientInfo(),
		)
		assert.NoError(t, err)
		_, err = m.User.Read(context.Background(), "123")
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
		s := httptest.NewServer(h)

		m, err := New(
			s.URL,
			WithInsecure(),
			WithAuth0ClientEnvEntry("foo", "bar"),
		)
		assert.NoError(t, err)
		_, err = m.User.Read(context.Background(), "123")
		assert.NoError(t, err)
	})

	t.Run("Handles when client info has been disabled", func(t *testing.T) {
		h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			assert.Equal(t, "", header)
		})
		s := httptest.NewServer(h)

		m, err := New(
			s.URL,
			WithInsecure(),
			WithNoAuth0ClientInfo(),
			WithAuth0ClientEnvEntry("foo", "bar"),
		)
		assert.NoError(t, err)
		_, err = m.User.Read(context.Background(), "123")
		assert.NoError(t, err)
	})
}

func TestApiCallContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the request.

	m, err := New(
		"http://localhost:8080",
		WithInsecure(),
	)

	assert.NoError(t, err)

	_, err = m.User.Read(ctx, "123")
	assert.ErrorIs(t, err, context.Canceled)
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

		s := httptest.NewServer(h)
		defer s.Close()

		m, err := New(
			s.URL,
			WithInsecure(),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(context.Background(), "123")
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

		s := httptest.NewServer(h)
		defer s.Close()

		m, err := New(
			s.URL,
			WithInsecure(),
			WithRetries(1, []int{http.StatusBadGateway}),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(context.Background(), "123")
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

		s := httptest.NewServer(h)
		defer s.Close()

		m, err := New(
			s.URL,
			WithInsecure(),
			WithNoRetries(),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(context.Background(), "123")
		assert.Equal(t, http.StatusBadGateway, err.(*managementError).StatusCode)
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

		s := httptest.NewServer(h)
		defer s.Close()

		m, err := New(
			s.URL,
			WithInsecure(),
			WithRetries(3, []int{http.StatusBadGateway}),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(ctx, "123")
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

		s := httptest.NewServer(h)
		defer s.Close()

		m, err := New(
			s.URL,
			WithInsecure(),
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

		_, err = m.User.Read(ctx, "123")
		assert.ErrorIs(t, err, context.Canceled)
		assert.Equal(t, int64(2), i.Load()) // 1 request should have been made before the context times out
	})
}

func TestApiCallContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(50 * time.Millisecond) // Delay until the deadline is exceeded.

	m, err := New(
		"http://localhost:8080",
		WithInsecure(),
	)
	assert.NoError(t, err)

	_, err = m.User.Read(ctx, "123")
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestCustomDomainHeader(t *testing.T) {
	t.Run("Option applies custom domain header", func(t *testing.T) {
		tests := []struct {
			name          string
			inputDomain   string
			expectedValue string
		}{
			{"sets valid custom domain", "my.custom.domain", "my.custom.domain"},
			{"sets empty custom domain", "", ""},
			{"sets domain with subdomain", "sub.domain.example.com", "sub.domain.example.com"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r, _ := http.NewRequest("GET", "/", nil)
				CustomDomainHeader(tt.inputDomain).apply(r)
				got := r.Header.Get("Auth0-Custom-Domain")
				assert.Equal(t, tt.expectedValue, got)
			})
		}
	})

	t.Run("Middleware applies custom domain header only on whitelisted paths", func(t *testing.T) {
		globalDomain := "global.custom.domain"
		whitelistedPath := "/api/v2/users"
		nonWhitelistedPath := "/api/v2/clients"

		t.Run("applies on whitelisted endpoint", func(t *testing.T) {
			var actualHeader string

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == whitelistedPath {
					actualHeader = r.Header.Get("Auth0-Custom-Domain")

					w.Write([]byte(`{"users":[]}`))

					return
				}

				http.NotFound(w, r)
			}))
			defer server.Close()

			m, err := New(
				server.URL,
				WithInsecure(),
				WithCustomDomainHeader(globalDomain),
			)
			require.NoError(t, err)

			_, err = m.User.List(context.Background())
			require.NoError(t, err)
			assert.Equal(t, globalDomain, actualHeader)
		})

		t.Run("does not apply on non-whitelisted endpoint", func(t *testing.T) {
			var actualHeader string

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == nonWhitelistedPath {
					actualHeader = r.Header.Get("Auth0-Custom-Domain")

					w.Write([]byte(`{"clients":[]}`))

					return
				}

				http.NotFound(w, r)
			}))
			defer server.Close()

			m, err := New(
				server.URL,
				WithInsecure(),
				WithCustomDomainHeader(globalDomain),
			)
			require.NoError(t, err)

			_, err = m.Client.List(context.Background())
			require.NoError(t, err)
			assert.Empty(t, actualHeader, "Header should not be set for non-whitelisted endpoint")
		})
	})
}

func TestApplyCustomDomainHeader_InvalidURI(t *testing.T) {
	m := &Management{
		customDomainHeader: "invalid.test.domain",
	}

	req, _ := http.NewRequest("GET", "/", nil)

	// Pass an invalid URI to trigger the error path
	m.applyCustomDomainHeader(req, "http://%41:8080/") // %41 is 'A', but in host:port, this is invalid

	// Header should not be set
	assert.Empty(t, req.Header.Get("Auth0-Custom-Domain"))
}
