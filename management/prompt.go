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

// MarshalJSON implements a custom Marshaler
func (c *CustomPrompt) MarshalJSON() ([]byte, error) {
	body := map[string]CustomPrompt{string(c.Segment): *c}
	return json.Marshal(body)
}

// UnmarshalJSON implements a custom Unmarshaler
func (c *CustomPrompt) UnmarshalJSON(data []byte) error {
	var body map[string]map[string]string
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	for k, v := range body {
		c.FormContentStart = v["form-content-start"]
		c.FormFooterEnd = v["form-content-end"]
		c.FormFooterStart = v["form-footer-start"]
		c.FormFooterEnd = v["form-footer-end"]
		c.SecondaryActionsStart = v["secondary-actions-start"]
		c.SecondaryActionsEnd = v["secondary-actions-end"]
		c.Segment = CustomPromptSegment(k)
	}
	return nil
}

// CustomPromptSegment defines the partials segment that we are managing.
type CustomPromptSegment string

const (
	// CustomPromptSignup represents the signup segment
	CustomPromptSignup CustomPromptSegment = "signup"

	// CustomPromptSignupID represents the signup-id segment
	CustomPromptSignupID CustomPromptSegment = "signup-id"

	// CustomPromptSignupPassword represents the signup-password segment
	CustomPromptSignupPassword CustomPromptSegment = "signup-password"

	// CustomPromptLogin represents the login segment
	CustomPromptLogin CustomPromptSegment = "login"

	// CustomPromptLoginID represents the login-id segment
	CustomPromptLoginID CustomPromptSegment = "login-id"

	// CustomPromptLoginPassword represents the login-password segment
	CustomPromptLoginPassword CustomPromptSegment = "login-password"
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
	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Segment), "partials"), c, opts...)
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
func (m *PromptManager) ReadCustomPartials(ctx context.Context, segment CustomPromptSegment, opts ...RequestOption) (c *CustomPrompt, err error) {
	if err := validateCustomPromptSegment(segment); err != nil {
		return nil, err
	}
	c = &CustomPrompt{Segment: segment}
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(segment), "partials"), c, opts...)
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
