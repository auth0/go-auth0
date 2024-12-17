package userBlocks

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

// DeleteUserBlocks Unblock by identifier
//
// https://auth0.com/docs/api/management/v2/#!/UserBlocks/delete_user_blocks
func (m *Manager) DeleteAll(ctx context.Context, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("user-blocks"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserBlocksById Unblock a user
//
// https://auth0.com/docs/api/management/v2/#!/UserBlocks/delete_user_blocks_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("user-blocks", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetUserBlocks Get blocks by identifier
//
// https://auth0.com/docs/api/management/v2/#!/UserBlocks/get_user_blocks
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.UserBlock, error) {
	var localVarReturnValue *models.UserBlock
	err := m.management.Request(ctx, "GET", m.management.URI("user-blocks"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUserBlocksById Get a user&#39;s blocks
//
// https://auth0.com/docs/api/management/v2/#!/UserBlocks/get_user_blocks_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.UserBlock, error) {
	var localVarReturnValue *models.UserBlock
	err := m.management.Request(ctx, "GET", m.management.URI("user-blocks", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
