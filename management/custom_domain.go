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

	// Deprecated: Primary field is no longer used and will be removed in a future release.
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

	// Deprecated: The verification status is no longer used and will be removed in a future release.
	// The custom domain verification method. The only allowed value is "txt".
	VerificationMethod *string `json:"verification_method,omitempty"`

	Verification *CustomDomainVerification `json:"verification,omitempty"`

	// The TLS version policy. Can be either "compatible" or "recommended".
	TLSPolicy *string `json:"tls_policy,omitempty"`

	// The HTTP header to fetch the client's IP address.
	CustomClientIPHeader *string `json:"custom_client_ip_header,omitempty"`

	// DomainMetadata holds custom metadata for the domain as key-value pairs (string to string, max 255 characters per value).
	// A maximum of 10 metadata entries is allowed.
	// To remove metadata, set each key's value to nil.
	DomainMetadata *map[string]interface{} `json:"domain_metadata,omitempty"`

	// The custom domain certificate.
	Certificate *CustomDomainCertificate `json:"certificate,omitempty"`
}

// CustomDomainCertificate represents the certificate details for a custom domain.
type CustomDomainCertificate struct {
	// Status indicates the current state of the certificate provisioning process.
	Status *string `json:"status,omitempty"`

	// ErrorMsg contains the error message if the provisioning process fails.
	ErrorMsg *string `json:"error_msg,omitempty"`

	// CertificateAuthority is the name of the certificate authority that issued the certificate.
	CertificateAuthority *string `json:"certificate_authority,omitempty"`

	// RenewsBefore specifies the date by which the certificate should be renewed.
	RenewsBefore *string `json:"renews_before,omitempty"`
}

// CustomDomainList is a list of CustomDomains.
type CustomDomainList struct {
	// List is the list of custom domains.
	List
	// CustomDomains is the list of custom domains.
	CustomDomains []*CustomDomain `json:"custom_domains"`
}

// CustomDomainVerification contains information related to verifying a custom domain.
type CustomDomainVerification struct {
	// Methods defines the list of domain verification methods used.
	Methods []map[string]interface{} `json:"methods,omitempty"`

	// Status represents the current status of the domain verification process.
	Status *string `json:"status,omitempty"`

	// ErrorMsg contains the error message, if any, from the last DNS verification check.
	ErrorMsg *string `json:"error_msg,omitempty"`

	// LastVerifiedAt indicates the last time the domain was successfully verified.
	LastVerifiedAt *string `json:"last_verified_at,omitempty"`
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

// ListWithPagination lists all custom domains with support for pagination.
// This method supports checkpoint pagination. Pagination options (e.g., 'from' and 'take')
// should be passed via opts. The returned CustomDomainList will contain the list of custom domains
// for the current page and pagination information.
//
// See: https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains
func (m *CustomDomainManager) ListWithPagination(ctx context.Context, opts ...RequestOption) (c *CustomDomainList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("custom-domains"), &c, applyListCheckpointDefaults(opts))
	return
}
