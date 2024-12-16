package grants

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

// DeleteGrantsById Delete a grant by id
//
// https://auth0.com/docs/api/management/v2/#!/Grants/delete_grants_by_id
func (m *Manager) DeleteGrantsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("grants", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGrantsByUserId Delete a grant by user_id
//
// https://auth0.com/docs/api/management/v2/#!/Grants/delete_grants_by_user_id
func (m *Manager) DeleteGrantsByUserId(ctx context.Context, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("grants"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetGrants Get grants
//
// https://auth0.com/docs/api/management/v2/#!/Grants/get_grants
func (m *Manager) GetGrants(ctx context.Context, opts ...management.RequestOption) (*models.GetGrants200Response, error) {
	var localVarReturnValue *models.GetGrants200Response
	err := m.management.Request(ctx, "GET", m.management.URI("grants"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
