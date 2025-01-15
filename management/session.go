package management

import "context"

// Session represents a session Struct.
type Session struct {
	ID               *string                `json:"id,omitempty"`
	UserID           *string                `json:"user_id,omitempty"`
	CreatedAt        *string                `json:"created_at,omitempty"`
	UpdatedAt        *string                `json:"updated_at,omitempty"`
	AuthenticatedAt  *string                `json:"authenticated_at,omitempty"`
	IdleExpiresAt    *string                `json:"idle_expires_at,omitempty"`
	ExpiresAt        *string                `json:"expires_at,omitempty"`
	LastInteractedAt *string                `json:"last_interacted_at,omitempty"`
	Device           *SessionDevice         `json:"device,omitempty"`
	Clients          []*SessionClient       `json:"clients,omitempty"`
	Authentication   *SessionAuthentication `json:"authentication,omitempty"`
}

// SessionDevice represents the device signal details.
type SessionDevice struct {
	InitialUserAgent *string `json:"initial_user_agent,omitempty"`
	InitialIP        *string `json:"initial_ip,omitempty"`
	InitialASN       *string `json:"initial_asn,omitempty"`
	LastUserAgent    *string `json:"last_user_agent,omitempty"`
	LastIP           *string `json:"last_ip,omitempty"`
	LastASN          *string `json:"last_asn,omitempty"`
}

// SessionClient represents the client signal details.
type SessionClient struct {
	ClientID *string `json:"client_id,omitempty"`
}

// SessionAuthentication represents the authentication signal details.
type SessionAuthentication struct {
	Methods []*SessionAuthenticationMethod `json:"methods,omitempty"`
}

// SessionAuthenticationMethod represents the authentication signal details.
type SessionAuthenticationMethod struct {
	Name      *string `json:"name,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Type      *string `json:"type,omitempty"`
}

// SessionManager manages sessions using the Management API.
type SessionManager manager

// Read retrieves a session by its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Sessions/get_session
func (m *SessionManager) Read(ctx context.Context, sessionID string, opts ...RequestOption) (r *Session, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("sessions", sessionID), &r, opts...)
	return
}

// Delete removes a session by its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Sessions/delete-session
func (m *SessionManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("sessions", id), nil, opts...)
}

// Revoke revokes a session by its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Sessions/revoke-session
func (m *SessionManager) Revoke(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("sessions", id, "revoke"), nil, opts...)
}
