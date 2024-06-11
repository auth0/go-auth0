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

	// The algorithm used to sign tokens ["HS256" or "RS256"].
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
