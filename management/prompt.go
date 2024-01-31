package management

import (
	"context"
	"fmt"
)

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

// CustomPrompt to be used for Custom Prompt Partials.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations
type CustomPrompt struct {
	FormContentStart      string `json:"form-content-start,omitempty"`
	FormContentEnd        string `json:"form-content-end,omitempty"`
	FormFooterStart       string `json:"form-footer-start,omitempty"`
	FormFooterEnd         string `json:"form-footer-end,omitempty"`
	SecondaryActionsStart string `json:"secondary-actions-start,omitempty"`
	SecondaryActionsEnd   string `json:"secondary-actions-end,omitempty"`

	// Segment for custom prompt
	Segment CustomPromptSegment `json:"-"`
}

// CustomPromptSegment defines the partials segment that we are managing.
type CustomPromptSegment string

const (
	CustomPromptSignup         CustomPromptSegment = "signup"
	CustomPromptSignupID       CustomPromptSegment = "signup-id"
	CustomPromptSignupPassword CustomPromptSegment = "signup-password"
	CustomPromptLogin          CustomPromptSegment = "login"
	CustomPromptLoginID        CustomPromptSegment = "login-id"
	CustomPromptLoginPassword  CustomPromptSegment = "login-password"
)

var validCustomPromptSegments = []CustomPromptSegment{
	CustomPromptSignup,
	CustomPromptSignupID,
	CustomPromptSignupPassword,
	CustomPromptLogin,
	CustomPromptLoginID,
	CustomPromptLoginPassword,
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

// CreateCustomPartials creates new custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) CreateCustomPartials(ctx context.Context, c *CustomPrompt, opts ...RequestOption) error {
	if err := validateCustomPromptSegment(c.Segment); err != nil {
		return err
	}
	return m.management.Request(ctx, "POST", m.management.URI("prompts", string(c.Segment), "partials"), c, opts...)
}

// UpdateCustomPartials updates custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) UpdateCustomPartials(ctx context.Context, c *CustomPrompt, opts ...RequestOption) error {
	if err := validateCustomPromptSegment(c.Segment); err != nil {
		return err
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Segment), "partials"), c, opts...)
}

// ReadCustomPartials reads custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) ReadCustomPartials(ctx context.Context, prompt CustomPromptSegment, opts ...RequestOption) (c *CustomPrompt, err error) {
	if err := validateCustomPromptSegment(prompt); err != nil {
		return nil, err
	}
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "partials"), &c, opts...)
	return
}

// DeleteCustomPartials deletes custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) DeleteCustomPartials(ctx context.Context, c *CustomPrompt, opts ...RequestOption) error {
	if err := validateCustomPromptSegment(c.Segment); err != nil {
		return err
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Segment), "partials"), &CustomPrompt{}, opts...)
}

func validateCustomPromptSegment(segment CustomPromptSegment) error {
	for _, p := range validCustomPromptSegments {
		if p == segment {
			return nil
		}
	}
	return fmt.Errorf("invalid custom segment: %s", segment)
}
