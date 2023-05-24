package management

import "context"

// EmailTemplate is used to customize emails.
//
// See https://auth0.com/docs/customize/email/email-templates
type EmailTemplate struct {
	// The template name. Can be one of "verify_email", "reset_email",
	// "welcome_email", "blocked_account", "stolen_credentials",
	// "enrollment_email", "change_password", "password_reset" or
	// "mfa_oob_code".
	Template *string `json:"template,omitempty"`

	// The body of the template.
	Body *string `json:"body,omitempty"`

	// The sender of the email.
	From *string `json:"from,omitempty"`

	// The URL to redirect the user to after a successful action.
	ResultURL *string `json:"resultUrl,omitempty"`

	// The subject of the email.
	Subject *string `json:"subject,omitempty"`

	// The syntax of the template body.
	Syntax *string `json:"syntax,omitempty"`

	// The lifetime in seconds that the link within the email will be valid for.
	URLLifetimeInSecoonds *int `json:"urlLifetimeInSeconds,omitempty"`

	// Whether or not the template is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Whether the `reset_email` and `verify_email` templates should include the user's
	// email address as the email parameter in the returnUrl (true) or whether no email
	// address should be included in the redirect (false). Defaults to true.
	IncludeEmailInRedirect *bool `json:"includeEmailInRedirect,omitempty"`
}

// EmailTemplateManager manages Auth0 EmailTemplate resources.
type EmailTemplateManager manager

// Create an email template.
//
// See: https://auth0.com/docs/api/management/v2#!/Email_Templates/post_email_templates
func (m *EmailTemplateManager) Create(ctx context.Context, e *EmailTemplate, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("email-templates"), e, opts...)
}

// Read an email template by pre-defined name.
//
// These names are `verify_email`, `reset_email`, `welcome_email`,
// `blocked_account`, `stolen_credentials`, `enrollment_email`, and
// `mfa_oob_code`.
//
// The names `change_password`, and `password_reset` are also supported for
// legacy scenarios.
//
// See: https://auth0.com/docs/api/management/v2#!/Email_Templates/get_email_templates_by_templateName
func (m *EmailTemplateManager) Read(ctx context.Context, template string, opts ...RequestOption) (e *EmailTemplate, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("email-templates", template), &e, opts...)
	return
}

// Update an email template.
//
// See: https://auth0.com/docs/api/management/v2#!/Email_Templates/patch_email_templates_by_templateName
func (m *EmailTemplateManager) Update(ctx context.Context, template string, e *EmailTemplate, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("email-templates", template), e, opts...)
}

// Replace an email template.
//
// See: https://auth0.com/docs/api/management/v2#!/Email_Templates/put_email_templates_by_templateName
func (m *EmailTemplateManager) Replace(ctx context.Context, template string, e *EmailTemplate, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PUT", m.management.URI("email-templates", template), e, opts...)
}
