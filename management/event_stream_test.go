package management

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestEventStreamManager_CreateTypeWebhook(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenEventStream := &EventStream{
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

	err := api.EventStream.Create(context.Background(), givenEventStream)
	assert.NoError(t, err)
	assert.NotEmpty(t, givenEventStream.GetID())
	t.Cleanup(func() {
		cleanUpEventStream(t, givenEventStream.GetID())
	})
}

func TestEventStreamManager_CreateTypeEventBridge(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenEventStream := &EventStream{
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

	err := api.EventStream.Create(context.Background(), givenEventStream)
	assert.NoError(t, err)
	assert.NotEmpty(t, givenEventStream.GetID())
	t.Cleanup(func() {
		cleanUpEventStream(t, givenEventStream.GetID())
	})
}

func TestEventStreamManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	givenEventStream := givenEventStream(t)

	updatedEventStream := &EventStream{
		Name:   auth0.String("updated-stream-name"),
		Status: auth0.String("enabled"),
		Subscriptions: &[]EventStreamSubscription{
			{
				EventStreamSubscriptionType: auth0.String("user.created"),
			},
		},
	}
	err := api.EventStream.Update(context.Background(), givenEventStream.GetID(), updatedEventStream)
	assert.NoError(t, err)

	retrievedEventStream, err := api.EventStream.Read(context.Background(), givenEventStream.GetID())
	assert.NoError(t, err)

	assert.Equal(t, retrievedEventStream.GetName(), updatedEventStream.GetName())
	assert.Equal(t, retrievedEventStream.GetStatus(), updatedEventStream.GetStatus())
	assert.Equal(t, retrievedEventStream.GetSubscriptions(), updatedEventStream.GetSubscriptions())
	t.Cleanup(func() {
		cleanUpEventStream(t, givenEventStream.GetID())
	})
}

func TestEventStreamManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	eventStream := givenEventStream(t)
	eventStreamList, err := api.EventStream.List(context.Background())
	assert.NoError(t, err)
	assert.Contains(t, eventStreamList.EventStreams, eventStream)
}

func TestEventStreamManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	eventStream := givenEventStream(t)
	readEventStream, err := api.EventStream.Read(context.Background(), eventStream.GetID())
	assert.NoError(t, err)
	assert.Equal(t, readEventStream, eventStream)
}

func TestEventStreamManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	eventStream := givenEventStream(t)

	err := api.EventStream.Delete(context.Background(), eventStream.GetID())
	assert.NoError(t, err)

	actualEventStream, err := api.EventStream.Read(context.Background(), eventStream.GetID())

	assert.Empty(t, actualEventStream)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenEventStream(t *testing.T) *EventStream {
	t.Helper()

	givenEventStream := &EventStream{
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

	err := api.EventStream.Create(context.Background(), givenEventStream)
	assert.NoError(t, err)
	assert.NotEmpty(t, givenEventStream.GetID())
	t.Cleanup(func() {
		cleanUpEventStream(t, givenEventStream.GetID())
	})

	return givenEventStream
}

func cleanUpEventStream(t *testing.T, id string) {
	t.Helper()

	err := api.EventStream.Delete(context.Background(), id)
	assert.NoError(t, err)
}
