package management

import (
	"context"
	"time"
)

// EncryptionKeyList is a list of encryption keys.
type EncryptionKeyList struct {
	List
	Keys []*EncryptionKey `json:"keys"`
}

// EncryptionKey is used for encrypting data.
type EncryptionKey struct {
	// Key ID
	KID *string `json:"kid,omitempty"`
	// Key type
	Type *string `json:"type,omitempty"`
	// Key state
	State *string `json:"state,omitempty"`
	// Key creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Key update timestamp
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of parent wrapping key
	ParentKID *string `json:"parent_kid,omitempty"`
	// Base64 encoded ciphertext of key material wrapped by public wrapping key
	WrappedKey *string `json:"wrapped_key,omitempty"`
}

// reset cleans up unnecessary fields based on the operation type.
func (k *EncryptionKey) reset(op string) {
	k.KID = nil
	k.CreatedAt = nil
	k.UpdatedAt = nil
	k.ParentKID = nil
	k.State = nil

	switch op {
	case "import":
		k.Type = nil
	case "create":
		k.WrappedKey = nil
	default:
		k.Type = nil
		k.WrappedKey = nil
	}
}

// WrappingKey is used for creating the public wrapping key.
type WrappingKey struct {
	// The public key of the wrapping key for uploading the customer provided root key.
	PublicKey *string `json:"public_key,omitempty"`
	// The algorithm to be used for wrapping the key. Normally CKM_RSA_AES_KEY_WRAP
	Algorithm *string `json:"algorithm,omitempty"`
}

// EncryptionKeyManager manages Auth0 EncryptionKey resources.
type EncryptionKeyManager manager

// Create an encryption key.
//
// See: https://auth0.com/docs/api/management/v2/keys/post-encryption
func (m *EncryptionKeyManager) Create(ctx context.Context, e *EncryptionKey, opts ...RequestOption) error {
	e.reset("create")
	return m.management.Request(ctx, "POST", m.management.URI("keys", "encryption"), e, opts...)
}

// List all encryption keys.
//
// See: https://auth0.com/docs/api/management/v2/keys/get-encryption-keys
func (m *EncryptionKeyManager) List(ctx context.Context, opts ...RequestOption) (ekl *EncryptionKeyList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("keys", "encryption"), &ekl, applyListDefaults(opts))
	return
}

// Read an encryption key by its key id.
//
// See: https://auth0.com/docs/api/management/v2/keys/get-encryption-key
func (m *EncryptionKeyManager) Read(ctx context.Context, kid string, opts ...RequestOption) (k *EncryptionKey, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("keys", "encryption", kid), &k, opts...)
	return
}

// Rekey the key hierarchy, Performs rekeying operation on the key hierarchy.
//
// See: https://auth0.com/docs/api/management/v2/keys/post-encryption-rekey
func (m *EncryptionKeyManager) Rekey(ctx context.Context, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("keys", "encryption", "rekey"), nil, opts...)
}

// Delete an encryption key by its key id.
//
// See: https://auth0.com/docs/api/management/v2/keys/delete-encryption-key
func (m *EncryptionKeyManager) Delete(ctx context.Context, kid string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("keys", "encryption", kid), nil, opts...)
}

// ImportWrappedKey Imports wrapped key material and activate encryption key
//
// See: https://auth0.com/docs/api/management/v2/keys/post-encryption-key
func (m *EncryptionKeyManager) ImportWrappedKey(ctx context.Context, e *EncryptionKey, opts ...RequestOption) error {
	id := *e.KID
	e.reset("import")

	return m.management.Request(ctx, "POST", m.management.URI("keys", "encryption", id), e, opts...)
}

// CreatePublicWrappingKey creates the public wrapping key to wrap your own encryption key material.
//
// See: https://auth0.com/docs/api/management/v2/keys/post-encryption-wrapping-key
func (m *EncryptionKeyManager) CreatePublicWrappingKey(ctx context.Context, kid string, opts ...RequestOption) (w *WrappingKey, err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("keys", "encryption", kid, "wrapping-key"), &w, opts...)
	return
}
