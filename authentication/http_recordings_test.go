package authentication

import (
	"net/http"
	"strings"
	"testing"

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

	removeSensitiveDataFromRecordings(recorderTransport)

	authAPI.http.Transport = recorderTransport

	t.Cleanup(func() {
		err := recorderTransport.Stop()
		require.NoError(t, err)
		authAPI.http.Transport = initialTransport
	})
}

func removeSensitiveDataFromRecordings(recorderTransport *recorder.Recorder) {
	recorderTransport.AddHook(
		func(i *cassette.Interaction) error {
			skip429Response(i)
			redactHeaders(i)

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
