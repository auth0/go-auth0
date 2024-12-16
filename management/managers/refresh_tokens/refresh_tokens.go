package refreshTokens

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

// DeleteRefreshToken Delete a refresh tokens
//
// https://auth0.com/docs/api/management/v2/#!/RefreshTokens/delete_refresh_token
func (m *Manager) DeleteRefreshToken(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("refresh-tokens", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetRefreshToken Get a refresh token
//
// https://auth0.com/docs/api/management/v2/#!/RefreshTokens/get_refresh_token
func (m *Manager) GetRefreshToken(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetRefreshToken200Response, error) {
	var localVarReturnValue *models.GetRefreshToken200Response
	err := m.management.Request(ctx, "GET", m.management.URI("refresh-tokens", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
