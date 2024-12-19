package usersByEmail

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

// GetByEmail Search Users by Email
//
// https://auth0.com/docs/api/management/v2/#!/UsersByEmail/get_users_by_email
func (m *Manager) GetByEmail(ctx context.Context, opts ...management.RequestOption) ([]*models.GetUsers200ResponseOneOfInner, error) {
	var localVarReturnValue []*models.GetUsers200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("users-by-email"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
