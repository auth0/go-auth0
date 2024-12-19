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

// DeletePermissions Remove permissions from a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/delete_role_permission_assignment
func (m *Manager) DeletePermissions(ctx context.Context, id string, postRolePermissionAssignmentRequest *models.PostRolePermissionAssignmentRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("roles", string(id), "permissions"), postRolePermissionAssignmentRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Delete Delete a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/delete_roles_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("roles", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetPermissions Get permissions granted by role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_role_permission
func (m *Manager) GetPermissions(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetRolePermission200Response, error) {
	var localVarReturnValue *models.GetRolePermission200Response
	err := m.management.Request(ctx, "GET", m.management.URI("roles", string(id), "permissions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUsers Get a role&#39;s users
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_role_user
func (m *Manager) GetUsers(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetRoleUser200Response, error) {
	var localVarReturnValue *models.GetRoleUser200Response
	err := m.management.Request(ctx, "GET", m.management.URI("roles", string(id), "users"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetAll Get roles
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_roles
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200Response, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200Response
	err := m.management.Request(ctx, "GET", m.management.URI("roles"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Get Get a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/get_roles_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("roles", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Update Update a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/patch_roles_by_id
func (m *Manager) Update(ctx context.Context, id string, roleUpdate *models.RoleUpdate, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200ResponseOneOfInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("roles", string(id)), roleUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// AddPermissions Associate permissions with a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/post_role_permission_assignment
func (m *Manager) AddPermissions(ctx context.Context, id string, postRolePermissionAssignmentRequest *models.PostRolePermissionAssignmentRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("roles", string(id), "permissions"), postRolePermissionAssignmentRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AssignUsers Assign users to a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/post_role_users
func (m *Manager) AssignUsers(ctx context.Context, id string, postRoleUsersRequest *models.PostRoleUsersRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("roles", string(id), "users"), postRoleUsersRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Create Create a role
//
// https://auth0.com/docs/api/management/v2/#!/Roles/post_roles
func (m *Manager) Create(ctx context.Context, roleCreate *models.RoleCreate, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200ResponseOneOfInner
	err := m.management.Request(ctx, "POST", m.management.URI("roles"), roleCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
