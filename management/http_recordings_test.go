package management

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/v2/cassette"
	"github.com/dnaeon/go-vcr/v2/recorder"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func setupHTTPRecordings(t *testing.T) {
	t.Helper()

	if !httpRecordingsEnabled {
		return
	}

	const recordingsDIR = "testdata/recordings/"

	initialTransport := m.http.Transport

	vcrTransport, err := recorder.NewAsMode(
		recordingsDIR+t.Name(),
		recorder.ModeReplaying,
		m.http.Transport,
	)
	require.NoError(t, err)

	removeSensitiveDataFromRecordings(t, vcrTransport)

	m.http.Transport = vcrTransport

	t.Cleanup(func() {
		err := vcrTransport.Stop()
		require.NoError(t, err)
		m.http.Transport = initialTransport
	})
}

func removeSensitiveDataFromRecordings(t *testing.T, vcr *recorder.Recorder) {
	requestHeaders := []string{"Authorization"}
	responseHeaders := []string{
		"Alt-Svc",
		"Cache-Control",
		"Cf-Cache-Status",
		"Cf-Ray",
		"Date",
		"Expect-Ct",
		"Ot-Baggage-Auth0-Request-Id",
		"Ot-Tracer-Sampled",
		"Ot-Tracer-Spanid",
		"Ot-Tracer-Traceid",
		"Server",
		"Set-Cookie",
		"Strict-Transport-Security",
		"Traceparent",
		"Tracestate",
		"Vary",
		"X-Content-Type-Options",
		"X-Ratelimit-Limit",
		"X-Ratelimit-Remaining",
		"X-Ratelimit-Reset",
	}

	vcr.AddFilter(func(i *cassette.Interaction) error {
		for _, header := range requestHeaders {
			delete(i.Request.Headers, header)
		}
		return nil
	})
	vcr.AddFilter(func(i *cassette.Interaction) error {
		for _, header := range responseHeaders {
			delete(i.Response.Headers, header)
		}
		return nil
	})
	vcr.AddSaveFilter(func(i *cassette.Interaction) error {
		redactSensitiveDataInSigningKey(t, i)
		redactSensitiveDataInClient(t, i)
		redactSensitiveDataInResourceServer(t, i)

		i.URL = strings.Replace(i.URL, domain, "go-auth0-dev.eu.auth0.com", -1)
		i.Duration = time.Millisecond

		return nil
	})
}

func redactSensitiveDataInSigningKey(t *testing.T, i *cassette.Interaction) {
	signingKey := &SigningKey{
		KID:         auth0.String("111111111111111111111"),
		Cert:        auth0.String("-----BEGIN CERTIFICATE-----\\r\\n[REDACTED]\\r\\n-----END CERTIFICATE-----"),
		PKCS7:       auth0.String("-----BEGIN PKCS7-----\\r\\n[REDACTED]\\r\\n-----END PKCS7-----"),
		Current:     auth0.Bool(true),
		Next:        auth0.Bool(false),
		Previous:    auth0.Bool(true),
		Fingerprint: auth0.String("[REDACTED]"),
		Thumbprint:  auth0.String("[REDACTED]"),
		Revoked:     auth0.Bool(false),
	}
	previousSigningKey := &SigningKey{
		KID:         auth0.String("222222222222222222222"),
		Cert:        auth0.String("-----BEGIN CERTIFICATE-----\\r\\n[REDACTED]\\r\\n-----END CERTIFICATE-----"),
		PKCS7:       auth0.String("-----BEGIN PKCS7-----\\r\\n[REDACTED]\\r\\n-----END PKCS7-----"),
		Current:     auth0.Bool(false),
		Next:        auth0.Bool(true),
		Previous:    auth0.Bool(true),
		Fingerprint: auth0.String("[REDACTED]"),
		Thumbprint:  auth0.String("[REDACTED]"),
		Revoked:     auth0.Bool(false),
	}

	if i.URL == "https://"+domain+"/api/v2/keys/signing" && i.Method == http.MethodGet {
		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)
		previousSigningKeyBody, err := json.Marshal(previousSigningKey)
		require.NoError(t, err)

		i.Response.Body = fmt.Sprintf(`[%s,%s]`, signingKeyBody, previousSigningKeyBody)
	}

	if strings.Contains(i.URL, "https://"+domain+"/api/v2/keys/signing/") && i.Method == http.MethodGet {
		i.URL = "https://" + domain + "/api/v2/keys/signing/111111111111111111111"

		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)

		i.Response.Body = string(signingKeyBody)
	}

	if strings.Contains(i.URL, "/revoke") && strings.Contains(i.URL, "keys/signing") && i.Method == http.MethodPut {
		i.URL = "https://" + domain + "/api/v2/keys/signing/111111111111111111111/revoke"

		signingKey.RevokedAt = auth0.Time(time.Now())
		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)

		i.Response.Body = string(signingKeyBody)
	}

	if strings.Contains(i.URL, "/rotate") && strings.Contains(i.URL, "keys/signing") && i.Method == http.MethodPost {
		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)

		i.Response.Body = string(signingKeyBody)
	}
}

func redactSensitiveDataInClient(t *testing.T, i *cassette.Interaction) {
	create := i.URL == "https://"+domain+"/api/v2/clients" && i.Method == http.MethodPost
	read := strings.Contains(i.URL, "https://"+domain+"/api/v2/clients/") && i.Method == http.MethodGet
	update := strings.Contains(i.URL, "https://"+domain+"/api/v2/clients/") && i.Method == http.MethodPatch
	rotateSecret := strings.Contains(i.URL, "clients") && strings.Contains(i.URL, "/rotate-secret")

	if create || read || update || rotateSecret {
		var client Client
		err := json.Unmarshal([]byte(i.Response.Body), &client)
		require.NoError(t, err)

		client.SigningKeys = nil
		client.ClientSecret = nil

		clientBody, err := json.Marshal(client)
		require.NoError(t, err)

		i.Response.Body = string(clientBody)
	}
}

func redactSensitiveDataInResourceServer(t *testing.T, i *cassette.Interaction) {
	create := i.URL == "https://"+domain+"/api/v2/resource-servers" && i.Method == http.MethodPost
	read := strings.Contains(i.URL, "https://"+domain+"/api/v2/resource-servers/") && i.Method == http.MethodGet
	update := strings.Contains(i.URL, "https://"+domain+"/api/v2/resource-servers/") && i.Method == http.MethodPatch

	if create || read || update {
		var rs ResourceServer
		err := json.Unmarshal([]byte(i.Response.Body), &rs)
		require.NoError(t, err)

		rs.SigningSecret = nil

		rsBody, err := json.Marshal(rs)
		require.NoError(t, err)

		i.Response.Body = string(rsBody)
	}
}
