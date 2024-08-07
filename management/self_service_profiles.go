package management

import (
	"context"
	"time"
)

// SelfServiceProfile is used to allow B2B customers to better manage
// their SSO Profiles
//
// See: https://auth0.com/docs/self-service-profile
type SelfServiceProfile struct {
	ID             *string           `json:"id,omitempty"`
	SAMLMappings   []*SAMLMappings   `json:"saml_mappings,omitempty"`
	OIDCScopes     []*string         `json:"oidc_scopes,omitempty"`
	CreatedAt      *time.Time        `json:"created_at,omitempty"`
	UpdatedAt      *time.Time        `json:"updated_at,omitempty"`
	Branding       *Branding         `json:"branding,omitempty"`
	UserAttributes []*UserAttributes `json:"user_attributes,omitempty"`
}

// UserAttributes is used to determine optional attributes.
// Only one between UserAttributes or (SAMLMappings,OIDCScopes) can be passed.
type UserAttributes struct {
	Name            *string          `json:"name,omitempty"`
	DisplayName     *string          `json:"display_name,omitempty"`
	Description     *string          `json:"description,omitempty"`
	IsOptional      *bool            `json:"is_optional,omitempty"`
	DefaultMappings *DefaultMappings `json:"default_mappings,omitempty"`
}

// SAMLMappings is used to determine saml settings.
type SAMLMappings struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsOptional  *bool   `json:"is_optional,omitempty"`
}

type DefaultMappings struct {
	Saml *string `json:"saml,omitempty"`
	Oidc *string `json:"oidc,omitempty"`
}

// SSOManager manages Auth0 Self Service Profile resources.
type SSOManager manager

// Create a new Self Service Profile.
func (m *SSOManager) Create(ctx context.Context, s *SelfServiceProfile, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("self-service-profiles"), &s, opts...)
	return
}

// List all Self Service Profiles.
func (m *SSOManager) List(ctx context.Context, opts ...RequestOption) (s []*SelfServiceProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles"), &s, applyListDefaults(opts))
	return
}

// Get a single Self Service Profile against the ID.
func (m *SSOManager) Read(ctx context.Context, id string, opts ...RequestOption) (s *SelfServiceProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles", id), &s, opts...)
	return
}

// Update an existing Self Service Profile against the ID.
func (m *SSOManager) Update(ctx context.Context, id string, s *SelfServiceProfile, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("self-service-profiles", id), s, opts...)
}

// Delete a Self Service Profile against the ID.
func (m *SSOManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("self-service-profiles", id), nil, opts...)
}
