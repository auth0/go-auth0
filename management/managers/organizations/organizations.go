package organizations

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

// DeleteEnabledConnectionsByConnectionId Delete connections from an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/delete_enabled_connections_by_connection_id
func (m *Manager) DeleteEnabledConnection(ctx context.Context, id string, connectionId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("organizations", string(id), "enabled_connections", string(connectionId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteInvitationsByInvitationId Delete an invitation to an Organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/delete_invitations_by_invitation_id
func (m *Manager) DeleteInvitation(ctx context.Context, id string, invitationId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("organizations", string(id), "invitations", string(invitationId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMembers Delete members from an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/delete_members
func (m *Manager) DeleteMembers(ctx context.Context, id string, deleteMembersRequest *models.DeleteMembersRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("organizations", string(id), "members"), deleteMembersRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrganizationMemberRoles Delete user roles from an Organization member
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/delete_organization_member_roles
func (m *Manager) DeleteMemberRoles(ctx context.Context, id string, userId string, deleteOrganizationMemberRolesRequest *models.DeleteOrganizationMemberRolesRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("organizations", string(id), "members", string(userId), "roles"), deleteOrganizationMemberRolesRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrganizationsById Delete organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/delete_organizations_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("organizations", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetEnabledConnections Get connections enabled for an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_enabled_connections
func (m *Manager) GetEnabledConnections(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetEnabledConnections200Response, error) {
	var localVarReturnValue *models.GetEnabledConnections200Response
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id), "enabled_connections"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetEnabledConnectionsByConnectionId Get an enabled connection for an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_enabled_connections_by_connection_id
func (m *Manager) GetEnabledConnection(ctx context.Context, id string, connectionId string, opts ...management.RequestOption) (*models.GetEnabledConnections200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetEnabledConnections200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id), "enabled_connections", string(connectionId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetInvitations Get invitations to an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_invitations
func (m *Manager) GetInvitations(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetInvitations200Response, error) {
	var localVarReturnValue *models.GetInvitations200Response
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id), "invitations"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetInvitationsByInvitationId Get a specific invitation to an Organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_invitations_by_invitation_id
func (m *Manager) GetInvitation(ctx context.Context, id string, invitationId string, opts ...management.RequestOption) (*models.GetInvitations200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetInvitations200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id), "invitations", string(invitationId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetMembers Get members who belong to an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_members
func (m *Manager) GetMembers(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetMembers200Response, error) {
	var localVarReturnValue *models.GetMembers200Response
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id), "members"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetNameByName Get organization by name
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_name_by_name
func (m *Manager) GetByName(ctx context.Context, name string, opts ...management.RequestOption) (*models.GetOrganizations200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizations200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", "name", string(name)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetOrganizationMemberRoles Get user roles assigned to an Organization member
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_organization_member_roles
func (m *Manager) GetMemberRoles(ctx context.Context, id string, userId string, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200Response, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200Response
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id), "members", string(userId), "roles"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetOrganizations Get organizations
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_organizations
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetOrganizations200Response, error) {
	var localVarReturnValue *models.GetOrganizations200Response
	err := m.management.Request(ctx, "GET", m.management.URI("organizations"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetOrganizationsById Get organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/get_organizations_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetOrganizations200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizations200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("organizations", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchEnabledConnectionsByConnectionId Update the Connection of an Organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/patch_enabled_connections_by_connection_id
func (m *Manager) UpdateEnabledConnection(ctx context.Context, id string, connectionId string, patchEnabledConnectionsByConnectionIdRequest *models.PatchEnabledConnectionsByConnectionIdRequest, opts ...management.RequestOption) (*models.GetEnabledConnections200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetEnabledConnections200ResponseOneOfInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("organizations", string(id), "enabled_connections", string(connectionId)), patchEnabledConnectionsByConnectionIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchOrganizationsById Modify an Organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/patch_organizations_by_id
func (m *Manager) Update(ctx context.Context, id string, patchOrganizationsByIdRequest *models.PatchOrganizationsByIdRequest, opts ...management.RequestOption) (*models.GetOrganizations200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetOrganizations200ResponseOneOfInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("organizations", string(id)), patchOrganizationsByIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostEnabledConnections Add connections to an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/post_enabled_connections
func (m *Manager) AddEnabledConnection(ctx context.Context, id string, postEnabledConnectionsRequest *models.PostEnabledConnectionsRequest, opts ...management.RequestOption) (*models.GetEnabledConnections200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetEnabledConnections200ResponseOneOfInner
	err := m.management.Request(ctx, "POST", m.management.URI("organizations", string(id), "enabled_connections"), postEnabledConnectionsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostInvitations Create invitations to an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/post_invitations
func (m *Manager) CreateInvitation(ctx context.Context, id string, postInvitationsRequest *models.PostInvitationsRequest, opts ...management.RequestOption) (*models.GetInvitations200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetInvitations200ResponseOneOfInner
	err := m.management.Request(ctx, "POST", m.management.URI("organizations", string(id), "invitations"), postInvitationsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostMembers Add members to an organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/post_members
func (m *Manager) AddMembers(ctx context.Context, id string, postMembersRequest *models.PostMembersRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("organizations", string(id), "members"), postMembersRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostOrganizationMemberRoles Assign user roles to an Organization member
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/post_organization_member_roles
func (m *Manager) AddMemberRoles(ctx context.Context, id string, userId string, postOrganizationMemberRolesRequest *models.PostOrganizationMemberRolesRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("organizations", string(id), "members", string(userId), "roles"), postOrganizationMemberRolesRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostOrganizations Create an Organization
//
// https://auth0.com/docs/api/management/v2/#!/Organizations/post_organizations
func (m *Manager) Create(ctx context.Context, postOrganizationsRequest *models.PostOrganizationsRequest, opts ...management.RequestOption) (*models.PostOrganizations201Response, error) {
	var localVarReturnValue *models.PostOrganizations201Response
	err := m.management.Request(ctx, "POST", m.management.URI("organizations"), postOrganizationsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
