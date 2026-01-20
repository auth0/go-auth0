package management

import "context"

// ClientGrant is a method through which applications can gain Access Tokens.
//
// See: https://auth0.com/docs/get-started/applications/application-grant-types
type ClientGrant struct {
	// A generated string identifying the client grant.
	ID *string `json:"id,omitempty"`

	// The identifier of the client.
	ClientID *string `json:"client_id,omitempty"`

	// The audience.
	Audience *string `json:"audience,omitempty"`

	// The list of permissions (scopes) that are granted to the client.
	Scope *[]string `json:"scope,omitempty"`

	//When enabled, all scopes configured on the resource server are allowed for by this client grant.
	//Scope param not accepted in request/returned in response when allow_all_scopes is true
	AllowAllScopes *bool `json:"allow_all_scopes,omitempty"`

	// If enabled, any organization can be used with this grant.
	// If disabled (default), the grant must be explicitly assigned to the desired organizations.
	AllowAnyOrganization *bool `json:"allow_any_organization,omitempty"`

	// Defines whether organizations can be used with client credentials exchanges for this grant.
	// Can be one of `deny`, `allow`, or `require`. Defaults to `deny` when not defined.
	OrganizationUsage *string `json:"organization_usage,omitempty"`

	// SubjectType defines the type of subject for this grant.
	// Can be one of `client` or `user`. Defaults to `client` when not defined.
	SubjectType *string `json:"subject_type,omitempty"`

	// AuthorizationDetailsTypes defines the types of authorization details allowed for this client grant.
	AuthorizationDetailsTypes *[]string `json:"authorization_details_types,omitempty"`

	// IsSystem indicates whether this grant is a special grant created by Auth0. It cannot be modified or deleted directly.
	IsSystem *bool `json:"is_system,omitempty"`
}

// ClientGrantList is a list of ClientGrants.
type ClientGrantList struct {
	List
	ClientGrants []*ClientGrant `json:"client_grants"`
}

// ClientGrantManager manages Auth0 ClientGrant resources.
type ClientGrantManager manager

// Create a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/post_client_grants
func (m *ClientGrantManager) Create(ctx context.Context, g *ClientGrant, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("client-grants"), g, opts...)
}

// Read a client grant by its ID.
func (m *ClientGrantManager) Read(ctx context.Context, id string, opts ...RequestOption) (cg *ClientGrant, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("client-grants", id), &cg, opts...)
	return
}

// Update a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/patch_client_grants_by_id
func (m *ClientGrantManager) Update(ctx context.Context, id string, g *ClientGrant, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("client-grants", id), g, opts...)
}

// Delete a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/delete_client_grants_by_id
func (m *ClientGrantManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("client-grants", id), nil, opts...)
}

// List client grants.
//
// This method forces the `include_totals=true` and defaults to `per_page=50` if
// not provided.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/get_client_grants
func (m *ClientGrantManager) List(ctx context.Context, opts ...RequestOption) (gs *ClientGrantList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("client-grants"), &gs, applyListDefaults(opts))
	return
}

// Organizations lists client grants associated to an organization.
//
// This method forces the `include_totals=true` and defaults to `per_page=50` if
// not provided.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
func (m *ClientGrantManager) Organizations(ctx context.Context, id string, opts ...RequestOption) (o *OrganizationList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("client-grants", id, "organizations"), &o, applyListDefaults(opts))
	return
}
