package roles

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/auth0/go-auth0/management/models"
)

// Manager defines the base manager interface
type Manager struct {
	management *management.Management
}

// NewManager creates a new manager for  operations
func NewManager(mgmt *management.Management) *Manager {
	return &Manager{
		management: mgmt,
	}
}

// DeleteRolePermissionAssignment Remove permissions from a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/delete_role_permission_assignment
func (m *Manager) DeleteRolePermissionAssignment(ctx context.Context, id string, postRolePermissionAssignmentRequest *models.PostRolePermissionAssignmentRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("roles", management.SafeString(id), "permissions"), postRolePermissionAssignmentRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRolesById Delete a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/delete_roles_by_id
func (m *Manager) DeleteRolesById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("roles", management.SafeString(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetRolePermission Get permissions granted by role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_role_permission
func (m *Manager) GetRolePermission(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetRolePermission200Response, error) {
	var localVarReturnValue *models.GetRolePermission200Response
	err := m.management.Request(ctx, "GET", m.management.URI("roles", management.SafeString(id), "permissions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetRoleUser Get a role&#39;s users
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_role_user
func (m *Manager) GetRoleUser(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetRoleUser200Response, error) {
	var localVarReturnValue *models.GetRoleUser200Response
	err := m.management.Request(ctx, "GET", m.management.URI("roles", management.SafeString(id), "users"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetRoles Get roles
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_roles
func (m *Manager) GetRoles(ctx context.Context, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200Response, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200Response
	err := m.management.Request(ctx, "GET", m.management.URI("roles"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetRolesById Get a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_roles_by_id
func (m *Manager) GetRolesById(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("roles", management.SafeString(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchRolesById Update a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/patch_roles_by_id
func (m *Manager) PatchRolesById(ctx context.Context, id string, roleUpdate *models.RoleUpdate, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200ResponseOneOfInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("roles", management.SafeString(id)), roleUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostRolePermissionAssignment Associate permissions with a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/post_role_permission_assignment
func (m *Manager) PostRolePermissionAssignment(ctx context.Context, id string, postRolePermissionAssignmentRequest *models.PostRolePermissionAssignmentRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("roles", management.SafeString(id), "permissions"), postRolePermissionAssignmentRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostRoleUsers Assign users to a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/post_role_users
func (m *Manager) PostRoleUsers(ctx context.Context, id string, postRoleUsersRequest *models.PostRoleUsersRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("roles", management.SafeString(id), "users"), postRoleUsersRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostRoles Create a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/post_roles
func (m *Manager) PostRoles(ctx context.Context, roleCreate *models.RoleCreate, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200ResponseOneOfInner
	err := m.management.Request(ctx, "POST", m.management.URI("roles"), roleCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
