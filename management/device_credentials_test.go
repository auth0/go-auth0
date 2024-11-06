package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestDeviceCredentials_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedDeviceCredential := &DeviceCredential{
		DeviceName: auth0.String("TestDevice"),
		Type:       auth0.String("public_key"),
		Value:      auth0.String("ABCD"),
		DeviceID:   auth0.String("test_device"),
		ClientID:   auth0.String("test_client"),
	}

	err := api.DeviceCredentials.Create(context.Background(), expectedDeviceCredential)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedDeviceCredential.GetID())

	t.Cleanup(func() {
		cleanupDeviceCredential(t, expectedDeviceCredential.GetID())
	})
}

func TestDeviceCredentials_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedDeviceCredential := givenADeviceCredential(t)

	deviceCredentialList, err := api.DeviceCredentials.List(context.Background(), IncludeFields("id"))

	assert.NoError(t, err)
	assert.Contains(t, deviceCredentialList.DeviceCredentials, &DeviceCredential{ID: expectedDeviceCredential.ID})
}

func TestDeviceCredentials_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedDeviceCredential := givenADeviceCredential(t)

	err := api.DeviceCredentials.Delete(context.Background(), expectedDeviceCredential.GetID())
	assert.NoError(t, err)

	actualDeviceCredentials, err := api.DeviceCredentials.List(context.Background())
	assert.NoError(t, err)
	assert.Empty(t, actualDeviceCredentials.DeviceCredentials)
}

func givenADeviceCredential(t *testing.T) *DeviceCredential {
	t.Helper()

	deviceCredential := &DeviceCredential{
		DeviceName: auth0.String("TestDevice"),
		Type:       auth0.String("refresh_token"),
		Value:      auth0.String("ABCD"),
		DeviceID:   auth0.String("test_device"),
		ClientID:   auth0.String("test_client"),
	}
	err := api.DeviceCredentials.Create(context.Background(), deviceCredential)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupDeviceCredential(t, deviceCredential.GetID())
	})

	return deviceCredential
}

func cleanupDeviceCredential(t *testing.T, id string) {
	t.Helper()

	err := api.DeviceCredentials.Delete(context.Background(), id)
	require.NoError(t, err)
}
