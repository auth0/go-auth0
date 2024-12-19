package sessions

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

// Delete Delete session
//
// https://auth0.com/docs/api/management/v2/#!/Sessions/delete_session
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("sessions", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Get Get session
//
// https://auth0.com/docs/api/management/v2/#!/Sessions/get_session
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetSession200Response, error) {
	var localVarReturnValue *models.GetSession200Response
	err := m.management.Request(ctx, "GET", m.management.URI("sessions", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
