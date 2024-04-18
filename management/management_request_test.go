package management

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestNewRequest_Get(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)

	request, err := api.NewRequest(context.Background(), "GET", api.URI("clients", expectedClient.GetClientID()), nil)
	assert.NoError(t, err)

	assert.Equal(t, "application/json", request.Header.Get("Content-Type"))
	requestBody, err := io.ReadAll(request.Body)
	assert.NoError(t, err)
	assert.Empty(t, requestBody)

	response, err := api.Do(request)
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	var respClient Client

	responseBody, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseBody)
	if err := json.Unmarshal(responseBody, &respClient); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	t.Cleanup(func() {
		cleanupClient(t, respClient.GetClientID())
	})
}
func TestNewRequest_Post(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := &Client{
		Name:        auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
	}

	request, err := api.NewRequest(context.Background(), "POST", api.URI("clients"), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, request)
	assert.Equal(t, "application/json", request.Header.Get("Content-Type"))

	response, err := api.Do(request)
	assert.NoError(t, err)
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		t.Fatalf("Received error response: %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var respClient Client
	if err := json.Unmarshal(responseBody, &respClient); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	t.Cleanup(func() {
		cleanupClient(t, respClient.GetClientID())
	})
}

func TestNewRequest_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)

	request, err := api.NewRequest(context.Background(), "DELETE", api.URI("clients", expectedClient.GetClientID()), nil)
	assert.NoError(t, err)
	requestBody, err := io.ReadAll(request.Body)
	assert.NoError(t, err)
	assert.Empty(t, requestBody)

	response, err := api.Do(request)
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
