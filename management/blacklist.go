package management

import "context"

// BlacklistToken is a token that has been blacklisted.
type BlacklistToken struct {
	// The "aud" (audience) claim identifies the
	// recipients that the JWT is intended for.
	//
	// See: https://tools.ietf.org/html/rfc7519#section-4.1.3
	Audience string `json:"aud,omitempty"`

	// The "jti" (JWT ID) claim provides a unique
	// (within "aud") identifier for the JWT.
	//
	// See: https://tools.ietf.org/html/rfc7519#section-4.1.7
	JTI string `json:"jti,omitempty"`
}

// BlacklistManager manages Auth0 BlacklistToken resources.
type BlacklistManager manager

// List tokens that are blacklisted.
//
// Note: The JWT specification states that the `jti` field can be used to
// prevent replay attacks. Though Auth0 tokens do not include a `jti`, you can
// nevertheless blacklist a `jti` to prevent a token being used more than a
// predetermined number of times. This behavior is similar to implementing a
// nonce (where the token's signature can be thought of as the nonce). If a
// token gets stolen, it (or the tokens issued after it) should be blacklisted
// and let expire.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
// See: https://auth0.com/docs/api/management/v2#!/Blacklists/get_tokens
func (m *BlacklistManager) List(ctx context.Context, opts ...RequestOption) (bl []*BlacklistToken, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("blacklists", "tokens"), &bl, applyListDefaults(opts))
	return
}

// Create a blacklist for a token.
//
// See: https://auth0.com/docs/api/management/v2#!/Blacklists/post_tokens
func (m *BlacklistManager) Create(ctx context.Context, t *BlacklistToken, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("blacklists", "tokens"), t, opts...)
}
