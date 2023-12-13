package mfa

import (
	"github.com/auth0/go-auth0/authentication/oauth"
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
