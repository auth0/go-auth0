package branding

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

// DeleteBrandingTheme Delete branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/delete_branding_theme
func (m *Manager) DeleteBrandingTheme(ctx context.Context, themeId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("branding", "themes", management.SafeString(themeId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniversalLogin Delete template for New Universal Login Experience
//
// https://auth0.com/docs/api/management/v2/#!/Branding/delete_universal_login
func (m *Manager) DeleteUniversalLogin(ctx context.Context, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("branding", "templates", "universal-login"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetBranding Get branding settings
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_branding
func (m *Manager) GetBranding(ctx context.Context, opts ...management.RequestOption) (*models.GetBranding200Response, error) {
	var localVarReturnValue *models.GetBranding200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetBrandingTheme Get branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_branding_theme
func (m *Manager) GetBrandingTheme(ctx context.Context, themeId string, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding", "themes", management.SafeString(themeId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetDefaultBrandingTheme Get default branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_default_branding_theme
func (m *Manager) GetDefaultBrandingTheme(ctx context.Context, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding", "themes", "default"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUniversalLogin Get template for New Universal Login Experience
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_universal_login
func (m *Manager) GetUniversalLogin(ctx context.Context, opts ...management.RequestOption) (*models.GetUniversalLogin200Response, error) {
	var localVarReturnValue *models.GetUniversalLogin200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding", "templates", "universal-login"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchBranding Update branding settings
//
// https://auth0.com/docs/api/management/v2/#!/Branding/patch_branding
func (m *Manager) PatchBranding(ctx context.Context, patchBrandingRequest *models.PatchBrandingRequest, opts ...management.RequestOption) (*models.GetBranding200Response, error) {
	var localVarReturnValue *models.GetBranding200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("branding"), patchBrandingRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchBrandingTheme Update branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/patch_branding_theme
func (m *Manager) PatchBrandingTheme(ctx context.Context, themeId string, postBrandingThemeRequest *models.PostBrandingThemeRequest, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("branding", "themes", management.SafeString(themeId)), postBrandingThemeRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostBrandingTheme Create branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/post_branding_theme
func (m *Manager) PostBrandingTheme(ctx context.Context, postBrandingThemeRequest *models.PostBrandingThemeRequest, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "POST", m.management.URI("branding", "themes"), postBrandingThemeRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutUniversalLogin Set template for New Universal Login Experience
//
// https://auth0.com/docs/api/management/v2/#!/Branding/put_universal_login
func (m *Manager) PutUniversalLogin(ctx context.Context, putUniversalLoginRequest *models.PutUniversalLoginRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "PUT", m.management.URI("branding", "templates", "universal-login"), putUniversalLoginRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
