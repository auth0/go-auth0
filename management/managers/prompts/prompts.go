package prompts

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

// GetCustomTextByLanguage Get custom text for a prompt
//
// https://auth0.com/docs/api/management/v2/#!/Prompts/get_custom_text_by_language
func (m *Manager) GetCustomTextByLanguage(ctx context.Context, prompt models.GetCustomTextByLanguagePromptParameter, language models.GetCustomTextByLanguageLanguageParameter, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "custom-text", string(language)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// GetPartials Get partials for a prompt
//
// https://auth0.com/docs/api/management/v2/#!/Prompts/get_partials
func (m *Manager) GetPartials(ctx context.Context, prompt models.GetPartialsPromptParameter, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "partials"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// Get Get prompt settings
//
// https://auth0.com/docs/api/management/v2/#!/Prompts/get_prompts
func (m *Manager) Get(ctx context.Context, opts ...management.RequestOption) (*models.PromptsSettings, error) {
	var localVarReturnValue *models.PromptsSettings
	err := m.management.Request(ctx, "GET", m.management.URI("prompts"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Update Update prompt settings
//
// https://auth0.com/docs/api/management/v2/#!/Prompts/patch_prompts
func (m *Manager) Update(ctx context.Context, promptsSettingsUpdate *models.PromptsSettingsUpdate, opts ...management.RequestOption) (*models.PromptsSettings, error) {
	var localVarReturnValue *models.PromptsSettings
	err := m.management.Request(ctx, "PATCH", m.management.URI("prompts"), promptsSettingsUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateCustomTextByLanguage Set custom text for a specific prompt
//
// https://auth0.com/docs/api/management/v2/#!/Prompts/put_custom_text_by_language
func (m *Manager) UpdateCustomTextByLanguage(ctx context.Context, prompt models.GetCustomTextByLanguagePromptParameter, language models.GetCustomTextByLanguageLanguageParameter, requestBody map[string]interface{}, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "PUT", m.management.URI("prompts", string(prompt), "custom-text", string(language)), requestBody, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePartials Set partials for a prompt
//
// https://auth0.com/docs/api/management/v2/#!/Prompts/put_partials
func (m *Manager) UpdatePartials(ctx context.Context, prompt models.GetPartialsPromptParameter, requestBody map[string]interface{}, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "PUT", m.management.URI("prompts", string(prompt), "partials"), requestBody, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
