package management

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

// PromptManager is used for managing a Prompt.
type PromptManager struct {
	*Management
}

func newPromptManager(m *Management) *PromptManager {
	return &PromptManager{m}
}

// Read retrieves prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_prompts
func (m *PromptManager) Read(opts ...RequestOption) (p *Prompt, err error) {
	err = m.Request("GET", m.URI("prompts"), &p, opts...)
	return
}

// Update prompts settings.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/patch_prompts
func (m *PromptManager) Update(p *Prompt, opts ...RequestOption) error {
	return m.Request("PATCH", m.URI("prompts"), p, opts...)
}

// CustomText retrieves the custom text for a specific prompt and language.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/get_custom_text_by_language
func (m *PromptManager) CustomText(p string, l string, opts ...RequestOption) (t map[string]interface{}, err error) {
	err = m.Request("GET", m.URI("prompts", p, "custom-text", l), &t, opts...)
	return
}

// SetCustomText sets the custom text for a specific prompt. Existing texts will be overwritten.
//
// See: https://auth0.com/docs/api/management/v2#!/Prompts/put_custom_text_by_language
func (m *PromptManager) SetCustomText(p string, l string, b map[string]interface{}, opts ...RequestOption) (err error) {
	err = m.Request("PUT", m.URI("prompts", p, "custom-text", l), &b, opts...)
	return
}
