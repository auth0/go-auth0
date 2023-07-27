package management

import (
	"context"
	"time"
)

// SigningKey is used for signing tokens.
type SigningKey struct {
	// The key id of the signing key.
	KID *string `json:"kid,omitempty"`

	// The public certificate of the signing key.
	Cert *string `json:"cert,omitempty"`

	// The public certificate of the signing key in pkcs7 format.
	PKCS7 *string `json:"pkcs7,omitempty"`

	// True if the key is the the current key.
	Current *bool `json:"current,omitempty"`

	// True if the key is the the next key.
	Next *bool `json:"next,omitempty"`

	// True if the key is the the previous key.
	Previous *bool `json:"previous,omitempty"`

	// The date and time when the key became the current key.
	CurrentSince *time.Time `json:"current_since,omitempty"`

	// The date and time when the current key was rotated.
	CurrentUntil *time.Time `json:"current_until,omitempty"`

	// The cert fingerprint.
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The cert thumbprint.
	Thumbprint *string `json:"thumbprint,omitempty"`

	// True if the key is revoked.
	Revoked *bool `json:"revoked,omitempty"`

	// The date and time when the key was revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// SigningKeyManager manages Auth0 SigningKey resources.
type SigningKeyManager manager

// List all Application Signing Keys.
//
// See: https://auth0.com/docs/api/management/v2#!/Keys/get_signing_keys
func (m *SigningKeyManager) List(ctx context.Context, opts ...RequestOption) (ks []*SigningKey, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("keys", "signing"), &ks, opts...)
	return
}

// Read an Application Signing Key by its key id.
//
// See: https://auth0.com/docs/api/management/v2#!/Keys/get_signing_key
func (m *SigningKeyManager) Read(ctx context.Context, kid string, opts ...RequestOption) (k *SigningKey, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("keys", "signing", kid), &k, opts...)
	return
}

// Rotate the Application Signing Key.
//
// See: https://auth0.com/docs/api/management/v2#!/Keys/post_signing_keys
func (m *SigningKeyManager) Rotate(ctx context.Context, opts ...RequestOption) (k *SigningKey, err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("keys", "signing", "rotate"), &k, opts...)
	return
}

// Revoke an Application Signing Key by its key id.
//
// See: https://auth0.com/docs/api/management/v2#!/Keys/put_signing_keys
func (m *SigningKeyManager) Revoke(ctx context.Context, kid string, opts ...RequestOption) (k *SigningKey, err error) {
	err = m.management.Request(ctx, "PUT", m.management.URI("keys", "signing", kid, "revoke"), &k, opts...)
	return
}
