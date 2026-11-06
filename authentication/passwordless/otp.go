package passwordless

import "github.com/auth0/go-auth0/v3/authentication/oauth"

// ChallengeWithEmailRequest defines the request body for starting a passwordless OTP flow
// on a database connection using an email address.
type ChallengeWithEmailRequest struct {
	oauth.ClientAuthentication
	// The name of the database connection. Required.
	Connection string `json:"connection,omitempty"`
	// The user's email address. Required.
	Email string `json:"email,omitempty"`
	// Whether a new user should be created if one does not already exist. Always sent, defaults to false.
	AllowSignup bool `json:"allow_signup"`
}

// ChallengeWithPhoneNumberRequest defines the request body for starting a passwordless OTP flow
// on a database connection using a phone number.
type ChallengeWithPhoneNumberRequest struct {
	oauth.ClientAuthentication
	// The name of the database connection. Required.
	Connection string `json:"connection,omitempty"`
	// The user's phone number in E.164 format (e.g. +14155550100). Required.
	PhoneNumber string `json:"phone_number,omitempty"`
	// How the code is delivered, either `text` or `voice`. If omitted, the server defaults to `text`.
	DeliveryMethod string `json:"delivery_method,omitempty"`
	// Whether a new user should be created if one does not already exist. Always sent, defaults to false.
	AllowSignup bool `json:"allow_signup"`
}

// ChallengeResponse defines the response from the `ChallengeWithEmail` and `ChallengeWithPhoneNumber` requests.
type ChallengeResponse struct {
	// An opaque session identifier to be passed to `LoginWithOTPChallenge`. Callers must not parse or depend on its contents.
	AuthSession string `json:"auth_session,omitempty"`
}

// LoginWithOTPChallengeRequest defines the request body for exchanging a code and the auth session
// returned by `ChallengeWithEmail` or `ChallengeWithPhoneNumber` for a token.
type LoginWithOTPChallengeRequest struct {
	oauth.ClientAuthentication
	// The auth session returned by the challenge request. Required.
	AuthSession string
	// The one-time code the user received. Required.
	OTP string
	// Use `openid` to get an ID token, or `openid profile email` to also include user profile information in the ID token.
	Scope string
	// API Identifier of the API for which you want to get an access token.
	Audience string
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string
}
