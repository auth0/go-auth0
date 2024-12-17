package keys

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

// GetSigningKey Get an Application Signing Key by its key id
//
// https://auth0.com/docs/api/management/v2/#!/Keys/get_signing_key
func (m *Manager) Get(ctx context.Context, kid string, opts ...management.RequestOption) (*models.GetSigningKeys200ResponseInner, error) {
	var localVarReturnValue *models.GetSigningKeys200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("keys", "signing", string(kid)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSigningKeys Get all Application Signing Keys
//
// https://auth0.com/docs/api/management/v2/#!/Keys/get_signing_keys
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) ([]*models.GetSigningKeys200ResponseInner, error) {
	var localVarReturnValue []*models.GetSigningKeys200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("keys", "signing"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostSigningKeys Rotate the Application Signing Key
//
// https://auth0.com/docs/api/management/v2/#!/Keys/post_signing_keys
func (m *Manager) Rotate(ctx context.Context, opts ...management.RequestOption) (*models.PostSigningKeys201Response, error) {
	var localVarReturnValue *models.PostSigningKeys201Response
	err := m.management.Request(ctx, "POST", m.management.URI("keys", "signing", "rotate"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutSigningKeys Revoke an Application Signing Key by its key id
//
// https://auth0.com/docs/api/management/v2/#!/Keys/put_signing_keys
func (m *Manager) Revoke(ctx context.Context, kid string, opts ...management.RequestOption) (*models.PutSigningKeys200Response, error) {
	var localVarReturnValue *models.PutSigningKeys200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("keys", "signing", string(kid), "revoke"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
