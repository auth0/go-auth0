package oauth

// ClientAuthentication defines the authentication options that can be overridden per request.
type ClientAuthentication struct {
	// ClientID to use for the specific request.
	ClientID string
	// ClientSecret to use for the specific request .Required when Client Secret Basic or Client
	// Secret Post is the application authentication method.
	ClientSecret string
}

// TokenSet defines the response of the OAuth endpoints.
type TokenSet struct {
	// The access token.
	AccessToken string `json:"access_token,omitempty"`
	// The duration in seconds. that the access token is valid.
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// The user's ID token.
	IDToken string `json:"id_token,omitempty"`
	// The refresh token, only available if `offline_access` scope was provided.
	RefreshToken string `json:"refresh_token,omitempty"`
	// The type of the access token.
	TokenType string `json:"token_type,omitempty"`
}

// LoginWithPasswordRequest defines the request body for logging in with a password grant.
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
	// String value of the realm the user belongs. Set this if you want to add realm support at this grant.
	Realm string
	// Extra parameters to be merged into the request body. Values set here will override any existing values
	ExtraParameters map[string]string
}
