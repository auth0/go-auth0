package resourceServers

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

// DeleteResourceServersById Delete a resource server
//
// https://auth0.com/docs/api/management/v2/#!/ResourceServers/delete_resource_servers_by_id
func (m *Manager) DeleteResourceServersById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("resource-servers", management.SafeString(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetResourceServers Get resource servers
//
// https://auth0.com/docs/api/management/v2/#!/ResourceServers/get_resource_servers
func (m *Manager) GetResourceServers(ctx context.Context, opts ...management.RequestOption) (*models.GetResourceServers200Response, error) {
	var localVarReturnValue *models.GetResourceServers200Response
	err := m.management.Request(ctx, "GET", m.management.URI("resource-servers"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetResourceServersById Get a resource server
//
// https://auth0.com/docs/api/management/v2/#!/ResourceServers/get_resource_servers_by_id
func (m *Manager) GetResourceServersById(ctx context.Context, id string, opts ...management.RequestOption) (*models.ResourceServer, error) {
	var localVarReturnValue *models.ResourceServer
	err := m.management.Request(ctx, "GET", m.management.URI("resource-servers", management.SafeString(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchResourceServersById Update a resource server
//
// https://auth0.com/docs/api/management/v2/#!/ResourceServers/patch_resource_servers_by_id
func (m *Manager) PatchResourceServersById(ctx context.Context, id string, resourceServerUpdate *models.ResourceServerUpdate, opts ...management.RequestOption) (*models.ResourceServer, error) {
	var localVarReturnValue *models.ResourceServer
	err := m.management.Request(ctx, "PATCH", m.management.URI("resource-servers", management.SafeString(id)), resourceServerUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostResourceServers Create a resource server
//
// https://auth0.com/docs/api/management/v2/#!/ResourceServers/post_resource_servers
func (m *Manager) PostResourceServers(ctx context.Context, resourceServerCreate *models.ResourceServerCreate, opts ...management.RequestOption) (*models.ResourceServer, error) {
	var localVarReturnValue *models.ResourceServer
	err := m.management.Request(ctx, "POST", m.management.URI("resource-servers"), resourceServerCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
