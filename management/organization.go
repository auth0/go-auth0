package management

import "context"

// Organization is used to allow B2B customers to better manage
// their partners and customers, and to customize the ways that
// end-users access their applications.
//
// See: https://auth0.com/docs/manage-users/organizations
type Organization struct {
	// Organization identifier
	ID *string `json:"id,omitempty"`

	// Name of this organization.
	Name *string `json:"name,omitempty"`

	// DisplayName of this organization.
	DisplayName *string `json:"display_name,omitempty"`

	// Branding defines how to style the login pages
	Branding *OrganizationBranding `json:"branding,omitempty"`

	// Metadata associated with the organization, in the form of an object with
	// string values (max 255 chars). Maximum of 10 metadata properties allowed.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// OrganizationBranding holds branding information for an Organization.
type OrganizationBranding struct {
	// URL of logo to display on login page.
	LogoURL *string `json:"logo_url,omitempty"`

	// Color scheme used to customize the login pages.
	Colors *map[string]string `json:"colors,omitempty"`
}

// OrganizationMember holds member information for an Organization.
type OrganizationMember struct {
	UserID  *string `json:"user_id,omitempty"`
	Picture *string `json:"picture,omitempty"`
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty"`
	// Roles is only included when the field is requested using the `IncludeFields` RequestOption.
	// All fields that are required must also be included in the `IncludeFields` call so to requests
	// all fields use `IncludeFields("user_id", "picture", "name", "email", "roles")`.
	Roles []*OrganizationMemberListRole `json:"roles,omitempty"`
}

// OrganizationMemberListRole holds member role information.
type OrganizationMemberListRole struct {
	// ID for this role.
	ID *string `json:"id,omitempty"`

	// Name of the role.
	Name *string `json:"name,omitempty"`
}

// OrganizationConnection holds connection information for an Organization.
type OrganizationConnection struct {
	// ID of the connection.
	ConnectionID *string `json:"connection_id,omitempty"`

	// When true, all users that log in with this connection will be
	// automatically granted membership in the organization. When false, users
	// must be granted membership in the organization before logging in with
	// this connection.
	AssignMembershipOnLogin *bool `json:"assign_membership_on_login,omitempty"`

	// Connection details
	Connection *OrganizationConnectionDetails `json:"connection,omitempty"`

	// Determines whether a connection should be displayed on this organization’s login prompt.
	// Only applicable for enterprise connections. Default: true.
	ShowAsButton *bool `json:"show_as_button,omitempty"`
}

// OrganizationConnectionDetails holds connection details for an Organization.
type OrganizationConnectionDetails struct {
	// The name of the enabled connection.
	Name *string `json:"name,omitempty"`

	// The strategy of the enabled connection.
	Strategy *string `json:"strategy,omitempty"`
}

// OrganizationInvitationInviter holds the name of the inviter.
type OrganizationInvitationInviter struct {
	// The inviters' name.
	Name *string `json:"name,omitempty"`
}

// OrganizationInvitationInvitee holds the name of the invitee.
type OrganizationInvitationInvitee struct {
	// The invitee's email.
	Email *string `json:"email,omitempty"`
}

// OrganizationInvitation holds information on the invitation to an Organization.
type OrganizationInvitation struct {
	// The id of the user invitation.
	ID *string `json:"id,omitempty"`

	// Organization identifier
	OrganizationID *string `json:"organization_id,omitempty"`

	Inviter *OrganizationInvitationInviter `json:"inviter,omitempty"`

	Invitee *OrganizationInvitationInvitee `json:"invitee,omitempty"`

	// The invitation url to be send to the invitee.
	InvitationURL *string `json:"invitation_url,omitempty"`

	// The ISO 8601 formatted timestamp representing the creation time of the
	// invitation.
	CreatedAt *string `json:"created_at,omitempty"`

	// Number of seconds for which the invitation is valid before expiration. If
	// unspecified or set to 0, this value defaults to 604800 seconds (7 days).
	// Max value: 2592000 seconds (30 days).
	TTLSec *int `json:"ttl_sec,omitempty"`

	// The ISO 8601 formatted timestamp representing the expiration time of the
	// invitation.
	ExpiresAt *string `json:"expires_at,omitempty"`

	// Auth0 client ID. Used to resolve the application's login initiation
	// endpoint.
	ClientID *string `json:"client_id,omitempty"`

	// The id of the connection to force invitee to authenticate with.
	ConnectionID *string `json:"connection_id,omitempty"`

	// Data related to the user that does affect the application's core
	// functionality.
	AppMetadata map[string]interface{} `json:"app_metadata,omitempty"`

	// Data related to the user that does not affect the application's core
	// functionality.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`

	// List of roles IDs to associated with the user.
	Roles []string `json:"roles,omitempty"`

	// The id of the invitation ticket
	TicketID *string `json:"ticket_id,omitempty"`

	// Whether the user will receive an invitation email (true) or no email
	// (false), true by default
	SendInvitationEmail *bool `json:"send_invitation_email,omitempty"`
}

// OrganizationMemberRole holds member role information.
type OrganizationMemberRole struct {
	// ID for this role.
	ID *string `json:"id,omitempty"`

	// Name of the role.
	Name *string `json:"name,omitempty"`

	// Description of the role.
	Description *string `json:"description,omitempty"`
}

// OrganizationMemberRoleList is a list of OrganizationMemberRoles.
type OrganizationMemberRoleList struct {
	List
	Roles []OrganizationMemberRole `json:"roles"`
}

// OrganizationInvitationList is a list of OrganizationInvitations.
type OrganizationInvitationList struct {
	List
	OrganizationInvitations []*OrganizationInvitation `json:"invitations"`
}

// OrganizationConnectionList is a list of OrganizationConnection.
type OrganizationConnectionList struct {
	List
	OrganizationConnections []*OrganizationConnection `json:"enabled_connections"`
}

// OrganizationMemberList is a list of OrganizationMembers.
type OrganizationMemberList struct {
	List
	Members []OrganizationMember `json:"members"`
}

// OrganizationList is a list of Organizations.
type OrganizationList struct {
	List
	Organizations []*Organization `json:"organizations"`
}

// OrganizationManager is used for managing an Organization.
type OrganizationManager manager

// List available organizations.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_organizations
func (m *OrganizationManager) List(ctx context.Context, opts ...RequestOption) (o *OrganizationList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations"), &o, applyListDefaults(opts))
	return
}

// Create an Organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/post_organizations
func (m *OrganizationManager) Create(ctx context.Context, o *Organization, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("organizations"), &o, opts...)
	return
}

// Get a specific organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_organizations_by_id
func (m *OrganizationManager) Read(ctx context.Context, id string, opts ...RequestOption) (o *Organization, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id), &o, opts...)
	return
}

// Delete a specific organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/delete_organizations_by_id
func (m *OrganizationManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("organizations", id), nil, opts...)
	return
}

// Update an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/patch_organizations_by_id
func (m *OrganizationManager) Update(ctx context.Context, id string, o *Organization, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PATCH", m.management.URI("organizations", id), &o, opts...)
	return
}

// ReadByName retrieves a specific organization by name.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_name_by_name
func (m *OrganizationManager) ReadByName(ctx context.Context, name string, opts ...RequestOption) (o *Organization, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", "name", name), &o, opts...)
	return
}

// Connections retrieves connections enabled for an organization.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_enabled_connections
func (m *OrganizationManager) Connections(ctx context.Context, id string, opts ...RequestOption) (c *OrganizationConnectionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "enabled_connections"), &c, applyListDefaults(opts))
	return
}

// AddConnection adds connections to an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/post_enabled_connections
func (m *OrganizationManager) AddConnection(ctx context.Context, id string, c *OrganizationConnection, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("organizations", id, "enabled_connections"), &c, opts...)
	return
}

// Connection retrieves an enabled connection for an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_enabled_connections_by_connectionId
func (m *OrganizationManager) Connection(ctx context.Context, id string, connectionID string, opts ...RequestOption) (c *OrganizationConnection, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "enabled_connections", connectionID), &c, opts...)
	return
}

// DeleteConnection deletes connections from an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/delete_enabled_connections_by_connectionId
func (m *OrganizationManager) DeleteConnection(ctx context.Context, id string, connectionID string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("organizations", id, "enabled_connections", connectionID), nil, opts...)
	return
}

// UpdateConnection updates an enabled_connection belonging to an Organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/patch_enabled_connections_by_connectionId
func (m *OrganizationManager) UpdateConnection(ctx context.Context, id string, connectionID string, c *OrganizationConnection, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PATCH", m.management.URI("organizations", id, "enabled_connections", connectionID), &c, opts...)
	return
}

// Invitations retrieves invitations to organization.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
// Note that when paginating this response the `HasNext` helper cannot be used, so instead check the length of the returned list
// manually and break when there are 0 entries. See https://github.com/auth0/go-auth0/issues/48 for more context.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_invitations
func (m *OrganizationManager) Invitations(ctx context.Context, id string, opts ...RequestOption) (i *OrganizationInvitationList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "invitations"), &i, applyListDefaults(opts))
	return
}

// CreateInvitation creates invitations to an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/post_invitations
func (m *OrganizationManager) CreateInvitation(ctx context.Context, id string, i *OrganizationInvitation, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("organizations", id, "invitations"), &i, opts...)
	return
}

// Invitation retrieves an invitation to an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_invitations_by_invitation_id
func (m *OrganizationManager) Invitation(ctx context.Context, id string, invitationID string, opts ...RequestOption) (i *OrganizationInvitation, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "invitations", invitationID), &i, opts...)
	return
}

// DeleteInvitation deletes an invitation to an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/delete_invitations_by_invitation_id
func (m *OrganizationManager) DeleteInvitation(ctx context.Context, id string, invitationID string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("organizations", id, "invitations", invitationID), nil, opts...)
	return
}

// Members lists organization members.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_members
func (m *OrganizationManager) Members(ctx context.Context, id string, opts ...RequestOption) (o *OrganizationMemberList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "members"), &o, applyListDefaults(opts))
	return
}

// AddMembers adds members to an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/post_members
func (m *OrganizationManager) AddMembers(ctx context.Context, id string, memberIDs []string, opts ...RequestOption) (err error) {
	body := struct {
		Members []string `json:"members"`
	}{
		Members: memberIDs,
	}
	err = m.management.Request(ctx, "POST", m.management.URI("organizations", id, "members"), &body, opts...)
	return
}

// DeleteMembers deletes members from an organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/delete_members
func (m *OrganizationManager) DeleteMembers(ctx context.Context, id string, memberIDs []string, opts ...RequestOption) (err error) {
	body := struct {
		Members []string `json:"members"`
	}{
		Members: memberIDs,
	}
	err = m.management.Request(ctx, "DELETE", m.management.URI("organizations", id, "members"), &body, opts...)
	return
}

// MemberRoles retrieves the roles assigned to an organization member.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/get_organization_member_roles
func (m *OrganizationManager) MemberRoles(ctx context.Context, id string, memberID string, opts ...RequestOption) (r *OrganizationMemberRoleList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "members", memberID, "roles"), &r, applyListDefaults(opts))
	return
}

// AssignMemberRoles assigns one or more roles to a given user that will be applied in the context of the provided organization
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/post_organization_member_roles
func (m *OrganizationManager) AssignMemberRoles(ctx context.Context, id string, memberID string, roles []string, opts ...RequestOption) (err error) {
	body := struct {
		Roles []string `json:"roles"`
	}{
		Roles: roles,
	}
	err = m.management.Request(ctx, "POST", m.management.URI("organizations", id, "members", memberID, "roles"), &body, opts...)
	return
}

// DeleteMemberRoles removes one or more roles from a given user in the context of the provided organization.
//
// See: https://auth0.com/docs/api/management/v2/#!/Organizations/delete_organization_member_roles
func (m *OrganizationManager) DeleteMemberRoles(ctx context.Context, id string, memberID string, roles []string, opts ...RequestOption) (err error) {
	body := struct {
		Roles []string `json:"roles"`
	}{
		Roles: roles,
	}
	err = m.management.Request(ctx, "DELETE", m.management.URI("organizations", id, "members", memberID, "roles"), &body, opts...)
	return
}

// ClientGrants retrieves the client grants assigned to an organization.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
func (m *OrganizationManager) ClientGrants(ctx context.Context, id string, opts ...RequestOption) (g *ClientGrantList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("organizations", id, "client-grants"), &g, applyListDefaults(opts))
	return
}

// AssociateClientGrant assigns a client grant to an organization.
func (m *OrganizationManager) AssociateClientGrant(ctx context.Context, id string, grantID string, opts ...RequestOption) (err error) {
	body := struct {
		GrantID string `json:"grant_id"`
	}{
		GrantID: grantID,
	}
	err = m.management.Request(ctx, "POST", m.management.URI("organizations", id, "client-grants"), &body, opts...)
	return
}

// RemoveClientGrant removes a client grant from an organization.
func (m *OrganizationManager) RemoveClientGrant(ctx context.Context, id string, grantID string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("organizations", id, "client-grants", grantID), nil, opts...)
	return
}
