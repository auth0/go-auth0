package management

import (
	"context"
	"encoding/json"
	"fmt"
)

// RenderingMode is a type that represents the rendering mode.
type RenderingMode string

var (
	// RenderingModeStandard represents the standard rendering mode.
	RenderingModeStandard RenderingMode = "standard"

	// RenderingModeAdvanced represents the advanced rendering mode.
	RenderingModeAdvanced RenderingMode = "advanced"
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

	// PromptLoginEmailVerification represents the login-email-verification prompt.
	PromptLoginEmailVerification PromptType = "login-email-verification"

	// PromptPhoneIdentifierEnrollment represents the phone-identifier-enrollment prompt.
	PromptPhoneIdentifierEnrollment PromptType = "phone-identifier-enrollment"

	// PromptPhoneIdentifierChallenge represents the phone-identifier-challenge prompt.
	PromptPhoneIdentifierChallenge PromptType = "phone-identifier-challenge"

	// PromptEmailIdentifierChallenge represents the email-identifier-challenge prompt.
	PromptEmailIdentifierChallenge PromptType = "email-identifier-challenge"

	// PromptResetPassword represents the reset-password prompt.
	PromptResetPassword PromptType = "reset-password"

	// PromptCustomForm represents the custom-form prompt.
	PromptCustomForm PromptType = "custom-form"

	// PromptConsent represents the consent prompt.
	PromptConsent PromptType = "consent"

	// PromptCustomizedConsent represents the customized-consent prompt.
	PromptCustomizedConsent PromptType = "customized-consent"

	// PromptLogout represents the logout prompt.
	PromptLogout PromptType = "logout"

	// PromptMFAPush represents the mfa-push prompt.
	PromptMFAPush PromptType = "mfa-push"

	// PromptMFAOTP represents the mfa-otp prompt.
	PromptMFAOTP PromptType = "mfa-otp"

	// PromptMFAVoice represents the mfa-voice prompt.
	PromptMFAVoice PromptType = "mfa-voice"

	// PromptMFAPhone represents the mfa-phone prompt.
	PromptMFAPhone PromptType = "mfa-phone"

	// PromptMFAWebAuthn represents the mfa-webauthn prompt.
	PromptMFAWebAuthn PromptType = "mfa-webauthn"

	// PromptMFASMS represents the mfa-sms prompt.
	PromptMFASMS PromptType = "mfa-sms"

	// PromptMFAEmail represents the mfa-email prompt.
	PromptMFAEmail PromptType = "mfa-email"

	// PromptMFARecoveryCode represents the mfa-recovery-code prompt.
	PromptMFARecoveryCode PromptType = "mfa-recovery-code"

	// PromptMFA represents the mfa prompt.
	PromptMFA PromptType = "mfa"

	// PromptStatus represents the status prompt.
	PromptStatus PromptType = "status"

	// PromptDeviceFlow represents the device-flow prompt.
	PromptDeviceFlow PromptType = "device-flow"

	// PromptEmailVerification represents the email-verification prompt.
	PromptEmailVerification PromptType = "email-verification"

	// PromptEmailOTPChallenge represents the email-otp-challenge prompt.
	PromptEmailOTPChallenge PromptType = "email-otp-challenge"

	// PromptOrganizations represents the organizations prompt.
	PromptOrganizations PromptType = "organizations"

	// PromptInvitation represents the invitation prompt.
	PromptInvitation PromptType = "invitation"

	// PromptCommon represents the common prompt.
	PromptCommon PromptType = "common"

	// PromptPasskeys represents the passkeys prompt.
	PromptPasskeys PromptType = "passkeys"

	// PromptCaptcha represents the captcha prompt.
	PromptCaptcha PromptType = "captcha"

	// PromptBruteForceProtection represents the brute-force-protection prompt.
	PromptBruteForceProtection PromptType = "brute-force-protection"
)

var allowedPromptsWithPartials = []PromptType{
	PromptSignup,
	PromptSignupID,
	PromptSignupPassword,
	PromptLogin,
	PromptLoginID,
	PromptLoginPassword,
	PromptLoginPasswordLess,
	PromptCustomizedConsent,
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

	// ScreenLoginPasswordlessEmailCode represents the login-passwordless-email-code screen.
	ScreenLoginPasswordlessEmailCode ScreenName = "login-passwordless-email-code"

	// ScreenLoginPasswordlessEmailLink represents the login-passwordless-email-link screen.
	ScreenLoginPasswordlessEmailLink ScreenName = "login-passwordless-email-link"

	// ScreenLoginPasswordlessSMSOTP represents the login-passwordless-sms-otp screen.
	ScreenLoginPasswordlessSMSOTP ScreenName = "login-passwordless-sms-otp"

	// ScreenLoginEmailVerification represents the login-email-verification screen.
	ScreenLoginEmailVerification ScreenName = "login-email-verification"

	// ScreenSignup represents the signup screen.
	ScreenSignup ScreenName = "signup"

	// ScreenSignupID represents the signup-id screen.
	ScreenSignupID ScreenName = "signup-id"

	// ScreenSignupPassword represents the signup-password screen.
	ScreenSignupPassword ScreenName = "signup-password"

	// ScreenPhoneIdentifierEnrollment represents the phone-identifier-enrollment screen.
	ScreenPhoneIdentifierEnrollment ScreenName = "phone-identifier-enrollment"

	// ScreenPhoneIdentifierChallenge represents the phone-identifier-challenge screen.
	ScreenPhoneIdentifierChallenge ScreenName = "phone-identifier-challenge"

	// ScreenEmailIdentifierChallenge represents the email-identifier-challenge screen.
	ScreenEmailIdentifierChallenge ScreenName = "email-identifier-challenge"

	// ScreenResetPasswordRequest represents the reset-password-request screen.
	ScreenResetPasswordRequest ScreenName = "reset-password-request"

	// ScreenResetPasswordEmail represents the reset-password-email screen.
	ScreenResetPasswordEmail ScreenName = "reset-password-email"

	// ScreenResetPassword represents the reset-password screen.
	ScreenResetPassword ScreenName = "reset-password"

	// ScreenResetPasswordSuccess represents the reset-password-success screen.
	ScreenResetPasswordSuccess ScreenName = "reset-password-success"

	// ScreenResetPasswordError represents the reset-password-error screen.
	ScreenResetPasswordError ScreenName = "reset-password-error"

	// ScreenResetPasswordMFAEmailChallenge represents the reset-password-mfa-email-challenge screen.
	ScreenResetPasswordMFAEmailChallenge ScreenName = "reset-password-mfa-email-challenge"

	// ScreenResetPasswordMFAOTPChallenge represents the reset-password-mfa-otp-challenge screen.
	ScreenResetPasswordMFAOTPChallenge ScreenName = "reset-password-mfa-otp-challenge"

	// ScreenResetPasswordMFAPhoneChallenge represents the reset-password-mfa-phone-challenge screen.
	ScreenResetPasswordMFAPhoneChallenge ScreenName = "reset-password-mfa-phone-challenge"

	// ScreenResetPasswordMFAPushChallengePush represents the reset-password-mfa-push-challenge-push screen.
	ScreenResetPasswordMFAPushChallengePush ScreenName = "reset-password-mfa-push-challenge-push"

	// ScreenResetPasswordMFARecoveryCodeChallenge represents the reset-password-mfa-recovery-code-challenge screen.
	ScreenResetPasswordMFARecoveryCodeChallenge ScreenName = "reset-password-mfa-recovery-code-challenge"

	// ScreenResetPasswordMFASMSChallenge represents the reset-password-mfa-sms-challenge screen.
	ScreenResetPasswordMFASMSChallenge ScreenName = "reset-password-mfa-sms-challenge"

	// ScreenResetPasswordMFAVoiceChallenge represents the reset-password-mfa-voice-challenge screen.
	ScreenResetPasswordMFAVoiceChallenge ScreenName = "reset-password-mfa-voice-challenge"

	// ScreenResetPasswordMFAWebAuthnPlatformChallenge represents the reset-password-mfa-webauthn-platform-challenge screen.
	ScreenResetPasswordMFAWebAuthnPlatformChallenge ScreenName = "reset-password-mfa-webauthn-platform-challenge"

	// ScreenResetPasswordMFAWebAuthnRoamingChallenge represents the reset-password-mfa-webauthn-roaming-challenge screen.
	ScreenResetPasswordMFAWebAuthnRoamingChallenge ScreenName = "reset-password-mfa-webauthn-roaming-challenge"

	// ScreenCustomForm represents the custom-form screen.
	ScreenCustomForm ScreenName = "custom-form"

	// ScreenConsent represents the consent screen.
	ScreenConsent ScreenName = "consent"

	// ScreenCustomizedConsent represents the customized-consent screen.
	ScreenCustomizedConsent ScreenName = "customized-consent"

	// ScreenLogout represents the logout screen.
	ScreenLogout ScreenName = "logout"

	// ScreenLogoutComplete represents the logout-complete screen.
	ScreenLogoutComplete ScreenName = "logout-complete"

	// ScreenLogoutAborted represents the logout-aborted screen.
	ScreenLogoutAborted ScreenName = "logout-aborted"

	// ScreenMFAPushWelcome represents the mfa-push-welcome screen.
	ScreenMFAPushWelcome ScreenName = "mfa-push-welcome"

	// ScreenMFAPushEnrollmentQR represents the mfa-push-enrollment-qr screen.
	ScreenMFAPushEnrollmentQR ScreenName = "mfa-push-enrollment-qr"

	// ScreenMFAPushEnrollmentCode represents the mfa-push-enrollment-code screen.
	ScreenMFAPushEnrollmentCode ScreenName = "mfa-push-enrollment-code"

	// ScreenMFAPushSuccess represents the mfa-push-success screen.
	ScreenMFAPushSuccess ScreenName = "mfa-push-success"

	// ScreenMFAPushChallengePush represents the mfa-push-challenge-push screen.
	ScreenMFAPushChallengePush ScreenName = "mfa-push-challenge-push"

	// ScreenMFAPushList represents the mfa-push-list screen.
	ScreenMFAPushList ScreenName = "mfa-push-list"

	// ScreenMFAOTPEnrollmentQR represents the mfa-otp-enrollment-qr screen.
	ScreenMFAOTPEnrollmentQR ScreenName = "mfa-otp-enrollment-qr"

	// ScreenMFAOTPEnrollmentCode represents the mfa-otp-enrollment-code screen.
	ScreenMFAOTPEnrollmentCode ScreenName = "mfa-otp-enrollment-code"

	// ScreenMFAOTPChallenge represents the mfa-otp-challenge screen.
	ScreenMFAOTPChallenge ScreenName = "mfa-otp-challenge"

	// ScreenMFAVoiceEnrollment represents the mfa-voice-enrollment screen.
	ScreenMFAVoiceEnrollment ScreenName = "mfa-voice-enrollment"

	// ScreenMFAVoiceChallenge represents the mfa-voice-challenge screen.
	ScreenMFAVoiceChallenge ScreenName = "mfa-voice-challenge"

	// ScreenMFAPhoneChallenge represents the mfa-phone-challenge screen.
	ScreenMFAPhoneChallenge ScreenName = "mfa-phone-challenge"

	// ScreenMFAPhoneEnrollment represents the mfa-phone-enrollment screen.
	ScreenMFAPhoneEnrollment ScreenName = "mfa-phone-enrollment"

	// ScreenMFAWebAuthnPlatformEnrollment represents the mfa-webauthn-platform-enrollment screen.
	ScreenMFAWebAuthnPlatformEnrollment ScreenName = "mfa-webauthn-platform-enrollment"

	// ScreenMFAWebAuthnRoamingEnrollment represents the mfa-webauthn-roaming-enrollment screen.
	ScreenMFAWebAuthnRoamingEnrollment ScreenName = "mfa-webauthn-roaming-enrollment"

	// ScreenMFAWebAuthnPlatformChallenge represents the mfa-webauthn-platform-challenge screen.
	ScreenMFAWebAuthnPlatformChallenge ScreenName = "mfa-webauthn-platform-challenge"

	// ScreenMFAWebAuthnRoamingChallenge represents the mfa-webauthn-roaming-challenge screen.
	ScreenMFAWebAuthnRoamingChallenge ScreenName = "mfa-webauthn-roaming-challenge"

	// ScreenMFAWebAuthnChangeKeyNickname represents the mfa-webauthn-change-key-nickname screen.
	ScreenMFAWebAuthnChangeKeyNickname ScreenName = "mfa-webauthn-change-key-nickname"

	// ScreenMFAWebAuthnEnrollmentSuccess represents the mfa-webauthn-enrollment-success screen.
	ScreenMFAWebAuthnEnrollmentSuccess ScreenName = "mfa-webauthn-enrollment-success"

	// ScreenMFAWebAuthnError represents the mfa-webauthn-error screen.
	ScreenMFAWebAuthnError ScreenName = "mfa-webauthn-error"

	// ScreenMFAWebAuthnNotAvailableError represents the mfa-webauthn-not-available-error screen.
	ScreenMFAWebAuthnNotAvailableError ScreenName = "mfa-webauthn-not-available-error"

	// ScreenMFACountryCodes represents the mfa-country-codes screen.
	ScreenMFACountryCodes ScreenName = "mfa-country-codes"

	// ScreenMFASMSEnrollment represents the mfa-sms-enrollment screen.
	ScreenMFASMSEnrollment ScreenName = "mfa-sms-enrollment"

	// ScreenMFASMSChallenge represents the mfa-sms-challenge screen.
	ScreenMFASMSChallenge ScreenName = "mfa-sms-challenge"

	// ScreenMFASMSList represents the mfa-sms-list screen.
	ScreenMFASMSList ScreenName = "mfa-sms-list"

	// ScreenMFAEmailChallenge represents the mfa-email-challenge screen.
	ScreenMFAEmailChallenge ScreenName = "mfa-email-challenge"

	// ScreenMFAEmailList represents the mfa-email-list screen.
	ScreenMFAEmailList ScreenName = "mfa-email-list"

	// ScreenMFARecoveryCodeEnrollment represents the mfa-recovery-code-enrollment screen.
	ScreenMFARecoveryCodeEnrollment ScreenName = "mfa-recovery-code-enrollment"

	// ScreenMFARecoveryCodeChallenge represents the mfa-recovery-code-challenge screen.
	ScreenMFARecoveryCodeChallenge ScreenName = "mfa-recovery-code-challenge"

	// ScreenMFADetectBrowserCapabilities represents the mfa-detect-browser-capabilities screen.
	ScreenMFADetectBrowserCapabilities ScreenName = "mfa-detect-browser-capabilities"

	// ScreenMFAEnrollResult represents the mfa-enroll-result screen.
	ScreenMFAEnrollResult ScreenName = "mfa-enroll-result"

	// ScreenMFALoginOptions represents the mfa-login-options screen.
	ScreenMFALoginOptions ScreenName = "mfa-login-options"

	// ScreenMFABeginEnrollOptions represents the mfa-begin-enroll-options screen.
	ScreenMFABeginEnrollOptions ScreenName = "mfa-begin-enroll-options"

	// ScreenMFARecoveryCodeChallengeNewCode represents the mfa-recovery-code-challenge-new-code screen.
	ScreenMFARecoveryCodeChallengeNewCode ScreenName = "mfa-recovery-code-challenge-new-code"

	// ScreenStatus represents the status screen.
	ScreenStatus ScreenName = "status"

	// ScreenDeviceCodeActivation represents the device-code-activation screen.
	ScreenDeviceCodeActivation ScreenName = "device-code-activation"

	// ScreenDeviceCodeActivationAllowed represents the device-code-activation-allowed screen.
	ScreenDeviceCodeActivationAllowed ScreenName = "device-code-activation-allowed"

	// ScreenDeviceCodeActivationDenied represents the device-code-activation-denied screen.
	ScreenDeviceCodeActivationDenied ScreenName = "device-code-activation-denied"

	// ScreenDeviceCodeConfirmation represents the device-code-confirmation screen.
	ScreenDeviceCodeConfirmation ScreenName = "device-code-confirmation"

	// ScreenEmailVerificationResult represents the email-verification-result screen.
	ScreenEmailVerificationResult ScreenName = "email-verification-result"

	// ScreenEmailOTPChallenge represents the email-otp-challenge screen.
	ScreenEmailOTPChallenge ScreenName = "email-otp-challenge"

	// ScreenOrganizationSelection represents the organization-selection screen.
	ScreenOrganizationSelection ScreenName = "organization-selection"

	// ScreenOrganizationPicker represents the organization-picker screen.
	ScreenOrganizationPicker ScreenName = "organization-picker"

	// ScreenAcceptInvitation represents the accept-invitation screen.
	ScreenAcceptInvitation ScreenName = "accept-invitation"

	// ScreenRedeemTicket represents the redeem-ticket screen.
	ScreenRedeemTicket ScreenName = "redeem-ticket"

	// ScreenPasskeyEnrollment represents the passkey-enrollment screen.
	ScreenPasskeyEnrollment ScreenName = "passkey-enrollment"

	// ScreenPasskeyEnrollmentLocal represents the passkey-enrollment-local screen.
	ScreenPasskeyEnrollmentLocal ScreenName = "passkey-enrollment-local"

	// ScreenInterstitialCaptcha represents the interstitial-captcha screen.
	ScreenInterstitialCaptcha ScreenName = "interstitial-captcha"

	// ScreenBruteForceProtectionUnblock represents the brute-force-protection-unblock screen.
	ScreenBruteForceProtectionUnblock ScreenName = "brute-force-protection-unblock"

	// ScreenBruteForceProtectionUnblockFailure represents the brute-force-protection-unblock-failure.
	ScreenBruteForceProtectionUnblockFailure ScreenName = "brute-force-protection-unblock-failure"

	// ScreenBruteForceProtectionUnblockSuccess represents the brute-force-protection-unblock-success.
	ScreenBruteForceProtectionUnblockSuccess ScreenName = "brute-force-protection-unblock-success"
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

	// InsertionPointSecondaryActionsStart represents the secondary-actions-start insertion point.
	InsertionPointSecondaryActionsStart InsertionPoint = "secondary-actions-start"

	// InsertionPointSecondaryActionsEnd represents the secondary-actions-end insertion point.
	InsertionPointSecondaryActionsEnd InsertionPoint = "secondary-actions-end"

	// InsertionPointFormContent represents the form-content insertion point.
	InsertionPointFormContent InsertionPoint = "form-content"
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

// PromptRendering is used to retrieve and set the settings for the ACUL.
type PromptRendering struct {
	Tenant                  *string        `json:"tenant,omitempty"`
	Prompt                  *PromptType    `json:"prompt,omitempty"`
	Screen                  *ScreenName    `json:"screen,omitempty"`
	RenderingMode           *RenderingMode `json:"rendering_mode,omitempty"`
	ContextConfiguration    *[]string      `json:"context_configuration,omitempty"`
	DefaultHeadTagsDisabled *bool          `json:"default_head_tags_disabled,omitempty"`
	HeadTags                []interface{}  `json:"head_tags,omitempty"`
	// Filters are optional filters to apply rendering rules to specific entities.
	Filters *PromptRenderingFilters `json:"filters,omitempty"`
	// UsePageTemplate indicates whether to use the page template with ACUL.
	UsePageTemplate *bool `json:"use_page_template,omitempty"`
}

// PromptRenderingFilters is used to filter rendering rules based on specific entities.
type PromptRenderingFilters struct {
	// MatchType defines the type of match to apply.
	MatchType *string `json:"match_type,omitempty"`
	// Clients is a list of client filters.
	Clients *[]PromptRenderingFilter `json:"clients,omitempty"`
	// Organizations is a list of organization filters.
	Organizations *[]PromptRenderingFilter `json:"organizations,omitempty"`
	// Domains is a list of domain filters.
	Domains *[]PromptRenderingFilter `json:"domains,omitempty"`
}

// PromptRenderingFilter is used to filter rendering rules based on specific entities.
type PromptRenderingFilter struct {
	// ID is the identifier of the entity.
	ID *string `json:"id,omitempty"`
	// Metadata is the metadata associated with the entity.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
}

// PromptRenderingList is a list of prompt rendering settings.
type PromptRenderingList struct {
	List
	PromptRenderings []*PromptRendering `json:"configs"`
}

// MarshalJSON implements a custom [json.Marshaler].
func (c *PromptRendering) MarshalJSON() ([]byte, error) {
	type RenderingSubSet struct {
		RenderingMode           *RenderingMode          `json:"rendering_mode,omitempty"`
		ContextConfiguration    *[]string               `json:"context_configuration,omitempty"`
		DefaultHeadTagsDisabled *bool                   `json:"default_head_tags_disabled,omitempty"`
		HeadTags                []interface{}           `json:"head_tags,omitempty"`
		Filters                 *PromptRenderingFilters `json:"filters,omitempty"`
		UsePageTemplate         *bool                   `json:"use_page_template,omitempty"`
	}

	return json.Marshal(&RenderingSubSet{
		RenderingMode:           c.RenderingMode,
		ContextConfiguration:    c.ContextConfiguration,
		DefaultHeadTagsDisabled: c.DefaultHeadTagsDisabled,
		HeadTags:                c.HeadTags,
		Filters:                 c.Filters,
		UsePageTemplate:         c.UsePageTemplate,
	})
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
// See: https://auth0.com/docs/sign-up-prompt-customizations#use-the-api-to-edit-custom-prompts
func (m *PromptManager) DeletePartials(ctx context.Context, prompt PromptType, opts ...RequestOption) error {
	if err := guardAgainstPromptTypesWithNoPartials(prompt); err != nil {
		return err
	}

	return m.management.Request(ctx, "PUT", m.management.URI("prompts", string(prompt), "partials"), &PromptScreenPartials{}, opts...)
}

func guardAgainstPromptTypesWithNoPartials(prompt PromptType) error {
	for _, p := range allowedPromptsWithPartials {
		if p == prompt {
			return nil
		}
	}

	return fmt.Errorf("cannot customize partials for prompt: %q", prompt)
}

// ReadRendering retrieves the settings for the ACUL.
//
// See: https://auth0.com/docs/api/management/v2/prompts/get-rendering
func (m *PromptManager) ReadRendering(ctx context.Context, prompt PromptType, screen ScreenName, opts ...RequestOption) (c *PromptRendering, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", string(prompt), "screen", string(screen), "rendering"), &c, opts...)
	return
}
func (c *PromptRendering) cleanForPatch() *PromptRendering {
	if c.RenderingMode != nil && *c.RenderingMode == RenderingModeStandard {
		return &PromptRendering{
			RenderingMode: c.RenderingMode,
		}
	}

	return c
}

// UpdateRendering updates the settings for the ACUL.
//
// See: https://auth0.com/docs/api/management/v2/prompts/patch-rendering
func (m *PromptManager) UpdateRendering(ctx context.Context, prompt PromptType, screen ScreenName, c *PromptRendering, opts ...RequestOption) error {
	if c != nil {
		c = c.cleanForPatch()
	}

	return m.management.Request(ctx, "PATCH", m.management.URI("prompts", string(prompt), "screen", string(screen), "rendering"), c, opts...)
}

// ListRendering lists the settings for all the ACUL.
//
// See: https://auth0.com/docs/api/management/v2/prompts/get-all-rendering
func (m *PromptManager) ListRendering(ctx context.Context, opts ...RequestOption) (c *PromptRenderingList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("prompts", "rendering"), &c, applyListDefaults(opts))
	return
}
