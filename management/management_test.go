package management

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var (
	domain                = os.Getenv("AUTH0_DOMAIN")
	clientID              = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret          = os.Getenv("AUTH0_CLIENT_SECRET")
	debug                 = os.Getenv("AUTH0_DEBUG")
	httpRecordings        = os.Getenv("AUTH0_HTTP_RECORDINGS")
	httpRecordingsEnabled = false
	m                     = &Management{}
)

func init() {
	httpRecordingsEnabled = httpRecordings == "true" || httpRecordings == "1" || httpRecordings == "on"
	initTestManagement()
}

func initTestManagement() {
	var err error

	m, err = New(
		domain,
		WithClientCredentials(clientID, clientSecret),
		WithDebug(debug == "true" || debug == "1" || debug == "on"),
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

func TestOptionDefauls(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	applyListDefaults([]RequestOption{
		PerPage(20),          // This should be persisted (default is 50).
		IncludeTotals(false), // This should be altered to true by withListDefaults.
	}).apply(r)

	v := r.URL.Query()

	perPage := v.Get("per_page")
	assert.Equal(t, "20", perPage)

	includeTotals := v.Get("include_totals")
	assert.Equal(t, "true", includeTotals)
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

func TestRequestOptionContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the request.

	err := m.Request("GET", "/", nil, Context(ctx))
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected err to be context.Canceled, got %v", err)
	}
}

func TestRequestOptionContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	time.Sleep(50 * time.Millisecond) // Delay until the deadline is exceeded.

	err := m.Request("GET", "/", nil, Context(ctx))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected err to be context.DeadlineExceeded, got %v", err)
	}
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

	u, err := m.User.Read("123")
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
			actual := m.URI(testCase.given...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
