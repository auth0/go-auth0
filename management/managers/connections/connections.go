package connections

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

// Delete Delete a connection
//
// https://auth0.com/docs/api/management/v2/#!/Connections/delete_connections_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("connections", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteScimConfiguration Delete a connection&#39;s SCIM configuration
//
// https://auth0.com/docs/api/management/v2/#!/Connections/delete_scim_configuration
func (m *Manager) DeleteScimConfiguration(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("connections", string(id), "scim-configuration"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteScimToken Delete a connection&#39;s SCIM token
//
// https://auth0.com/docs/api/management/v2/#!/Connections/delete_tokens_by_token_id
func (m *Manager) DeleteScimToken(ctx context.Context, id string, tokenId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("connections", string(id), "scim-configuration", "tokens", string(tokenId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByEmail Delete a connection user
//
// https://auth0.com/docs/api/management/v2/#!/Connections/delete_users_by_email
func (m *Manager) DeleteUserByEmail(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("connections", string(id), "users"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAll Get all connections
//
// https://auth0.com/docs/api/management/v2/#!/Connections/get_connections
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetConnections200Response, error) {
	var localVarReturnValue *models.GetConnections200Response
	err := m.management.Request(ctx, "GET", m.management.URI("connections"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Get Get a connection
//
// https://auth0.com/docs/api/management/v2/#!/Connections/get_connections_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.Connection, error) {
	var localVarReturnValue *models.Connection
	err := m.management.Request(ctx, "GET", m.management.URI("connections", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetDefaultScimMapping Get a connection&#39;s default SCIM mapping
//
// https://auth0.com/docs/api/management/v2/#!/Connections/get_default_mapping
func (m *Manager) GetDefaultScimMapping(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetDefaultMapping200Response, error) {
	var localVarReturnValue *models.GetDefaultMapping200Response
	err := m.management.Request(ctx, "GET", m.management.URI("connections", string(id), "scim-configuration", "default-mapping"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetScimConfiguration Get a connection&#39;s SCIM configuration
//
// https://auth0.com/docs/api/management/v2/#!/Connections/get_scim_configuration
func (m *Manager) GetScimConfiguration(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetScimConfiguration200Response, error) {
	var localVarReturnValue *models.GetScimConfiguration200Response
	err := m.management.Request(ctx, "GET", m.management.URI("connections", string(id), "scim-configuration"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetScimTokens Get a connection&#39;s SCIM tokens
//
// https://auth0.com/docs/api/management/v2/#!/Connections/get_scim_tokens
func (m *Manager) GetScimTokens(ctx context.Context, id string, opts ...management.RequestOption) ([]*models.GetScimTokens200ResponseInner, error) {
	var localVarReturnValue []*models.GetScimTokens200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("connections", string(id), "scim-configuration", "tokens"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// CheckStatus Check connection status
//
// https://auth0.com/docs/api/management/v2/#!/Connections/get_status
func (m *Manager) CheckStatus(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "GET", m.management.URI("connections", string(id), "status"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Update Update a connection
//
// https://auth0.com/docs/api/management/v2/#!/Connections/patch_connections_by_id
func (m *Manager) Update(ctx context.Context, id string, connectionUpdate *models.ConnectionUpdate, opts ...management.RequestOption) (*models.Connection, error) {
	var localVarReturnValue *models.Connection
	err := m.management.Request(ctx, "PATCH", m.management.URI("connections", string(id)), connectionUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateScimConfiguration Patch a connection&#39;s SCIM configuration
//
// https://auth0.com/docs/api/management/v2/#!/Connections/patch_scim_configuration
func (m *Manager) UpdateScimConfiguration(ctx context.Context, id string, patchScimConfigurationRequest *models.PatchScimConfigurationRequest, opts ...management.RequestOption) (*models.GetScimConfiguration200Response, error) {
	var localVarReturnValue *models.GetScimConfiguration200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("connections", string(id), "scim-configuration"), patchScimConfigurationRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Create Create a connection
//
// https://auth0.com/docs/api/management/v2/#!/Connections/post_connections
func (m *Manager) Create(ctx context.Context, connectionCreate *models.ConnectionCreate, opts ...management.RequestOption) (*models.Connection, error) {
	var localVarReturnValue *models.Connection
	err := m.management.Request(ctx, "POST", m.management.URI("connections"), connectionCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// CreateScimConfiguration Create a SCIM configuration
//
// https://auth0.com/docs/api/management/v2/#!/Connections/post_scim_configuration
func (m *Manager) CreateScimConfiguration(ctx context.Context, id string, postScimConfigurationRequest *models.PostScimConfigurationRequest, opts ...management.RequestOption) (*models.GetScimConfiguration200Response, error) {
	var localVarReturnValue *models.GetScimConfiguration200Response
	err := m.management.Request(ctx, "POST", m.management.URI("connections", string(id), "scim-configuration"), postScimConfigurationRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// CreateScimToken Create a SCIM Token
//
// https://auth0.com/docs/api/management/v2/#!/Connections/post_scim_token
func (m *Manager) CreateScimToken(ctx context.Context, id string, postScimTokenRequest *models.PostScimTokenRequest, opts ...management.RequestOption) (*models.PostScimToken201Response, error) {
	var localVarReturnValue *models.PostScimToken201Response
	err := m.management.Request(ctx, "POST", m.management.URI("connections", string(id), "scim-configuration", "tokens"), postScimTokenRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
