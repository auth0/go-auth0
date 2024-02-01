package management

import (
	"context"
	"encoding/json"
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

// PartialsPrompt to be used for Custom Prompt Partials.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations
type PartialsPrompt struct {
	FormContentStart      string `json:"form-content-start,omitempty"`
	FormContentEnd        string `json:"form-content-end,omitempty"`
	FormFooterStart       string `json:"form-footer-start,omitempty"`
	FormFooterEnd         string `json:"form-footer-end,omitempty"`
	SecondaryActionsStart string `json:"secondary-actions-start,omitempty"`
	SecondaryActionsEnd   string `json:"secondary-actions-end,omitempty"`

	// Segment for custom prompt
	Segment PartialsPromptSegment `json:"-"`
}

// MarshalJSON implements a custom Marshaler.
func (c *PartialsPrompt) MarshalJSON() ([]byte, error) {
	body := map[string]PartialsPrompt{string(c.Segment): *c}
	return json.Marshal(body)
}

// UnmarshalJSON implements a custom Unmarshaler.
func (c *PartialsPrompt) UnmarshalJSON(data []byte) error {
	var body map[string]struct{ PartialsPrompt }
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	for k, v := range body {
		*c = v.PartialsPrompt
		c.Segment = PartialsPromptSegment(k)
	}
	return nil
}

// PartialsPromptSegment defines the partials segment that we are managing.
type PartialsPromptSegment string

const (
	// PartialsPromptSignup represents the signup segment.
	PartialsPromptSignup PartialsPromptSegment = "signup"

	// PartialsPromptSignupID represents the signup-id segment.
	PartialsPromptSignupID PartialsPromptSegment = "signup-id"

	// PartialsPromptSignupPassword represents the signup-password segment.
	PartialsPromptSignupPassword PartialsPromptSegment = "signup-password"

	// PartialsPromptLogin represents the login segment.
	PartialsPromptLogin PartialsPromptSegment = "login"

	// PartialsPromptLoginID represents the login-id segment.
	PartialsPromptLoginID PartialsPromptSegment = "login-id"

	// PartialsPromptLoginPassword represents the login-password segment.
	PartialsPromptLoginPassword PartialsPromptSegment = "login-password"
)

var validPartialsPromptSegments = []PartialsPromptSegment{
	PartialsPromptSignup,
	PartialsPromptSignupID,
	PartialsPromptSignupPassword,
	PartialsPromptLogin,
	PartialsPromptLoginID,
	PartialsPromptLoginPassword,
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
func (m *PromptManager) CreatePartials(ctx context.Context, c *PartialsPrompt, opts ...RequestOption) error {
	if err := validatePartialsPromptSegment(c.Segment); err != nil {
		return err
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Segment), "partials"), c, opts...)
}

// UpdatePartials updates custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) UpdatePartials(ctx context.Context, c *PartialsPrompt, opts ...RequestOption) error {
	if err := validatePartialsPromptSegment(c.Segment); err != nil {
		return err
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Segment), "partials"), c, opts...)
}

// ReadPartials reads custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) ReadPartials(ctx context.Context, segment PartialsPromptSegment, opts ...RequestOption) (c *PartialsPrompt, err error) {
	if err := validatePartialsPromptSegment(segment); err != nil {
		return nil, err
	}
	c = &PartialsPrompt{Segment: segment}
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(segment), "partials"), c, opts...)
	return
}

// DeletePartials deletes custom prompt partials for a given segment.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) DeletePartials(ctx context.Context, c *PartialsPrompt, opts ...RequestOption) error {
	if err := validatePartialsPromptSegment(c.Segment); err != nil {
		return err
	}
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Segment), "partials"), &PartialsPrompt{}, opts...)
}

func validatePartialsPromptSegment(segment PartialsPromptSegment) error {
	for _, p := range validPartialsPromptSegments {
		if p == segment {
			return nil
		}
	}
	return fmt.Errorf("invalid custom segment: %s", segment)
}
