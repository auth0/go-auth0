package management

import "context"

// Role is used to assign roles to a User.
type Role struct {
	// A unique ID for the role.
	ID *string `json:"id,omitempty"`

	// The name of the role created.
	Name *string `json:"name,omitempty"`

	// A description of the role created.
	Description *string `json:"description,omitempty"`
}

// RoleList holds a list of Roles.
type RoleList struct {
	List
	Roles []*Role `json:"roles"`
}

// Permission is granted to a Role.
type Permission struct {
	// The resource server that the permission is attached to.
	ResourceServerIdentifier *string `json:"resource_server_identifier,omitempty"`

	// The name of the resource server.
	ResourceServerName *string `json:"resource_server_name,omitempty"`

	// The name of the permission.
	Name *string `json:"permission_name,omitempty"`

	// The description of the permission.
	Description *string `json:"description,omitempty"`
}

// PermissionList holds a list of Permissions.
type PermissionList struct {
	List
	Permissions []*Permission `json:"permissions"`
}

// RoleManager manages Auth0 Role resources.
type RoleManager manager

// Create a new role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/post_roles
func (m *RoleManager) Create(ctx context.Context, r *Role, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("roles"), r, opts...)
}

// Retrieve a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_roles_by_id
func (m *RoleManager) Read(ctx context.Context, id string, opts ...RequestOption) (r *Role, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("roles", id), &r, opts...)
	return
}

// Update a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/patch_roles_by_id
func (m *RoleManager) Update(ctx context.Context, id string, r *Role, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("roles", id), r, opts...)
}

// Delete a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/delete_roles_by_id
func (m *RoleManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	// Deleting a role results in a 200 status code instead of 204 which
	// triggers decoding of the response payload.
	//
	// In order to avoid Unmarshal(nil) errors, we pass an empty &Role{}.
	return m.management.Request(ctx, "DELETE", m.management.URI("roles", id), &Role{}, opts...)
}

// List roles that can be assigned to users or groups.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_roles
func (m *RoleManager) List(ctx context.Context, opts ...RequestOption) (r *RoleList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("roles"), &r, applyListDefaults(opts))
	return
}

// AssignUsers assigns users to a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/post_role_users
func (m *RoleManager) AssignUsers(ctx context.Context, id string, users []*User, opts ...RequestOption) error {
	u := make(map[string][]*string)
	u["users"] = make([]*string, len(users))
	for i, user := range users {
		u["users"][i] = user.ID
	}
	return m.management.Request(ctx, "POST", m.management.URI("roles", id, "users"), &u, opts...)
}

// Users retrieves users associated with a role.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_role_user
func (m *RoleManager) Users(ctx context.Context, id string, opts ...RequestOption) (u *UserList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("roles", id, "users"), &u, applyListDefaults(opts))
	return
}

// AssociatePermissions associates permissions to a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/post_role_permission_assignment
func (m *RoleManager) AssociatePermissions(ctx context.Context, id string, permissions []*Permission, opts ...RequestOption) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return m.management.Request(ctx, "POST", m.management.URI("roles", id, "permissions"), &p, opts...)
}

// Permissions retrieves permissions granted by a role.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_role_permission
func (m *RoleManager) Permissions(ctx context.Context, id string, opts ...RequestOption) (p *PermissionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("roles", id, "permissions"), &p, applyListDefaults(opts))
	return
}

// RemovePermissions removes permissions associated to a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/delete_role_permission_assignment
func (m *RoleManager) RemovePermissions(ctx context.Context, id string, permissions []*Permission, opts ...RequestOption) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return m.management.Request(ctx, "DELETE", m.management.URI("roles", id, "permissions"), &p, opts...)
}
