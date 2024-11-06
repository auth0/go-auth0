package management

import "context"

// DeviceCredential is a device credential.
type DeviceCredential struct {
	// ID of this device credential
	ID *string `json:"id,omitempty"`

	// The id of the client.
	ClientID *string `json:"client_id,omitempty"`

	// The id of the user.
	UserID *string `json:"user_id,omitempty"`

	// User agent for this device
	DeviceName *string `json:"device_name,omitempty"`

	// Unique identifier for the device. NOTE: This field is generally not populated for refresh_tokens and rotating_refresh_tokens
	DeviceID *string `json:"device_id,omitempty"`

	// Type of credential. Can be public_key, refresh_token, or rotating_refresh_token
	Type *string `json:"type,omitempty"`

	// Base64 encoded string containing the credential
	Value *string `json:"value,omitempty"`
}

// DeviceCredentialList is a list of DeviceCredentials.
type DeviceCredentialList struct {
	List
	DeviceCredentials []*DeviceCredential `json:"device_credentials"`
}

// DeviceCredentialsManager manages Auth0 device-credentials resources.
type DeviceCredentialsManager manager

// Create a device credential public key to manage refresh token rotation for a given user_id
// Type of credential must be "public_key".
//
// See: https://auth0.com/docs/api/management/v2/device-credentials/post-device-credentials
func (m *DeviceCredentialsManager) Create(ctx context.Context, d *DeviceCredential, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("device-credentials"), d, opts...)
}

// List device credential information (public_key, refresh_token, or rotating_refresh_token) associated with a specific user.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/device-credentials/get-device-credentials
func (m *DeviceCredentialsManager) List(ctx context.Context, opts ...RequestOption) (d *DeviceCredentialList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("device-credentials"), &d, applyListDefaults(opts))
	return
}

// Delete a device credential (such as a refresh token or public key) with the given ID.
//
// See: https://auth0.com/docs/api/management/v2/device-credentials/delete-device-credentials-by-id
func (m *DeviceCredentialsManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("device-credentials", id), nil, opts...)
}
