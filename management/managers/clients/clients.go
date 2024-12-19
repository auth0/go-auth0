package clients

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

// Delete Delete a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/delete_clients_by_id
func (m *Manager) Delete(ctx context.Context, clientId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("clients", string(clientId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCredential Delete a client credential
//
// https://auth0.com/docs/api/management/v2/#!/Clients/delete_credentials_by_credential_id
func (m *Manager) DeleteCredential(ctx context.Context, clientId string, credentialId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("clients", string(clientId), "credentials", string(credentialId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAll Get clients
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_clients
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetClients200Response, error) {
	var localVarReturnValue *models.GetClients200Response
	err := m.management.Request(ctx, "GET", m.management.URI("clients"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Get Get client by ID
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_clients_by_id
func (m *Manager) Get(ctx context.Context, clientId string, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "GET", m.management.URI("clients", string(clientId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetCredentials Get client credentials
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_credentials
func (m *Manager) GetCredentials(ctx context.Context, clientId string, opts ...management.RequestOption) ([]*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue []*models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("clients", string(clientId), "credentials"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetCredential Get client credential details
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_credentials_by_credential_id
func (m *Manager) GetCredential(ctx context.Context, clientId string, credentialId string, opts ...management.RequestOption) (*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue *models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("clients", string(clientId), "credentials", string(credentialId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Update Update a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/patch_clients_by_id
func (m *Manager) Update(ctx context.Context, clientId string, clientUpdate *models.ClientUpdate, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "PATCH", m.management.URI("clients", string(clientId)), clientUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateCredential Update a client credential
//
// https://auth0.com/docs/api/management/v2/#!/Clients/patch_credentials_by_credential_id
func (m *Manager) UpdateCredential(ctx context.Context, clientId string, credentialId string, patchCredentialsByCredentialIdRequest *models.PatchCredentialsByCredentialIdRequest, opts ...management.RequestOption) (*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue *models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("clients", string(clientId), "credentials", string(credentialId)), patchCredentialsByCredentialIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Create Create a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/post_clients
func (m *Manager) Create(ctx context.Context, clientCreate *models.ClientCreate, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "POST", m.management.URI("clients"), clientCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// CreateCredential Create a client credential
//
// https://auth0.com/docs/api/management/v2/#!/Clients/post_credentials
func (m *Manager) CreateCredential(ctx context.Context, clientId string, postCredentialsRequest *models.PostCredentialsRequest, opts ...management.RequestOption) (*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue *models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "POST", m.management.URI("clients", string(clientId), "credentials"), postCredentialsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// RotateClientSecret Rotate a client secret
//
// https://auth0.com/docs/api/management/v2/#!/Clients/post_rotate_secret
func (m *Manager) RotateClientSecret(ctx context.Context, clientId string, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "POST", m.management.URI("clients", string(clientId), "rotate-secret"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
