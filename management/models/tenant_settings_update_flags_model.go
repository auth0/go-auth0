/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// TenantSettingsUpdateFlags Flags used to change the behavior of this tenant.
type TenantSettingsUpdateFlags struct {
	// Whether to use the older v1 change password flow (true, not recommended except for backward compatibility) or the newer safer flow (false, recommended).
	ChangePwdFlowV1 *bool `json:"change_pwd_flow_v1,omitempty"`
	// Whether all current connections should be enabled when a new client (application) is created (true, default) or not (false).
	EnableClientConnections *bool `json:"enable_client_connections,omitempty"`
	// Whether the APIs section is enabled (true) or disabled (false).
	EnableApisSection *bool `json:"enable_apis_section,omitempty"`
	// Whether advanced API Authorization scenarios are enabled (true) or disabled (false).
	EnablePipeline2 *bool `json:"enable_pipeline2,omitempty"`
	//  Whether third-party developers can <a href='https://auth0.com/docs/api-auth/dynamic-client-registration'>dynamically register</a> applications for your APIs (true) or not (false). This flag enables dynamic client registration.
	EnableDynamicClientRegistration *bool `json:"enable_dynamic_client_registration,omitempty"`
	// Whether emails sent by Auth0 for change password, verification etc. should use your verified custom domain (true) or your auth0.com sub-domain (false).  Affects all emails, links, and URLs. Email will fail if the custom domain is not verified.
	EnableCustomDomainInEmails *bool `json:"enable_custom_domain_in_emails,omitempty"`
	// Whether the legacy `/tokeninfo` endpoint is enabled for your account (true) or unavailable (false).
	AllowLegacyTokeninfoEndpoint *bool `json:"allow_legacy_tokeninfo_endpoint,omitempty"`
	// Whether ID tokens and the userinfo endpoint includes a complete user profile (true) or only OpenID Connect claims (false).
	EnableLegacyProfile *bool `json:"enable_legacy_profile,omitempty"`
	// Whether ID tokens can be used to authorize some types of requests to API v2 (true) not not (false).
	EnableIdtokenApi2 *bool `json:"enable_idtoken_api2,omitempty"`
	// Whether the public sign up process shows a user_exists error (true) or a generic error (false) if the user already exists.
	EnablePublicSignupUserExistsError *bool `json:"enable_public_signup_user_exists_error,omitempty"`
	//  Whether the legacy delegation endpoint will be enabled for your account (true) or not available (false).
	AllowLegacyDelegationGrantTypes *bool `json:"allow_legacy_delegation_grant_types,omitempty"`
	// Whether the legacy `auth/ro` endpoint (used with resource owner password and passwordless features) will be enabled for your account (true) or not available (false).
	AllowLegacyRoGrantTypes *bool `json:"allow_legacy_ro_grant_types,omitempty"`
	// Whether users are prompted to confirm log in before SSO redirection (false) or are not prompted (true).
	EnableSso *bool `json:"enable_sso,omitempty"`
	// Whether classic Universal Login prompts include additional security headers to prevent clickjacking (true) or no safeguard (false).
	DisableClickjackProtectionHeaders *bool `json:"disable_clickjack_protection_headers,omitempty"`
	// Do not Publish Enterprise Connections Information with IdP domains on the lock configuration file.
	NoDiscloseEnterpriseConnections *bool `json:"no_disclose_enterprise_connections,omitempty"`
	// If true, SMS phone numbers will not be obfuscated in Management API GET calls.
	DisableManagementApiSmsObfuscation *bool `json:"disable_management_api_sms_obfuscation,omitempty"`
	// Enforce client authentication for passwordless start.
	EnforceClientAuthenticationOnPasswordlessStart *bool `json:"enforce_client_authentication_on_passwordless_start,omitempty"`
	// Changes email_verified behavior for Azure AD/ADFS connections when enabled. Sets email_verified to false otherwise.
	TrustAzureAdfsEmailVerifiedConnectionProperty *bool `json:"trust_azure_adfs_email_verified_connection_property,omitempty"`
	// Enables the email verification flow during login for Azure AD and ADFS connections.
	EnableAdfsWaadEmailVerification *bool `json:"enable_adfs_waad_email_verification,omitempty"`
	// Delete underlying grant when a Refresh Token is revoked via the Authentication API.
	RevokeRefreshTokenGrant *bool `json:"revoke_refresh_token_grant,omitempty"`
	// Enables beta access to log streaming changes.
	DashboardLogStreamsNext *bool `json:"dashboard_log_streams_next,omitempty"`
	// Enables new insights activity page view.
	DashboardInsightsView *bool `json:"dashboard_insights_view,omitempty"`
	// Disables SAML fields map fix for bad mappings with repeated attributes.
	DisableFieldsMapFix *bool `json:"disable_fields_map_fix,omitempty"`
	// Used to allow users to pick what factor to enroll of the available MFA factors.
	MfaShowFactorListOnEnrollment *bool `json:"mfa_show_factor_list_on_enrollment,omitempty"`
	// Removes alg property from jwks .well-known endpoint
	RemoveAlgFromJwks *bool `json:"remove_alg_from_jwks,omitempty"`
}

// GetChangePwdFlowV1 returns the ChangePwdFlowV1 field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetChangePwdFlowV1() bool {
	if o == nil || IsNil(o.ChangePwdFlowV1) {
		var ret bool
		return ret
	}
	return *o.ChangePwdFlowV1
}

// GetChangePwdFlowV1Ok returns a tuple with the ChangePwdFlowV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetChangePwdFlowV1Ok() (*bool, bool) {
	if o == nil || IsNil(o.ChangePwdFlowV1) {
		return nil, false
	}
	return o.ChangePwdFlowV1, true
}

// HasChangePwdFlowV1 returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasChangePwdFlowV1() bool {
	if o != nil && !IsNil(o.ChangePwdFlowV1) {
		return true
	}

	return false
}

// SetChangePwdFlowV1 gets a reference to the given bool and assigns it to the ChangePwdFlowV1 field.
func (o *TenantSettingsUpdateFlags) SetChangePwdFlowV1(v bool) {
	o.ChangePwdFlowV1 = &v
}

// GetEnableClientConnections returns the EnableClientConnections field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableClientConnections() bool {
	if o == nil || IsNil(o.EnableClientConnections) {
		var ret bool
		return ret
	}
	return *o.EnableClientConnections
}

// GetEnableClientConnectionsOk returns a tuple with the EnableClientConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableClientConnectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClientConnections) {
		return nil, false
	}
	return o.EnableClientConnections, true
}

// HasEnableClientConnections returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableClientConnections() bool {
	if o != nil && !IsNil(o.EnableClientConnections) {
		return true
	}

	return false
}

// SetEnableClientConnections gets a reference to the given bool and assigns it to the EnableClientConnections field.
func (o *TenantSettingsUpdateFlags) SetEnableClientConnections(v bool) {
	o.EnableClientConnections = &v
}

// GetEnableApisSection returns the EnableApisSection field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableApisSection() bool {
	if o == nil || IsNil(o.EnableApisSection) {
		var ret bool
		return ret
	}
	return *o.EnableApisSection
}

// GetEnableApisSectionOk returns a tuple with the EnableApisSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableApisSectionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableApisSection) {
		return nil, false
	}
	return o.EnableApisSection, true
}

// HasEnableApisSection returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableApisSection() bool {
	if o != nil && !IsNil(o.EnableApisSection) {
		return true
	}

	return false
}

// SetEnableApisSection gets a reference to the given bool and assigns it to the EnableApisSection field.
func (o *TenantSettingsUpdateFlags) SetEnableApisSection(v bool) {
	o.EnableApisSection = &v
}

// GetEnablePipeline2 returns the EnablePipeline2 field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnablePipeline2() bool {
	if o == nil || IsNil(o.EnablePipeline2) {
		var ret bool
		return ret
	}
	return *o.EnablePipeline2
}

// GetEnablePipeline2Ok returns a tuple with the EnablePipeline2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnablePipeline2Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnablePipeline2) {
		return nil, false
	}
	return o.EnablePipeline2, true
}

// HasEnablePipeline2 returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnablePipeline2() bool {
	if o != nil && !IsNil(o.EnablePipeline2) {
		return true
	}

	return false
}

// SetEnablePipeline2 gets a reference to the given bool and assigns it to the EnablePipeline2 field.
func (o *TenantSettingsUpdateFlags) SetEnablePipeline2(v bool) {
	o.EnablePipeline2 = &v
}

// GetEnableDynamicClientRegistration returns the EnableDynamicClientRegistration field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableDynamicClientRegistration() bool {
	if o == nil || IsNil(o.EnableDynamicClientRegistration) {
		var ret bool
		return ret
	}
	return *o.EnableDynamicClientRegistration
}

// GetEnableDynamicClientRegistrationOk returns a tuple with the EnableDynamicClientRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableDynamicClientRegistrationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDynamicClientRegistration) {
		return nil, false
	}
	return o.EnableDynamicClientRegistration, true
}

// HasEnableDynamicClientRegistration returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableDynamicClientRegistration() bool {
	if o != nil && !IsNil(o.EnableDynamicClientRegistration) {
		return true
	}

	return false
}

// SetEnableDynamicClientRegistration gets a reference to the given bool and assigns it to the EnableDynamicClientRegistration field.
func (o *TenantSettingsUpdateFlags) SetEnableDynamicClientRegistration(v bool) {
	o.EnableDynamicClientRegistration = &v
}

// GetEnableCustomDomainInEmails returns the EnableCustomDomainInEmails field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableCustomDomainInEmails() bool {
	if o == nil || IsNil(o.EnableCustomDomainInEmails) {
		var ret bool
		return ret
	}
	return *o.EnableCustomDomainInEmails
}

// GetEnableCustomDomainInEmailsOk returns a tuple with the EnableCustomDomainInEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableCustomDomainInEmailsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCustomDomainInEmails) {
		return nil, false
	}
	return o.EnableCustomDomainInEmails, true
}

// HasEnableCustomDomainInEmails returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableCustomDomainInEmails() bool {
	if o != nil && !IsNil(o.EnableCustomDomainInEmails) {
		return true
	}

	return false
}

// SetEnableCustomDomainInEmails gets a reference to the given bool and assigns it to the EnableCustomDomainInEmails field.
func (o *TenantSettingsUpdateFlags) SetEnableCustomDomainInEmails(v bool) {
	o.EnableCustomDomainInEmails = &v
}

// GetAllowLegacyTokeninfoEndpoint returns the AllowLegacyTokeninfoEndpoint field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetAllowLegacyTokeninfoEndpoint() bool {
	if o == nil || IsNil(o.AllowLegacyTokeninfoEndpoint) {
		var ret bool
		return ret
	}
	return *o.AllowLegacyTokeninfoEndpoint
}

// GetAllowLegacyTokeninfoEndpointOk returns a tuple with the AllowLegacyTokeninfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetAllowLegacyTokeninfoEndpointOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLegacyTokeninfoEndpoint) {
		return nil, false
	}
	return o.AllowLegacyTokeninfoEndpoint, true
}

// HasAllowLegacyTokeninfoEndpoint returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasAllowLegacyTokeninfoEndpoint() bool {
	if o != nil && !IsNil(o.AllowLegacyTokeninfoEndpoint) {
		return true
	}

	return false
}

// SetAllowLegacyTokeninfoEndpoint gets a reference to the given bool and assigns it to the AllowLegacyTokeninfoEndpoint field.
func (o *TenantSettingsUpdateFlags) SetAllowLegacyTokeninfoEndpoint(v bool) {
	o.AllowLegacyTokeninfoEndpoint = &v
}

// GetEnableLegacyProfile returns the EnableLegacyProfile field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableLegacyProfile() bool {
	if o == nil || IsNil(o.EnableLegacyProfile) {
		var ret bool
		return ret
	}
	return *o.EnableLegacyProfile
}

// GetEnableLegacyProfileOk returns a tuple with the EnableLegacyProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableLegacyProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLegacyProfile) {
		return nil, false
	}
	return o.EnableLegacyProfile, true
}

// HasEnableLegacyProfile returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableLegacyProfile() bool {
	if o != nil && !IsNil(o.EnableLegacyProfile) {
		return true
	}

	return false
}

// SetEnableLegacyProfile gets a reference to the given bool and assigns it to the EnableLegacyProfile field.
func (o *TenantSettingsUpdateFlags) SetEnableLegacyProfile(v bool) {
	o.EnableLegacyProfile = &v
}

// GetEnableIdtokenApi2 returns the EnableIdtokenApi2 field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableIdtokenApi2() bool {
	if o == nil || IsNil(o.EnableIdtokenApi2) {
		var ret bool
		return ret
	}
	return *o.EnableIdtokenApi2
}

// GetEnableIdtokenApi2Ok returns a tuple with the EnableIdtokenApi2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableIdtokenApi2Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableIdtokenApi2) {
		return nil, false
	}
	return o.EnableIdtokenApi2, true
}

// HasEnableIdtokenApi2 returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableIdtokenApi2() bool {
	if o != nil && !IsNil(o.EnableIdtokenApi2) {
		return true
	}

	return false
}

// SetEnableIdtokenApi2 gets a reference to the given bool and assigns it to the EnableIdtokenApi2 field.
func (o *TenantSettingsUpdateFlags) SetEnableIdtokenApi2(v bool) {
	o.EnableIdtokenApi2 = &v
}

// GetEnablePublicSignupUserExistsError returns the EnablePublicSignupUserExistsError field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnablePublicSignupUserExistsError() bool {
	if o == nil || IsNil(o.EnablePublicSignupUserExistsError) {
		var ret bool
		return ret
	}
	return *o.EnablePublicSignupUserExistsError
}

// GetEnablePublicSignupUserExistsErrorOk returns a tuple with the EnablePublicSignupUserExistsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnablePublicSignupUserExistsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePublicSignupUserExistsError) {
		return nil, false
	}
	return o.EnablePublicSignupUserExistsError, true
}

// HasEnablePublicSignupUserExistsError returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnablePublicSignupUserExistsError() bool {
	if o != nil && !IsNil(o.EnablePublicSignupUserExistsError) {
		return true
	}

	return false
}

// SetEnablePublicSignupUserExistsError gets a reference to the given bool and assigns it to the EnablePublicSignupUserExistsError field.
func (o *TenantSettingsUpdateFlags) SetEnablePublicSignupUserExistsError(v bool) {
	o.EnablePublicSignupUserExistsError = &v
}

// GetAllowLegacyDelegationGrantTypes returns the AllowLegacyDelegationGrantTypes field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetAllowLegacyDelegationGrantTypes() bool {
	if o == nil || IsNil(o.AllowLegacyDelegationGrantTypes) {
		var ret bool
		return ret
	}
	return *o.AllowLegacyDelegationGrantTypes
}

// GetAllowLegacyDelegationGrantTypesOk returns a tuple with the AllowLegacyDelegationGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetAllowLegacyDelegationGrantTypesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLegacyDelegationGrantTypes) {
		return nil, false
	}
	return o.AllowLegacyDelegationGrantTypes, true
}

// HasAllowLegacyDelegationGrantTypes returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasAllowLegacyDelegationGrantTypes() bool {
	if o != nil && !IsNil(o.AllowLegacyDelegationGrantTypes) {
		return true
	}

	return false
}

// SetAllowLegacyDelegationGrantTypes gets a reference to the given bool and assigns it to the AllowLegacyDelegationGrantTypes field.
func (o *TenantSettingsUpdateFlags) SetAllowLegacyDelegationGrantTypes(v bool) {
	o.AllowLegacyDelegationGrantTypes = &v
}

// GetAllowLegacyRoGrantTypes returns the AllowLegacyRoGrantTypes field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetAllowLegacyRoGrantTypes() bool {
	if o == nil || IsNil(o.AllowLegacyRoGrantTypes) {
		var ret bool
		return ret
	}
	return *o.AllowLegacyRoGrantTypes
}

// GetAllowLegacyRoGrantTypesOk returns a tuple with the AllowLegacyRoGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetAllowLegacyRoGrantTypesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLegacyRoGrantTypes) {
		return nil, false
	}
	return o.AllowLegacyRoGrantTypes, true
}

// HasAllowLegacyRoGrantTypes returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasAllowLegacyRoGrantTypes() bool {
	if o != nil && !IsNil(o.AllowLegacyRoGrantTypes) {
		return true
	}

	return false
}

// SetAllowLegacyRoGrantTypes gets a reference to the given bool and assigns it to the AllowLegacyRoGrantTypes field.
func (o *TenantSettingsUpdateFlags) SetAllowLegacyRoGrantTypes(v bool) {
	o.AllowLegacyRoGrantTypes = &v
}

// GetEnableSso returns the EnableSso field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableSso() bool {
	if o == nil || IsNil(o.EnableSso) {
		var ret bool
		return ret
	}
	return *o.EnableSso
}

// GetEnableSsoOk returns a tuple with the EnableSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSso) {
		return nil, false
	}
	return o.EnableSso, true
}

// HasEnableSso returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableSso() bool {
	if o != nil && !IsNil(o.EnableSso) {
		return true
	}

	return false
}

// SetEnableSso gets a reference to the given bool and assigns it to the EnableSso field.
func (o *TenantSettingsUpdateFlags) SetEnableSso(v bool) {
	o.EnableSso = &v
}

// GetDisableClickjackProtectionHeaders returns the DisableClickjackProtectionHeaders field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetDisableClickjackProtectionHeaders() bool {
	if o == nil || IsNil(o.DisableClickjackProtectionHeaders) {
		var ret bool
		return ret
	}
	return *o.DisableClickjackProtectionHeaders
}

// GetDisableClickjackProtectionHeadersOk returns a tuple with the DisableClickjackProtectionHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetDisableClickjackProtectionHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableClickjackProtectionHeaders) {
		return nil, false
	}
	return o.DisableClickjackProtectionHeaders, true
}

// HasDisableClickjackProtectionHeaders returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasDisableClickjackProtectionHeaders() bool {
	if o != nil && !IsNil(o.DisableClickjackProtectionHeaders) {
		return true
	}

	return false
}

// SetDisableClickjackProtectionHeaders gets a reference to the given bool and assigns it to the DisableClickjackProtectionHeaders field.
func (o *TenantSettingsUpdateFlags) SetDisableClickjackProtectionHeaders(v bool) {
	o.DisableClickjackProtectionHeaders = &v
}

// GetNoDiscloseEnterpriseConnections returns the NoDiscloseEnterpriseConnections field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetNoDiscloseEnterpriseConnections() bool {
	if o == nil || IsNil(o.NoDiscloseEnterpriseConnections) {
		var ret bool
		return ret
	}
	return *o.NoDiscloseEnterpriseConnections
}

// GetNoDiscloseEnterpriseConnectionsOk returns a tuple with the NoDiscloseEnterpriseConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetNoDiscloseEnterpriseConnectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.NoDiscloseEnterpriseConnections) {
		return nil, false
	}
	return o.NoDiscloseEnterpriseConnections, true
}

// HasNoDiscloseEnterpriseConnections returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasNoDiscloseEnterpriseConnections() bool {
	if o != nil && !IsNil(o.NoDiscloseEnterpriseConnections) {
		return true
	}

	return false
}

// SetNoDiscloseEnterpriseConnections gets a reference to the given bool and assigns it to the NoDiscloseEnterpriseConnections field.
func (o *TenantSettingsUpdateFlags) SetNoDiscloseEnterpriseConnections(v bool) {
	o.NoDiscloseEnterpriseConnections = &v
}

// GetDisableManagementApiSmsObfuscation returns the DisableManagementApiSmsObfuscation field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetDisableManagementApiSmsObfuscation() bool {
	if o == nil || IsNil(o.DisableManagementApiSmsObfuscation) {
		var ret bool
		return ret
	}
	return *o.DisableManagementApiSmsObfuscation
}

// GetDisableManagementApiSmsObfuscationOk returns a tuple with the DisableManagementApiSmsObfuscation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetDisableManagementApiSmsObfuscationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableManagementApiSmsObfuscation) {
		return nil, false
	}
	return o.DisableManagementApiSmsObfuscation, true
}

// HasDisableManagementApiSmsObfuscation returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasDisableManagementApiSmsObfuscation() bool {
	if o != nil && !IsNil(o.DisableManagementApiSmsObfuscation) {
		return true
	}

	return false
}

// SetDisableManagementApiSmsObfuscation gets a reference to the given bool and assigns it to the DisableManagementApiSmsObfuscation field.
func (o *TenantSettingsUpdateFlags) SetDisableManagementApiSmsObfuscation(v bool) {
	o.DisableManagementApiSmsObfuscation = &v
}

// GetEnforceClientAuthenticationOnPasswordlessStart returns the EnforceClientAuthenticationOnPasswordlessStart field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnforceClientAuthenticationOnPasswordlessStart() bool {
	if o == nil || IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		var ret bool
		return ret
	}
	return *o.EnforceClientAuthenticationOnPasswordlessStart
}

// GetEnforceClientAuthenticationOnPasswordlessStartOk returns a tuple with the EnforceClientAuthenticationOnPasswordlessStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnforceClientAuthenticationOnPasswordlessStartOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		return nil, false
	}
	return o.EnforceClientAuthenticationOnPasswordlessStart, true
}

// HasEnforceClientAuthenticationOnPasswordlessStart returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnforceClientAuthenticationOnPasswordlessStart() bool {
	if o != nil && !IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		return true
	}

	return false
}

// SetEnforceClientAuthenticationOnPasswordlessStart gets a reference to the given bool and assigns it to the EnforceClientAuthenticationOnPasswordlessStart field.
func (o *TenantSettingsUpdateFlags) SetEnforceClientAuthenticationOnPasswordlessStart(v bool) {
	o.EnforceClientAuthenticationOnPasswordlessStart = &v
}

// GetTrustAzureAdfsEmailVerifiedConnectionProperty returns the TrustAzureAdfsEmailVerifiedConnectionProperty field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetTrustAzureAdfsEmailVerifiedConnectionProperty() bool {
	if o == nil || IsNil(o.TrustAzureAdfsEmailVerifiedConnectionProperty) {
		var ret bool
		return ret
	}
	return *o.TrustAzureAdfsEmailVerifiedConnectionProperty
}

// GetTrustAzureAdfsEmailVerifiedConnectionPropertyOk returns a tuple with the TrustAzureAdfsEmailVerifiedConnectionProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetTrustAzureAdfsEmailVerifiedConnectionPropertyOk() (*bool, bool) {
	if o == nil || IsNil(o.TrustAzureAdfsEmailVerifiedConnectionProperty) {
		return nil, false
	}
	return o.TrustAzureAdfsEmailVerifiedConnectionProperty, true
}

// HasTrustAzureAdfsEmailVerifiedConnectionProperty returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasTrustAzureAdfsEmailVerifiedConnectionProperty() bool {
	if o != nil && !IsNil(o.TrustAzureAdfsEmailVerifiedConnectionProperty) {
		return true
	}

	return false
}

// SetTrustAzureAdfsEmailVerifiedConnectionProperty gets a reference to the given bool and assigns it to the TrustAzureAdfsEmailVerifiedConnectionProperty field.
func (o *TenantSettingsUpdateFlags) SetTrustAzureAdfsEmailVerifiedConnectionProperty(v bool) {
	o.TrustAzureAdfsEmailVerifiedConnectionProperty = &v
}

// GetEnableAdfsWaadEmailVerification returns the EnableAdfsWaadEmailVerification field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetEnableAdfsWaadEmailVerification() bool {
	if o == nil || IsNil(o.EnableAdfsWaadEmailVerification) {
		var ret bool
		return ret
	}
	return *o.EnableAdfsWaadEmailVerification
}

// GetEnableAdfsWaadEmailVerificationOk returns a tuple with the EnableAdfsWaadEmailVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetEnableAdfsWaadEmailVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAdfsWaadEmailVerification) {
		return nil, false
	}
	return o.EnableAdfsWaadEmailVerification, true
}

// HasEnableAdfsWaadEmailVerification returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasEnableAdfsWaadEmailVerification() bool {
	if o != nil && !IsNil(o.EnableAdfsWaadEmailVerification) {
		return true
	}

	return false
}

// SetEnableAdfsWaadEmailVerification gets a reference to the given bool and assigns it to the EnableAdfsWaadEmailVerification field.
func (o *TenantSettingsUpdateFlags) SetEnableAdfsWaadEmailVerification(v bool) {
	o.EnableAdfsWaadEmailVerification = &v
}

// GetRevokeRefreshTokenGrant returns the RevokeRefreshTokenGrant field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetRevokeRefreshTokenGrant() bool {
	if o == nil || IsNil(o.RevokeRefreshTokenGrant) {
		var ret bool
		return ret
	}
	return *o.RevokeRefreshTokenGrant
}

// GetRevokeRefreshTokenGrantOk returns a tuple with the RevokeRefreshTokenGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetRevokeRefreshTokenGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.RevokeRefreshTokenGrant) {
		return nil, false
	}
	return o.RevokeRefreshTokenGrant, true
}

// HasRevokeRefreshTokenGrant returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasRevokeRefreshTokenGrant() bool {
	if o != nil && !IsNil(o.RevokeRefreshTokenGrant) {
		return true
	}

	return false
}

// SetRevokeRefreshTokenGrant gets a reference to the given bool and assigns it to the RevokeRefreshTokenGrant field.
func (o *TenantSettingsUpdateFlags) SetRevokeRefreshTokenGrant(v bool) {
	o.RevokeRefreshTokenGrant = &v
}

// GetDashboardLogStreamsNext returns the DashboardLogStreamsNext field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetDashboardLogStreamsNext() bool {
	if o == nil || IsNil(o.DashboardLogStreamsNext) {
		var ret bool
		return ret
	}
	return *o.DashboardLogStreamsNext
}

// GetDashboardLogStreamsNextOk returns a tuple with the DashboardLogStreamsNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetDashboardLogStreamsNextOk() (*bool, bool) {
	if o == nil || IsNil(o.DashboardLogStreamsNext) {
		return nil, false
	}
	return o.DashboardLogStreamsNext, true
}

// HasDashboardLogStreamsNext returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasDashboardLogStreamsNext() bool {
	if o != nil && !IsNil(o.DashboardLogStreamsNext) {
		return true
	}

	return false
}

// SetDashboardLogStreamsNext gets a reference to the given bool and assigns it to the DashboardLogStreamsNext field.
func (o *TenantSettingsUpdateFlags) SetDashboardLogStreamsNext(v bool) {
	o.DashboardLogStreamsNext = &v
}

// GetDashboardInsightsView returns the DashboardInsightsView field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetDashboardInsightsView() bool {
	if o == nil || IsNil(o.DashboardInsightsView) {
		var ret bool
		return ret
	}
	return *o.DashboardInsightsView
}

// GetDashboardInsightsViewOk returns a tuple with the DashboardInsightsView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetDashboardInsightsViewOk() (*bool, bool) {
	if o == nil || IsNil(o.DashboardInsightsView) {
		return nil, false
	}
	return o.DashboardInsightsView, true
}

// HasDashboardInsightsView returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasDashboardInsightsView() bool {
	if o != nil && !IsNil(o.DashboardInsightsView) {
		return true
	}

	return false
}

// SetDashboardInsightsView gets a reference to the given bool and assigns it to the DashboardInsightsView field.
func (o *TenantSettingsUpdateFlags) SetDashboardInsightsView(v bool) {
	o.DashboardInsightsView = &v
}

// GetDisableFieldsMapFix returns the DisableFieldsMapFix field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetDisableFieldsMapFix() bool {
	if o == nil || IsNil(o.DisableFieldsMapFix) {
		var ret bool
		return ret
	}
	return *o.DisableFieldsMapFix
}

// GetDisableFieldsMapFixOk returns a tuple with the DisableFieldsMapFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetDisableFieldsMapFixOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableFieldsMapFix) {
		return nil, false
	}
	return o.DisableFieldsMapFix, true
}

// HasDisableFieldsMapFix returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasDisableFieldsMapFix() bool {
	if o != nil && !IsNil(o.DisableFieldsMapFix) {
		return true
	}

	return false
}

// SetDisableFieldsMapFix gets a reference to the given bool and assigns it to the DisableFieldsMapFix field.
func (o *TenantSettingsUpdateFlags) SetDisableFieldsMapFix(v bool) {
	o.DisableFieldsMapFix = &v
}

// GetMfaShowFactorListOnEnrollment returns the MfaShowFactorListOnEnrollment field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetMfaShowFactorListOnEnrollment() bool {
	if o == nil || IsNil(o.MfaShowFactorListOnEnrollment) {
		var ret bool
		return ret
	}
	return *o.MfaShowFactorListOnEnrollment
}

// GetMfaShowFactorListOnEnrollmentOk returns a tuple with the MfaShowFactorListOnEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetMfaShowFactorListOnEnrollmentOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaShowFactorListOnEnrollment) {
		return nil, false
	}
	return o.MfaShowFactorListOnEnrollment, true
}

// HasMfaShowFactorListOnEnrollment returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasMfaShowFactorListOnEnrollment() bool {
	if o != nil && !IsNil(o.MfaShowFactorListOnEnrollment) {
		return true
	}

	return false
}

// SetMfaShowFactorListOnEnrollment gets a reference to the given bool and assigns it to the MfaShowFactorListOnEnrollment field.
func (o *TenantSettingsUpdateFlags) SetMfaShowFactorListOnEnrollment(v bool) {
	o.MfaShowFactorListOnEnrollment = &v
}

// GetRemoveAlgFromJwks returns the RemoveAlgFromJwks field value if set, zero value otherwise.
func (o *TenantSettingsUpdateFlags) GetRemoveAlgFromJwks() bool {
	if o == nil || IsNil(o.RemoveAlgFromJwks) {
		var ret bool
		return ret
	}
	return *o.RemoveAlgFromJwks
}

// GetRemoveAlgFromJwksOk returns a tuple with the RemoveAlgFromJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsUpdateFlags) GetRemoveAlgFromJwksOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveAlgFromJwks) {
		return nil, false
	}
	return o.RemoveAlgFromJwks, true
}

// HasRemoveAlgFromJwks returns a boolean if a field has been set.
func (o *TenantSettingsUpdateFlags) HasRemoveAlgFromJwks() bool {
	if o != nil && !IsNil(o.RemoveAlgFromJwks) {
		return true
	}

	return false
}

// SetRemoveAlgFromJwks gets a reference to the given bool and assigns it to the RemoveAlgFromJwks field.
func (o *TenantSettingsUpdateFlags) SetRemoveAlgFromJwks(v bool) {
	o.RemoveAlgFromJwks = &v
}

func (o TenantSettingsUpdateFlags) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantSettingsUpdateFlags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangePwdFlowV1) {
		toSerialize["change_pwd_flow_v1"] = o.ChangePwdFlowV1
	}
	if !IsNil(o.EnableClientConnections) {
		toSerialize["enable_client_connections"] = o.EnableClientConnections
	}
	if !IsNil(o.EnableApisSection) {
		toSerialize["enable_apis_section"] = o.EnableApisSection
	}
	if !IsNil(o.EnablePipeline2) {
		toSerialize["enable_pipeline2"] = o.EnablePipeline2
	}
	if !IsNil(o.EnableDynamicClientRegistration) {
		toSerialize["enable_dynamic_client_registration"] = o.EnableDynamicClientRegistration
	}
	if !IsNil(o.EnableCustomDomainInEmails) {
		toSerialize["enable_custom_domain_in_emails"] = o.EnableCustomDomainInEmails
	}
	if !IsNil(o.AllowLegacyTokeninfoEndpoint) {
		toSerialize["allow_legacy_tokeninfo_endpoint"] = o.AllowLegacyTokeninfoEndpoint
	}
	if !IsNil(o.EnableLegacyProfile) {
		toSerialize["enable_legacy_profile"] = o.EnableLegacyProfile
	}
	if !IsNil(o.EnableIdtokenApi2) {
		toSerialize["enable_idtoken_api2"] = o.EnableIdtokenApi2
	}
	if !IsNil(o.EnablePublicSignupUserExistsError) {
		toSerialize["enable_public_signup_user_exists_error"] = o.EnablePublicSignupUserExistsError
	}
	if !IsNil(o.AllowLegacyDelegationGrantTypes) {
		toSerialize["allow_legacy_delegation_grant_types"] = o.AllowLegacyDelegationGrantTypes
	}
	if !IsNil(o.AllowLegacyRoGrantTypes) {
		toSerialize["allow_legacy_ro_grant_types"] = o.AllowLegacyRoGrantTypes
	}
	if !IsNil(o.EnableSso) {
		toSerialize["enable_sso"] = o.EnableSso
	}
	if !IsNil(o.DisableClickjackProtectionHeaders) {
		toSerialize["disable_clickjack_protection_headers"] = o.DisableClickjackProtectionHeaders
	}
	if !IsNil(o.NoDiscloseEnterpriseConnections) {
		toSerialize["no_disclose_enterprise_connections"] = o.NoDiscloseEnterpriseConnections
	}
	if !IsNil(o.DisableManagementApiSmsObfuscation) {
		toSerialize["disable_management_api_sms_obfuscation"] = o.DisableManagementApiSmsObfuscation
	}
	if !IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		toSerialize["enforce_client_authentication_on_passwordless_start"] = o.EnforceClientAuthenticationOnPasswordlessStart
	}
	if !IsNil(o.TrustAzureAdfsEmailVerifiedConnectionProperty) {
		toSerialize["trust_azure_adfs_email_verified_connection_property"] = o.TrustAzureAdfsEmailVerifiedConnectionProperty
	}
	if !IsNil(o.EnableAdfsWaadEmailVerification) {
		toSerialize["enable_adfs_waad_email_verification"] = o.EnableAdfsWaadEmailVerification
	}
	if !IsNil(o.RevokeRefreshTokenGrant) {
		toSerialize["revoke_refresh_token_grant"] = o.RevokeRefreshTokenGrant
	}
	if !IsNil(o.DashboardLogStreamsNext) {
		toSerialize["dashboard_log_streams_next"] = o.DashboardLogStreamsNext
	}
	if !IsNil(o.DashboardInsightsView) {
		toSerialize["dashboard_insights_view"] = o.DashboardInsightsView
	}
	if !IsNil(o.DisableFieldsMapFix) {
		toSerialize["disable_fields_map_fix"] = o.DisableFieldsMapFix
	}
	if !IsNil(o.MfaShowFactorListOnEnrollment) {
		toSerialize["mfa_show_factor_list_on_enrollment"] = o.MfaShowFactorListOnEnrollment
	}
	if !IsNil(o.RemoveAlgFromJwks) {
		toSerialize["remove_alg_from_jwks"] = o.RemoveAlgFromJwks
	}
	return toSerialize, nil
}

type NullableTenantSettingsUpdateFlags struct {
	value *TenantSettingsUpdateFlags
	isSet bool
}

func (v NullableTenantSettingsUpdateFlags) Get() *TenantSettingsUpdateFlags {
	return v.value
}

func (v *NullableTenantSettingsUpdateFlags) Set(val *TenantSettingsUpdateFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantSettingsUpdateFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantSettingsUpdateFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantSettingsUpdateFlags(val *TenantSettingsUpdateFlags) *NullableTenantSettingsUpdateFlags {
	return &NullableTenantSettingsUpdateFlags{value: val, isSet: true}
}

func (v NullableTenantSettingsUpdateFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantSettingsUpdateFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
