package blacklists

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

// GetAll Get blacklisted tokens
//
// https://auth0.com/docs/api/management/v2/#!/Blacklists/get_tokens
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) ([]*models.Token, error) {
	var localVarReturnValue []*models.Token
	err := m.management.Request(ctx, "GET", m.management.URI("blacklists", "tokens"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Add Blacklist a token
//
// https://auth0.com/docs/api/management/v2/#!/Blacklists/post_tokens
func (m *Manager) Add(ctx context.Context, token *models.Token, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("blacklists", "tokens"), token, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
