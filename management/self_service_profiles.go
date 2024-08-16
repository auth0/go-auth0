package management

import (
	"context"
	"encoding/json"
	"time"
)

// SelfServiceProfile allows to configure SSO so that
// customers can independently set up
// SSO and sign in to your application.
type SelfServiceProfile struct {
	ID *string `json:"id,omitempty"`

	// List of attributes to be mapped that
	// will be shown to the user during the SS-SSO flow.
	UserAttributes []*UserAttributes `json:"user_attributes,omitempty"`
	CreatedAt      *time.Time        `json:"created_at,omitempty"`
	UpdatedAt      *time.Time        `json:"updated_at,omitempty"`

	// Branding scheme for the profile.
	Branding *Branding `json:"branding,omitempty"`
}

// UserAttributes is used to determine optional attributes.
type UserAttributes struct {
	// Identifier of this attribute.
	Name *string `json:"name"`

	// Description of this attribute.
	Description *string `json:"description"`

	// Determines if this attribute is required.
	IsOptional *bool `json:"is_optional"`
}

// SSOTicket is used to created self-service ticket for a set of clients and organizations.
type SSOTicket struct {
	// If provided, this will allow editing of the
	// provided connection during the SSO Flow.
	ConnectionID *string `json:"connection_id,omitempty"`

	// If provided, this will create a new connection
	// for the SSO flow with the given configuration.
	ConnectionConfig *ConnectionConfig `json:"connection_config,omitempty"`

	// List of client_ids that the
	// connection will be enabled for.
	EnabledClients []*string `json:"enabled_clients,omitempty"`

	// List of organizations that the
	// connection will be enabled for.
	EnabledOrganizations []*EnabledOrganizations `json:"enabled_organizations,omitempty"`

	// The ticket that is generated.
	Ticket *string `json:"ticket,omitempty"`
}

// ConnectionConfig sets the configuration for SSOTicket.
type ConnectionConfig struct {
	// The name of the connection that will be
	// created as a part of the SSO flow.
	Name string `json:"name,omitempty"`
}

// EnabledOrganizations is the list of Organizations associated with the SSO Ticket.
type EnabledOrganizations struct {
	// Organization identifier.
	OrganizationID string `json:"organization_id,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface.
func (ssp *SelfServiceProfile) MarshalJSON() ([]byte, error) {
	type SelfServiceProfileSubset struct {
		UserAttributes []*UserAttributes `json:"user_attributes,omitempty"`
		Branding       *Branding         `json:"branding,omitempty"`
	}

	return json.Marshal(&SelfServiceProfileSubset{
		UserAttributes: ssp.UserAttributes,
		Branding:       ssp.Branding,
	})
}

// SelfServiceProfileManager manages Auth0 Self Service Profile resources.
type SelfServiceProfileManager manager

// Create a new Self Service Profile.
func (m *SelfServiceProfileManager) Create(ctx context.Context, s *SelfServiceProfile, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("self-service-profiles"), &s, opts...)
	return
}

// List all Self Service Profiles.
func (m *SelfServiceProfileManager) List(ctx context.Context, opts ...RequestOption) (s []*SelfServiceProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles"), &s, applyListDefaults(opts))
	return
}

// Get a single Self Service Profile against the ID.
func (m *SelfServiceProfileManager) Read(ctx context.Context, id string, opts ...RequestOption) (s *SelfServiceProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("self-service-profiles", id), &s, opts...)
	return
}

// Update an existing Self Service Profile against the ID.
func (m *SelfServiceProfileManager) Update(ctx context.Context, id string, s *SelfServiceProfile, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("self-service-profiles", id), s, opts...)
}

// Delete a Self Service Profile against the ID.
func (m *SelfServiceProfileManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("self-service-profiles", id), nil, opts...)
}

// CreateTicket creates a sso-access ticket to initiate the Self Service SSO Flow.
func (m *SelfServiceProfileManager) CreateTicket(ctx context.Context, id string, t *SSOTicket, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("self-service-profiles", id, "sso-ticket"), t, opts...)
	return
}
