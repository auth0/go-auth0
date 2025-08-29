package management

import "context"

// ResourceServer is an entity that represents an external resource, capable of
// accepting and responding to protected resource requests made by applications.
type ResourceServer struct {
	// A generated string identifying the resource server.
	ID *string `json:"id,omitempty"`

	// The name of the resource server. Must contain at least one character.
	// Does not allow '<' or '>'
	Name *string `json:"name,omitempty"`

	// The identifier of the resource server.
	Identifier *string `json:"identifier,omitempty"`

	// Scopes supported by the resource server.
	Scopes *[]ResourceServerScope `json:"scopes,omitempty"`

	// Algorithm used to sign JWTs. Can be `HS256` or `RS256`. `PS256` available via addon.
	SigningAlgorithm *string `json:"signing_alg,omitempty"`

	// The secret used to sign tokens when using symmetric algorithms.
	SigningSecret *string `json:"signing_secret,omitempty"`

	// Allows issuance of refresh tokens for this entity.
	AllowOfflineAccess *bool `json:"allow_offline_access,omitempty"`

	// The amount of time in seconds that the token will be valid after being
	// issued.
	TokenLifetime *int `json:"token_lifetime,omitempty"`

	// The amount of time in seconds that the token will be valid after being
	// issued from browser based flows. Value cannot be larger than
	// token_lifetime.
	TokenLifetimeForWeb *int `json:"token_lifetime_for_web,omitempty"`

	// Flag this entity as capable of skipping consent.
	SkipConsentForVerifiableFirstPartyClients *bool `json:"skip_consent_for_verifiable_first_party_clients,omitempty"`

	// A URI from which to retrieve JWKs for this resource server used for
	// verifying the JWT sent to Auth0 for token introspection.
	VerificationLocation *string `json:"verificationLocation,omitempty"`

	Options *map[string]string `json:"options,omitempty"`

	// Enables the enforcement of the authorization policies.
	EnforcePolicies *bool `json:"enforce_policies,omitempty"`

	// TokenDialect specifies the dialect of access tokens that should be issued for this resource server.
	//
	// Available options:
	//   - "access_token": A JWT containing standard Auth0 claims.
	//   - "rfc9068_profile": A JWT conforming to the IETF JWT Access Token Profile.
	//   - "access_token_authz": A JWT containing standard Auth0 claims, including RBAC permissions claims.
	//   - "rfc9068_profile_authz": A JWT conforming to the IETF JWT Access Token Profile, including RBAC permissions claims.
	//
	//  Note:  RBAC permissions claims are available if RBAC (enforce_policies) is enabled for this API."
	// For more details, see the Access Token Profiles documentation : https://auth0.com/docs/secure/tokens/access-tokens/access-token-profiles.
	TokenDialect *string `json:"token_dialect,omitempty"`

	// ConsentPolicy specifies the consent policy for the resource server.
	//
	// Available options:
	//   - "transactional-authorization-with-mfa"
	//   - null
	//
	// To unset values (set to null), use a PATCH request like this:
	//
	// PATCH /api/v2/resource-servers/{id}
	//
	// {
	//   "consent_policy": null
	// }
	//
	// For more details on making custom requests, refer to the Auth0 Go SDK examples:
	// https://github.com/auth0/go-auth0/blob/main/EXAMPLES.md#providing-a-custom-user-struct
	ConsentPolicy *string `json:"consent_policy,omitempty"`

	// The list of authorization details for the resource server.
	//
	// To unset values (set to null), use a PATCH request like this:
	//
	// PATCH /api/v2/resource-servers/{id}
	//
	// {
	//	 "authorization_details": null
	// }
	//
	// For more details on making custom requests, refer to the Auth0 Go SDK examples:
	// https://github.com/auth0/go-auth0/blob/main/EXAMPLES.md#providing-a-custom-user-struct
	AuthorizationDetails *[]ResourceServerAuthorizationDetails `json:"authorization_details,omitempty"`

	// TokenEncryption specifies the token encryption for the resource server.
	//
	// Available options:
	//   - "compact-nested-jwe"
	//   - null
	TokenEncryption *ResourceServerTokenEncryption `json:"token_encryption,omitempty"`

	// Proof-of-Possession configuration for access tokens.
	//
	// To unset values (set to null), use a PATCH request like this:
	//
	// PATCH /api/v2/resource-servers/{id}
	//
	// {
	//	 "proof_of_possession": null
	// }
	//
	// For more details on making custom requests, refer to the Auth0 Go SDK examples:
	// https://github.com/auth0/go-auth0/blob/main/EXAMPLES.md#providing-a-custom-user-struct
	ProofOfPossession *ResourceServerProofOfPossession `json:"proof_of_possession,omitempty"`

	// SubjectTypeAuthorization defines the authorization policies for user and client flows.
	SubjectTypeAuthorization *ResourceServerSubjectTypeAuthorization `json:"subject_type_authorization,omitempty"`
}

// ResourceServerSubjectTypeAuthorization defines the authorization policies for user and client flows.
type ResourceServerSubjectTypeAuthorization struct {
	// User authorization policies for the resource server.
	User *ResourceServerSubjectTypeAuthorizationUser `json:"user,omitempty"`

	// Client authorization policies for the resource server.
	Client *ResourceServerSubjectTypeAuthorizationClient `json:"client,omitempty"`
}

// ResourceServerSubjectTypeAuthorizationUser defines the authorization policies for user-initiated flows.
type ResourceServerSubjectTypeAuthorizationUser struct {
	// Policy defines the user flows policy for the resource server.
	//
	// Available options:
	//   - "allow_all": Allows all user-initiated flows.
	//   - "deny_all": Denies all user-initiated flows.
	//   - "require_client_grant": Requires a client grant for user-initiated flows.
	Policy *string `json:"policy,omitempty"`
}

// ResourceServerSubjectTypeAuthorizationClient defines the authorization policies for client-initiated flows.
type ResourceServerSubjectTypeAuthorizationClient struct {
	// Policy defines the client flows policy for the resource server.
	//
	// Available options:
	//   - "deny_all": Denies all client-initiated flows.
	//   - "require_client_grant": Requires a client grant for client-initiated flows.
	Policy *string `json:"policy,omitempty"`
}

// ResourceServerAuthorizationDetails specifies the authorization details for the resource server.
type ResourceServerAuthorizationDetails struct {
	// The authorization_detail type identifier.
	Type *string `json:"type,omitempty"`
}

// ResourceServerProofOfPossession specifies the proof-of-possession configuration for access tokens.
type ResourceServerProofOfPossession struct {
	// Intended mechanism for Proof-of-Possession.
	//
	// Available options:
	//   - "mtls"
	//   - "dpop"
	Mechanism *string `json:"mechanism,omitempty"`

	// Whether the use of Proof-of-Possession is required for the resource server.
	Required *bool `json:"required,omitempty"`
}

// ResourceServerTokenEncryption specifies the token encryption for the resource server.
type ResourceServerTokenEncryption struct {
	// Format of the encrypted JWT payload.
	Format *string `json:"format,omitempty"`

	// EncryptionKey specifies the encryption key for the token encryption.
	EncryptionKey *ResourceServerTokenEncryptionKey `json:"encryption_key,omitempty"`
}

// ResourceServerTokenEncryptionKey specifies the encryption key for the token encryption.
type ResourceServerTokenEncryptionKey struct {
	// Name of the encryption key.
	Name *string `json:"name,omitempty"`

	// Algorithm used to encrypt the token.
	Alg *string `json:"alg,omitempty"`

	// Key ID.
	Kid *string `json:"kid,omitempty"`

	// PEM-formatted public key. Must be JSON escaped
	Pem *string `json:"pem,omitempty"`
}

// ResourceServerScope defines the specific actions, resource servers can be allowed to do.
type ResourceServerScope struct {
	// The scope name. Use the format <action>:<resource>.
	// For example 'delete:client_grants'.
	Value *string `json:"value,omitempty"`

	// Description of the scope.
	Description *string `json:"description,omitempty"`
}

// ResourceServerList is a list of ResourceServers.
type ResourceServerList struct {
	List
	ResourceServers []*ResourceServer `json:"resource_servers"`
}

// ResourceServerManager is used for managing a ResourceServer.
type ResourceServerManager manager

// Create a resource server.
//
// See: https://auth0.com/docs/api/management/v2#!/Resource_Servers/post_resource_servers
func (m *ResourceServerManager) Create(ctx context.Context, rs *ResourceServer, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("resource-servers"), rs, opts...)
}

// Read retrieves a resource server by its id or audience.
//
// See: https://auth0.com/docs/api/management/v2#!/Resource_Servers/get_resource_servers_by_id
func (m *ResourceServerManager) Read(ctx context.Context, id string, opts ...RequestOption) (rs *ResourceServer, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("resource-servers", id), &rs, opts...)
	return
}

// Update a resource server.
//
// See: https://auth0.com/docs/api/management/v2#!/Resource_Servers/patch_resource_servers_by_id
func (m *ResourceServerManager) Update(ctx context.Context, id string, rs *ResourceServer, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("resource-servers", id), rs, opts...)
}

// Delete a resource server.
//
// See: https://auth0.com/docs/api/management/v2#!/Resource_Servers/delete_resource_servers_by_id
func (m *ResourceServerManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("resource-servers", id), nil, opts...)
}

// List resource server.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Resource_Servers/get_resource_servers
func (m *ResourceServerManager) List(ctx context.Context, opts ...RequestOption) (rl *ResourceServerList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("resource-servers"), &rl, applyListDefaults(opts))
	return
}
