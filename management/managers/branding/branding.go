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

// DeleteTheme Delete branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/delete_branding_theme
func (m *Manager) DeleteTheme(ctx context.Context, themeId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("branding", "themes", string(themeId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUniversalLoginTemplate Delete template for New Universal Login Experience
//
// https://auth0.com/docs/api/management/v2/#!/Branding/delete_universal_login
func (m *Manager) DeleteUniversalLoginTemplate(ctx context.Context, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("branding", "templates", "universal-login"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetSettings Get branding settings
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_branding
func (m *Manager) GetSettings(ctx context.Context, opts ...management.RequestOption) (*models.GetBranding200Response, error) {
	var localVarReturnValue *models.GetBranding200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetTheme Get branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_branding_theme
func (m *Manager) GetTheme(ctx context.Context, themeId string, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding", "themes", string(themeId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetDefaultTheme Get default branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_default_branding_theme
func (m *Manager) GetDefaultTheme(ctx context.Context, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding", "themes", "default"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUniversalLoginTemplate Get template for New Universal Login Experience
//
// https://auth0.com/docs/api/management/v2/#!/Branding/get_universal_login
func (m *Manager) GetUniversalLoginTemplate(ctx context.Context, opts ...management.RequestOption) (*models.GetUniversalLogin200Response, error) {
	var localVarReturnValue *models.GetUniversalLogin200Response
	err := m.management.Request(ctx, "GET", m.management.URI("branding", "templates", "universal-login"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateSettings Update branding settings
//
// https://auth0.com/docs/api/management/v2/#!/Branding/patch_branding
func (m *Manager) UpdateSettings(ctx context.Context, patchBrandingRequest *models.PatchBrandingRequest, opts ...management.RequestOption) (*models.GetBranding200Response, error) {
	var localVarReturnValue *models.GetBranding200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("branding"), patchBrandingRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateTheme Update branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/patch_branding_theme
func (m *Manager) UpdateTheme(ctx context.Context, themeId string, postBrandingThemeRequest *models.PostBrandingThemeRequest, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("branding", "themes", string(themeId)), postBrandingThemeRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// CreateTheme Create branding theme
//
// https://auth0.com/docs/api/management/v2/#!/Branding/post_branding_theme
func (m *Manager) CreateTheme(ctx context.Context, postBrandingThemeRequest *models.PostBrandingThemeRequest, opts ...management.RequestOption) (*models.PostBrandingTheme200Response, error) {
	var localVarReturnValue *models.PostBrandingTheme200Response
	err := m.management.Request(ctx, "POST", m.management.URI("branding", "themes"), postBrandingThemeRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetUniversalLoginTemplate Set template for New Universal Login Experience
//
// https://auth0.com/docs/api/management/v2/#!/Branding/put_universal_login
func (m *Manager) SetUniversalLoginTemplate(ctx context.Context, putUniversalLoginRequest *models.PutUniversalLoginRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "PUT", m.management.URI("branding", "templates", "universal-login"), putUniversalLoginRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
