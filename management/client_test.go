package management

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestClient_Create(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClient := &Client{
		Name:        auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
	}

	err := m.Client.Create(expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())

	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestClient_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClient := givenAClient(t)

	actualClient, err := m.Client.Read(expectedClient.GetClientID())

	assert.NoError(t, err)
	assert.Equal(t, expectedClient.GetName(), actualClient.GetName())
}

func TestClient_Update(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClient := givenAClient(t)

	expectedDescription := "This is more than just a test client."
	expectedClient.Description = &expectedDescription

	clientID := expectedClient.GetClientID()
	expectedClient.ClientID = nil                       // Read-Only: Additional properties not allowed.
	expectedClient.SigningKeys = nil                    // Read-Only: Additional properties not allowed.
	expectedClient.JWTConfiguration.SecretEncoded = nil // Read-Only: Additional properties not allowed.
	expectedClient.ClientSecret = nil

	err := m.Client.Update(clientID, expectedClient)

	assert.NoError(t, err)
	assert.Equal(t, expectedDescription, *expectedClient.Description)
}

func TestClient_Delete(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClient := givenAClient(t)

	err := m.Client.Delete(expectedClient.GetClientID())
	assert.NoError(t, err)

	actualClient, err := m.Client.Read(expectedClient.GetClientID())

	assert.Empty(t, actualClient)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestClient_List(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClient := givenAClient(t)

	clientList, err := m.Client.List(IncludeFields("client_id"))

	assert.NoError(t, err)
	assert.Contains(t, clientList.Clients, &Client{ClientID: expectedClient.ClientID})
}

func TestClient_RotateSecret(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClient := givenAClient(t)

	oldSecret := expectedClient.GetClientSecret()
	actualClient, err := m.Client.RotateSecret(expectedClient.GetClientID())

	assert.NoError(t, err)
	assert.NotEqual(t, oldSecret, actualClient.GetClientSecret())
}

func TestJWTConfiguration(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		for clientJWTConfiguration, expected := range map[*ClientJWTConfiguration]string{
			{}:                                   `{}`,
			{LifetimeInSeconds: auth0.Int(1000)}: `{"lifetime_in_seconds":1000}`,
		} {
			jsonBody, err := json.Marshal(clientJWTConfiguration)
			assert.NoError(t, err)
			assert.Equal(t, string(jsonBody), expected)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		for jsonBody, expected := range map[string]*ClientJWTConfiguration{
			`{}`:                             {LifetimeInSeconds: nil},
			`{"lifetime_in_seconds":1000}`:   {LifetimeInSeconds: auth0.Int(1000)},
			`{"lifetime_in_seconds":"1000"}`: {LifetimeInSeconds: auth0.Int(1000)},
		} {
			var actual ClientJWTConfiguration
			err := json.Unmarshal([]byte(jsonBody), &actual)

			assert.NoError(t, err)
			assert.Equal(t, &actual, expected)
		}
	})
}

func givenAClient(t *testing.T) *Client {
	t.Helper()

	client := &Client{
		Name:              auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description:       auth0.String("This is just a test client."),
		OrganizationUsage: auth0.String("allow"),
	}

	err := m.Client.Create(client)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupClient(t, client.GetClientID())
	})

	return client
}

func cleanupClient(t *testing.T, clientID string) {
	t.Helper()

	err := m.Client.Delete(clientID)
	require.NoError(t, err)
}
