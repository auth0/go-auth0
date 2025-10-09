package passwordless

import "github.com/auth0/go-auth0/v2/authentication/oauth"

// SendEmailRequest defines the request body for starting a passwordless flow via email.
type SendEmailRequest struct {
	oauth.ClientAuthentication
	Connection string `json:"connection,omitempty"`
	// The users's email address.
	Email string `json:"email,omitempty"`
	// Use `link` to send a link or `code` to send a verification code. If omitted, a `link` will be sent.
	Send string `json:"send,omitempty"`
	// Append or override the link parameters (like `scope`, `redirect_uri`, `protocol`, `response_type`), when you send a link.
	AuthParams map[string]interface{} `json:"authParams,omitempty"`
}

// SendEmailResponse defines the response from the `SendEmail` request.
type SendEmailResponse struct {
	// The identifier of the request.
	ID string `json:"_id,omitempty"`
	// The user's email address.
	Email string `json:"email,omitempty"`
	// Whether the user's email address is verified.
	EmailVerified bool `json:"email_verified,omitempty"`
}

// LoginWithEmailRequest defines the request body for exchanging a code requested by `SendEmail` for a token.
type LoginWithEmailRequest struct {
	oauth.ClientAuthentication
	// API Identifier of the API for which you want to get an access token.
	Audience string `json:"audience,omitempty"`
	// The user's verification code.
	Code string `json:"otp,omitempty"`
	// The user's email address.
	Email     string `json:"username,omitempty"`
	GrantType string `json:"grant_type,omitempty"`
	Realm     string `json:"realm,omitempty"`
	// Use `openid` to get an ID token, or `openid profile email` to also include user profile information in the ID token.
	Scope string `json:"scope,omitempty"`
}

// SendSMSRequest defines the request body for starting a passwordless flow via email.
type SendSMSRequest struct {
	oauth.ClientAuthentication
	Connection string `json:"connection,omitempty"`
	// The user's phone number.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Append or override the link parameters (like `scope`, `redirect_uri`, `protocol`, `response_type`), when you send a link.
	AuthParams map[string]interface{} `json:"authParams,omitempty"`
}

// SendSMSResponse defines the response from the `SendSMS` request.
type SendSMSResponse struct {
	// The identifier of the request.
	ID string `json:"_id,omitempty"`
	// The user's phone number.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Whether the users phone number is verified.
	PhoneVerified bool `json:"phone_verified,omitempty"`
}

// LoginWithSMSRequest defines the request body for exchanging a code requested by `SendSMS` for a token.
type LoginWithSMSRequest struct {
	oauth.ClientAuthentication
	// API Identifier of the API for which you want to get an access token.
	Audience string `json:"audience,omitempty"`
	// The user's verification code.
	Code string `json:"otp,omitempty"`
	// The user's phone number.
	PhoneNumber string `json:"username,omitempty"`
	GrantType   string `json:"grant_type,omitempty"`
	Realm       string `json:"realm,omitempty"`
	// Use `openid` to get an ID token, or `openid profile email` to also include user profile information in the ID token.
	Scope string `json:"scope,omitempty"`
}
