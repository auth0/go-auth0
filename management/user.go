package management

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// User represents an Auth0 user resource.
//
// See: https://auth0.com/docs/users
type User struct {
	// The users' identifier.
	ID *string `json:"user_id,omitempty"`

	// The connection the user belongs to.
	Connection *string `json:"connection,omitempty"`

	// The users' email.
	Email *string `json:"email,omitempty"`

	// The users' name.
	Name *string `json:"name,omitempty"`

	// The users' given name.
	GivenName *string `json:"given_name,omitempty"`

	// The users' family name.
	FamilyName *string `json:"family_name,omitempty"`

	// The users' username. Only valid if the connection requires a username.
	Username *string `json:"username,omitempty"`

	// The users' nickname.
	Nickname *string `json:"nickname,omitempty"`

	// The screen name, handle, or alias that this user identifies themselves with.
	ScreenName *string `json:"screen_name,omitempty"`

	// The user-defined UTF-8 string describing their account.
	Description *string `json:"description,omitempty"`

	// The user-defined location for this account’s profile.
	Location *string `json:"location,omitempty"`

	// The users' password (mandatory for non SMS connections)
	Password *string `json:"password,omitempty"`

	// The users' phone number (following the E.164 recommendation).
	// Only valid for users to be added to SMS connections.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// The time the user was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The last time the user was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The last time the user has logged in.
	LastLogin *time.Time `json:"last_login,omitempty"`

	// The last time the user had their password reset.
	// Only available for Database connection users.
	LastPasswordReset *time.Time `json:"last_password_reset,omitempty"`

	// UserMetadata holds data that the user has read/write access to.
	// For example color_preference, blog_url, etc.
	UserMetadata *map[string]interface{} `json:"user_metadata,omitempty"`

	// Identities is a list of user identities for when accounts are linked.
	Identities []*UserIdentity `json:"identities,omitempty"`

	// True if the user's email is verified, false otherwise. If it is true then
	// the user will not receive a verification email, unless verify_email: true
	// was specified.
	EmailVerified *bool `json:"-"`

	// If true, the user will receive a verification email after creation, even
	// if created with email_verified set to true. If false, the user will not
	// receive a verification email, even if created with email_verified set to
	// false. If unspecified, defaults to the behavior determined by the value
	// of email_verified.
	VerifyEmail *bool `json:"verify_email,omitempty"`

	// True if the user's phone number is verified, false otherwise. When the
	// user is added to an SMS connection, they will not receive a verification
	// SMS if this is true.
	PhoneVerified *bool `json:"phone_verified,omitempty"`

	// AppMetadata holds data that the user has read-only access to.
	// For example roles, permissions, vip, etc.
	// NOTE: Roles added to AppMetadata are not integrated with Auth0 Role-Based Access Control (RBAC).
	// For RBAC, see the functions User.Roles, User.AssignRoles, and User.RemoveRoles.
	AppMetadata *map[string]interface{} `json:"app_metadata,omitempty"`

	// The user's picture url.
	Picture *string `json:"picture,omitempty"`

	// A URL provided by the user in association with their profile.
	URL *string `json:"url,omitempty"`

	// True if the user is blocked from the application, false if the user is enabled.
	Blocked *bool `json:"blocked,omitempty"`

	// Last IP address from which this user logged in. Read only, cannot be modified.
	LastIP *string `json:"last_ip,omitempty"`

	// Total number of logins this user has performed. Read only, cannot be modified.
	LoginsCount *int64 `json:"logins_count,omitempty"`

	// List of multi-factor authentication providers with which this user has enrolled.
	Multifactor *[]string `json:"multifactor,omitempty"`

	// Auth0 client ID. Only valid when updating email address.
	ClientID *string `json:"client_id,omitempty"`
}

// UnmarshalJSON is a custom deserializer for the User type.
//
// We have to use a custom one due to possible inconsistencies in value types.
func (u *User) UnmarshalJSON(b []byte) error {
	type user User
	type userAlias struct {
		*user
		RawEmailVerified interface{} `json:"email_verified,omitempty"`
	}

	alias := &userAlias{(*user)(u), nil}

	err := json.Unmarshal(b, alias)
	if err != nil {
		return err
	}

	if alias.RawEmailVerified != nil {
		var emailVerified bool
		switch rawEmailVerified := alias.RawEmailVerified.(type) {
		case bool:
			emailVerified = rawEmailVerified
		case string:
			emailVerified, err = strconv.ParseBool(rawEmailVerified)
			if err != nil {
				return err
			}
		default:
			panic(reflect.TypeOf(rawEmailVerified))
		}
		alias.EmailVerified = &emailVerified
	}

	return nil
}

// MarshalJSON is a custom serializer for the User type.
func (u *User) MarshalJSON() ([]byte, error) {
	type user User
	type userAlias struct {
		*user
		RawEmailVerified interface{} `json:"email_verified,omitempty"`
	}

	alias := &userAlias{user: (*user)(u)}
	if u.EmailVerified != nil {
		alias.RawEmailVerified = u.EmailVerified
	}

	return json.Marshal(alias)
}

// UserIdentityLink contains the data needed for linking an identity to a given user.
type UserIdentityLink struct {
	// Connection id of the secondary user account being linked when more than one auth0 database provider exists.
	ConnectionID *string `json:"connection_id,omitempty"`

	// Secondary account user id.
	UserID *string `json:"user_id,omitempty"`

	// Identity provider of the secondary user account being linked.
	Provider *string `json:"provider,omitempty"`

	// LinkWith requires the authenticated primary account's JWT in the Authorization header.
	// Must be a JWT for the secondary account being linked. If sending this parameter,
	// provider, user_id, and connection_id must not be sent.
	LinkWith *string `json:"link_with,omitempty"`
}

// UserIdentity holds values that validate a User's identity.
type UserIdentity struct {
	Connection        *string                 `json:"connection,omitempty"`
	UserID            *string                 `json:"-"`
	Provider          *string                 `json:"provider,omitempty"`
	IsSocial          *bool                   `json:"isSocial,omitempty"`
	AccessToken       *string                 `json:"access_token,omitempty"`
	AccessTokenSecret *string                 `json:"access_token_secret,omitempty"`
	RefreshToken      *string                 `json:"refresh_token,omitempty"`
	ProfileData       *map[string]interface{} `json:"profileData,omitempty"`
}

// UnmarshalJSON is a custom deserializer for the UserIdentity type.
//
// We have to use a custom one due to a bug in the Auth0 Management API which
// might return a number for `user_id` instead of a string.
//
// See https://community.auth0.com/t/users-user-id-returns-inconsistent-type-for-identities-user-id/39236
func (i *UserIdentity) UnmarshalJSON(b []byte) error {
	type userIdentity UserIdentity
	type userIdentityAlias struct {
		*userIdentity
		RawUserID interface{} `json:"user_id,omitempty"`
	}

	alias := &userIdentityAlias{(*userIdentity)(i), nil}

	err := json.Unmarshal(b, alias)
	if err != nil {
		return err
	}

	if alias.RawUserID != nil {
		var id string
		switch rawID := alias.RawUserID.(type) {
		case string:
			id = rawID
		case float64:
			id = strconv.Itoa(int(rawID))
		default:
			panic(reflect.TypeOf(rawID))
		}
		alias.UserID = &id
	}

	return nil
}

// MarshalJSON is a custom serializer for the UserIdentity type.
func (i *UserIdentity) MarshalJSON() ([]byte, error) {
	type userIdentity UserIdentity
	type userIdentityAlias struct {
		*userIdentity
		RawUserID interface{} `json:"user_id,omitempty"`
	}

	alias := &userIdentityAlias{userIdentity: (*userIdentity)(i)}
	if i.UserID != nil {
		alias.RawUserID = i.UserID
	}

	return json.Marshal(alias)
}

type userBlock struct {
	BlockedFor []*UserBlock `json:"blocked_for,omitempty"`
}

// UserBlock keeps track of a blocked IP for the login identifier associated with a User.
type UserBlock struct {
	Identifier *string `json:"identifier,omitempty"`
	IP         *string `json:"ip,omitempty"`
}

// UserRecoveryCode represents a User's multi-factor authentication recovery code.
type UserRecoveryCode struct {
	RecoveryCode *string `json:"recovery_code,omitempty"`
}

// UserEnrollment contains information about the Guardian enrollments for the user.
type UserEnrollment struct {
	// Authentication method for this enrollment. Can be `authentication`, `guardian`, or `sms`.
	AuthMethod *string `json:"auth_method,omitempty"`

	// Start date and time of this enrollment.
	EnrolledAt *time.Time `json:"enrolled_at,omitempty"`

	// ID of this enrollment.
	ID *string `json:"id,omitempty"`

	// Device identifier (usually phone identifier) of this enrollment.
	Identifier *string `json:"identifier,omitempty"`

	// Last authentication date and time of this enrollment.
	LastAuth *time.Time `json:"last_auth,omitempty"`

	// Name of enrollment (usually phone number).
	Name *string `json:"name,omitempty"`

	// Phone number for this enrollment.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// Status of this enrollment. Can be `pending` or `confirmed`.
	Status *string `json:"status,omitempty"`

	// Type of enrollment.
	Type *string `json:"type,omitempty"`
}

// UserList is an envelope struct which is used when calling List() or Search() methods.
//
// It holds metadata such as the total result count, starting offset and limit.
type UserList struct {
	List
	Users []*User `json:"users"`
}

// AuthenticationMethod belonging to a user.
//
// See: https://auth0.com/docs/secure/multi-factor-authentication/manage-mfa-auth0-apis/manage-authentication-methods-with-management-api
type AuthenticationMethod struct {
	// The ID of the authentication method (auto generated).
	ID *string `json:"id,omitempty"`

	// The type of the authentication method. Should be one of "phone", "email", "totp", "webauthn-roaming", or "passkey".
	Type *string `json:"type,omitempty"`

	// The authentication method status.
	Confirmed *bool `json:"confirmed,omitempty"`

	// A human-readable label to identify the authentication method.
	Name *string `json:"name,omitempty"`

	// The ID of a linked authentication method. Linked authentication methods will be deleted together.
	LinkID *string `json:"link_id,omitempty"`

	// Applies to phone authentication methods only. The destination phone number used to send verification codes via text and voice.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// Applies to email authentication method only. The email address used to send verification messages.
	Email *string `json:"email,omitempty"`

	// Applies to webauthn authentication methods only. The ID of the generated credential.
	KeyID *string `json:"key_id,omitempty"`

	// Applies to webauthn authentication methods only. The public key.
	PublicKey *string `json:"public_key,omitempty"`

	// Authenticator creation date.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Enrollment date.
	EnrolledAt *time.Time `json:"enrolled_at,omitempty"`

	// Last authentication.
	LastAuthedAt *time.Time `json:"last_auth_at,omitempty"`

	// Base32 encoded secret for TOTP generation.
	TOTPSecret *string `json:"totp_secret,omitempty"`

	// The authentication method preferred for phone authenticators.
	PreferredAuthenticationMethod *string `json:"preferred_authentication_method,omitempty"`

	// Applies to email webauthn authenticators only. The relying party identifier.
	RelyingPartyIdentifier *string `json:"relying_party_identifier,omitempty"`

	AuthenticationMethods *[]AuthenticationMethodReference `json:"authentication_methods,omitempty"`

	// Applies to passkeys only. The kind of device the credential is stored on as defined by backup eligibility.
	// "single_device" credentials cannot be backed up and synced to another device,
	// "multi_device" credentials can be backed up if enabled by the end-user.
	CredentialDeviceType *string `json:"credential_device_type,omitempty"`

	// Applies to passkeys only. Whether the credential was backed up.
	CredentialBackedUp *bool `json:"credential_backed_up,omitempty"`

	// Applies to passkeys only. The ID of the user identity linked with the authentication method.
	IdentityUserID *string `json:"identity_user_id,omitempty"`

	// Applies to passkeys only. The user-agent of the browser used to create the passkey.
	UserAgent *string `json:"user_agent,omitempty"`
}

// AuthenticationMethodReference used within the AuthenticationMethod.
type AuthenticationMethodReference struct {
	// The ID of the authentication method (auto generated).
	ID *string `json:"id,omitempty"`
	// The type of the authentication method.
	Type *string `json:"type,omitempty"`
}

// AuthenticationMethodList is an envelope struct which is used when calling GetAuthenticationMethods().
//
// It holds metadata such as the total result count, starting offset and limit.
type AuthenticationMethodList struct {
	List
	Authenticators []*AuthenticationMethod `json:"authenticators,omitempty"`
}

// UserManager manages Auth0 User resources.
type UserManager manager

// newUserManager returns a new instance of a user manager.

// Create a new user. It works only for database and passwordless connections.
//
// The samples on the right show you every attribute that could be used. The
// attribute connection is always mandatory but depending on the type of
// connection you are using there could be others too. For instance, database
// connections require `email` and `password`.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_users
func (m *UserManager) Create(ctx context.Context, u *User, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("users"), u, opts...)
}

// Read user details for a given user_id.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_users_by_id
func (m *UserManager) Read(ctx context.Context, id string, opts ...RequestOption) (u *User, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", id), &u, opts...)
	return
}

// Update user.
//
// The following attributes can be updated at the root level:
//
// - `app_metadata`
// - `blocked`
// - `email`
// - `email_verified`
// - `family_name`
// - `given_name`
// - `name`
// - `nickname`
// - `password`
// - `phone_number`
// - `phone_verified`
// - `picture`
// - `username`
// - `user_metadata`
// - `verify_email`
//
// See: https://auth0.com/docs/api/management/v2#!/Users/patch_users_by_id
func (m *UserManager) Update(ctx context.Context, id string, u *User, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("users", id), u, opts...)
}

// Delete a single user based on its id.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_users_by_id
func (m *UserManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("users", id), nil, opts...)
}

// List users. This method forces the `include_totals` option.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_users
func (m *UserManager) List(ctx context.Context, opts ...RequestOption) (ul *UserList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users"), &ul, applyListDefaults(opts))
	return
}

// Search is an alias for List.
func (m *UserManager) Search(ctx context.Context, opts ...RequestOption) (ul *UserList, err error) {
	return m.List(ctx, opts...)
}

// ListByEmail retrieves all users matching a given email.
//
// If Auth0 is the identify provider (idP), the email address associated with a
// user is saved in lower case, regardless of how you initially provided it.
// For example, if you register a user as JohnSmith@example.com, Auth0 saves the
// user's email as johnsmith@example.com.
//
// In cases where Auth0 is not the idP, the `email` is stored based on the rules
// of idP, so make sure the search is made using the correct capitalization.
//
// When using this endpoint, make sure that you are searching for users via
// email addresses using the correct case.
//
// See: https://auth0.com/docs/api/management/v2#!/Users_By_Email/get_users_by_email
func (m *UserManager) ListByEmail(ctx context.Context, email string, opts ...RequestOption) (us []*User, err error) {
	opts = append(opts, Parameter("email", email))
	err = m.management.Request(ctx, "GET", m.management.URI("users-by-email"), &us, opts...)
	return
}

// Roles lists roles associated with a user.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_user_roles
func (m *UserManager) Roles(ctx context.Context, id string, opts ...RequestOption) (r *RoleList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", id, "roles"), &r, applyListDefaults(opts))
	return
}

// AssignRoles assigns roles to a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_user_roles
func (m *UserManager) AssignRoles(ctx context.Context, id string, roles []*Role, opts ...RequestOption) error {
	r := make(map[string][]*string)
	r["roles"] = make([]*string, len(roles))
	for i, role := range roles {
		r["roles"][i] = role.ID
	}
	return m.management.Request(ctx, "POST", m.management.URI("users", id, "roles"), &r, opts...)
}

// RemoveRoles removes any roles associated to a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_user_roles
func (m *UserManager) RemoveRoles(ctx context.Context, id string, roles []*Role, opts ...RequestOption) error {
	r := make(map[string][]*string)
	r["roles"] = make([]*string, len(roles))
	for i, role := range roles {
		r["roles"][i] = role.ID
	}
	return m.management.Request(ctx, "DELETE", m.management.URI("users", id, "roles"), &r, opts...)
}

// Permissions lists the permissions associated to the user.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_permissions
func (m *UserManager) Permissions(ctx context.Context, id string, opts ...RequestOption) (p *PermissionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", id, "permissions"), &p, applyListDefaults(opts))
	return
}

// AssignPermissions assigns permissions to the user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_permissions
func (m *UserManager) AssignPermissions(ctx context.Context, id string, permissions []*Permission, opts ...RequestOption) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return m.management.Request(ctx, "POST", m.management.URI("users", id, "permissions"), &p, opts...)
}

// RemovePermissions removes any permissions associated to a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_permissions
func (m *UserManager) RemovePermissions(ctx context.Context, id string, permissions []*Permission, opts ...RequestOption) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return m.management.Request(ctx, "DELETE", m.management.URI("users", id, "permissions"), &p, opts...)
}

// Blocks retrieves a list of blocked IP addresses of a particular user using the
// user ID.
//
// See: https://auth0.com/docs/api/management/v2#!/User_Blocks/get_user_blocks_by_id
func (m *UserManager) Blocks(ctx context.Context, id string, opts ...RequestOption) ([]*UserBlock, error) {
	b := new(userBlock)
	err := m.management.Request(ctx, "GET", m.management.URI("user-blocks", id), &b, opts...)
	return b.BlockedFor, err
}

// BlocksByIdentifier retrieves a list of blocked IP addresses of a particular
// user using any of the user identifiers: username, phone number or email.
//
// See: https://auth0.com/docs/api/management/v2#!/User_Blocks/get_user_blocks
func (m *UserManager) BlocksByIdentifier(ctx context.Context, identifier string, opts ...RequestOption) ([]*UserBlock, error) {
	b := new(userBlock)
	opts = append(opts, Parameter("identifier", identifier))
	err := m.management.Request(ctx, "GET", m.management.URI("user-blocks"), &b, opts...)
	return b.BlockedFor, err
}

// Unblock a user that was blocked due to an excessive amount of incorrectly
// provided credentials using the user ID.
//
// Note: This endpoint does not unblock users that were blocked by admins.
//
// See: https://auth0.com/docs/api/management/v2#!/User_Blocks/delete_user_blocks_by_id
func (m *UserManager) Unblock(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("user-blocks", id), nil, opts...)
}

// UnblockByIdentifier a user that was blocked due to an excessive amount of incorrectly
// provided credentials using any of the user identifiers: username, phone number or email.
//
// Note: This endpoint does not unblock users that were blocked by admins.
//
// See: https://auth0.com/docs/api/management/v2#!/User_Blocks/delete_user_blocks
func (m *UserManager) UnblockByIdentifier(ctx context.Context, identifier string, opts ...RequestOption) error {
	opts = append(opts, Parameter("identifier", identifier))
	return m.management.Request(ctx, "DELETE", m.management.URI("user-blocks"), nil, opts...)
}

// Enrollments retrieves all Guardian enrollments for a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_enrollments
func (m *UserManager) Enrollments(ctx context.Context, id string, opts ...RequestOption) (enrolls []*UserEnrollment, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", id, "enrollments"), &enrolls, opts...)
	return
}

// RegenerateRecoveryCode removes the current multi-factor authentication recovery code and generate a new one.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_recovery_code_regeneration
func (m *UserManager) RegenerateRecoveryCode(ctx context.Context, id string, opts ...RequestOption) (*UserRecoveryCode, error) {
	r := new(UserRecoveryCode)
	err := m.management.Request(ctx, "POST", m.management.URI("users", id, "recovery-code-regeneration"), &r, opts...)
	return r, err
}

// InvalidateRememberBrowser invalidates all remembered browsers across all authentication factors for a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_invalidate_remember_browser
func (m *UserManager) InvalidateRememberBrowser(ctx context.Context, id string, opts ...RequestOption) error {
	uri := m.management.URI(
		"users",
		id,
		"multifactor",
		"actions",
		"invalidate-remember-browser",
	)
	err := m.management.Request(ctx, "POST", uri, nil, opts...)
	return err
}

// Link links two user accounts together forming a primary and secondary relationship.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_identities
func (m *UserManager) Link(ctx context.Context, id string, il *UserIdentityLink, opts ...RequestOption) (uIDs []UserIdentity, err error) {
	request, err := m.management.NewRequest(ctx, "POST", m.management.URI("users", id, "identities"), il, opts...)
	if err != nil {
		return uIDs, err
	}

	response, err := m.management.Do(request)
	if err != nil {
		return uIDs, err
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		return uIDs, newError(response)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return uIDs, fmt.Errorf("failed to read the response body: %w", err)
	}

	if len(responseBody) > 0 {
		if err = json.Unmarshal(responseBody, &uIDs); err != nil {
			return uIDs, fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return uIDs, nil
}

// Unlink unlinks an identity from a user making it a separate account again.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_user_identity_by_user_id
func (m *UserManager) Unlink(ctx context.Context, id, provider, userID string, opts ...RequestOption) (uIDs []UserIdentity, err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("users", id, "identities", provider, userID), &uIDs, opts...)
	return
}

// Organizations lists user's organizations.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_organizations
func (m *UserManager) Organizations(ctx context.Context, id string, opts ...RequestOption) (p *OrganizationList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", id, "organizations"), &p, applyListDefaults(opts))
	return
}

// ListAuthenticationMethods retrieves a list of authentication methods.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_authentication_methods
func (m *UserManager) ListAuthenticationMethods(ctx context.Context, userID string, opts ...RequestOption) (a *AuthenticationMethodList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", userID, "authentication-methods"), &a, applyListDefaults(opts))
	return
}

// GetAuthenticationMethodByID gets a specific authentication method for a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_authentication_methods_by_authentication_method_id
func (m *UserManager) GetAuthenticationMethodByID(ctx context.Context, userID string, id string, opts ...RequestOption) (a *AuthenticationMethod, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("users", userID, "authentication-methods", id), &a, opts...)
	return
}

// CreateAuthenticationMethod creates an authentication method for a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_authentication_methods
func (m *UserManager) CreateAuthenticationMethod(ctx context.Context, userID string, a *AuthenticationMethod, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("users", userID, "authentication-methods"), &a, opts...)
	return
}

// UpdateAllAuthenticationMethods updates all authentication methods by replacing them with the given ones.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/put_authentication_methods
func (m *UserManager) UpdateAllAuthenticationMethods(ctx context.Context, userID string, a *[]AuthenticationMethod, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PUT", m.management.URI("users", userID, "authentication-methods"), &a, opts...)
	return
}

// UpdateAuthenticationMethod updates an authentication method by ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/patch_authentication_methods_by_authentication_method_id
func (m *UserManager) UpdateAuthenticationMethod(ctx context.Context, userID string, id string, a *AuthenticationMethod, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PATCH", m.management.URI("users", userID, "authentication-methods", id), &a, opts...)
	return
}

// DeleteAuthenticationMethod deletes an authentication method by ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_authentication_methods_by_authentication_method_id
func (m *UserManager) DeleteAuthenticationMethod(ctx context.Context, userID string, id string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("users", userID, "authentication-methods", id), nil, opts...)
	return
}

// DeleteAllAuthenticationMethods deletes all authentication methods for the given user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_authentication_methods
func (m *UserManager) DeleteAllAuthenticationMethods(ctx context.Context, userID string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("users", userID, "authentication-methods"), nil, opts...)
	return
}
