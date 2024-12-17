package emails

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

// GetProvider Get the email provider
//
// https://auth0.com/docs/api/management/v2/#!/Emails/get_provider
func (m *Manager) Get(ctx context.Context, opts ...management.RequestOption) (*models.EmailProvider, error) {
	var localVarReturnValue *models.EmailProvider
	err := m.management.Request(ctx, "GET", m.management.URI("emails", "provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchProvider Update the email provider
//
// https://auth0.com/docs/api/management/v2/#!/Emails/patch_provider
func (m *Manager) Update(ctx context.Context, emailProviderUpdate *models.EmailProviderUpdate, opts ...management.RequestOption) (*models.EmailProvider, error) {
	var localVarReturnValue *models.EmailProvider
	err := m.management.Request(ctx, "PATCH", m.management.URI("emails", "provider"), emailProviderUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostProvider Configure the email provider
//
// https://auth0.com/docs/api/management/v2/#!/Emails/post_provider
func (m *Manager) Configure(ctx context.Context, emailProviderCreate *models.EmailProviderCreate, opts ...management.RequestOption) (*models.EmailProvider, error) {
	var localVarReturnValue *models.EmailProvider
	err := m.management.Request(ctx, "POST", m.management.URI("emails", "provider"), emailProviderCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
