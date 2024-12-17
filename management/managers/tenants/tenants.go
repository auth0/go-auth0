package tenants

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

// PatchSettings Update tenant settings
//
// https://auth0.com/docs/api/management/v2/#!/Tenants/patch_settings
func (m *Manager) UpdateSettings(ctx context.Context, tenantSettingsUpdate *models.TenantSettingsUpdate, opts ...management.RequestOption) (*models.TenantSettings, error) {
	var localVarReturnValue *models.TenantSettings
	err := m.management.Request(ctx, "PATCH", m.management.URI("tenants", "settings"), tenantSettingsUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// TenantSettingsRoute Get tenant settings
//
// https://auth0.com/docs/api/management/v2/#!/Tenants/tenant_settings_route
func (m *Manager) GetSettings(ctx context.Context, opts ...management.RequestOption) (*models.TenantSettings, error) {
	var localVarReturnValue *models.TenantSettings
	err := m.management.Request(ctx, "GET", m.management.URI("tenants", "settings"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
