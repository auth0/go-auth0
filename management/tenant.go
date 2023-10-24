package management

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/auth0/go-auth0"
)

// Tenant represents an Auth0 Tenant.
type Tenant struct {
	// Change password page settings
	ChangePassword *TenantChangePassword `json:"change_password,omitempty"`

	// Guardian MFA page settings
	GuardianMFAPage *TenantGuardianMFAPage `json:"guardian_mfa_page,omitempty"`

	// Default audience for API Authorization
	DefaultAudience *string `json:"default_audience,omitempty"`

	// Name of the connection that will be used for password grants at the token
	// endpoint. Only the following connection types are supported: LDAP, AD,
	// Database Connections, Passwordless, Windows Azure Active Directory, ADFS.
	DefaultDirectory *string `json:"default_directory,omitempty"`

	ErrorPage *TenantErrorPage `json:"error_page,omitempty"`

	DeviceFlow *TenantDeviceFlow `json:"device_flow,omitempty"`

	Flags *TenantFlags `json:"flags,omitempty"`

	// The friendly name of the tenant
	FriendlyName *string `json:"friendly_name,omitempty"`

	// The URL of the tenant logo (recommended size: 150x150)
	PictureURL *string `json:"picture_url,omitempty"`

	// User support email
	SupportEmail *string `json:"support_email,omitempty"`

	// User support URL
	SupportURL *string `json:"support_url,omitempty"`

	// Used to store additional metadata
	UniversalLogin *TenantUniversalLogin `json:"universal_login,omitempty"`

	// A set of URLs that are valid to redirect to after logout from Auth0.
	AllowedLogoutURLs *[]string `json:"allowed_logout_urls,omitempty"`

	// Login session lifetime, how long the session will stay valid (hours).
	//
	// When marshalling, values are rounded to the nearest integer. If the value
	// is smaller than 1, hours are transformed to minutes and marshaled as
	// session_lifetime_in_minutes instead.
	SessionLifetime *float64 `json:"session_lifetime,omitempty"`

	// Force a user to login after they have been inactive for the specified
	// number (hours).
	//
	// When marshalling, values are rounded to the nearest integer. If the value
	// is smaller than 1, hours are transformed to minutes and marshaled as
	// idle_session_lifetime_in_minutes instead.
	IdleSessionLifetime *float64 `json:"idle_session_lifetime,omitempty"`

	// The selected sandbox version to be used for the extensibility environment
	SandboxVersion *string `json:"sandbox_version,omitempty"`

	// A set of available sandbox versions for the extensibility environment
	SandboxVersionAvailable *[]string `json:"sandbox_versions_available,omitempty"`

	// The default absolute redirection uri, must be https and cannot contain a
	// fragment.
	DefaultRedirectionURI *string `json:"default_redirection_uri,omitempty"`

	// Supported locales for the UI
	EnabledLocales *[]string `json:"enabled_locales,omitempty"`

	SessionCookie *TenantSessionCookie `json:"session_cookie,omitempty"`

	// Sessions related settings for the tenant.
	Sessions *TenantSessions `json:"sessions,omitempty"`

	// If `true`, allows accepting an organization name or organization ID on auth endpoints.
	AllowOrgNameInAuthAPI *bool `json:"allow_organization_name_in_authentication_api,omitempty"`

	// If `true`, flexible factors will be enabled for MFA in the PostLogin action.
	CustomizeMFAInPostLoginAction *bool `json:"customize_mfa_in_postlogin_action,omitempty"`
}

// MarshalJSON is a custom serializer for the Tenant type.
func (t *Tenant) MarshalJSON() ([]byte, error) {
	type tenant Tenant
	type tenantWrapper struct {
		*tenant
		SessionLifetimeInMinutes     *int `json:"session_lifetime_in_minutes,omitempty"`
		IdleSessionLifetimeInMinutes *int `json:"idle_session_lifetime_in_minutes,omitempty"`
	}

	w := &tenantWrapper{(*tenant)(t), nil, nil}

	if t.SessionLifetime != nil {
		sessionLifetime := t.GetSessionLifetime()

		if sessionLifetime < 1 {
			w.SessionLifetimeInMinutes = auth0.Int(int(math.Round(sessionLifetime * 60.0)))
			w.SessionLifetime = nil
			defer func() { w.SessionLifetime = &sessionLifetime }()
		} else {
			w.SessionLifetime = auth0.Float64(math.Round(sessionLifetime))
		}
	}

	if t.IdleSessionLifetime != nil {
		idleSessionLifetime := t.GetIdleSessionLifetime()

		if idleSessionLifetime < 1 {
			w.IdleSessionLifetimeInMinutes = auth0.Int(int(math.Round(idleSessionLifetime * 60.0)))
			w.IdleSessionLifetime = nil
			defer func() { w.IdleSessionLifetime = &idleSessionLifetime }()
		} else {
			w.IdleSessionLifetime = auth0.Float64(math.Round(idleSessionLifetime))
		}
	}

	return json.Marshal(w)
}

// TenantChangePassword holds settings for the change password page.
type TenantChangePassword struct {
	// True to use the custom change password html, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`

	// Replace default change password page with a custom HTML (Liquid syntax is
	// supported).
	HTML *string `json:"html,omitempty"`
}

// TenantGuardianMFAPage holds settings for guardian mfa page.
type TenantGuardianMFAPage struct {
	// True to use the custom html for Guardian page, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
	// Replace default Guardian page with a custom HTML (Liquid syntax is
	// supported).
	HTML *string `json:"html,omitempty"`
}

// TenantErrorPage holds settings for the error page.
type TenantErrorPage struct {
	// Replace default error page with a custom HTML (Liquid syntax is
	// supported).
	HTML *string `json:"html,omitempty"`
	// True to show link to log as part of the default error page, false
	// otherwise (default: true).
	ShowLogLink *bool `json:"show_log_link,omitempty"`
	// Redirect to specified url instead of show the default error page
	URL *string `json:"url,omitempty"`
}

// TenantFlags holds information on flag toggles.
type TenantFlags struct {
	// This flag determines whether all current connections shall be enabled
	// when a new client is created. Default value is true.
	EnableClientConnections *bool `json:"enable_client_connections,omitempty"`

	// This flag enables the API section in the Auth0 Management Dashboard.
	EnableAPIsSection *bool `json:"enable_apis_section,omitempty"`

	// If set to true all Impersonation functionality is disabled for the
	// Tenant. This is a read-only attribute.
	DisableImpersonation *bool `json:"disable_impersonation,omitempty"`

	// This flag enables advanced API Authorization scenarios.
	EnablePipeline2 *bool `json:"enable_pipeline2,omitempty"`

	// This flag enables dynamic client registration.
	EnableDynamicClientRegistration *bool `json:"enable_dynamic_client_registration,omitempty"`

	// If enabled, All your email links and urls will use your configured custom
	// domain. If no custom domain is found the email operation will fail.
	EnableCustomDomainInEmails *bool `json:"enable_custom_domain_in_emails,omitempty"`

	// If enabled, users will not be prompted to confirm log in before SSO
	// redirection.
	EnableSSO *bool `json:"enable_sso,omitempty"`

	// Whether the `EnableSSO` setting can be changed.
	AllowChangingEnableSSO *bool `json:"allow_changing_enable_sso,omitempty"`

	// If enabled, activate the new look and feel for Universal Login
	UniversalLogin *bool `json:"universal_login,omitempty"`

	// If enabled, the legacy Logs Search Engine V2 will be enabled for your
	// account.
	//
	// Turn it off to opt-in for the latest Logs Search Engine V3.
	EnableLegacyLogsSearchV2 *bool `json:"enable_legacy_logs_search_v2,omitempty"`

	// If enabled, additional HTTP security headers will not be included in the
	// response to prevent embedding of the Universal Login prompts in an
	// IFRAME.
	DisableClickjackProtectionHeaders *bool `json:"disable_clickjack_protection_headers,omitempty"`

	// If enabled, this will use a generic response in the public signup API
	// which will prevent users from being able to find out if an e-mail address
	// or username has previously registered.
	EnablePublicSignupUserExistsError *bool `json:"enable_public_signup_user_exists_error,omitempty"`

	// If enabled, this will use the scope description when generating a consent
	// prompt. Otherwise, the scope name is used.
	UseScopeDescriptionsForConsent *bool `json:"use_scope_descriptions_for_consent,omitempty"`

	// Whether the legacy delegation endpoint will be enabled for your account (true) or not available (false).
	AllowLegacyDelegationGrantTypes *bool `json:"allow_legacy_delegation_grant_types,omitempty"`

	// Whether the legacy `auth/ro` endpoint (used with resource owner password and passwordless features) will be enabled for your account (true) or not available (false).
	AllowLegacyROGrantTypes *bool `json:"allow_legacy_ro_grant_types,omitempty"`

	// If enabled, customers can use Tokeninfo Endpoint, otherwise they can not use it.
	AllowLegacyTokenInfoEndpoint *bool `json:"allow_legacy_tokeninfo_endpoint,omitempty"`

	// Whether ID tokens and the userinfo endpoint includes a complete user profile (true) or only OpenID Connect claims (false).
	EnableLegacyProfile *bool `json:"enable_legacy_profile,omitempty"`

	// Whether ID tokens can be used to authorize some types of requests to API v2 (true) not not (false).
	EnableIDTokenAPI2 *bool `json:"enable_idtoken_api2,omitempty"`

	// Do not Publish Enterprise Connections Information with IdP domains on the lock configuration file.
	NoDisclosureEnterpriseConnections *bool `json:"no_disclose_enterprise_connections,omitempty"`

	// If true, SMS phone numbers will not be obfuscated in Management API GET calls.
	DisableManagementAPISMSObfuscation *bool `json:"disable_management_api_sms_obfuscation,omitempty"`

	// If enabled, users will be presented with an email verification prompt during their first login when using Azure AD or ADFS connections
	EnableADFSWAADEmailVerification *bool `json:"enable_adfs_waad_email_verification,omitempty"`

	// Delete underlying grant when a Refresh Token is revoked via the Authentication API.
	RevokeRefreshTokenGrant *bool `json:"revoke_refresh_token_grant,omitempty"`

	// Enables beta access to log streaming changes
	DashboardLogStreams *bool `json:"dashboard_log_streams_next,omitempty"`

	// Enables new insights activity page view
	DashboardInsightsView *bool `json:"dashboard_insights_view,omitempty"`

	// Disables SAML fields map fix for bad mappings with repeated attributes
	DisableFieldsMapFix *bool `json:"disable_fields_map_fix,omitempty"`

	// Used to allow users to pick what factor to enroll of the available MFA factors.
	MFAShowFactorListOnEnrollment *bool `json:"mfa_show_factor_list_on_enrollment,omitempty"`

	// If `true`, all Clients will be required to use Pushed Authorization Requests.
	// This feature currently must be enabled for your tenant.
	RequirePushedAuthorizationRequests *bool `json:"require_pushed_authorization_requests,omitempty"`
}

// TenantUniversalLogin holds universal login settings.
type TenantUniversalLogin struct {
	Colors *TenantUniversalLoginColors `json:"colors,omitempty"`
}

// TenantUniversalLoginColors holds universal login color settings.
type TenantUniversalLoginColors struct {
	// Primary button background color
	Primary *string `json:"primary,omitempty"`

	// Page background color.
	//
	// Only one of PageBackground and PageBackgroundGradient should be set. If
	// both fields are set, PageBackground takes priority.
	PageBackground *string `json:"-"`

	// Page background gradient.
	//
	// Only one of PageBackground and PageBackgroundGradient should be set. If
	// both fields are set, PageBackground takes priority.
	PageBackgroundGradient *BrandingPageBackgroundGradient `json:"-"`
}

// MarshalJSON is a custom serializer for the TenantUniversalLoginColors type.
func (c *TenantUniversalLoginColors) MarshalJSON() ([]byte, error) {
	type colors TenantUniversalLoginColors
	type colorsWrapper struct {
		*colors
		RawPageBackground interface{} `json:"page_background,omitempty"`
	}

	alias := &colorsWrapper{(*colors)(c), nil}

	switch {
	case c.PageBackground != nil && c.PageBackgroundGradient != nil:
		return nil, fmt.Errorf("only one of PageBackground and PageBackgroundGradient is allowed")
	case c.PageBackground != nil:
		alias.RawPageBackground = c.PageBackground
	case c.PageBackgroundGradient != nil:
		alias.RawPageBackground = c.PageBackgroundGradient
	}

	return json.Marshal(alias)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the json field page_background, which can either
// be a hex color string, or an object describing a gradient.
func (c *TenantUniversalLoginColors) UnmarshalJSON(data []byte) error {
	type colors BrandingColors
	type colorsWrapper struct {
		*colors
		RawPageBackground json.RawMessage `json:"page_background,omitempty"`
	}

	alias := &colorsWrapper{(*colors)(c), nil}

	err := json.Unmarshal(data, alias)
	if err != nil {
		return err
	}

	if alias.RawPageBackground != nil {
		var v interface{}
		err = json.Unmarshal(alias.RawPageBackground, &v)
		if err != nil {
			return err
		}

		switch rawPageBackground := v.(type) {
		case string:
			c.PageBackground = &rawPageBackground

		case map[string]interface{}:
			var gradient BrandingPageBackgroundGradient
			err = json.Unmarshal(alias.RawPageBackground, &gradient)
			if err != nil {
				return err
			}
			c.PageBackgroundGradient = &gradient

		default:
			return fmt.Errorf("unexpected type for field page_background")
		}
	}

	return nil
}

// TenantDeviceFlow holds settings for the device flow.
type TenantDeviceFlow struct {
	// The character set for generating a User Code ['base20' or 'digits']
	Charset *string `json:"charset,omitempty"`

	// The mask used to format the generated User Code to a friendly, readable
	// format with possible spaces or hyphens
	Mask *string `json:"mask,omitempty"`
}

// TenantSessionCookie manages behavior of the tenant's session cookie, accepts either 'persistent' or 'non-persistent'.
type TenantSessionCookie struct {
	Mode *string `json:"mode,omitempty"`
}

// TenantSessions manages sessions related settings for the tenant.
type TenantSessions struct {
	// Whether to bypass prompting logic (false) when performing OIDC Logout.
	OIDCLogoutPromptEnabled *bool `json:"oidc_logout_prompt_enabled,omitempty"`
}

// TenantManager manages Auth0 Tenant resources.
type TenantManager manager

// Read tenant settings.
// A list of fields to include or exclude may also be specified.
//
// See: https://auth0.com/docs/api/management/v2#!/Tenants/get_settings
func (m *TenantManager) Read(ctx context.Context, opts ...RequestOption) (t *Tenant, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("tenants", "settings"), &t, opts...)
	return
}

// Update settings for a tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Tenants/patch_settings
func (m *TenantManager) Update(ctx context.Context, t *Tenant, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("tenants", "settings"), t, opts...)
}
