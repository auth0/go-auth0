package management

import (
	"context"
	"time"
)

// EventStream is used for event stream data.
type EventStream struct {
	ID            *string                    `json:"id,omitempty"`
	Status        *string                    `json:"status,omitempty"`
	Name          *string                    `json:"name,omitempty"`
	Subscriptions *[]EventStreamSubscription `json:"subscriptions,omitempty"`
	CreatedAt     *time.Time                 `json:"created_at,omitempty"`
	UpdatedAt     *time.Time                 `json:"updated_at,omitempty"`
	Destination   *EventStreamDestination    `json:"destination,omitempty"`
}

// EventStreamList is a list of Event Streams.
type EventStreamList struct {
	List
	EventStreams []*EventStream `json:"eventStreams,omitempty"`
}

// EventStreamSubscription represents subscription information.
type EventStreamSubscription struct {
	EventStreamSubscriptionType *string `json:"event_type,omitempty"`
}

// EventStreamDestination represents destination details.
type EventStreamDestination struct {
	EventStreamDestinationType          *string                `json:"type,omitempty"`
	EventStreamDestinationConfiguration map[string]interface{} `json:"configuration,omitempty"`
}

// reset cleans up unnecessary fields based on the operation type.
func (e *EventStream) reset(op string) {
	e.ID = nil
	e.CreatedAt = nil
	e.UpdatedAt = nil

	switch op {
	case "create":
		e.Status = nil
	default:
	}
}

// EventStreamManager manages Auth0 Event Stream resources.
type EventStreamManager manager

// Create an Event Stream.
func (m *EventStreamManager) Create(ctx context.Context, e *EventStream, opts ...RequestOption) error {
	e.reset("create")
	return m.management.Request(ctx, "POST", m.management.URI("event-streams"), e, opts...)
}

// Update an Event Stream.
func (m *EventStreamManager) Update(ctx context.Context, id string, e *EventStream, opts ...RequestOption) error {
	e.reset("update")
	return m.management.Request(ctx, "PATCH", m.management.URI("event-streams", id), e, opts...)
}

// List all Event Streams.
func (m *EventStreamManager) List(ctx context.Context, opts ...RequestOption) (esl *EventStreamList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("event-streams"), &esl, opts...)
	return
}

// Read an Event Stream by its id.
func (m *EventStreamManager) Read(ctx context.Context, id string, opts ...RequestOption) (k *EventStream, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("event-streams", id), &k, opts...)
	return
}

// Delete an Event Stream by its id.
func (m *EventStreamManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("event-streams", id), nil, opts...)
}
