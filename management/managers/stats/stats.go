package stats

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

// GetActiveUsers Get active users count
//
// https://auth0.com/docs/api/management/v2/#!/Stats/get_active_users
func (m *Manager) GetActiveUsers(ctx context.Context, opts ...management.RequestOption) (float32, error) {
	var localVarReturnValue float32
	err := m.management.Request(ctx, "GET", m.management.URI("stats", "active-users"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// GetDaily Get daily stats
//
// https://auth0.com/docs/api/management/v2/#!/Stats/get_daily
func (m *Manager) GetDaily(ctx context.Context, opts ...management.RequestOption) ([]*models.StatsEntry, error) {
	var localVarReturnValue []*models.StatsEntry
	err := m.management.Request(ctx, "GET", m.management.URI("stats", "daily"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
