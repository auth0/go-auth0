package management

import (
	"context"
	"strings"
	"time"
)

// EventStreamManager manages Auth0 Event Stream resources.
type EventStreamManager manager

/* ------------------------------------------------ CRUD ---------------------------------------------------------*/

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

// Create an Event Stream.
func (m *EventStreamManager) Create(ctx context.Context, e *EventStream, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("event-streams"), e, opts...)
}

// Update an Event Stream.
func (m *EventStreamManager) Update(ctx context.Context, id string, e *EventStream, opts ...RequestOption) error {
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

/* ------------------------------------------------ UTILS ---------------------------------------------------------*/

// WithEventTypes returns a RequestOption that sets the `event_types` query parameter
// by joining the provided event type strings with commas.
// This is useful for filtering results by multiple event types in endpoints like /deliveries.
func WithEventTypes(types ...string) RequestOption {
	return Parameter("event_types", strings.Join(types, ","))
}

/* ------------------------------------------------ DELIVERY ---------------------------------------------------------*/

// BulkRedeliverRequest specifies filters for bulk redelivery of events.
type BulkRedeliverRequest struct {
	EventTypes *[]string `json:"event_types,omitempty"`
	DateFrom   *string   `json:"date_from,omitempty"`
	DateTo     *string   `json:"date_to,omitempty"`
}

// EventDeliveryList represents the response from the /deliveries endpoint.
type EventDeliveryList struct {
	List
	Deliveries []*EventDelivery `json:"deliveries"`
}

// EventDelivery represents an individual delivery object.
type EventDelivery struct {
	ID            *string            `json:"id,omitempty"`
	EventStreamID *string            `json:"event_stream_id,omitempty"`
	EventType     *string            `json:"event_type,omitempty"`
	Status        *string            `json:"status,omitempty"`
	Attempts      []*DeliveryAttempt `json:"attempts"`
	Event         *DeliveryEvent     `json:"event,omitempty"`
}

// DeliveryAttempt represents an individual delivery attempt.
type DeliveryAttempt struct {
	Status       *string    `json:"status,omitempty"`
	ErrorMessage *string    `json:"error_message,omitempty"`
	Timestamp    *time.Time `json:"timestamp,omitempty"`
	Duration     *float64   `json:"duration,omitempty"`
}

// DeliveryEvent represents the original event data structure.
type DeliveryEvent struct {
	ID          *string                `json:"id,omitempty"`
	Source      *string                `json:"source,omitempty"`
	SpecVersion *string                `json:"specversion,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Time        *time.Time             `json:"time,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	A0Tenant    *string                `json:"a0tenant,omitempty"`
	A0Stream    *string                `json:"a0stream,omitempty"`
	A0Purpose   *string                `json:"a0purpose,omitempty"`
}

// EventDeliveryListQueryParams holds query parameters for delivery listing.
type EventDeliveryListQueryParams struct {
	EventTypes *[]string  `url:"event_types,omitempty"`
	DateFrom   *time.Time `url:"date_from,omitempty"`
	DateTo     *time.Time `url:"date_to,omitempty"`
	From       *string    `url:"from,omitempty"`
	Take       *int       `url:"take,omitempty"`
}

// ReadDelivery fetches a single failed delivery event by eventStreamID and eventID.
func (m *EventStreamManager) ReadDelivery(ctx context.Context, eventStreamID, eventID string, opts ...RequestOption) (ed *EventDelivery, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("event-streams", eventStreamID, "deliveries", eventID), &ed, opts...)
	return
}

// ListDeliveries retrieves failed deliveries for a given Event Stream ID.
//
// Valid query parameters include:
//   - "event_types": Comma-separated list of event types to filter by (e.g. "user.created,organization.created")
//   - "date_from": Start datetime (ISO 8601)
//   - "date_to": End datetime (ISO 8601)
//   - "take": Max results per page (int)
//   - "from": Pagination token returned in `nextPageToken` via ListDeliveries call
func (m *EventStreamManager) ListDeliveries(ctx context.Context, id string, opts ...RequestOption) (edl *EventDeliveryList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("event-streams", id, "deliveries"), &edl, applyListCheckpointDefaults(opts))
	return
}

// Redeliver attempts to resend a previously failed delivery.
func (m *EventStreamManager) Redeliver(ctx context.Context, eventStreamID, eventID string, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("event-streams", eventStreamID, "redeliver", eventID), nil, opts...)
}

// RedeliverMany allow attempting delivery for failed events - Filters are combined using AND logic.
func (m *EventStreamManager) RedeliverMany(ctx context.Context, eventStreamID string, req *BulkRedeliverRequest, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("event-streams", eventStreamID, "redeliver"), req, opts...)
}

/* ------------------------------------------------ TEST ---------------------------------------------------------*/

// TestEvent represents both the request payload and the response from the Test endpoint.
type TestEvent struct {
	// Request Fields
	EventType *string                `json:"event_type"`     // Required for request
	Data      map[string]interface{} `json:"data,omitempty"` // Optional in request

	// Response Fields
	ID            *string            `json:"id,omitempty"`
	EventStreamID *string            `json:"event_stream_id,omitempty"`
	Status        *string            `json:"status,omitempty"`
	Attempts      []*DeliveryAttempt `json:"attempts,omitempty"`
	Event         *DeliveryEvent     `json:"event,omitempty"`
}

// Test triggers a test event on the given event stream.
func (m *EventStreamManager) Test(ctx context.Context, id string, testEvent *TestEvent, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("event-streams", id, "test"), &testEvent, opts...)
}

/* ------------------------------------------------ STATS ---------------------------------------------------------*/

// EventStreamStats represents the statistics returned for an event stream.
type EventStreamStats struct {
	ID      *string        `json:"id,omitempty"`
	Name    *string        `json:"name,omitempty"`
	Window  *StatsWindow   `json:"window,omitempty"`
	Buckets []*time.Time   `json:"buckets,omitempty"`
	Metrics []*StatsMetric `json:"metrics,omitempty"`
}

// StatsWindow defines the time range and interval used for stats calculation.
type StatsWindow struct {
	DateFrom       *time.Time     `json:"date_from,omitempty"`
	DateTo         *time.Time     `json:"date_to,omitempty"`
	BucketInterval *StatsInterval `json:"bucket_interval,omitempty"`
}

// StatsInterval represents the size of each bucket interval in seconds.
type StatsInterval struct {
	ScaleFactor *int `json:"scale_factor,omitempty"`
}

// StatsMetric represents a specific metric tracked over time buckets.
type StatsMetric struct {
	Name        *string `json:"name,omitempty"`
	WindowTotal *int    `json:"window_total,omitempty"`
	Type        *string `json:"type,omitempty"` // e.g., "sum"
	Data        []*int  `json:"data,omitempty"` // aligns index-wise with Buckets
}

// Stats retrieves statistics for the specified event stream, including metrics such as
// deliveries over a given time window. Optional query parameters like `date_from` and
// `date_to` can be passed using RequestOptions.
func (m *EventStreamManager) Stats(ctx context.Context, id string, opts ...RequestOption) (stats *EventStreamStats, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("event-streams", id, "stats"), &stats, opts...)
	return
}
