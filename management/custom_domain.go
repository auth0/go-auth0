package management

import "context"

// CustomDomain to be used on authentication pages.
//
// See: https://auth0.com/docs/customize/custom-domains
type CustomDomain struct {
	// The id of the custom domain.
	ID *string `json:"custom_domain_id,omitempty"`

	// The custom domain.
	Domain *string `json:"domain,omitempty"`

	// The custom domain provisioning type.
	// Can be either "auth0_managed_certs" or "self_managed_certs".
	Type *string `json:"type,omitempty"`

	// Primary is true if the domain was marked as "primary", false otherwise.
	Primary *bool `json:"primary,omitempty"`

	// The custom domain configuration status. Can be any of the following:
	//
	// "disabled", "pending", "pending_verification" or "ready"
	Status *string `json:"status,omitempty"`

	// The origin domain name that the CNAME or reverse proxy should be pointed
	// to.
	OriginDomainName *string `json:"origin_domain_name,omitempty"`

	// For self-managed certificates, when the verification completes for the
	// first time, this field will be set to the value the reverse proxy should
	// send in the cname-api-key header field.
	CNAMEAPIKey *string `json:"cname_api_key,omitempty"`

	// The custom domain verification method. The only allowed value is "txt".
	VerificationMethod *string `json:"verification_method,omitempty"`

	Verification *CustomDomainVerification `json:"verification,omitempty"`

	// The TLS version policy. Can be either "compatible" or "recommended".
	TLSPolicy *string `json:"tls_policy,omitempty"`

	// The HTTP header to fetch the client's IP address.
	CustomClientIPHeader *string `json:"custom_client_ip_header,omitempty"`
}

// CustomDomainVerification is used to verify a CustomDomain.
type CustomDomainVerification struct {
	// The custom domain verification methods.
	Methods []map[string]interface{} `json:"methods,omitempty"`
}

// CustomDomainManager manages Auth0 CustomDomain resources.
type CustomDomainManager manager

// Create a new custom domain.
//
// Note: The custom domain will need to be verified before it starts accepting
// requests.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/post_custom_domains
func (m *CustomDomainManager) Create(ctx context.Context, c *CustomDomain, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("custom-domains"), c, opts...)
}

// Update a custom domain.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/patch_custom_domains_by_id
func (m *CustomDomainManager) Update(ctx context.Context, id string, c *CustomDomain, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("custom-domains", id), c, opts...)
}

// Read a custom domain configuration and status.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains_by_id
func (m *CustomDomainManager) Read(ctx context.Context, id string, opts ...RequestOption) (c *CustomDomain, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("custom-domains", id), &c, opts...)
	return
}

// Verify a custom domain.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/post_verify
func (m *CustomDomainManager) Verify(ctx context.Context, id string, opts ...RequestOption) (c *CustomDomain, err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("custom-domains", id, "verify"), &c, opts...)
	return
}

// Delete a custom domain and stop serving requests for it.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/delete_custom_domains_by_id
func (m *CustomDomainManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("custom-domains", id), nil, opts...)
}

// List all custom domains.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains
func (m *CustomDomainManager) List(ctx context.Context, opts ...RequestOption) (c []*CustomDomain, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("custom-domains"), &c, opts...)
	return
}
