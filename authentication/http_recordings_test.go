package authentication

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"

	"github.com/auth0/go-auth0/authentication/oauth"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

const (
	recordingsDIR    = "./../test/data/recordings/authentication/"
	recordingsDomain = "go-auth0-dev.eu.auth0.com"
)

func configureHTTPTestRecordings(t *testing.T) {
	t.Helper()

	if !httpRecordingsEnabled {
		return
	}

	initialTransport := authAPI.http.Transport

	recorderTransport, err := recorder.NewWithOptions(
		&recorder.Options{
			CassetteName:       recordingsDIR + t.Name(),
			Mode:               recorder.ModeRecordOnce,
			RealTransport:      authAPI.http.Transport,
			SkipRequestLatency: true,
		},
	)
	require.NoError(t, err)

	removeSensitiveDataFromRecordings(t, recorderTransport)

	authAPI.http.Transport = recorderTransport

	// Set a custom matcher that will ensure the request body matches the recording.
	recorderTransport.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		if r.Body == nil || r.Body == http.NoBody {
			return cassette.DefaultMatcher(r, i)
		}

		defer func() {
			err = r.Body.Close()
			require.NoError(t, err)
		}()

		reqBody, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		rb := strings.TrimSpace(string(reqBody))

		bodyMatches := false
		switch r.Header.Get("Content-Type") {
		case "application/json":
			v := map[string]string{}
			err := json.Unmarshal([]byte(rb), &v)
			require.NoError(t, err)

			if v["client_assertion"] != "" {
				verifyClientAssertion(t, v["client_assertion"])
				v["client_assertion"] = "test-client_assertion"
			}

			body, err := json.Marshal(v)
			require.NoError(t, err)

			bodyMatches = assert.JSONEq(t, i.Body, string(body))
			break
		case "application/x-www-form-urlencoded":
			err = r.ParseForm()
			require.NoError(t, err)

			if r.Form.Has("client_assertion") {
				verifyClientAssertion(t, r.Form.Get("client_assertion"))
				r.Form.Set("client_assertion", "test-client_assertion")
			}

			bodyMatches = assert.Equal(t, i.Form, r.Form)
			break
		default:
			bodyMatches = rb == i.Body
		}

		return r.Method == i.Method && r.URL.String() == i.URL && bodyMatches
	})

	t.Cleanup(func() {
		err := recorderTransport.Stop()
		require.NoError(t, err)
		authAPI.http.Transport = initialTransport
	})
}

func removeSensitiveDataFromRecordings(t *testing.T, recorderTransport *recorder.Recorder) {
	recorderTransport.AddHook(
		func(i *cassette.Interaction) error {
			skip429Response(i)
			redactHeaders(i)
			redactClientAuth(t, i)
			redactTokens(t, i)

			// Redact domain should always be ran last
			redactDomain(i, domain)

			return nil
		},
		recorder.BeforeSaveHook,
	)
}

func skip429Response(i *cassette.Interaction) {
	if i.Response.Code == http.StatusTooManyRequests {
		i.DiscardOnSave = true
	}
}

func redactHeaders(i *cassette.Interaction) {
	allowedHeaders := map[string]bool{
		"Content-Type": true,
		"User-Agent":   true,
	}

	for header := range i.Request.Headers {
		if _, ok := allowedHeaders[header]; !ok {
			delete(i.Request.Headers, header)
		}
	}
	for header := range i.Response.Headers {
		if _, ok := allowedHeaders[header]; !ok {
			delete(i.Response.Headers, header)
		}
	}
}

func redactClientAuth(t *testing.T, i *cassette.Interaction) {
	contentType := i.Request.Headers.Get("Content-Type")
	clientAuthParams := map[string]bool{
		"client_id":        true,
		"client_secret":    true,
		"client_assertion": true,
	}

	if contentType == "application/x-www-form-urlencoded" {
		for param := range clientAuthParams {
			if i.Request.Form.Has(param) {
				i.Request.Form.Set(param, "test-"+param)
			}
		}

		i.Request.Body = i.Request.Form.Encode()
	} else if contentType == "application/json" {
		jsonBody := map[string]interface{}{}

		err := json.Unmarshal([]byte(i.Request.Body), &jsonBody)
		require.NoError(t, err)

		for k := range jsonBody {
			if _, ok := clientAuthParams[k]; ok {
				jsonBody[k] = "test-" + k
			}
		}

		body, err := json.Marshal(jsonBody)
		require.NoError(t, err)

		i.Request.Body = string(body)
	}
}

func redactTokens(t *testing.T, i *cassette.Interaction) {
	if i.Response.Headers.Get("Content-Type") != "application/json" {
		return
	}

	tokenSet := &oauth.TokenSet{}

	err := json.Unmarshal([]byte(i.Response.Body), tokenSet)
	require.NoError(t, err)

	tokenSet.AccessToken = "test-access-token"
	tokenSet.IDToken = "test-id-token"

	if tokenSet.RefreshToken != "" {
		tokenSet.RefreshToken = "test-refresh-token"
	}

	body, err := json.Marshal(tokenSet)
	require.NoError(t, err)

	i.Response.Body = string(body)
}

func redactDomain(i *cassette.Interaction, domain string) {
	i.Request.Host = strings.ReplaceAll(i.Request.Host, domain, recordingsDomain)
	i.Request.URL = strings.ReplaceAll(i.Request.URL, domain, recordingsDomain)

	domainParts := strings.Split(domain, ".")

	i.Response.Body = strings.ReplaceAll(i.Response.Body, domainParts[0], recordingsDomain)
	i.Request.Body = strings.ReplaceAll(i.Request.Body, domainParts[0], recordingsDomain)
}

func verifyClientAssertion(t *testing.T, clientAssertion string) {
	key, err := jwk.ParseKey([]byte(jwtPublicKey), jwk.WithPEM(true))
	require.NoError(t, err)

	_, err = jws.Verify([]byte(clientAssertion), jws.WithKey(jwa.RS256, key))

	require.NoError(t, err)
}
