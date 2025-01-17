package management

import (
	"context"
	"time"
)

// RefreshTokenList represents a list of user refresh tokens.
type RefreshTokenList struct {
	List
	Tokens []*RefreshToken `json:"tokens,omitempty"`
}

// RefreshToken represents a refresh token.
type RefreshToken struct {
	// ID of the refresh token
	ID *string `json:"id,omitempty"`
	// ID of the user which can be used when interacting with other APIs.
	UserID *string `json:"user_id,omitempty"`
	// The date and time when the refresh token was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when the refresh token was last used
	IdleExpiresAt *time.Time `json:"idle_expires_at,omitempty"`
	// The date and time when the refresh token will expire
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Device used while issuing/exchanging the refresh token
	Device *RefreshTokenDevice `json:"device,omitempty"`
	// ID of the client application granted with this refresh token
	ClientID *string `json:"client_id,omitempty"`
	// ID of the authenticated session used to obtain this refresh-token
	SessionID *string `json:"session_id,omitempty"`
	// True if the token is a rotating refresh token
	Rotating *bool `json:"rotating,omitempty"`
	// A list of the resource server IDs associated to this refresh-token and their granted scopes
	ResourceServer []*RefreshTokenResourceServer `json:"resource_servers,omitempty"`
	// The date and time when the refresh token was last exchanged
	LastExchangedAt *string `json:"last_exchanged_at,omitempty"`
}

// RefreshTokenDevice represents the device associated with a refresh token.
type RefreshTokenDevice struct {
	// First IP address associated with the refresh token
	InitialIP *string `json:"initial_ip,omitempty"`
	// First autonomous system number associated with the refresh token
	InitialASN *string `json:"initial_asn,omitempty"`
	// First user agent associated with the refresh token
	InitialUserAgent *string `json:"initial_user_agent,omitempty"`
	// Last IP address associated with the refresh token
	LastIP *string `json:"last_ip,omitempty"`
	// Last autonomous system number associated with the refresh token
	LastASN *string `json:"last_asn,omitempty"`
	// Last user agent associated with the refresh token
	LastUserAgent *string `json:"last_user_agent,omitempty"`
}

// RefreshTokenResourceServer represents the resource server associated with a refresh token.
type RefreshTokenResourceServer struct {
	Audience *string `json:"audience,omitempty"`
	Scopes   *string `json:"scopes,omitempty"`
}

// RefreshTokenManager manages refresh tokens using the Management API.
type RefreshTokenManager manager

// Read retrieves a refresh token by its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Refresh_Tokens/get-refresh-token
func (m *RefreshTokenManager) Read(ctx context.Context, refreshTokenID string, opts ...RequestOption) (r *RefreshToken, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("refresh-tokens", refreshTokenID), &r, opts...)
	return
}

// Delete removes a refresh token by its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Refresh_Tokens/delete_refresh_token
func (m *RefreshTokenManager) Delete(ctx context.Context, refreshTokenID string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("refresh-tokens", refreshTokenID), nil, opts...)
}
