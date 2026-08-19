package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internal "github.com/auth0/go-auth0/v3/internal/client"
	"github.com/auth0/go-auth0/v3/management/option"
)

func TestNewBaseURL(t *testing.T) {
	tests := []struct {
		domain      string
		expectedURL string
	}{
		{"example.com", "https://example.com/api/v2"},
		{"example.com/", "https://example.com/api/v2"},
		{"https://example.com", "https://example.com/api/v2"},
		{"https://example.com/", "https://example.com/api/v2"},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			m, err := New(tt.domain)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedURL, m.baseURL)
		})
	}
}

// doerHTTPClient is a core.HTTPClient that is not an *http.Client.
type doerHTTPClient struct {
	called bool
}

func (d *doerHTTPClient) Do(req *http.Request) (*http.Response, error) {
	d.called = true

	return http.DefaultClient.Do(req)
}

// markerTransport identifies the transport a client was created with.
type markerTransport struct{}

func (m *markerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(req)
}

// headerRecorder serves requests and keeps the headers of the last one.
type headerRecorder struct {
	server *httptest.Server
	header http.Header
}

func newHeaderRecorder(t *testing.T) *headerRecorder {
	t.Helper()

	recorder := &headerRecorder{}
	recorder.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder.header = r.Header.Clone()

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"clients":[]}`))
	}))

	t.Cleanup(recorder.server.Close)

	return recorder
}

// domain returns the host of the test server, which New expects without a scheme.
func (h *headerRecorder) domain() string {
	return strings.TrimPrefix(h.server.URL, "http://")
}

// auth0Client decodes the "Auth0-Client" header of the recorded request.
func (h *headerRecorder) auth0Client(t *testing.T) *internal.Auth0ClientInfo {
	t.Helper()

	encoded := h.header.Get("Auth0-Client")
	require.NotEmpty(t, encoded, `expected an "Auth0-Client" header`)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	require.NoError(t, err)

	var clientInfo internal.Auth0ClientInfo
	require.NoError(t, json.Unmarshal(decoded, &clientInfo))

	return &clientInfo
}

// listClients issues a request so that the headers set by the transports can be asserted.
func listClients(t *testing.T, recorder *headerRecorder, options ...option.RequestOption) {
	t.Helper()

	m, err := New(
		recorder.domain(),
		append([]option.RequestOption{option.WithInsecure(), option.WithToken("token")}, options...)...,
	)
	require.NoError(t, err)

	_, err = m.Clients.List(t.Context(), nil)
	require.NoError(t, err)
}

func TestNewSendsClientInfo(t *testing.T) {
	recorder := newHeaderRecorder(t)

	listClients(t, recorder)

	assert.Equal(t, internal.UserAgent, recorder.header.Get("User-Agent"))
	assert.Equal(t, internal.DefaultAuth0ClientInfo.Name, recorder.auth0Client(t).Name)
}

func TestNewSendsClientInfoWithCustomHTTPClient(t *testing.T) {
	recorder := newHeaderRecorder(t)

	listClients(t, recorder, option.WithHTTPClient(&http.Client{}))

	assert.Equal(t, internal.UserAgent, recorder.header.Get("User-Agent"))
	assert.Equal(t, internal.DefaultAuth0ClientInfo.Name, recorder.auth0Client(t).Name)
}

func TestNewSendsClientInfoWithCustomHTTPClientImplementation(t *testing.T) {
	recorder := newHeaderRecorder(t)

	httpClient := &doerHTTPClient{}
	listClients(t, recorder, option.WithHTTPClient(httpClient))

	assert.True(t, httpClient.called, "expected the custom client to issue the request")
	assert.Equal(t, internal.UserAgent, recorder.header.Get("User-Agent"))
	assert.Equal(t, internal.DefaultAuth0ClientInfo.Name, recorder.auth0Client(t).Name)
}

func TestNewSendsClientInfoEnvEntriesWithCustomHTTPClient(t *testing.T) {
	recorder := newHeaderRecorder(t)

	listClients(
		t,
		recorder,
		option.WithHTTPClient(&http.Client{}),
		option.WithAuth0ClientEnvEntry("terraform-provider-auth0", "1.0.0"),
	)

	assert.Equal(t, "1.0.0", recorder.auth0Client(t).Env["terraform-provider-auth0"])
}

func TestNewOmitsClientInfoWithCustomHTTPClient(t *testing.T) {
	recorder := newHeaderRecorder(t)

	listClients(t, recorder, option.WithHTTPClient(&http.Client{}), option.WithNoAuth0ClientInfo())

	assert.Empty(t, recorder.header.Get("Auth0-Client"))
	assert.Equal(t, internal.UserAgent, recorder.header.Get("User-Agent"))
}

func TestNewKeepsCustomHTTPClientTransport(t *testing.T) {
	recorder := newHeaderRecorder(t)

	httpClient := &http.Client{
		Transport: internal.RoundTripFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set("X-Custom-Transport", "called")

			return http.DefaultTransport.RoundTrip(req)
		}),
	}

	listClients(t, recorder, option.WithHTTPClient(httpClient))

	assert.Equal(t, "called", recorder.header.Get("X-Custom-Transport"))
	assert.Equal(t, internal.DefaultAuth0ClientInfo.Name, recorder.auth0Client(t).Name)
}

func TestNewDoesNotModifyCustomHTTPClient(t *testing.T) {
	httpClient := &http.Client{
		Transport: &markerTransport{},
		Timeout:   42 * time.Second,
	}

	m, err := New("example.com", option.WithHTTPClient(httpClient))
	require.NoError(t, err)

	// The transports are layered onto a clone, so the caller's client is left alone.
	assert.IsType(t, &markerTransport{}, httpClient.Transport)

	// The clone keeps the settings of the client it was made from.
	usedClient, ok := m.options.HTTPClient.(*http.Client)
	require.True(t, ok, "expected an *http.Client to be used")
	assert.Equal(t, 42*time.Second, usedClient.Timeout)
	assert.NotSame(t, httpClient, usedClient)
	assert.NotEqual(t, "*client.markerTransport", fmt.Sprintf("%T", usedClient.Transport))
}
