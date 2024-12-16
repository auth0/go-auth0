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

// DeleteClientsById Delete a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/delete_clients_by_id
func (m *Manager) DeleteClientsById(ctx context.Context, clientId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("clients", management.SafeString(clientId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCredentialsByCredentialId Delete a client credential
//
// https://auth0.com/docs/api/management/v2/#!/Clients/delete_credentials_by_credential_id
func (m *Manager) DeleteCredentialsByCredentialId(ctx context.Context, clientId string, credentialId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("clients", management.SafeString(clientId), "credentials", management.SafeString(credentialId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetClients Get clients
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_clients
func (m *Manager) GetClients(ctx context.Context, opts ...management.RequestOption) (*models.GetClients200Response, error) {
	var localVarReturnValue *models.GetClients200Response
	err := m.management.Request(ctx, "GET", m.management.URI("clients"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetClientsById Get a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_clients_by_id
func (m *Manager) GetClientsById(ctx context.Context, clientId string, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "GET", m.management.URI("clients", management.SafeString(clientId)), nil, &localVarReturnValue, opts...)
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
	err := m.management.Request(ctx, "GET", m.management.URI("clients", management.SafeString(clientId), "credentials"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetCredentialsByCredentialId Get client credential details
//
// https://auth0.com/docs/api/management/v2/#!/Clients/get_credentials_by_credential_id
func (m *Manager) GetCredentialsByCredentialId(ctx context.Context, clientId string, credentialId string, opts ...management.RequestOption) (*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue *models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("clients", management.SafeString(clientId), "credentials", management.SafeString(credentialId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchClientsById Update a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/patch_clients_by_id
func (m *Manager) PatchClientsById(ctx context.Context, clientId string, clientUpdate *models.ClientUpdate, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "PATCH", m.management.URI("clients", management.SafeString(clientId)), clientUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchCredentialsByCredentialId Update a client credential
//
// https://auth0.com/docs/api/management/v2/#!/Clients/patch_credentials_by_credential_id
func (m *Manager) PatchCredentialsByCredentialId(ctx context.Context, clientId string, credentialId string, patchCredentialsByCredentialIdRequest *models.PatchCredentialsByCredentialIdRequest, opts ...management.RequestOption) (*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue *models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("clients", management.SafeString(clientId), "credentials", management.SafeString(credentialId)), patchCredentialsByCredentialIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostClients Create a client
//
// https://auth0.com/docs/api/management/v2/#!/Clients/post_clients
func (m *Manager) PostClients(ctx context.Context, clientCreate *models.ClientCreate, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "POST", m.management.URI("clients"), clientCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostCredentials Create a client credential
//
// https://auth0.com/docs/api/management/v2/#!/Clients/post_credentials
func (m *Manager) PostCredentials(ctx context.Context, clientId string, postCredentialsRequest *models.PostCredentialsRequest, opts ...management.RequestOption) (*models.GetCredentials200ResponseInner, error) {
	var localVarReturnValue *models.GetCredentials200ResponseInner
	err := m.management.Request(ctx, "POST", m.management.URI("clients", management.SafeString(clientId), "credentials"), postCredentialsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostRotateSecret Rotate a client secret
//
// https://auth0.com/docs/api/management/v2/#!/Clients/post_rotate_secret
func (m *Manager) PostRotateSecret(ctx context.Context, clientId string, opts ...management.RequestOption) (*models.Client, error) {
	var localVarReturnValue *models.Client
	err := m.management.Request(ctx, "POST", m.management.URI("clients", management.SafeString(clientId), "rotate-secret"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
