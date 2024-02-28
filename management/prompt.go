package management

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	// PromptSignup represents the signup prompt.
	PromptSignup PromptType = "signup"

	// PromptSignupID represents the signup-id prompt.
	PromptSignupID PromptType = "signup-id"

	// PromptSignupPassword represents the signup-password prompt.
	PromptSignupPassword PromptType = "signup-password"

	// PromptLogin represents the login prompt.
	PromptLogin PromptType = "login"

	// PromptLoginID represents the login-id prompt.
	PromptLoginID PromptType = "login-id"

	// PromptLoginPassword represents the login-password prompt.
	PromptLoginPassword PromptType = "login-password"
)

var allowedPromptsWithPartials = []PromptType{
	PromptSignup,
	PromptSignupID,
	PromptSignupPassword,
	PromptLogin,
	PromptLoginID,
	PromptLoginPassword,
}

// PromptType defines the prompt that we are managing.
type PromptType string

// Prompt is used within the Login Page.
//
// See: https://auth0.com/docs/customize/universal-login-pages/customize-login-text-prompts
type Prompt struct {
	// Which login experience to use. Can be `new` or `classic`.
	UniversalLoginExperience string `json:"universal_login_experience,omitempty"`

	// IdentifierFirst determines if the login screen prompts for just the identifier, identifier and password first.
	IdentifierFirst *bool `json:"identifier_first,omitempty"`

	// WebAuthnPlatformFirstFactor determines if the login screen uses identifier and biometrics first.
	WebAuthnPlatformFirstFactor *bool `json:"webauthn_platform_first_factor,omitempty"`
}

// PromptPartials to be used for Custom Prompt Partials.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations
type PromptPartials struct {
	FormContentStart      string `json:"form-content-start,omitempty"`
	FormContentEnd        string `json:"form-content-end,omitempty"`
	FormFooterStart       string `json:"form-footer-start,omitempty"`
	FormFooterEnd         string `json:"form-footer-end,omitempty"`
	SecondaryActionsStart string `json:"secondary-actions-start,omitempty"`
	SecondaryActionsEnd   string `json:"secondary-actions-end,omitempty"`

	Prompt PromptType `json:"-"`
}

// MarshalJSON implements a custom [json.Marshaler].
func (c *PromptPartials) MarshalJSON() ([]byte, error) {
	body := map[string]PromptPartials{
		string(c.Prompt): *c,
	}
	return json.Marshal(body)
}

// UnmarshalJSON implements a custom [json.Unmarshaler].
func (c *PromptPartials) UnmarshalJSON(data []byte) error {
	var body map[string]struct {
		PromptPartials
	}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	for k, v := range body {
		*c = v.PromptPartials
		c.Prompt = PromptType(k)
	}

	return nil
}

// PromptManager is used for managing a Prompt.
type PromptManager manager

// Read retrieves prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_prompts
func (m *PromptManager) Read(ctx context.Context, opts ...RequestOption) (p *Prompt, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("prompts"), &p, opts...)
	return
}

// Update prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/patch_prompts
func (m *PromptManager) Update(ctx context.Context, p *Prompt, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("prompts"), p, opts...)
}

// CustomText retrieves the custom text for a specific prompt and language.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) CustomText(ctx context.Context, p string, l string, opts ...RequestOption) (t map[string]interface{}, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", p, "custom-text", l), &t, opts...)
	return
}

// SetCustomText sets the custom text for a specific prompt. Existing texts will be overwritten.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) SetCustomText(ctx context.Context, p string, l string, b map[string]interface{}, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PUT", m.management.URI("prompts", p, "custom-text", l), &b, opts...)
	return
}

// CreatePartials creates new custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) CreatePartials(ctx context.Context, c *PromptPartials, opts ...RequestOption) error {
	if err := guardAgainstPromptTypesWithNoPartials(c.Prompt); err != nil {
		return err
	}

	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Prompt), "partials"), c, opts...)
}

// UpdatePartials updates custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) UpdatePartials(ctx context.Context, c *PromptPartials, opts ...RequestOption) error {
	if err := guardAgainstPromptTypesWithNoPartials(c.Prompt); err != nil {
		return err
	}

	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Prompt), "partials"), c, opts...)
}

// ReadPartials reads custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) ReadPartials(ctx context.Context, prompt PromptType, opts ...RequestOption) (c *PromptPartials, err error) {
	if err := guardAgainstPromptTypesWithNoPartials(prompt); err != nil {
		return nil, err
	}

	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "partials"), &c, opts...)

	if c == nil {
		c = &PromptPartials{
			Prompt: prompt,
		}
	}

	return
}

// DeletePartials deletes custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) DeletePartials(ctx context.Context, prompt PromptType, opts ...RequestOption) error {
	if err := guardAgainstPromptTypesWithNoPartials(prompt); err != nil {
		return err
	}

	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(prompt), "partials"), &PromptPartials{Prompt: prompt}, opts...)
}

func guardAgainstPromptTypesWithNoPartials(prompt PromptType) error {
	for _, p := range allowedPromptsWithPartials {
		if p == prompt {
			return nil
		}
	}

	return fmt.Errorf("cannot customize partials for prompt: %q", prompt)
}
