package passwordless

import "github.com/auth0/go-auth0/authentication/oauth"

// SendEmailRequest defines the request body for starting a passwordless flow via email.
type SendEmailRequest struct {
	oauth.ClientAuthentication
	Connection string                 `json:"connection,omitempty"`
	Email      string                 `json:"email,omitempty"`
	Send       string                 `json:"send,omitempty"`
	AuthParams map[string]interface{} `json:"authParams,omitempty"`
}

// SendEmailResponse defines the response from the `SendEmail` request.
type SendEmailResponse struct {
	ID            string `json:"_id,omitempty"`
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
}

// LoginWithEmailRequest defines the request body for exchanging a code requested by `SendEmail` for a token.
type LoginWithEmailRequest struct {
	oauth.ClientAuthentication
	Audience  string `json:"audience,omitempty"`
	Code      string `json:"otp,omitempty"`
	Email     string `json:"username,omitempty"`
	GrantType string `json:"grant_type,omitempty"`
	Realm     string `json:"realm,omitempty"`
	Scope     string `json:"scope,omitempty"`
}

// SendSMSRequest defines the request body for starting a passwordless flow via email.
type SendSMSRequest struct {
	oauth.ClientAuthentication
	Connection  string                 `json:"connection,omitempty"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	AuthParams  map[string]interface{} `json:"authParams,omitempty"`
}

// SendSMSResponse defines the response from the `SendSMS` request.
type SendSMSResponse struct {
	ID            string `json:"_id,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	PhoneVerified bool   `json:"phone_verified,omitempty"`
}

// LoginWithSMSRequest defines the request body for exchanging a code requested by `SendSMS` for a token.
type LoginWithSMSRequest struct {
	oauth.ClientAuthentication
	Audience    string `json:"audience,omitempty"`
	Code        string `json:"otp,omitempty"`
	PhoneNumber string `json:"username,omitempty"`
	GrantType   string `json:"grant_type,omitempty"`
	Realm       string `json:"realm,omitempty"`
	Scope       string `json:"scope,omitempty"`
}
