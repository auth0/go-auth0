package management

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/auth0/go-auth0"
)

const (
	recordingsDIR    = "./../test/data/recordings/"
	recordingsDomain = "go-auth0-dev.eu.auth0.com"
)

func configureHTTPTestRecordings(t *testing.T) {
	t.Helper()

	if !httpRecordingsEnabled {
		return
	}

	initialTransport := api.http.Transport

	recorderTransport, err := recorder.NewWithOptions(
		&recorder.Options{
			CassetteName:       recordingsDIR + t.Name(),
			Mode:               recorder.ModeRecordOnce,
			RealTransport:      api.http.Transport,
			SkipRequestLatency: true,
		},
	)
	require.NoError(t, err)

	removeSensitiveDataFromRecordings(t, recorderTransport)

	api.http.Transport = recorderTransport

	t.Cleanup(func() {
		err := recorderTransport.Stop()
		require.NoError(t, err)

		api.http.Transport = initialTransport
	})
}

func removeSensitiveDataFromRecordings(t *testing.T, recorderTransport *recorder.Recorder) {
	recorderTransport.AddHook(
		func(i *cassette.Interaction) error {
			skip429Response(i)
			redactHeaders(i)
			redactSensitiveDataInSigningKey(t, i)
			redactSensitiveDataInClient(t, i)
			redactSensitiveDataInResourceServer(t, i)
			redactSensitiveDataInLogs(t, i)
			redactSensitiveDataInConnectionSCIMToken(t, i)

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

func redactDomain(i *cassette.Interaction, domain string) {
	i.Request.Host = strings.ReplaceAll(i.Request.Host, domain, recordingsDomain)
	i.Request.URL = strings.ReplaceAll(i.Request.URL, domain, recordingsDomain)

	domainParts := strings.Split(domain, ".")

	i.Response.Body = strings.ReplaceAll(i.Response.Body, domainParts[0], recordingsDomain)
	i.Request.Body = strings.ReplaceAll(i.Request.Body, domainParts[0], recordingsDomain)
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

	if i.Request.URL == "https://"+domain+"/api/v2/keys/signing" && i.Request.Method == http.MethodGet {
		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)
		previousSigningKeyBody, err := json.Marshal(previousSigningKey)
		require.NoError(t, err)

		i.Response.Body = fmt.Sprintf(`[%s,%s]`, signingKeyBody, previousSigningKeyBody)

		return
	}

	isSigningKeysURL := strings.Contains(i.Request.URL, "https://"+domain+"/api/v2/keys/signing")
	if isSigningKeysURL && i.Request.Method == http.MethodGet {
		i.Request.URL = "https://" + domain + "/api/v2/keys/signing/111111111111111111111"

		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)

		i.Response.Body = string(signingKeyBody)

		return
	}

	if isSigningKeysURL && strings.Contains(i.Request.URL, "/revoke") && i.Request.Method == http.MethodPut {
		i.Request.URL = "https://" + domain + "/api/v2/keys/signing/111111111111111111111/revoke"

		signingKey.RevokedAt = auth0.Time(time.Now())
		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)

		i.Response.Body = string(signingKeyBody)

		return
	}

	if isSigningKeysURL && strings.Contains(i.Request.URL, "/rotate") && i.Request.Method == http.MethodPost {
		signingKeyBody, err := json.Marshal(signingKey)
		require.NoError(t, err)

		i.Response.Body = string(signingKeyBody)

		return
	}
}

func redactSensitiveDataInConnectionSCIMToken(t *testing.T, i *cassette.Interaction) {
	isTokenURL := strings.Contains(i.Request.URL, "https://"+domain+"/api/v2/connections") && strings.Contains(i.Request.URL, "scim-configuration/tokens")

	create := isTokenURL && i.Request.Method == http.MethodPost
	if create {
		var token SCIMToken
		err := json.Unmarshal([]byte(i.Response.Body), &token)
		require.NoError(t, err)

		redacted := "[REDACTED]"
		token.Token = &redacted

		tokenBody, err := json.Marshal(token)
		require.NoError(t, err)

		i.Response.Body = string(tokenBody)
	}
}

func redactSensitiveDataInClient(t *testing.T, i *cassette.Interaction) {
	isClientURL := strings.Contains(i.Request.URL, "https://"+domain+"/api/v2/clients") &&
		!strings.Contains(i.Request.URL, "/connections")
	create := isClientURL && i.Request.Method == http.MethodPost
	read := isClientURL && i.Request.Method == http.MethodGet
	update := isClientURL && i.Request.Method == http.MethodPatch
	rotateSecret := isClientURL && strings.Contains(i.Request.URL, "/rotate-secret")
	list := isClientURL && strings.Contains(i.Request.URL, "include_totals")
	credentials := isClientURL && strings.Contains(i.Request.URL, "/credentials")

	if (create || read || update || rotateSecret) && !list && !credentials {
		var client Client
		err := json.Unmarshal([]byte(i.Response.Body), &client)
		require.NoError(t, err)

		redacted := "[REDACTED]"
		if rotateSecret {
			redacted = "[ROTATED-REDACTED]"
		}

		client.SigningKeys = []map[string]string{
			{"cert": redacted},
		}
		client.ClientSecret = &redacted

		clientBody, err := json.Marshal(client)
		require.NoError(t, err)

		i.Response.Body = string(clientBody)
	}
}

func redactSensitiveDataInResourceServer(t *testing.T, i *cassette.Interaction) {
	isResourceServerURL := strings.Contains(i.Request.URL, "https://"+domain+"/api/v2/resource-servers")
	create := isResourceServerURL && i.Request.Method == http.MethodPost
	read := isResourceServerURL && i.Request.Method == http.MethodGet
	update := isResourceServerURL && i.Request.Method == http.MethodPatch
	list := isResourceServerURL && strings.Contains(i.Request.URL, "include_totals")

	if (create || read || update) && !list {
		var rs ResourceServer
		err := json.Unmarshal([]byte(i.Response.Body), &rs)
		require.NoError(t, err)

		rs.SigningSecret = nil

		rsBody, err := json.Marshal(rs)
		require.NoError(t, err)

		i.Response.Body = string(rsBody)
	}
}

func redactSensitiveDataInLogs(t *testing.T, i *cassette.Interaction) {
	isMultipleLogsURL := strings.Contains(i.Request.URL, "https://"+domain+"/api/v2/logs?")
	if isMultipleLogsURL {
		var logs []Log
		err := json.Unmarshal([]byte(i.Response.Body), &logs)
		require.NoError(t, err)

		for i, log := range logs {
			log.IP = auth0.String("[REDACTED]")
			logs[i] = log
		}

		logsBody, err := json.Marshal(logs)
		require.NoError(t, err)

		i.Response.Body = string(logsBody)
	}

	isSingleLogURL := strings.Contains(i.Request.URL, "https://"+domain+"/api/v2/logs/")
	if isSingleLogURL {
		var log Log
		err := json.Unmarshal([]byte(i.Response.Body), &log)
		require.NoError(t, err)

		log.IP = auth0.String("[REDACTED]")

		logsBody, err := json.Marshal(log)
		require.NoError(t, err)

		i.Response.Body = string(logsBody)
	}
}
