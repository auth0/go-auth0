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

	// PromptLoginPasswordLess represents the login-passwordless prompt.
	PromptLoginPasswordLess PromptType = "login-passwordless"
)

var allowedPromptsWithPartials = []PromptType{
	PromptSignup,
	PromptSignupID,
	PromptSignupPassword,
	PromptLogin,
	PromptLoginID,
	PromptLoginPassword,
	PromptLoginPasswordLess,
}

// PromptType defines the prompt that we are managing.
type PromptType string

// ScreenName is a type that represents the name of a screen.
type ScreenName string

// InsertionPoint is a type that represents the insertion point of a screen.
type InsertionPoint string

const (
	// ScreenLogin represents the login screen.
	ScreenLogin ScreenName = "login"

	// ScreenLoginID represents the login-id screen.
	ScreenLoginID ScreenName = "login-id"

	// ScreenLoginPassword represents the login-password screen.
	ScreenLoginPassword ScreenName = "login-password"

	// ScreenSignup represents the signup screen.
	ScreenSignup ScreenName = "signup"

	// ScreenSignupID represents the signup-id screen.
	ScreenSignupID ScreenName = "signup-id"

	// ScreenSignupPassword represents the signup-password screen.
	ScreenSignupPassword ScreenName = "signup-password"

	// ScreenLoginPasswordlessSMSOTP represents the login-passwordless-sms-otp screen.
	ScreenLoginPasswordlessSMSOTP ScreenName = "login-passwordless-sms-otp"

	// ScreenLoginPasswordlessEmailCode represents the login-passwordless-email-code screen.
	ScreenLoginPasswordlessEmailCode ScreenName = "login-passwordless-email-code"
)

const (
	// InsertionPointFormContentStart represents the form-content-start insertion point.
	InsertionPointFormContentStart InsertionPoint = "form-content-start"

	// InsertionPointFormContentEnd represents the form-content-end insertion point.
	InsertionPointFormContentEnd InsertionPoint = "form-content-end"

	// InsertionPointFormFooterStart represents the form-footer-start insertion point.
	InsertionPointFormFooterStart InsertionPoint = "form-footer-start"

	// InsertionPointFormFooterEnd represents the form-footer-end insertion point.
	InsertionPointFormFooterEnd InsertionPoint = "form-footer-end"

	// InsertionPointSecondaryActionsStart represents the primary-actions-start insertion point.
	InsertionPointSecondaryActionsStart InsertionPoint = "secondary-actions-start"

	// InsertionPointSecondaryActionsEnd represents the primary-actions-end insertion point.
	InsertionPointSecondaryActionsEnd InsertionPoint = "secondary-actions-end"
)

// ScreenPartials is a map of insertion points to partials.
type ScreenPartials struct {
	// Define InsertionPoints for the screen partials here
	Content map[InsertionPoint]string
}

// PromptScreenPartials is a map of screen names to insertion points.
type PromptScreenPartials map[ScreenName]map[InsertionPoint]string

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
// Deprecated: Use [PromptScreenPartials] instead.
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
// Deprecated: Use [ SetPartials ] instead. The [ SetPartials ] method is preferred for setting prompt partials and provides a more consistent API.
//
// To create a partial with a different screen name and prompt name, use the [ SetPartials ] method with the [PromptScreenPartials] struct.
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
// Deprecated: Use [ SetPartials ] instead. The [ SetPartials ] method offers more flexibility and is the recommended approach for updating prompt partials.
//
// To update a partial with a different screen name and prompt name, use the [ SetPartials ] method with the [PromptScreenPartials] struct.
//
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) UpdatePartials(ctx context.Context, c *PromptPartials, opts ...RequestOption) error {
	if err := guardAgainstPromptTypesWithNoPartials(c.Prompt); err != nil {
		return err
	}

	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(c.Prompt), "partials"), c, opts...)
}

// GetPartials retrieves custom prompt partials for a given segment.
//
// See : https://auth0.com/docs/api/management/v2/prompts/get-partials
func (m *PromptManager) GetPartials(ctx context.Context, prompt PromptType, opts ...RequestOption) (c *PromptScreenPartials, err error) {
	if err := guardAgainstPromptTypesWithNoPartials(prompt); err != nil {
		return nil, err
	}

	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "partials"), &c, opts...)

	return
}

// SetPartials sets custom prompt partials for a given segment.
//
// See : https://auth0.com/docs/api/management/v2/prompts/put-partials
func (m *PromptManager) SetPartials(ctx context.Context, prompt PromptType, c *PromptScreenPartials, opts ...RequestOption) error {
	if err := guardAgainstPromptTypesWithNoPartials(prompt); err != nil {
		return err
	}

	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(prompt), "partials"), &c, opts...)
}

// ReadPartials reads custom prompt partials for a given segment.
//
// Deprecated: Use [ GetPartials ] instead. The [ GetPartials ] method provides the same functionality with improved support and additional features.
//
// If there are multiple screen partials for a prompt, this method will return only the first screen partial. To retrieve all screen partials for a prompt, use the [ GetPartials ] method.
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
// Deprecated: Use [ SetPartials ] with an empty [PromptScreenPartials] struct instead. The [ SetPartials ] method now handles deletion as well.
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
