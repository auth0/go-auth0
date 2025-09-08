package management

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestEventStreamManager_CreateTypeWebhook(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedEventStream := &EventStream{
		Name: auth0.String("webhook-example"),
		Subscriptions: &[]EventStreamSubscription{
			{
				EventStreamSubscriptionType: auth0.String("user.created"),
			},
			{
				EventStreamSubscriptionType: auth0.String("user.updated"),
			},
			{
				EventStreamSubscriptionType: auth0.String("user.deleted"),
			},
		},
		Destination: &EventStreamDestination{
			EventStreamDestinationType: auth0.String("webhook"),
			EventStreamDestinationConfiguration: map[string]interface{}{
				"webhook_endpoint": "https://eof28wtn4v4506o.m.pipedream.net",
				"webhook_authorization": &map[string]string{
					"method": "bearer",
					"token":  "123456789",
				},
			},
		},
	}

	err := api.EventStream.Create(context.Background(), expectedEventStream)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedEventStream.GetID())
	t.Cleanup(func() {
		cleanUpEventStream(t, expectedEventStream.GetID())
	})
}

func TestEventStreamManager_CreateTypeEventBridge(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedEventStream := &EventStream{
		Name: auth0.String("eventbridge-example"),
		Subscriptions: &[]EventStreamSubscription{
			{
				EventStreamSubscriptionType: auth0.String("user.created"),
			},
			{
				EventStreamSubscriptionType: auth0.String("user.deleted"),
			},
		},
		Destination: &EventStreamDestination{
			EventStreamDestinationType: auth0.String("eventbridge"),
			EventStreamDestinationConfiguration: map[string]interface{}{
				"aws_account_id": "242109005777",
				"aws_region":     "us-east-2",
			},
		},
	}

	err := api.EventStream.Create(context.Background(), expectedEventStream)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedEventStream.GetID())
	t.Cleanup(func() {
		cleanUpEventStream(t, expectedEventStream.GetID())
	})
}

func TestEventStreamManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedEventStream := givenAnEventStream(t)

	updatedEventStream := &EventStream{
		Name:   auth0.String("updated-stream-name"),
		Status: auth0.String("enabled"),
		Subscriptions: &[]EventStreamSubscription{
			{
				EventStreamSubscriptionType: auth0.String("user.created"),
			},
		},
	}
	err := api.EventStream.Update(context.Background(), expectedEventStream.GetID(), updatedEventStream)
	assert.NoError(t, err)

	retrievedEventStream, err := api.EventStream.Read(context.Background(), expectedEventStream.GetID())
	assert.NoError(t, err)

	assert.Equal(t, retrievedEventStream.GetName(), updatedEventStream.GetName())
	assert.Equal(t, retrievedEventStream.GetStatus(), updatedEventStream.GetStatus())
	assert.Equal(t, retrievedEventStream.GetSubscriptions(), updatedEventStream.GetSubscriptions())
	t.Cleanup(func() {
		cleanUpEventStream(t, expectedEventStream.GetID())
	})
}

func TestEventStreamManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedEventStream := givenAnEventStream(t)
	eventStreamList, err := api.EventStream.List(context.Background())
	assert.NoError(t, err)
	assert.Contains(t, eventStreamList.EventStreams, expectedEventStream)
}

func TestEventStreamManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedEventStream := givenAnEventStream(t)
	readEventStream, err := api.EventStream.Read(context.Background(), expectedEventStream.GetID())
	assert.NoError(t, err)
	assert.Equal(t, readEventStream, expectedEventStream)
}

func TestEventStreamManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedEventStream := givenAnEventStream(t)

	err := api.EventStream.Delete(context.Background(), expectedEventStream.GetID())
	assert.NoError(t, err)

	actualEventStream, err := api.EventStream.Read(context.Background(), expectedEventStream.GetID())

	assert.Empty(t, actualEventStream)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenAnEventStream(t *testing.T) *EventStream {
	t.Helper()

	expectedEventStream := &EventStream{
		Name: auth0.String("eventbridge-example"),
		Subscriptions: &[]EventStreamSubscription{
			{
				EventStreamSubscriptionType: auth0.String("user.created"),
			},
			{
				EventStreamSubscriptionType: auth0.String("user.deleted"),
			},
		},
		Destination: &EventStreamDestination{
			EventStreamDestinationType: auth0.String("eventbridge"),
			EventStreamDestinationConfiguration: map[string]interface{}{
				"aws_account_id": "242109005777",
				"aws_region":     "us-east-2",
			},
		},
	}

	err := api.EventStream.Create(context.Background(), expectedEventStream)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedEventStream.GetID())
	t.Cleanup(func() {
		cleanUpEventStream(t, expectedEventStream.GetID())
	})

	return expectedEventStream
}

func givenAFailingWebhookEventStream(t *testing.T) *EventStream {
	t.Helper()

	stream := &EventStream{
		Name: auth0.String("failing-webhook-test-stream"),
		Subscriptions: &[]EventStreamSubscription{
			{
				EventStreamSubscriptionType: auth0.String("user.created"),
			},
			{
				EventStreamSubscriptionType: auth0.String("user.updated"),
			},
		},
		Status: auth0.String("enabled"),
		Destination: &EventStreamDestination{
			EventStreamDestinationType: auth0.String("webhook"),
			EventStreamDestinationConfiguration: map[string]interface{}{
				"webhook_endpoint": "https://api.hooklistener.com/w/random-turquoise-mead-xb3v",
				"webhook_authorization": map[string]interface{}{
					"method": "bearer",
					"token":  "123456789",
				},
			},
		},
	}

	err := api.EventStream.Create(context.Background(), stream)
	assert.NoError(t, err)
	assert.NotEmpty(t, stream.GetID())

	t.Cleanup(func() {
		cleanUpEventStream(t, stream.GetID())
	})

	return stream
}

/*
| Endpoint                 | Operation           | Preconditions                   | Validations                                  |
| ------------------------ | ------------------- | ------------------------------- | -------------------------------------------- |
| `/test`                  | Trigger fake events | Webhook stream w/ subscriptions | Validate event created, status=pending       |
| `/deliveries`            | List failures       | Trigger test events             | Check count, type filter, status             |
| `/deliveries/{event_id}` | View event details  | Event ID from above             | Check attempt history, error, timestamps     |
| `/redeliver`             | Bulk redelivery     | Failed events exist             | Triggered again, stats window total += N     |
| `/redeliver/{event_id}`  | Single redelivery   | Specific failed event ID        | +1 attempt shown in `/deliveries/{event_id}` |
| `/stats`                 | Metrics             | Events delivered                | Buckets, window range, failed delivery count |
*/

func TestEventStreamManager_Integration(t *testing.T) {
	configureHTTPTestRecordings(t)
	stream := givenAFailingWebhookEventStream(t)

	t.Run("Trigger Test Events", func(t *testing.T) {
		eventTypes := []string{"user.created", "user.created", "user.created", "user.updated"}

		for _, et := range eventTypes {
			testEvent := &TestEvent{EventType: auth0.String(et)}
			err := api.EventStream.Test(context.Background(), stream.GetID(), testEvent)

			time.Sleep(1 * time.Second)
			require.NoError(t, err)
			require.Equal(t, "pending", testEvent.GetStatus())
		}
	})

	t.Run("Validate Failed Deliveries", func(t *testing.T) {
		time.Sleep(60 * time.Second) // let async delivery fail

		deliveryList, err := api.EventStream.ListDeliveries(
			context.Background(),
			stream.GetID(),
			WithEventTypes("user.created", "user.updated"))
		require.NoError(t, err)
		require.Len(t, deliveryList.Deliveries, 4)
	})

	t.Run("Check Stats Before Redelivery", func(t *testing.T) {
		time.Sleep(20 * time.Second)

		stats, err := api.EventStream.Stats(context.Background(), stream.GetID())
		require.NoError(t, err)
		require.Equal(t, "auth0.event_streams.failed_deliveries", stats.Metrics[1].GetName())
		require.Equal(t, 4, stats.Metrics[1].GetWindowTotal())
	})

	t.Run("Bulk Redelivery", func(t *testing.T) {
		time.Sleep(20 * time.Second)

		req := &BulkRedeliverRequest{}
		err := api.EventStream.RedeliverMany(context.Background(), stream.GetID(), req)
		require.NoError(t, err)
	})

	t.Run("Check Stats After Redelivery", func(t *testing.T) {
		time.Sleep(20 * time.Second)

		stats, err := api.EventStream.Stats(context.Background(), stream.GetID())
		require.NoError(t, err)
		require.Equal(t, 6, stats.Metrics[1].GetWindowTotal())
	})

	var updateUserEventID string

	t.Run("Find user.updated Event ID", func(t *testing.T) {
		time.Sleep(20 * time.Second)

		list, err := api.EventStream.ListDeliveries(
			context.Background(),
			stream.GetID(),
			WithEventTypes("user.updated"))
		require.NoError(t, err)
		require.Len(t, list.Deliveries, 1)
		updateUserEventID = list.Deliveries[0].GetID()
	})

	t.Run("Single Redeliver", func(t *testing.T) {
		time.Sleep(20 * time.Second)

		err := api.EventStream.Redeliver(context.Background(), stream.GetID(), updateUserEventID)
		require.NoError(t, err)
	})

	t.Run("Validate Delivery Attempts Increased", func(t *testing.T) {
		time.Sleep(20 * time.Second)

		ev, err := api.EventStream.ReadDelivery(context.Background(), stream.GetID(), updateUserEventID)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(ev.Attempts), 2)
	})
}

func cleanUpEventStream(t *testing.T, id string) {
	t.Helper()

	err := api.EventStream.Delete(context.Background(), id)
	assert.NoError(t, err)
}
