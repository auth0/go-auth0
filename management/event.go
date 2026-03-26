package management

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// EventManager manages Auth0 Event resources for streaming events.
type EventManager manager

// Event represents a single event received from the SSE stream.
type Event struct {
	ID          *string                `json:"id,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Source      *string                `json:"source,omitempty"`
	SpecVersion *string                `json:"specversion,omitempty"`
	Time        *time.Time             `json:"time,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	A0Tenant    *string                `json:"a0tenant,omitempty"`
}

// SSEMessage represents the wrapper structure for SSE messages from Auth0.
// Each message contains an offset for pagination and the actual event.
type SSEMessage struct {
	Offset *string `json:"offset,omitempty"`
	Event  *Event  `json:"event,omitempty"`
}

// EventStreamReader provides an interface to read events from an SSE stream.
type EventStreamReader struct {
	reader   *bufio.Reader
	response *http.Response
	ctx      context.Context
	cancel   context.CancelFunc
}

// Read reads the next event from the stream.
// It returns io.EOF when the stream is closed or context is cancelled.
func (r *EventStreamReader) Read() (*SSEMessage, error) {
	// Check if context is already cancelled.
	select {
	case <-r.ctx.Done():
		return nil, r.ctx.Err()
	default:
	}

	for {
		line, err := r.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// If we have partial data before EOF, try to parse it.
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "data:") {
					data := strings.TrimPrefix(line, "data:")
					data = strings.TrimSpace(data)
					if data != "" {
						return parseEventData(data)
					}
				}
				return nil, io.EOF
			}
			// Check if error is due to context cancellation.
			select {
			case <-r.ctx.Done():
				return nil, r.ctx.Err()
			default:
				return nil, fmt.Errorf("failed to read from stream: %w", err)
			}
		}

		line = strings.TrimSpace(line)

		// Skip empty lines (SSE uses blank lines as event delimiters).
		if line == "" {
			continue
		}

		// Parse SSE data fields.
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimPrefix(line, "data:")
			data = strings.TrimSpace(data)

			if data == "" {
				continue
			}

			return parseEventData(data)
		}
	}
}

// Close closes the underlying response body and releases resources.
func (r *EventStreamReader) Close() error {
	if r.cancel != nil {
		r.cancel()
	}
	if r.response != nil && r.response.Body != nil {
		return r.response.Body.Close()
	}
	return nil
}

func parseEventData(data string) (*SSEMessage, error) {
	if data == "" {
		return nil, nil
	}

	var msg SSEMessage
	if err := json.Unmarshal([]byte(data), &msg); err != nil {
		return nil, fmt.Errorf("failed to parse event data: %w", err)
	}

	return &msg, nil
}

// Stream opens a Server-Sent Events (SSE) connection to the /api/v2/event endpoint
// and returns an EventStreamReader that can be used to read events as they arrive.
//
// The caller is responsible for closing the returned EventStreamReader when done.
// The context passed to Stream controls the lifetime of the connection - when
// cancelled, the stream will be closed and Read() will return a context error.
//
// Example usage for CLI tailing:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// Handle Ctrl+C
//	sigCh := make(chan os.Signal, 1)
//	signal.Notify(sigCh, os.Interrupt)
//	go func() {
//	    <-sigCh
//	    cancel()
//	}()
//
//	reader, err := api.Event.Stream(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer reader.Close()
//
//	for {
//	    msg, err := reader.Read()
//	    if err != nil {
//	        if err == context.Canceled {
//	            fmt.Println("Stream cancelled")
//	            break
//	        }
//	        if err == io.EOF {
//	            fmt.Println("Stream ended")
//	            break
//	        }
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Event: %s (type: %s)\n", msg.Event.GetID(), msg.Event.GetType())
//	}
func (m *EventManager) Stream(ctx context.Context, opts ...RequestOption) (*EventStreamReader, error) {
	// Create a cancellable context for the stream.
	streamCtx, cancel := context.WithCancel(ctx)

	request, err := m.management.NewRequest(streamCtx, http.MethodGet, m.management.URI("event"), nil, opts...)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Accept header to request SSE stream.
	request.Header.Set("Accept", "text/event-stream")

	response, err := m.management.Do(request)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to connect to event stream: %w", err)
	}

	if response.StatusCode >= http.StatusBadRequest {
		cancel()
		defer response.Body.Close()
		return nil, newError(response)
	}

	return &EventStreamReader{
		reader:   bufio.NewReader(response.Body),
		response: response,
		ctx:      streamCtx,
		cancel:   cancel,
	}, nil
}
