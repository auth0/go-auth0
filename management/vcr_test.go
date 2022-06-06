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

func setupVCR(t *testing.T) {
	t.Helper()

	if !vcrTestsEnabled {
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
		i.URL = strings.Replace(i.URL, domain, "go-auth0-dev.eu.auth0.com", -1)
		i.Duration = time.Millisecond

		return nil
	})
}
