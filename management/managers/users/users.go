package users

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

// DeleteAuthenticationMethods Delete all authentication methods for the given user
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_authentication_methods
func (m *Manager) DeleteAuthenticationMethods(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "authentication-methods"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAuthenticationMethodsByAuthenticationMethodId Delete an authentication method by ID
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_authentication_methods_by_authentication_method_id
func (m *Manager) DeleteAuthenticationMethod(ctx context.Context, id string, authenticationMethodId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "authentication-methods", string(authenticationMethodId)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAuthenticators Delete All Authenticators
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_authenticators
func (m *Manager) DeleteAllAuthenticators(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "authenticators"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMultifactorByProvider Delete a User&#39;s Multi-factor Provider
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_multifactor_by_provider
func (m *Manager) DeleteMultifactorProvider(ctx context.Context, id string, provider models.DeleteMultifactorByProviderProviderParameter, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "multifactor", string(provider)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeletePermissions Remove Permissions from a User
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_permissions
func (m *Manager) DeletePermissions(ctx context.Context, id string, deletePermissionsRequest *models.DeletePermissionsRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "permissions"), deletePermissionsRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRefreshTokensForUser Delete refresh tokens for a user
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_refresh_tokens_for_user
func (m *Manager) DeleteRefreshTokens(ctx context.Context, userId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(userId), "refresh-tokens"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSessionsForUser Delete sessions for user
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_sessions_for_user
func (m *Manager) DeleteSessions(ctx context.Context, userId string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(userId), "sessions"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserIdentityByUserId Unlink a User Identity
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_user_identity_by_user_id
func (m *Manager) Unlink(ctx context.Context, id string, provider models.DeleteUserIdentityByUserIdProviderParameter, userId string, opts ...management.RequestOption) ([]*models.DeleteUserIdentityByUserId200ResponseInner, error) {
	var localVarReturnValue []*models.DeleteUserIdentityByUserId200ResponseInner
	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "identities", string(provider), string(userId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// DeleteUserRoles Removes roles from a user
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_user_roles
func (m *Manager) DeleteRoles(ctx context.Context, id string, deleteUserRolesRequest *models.DeleteUserRolesRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id), "roles"), deleteUserRolesRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUsersById Delete a User
//
// https://auth0.com/docs/api/management/v2/#!/Users/delete_users_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("users", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAuthenticationMethods Get a list of authentication methods
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_authentication_methods
func (m *Manager) GetAuthenticationMethods(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetAuthenticationMethods200Response, error) {
	var localVarReturnValue *models.GetAuthenticationMethods200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "authentication-methods"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetAuthenticationMethodsByAuthenticationMethodId Get an authentication method by ID
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_authentication_methods_by_authentication_method_id
func (m *Manager) GetAuthenticationMethod(ctx context.Context, id string, authenticationMethodId string, opts ...management.RequestOption) (*models.GetAuthenticationMethods200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetAuthenticationMethods200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "authentication-methods", string(authenticationMethodId)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetEnrollments Get the First Confirmed Multi-factor Authentication Enrollment
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_enrollments
func (m *Manager) GetEnrollments(ctx context.Context, id string, opts ...management.RequestOption) ([]*models.UserEnrollment, error) {
	var localVarReturnValue []*models.UserEnrollment
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "enrollments"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetLogsByUser Get user&#39;s log events
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_logs_by_user
func (m *Manager) GetLogs(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetLogs200Response, error) {
	var localVarReturnValue *models.GetLogs200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "logs"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPermissions Get a User&#39;s Permissions
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_permissions
func (m *Manager) GetPermissions(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetPermissions200Response, error) {
	var localVarReturnValue *models.GetPermissions200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "permissions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetRefreshTokensForUser Get refresh tokens for a user
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_refresh_tokens_for_user
func (m *Manager) GetRefreshTokens(ctx context.Context, userId string, opts ...management.RequestOption) (*models.GetRefreshTokensForUser200Response, error) {
	var localVarReturnValue *models.GetRefreshTokensForUser200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(userId), "refresh-tokens"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSessionsForUser Get sessions for user
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_sessions_for_user
func (m *Manager) GetSessions(ctx context.Context, userId string, opts ...management.RequestOption) (*models.GetSessionsForUser200Response, error) {
	var localVarReturnValue *models.GetSessionsForUser200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(userId), "sessions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUserOrganizations List user&#39;s organizations
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_user_organizations
func (m *Manager) GetUserOrganizations(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetUserOrganizations200Response, error) {
	var localVarReturnValue *models.GetUserOrganizations200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "organizations"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUserRoles Get a user&#39;s roles
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_user_roles
func (m *Manager) GetRoles(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetOrganizationMemberRoles200Response, error) {
	var localVarReturnValue *models.GetOrganizationMemberRoles200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id), "roles"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUsers List or Search Users
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_users
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetUsers200Response, error) {
	var localVarReturnValue *models.GetUsers200Response
	err := m.management.Request(ctx, "GET", m.management.URI("users"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetUsersById Get a User
//
// https://auth0.com/docs/api/management/v2/#!/Users/get_users_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetUsers200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetUsers200ResponseOneOfInner
	err := m.management.Request(ctx, "GET", m.management.URI("users", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchAuthenticationMethodsByAuthenticationMethodId Update an authentication method
//
// https://auth0.com/docs/api/management/v2/#!/Users/patch_authentication_methods_by_authentication_method_id
func (m *Manager) UpdateAuthenticationMethod(ctx context.Context, id string, authenticationMethodId string, patchAuthenticationMethodsByAuthenticationMethodIdRequest *models.PatchAuthenticationMethodsByAuthenticationMethodIdRequest, opts ...management.RequestOption) (*models.PatchAuthenticationMethodsByAuthenticationMethodId200Response, error) {
	var localVarReturnValue *models.PatchAuthenticationMethodsByAuthenticationMethodId200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("users", string(id), "authentication-methods", string(authenticationMethodId)), patchAuthenticationMethodsByAuthenticationMethodIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchUsersById Update a User
//
// https://auth0.com/docs/api/management/v2/#!/Users/patch_users_by_id
func (m *Manager) Update(ctx context.Context, id string, userUpdate *models.UserUpdate, opts ...management.RequestOption) (*models.GetUsers200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetUsers200ResponseOneOfInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("users", string(id)), userUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostAuthenticationMethods Creates an authentication method for a given user
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_authentication_methods
func (m *Manager) CreateAuthenticationMethod(ctx context.Context, id string, postAuthenticationMethodsRequest *models.PostAuthenticationMethodsRequest, opts ...management.RequestOption) (*models.PostAuthenticationMethods201Response, error) {
	var localVarReturnValue *models.PostAuthenticationMethods201Response
	err := m.management.Request(ctx, "POST", m.management.URI("users", string(id), "authentication-methods"), postAuthenticationMethodsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostIdentities Link a User Account
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_identities
func (m *Manager) Link(ctx context.Context, id string, postIdentitiesRequest *models.PostIdentitiesRequest, opts ...management.RequestOption) ([]*models.UserIdentity, error) {
	var localVarReturnValue []*models.UserIdentity
	err := m.management.Request(ctx, "POST", m.management.URI("users", string(id), "identities"), postIdentitiesRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostInvalidateRememberBrowser Invalidate All Remembered Browsers for Multi-factor Authentication
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_invalidate_remember_browser
func (m *Manager) InvalidateRememberBrowser(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("users", string(id), "multifactor", "actions", "invalidate-remember-browser"), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostPermissions Assign Permissions to a User
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_permissions
func (m *Manager) AssignPermissions(ctx context.Context, id string, postPermissionsRequest *models.PostPermissionsRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("users", string(id), "permissions"), postPermissionsRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostRecoveryCodeRegeneration Generate New Multi-factor Authentication Recovery Code
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_recovery_code_regeneration
func (m *Manager) RegenerateRecoveryCode(ctx context.Context, id string, opts ...management.RequestOption) (*models.PostRecoveryCodeRegeneration200Response, error) {
	var localVarReturnValue *models.PostRecoveryCodeRegeneration200Response
	err := m.management.Request(ctx, "POST", m.management.URI("users", string(id), "recovery-code-regeneration"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostUserRoles Assign roles to a user
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_user_roles
func (m *Manager) AssignRoles(ctx context.Context, id string, postUserRolesRequest *models.PostUserRolesRequest, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "POST", m.management.URI("users", string(id), "roles"), postUserRolesRequest, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// PostUsers Create a User
//
// https://auth0.com/docs/api/management/v2/#!/Users/post_users
func (m *Manager) Create(ctx context.Context, userCreate *models.UserCreate, opts ...management.RequestOption) (*models.GetUsers200ResponseOneOfInner, error) {
	var localVarReturnValue *models.GetUsers200ResponseOneOfInner
	err := m.management.Request(ctx, "POST", m.management.URI("users"), userCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutAuthenticationMethods Update all authentication methods by replacing them with the given ones
//
// https://auth0.com/docs/api/management/v2/#!/Users/put_authentication_methods
func (m *Manager) UpdateAuthenticationMethods(ctx context.Context, id string, putAuthenticationMethodsRequestInner []*models.PutAuthenticationMethodsRequestInner, opts ...management.RequestOption) ([]*models.PutAuthenticationMethods200ResponseInner, error) {
	var localVarReturnValue []*models.PutAuthenticationMethods200ResponseInner
	err := m.management.Request(ctx, "PUT", m.management.URI("users", string(id), "authentication-methods"), putAuthenticationMethodsRequestInner, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
