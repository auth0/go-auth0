package oauth

import "time"

// ClientAuthentication defines the authentication options that can be overridden per request.
type ClientAuthentication struct {
	// ClientID to use for the specific request.
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret to use for the specific request. Required when Client Secret Basic or Client
	// Secret Post is the application authentication method.
	ClientSecret string `json:"client_secret,omitempty"`
	// ClientAssertion to use for the specific request. Required if `Private Key JWT` is the
	// authentication method.
	ClientAssertion string `json:"client_assertion,omitempty"`
	// ClientAssertionType to use for the specific request. Required if you are passing your own
	// ClientAssertion.
	ClientAssertionType string `json:"client_assertion_type,omitempty"`
}

// TokenSet defines the response of the OAuth endpoints.
type TokenSet struct {
	// The access token.
	AccessToken string `json:"access_token,omitempty"`
	// The duration in seconds that the access token is valid for.
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// The user's ID token.
	IDToken string `json:"id_token,omitempty"`
	// The refresh token, only available if `offline_access` scope was provided.
	RefreshToken string `json:"refresh_token,omitempty"`
	// String value of the different scopes the application is asking for.
	// Multiple scopes are separated with whitespace.
	Scope string `json:"scope,omitempty"`
	// The type of the access token.
	TokenType string `json:"token_type,omitempty"`
}

// LoginWithPasswordRequest defines the request body for logging in with the Password grant.
type LoginWithPasswordRequest struct {
	ClientAuthentication
	// The user's username.
	Username string
	// The user's password.
	Password string
	// String value of the different scopes the application is asking for. Multiple scopes are separated with whitespace.
	Scope string
	// The unique identifier of the target API you want to access.
	Audience string
	// String value of the realm the user belongs. Set this if you want to add realm support to this grant.
	Realm string
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string
}

// LoginWithAuthCodeRequest defines the request body for logging in with the Authorization Code grant.
type LoginWithAuthCodeRequest struct {
	ClientAuthentication
	// The Authorization Code received from the initial /authorize call.
	Code string
	// This is required only if it was set at the GET /authorize endpoint. The values must match.
	RedirectURI string
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string
}

// LoginWithAuthCodeWithPKCERequest defines the request body for logging in with the Authorization Code with
// Proof Key for Code Exchange grant.
type LoginWithAuthCodeWithPKCERequest struct {
	ClientAuthentication
	// The Authorization Code received from the initial /authorize call.
	Code string
	// Cryptographically random key that was used to generate the code_challenge passed to /authorize.
	CodeVerifier string
	// This is required only if it was set at the GET /authorize endpoint. The values must match.
	RedirectURI string
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string
}

// LoginWithClientCredentialsRequest defines the request body for logging in with Authorization Code grant.
type LoginWithClientCredentialsRequest struct {
	ClientAuthentication
	// The unique identifier of the target API you want to access.
	Audience string
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string
}

// RefreshTokenRequest defines the request body for logging in with Authorization Code grant.
type RefreshTokenRequest struct {
	ClientAuthentication
	// The refresh token to use.
	RefreshToken string
	// 	A space-delimited list of requested scope permissions. If not sent, the original scopes will be used;
	// otherwise you can request a reduced set of scopes. Note that this must be URL encoded.
	Scope string
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string
}

// RevokeRefreshTokenRequest defines the request body for logging in with Authorization Code grant.
type RevokeRefreshTokenRequest struct {
	ClientAuthentication
	// The refresh token you want to revoke.
	Token string `json:"token,omitempty"`
	// Extra parameters to be merged into the request body. Values set here will override any existing values.
	ExtraParameters map[string]string `json:"-"`
}

// IDTokenValidationOptions allows validating optional claims that might not always be in the ID token.
type IDTokenValidationOptions struct {
	MaxAge       time.Duration
	Nonce        string
	Organization string
}
