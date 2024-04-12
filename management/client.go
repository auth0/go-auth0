package management

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/auth0/go-auth0"
)

// Client is an application or a sso integration.
//
// See: https://auth0.com/docs/get-started/applications
type Client struct {
	// The name of the client.
	Name *string `json:"name,omitempty"`

	// Free text description of the purpose of the Client.
	// Max character length is 140.
	Description *string `json:"description,omitempty"`

	// The ID of the client.
	ClientID *string `json:"client_id,omitempty"`

	// The client secret, it must not be public.
	ClientSecret *string `json:"client_secret,omitempty"`

	// The type of application this client represents.
	AppType *string `json:"app_type,omitempty"`

	// The URL of the client logo (recommended size: 150x150).
	LogoURI *string `json:"logo_uri,omitempty"`

	// Whether this client a first party client or not.
	IsFirstParty *bool `json:"is_first_party,omitempty"`

	// Set header `auth0-forwarded-for` as trusted to be used as source
	// of end user ip for brute-force-protection on token endpoint.
	IsTokenEndpointIPHeaderTrusted *bool `json:"is_token_endpoint_ip_header_trusted,omitempty"`

	// Whether this client will conform to strict OIDC specifications.
	OIDCConformant *bool `json:"oidc_conformant,omitempty"`

	// The URLs that Auth0 can use to as a callback for the client.
	Callbacks      *[]string `json:"callbacks,omitempty"`
	AllowedOrigins *[]string `json:"allowed_origins,omitempty"`

	// A set of URLs that represents valid web origins for use with web message response mode.
	WebOrigins        *[]string               `json:"web_origins,omitempty"`
	ClientAliases     *[]string               `json:"client_aliases,omitempty"`
	AllowedClients    *[]string               `json:"allowed_clients,omitempty"`
	AllowedLogoutURLs *[]string               `json:"allowed_logout_urls,omitempty"`
	JWTConfiguration  *ClientJWTConfiguration `json:"jwt_configuration,omitempty"`

	// Client signing keys.
	SigningKeys   []map[string]string `json:"signing_keys,omitempty"`
	EncryptionKey *map[string]string  `json:"encryption_key,omitempty"`
	SSO           *bool               `json:"sso,omitempty"`

	// True to disable Single Sign On, false otherwise (default: false).
	SSODisabled *bool `json:"sso_disabled,omitempty"`

	// True if this client can be used to make cross-origin authentication
	// requests, false otherwise (default: false).
	CrossOriginAuth *bool `json:"cross_origin_authentication,omitempty"`

	// List of acceptable Grant Types for this Client.
	GrantTypes *[]string `json:"grant_types,omitempty"`

	// URL for the location in your site where the cross origin verification
	// takes place for the cross-origin auth flow when performing Auth in your
	// own domain instead of Auth0 hosted login page.
	CrossOriginLocation *string `json:"cross_origin_loc,omitempty"`

	// True if the custom login page is to be used, false otherwise. Defaults to true.
	CustomLoginPageOn      *bool         `json:"custom_login_page_on,omitempty"`
	CustomLoginPage        *string       `json:"custom_login_page,omitempty"`
	CustomLoginPagePreview *string       `json:"custom_login_page_preview,omitempty"`
	FormTemplate           *string       `json:"form_template,omitempty"`
	Addons                 *ClientAddons `json:"addons,omitempty"`

	// Defines the requested authentication method for the token endpoint.
	// Possible values are:
	// 	'none' (public client without a client secret),
	// 	'client_secret_post' (client uses HTTP POST parameters) or
	// 	'client_secret_basic' (client uses HTTP Basic)
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`

	// Metadata associated with the client, in the form of an object with string values (max 255 chars).
	// Maximum of 10 metadata properties allowed. Field names (max 255 chars) are alphanumeric and may
	// only include the following special characters: :,-+=_*?"/\()<>@ [Tab] [Space].
	// To remove a key, the value needs to be sent as null.
	ClientMetadata *map[string]interface{} `json:"client_metadata,omitempty"`

	Mobile *ClientMobile `json:"mobile,omitempty"`

	// Initiate login uri, must be https and cannot contain a fragment.
	InitiateLoginURI *string `json:"initiate_login_uri,omitempty"`

	NativeSocialLogin *ClientNativeSocialLogin `json:"native_social_login,omitempty"`
	RefreshToken      *ClientRefreshToken      `json:"refresh_token,omitempty"`

	OrganizationUsage           *string `json:"organization_usage,omitempty"`
	OrganizationRequireBehavior *string `json:"organization_require_behavior,omitempty"`

	ClientAuthenticationMethods *ClientAuthenticationMethods `json:"client_authentication_methods,omitempty"`

	// If `true` then the client will require Pushed Authorization Requests.
	// This feature currently must be enabled for your tenant.
	RequirePushedAuthorizationRequests *bool `json:"require_pushed_authorization_requests,omitempty"`

	// URLs that are valid to call back from Auth0 for OIDC backchannel logout.
	// This feature currently must be enabled for your tenant.
	// Deprecated: use OIDCLogout instead of OIDCBackchannelLogout.
	OIDCBackchannelLogout *OIDCBackchannelLogout `json:"oidc_backchannel_logout,omitempty"`

	// URLs that are valid to call back from Auth0 for OIDC logout.
	OIDCLogout *OIDCLogout `json:"oidc_logout,omitempty"`
}

// ClientJWTConfiguration is used to configure JWT settings for our Client.
type ClientJWTConfiguration struct {
	// The amount of seconds the JWT will be valid (affects exp claim)
	LifetimeInSeconds *int `json:"-"`

	// True if the client secret is base64 encoded, false otherwise. Defaults to
	// true
	SecretEncoded *bool `json:"secret_encoded,omitempty"`

	Scopes *map[string]string `json:"scopes,omitempty"`

	// Algorithm used to sign JWTs. Can be "HS256" or "RS256"
	Algorithm *string `json:"alg,omitempty"`
}

// ClientNativeSocialLogin is used to configure Native Social Login for our Client.
type ClientNativeSocialLogin struct {
	// Native Social Login support for the Apple connection
	Apple *ClientNativeSocialLoginSupportEnabled `json:"apple,omitempty"`

	// Native Social Login support for the Facebook connection
	Facebook *ClientNativeSocialLoginSupportEnabled `json:"facebook,omitempty"`
}

// ClientNativeSocialLoginSupportEnabled used to indicate if support is enabled or not.
type ClientNativeSocialLoginSupportEnabled struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// ClientMobile is used to configure mobile app settings.
type ClientMobile struct {
	Android *ClientMobileAndroid `json:"android,omitempty"`
	IOS     *ClientMobileIOS     `json:"ios,omitempty"`
}

// ClientMobileAndroid is used to configure Android app settings.
type ClientMobileAndroid struct {
	AppPackageName *string   `json:"app_package_name,omitempty"`
	KeyHashes      *[]string `json:"sha256_cert_fingerprints,omitempty"`
}

// ClientMobileIOS is used to configure iOS app settings.
type ClientMobileIOS struct {
	TeamID *string `json:"team_id,omitempty"`
	AppID  *string `json:"app_bundle_identifier,omitempty"`
}

// ClientRefreshToken is used to configure the Refresh Token settings for our Client.
type ClientRefreshToken struct {
	// Refresh token rotation type. Can be "rotating" or "non-rotating".
	RotationType *string `json:"rotation_type,omitempty"`

	// Refresh token expiration type. Can be "expiring" or "non-expiring".
	ExpirationType *string `json:"expiration_type,omitempty"`

	// Period in seconds where the previous refresh token can be exchanged
	// without triggering breach detection.
	Leeway *int `json:"leeway,omitempty"`

	// Period in seconds for which refresh tokens will remain valid
	TokenLifetime *int `json:"token_lifetime,omitempty"`

	// Whether the refresh tokens should remain valid indefinitely.
	// If false, "TokenLifetime" should be set.
	InfiniteTokenLifetime *bool `json:"infinite_token_lifetime,omitempty"`

	// Whether the inactive refresh tokens should remain valid indefinitely.
	InfiniteIdleTokenLifetime *bool `json:"infinite_idle_token_lifetime,omitempty"`

	// Period in seconds after which inactive refresh tokens will expire.
	IdleTokenLifetime *int `json:"idle_token_lifetime,omitempty"`
}

// Credential is used to configure Client Credentials.
type Credential struct {
	// The ID of the credential.
	ID *string `json:"id,omitempty"`
	// The name of the credential
	Name *string `json:"name,omitempty"`
	// The key identifier of the credential.
	KeyID *string `json:"kid,omitempty"`
	// The type of credential.
	CredentialType *string `json:"credential_type,omitempty"`
	// PEM-formatted public key or X509 certificate.
	PEM *string `json:"pem,omitempty"`
	// Algorithm which will be used with the credential.
	Algorithm *string `json:"alg,omitempty"`
	// Parse expiry from x509 certificate. If `true`, attempts to parse the expiry date from the provided PEM.
	ParseExpiryFromCert *bool `json:"parse_expiry_from_cert,omitempty"`
	// The time that this credential was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time that this credential was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The time that this credential will expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// ClientAuthenticationMethods defines client authentication method settings for the client.
type ClientAuthenticationMethods struct {
	PrivateKeyJWT *PrivateKeyJWT `json:"private_key_jwt,omitempty"`
}

// PrivateKeyJWT defines the `private_key_jwt` client authentication method settings for the client.
type PrivateKeyJWT struct {
	Credentials *[]Credential `json:"credentials,omitempty"`
}

// OIDCBackchannelLogout defines the `oidc_backchannel_logout` settings for the client.
// Deprecated: use OIDCLogout instead of OIDCBackchannelLogout.
type OIDCBackchannelLogout struct {
	BackChannelLogoutURLs *[]string `json:"backchannel_logout_urls,omitempty"`
}

// OIDCLogout defines the `oidc_logout` settings for the client.
type OIDCLogout struct {
	BackChannelLogoutURLs       *[]string                    `json:"backchannel_logout_urls,omitempty"`
	BackChannelLogoutInitiators *BackChannelLogoutInitiators `json:"backchannel_logout_initiators,omitempty"`
}

// BackChannelLogoutInitiators defines the setting for OIDC logout initiators for a Client.
type BackChannelLogoutInitiators struct {
	Mode               *string   `json:"mode,omitempty"`
	SelectedInitiators *[]string `json:"selected_initiators,omitempty"`
}

// ClientAddons defines the `addons` settings for a Client.
type ClientAddons struct {
	AWS                  *AWSClientAddon                  `json:"aws,omitempty"`
	AzureBlob            *AzureBlobClientAddon            `json:"azure_blob,omitempty"`
	AzureSB              *AzureSBClientAddon              `json:"azure_sb,omitempty"`
	Box                  *BoxClientAddon                  `json:"box,omitempty"`
	CloudBees            *CloudBeesClientAddon            `json:"cloudbees,omitempty"`
	Concur               *ConcurClientAddon               `json:"concur,omitempty"`
	Dropbox              *DropboxClientAddon              `json:"dropbox,omitempty"`
	EchoSign             *EchoSignClientAddon             `json:"echosign,omitempty"`
	Egnyte               *EgnyteClientAddon               `json:"egnyte,omitempty"`
	Firebase             *FirebaseClientAddon             `json:"firebase,omitempty"`
	Layer                *LayerClientAddon                `json:"layer,omitempty"`
	MSCRM                *MSCRMClientAddon                `json:"mscrm,omitempty"`
	NewRelic             *NewRelicClientAddon             `json:"newrelic,omitempty"`
	Office365            *Office365ClientAddon            `json:"office365,omitempty"`
	RMS                  *RMSClientAddon                  `json:"rms,omitempty"`
	Salesforce           *SalesforceClientAddon           `json:"salesforce,omitempty"`
	SalesforceAPI        *SalesforceAPIClientAddon        `json:"salesforce_api,omitempty"`
	SalesforceSandboxAPI *SalesforceSandboxAPIClientAddon `json:"salesforce_sandbox_api,omitempty"`
	// SAML2 Addon configuration. Set this property to indicate that the Addon should be enabled, all settings are optional.
	// The first entry in `Callbacks` should be the URL to post the SAML Token to.
	SAML2          *SAML2ClientAddon          `json:"samlp,omitempty"`
	SAPAPI         *SAPAPIClientAddon         `json:"sap_api,omitempty"`
	SharePoint     *SharePointClientAddon     `json:"sharepoint,omitempty"`
	Sentry         *SentryClientAddon         `json:"sentry,omitempty"`
	Slack          *SlackClientAddon          `json:"slack,omitempty"`
	SpringCM       *SpringCMClientAddon       `json:"springcm,omitempty"`
	SSOIntegration *SSOIntegrationClientAddon `json:"sso_integration,omitempty"`
	WAMS           *WAMSClientAddon           `json:"wams,omitempty"`
	// WS-Fed Addon. Set this property to indicate the Addon should be enabled and then store the
	// configuration in `Callbacks` and `ClientAliases` properties on the Client.
	// The first entry in `Callbacks` should be the URL to post the SAML Token to.
	// ClientAliases should include the Realm which is the identifier sent by the application.
	WSFED   *WSFEDClientAddon   `json:"wsfed,omitempty"`
	Zendesk *ZendeskClientAddon `json:"zendesk,omitempty"`
	Zoom    *ZoomClientAddon    `json:"zoom,omitempty"`
}

// AWSClientAddon defines the `aws` settings for a client.
type AWSClientAddon struct {
	Principal         *string `json:"principal,omitempty"`
	Role              *string `json:"role,omitempty"`
	LifetimeInSeconds *int    `json:"lifetime_in_seconds,omitempty"`
}

// AzureBlobClientAddon defines the `azure_blob` settings for a client.
type AzureBlobClientAddon struct {
	AccountName      *string `json:"accountName,omitempty"`
	StorageAccessKey *string `json:"storageAccessKey,omitempty"`
	ContainerName    *string `json:"containerName,omitempty"`
	BlobName         *string `json:"blobName,omitempty"`
	Expiration       *int    `json:"expiration,omitempty"`
	SignedIdentifier *string `json:"signedIdentifier,omitempty"`
	BlobRead         *bool   `json:"blob_read,omitempty"`
	BlobWrite        *bool   `json:"blob_write,omitempty"`
	BlobDelete       *bool   `json:"blob_delete,omitempty"`
	ContainerRead    *bool   `json:"container_read,omitempty"`
	ContainerWrite   *bool   `json:"container_write,omitempty"`
	ContainerDelete  *bool   `json:"container_delete,omitempty"`
	ContainerList    *bool   `json:"container_list,omitempty"`
}

// AzureSBClientAddon defines the `azure_sb` settings for a client.
type AzureSBClientAddon struct {
	Namespace  *string `json:"namespace,omitempty"`
	SASKeyName *string `json:"sasKeyName,omitempty"`
	SASKey     *string `json:"sasKey,omitempty"`
	EntityPath *string `json:"entityPath,omitempty"`
	Expiration *int    `json:"expiration,omitempty"`
}

// BoxClientAddon defines the `box` settings for a client.
type BoxClientAddon struct {
}

// CloudBeesClientAddon defines the `cloudbees` settings for a client.
type CloudBeesClientAddon struct {
}

// ConcurClientAddon defines the `concur` settings for a client.
type ConcurClientAddon struct {
}

// DropboxClientAddon defines the `dropbox` settings for a client.
type DropboxClientAddon struct {
}

// EchoSignClientAddon defines the `echosign` settings for a client.
type EchoSignClientAddon struct {
	Domain *string `json:"domain,omitempty"`
}

// EgnyteClientAddon defines the `egnyte` settings for a client.
type EgnyteClientAddon struct {
	Domain *string `json:"domain,omitempty"`
}

// FirebaseClientAddon defines the `firebase` settings for a client.
type FirebaseClientAddon struct {
	Secret            *string `json:"secret,omitempty"`
	PrivateKeyID      *string `json:"private_key_id,omitempty"`
	PrivateKey        *string `json:"private_key,omitempty"`
	ClientEmail       *string `json:"client_email,omitempty"`
	LifetimeInSeconds *int    `json:"lifetime_in_seconds,omitempty"`
}

// LayerClientAddon defines the `layer` settings for a client.
type LayerClientAddon struct {
	ProviderID *string `json:"providerId,omitempty"`
	KeyID      *string `json:"keyId,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
	Principal  *string `json:"principal,omitempty"`
	Expiration *int    `json:"expiration,omitempty"`
}

// MSCRMClientAddon defines the `mscrm` settings for a client.
type MSCRMClientAddon struct {
	URL *string `json:"url,omitempty"`
}

// NewRelicClientAddon defines the `newrelic` settings for a client.
type NewRelicClientAddon struct {
	Account *string `json:"account,omitempty"`
}

// Office365ClientAddon defines the `office365` settings for a client.
type Office365ClientAddon struct {
	Domain     *string `json:"domain,omitempty"`
	Connection *string `json:"connection,omitempty"`
}

// RMSClientAddon defines the `rms` settings for a client.
type RMSClientAddon struct {
	URL *string `json:"url,omitempty"`
}

// SalesforceClientAddon defines the `salesforce` settings for a client.
type SalesforceClientAddon struct {
	EntityID *string `json:"entity_id,omitempty"`
}

// SalesforceAPIClientAddon defines the `salesforce_api` settings for a client.
type SalesforceAPIClientAddon struct {
	ClientID            *string `json:"clientid,omitempty"`
	Principal           *string `json:"principal,omitempty"`
	CommunityName       *string `json:"communityName,omitempty"`
	CommunityURLSection *string `json:"community_url_section,omitempty"`
}

// SalesforceSandboxAPIClientAddon defines the `salesforce_sandbox_api` settings for a client.
type SalesforceSandboxAPIClientAddon struct {
	ClientID            *string `json:"clientid,omitempty"`
	Principal           *string `json:"principal,omitempty"`
	CommunityName       *string `json:"communityName,omitempty"`
	CommunityURLSection *string `json:"community_url_section,omitempty"`
}

// SAML2ClientAddon defines the `SAML2` settings for a Client.
type SAML2ClientAddon struct {
	// The mappings between the Auth0 user profile and the output attributes on the SAML Assertion.
	// Each "name" represents the property name on the Auth0 user profile.
	// Each "value" is the name (including namespace) for the resulting SAML attribute in the assertion.
	Mappings *map[string]string `json:"mappings,omitempty"`
	// The audience of the SAML Assertion.
	Audience *string `json:"audience,omitempty"`
	// The recipient of the SAML Assertion.
	Recipient *string `json:"recipient,omitempty"`
	// Whether or not a UPN claim should be created.
	CreateUPNClaim *bool `json:"createUpnClaim,omitempty"`
	// If `PassthroughClaimsWithNoMapping` is true and this is false, for each claim that is not mapped to the common profile Auth0 will add a prefix
	// 	http://schema.auth0.com	. If true it will passthrough the claim as-is.
	MapUnknownClaimsAsIs *bool `json:"mapUnknownClaimsAsIs,omitempty"`
	// If true, for each claim that is not mapped to the common profile, Auth0 will passthrough those in the output assertion.
	// If false, those claims won't be mapped.
	PassthroughClaimsWithNoMapping *bool `json:"passthroughClaimsWithNoMapping,omitempty"`
	// If true, it will will add more information in the token like the provider used (google, adfs, ad, etc.) and the access_token if available.
	MapIdentities *bool `json:"mapIdentities,omitempty"`
	// Signature algorithm to sign the SAML Assertion or response.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	// Digest algorithm to calculate digest of the SAML Assertion or response.
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty"`
	Issuer          *string `json:"issuer,omitempty"`
	// Destination of the SAML Response. If not specified, it will be AssertionConsumerUrlof SAMLRequest or Callback URL if there was no SAMLRequest.
	Destination *string `json:"destination,omitempty"`
	// Expiration of the token.
	LifetimeInSeconds *int `json:"lifetimeInSeconds,omitempty"`
	// Whether or not the SAML Response should be signed. By default the SAML Assertion will be signed, but not the SAML Response.
	// If true, SAML Response will be signed instead of SAML Assertion.
	SignResponse         *bool   `json:"signResponse,omitempty"`
	NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`
	// Auth0 will try each of the attributes of this array in order. If one of them has a value, it will use that for the Subject/NameID
	NameIdentifierProbes *[]string `json:"nameIdentifierProbes,omitempty"`
	AuthnContextClassRef *string   `json:"authnContextClassRef,omitempty"`
	// When set to true, the xs:type of the element is inferred. Types are xs:string, xs:boolean, xs:double, and xs:anyType.
	// When set to false all xs:type are xs:anyType
	TypedAttributes *bool `json:"typedAttributes,omitempty"`
	// When set to true, the NameFormat is inferred based on the attribute name.
	// NameFormat values are urn:oasis:names:tc:SAML:2.0:attrname-format:uri, urn:oasis:names:tc:SAML:2.0:attrname-format:basic,
	// and urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified.
	// If set to false, the attribute NameFormat is not set in the assertion.
	IncludeAttributeNameFormat *bool `json:"includeAttributeNameFormat,omitempty"`
	// Indicates the protocol binding used for SAML logout responses.
	// By default Auth0 uses HTTP-POST, but you can switch to HTTP-Redirect by setting to `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect`.
	Binding *string `json:"binding,omitempty"`
	// Optionally indicates the public key certificate used to validate SAML requests. If set, SAML requests will be required to be signed.
	SigningCert *string `json:"signingCert,omitempty"`
	//  An object that controls SAML logout behavior.
	Logout *SAML2ClientAddonLogout `json:"logout,omitempty"`
}

// SAML2ClientAddonLogout defines the `logout` settings for the SAML2Addon.
type SAML2ClientAddonLogout struct {
	// The service provider (client application)'s Single Logout Service URL, where Auth0 will send logout requests and responses
	Callback *string `json:"callback,omitempty"`
	// Controls whether Auth0 should notify service providers of session termination
	SLOEnabled *bool `json:"slo_enabled,omitempty"`
}

// SAPAPIClientAddon defines the `sap` settings for a client.
type SAPAPIClientAddon struct {
	ClientID             *string `json:"clientid,omitempty"`
	UsernameAttribute    *string `json:"usernameAttribute,omitempty"`
	TokenEndpointURL     *string `json:"tokenEndpointUrl,omitempty"`
	Scope                *string `json:"scope,omitempty"`
	ServicePassword      *string `json:"servicePassword,omitempty"`
	NameIdentifierFormat *string `json:"nameIdentifierFormat,omitempty"`
}

// SentryClientAddon defines the `sentry` settings for a client.
type SentryClientAddon struct {
	OrgSlug *string `json:"org_slug,omitempty"`
	BaseURL *string `json:"base_url,omitempty"`
}

// SharePointClientAddon defines the `sharepoint` settings for a client.
type SharePointClientAddon struct {
	URL         *string   `json:"url,omitempty"`
	ExternalURL *[]string `json:"external_url,omitempty"`
}

// SlackClientAddon defines the `slack` settings for a client.
type SlackClientAddon struct {
	Team *string `json:"team,omitempty"`
}

// SpringCMClientAddon defines the `springcm` settings for a client.
type SpringCMClientAddon struct {
	ACSURL *string `json:"acsurl,omitempty"`
}

// SSOIntegrationClientAddon defines the `sso_integration` settings for a client.
type SSOIntegrationClientAddon struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// WAMSClientAddon defines the `wams` settings for a client.
type WAMSClientAddon struct {
	Masterkey *string `json:"masterkey,omitempty"`
}

// ZendeskClientAddon defines the `zendesk` settings for a client.
type ZendeskClientAddon struct {
	AccountName *string `json:"accountName,omitempty"`
}

// ZoomClientAddon defines the `zoom` settings for a client.
type ZoomClientAddon struct {
	Account *string `json:"account,omitempty"`
}

// WSFEDClientAddon is an empty struct used to indicate that the WS-FED Addon should be enabled.
// Configuration for this Addon is stored in the `Callbacks` and `ClientAliases` properties on the Client.
type WSFEDClientAddon struct {
}

// ClientList is a list of Clients.
type ClientList struct {
	List
	Clients []*Client `json:"clients"`
}

// ClientManager manages Auth0 Client resources.
type ClientManager manager

// Create a new client application.
//
// See: https://auth0.com/docs/api/management/v2#!/Clients/post_clients
func (m *ClientManager) Create(ctx context.Context, c *Client, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("clients"), c, opts...)
}

// Read a client by its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Clients/get_clients_by_id
func (m *ClientManager) Read(ctx context.Context, id string, opts ...RequestOption) (c *Client, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("clients", id), &c, opts...)
	return
}

// List client applications.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Clients/get_clients
func (m *ClientManager) List(ctx context.Context, opts ...RequestOption) (c *ClientList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("clients"), &c, applyListDefaults(opts))
	return
}

// Update a client.
//
// See: https://auth0.com/docs/api/management/v2#!/Clients/patch_clients_by_id
func (m *ClientManager) Update(ctx context.Context, id string, c *Client, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("clients", id), c, opts...)
}

// RotateSecret rotates a client secret.
//
// See: https://auth0.com/docs/api/management/v2#!/Clients/post_rotate_secret
func (m *ClientManager) RotateSecret(ctx context.Context, id string, opts ...RequestOption) (c *Client, err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("clients", id, "rotate-secret"), &c, opts...)
	return
}

// Delete a client and all its related assets (like rules, connections, etc)
// given its ID.
//
// See: https://auth0.com/docs/api/management/v2#!/Clients/delete_clients_by_id
func (m *ClientManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("clients", id), nil, opts...)
}

// CreateCredential creates a client application's client credential.
func (m *ClientManager) CreateCredential(ctx context.Context, clientID string, credential *Credential, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("clients", clientID, "credentials"), credential, opts...)
}

// UpdateCredential updates a client application's client credential expiry.
func (m *ClientManager) UpdateCredential(ctx context.Context, clientID, credentialID string, credential *Credential, opts ...RequestOption) error {
	credentialClone := &Credential{ExpiresAt: credential.ExpiresAt} // The API only accepts the expires_at property.

	err := m.management.Request(ctx, "PATCH", m.management.URI("clients", clientID, "credentials", credentialID), credentialClone, opts...)
	if err != nil {
		return err
	}

	credential.ID = credentialClone.ID
	credential.Name = credentialClone.Name
	credential.CredentialType = credentialClone.CredentialType
	credential.KeyID = credentialClone.KeyID
	credential.Algorithm = credentialClone.Algorithm
	credential.CreatedAt = credentialClone.CreatedAt
	credential.UpdatedAt = credentialClone.UpdatedAt
	credential.ExpiresAt = credentialClone.ExpiresAt
	// PEM and ParseExpiryFromCert don't get returned.

	return nil
}

// ListCredentials lists client credentials associated with the client application.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
func (m *ClientManager) ListCredentials(ctx context.Context, clientID string, opts ...RequestOption) (c []*Credential, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("clients", clientID, "credentials"), &c, applyListDefaults(opts))
	return
}

// GetCredential gets a client credentials object.
func (m *ClientManager) GetCredential(ctx context.Context, clientID string, credentialID string, opts ...RequestOption) (c *Credential, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("clients", clientID, "credentials", credentialID), &c, opts...)
	return
}

// DeleteCredential deletes a client credentials object.
func (m *ClientManager) DeleteCredential(ctx context.Context, clientID string, credentialID string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("clients", clientID, "credentials", credentialID), nil, opts...)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the json field lifetime_in_seconds, which can either
// be an int, or a string in older tenants.
func (jc *ClientJWTConfiguration) UnmarshalJSON(b []byte) error {
	type clientJWTConfiguration ClientJWTConfiguration
	type clientJWTConfigurationWrapper struct {
		*clientJWTConfiguration
		RawLifetimeInSeconds interface{} `json:"lifetime_in_seconds,omitempty"`
	}

	alias := &clientJWTConfigurationWrapper{(*clientJWTConfiguration)(jc), nil}

	err := json.Unmarshal(b, alias)
	if err != nil {
		return err
	}

	unexpectedTypeError := fmt.Errorf("unexpected type for field lifetime_in_seconds")

	if alias.RawLifetimeInSeconds != nil {
		switch rawLifetimeInSeconds := alias.RawLifetimeInSeconds.(type) {
		case float64:
			jc.LifetimeInSeconds = auth0.Int(int(rawLifetimeInSeconds))
		case string:
			value, err := strconv.Atoi(rawLifetimeInSeconds)
			if err != nil {
				return unexpectedTypeError
			}
			jc.LifetimeInSeconds = &value
		default:
			return unexpectedTypeError
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (jc *ClientJWTConfiguration) MarshalJSON() ([]byte, error) {
	type clientJWTConfiguration ClientJWTConfiguration
	type clientJWTConfigurationWrapper struct {
		*clientJWTConfiguration
		RawLifetimeInSeconds interface{} `json:"lifetime_in_seconds,omitempty"`
	}

	alias := &clientJWTConfigurationWrapper{(*clientJWTConfiguration)(jc), nil}
	if jc.LifetimeInSeconds != nil {
		alias.RawLifetimeInSeconds = jc.LifetimeInSeconds
	}

	return json.Marshal(alias)
}
