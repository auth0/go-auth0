package management

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0/internal/client"
)

var (
	domain                = os.Getenv("AUTH0_DOMAIN")
	clientID              = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret          = os.Getenv("AUTH0_CLIENT_SECRET")
	debug                 = os.Getenv("AUTH0_DEBUG")
	httpRecordings        = os.Getenv("AUTH0_HTTP_RECORDINGS")
	httpRecordingsEnabled = false
	api                   = &Management{}
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
		s := httptest.NewServer(h)

		m, err := New(
			s.URL,
			WithInsecure(),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(context.Background(), "123")

		assert.NoError(t, err)
	})

	t.Run("Allows passing custom Auth0ClientInfo", func(t *testing.T) {
		customClient := client.Auth0ClientInfo{Name: "test-client", Version: "1.0.0"}

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			assert.Equal(t, "eyJuYW1lIjoidGVzdC1jbGllbnQiLCJ2ZXJzaW9uIjoiMS4wLjAifQ==", header)
		})
		s := httptest.NewServer(h)

		m, err := New(
			s.URL,
			WithInsecure(),
			WithAuth0ClientInfo(customClient),
		)
		assert.NoError(t, err)

		_, err = m.User.Read(context.Background(), "123")

		assert.NoError(t, err)
	})

	t.Run("Allows disabling Auth0ClientInfo", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	t.Run("Allows passing extra env info with custom client", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Auth0-Client")
			assert.Equal(t, "eyJuYW1lIjoidGVzdC1jbGllbnQiLCJ2ZXJzaW9uIjoiMS4wLjAiLCJlbnYiOnsiZm9vIjoiYmFyIn19", header)
		})
		s := httptest.NewServer(h)
		customClient := client.Auth0ClientInfo{Name: "test-client", Version: "1.0.0"}

		m, err := New(
			s.URL,
			WithInsecure(),
			WithAuth0ClientInfo(customClient),
			WithAuth0ClientEnvEntry("foo", "bar"),
		)
		assert.NoError(t, err)
		_, err = m.User.Read(context.Background(), "123")
		assert.NoError(t, err)
	})

	t.Run("Handles when client info has been disabled", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func TestRateLimit(t *testing.T) {
	start := time.Now()
	first := true

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if first {
			w.WriteHeader(http.StatusTooManyRequests)
			first = !first
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

	elapsed := time.Since(start)
	assert.Greater(t, elapsed, 500*time.Millisecond)
	assert.NoError(t, err)
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
