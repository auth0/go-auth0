package management

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventStreamReader_Read(t *testing.T) {
	tests := []struct {
		name           string
		sseData        string
		expectedEvent  *Event
		expectedOffset string
		expectError    bool
	}{
		{
			name:    "Valid event with offset",
			sseData: "data: {\"offset\":\"abc123\",\"event\":{\"id\":\"evt_123\",\"type\":\"user.created\",\"source\":\"auth0\"}}\n",
			expectedEvent: &Event{
				ID:     stringPtr("evt_123"),
				Type:   stringPtr("user.created"),
				Source: stringPtr("auth0"),
			},
			expectedOffset: "abc123",
			expectError:    false,
		},
		{
			name:          "Empty stream",
			sseData:       "",
			expectedEvent: nil,
			expectError:   true, // EOF
		},
		{
			name:    "Event with nested data and a0tenant",
			sseData: "data: {\"offset\":\"xyz789\",\"event\":{\"id\":\"evt_456\",\"type\":\"user.updated\",\"source\":\"urn:auth0:test-tenant.us.auth0.com\",\"specversion\":\"1.0\",\"data\":{\"object\":{\"email\":\"test@example.com\"}},\"a0tenant\":\"test-tenant\"}}\n",
			expectedEvent: &Event{
				ID:       stringPtr("evt_456"),
				Type:     stringPtr("user.updated"),
				Source:   stringPtr("urn:auth0:test-tenant.us.auth0.com"),
				A0Tenant: stringPtr("test-tenant"),
			},
			expectedOffset: "xyz789",
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "text/event-stream", r.Header.Get("Accept"))

				w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")
				w.WriteHeader(http.StatusOK)

				flusher, ok := w.(http.Flusher)
				require.True(t, ok)

				_, _ = w.Write([]byte(tc.sseData))
				flusher.Flush()
			}))
			defer server.Close()

			api, err := New(
				server.URL,
				WithInsecure(),
			)
			require.NoError(t, err)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			reader, err := api.Event.Stream(ctx)
			require.NoError(t, err)
			defer reader.Close()

			msg, err := reader.Read()

			if tc.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, msg)
				require.NotNil(t, msg.Event)
				assert.Equal(t, tc.expectedEvent.GetID(), msg.Event.GetID())
				assert.Equal(t, tc.expectedEvent.GetType(), msg.Event.GetType())
				assert.Equal(t, tc.expectedOffset, msg.GetOffset())
				if tc.expectedEvent.A0Tenant != nil {
					assert.Equal(t, tc.expectedEvent.GetA0Tenant(), msg.Event.GetA0Tenant())
				}
			}
		})
	}
}

func TestEventStreamReader_MultipleEvents(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		events := []string{
			`data: {"offset":"off1","event":{"id":"evt_1","type":"user.created"}}`,
			``,
			`data: {"offset":"off2","event":{"id":"evt_2","type":"user.updated"}}`,
			``,
			`data: {"offset":"off3","event":{"id":"evt_3","type":"user.deleted"}}`,
			``,
		}

		flusher, ok := w.(http.Flusher)
		require.True(t, ok)

		for _, event := range events {
			_, _ = w.Write([]byte(event + "\n"))
		}
		flusher.Flush()
	}))
	defer server.Close()

	api, err := New(
		server.URL,
		WithInsecure(),
	)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reader, err := api.Event.Stream(ctx)
	require.NoError(t, err)
	defer reader.Close()

	expectedTypes := []string{"user.created", "user.updated", "user.deleted"}
	for i, expectedType := range expectedTypes {
		msg, err := reader.Read()
		require.NoError(t, err)
		require.NotNil(t, msg)
		require.NotNil(t, msg.Event)
		assert.Equal(t, expectedType, msg.Event.GetType(), "Event %d type mismatch", i+1)
	}
}

func TestEventStreamReader_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"access_denied","error_description":"Unauthorized","statusCode":401,"message":"Unauthorized"}`))
	}))
	defer server.Close()

	api, err := New(
		server.URL,
		WithInsecure(),
	)
	require.NoError(t, err)

	ctx := context.Background()

	_, err = api.Event.Stream(ctx)
	require.Error(t, err)

	// Verify it's a proper management error.
	var managementErr Error
	assert.ErrorAs(t, err, &managementErr)
	assert.Equal(t, http.StatusUnauthorized, managementErr.Status())
}

func TestEventStreamReader_Close(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		flusher, ok := w.(http.Flusher)
		require.True(t, ok)

		// Send one event then keep connection open.
		_, _ = w.Write([]byte("data: {\"offset\":\"off1\",\"event\":{\"id\":\"evt_1\"}}\n\n"))
		flusher.Flush()

		// Keep connection open until client closes.
		<-r.Context().Done()
	}))
	defer server.Close()

	api, err := New(
		server.URL,
		WithInsecure(),
	)
	require.NoError(t, err)

	ctx := context.Background()

	reader, err := api.Event.Stream(ctx)
	require.NoError(t, err)

	// Read the first event.
	msg, err := reader.Read()
	require.NoError(t, err)
	require.NotNil(t, msg)

	// Close should not error.
	err = reader.Close()
	assert.NoError(t, err)
}

func TestEventStreamReader_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		flusher, ok := w.(http.Flusher)
		require.True(t, ok)

		_, _ = w.Write([]byte("data: {\"offset\":\"off1\",\"event\":{\"id\":\"evt_1\"}}\n\n"))
		flusher.Flush()

		// Wait for context cancellation.
		<-r.Context().Done()
	}))
	defer server.Close()

	api, err := New(
		server.URL,
		WithInsecure(),
	)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())

	reader, err := api.Event.Stream(ctx)
	require.NoError(t, err)
	defer reader.Close()

	// Read first event successfully.
	msg, err := reader.Read()
	require.NoError(t, err)
	require.NotNil(t, msg)

	// Cancel context.
	cancel()

	// Subsequent reads should eventually error.
	_, err = reader.Read()
	assert.Error(t, err)
}

func TestEventManager_StreamEndpoint(t *testing.T) {
	var capturedPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("data: {\"offset\":\"off1\",\"event\":{}}\n\n"))
	}))
	defer server.Close()

	api, err := New(
		server.URL,
		WithInsecure(),
	)
	require.NoError(t, err)

	ctx := context.Background()
	reader, err := api.Event.Stream(ctx)
	require.NoError(t, err)
	defer reader.Close()

	// Verify the correct endpoint was called.
	assert.Equal(t, "/api/v2/event", capturedPath)
}

func TestEventStreamReader_ParseEventData(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		expectNil   bool
		expectError bool
	}{
		{
			name:      "Empty data",
			data:      "",
			expectNil: true,
		},
		{
			name:        "Invalid JSON",
			data:        "not json",
			expectError: true,
		},
		{
			name: "Valid data",
			data: `{"offset":"abc","event":{"id":"evt_1"}}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg, err := parseEventData(tc.data)
			if tc.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if tc.expectNil {
				assert.Nil(t, msg)
			} else {
				assert.NotNil(t, msg)
			}
		})
	}
}

func TestEvent_Getters(t *testing.T) {
	now := time.Now()
	event := &Event{
		ID:          stringPtr("evt_123"),
		Type:        stringPtr("user.created"),
		Source:      stringPtr("auth0"),
		SpecVersion: stringPtr("1.0"),
		Time:        &now,
		Data:        map[string]interface{}{"key": "value"},
		A0Tenant:    stringPtr("my-tenant"),
	}

	assert.Equal(t, "evt_123", event.GetID())
	assert.Equal(t, "user.created", event.GetType())
	assert.Equal(t, "auth0", event.GetSource())
	assert.Equal(t, "1.0", event.GetSpecVersion())
	assert.Equal(t, now, event.GetTime())
	assert.Equal(t, map[string]interface{}{"key": "value"}, event.GetData())
	assert.Equal(t, "my-tenant", event.GetA0Tenant())

	// Test nil event.
	var nilEvent *Event
	assert.Equal(t, "", nilEvent.GetID())
	assert.Equal(t, "", nilEvent.GetType())
	assert.Equal(t, "", nilEvent.GetSource())
	assert.Equal(t, "", nilEvent.GetSpecVersion())
	assert.Equal(t, time.Time{}, nilEvent.GetTime())
	assert.Nil(t, nilEvent.GetData())
	assert.Equal(t, "", nilEvent.GetA0Tenant())
}

func TestSSEMessage_Getters(t *testing.T) {
	msg := &SSEMessage{
		Offset: stringPtr("offset123"),
		Event:  &Event{ID: stringPtr("evt_1")},
	}

	assert.Equal(t, "offset123", msg.GetOffset())
	assert.NotNil(t, msg.GetEvent())
	assert.Equal(t, "evt_1", msg.GetEvent().GetID())

	// Test nil message.
	var nilMsg *SSEMessage
	assert.Equal(t, "", nilMsg.GetOffset())
	assert.Nil(t, nilMsg.GetEvent())
}

func TestEventStreamReader_ReadWithSplitData(t *testing.T) {
	// Test that the reader handles data lines that may come without trailing
	// newlines (EOF mid-stream).
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		// Write data without trailing newline to simulate abrupt stream end.
		_, _ = w.Write([]byte(`data: {"offset":"off1","event":{"id":"evt_1","type":"user.created"}}`))
	}))
	defer server.Close()

	api, err := New(
		server.URL,
		WithInsecure(),
	)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reader, err := api.Event.Stream(ctx)
	require.NoError(t, err)
	defer reader.Close()

	msg, err := reader.Read()
	// Should still parse the event even without trailing newline (EOF case).
	if err == nil {
		require.NotNil(t, msg)
		assert.Equal(t, "evt_1", msg.Event.GetID())
	} else {
		// io.EOF is also acceptable if the reader couldn't get a full line.
		assert.ErrorIs(t, err, io.EOF)
	}
}

func stringPtr(s string) *string {
	return &s
}
