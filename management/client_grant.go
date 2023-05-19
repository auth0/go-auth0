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

	Scope []string `json:"scope"`
}

// ClientGrantList is a list of ClientGrants.
type ClientGrantList struct {
	List
	ClientGrants []*ClientGrant `json:"client_grants"`
}

// ClientGrantManager manages Auth0 ClientGrant resources.
type ClientGrantManager struct {
	*Management
}

func newClientGrantManager(m *Management) *ClientGrantManager {
	return &ClientGrantManager{m}
}

// Create a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/post_client_grants
func (m *ClientGrantManager) Create(ctx context.Context, g *ClientGrant, opts ...RequestOption) (err error) {
	return m.Request(ctx, "POST", m.URI("client-grants"), g, opts...)
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
	return m.Request(ctx, "PATCH", m.URI("client-grants", id), g, opts...)
}

// Delete a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/delete_client_grants_by_id
func (m *ClientGrantManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.Request(ctx, "DELETE", m.URI("client-grants", id), nil, opts...)
}

// List all client grants.
//
// This method forces the `include_totals=true` and defaults to `per_page=50` if
// not provided.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/get_client_grants
func (m *ClientGrantManager) List(ctx context.Context, opts ...RequestOption) (gs *ClientGrantList, err error) {
	err = m.Request(ctx, "GET", m.URI("client-grants"), &gs, applyListDefaults(opts))
	return
}
