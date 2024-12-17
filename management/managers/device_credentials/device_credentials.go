package deviceCredentials

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/auth0/go-auth0/management/models"
)

// Manager defines the base manager interface
type Manager struct {
	management *management.Management
}

// NewManager creates a new manager for  operations
func NewManager(mgmt *management.Management) *Manager {
	return &Manager{
		management: mgmt,
	}
}

// DeleteDeviceCredentialsById Delete a device credential
//
// https://auth0.com/docs/api/management/v2/#!/DeviceCredentials/delete_device_credentials_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("device-credentials", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetDeviceCredentials Retrieve device credentials
//
// https://auth0.com/docs/api/management/v2/#!/DeviceCredentials/get_device_credentials
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetDeviceCredentials200Response, error) {
	var localVarReturnValue *models.GetDeviceCredentials200Response
	err := m.management.Request(ctx, "GET", m.management.URI("device-credentials"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostDeviceCredentials Create a device public key credential
//
// https://auth0.com/docs/api/management/v2/#!/DeviceCredentials/post_device_credentials
func (m *Manager) CreatePublicKey(ctx context.Context, deviceCredentialCreate *models.DeviceCredentialCreate, opts ...management.RequestOption) (*models.PostDeviceCredentials201Response, error) {
	var localVarReturnValue *models.PostDeviceCredentials201Response
	err := m.management.Request(ctx, "POST", m.management.URI("device-credentials"), deviceCredentialCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
