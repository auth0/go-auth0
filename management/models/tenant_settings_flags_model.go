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

// TenantSettingsFlags Flags used to change the behavior of this tenant.
type TenantSettingsFlags struct {
	// Whether to use the older v1 change password flow (true, not recommended except for backward compatibility) or the newer safer flow (false, recommended).
	ChangePwdFlowV1 *bool `json:"change_pwd_flow_v1,omitempty"`
	// Whether the APIs section is enabled (true) or disabled (false).
	EnableApisSection *bool `json:"enable_apis_section,omitempty"`
	// Whether the impersonation functionality has been disabled (true) or not (false). Read-only.
	DisableImpersonation *bool `json:"disable_impersonation,omitempty"`
	// Whether all current connections should be enabled when a new client (application) is created (true, default) or not (false).
	EnableClientConnections *bool `json:"enable_client_connections,omitempty"`
	// Whether advanced API Authorization scenarios are enabled (true) or disabled (false).
	EnablePipeline2 *bool `json:"enable_pipeline2,omitempty"`
	// If enabled, clients are able to add legacy delegation grants.
	AllowLegacyDelegationGrantTypes *bool `json:"allow_legacy_delegation_grant_types,omitempty"`
	// If enabled, clients are able to add legacy RO grants.
	AllowLegacyRoGrantTypes *bool `json:"allow_legacy_ro_grant_types,omitempty"`
	// Whether the legacy `/tokeninfo` endpoint is enabled for your account (true) or unavailable (false).
	AllowLegacyTokeninfoEndpoint *bool `json:"allow_legacy_tokeninfo_endpoint,omitempty"`
	// Whether ID tokens and the userinfo endpoint includes a complete user profile (true) or only OpenID Connect claims (false).
	EnableLegacyProfile *bool `json:"enable_legacy_profile,omitempty"`
	// Whether ID tokens can be used to authorize some types of requests to API v2 (true) not not (false).
	EnableIdtokenApi2 *bool `json:"enable_idtoken_api2,omitempty"`
	// Whether the public sign up process shows a user_exists error (true) or a generic error (false) if the user already exists.
	EnablePublicSignupUserExistsError *bool `json:"enable_public_signup_user_exists_error,omitempty"`
	// Whether users are prompted to confirm log in before SSO redirection (false) or are not prompted (true).
	EnableSso *bool `json:"enable_sso,omitempty"`
	// Whether the `enable_sso` setting can be changed (true) or not (false).
	AllowChangingEnableSso *bool `json:"allow_changing_enable_sso,omitempty"`
	// Whether classic Universal Login prompts include additional security headers to prevent clickjacking (true) or no safeguard (false).
	DisableClickjackProtectionHeaders *bool `json:"disable_clickjack_protection_headers,omitempty"`
	// Do not Publish Enterprise Connections Information with IdP domains on the lock configuration file.
	NoDiscloseEnterpriseConnections *bool `json:"no_disclose_enterprise_connections,omitempty"`
	// Enforce client authentication for passwordless start.
	EnforceClientAuthenticationOnPasswordlessStart *bool `json:"enforce_client_authentication_on_passwordless_start,omitempty"`
	// Enables the email verification flow during login for Azure AD and ADFS connections
	EnableAdfsWaadEmailVerification *bool `json:"enable_adfs_waad_email_verification,omitempty"`
	// Delete underlying grant when a Refresh Token is revoked via the Authentication API.
	RevokeRefreshTokenGrant *bool `json:"revoke_refresh_token_grant,omitempty"`
	// Enables beta access to log streaming changes
	DashboardLogStreamsNext *bool `json:"dashboard_log_streams_next,omitempty"`
	// Enables new insights activity page view
	DashboardInsightsView *bool `json:"dashboard_insights_view,omitempty"`
	// Disables SAML fields map fix for bad mappings with repeated attributes
	DisableFieldsMapFix *bool `json:"disable_fields_map_fix,omitempty"`
	// Used to allow users to pick what factor to enroll of the available MFA factors.
	MfaShowFactorListOnEnrollment *bool `json:"mfa_show_factor_list_on_enrollment,omitempty"`
	// Removes alg property from jwks .well-known endpoint
	RemoveAlgFromJwks *bool `json:"remove_alg_from_jwks,omitempty"`
	// Improves bot detection during signup in classic universal login
	ImprovedSignupBotDetectionInClassic *bool `json:"improved_signup_bot_detection_in_classic,omitempty"`
}

// GetChangePwdFlowV1 returns the ChangePwdFlowV1 field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetChangePwdFlowV1() bool {
	if o == nil || IsNil(o.ChangePwdFlowV1) {
		var ret bool
		return ret
	}
	return *o.ChangePwdFlowV1
}

// GetChangePwdFlowV1Ok returns a tuple with the ChangePwdFlowV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetChangePwdFlowV1Ok() (*bool, bool) {
	if o == nil || IsNil(o.ChangePwdFlowV1) {
		return nil, false
	}
	return o.ChangePwdFlowV1, true
}

// HasChangePwdFlowV1 returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasChangePwdFlowV1() bool {
	if o != nil && !IsNil(o.ChangePwdFlowV1) {
		return true
	}

	return false
}

// SetChangePwdFlowV1 gets a reference to the given bool and assigns it to the ChangePwdFlowV1 field.
func (o *TenantSettingsFlags) SetChangePwdFlowV1(v bool) {
	o.ChangePwdFlowV1 = &v
}

// GetEnableApisSection returns the EnableApisSection field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnableApisSection() bool {
	if o == nil || IsNil(o.EnableApisSection) {
		var ret bool
		return ret
	}
	return *o.EnableApisSection
}

// GetEnableApisSectionOk returns a tuple with the EnableApisSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnableApisSectionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableApisSection) {
		return nil, false
	}
	return o.EnableApisSection, true
}

// HasEnableApisSection returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnableApisSection() bool {
	if o != nil && !IsNil(o.EnableApisSection) {
		return true
	}

	return false
}

// SetEnableApisSection gets a reference to the given bool and assigns it to the EnableApisSection field.
func (o *TenantSettingsFlags) SetEnableApisSection(v bool) {
	o.EnableApisSection = &v
}

// GetDisableImpersonation returns the DisableImpersonation field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetDisableImpersonation() bool {
	if o == nil || IsNil(o.DisableImpersonation) {
		var ret bool
		return ret
	}
	return *o.DisableImpersonation
}

// GetDisableImpersonationOk returns a tuple with the DisableImpersonation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetDisableImpersonationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableImpersonation) {
		return nil, false
	}
	return o.DisableImpersonation, true
}

// HasDisableImpersonation returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasDisableImpersonation() bool {
	if o != nil && !IsNil(o.DisableImpersonation) {
		return true
	}

	return false
}

// SetDisableImpersonation gets a reference to the given bool and assigns it to the DisableImpersonation field.
func (o *TenantSettingsFlags) SetDisableImpersonation(v bool) {
	o.DisableImpersonation = &v
}

// GetEnableClientConnections returns the EnableClientConnections field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnableClientConnections() bool {
	if o == nil || IsNil(o.EnableClientConnections) {
		var ret bool
		return ret
	}
	return *o.EnableClientConnections
}

// GetEnableClientConnectionsOk returns a tuple with the EnableClientConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnableClientConnectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClientConnections) {
		return nil, false
	}
	return o.EnableClientConnections, true
}

// HasEnableClientConnections returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnableClientConnections() bool {
	if o != nil && !IsNil(o.EnableClientConnections) {
		return true
	}

	return false
}

// SetEnableClientConnections gets a reference to the given bool and assigns it to the EnableClientConnections field.
func (o *TenantSettingsFlags) SetEnableClientConnections(v bool) {
	o.EnableClientConnections = &v
}

// GetEnablePipeline2 returns the EnablePipeline2 field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnablePipeline2() bool {
	if o == nil || IsNil(o.EnablePipeline2) {
		var ret bool
		return ret
	}
	return *o.EnablePipeline2
}

// GetEnablePipeline2Ok returns a tuple with the EnablePipeline2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnablePipeline2Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnablePipeline2) {
		return nil, false
	}
	return o.EnablePipeline2, true
}

// HasEnablePipeline2 returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnablePipeline2() bool {
	if o != nil && !IsNil(o.EnablePipeline2) {
		return true
	}

	return false
}

// SetEnablePipeline2 gets a reference to the given bool and assigns it to the EnablePipeline2 field.
func (o *TenantSettingsFlags) SetEnablePipeline2(v bool) {
	o.EnablePipeline2 = &v
}

// GetAllowLegacyDelegationGrantTypes returns the AllowLegacyDelegationGrantTypes field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetAllowLegacyDelegationGrantTypes() bool {
	if o == nil || IsNil(o.AllowLegacyDelegationGrantTypes) {
		var ret bool
		return ret
	}
	return *o.AllowLegacyDelegationGrantTypes
}

// GetAllowLegacyDelegationGrantTypesOk returns a tuple with the AllowLegacyDelegationGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetAllowLegacyDelegationGrantTypesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLegacyDelegationGrantTypes) {
		return nil, false
	}
	return o.AllowLegacyDelegationGrantTypes, true
}

// HasAllowLegacyDelegationGrantTypes returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasAllowLegacyDelegationGrantTypes() bool {
	if o != nil && !IsNil(o.AllowLegacyDelegationGrantTypes) {
		return true
	}

	return false
}

// SetAllowLegacyDelegationGrantTypes gets a reference to the given bool and assigns it to the AllowLegacyDelegationGrantTypes field.
func (o *TenantSettingsFlags) SetAllowLegacyDelegationGrantTypes(v bool) {
	o.AllowLegacyDelegationGrantTypes = &v
}

// GetAllowLegacyRoGrantTypes returns the AllowLegacyRoGrantTypes field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetAllowLegacyRoGrantTypes() bool {
	if o == nil || IsNil(o.AllowLegacyRoGrantTypes) {
		var ret bool
		return ret
	}
	return *o.AllowLegacyRoGrantTypes
}

// GetAllowLegacyRoGrantTypesOk returns a tuple with the AllowLegacyRoGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetAllowLegacyRoGrantTypesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLegacyRoGrantTypes) {
		return nil, false
	}
	return o.AllowLegacyRoGrantTypes, true
}

// HasAllowLegacyRoGrantTypes returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasAllowLegacyRoGrantTypes() bool {
	if o != nil && !IsNil(o.AllowLegacyRoGrantTypes) {
		return true
	}

	return false
}

// SetAllowLegacyRoGrantTypes gets a reference to the given bool and assigns it to the AllowLegacyRoGrantTypes field.
func (o *TenantSettingsFlags) SetAllowLegacyRoGrantTypes(v bool) {
	o.AllowLegacyRoGrantTypes = &v
}

// GetAllowLegacyTokeninfoEndpoint returns the AllowLegacyTokeninfoEndpoint field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetAllowLegacyTokeninfoEndpoint() bool {
	if o == nil || IsNil(o.AllowLegacyTokeninfoEndpoint) {
		var ret bool
		return ret
	}
	return *o.AllowLegacyTokeninfoEndpoint
}

// GetAllowLegacyTokeninfoEndpointOk returns a tuple with the AllowLegacyTokeninfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetAllowLegacyTokeninfoEndpointOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLegacyTokeninfoEndpoint) {
		return nil, false
	}
	return o.AllowLegacyTokeninfoEndpoint, true
}

// HasAllowLegacyTokeninfoEndpoint returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasAllowLegacyTokeninfoEndpoint() bool {
	if o != nil && !IsNil(o.AllowLegacyTokeninfoEndpoint) {
		return true
	}

	return false
}

// SetAllowLegacyTokeninfoEndpoint gets a reference to the given bool and assigns it to the AllowLegacyTokeninfoEndpoint field.
func (o *TenantSettingsFlags) SetAllowLegacyTokeninfoEndpoint(v bool) {
	o.AllowLegacyTokeninfoEndpoint = &v
}

// GetEnableLegacyProfile returns the EnableLegacyProfile field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnableLegacyProfile() bool {
	if o == nil || IsNil(o.EnableLegacyProfile) {
		var ret bool
		return ret
	}
	return *o.EnableLegacyProfile
}

// GetEnableLegacyProfileOk returns a tuple with the EnableLegacyProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnableLegacyProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLegacyProfile) {
		return nil, false
	}
	return o.EnableLegacyProfile, true
}

// HasEnableLegacyProfile returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnableLegacyProfile() bool {
	if o != nil && !IsNil(o.EnableLegacyProfile) {
		return true
	}

	return false
}

// SetEnableLegacyProfile gets a reference to the given bool and assigns it to the EnableLegacyProfile field.
func (o *TenantSettingsFlags) SetEnableLegacyProfile(v bool) {
	o.EnableLegacyProfile = &v
}

// GetEnableIdtokenApi2 returns the EnableIdtokenApi2 field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnableIdtokenApi2() bool {
	if o == nil || IsNil(o.EnableIdtokenApi2) {
		var ret bool
		return ret
	}
	return *o.EnableIdtokenApi2
}

// GetEnableIdtokenApi2Ok returns a tuple with the EnableIdtokenApi2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnableIdtokenApi2Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableIdtokenApi2) {
		return nil, false
	}
	return o.EnableIdtokenApi2, true
}

// HasEnableIdtokenApi2 returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnableIdtokenApi2() bool {
	if o != nil && !IsNil(o.EnableIdtokenApi2) {
		return true
	}

	return false
}

// SetEnableIdtokenApi2 gets a reference to the given bool and assigns it to the EnableIdtokenApi2 field.
func (o *TenantSettingsFlags) SetEnableIdtokenApi2(v bool) {
	o.EnableIdtokenApi2 = &v
}

// GetEnablePublicSignupUserExistsError returns the EnablePublicSignupUserExistsError field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnablePublicSignupUserExistsError() bool {
	if o == nil || IsNil(o.EnablePublicSignupUserExistsError) {
		var ret bool
		return ret
	}
	return *o.EnablePublicSignupUserExistsError
}

// GetEnablePublicSignupUserExistsErrorOk returns a tuple with the EnablePublicSignupUserExistsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnablePublicSignupUserExistsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePublicSignupUserExistsError) {
		return nil, false
	}
	return o.EnablePublicSignupUserExistsError, true
}

// HasEnablePublicSignupUserExistsError returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnablePublicSignupUserExistsError() bool {
	if o != nil && !IsNil(o.EnablePublicSignupUserExistsError) {
		return true
	}

	return false
}

// SetEnablePublicSignupUserExistsError gets a reference to the given bool and assigns it to the EnablePublicSignupUserExistsError field.
func (o *TenantSettingsFlags) SetEnablePublicSignupUserExistsError(v bool) {
	o.EnablePublicSignupUserExistsError = &v
}

// GetEnableSso returns the EnableSso field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnableSso() bool {
	if o == nil || IsNil(o.EnableSso) {
		var ret bool
		return ret
	}
	return *o.EnableSso
}

// GetEnableSsoOk returns a tuple with the EnableSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnableSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSso) {
		return nil, false
	}
	return o.EnableSso, true
}

// HasEnableSso returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnableSso() bool {
	if o != nil && !IsNil(o.EnableSso) {
		return true
	}

	return false
}

// SetEnableSso gets a reference to the given bool and assigns it to the EnableSso field.
func (o *TenantSettingsFlags) SetEnableSso(v bool) {
	o.EnableSso = &v
}

// GetAllowChangingEnableSso returns the AllowChangingEnableSso field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetAllowChangingEnableSso() bool {
	if o == nil || IsNil(o.AllowChangingEnableSso) {
		var ret bool
		return ret
	}
	return *o.AllowChangingEnableSso
}

// GetAllowChangingEnableSsoOk returns a tuple with the AllowChangingEnableSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetAllowChangingEnableSsoOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChangingEnableSso) {
		return nil, false
	}
	return o.AllowChangingEnableSso, true
}

// HasAllowChangingEnableSso returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasAllowChangingEnableSso() bool {
	if o != nil && !IsNil(o.AllowChangingEnableSso) {
		return true
	}

	return false
}

// SetAllowChangingEnableSso gets a reference to the given bool and assigns it to the AllowChangingEnableSso field.
func (o *TenantSettingsFlags) SetAllowChangingEnableSso(v bool) {
	o.AllowChangingEnableSso = &v
}

// GetDisableClickjackProtectionHeaders returns the DisableClickjackProtectionHeaders field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetDisableClickjackProtectionHeaders() bool {
	if o == nil || IsNil(o.DisableClickjackProtectionHeaders) {
		var ret bool
		return ret
	}
	return *o.DisableClickjackProtectionHeaders
}

// GetDisableClickjackProtectionHeadersOk returns a tuple with the DisableClickjackProtectionHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetDisableClickjackProtectionHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableClickjackProtectionHeaders) {
		return nil, false
	}
	return o.DisableClickjackProtectionHeaders, true
}

// HasDisableClickjackProtectionHeaders returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasDisableClickjackProtectionHeaders() bool {
	if o != nil && !IsNil(o.DisableClickjackProtectionHeaders) {
		return true
	}

	return false
}

// SetDisableClickjackProtectionHeaders gets a reference to the given bool and assigns it to the DisableClickjackProtectionHeaders field.
func (o *TenantSettingsFlags) SetDisableClickjackProtectionHeaders(v bool) {
	o.DisableClickjackProtectionHeaders = &v
}

// GetNoDiscloseEnterpriseConnections returns the NoDiscloseEnterpriseConnections field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetNoDiscloseEnterpriseConnections() bool {
	if o == nil || IsNil(o.NoDiscloseEnterpriseConnections) {
		var ret bool
		return ret
	}
	return *o.NoDiscloseEnterpriseConnections
}

// GetNoDiscloseEnterpriseConnectionsOk returns a tuple with the NoDiscloseEnterpriseConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetNoDiscloseEnterpriseConnectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.NoDiscloseEnterpriseConnections) {
		return nil, false
	}
	return o.NoDiscloseEnterpriseConnections, true
}

// HasNoDiscloseEnterpriseConnections returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasNoDiscloseEnterpriseConnections() bool {
	if o != nil && !IsNil(o.NoDiscloseEnterpriseConnections) {
		return true
	}

	return false
}

// SetNoDiscloseEnterpriseConnections gets a reference to the given bool and assigns it to the NoDiscloseEnterpriseConnections field.
func (o *TenantSettingsFlags) SetNoDiscloseEnterpriseConnections(v bool) {
	o.NoDiscloseEnterpriseConnections = &v
}

// GetEnforceClientAuthenticationOnPasswordlessStart returns the EnforceClientAuthenticationOnPasswordlessStart field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnforceClientAuthenticationOnPasswordlessStart() bool {
	if o == nil || IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		var ret bool
		return ret
	}
	return *o.EnforceClientAuthenticationOnPasswordlessStart
}

// GetEnforceClientAuthenticationOnPasswordlessStartOk returns a tuple with the EnforceClientAuthenticationOnPasswordlessStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnforceClientAuthenticationOnPasswordlessStartOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		return nil, false
	}
	return o.EnforceClientAuthenticationOnPasswordlessStart, true
}

// HasEnforceClientAuthenticationOnPasswordlessStart returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnforceClientAuthenticationOnPasswordlessStart() bool {
	if o != nil && !IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		return true
	}

	return false
}

// SetEnforceClientAuthenticationOnPasswordlessStart gets a reference to the given bool and assigns it to the EnforceClientAuthenticationOnPasswordlessStart field.
func (o *TenantSettingsFlags) SetEnforceClientAuthenticationOnPasswordlessStart(v bool) {
	o.EnforceClientAuthenticationOnPasswordlessStart = &v
}

// GetEnableAdfsWaadEmailVerification returns the EnableAdfsWaadEmailVerification field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetEnableAdfsWaadEmailVerification() bool {
	if o == nil || IsNil(o.EnableAdfsWaadEmailVerification) {
		var ret bool
		return ret
	}
	return *o.EnableAdfsWaadEmailVerification
}

// GetEnableAdfsWaadEmailVerificationOk returns a tuple with the EnableAdfsWaadEmailVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetEnableAdfsWaadEmailVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAdfsWaadEmailVerification) {
		return nil, false
	}
	return o.EnableAdfsWaadEmailVerification, true
}

// HasEnableAdfsWaadEmailVerification returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasEnableAdfsWaadEmailVerification() bool {
	if o != nil && !IsNil(o.EnableAdfsWaadEmailVerification) {
		return true
	}

	return false
}

// SetEnableAdfsWaadEmailVerification gets a reference to the given bool and assigns it to the EnableAdfsWaadEmailVerification field.
func (o *TenantSettingsFlags) SetEnableAdfsWaadEmailVerification(v bool) {
	o.EnableAdfsWaadEmailVerification = &v
}

// GetRevokeRefreshTokenGrant returns the RevokeRefreshTokenGrant field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetRevokeRefreshTokenGrant() bool {
	if o == nil || IsNil(o.RevokeRefreshTokenGrant) {
		var ret bool
		return ret
	}
	return *o.RevokeRefreshTokenGrant
}

// GetRevokeRefreshTokenGrantOk returns a tuple with the RevokeRefreshTokenGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetRevokeRefreshTokenGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.RevokeRefreshTokenGrant) {
		return nil, false
	}
	return o.RevokeRefreshTokenGrant, true
}

// HasRevokeRefreshTokenGrant returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasRevokeRefreshTokenGrant() bool {
	if o != nil && !IsNil(o.RevokeRefreshTokenGrant) {
		return true
	}

	return false
}

// SetRevokeRefreshTokenGrant gets a reference to the given bool and assigns it to the RevokeRefreshTokenGrant field.
func (o *TenantSettingsFlags) SetRevokeRefreshTokenGrant(v bool) {
	o.RevokeRefreshTokenGrant = &v
}

// GetDashboardLogStreamsNext returns the DashboardLogStreamsNext field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetDashboardLogStreamsNext() bool {
	if o == nil || IsNil(o.DashboardLogStreamsNext) {
		var ret bool
		return ret
	}
	return *o.DashboardLogStreamsNext
}

// GetDashboardLogStreamsNextOk returns a tuple with the DashboardLogStreamsNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetDashboardLogStreamsNextOk() (*bool, bool) {
	if o == nil || IsNil(o.DashboardLogStreamsNext) {
		return nil, false
	}
	return o.DashboardLogStreamsNext, true
}

// HasDashboardLogStreamsNext returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasDashboardLogStreamsNext() bool {
	if o != nil && !IsNil(o.DashboardLogStreamsNext) {
		return true
	}

	return false
}

// SetDashboardLogStreamsNext gets a reference to the given bool and assigns it to the DashboardLogStreamsNext field.
func (o *TenantSettingsFlags) SetDashboardLogStreamsNext(v bool) {
	o.DashboardLogStreamsNext = &v
}

// GetDashboardInsightsView returns the DashboardInsightsView field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetDashboardInsightsView() bool {
	if o == nil || IsNil(o.DashboardInsightsView) {
		var ret bool
		return ret
	}
	return *o.DashboardInsightsView
}

// GetDashboardInsightsViewOk returns a tuple with the DashboardInsightsView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetDashboardInsightsViewOk() (*bool, bool) {
	if o == nil || IsNil(o.DashboardInsightsView) {
		return nil, false
	}
	return o.DashboardInsightsView, true
}

// HasDashboardInsightsView returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasDashboardInsightsView() bool {
	if o != nil && !IsNil(o.DashboardInsightsView) {
		return true
	}

	return false
}

// SetDashboardInsightsView gets a reference to the given bool and assigns it to the DashboardInsightsView field.
func (o *TenantSettingsFlags) SetDashboardInsightsView(v bool) {
	o.DashboardInsightsView = &v
}

// GetDisableFieldsMapFix returns the DisableFieldsMapFix field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetDisableFieldsMapFix() bool {
	if o == nil || IsNil(o.DisableFieldsMapFix) {
		var ret bool
		return ret
	}
	return *o.DisableFieldsMapFix
}

// GetDisableFieldsMapFixOk returns a tuple with the DisableFieldsMapFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetDisableFieldsMapFixOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableFieldsMapFix) {
		return nil, false
	}
	return o.DisableFieldsMapFix, true
}

// HasDisableFieldsMapFix returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasDisableFieldsMapFix() bool {
	if o != nil && !IsNil(o.DisableFieldsMapFix) {
		return true
	}

	return false
}

// SetDisableFieldsMapFix gets a reference to the given bool and assigns it to the DisableFieldsMapFix field.
func (o *TenantSettingsFlags) SetDisableFieldsMapFix(v bool) {
	o.DisableFieldsMapFix = &v
}

// GetMfaShowFactorListOnEnrollment returns the MfaShowFactorListOnEnrollment field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetMfaShowFactorListOnEnrollment() bool {
	if o == nil || IsNil(o.MfaShowFactorListOnEnrollment) {
		var ret bool
		return ret
	}
	return *o.MfaShowFactorListOnEnrollment
}

// GetMfaShowFactorListOnEnrollmentOk returns a tuple with the MfaShowFactorListOnEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetMfaShowFactorListOnEnrollmentOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaShowFactorListOnEnrollment) {
		return nil, false
	}
	return o.MfaShowFactorListOnEnrollment, true
}

// HasMfaShowFactorListOnEnrollment returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasMfaShowFactorListOnEnrollment() bool {
	if o != nil && !IsNil(o.MfaShowFactorListOnEnrollment) {
		return true
	}

	return false
}

// SetMfaShowFactorListOnEnrollment gets a reference to the given bool and assigns it to the MfaShowFactorListOnEnrollment field.
func (o *TenantSettingsFlags) SetMfaShowFactorListOnEnrollment(v bool) {
	o.MfaShowFactorListOnEnrollment = &v
}

// GetRemoveAlgFromJwks returns the RemoveAlgFromJwks field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetRemoveAlgFromJwks() bool {
	if o == nil || IsNil(o.RemoveAlgFromJwks) {
		var ret bool
		return ret
	}
	return *o.RemoveAlgFromJwks
}

// GetRemoveAlgFromJwksOk returns a tuple with the RemoveAlgFromJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetRemoveAlgFromJwksOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveAlgFromJwks) {
		return nil, false
	}
	return o.RemoveAlgFromJwks, true
}

// HasRemoveAlgFromJwks returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasRemoveAlgFromJwks() bool {
	if o != nil && !IsNil(o.RemoveAlgFromJwks) {
		return true
	}

	return false
}

// SetRemoveAlgFromJwks gets a reference to the given bool and assigns it to the RemoveAlgFromJwks field.
func (o *TenantSettingsFlags) SetRemoveAlgFromJwks(v bool) {
	o.RemoveAlgFromJwks = &v
}

// GetImprovedSignupBotDetectionInClassic returns the ImprovedSignupBotDetectionInClassic field value if set, zero value otherwise.
func (o *TenantSettingsFlags) GetImprovedSignupBotDetectionInClassic() bool {
	if o == nil || IsNil(o.ImprovedSignupBotDetectionInClassic) {
		var ret bool
		return ret
	}
	return *o.ImprovedSignupBotDetectionInClassic
}

// GetImprovedSignupBotDetectionInClassicOk returns a tuple with the ImprovedSignupBotDetectionInClassic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsFlags) GetImprovedSignupBotDetectionInClassicOk() (*bool, bool) {
	if o == nil || IsNil(o.ImprovedSignupBotDetectionInClassic) {
		return nil, false
	}
	return o.ImprovedSignupBotDetectionInClassic, true
}

// HasImprovedSignupBotDetectionInClassic returns a boolean if a field has been set.
func (o *TenantSettingsFlags) HasImprovedSignupBotDetectionInClassic() bool {
	if o != nil && !IsNil(o.ImprovedSignupBotDetectionInClassic) {
		return true
	}

	return false
}

// SetImprovedSignupBotDetectionInClassic gets a reference to the given bool and assigns it to the ImprovedSignupBotDetectionInClassic field.
func (o *TenantSettingsFlags) SetImprovedSignupBotDetectionInClassic(v bool) {
	o.ImprovedSignupBotDetectionInClassic = &v
}

func (o TenantSettingsFlags) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantSettingsFlags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangePwdFlowV1) {
		toSerialize["change_pwd_flow_v1"] = o.ChangePwdFlowV1
	}
	if !IsNil(o.EnableApisSection) {
		toSerialize["enable_apis_section"] = o.EnableApisSection
	}
	if !IsNil(o.DisableImpersonation) {
		toSerialize["disable_impersonation"] = o.DisableImpersonation
	}
	if !IsNil(o.EnableClientConnections) {
		toSerialize["enable_client_connections"] = o.EnableClientConnections
	}
	if !IsNil(o.EnablePipeline2) {
		toSerialize["enable_pipeline2"] = o.EnablePipeline2
	}
	if !IsNil(o.AllowLegacyDelegationGrantTypes) {
		toSerialize["allow_legacy_delegation_grant_types"] = o.AllowLegacyDelegationGrantTypes
	}
	if !IsNil(o.AllowLegacyRoGrantTypes) {
		toSerialize["allow_legacy_ro_grant_types"] = o.AllowLegacyRoGrantTypes
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
	if !IsNil(o.EnableSso) {
		toSerialize["enable_sso"] = o.EnableSso
	}
	if !IsNil(o.AllowChangingEnableSso) {
		toSerialize["allow_changing_enable_sso"] = o.AllowChangingEnableSso
	}
	if !IsNil(o.DisableClickjackProtectionHeaders) {
		toSerialize["disable_clickjack_protection_headers"] = o.DisableClickjackProtectionHeaders
	}
	if !IsNil(o.NoDiscloseEnterpriseConnections) {
		toSerialize["no_disclose_enterprise_connections"] = o.NoDiscloseEnterpriseConnections
	}
	if !IsNil(o.EnforceClientAuthenticationOnPasswordlessStart) {
		toSerialize["enforce_client_authentication_on_passwordless_start"] = o.EnforceClientAuthenticationOnPasswordlessStart
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
	if !IsNil(o.ImprovedSignupBotDetectionInClassic) {
		toSerialize["improved_signup_bot_detection_in_classic"] = o.ImprovedSignupBotDetectionInClassic
	}
	return toSerialize, nil
}

type NullableTenantSettingsFlags struct {
	value *TenantSettingsFlags
	isSet bool
}

func (v NullableTenantSettingsFlags) Get() *TenantSettingsFlags {
	return v.value
}

func (v *NullableTenantSettingsFlags) Set(val *TenantSettingsFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantSettingsFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantSettingsFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantSettingsFlags(val *TenantSettingsFlags) *NullableTenantSettingsFlags {
	return &NullableTenantSettingsFlags{value: val, isSet: true}
}

func (v NullableTenantSettingsFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantSettingsFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
