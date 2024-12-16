package clientGrants

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

// DeleteClientGrantsById Delete client grant
//
// https://auth0.com/docs/api/management/v2/#!/ClientGrants/delete_client_grants_by_id
func (m *Manager) DeleteClientGrantsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("client-grants", management.SafeString(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetClientGrants Get client grants
//
// https://auth0.com/docs/api/management/v2/#!/ClientGrants/get_client_grants
func (m *Manager) GetClientGrants(ctx context.Context, opts ...management.RequestOption) (*models.GetClientGrants200Response, error) {
	var localVarReturnValue *models.GetClientGrants200Response
	err := m.management.Request(ctx, "GET", m.management.URI("client-grants"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchClientGrantsById Update client grant
//
// https://auth0.com/docs/api/management/v2/#!/ClientGrants/patch_client_grants_by_id
func (m *Manager) PatchClientGrantsById(ctx context.Context, id string, patchClientGrantsByIdRequest *models.PatchClientGrantsByIdRequest, opts ...management.RequestOption) (*models.ClientGrant, error) {
	var localVarReturnValue *models.ClientGrant
	err := m.management.Request(ctx, "PATCH", m.management.URI("client-grants", management.SafeString(id)), patchClientGrantsByIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostClientGrants Create client grant
//
// https://auth0.com/docs/api/management/v2/#!/ClientGrants/post_client_grants
func (m *Manager) PostClientGrants(ctx context.Context, clientGrantCreate *models.ClientGrantCreate, opts ...management.RequestOption) (*models.ClientGrant, error) {
	var localVarReturnValue *models.ClientGrant
	err := m.management.Request(ctx, "POST", m.management.URI("client-grants"), clientGrantCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
