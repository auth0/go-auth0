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

	// If enabled, any organization can be used with this grant.
	// If disabled (default), the grant must be explicitly assigned to the desired organizations.
	AllowAnyOrganization *bool `json:"allow_any_organization,omitempty"`

	// Defines whether organizations can be used with client credentials exchanges for this grant.
	// Can be one of `deny`, `allow`, or `require`. Defaults to `deny` when not defined.
	OrganizationUsage *string `json:"organization_usage,omitempty"`
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
//
// The Auth0 Management API does not offer a method to retrieve a client grant
// by id, we fake this by listing all client grants and matching by id on the
// client side. For this reason this method should be used with caution.
func (m *ClientGrantManager) Read(ctx context.Context, id string, opts ...RequestOption) (*ClientGrant, error) {
	var page int
	for {
		l, err := m.List(ctx, append(opts, Page(page))...)
		if err != nil {
			return nil, err
		}
		for _, g := range l.ClientGrants {
			if g.GetID() == id {
				return g, nil
			}
		}
		if !l.HasNext() {
			break
		}
		page++
	}
	return nil, &managementError{
		StatusCode: 404,
		Err:        "Not Found",
		Message:    "Client grant not found",
	}
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
