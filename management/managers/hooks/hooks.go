package hooks

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

// Delete Delete a hook
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/delete_hooks_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("hooks", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSecrets Delete hook secrets
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/delete_secrets
func (m *Manager) DeleteSecrets(ctx context.Context, id string, requestBody []*string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("hooks", string(id), "secrets"), requestBody, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAll Get hooks
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/get_hooks
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetHooks200Response, error) {
	var localVarReturnValue *models.GetHooks200Response
	err := m.management.Request(ctx, "GET", m.management.URI("hooks"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Get Get a hook
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/get_hooks_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.Hook, error) {
	var localVarReturnValue *models.Hook
	err := m.management.Request(ctx, "GET", m.management.URI("hooks", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSecrets Get hook secrets
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/get_secrets
func (m *Manager) GetSecrets(ctx context.Context, id string, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "GET", m.management.URI("hooks", string(id), "secrets"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// Update Update a hook
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/patch_hooks_by_id
func (m *Manager) Update(ctx context.Context, id string, hookUpdate *models.HookUpdate, opts ...management.RequestOption) (*models.Hook, error) {
	var localVarReturnValue *models.Hook
	err := m.management.Request(ctx, "PATCH", m.management.URI("hooks", string(id)), hookUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateSecrets Update hook secrets
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/patch_secrets
func (m *Manager) UpdateSecrets(ctx context.Context, id string, requestBody map[string]interface{}, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "PATCH", m.management.URI("hooks", string(id), "secrets"), requestBody, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Create Create a hook
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/post_hooks
func (m *Manager) Create(ctx context.Context, hookCreate *models.HookCreate, opts ...management.RequestOption) (*models.Hook, error) {
	var localVarReturnValue *models.Hook
	err := m.management.Request(ctx, "POST", m.management.URI("hooks"), hookCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// AddSecrets Add hook secrets
//
// https://auth0.com/docs/api/management/v2/#!/Hooks/post_secrets
func (m *Manager) AddSecrets(ctx context.Context, id string, requestBody map[string]interface{}, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("hooks", string(id), "secrets"), requestBody, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
