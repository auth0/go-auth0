package management

import "context"

// Grant is a way of retrieving an Access Token.
//
// See: https://auth0.com/docs/get-started/authentication-and-authorization-flow/which-oauth-2-0-flow-should-i-use
type Grant struct {
	// The id of the grant.
	ID *string `json:"id,omitempty"`

	// The id of the client.
	ClientID *string `json:"clientID,omitempty"`

	// The id of the user.
	UserID *string `json:"user_id"`

	// The grant's audience.
	Audience *string `json:"audience,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`
}

// GrantList is a list of Grants.
type GrantList struct {
	List
	Grants []*Grant `json:"grants"`
}

// GrantManager manages Auth0 Grant resources.
type GrantManager manager

// List the grants associated with your account.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Grants/get_grants
func (m *GrantManager) List(ctx context.Context, opts ...RequestOption) (g *GrantList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("grants"), &g, applyListDefaults(opts))
	return
}

// Delete revokes a grant associated with a user-id.
//
// https://auth0.com/docs/api/management/v2#!/Grants/delete_grants_by_id
func (m *GrantManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("grants", id), nil, opts...)
}
