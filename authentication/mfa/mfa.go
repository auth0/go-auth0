package mfa

import (
	"github.com/auth0/go-auth0/v2/authentication/oauth"
)

// ChallengeRequest defines the request body for requesting an MFA challenge.
type ChallengeRequest struct {
	oauth.ClientAuthentication
	// The token received from the `mfa_required` error.
	MFAToken string `json:"mfa_token,omitempty"`
	// A whitespace-separated list of the challenges types accepted by your application.
	// Accepted challenge types are "oob" or "otp". Excluding this parameter means that your
	// client application accepts all supported challenge types.
	ChallengeType string `json:"challenge_type,omitempty"`
	// The ID of the authenticator to challenge. You can get the ID by querying the list of
	// available authenticators for the user using `management.User.ListAuthenticationMethods`.
	AuthenticatorID string `json:"authenticator_id,omitempty"`
}

// ChallengeResponse defines the response body when requesting an MFA challenge.
type ChallengeResponse struct {
	// The type of challenge requested.
	ChallengeType string `json:"challenge_type,omitempty"`
	// The OOB code to use when calling `VerifyWithOOBRequest`
	// Only present when `ChallengeType` is "oob".
	OOBCode string `json:"oob_code,omitempty"`
	/// If included, then the user should be prompted for a `BindingCode` which should be included
	// in the `VerifyWithOOBRequest` provided to `VerifyWithOOB`.
	// Only present when `ChallengeType` is "oob".
	BindingMethod string `json:"binding_method,omitempty"`
}

// VerifyWithOTPRequest defines the request body for verifying an MFA challenge with OTP.
type VerifyWithOTPRequest struct {
	oauth.ClientAuthentication
	MFAToken string
	OTP      string
}

// VerifyWithOOBRequest defines the request body for verifying an MFA challenge with an OOB challenge.
type VerifyWithOOBRequest struct {
	oauth.ClientAuthentication
	MFAToken    string
	OOBCode     string
	BindingCode string
}

// VerifyWithRecoveryCodeRequest defines the request body for verifying an MFA challenge with a
// recovery code.
type VerifyWithRecoveryCodeRequest struct {
	oauth.ClientAuthentication
	MFAToken     string
	RecoveryCode string
}

// VerifyWithRecoveryCodeResponse defines the response when verifying with a recovery code.
type VerifyWithRecoveryCodeResponse struct {
	oauth.TokenSet
	// If present, a new recovery code that should be presented to the user to store.
	RecoveryCode string `json:"recovery_code,omitempty"`
}

// AddAuthenticatorRequest defines the request body for adding an authenticator.
type AddAuthenticatorRequest struct {
	oauth.ClientAuthentication
	// The type of authenticators supported by the client.
	// An array with values "otp" or "oob".
	AuthenticatorTypes []string `json:"authenticator_types"`
	// The type of OOB channels supported by the client.
	// An array with values "auth0", "sms", "voice".
	// Required if authenticator_types include oob.
	OOBChannels []string `json:"oob_channels,omitempty"`
	// The phone number to use for SMS or Voice.
	// Required if oob_channels includes sms or voice.
	PhoneNumber string `json:"phone_number,omitempty"`
}

// AddAuthenticatorResponse defines the response when adding an authenticator.
type AddAuthenticatorResponse struct {
	// If present, the OOB code that should be presented to the user to verify the authenticator.
	OOBCode string `json:"oob_code,omitempty"`
	// If present, a new recovery code that should be presented to the user to store.
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
	// The URI to generate a QR code for the authenticator.
	BarcodeURI string `json:"barcode_uri,omitempty"`
	// The secret to use for the OTP.
	Secret string `json:"secret,omitempty"`
	// The type of authenticator added.
	AuthenticatorType string `json:"authenticator_type,omitempty"`
	// Deprecated: Use OOBChannel instead.
	// The OOB channels supported by the authenticator.
	OOBChannels string `json:"oob_channels,omitempty"`
	// The OOB channel used for the authenticator.
	OOBChannel string `json:"oob_channel,omitempty"`
	// The binding method to use when verifying the authenticator.
	BindingMethod string `json:"binding_method,omitempty"`
}

// ListAuthenticatorsResponse defines the response when listing authenticators.
type ListAuthenticatorsResponse struct {
	ID                string `json:"id,omitempty"`
	AuthenticatorType string `json:"authenticator_type,omitempty"`
	// Deprecated: Use OOBChannel instead.
	OOBChannels string `json:"oob_channels,omitempty"`
	OOBChannel  string `json:"oob_channel,omitempty"`
	Name        string `json:"name,omitempty"`
	Active      bool   `json:"active,omitempty"`
}
