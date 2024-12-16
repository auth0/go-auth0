package customDomains

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

// DeleteCustomDomainsById Delete custom domain configuration
//
// https://auth0.com/docs/api/management/v2/#!/CustomDomains/delete_custom_domains_by_id
func (m *Manager) DeleteCustomDomainsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("custom-domains", management.SafeString(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetCustomDomains Get custom domains configurations
//
// https://auth0.com/docs/api/management/v2/#!/CustomDomains/get_custom_domains
func (m *Manager) GetCustomDomains(ctx context.Context, opts ...management.RequestOption) ([]*models.CustomDomain, error) {
	var localVarReturnValue []*models.CustomDomain
	err := m.management.Request(ctx, "GET", m.management.URI("custom-domains"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetCustomDomainsById Get custom domain configuration
//
// https://auth0.com/docs/api/management/v2/#!/CustomDomains/get_custom_domains_by_id
func (m *Manager) GetCustomDomainsById(ctx context.Context, id string, opts ...management.RequestOption) (*models.CustomDomain, error) {
	var localVarReturnValue *models.CustomDomain
	err := m.management.Request(ctx, "GET", m.management.URI("custom-domains", management.SafeString(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchCustomDomainsById Update custom domain configuration
//
// https://auth0.com/docs/api/management/v2/#!/CustomDomains/patch_custom_domains_by_id
func (m *Manager) PatchCustomDomainsById(ctx context.Context, id string, patchCustomDomainsByIdRequest *models.PatchCustomDomainsByIdRequest, opts ...management.RequestOption) (*models.PostCustomDomains201Response, error) {
	var localVarReturnValue *models.PostCustomDomains201Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("custom-domains", management.SafeString(id)), patchCustomDomainsByIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostCustomDomains Configure a new custom domain
//
// https://auth0.com/docs/api/management/v2/#!/CustomDomains/post_custom_domains
func (m *Manager) PostCustomDomains(ctx context.Context, postCustomDomainsRequest *models.PostCustomDomainsRequest, opts ...management.RequestOption) (*models.PostCustomDomains201Response, error) {
	var localVarReturnValue *models.PostCustomDomains201Response
	err := m.management.Request(ctx, "POST", m.management.URI("custom-domains"), postCustomDomainsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostVerify Verify a custom domain
//
// https://auth0.com/docs/api/management/v2/#!/CustomDomains/post_verify
func (m *Manager) PostVerify(ctx context.Context, id string, opts ...management.RequestOption) (*models.PostVerify200Response, error) {
	var localVarReturnValue *models.PostVerify200Response
	err := m.management.Request(ctx, "POST", m.management.URI("custom-domains", management.SafeString(id), "verify"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
