package management

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/auth0/go-auth0/internal/tag"
)

const (
	// ConnectionStrategyAuth0 constant.
	ConnectionStrategyAuth0 = "auth0"
	// ConnectionStrategyOkta constant.
	ConnectionStrategyOkta = "okta"
	// ConnectionStrategyGoogleOAuth2 constant.
	ConnectionStrategyGoogleOAuth2 = "google-oauth2"
	// ConnectionStrategyFacebook constant.
	ConnectionStrategyFacebook = "facebook"
	// ConnectionStrategyApple constant.
	ConnectionStrategyApple = "apple"
	// ConnectionStrategyLinkedin constant.
	ConnectionStrategyLinkedin = "linkedin"
	// ConnectionStrategyGitHub constant.
	ConnectionStrategyGitHub = "github"
	// ConnectionStrategyWindowsLive constant.
	ConnectionStrategyWindowsLive = "windowslive"
	// ConnectionStrategySalesforce constant.
	ConnectionStrategySalesforce = "salesforce"
	// ConnectionStrategySalesforceCommunity constant.
	ConnectionStrategySalesforceCommunity = "salesforce-community"
	// ConnectionStrategySalesforceSandbox constant.
	ConnectionStrategySalesforceSandbox = "salesforce-sandbox"
	// ConnectionStrategyEmail constant.
	ConnectionStrategyEmail = "email"
	// ConnectionStrategySMS constant.
	ConnectionStrategySMS = "sms"
	// ConnectionStrategyOIDC constant.
	ConnectionStrategyOIDC = "oidc"
	// ConnectionStrategyOAuth2 constant.
	ConnectionStrategyOAuth2 = "oauth2"
	// ConnectionStrategyAD constant.
	ConnectionStrategyAD = "ad"
	// ConnectionStrategyADFS constant.
	ConnectionStrategyADFS = "adfs"
	// ConnectionStrategyAzureAD constant.
	ConnectionStrategyAzureAD = "waad"
	// ConnectionStrategySAML constant.
	ConnectionStrategySAML = "samlp"
	// ConnectionStrategyGoogleApps constant.
	ConnectionStrategyGoogleApps = "google-apps"
	// ConnectionStrategyDropbox constant.
	ConnectionStrategyDropbox = "dropbox"
	// ConnectionStrategyBitBucket constant.
	ConnectionStrategyBitBucket = "bitbucket"
	// ConnectionStrategyPaypal constant.
	ConnectionStrategyPaypal = "paypal"
	// ConnectionStrategyTwitter constant.
	ConnectionStrategyTwitter = "twitter"
	// ConnectionStrategyAmazon constant.
	ConnectionStrategyAmazon = "amazon"
	// ConnectionStrategyYahoo constant.
	ConnectionStrategyYahoo = "yahoo"
	// ConnectionStrategyBox constant.
	ConnectionStrategyBox = "box"
	// ConnectionStrategyWordpress constant.
	ConnectionStrategyWordpress = "wordpress"
	// ConnectionStrategyShopify constant.
	ConnectionStrategyShopify = "shopify"
	// ConnectionStrategyCustom constant.
	ConnectionStrategyCustom = "custom"
	// ConnectionStrategyPingFederate constant.
	ConnectionStrategyPingFederate = "pingfederate"
	// ConnectionStrategyLine constant.
	ConnectionStrategyLine = "line"
)

var (
	// Enforcing the ConnectionOptionsScoper interface.
	_ ConnectionOptionsScoper = &ConnectionOptionsOkta{}
	_ ConnectionOptionsScoper = &ConnectionOptionsGoogleOAuth2{}
	_ ConnectionOptionsScoper = &ConnectionOptionsFacebook{}
	_ ConnectionOptionsScoper = &ConnectionOptionsApple{}
	_ ConnectionOptionsScoper = &ConnectionOptionsLinkedin{}
	_ ConnectionOptionsScoper = &ConnectionOptionsGitHub{}
	_ ConnectionOptionsScoper = &ConnectionOptionsWindowsLive{}
	_ ConnectionOptionsScoper = &ConnectionOptionsSalesforce{}
	_ ConnectionOptionsScoper = &ConnectionOptionsOIDC{}
	_ ConnectionOptionsScoper = &ConnectionOptionsOAuth2{}
	_ ConnectionOptionsScoper = &ConnectionOptionsAzureAD{}
	_ ConnectionOptionsScoper = &ConnectionOptionsPingFederate{}
	_ ConnectionOptionsScoper = &ConnectionOptionsSAML{}
	_ ConnectionOptionsScoper = &ConnectionOptionsGoogleApps{}
)

// ConnectionOptionsScoper is used to enforce being able to read and
// set scopes through the scope tag on connection options properties.
type ConnectionOptionsScoper interface {
	Scopes() []string
	SetScopes(enable bool, scopes ...string)
}

// Connection is the relationship between Auth0 and a source of users.
//
// See: https://auth0.com/docs/authenticate/identity-providers
type Connection struct {
	// A generated string identifying the connection.
	ID *string `json:"id,omitempty"`

	// The name of the connection. Must start and end with an alphanumeric
	// character and can only contain alphanumeric characters and '-'. Max
	// length 128.
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`

	// The identity provider identifier for the connection. Can be any of the
	// following:
	//
	// "ad", "adfs", "amazon", "dropbox", "bitbucket", "aol", "auth0-adldap",
	// "auth0-oidc", "auth0", "baidu", "bitly", "box", "custom", "daccount",
	// "dwolla", "email", "evernote-sandbox", "evernote", "exact", "facebook",
	// "fitbit", "flickr", "github", "google-apps", "google-oauth2", "guardian",
	//  "instagram", "ip", "linkedin", "miicard", "oauth1", "oauth2",
	// "office365", "paypal", "paypal-sandbox", "pingfederate",
	// "planningcenter", "renren", "salesforce-community", "salesforce-sandbox",
	//  "salesforce", "samlp", "sharepoint", "shopify", "sms", "soundcloud",
	// "thecity-sandbox", "thecity", "thirtysevensignals", "twitter", "untappd",
	//  "vkontakte", "waad", "weibo", "windowslive", "wordpress", "yahoo", "line",
	// "yammer", "okta" or "yandex".
	Strategy *string `json:"strategy,omitempty"`

	// True if the connection is domain level
	IsDomainConnection *bool `json:"is_domain_connection,omitempty"`

	// Options for validation.
	Options interface{} `json:"-"`

	// The identifiers of the clients for which the connection is to be
	// enabled. If the array is empty or the property is not specified, no
	// clients are enabled.
	EnabledClients *[]string `json:"enabled_clients,omitempty"`

	// Defines the realms for which the connection will be used (ie: email
	// domains). If the array is empty or the property is not specified, the
	// connection name will be added as realm.
	Realms *[]string `json:"realms,omitempty"`

	Metadata *map[string]string `json:"metadata,omitempty"`

	// Provisioning Ticket URL is Ticket URL for Active Directory/LDAP, etc.
	ProvisioningTicketURL *string `json:"provisioning_ticket_url,omitempty"`

	// Display connection as a button.
	ShowAsButton *bool `json:"show_as_button,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
func (c *Connection) MarshalJSON() ([]byte, error) {
	type connection Connection
	type connectionWrapper struct {
		*connection
		RawOptions json.RawMessage `json:"options,omitempty"`
	}

	w := &connectionWrapper{(*connection)(c), nil}

	if c.Options != nil {
		b, err := json.Marshal(c.Options)
		if err != nil {
			return nil, err
		}
		w.RawOptions = b
	}

	return json.Marshal(w)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Connection) UnmarshalJSON(b []byte) error {
	type connection Connection
	type connectionWrapper struct {
		*connection
		RawOptions json.RawMessage `json:"options,omitempty"`
	}

	w := &connectionWrapper{(*connection)(c), nil}

	err := json.Unmarshal(b, w)
	if err != nil {
		return err
	}

	if c.Strategy != nil {
		var v interface{}

		switch *c.Strategy {
		case ConnectionStrategyAuth0:
			v = &ConnectionOptions{}
		case ConnectionStrategyOkta:
			v = &ConnectionOptionsOkta{}
		case ConnectionStrategyGoogleOAuth2:
			v = &ConnectionOptionsGoogleOAuth2{}
		case ConnectionStrategyFacebook:
			v = &ConnectionOptionsFacebook{}
		case ConnectionStrategyApple:
			v = &ConnectionOptionsApple{}
		case ConnectionStrategyLinkedin:
			v = &ConnectionOptionsLinkedin{}
		case ConnectionStrategyGitHub:
			v = &ConnectionOptionsGitHub{}
		case ConnectionStrategyWindowsLive:
			v = &ConnectionOptionsWindowsLive{}
		case ConnectionStrategySalesforce,
			ConnectionStrategySalesforceCommunity,
			ConnectionStrategySalesforceSandbox:
			v = &ConnectionOptionsSalesforce{}
		case ConnectionStrategyEmail:
			v = &ConnectionOptionsEmail{}
		case ConnectionStrategySMS:
			v = &ConnectionOptionsSMS{}
		case ConnectionStrategyOIDC:
			v = &ConnectionOptionsOIDC{}
		case ConnectionStrategyOAuth2,
			ConnectionStrategyDropbox,
			ConnectionStrategyBitBucket,
			ConnectionStrategyPaypal,
			ConnectionStrategyTwitter,
			ConnectionStrategyAmazon,
			ConnectionStrategyYahoo,
			ConnectionStrategyBox,
			ConnectionStrategyWordpress,
			ConnectionStrategyShopify,
			ConnectionStrategyLine,
			ConnectionStrategyCustom:
			v = &ConnectionOptionsOAuth2{}
		case ConnectionStrategyAD:
			v = &ConnectionOptionsAD{}
		case ConnectionStrategyAzureAD:
			v = &ConnectionOptionsAzureAD{}
		case ConnectionStrategyADFS:
			v = &ConnectionOptionsADFS{}
		case ConnectionStrategyPingFederate:
			v = &ConnectionOptionsPingFederate{}
		case ConnectionStrategySAML:
			v = &ConnectionOptionsSAML{}
		case ConnectionStrategyGoogleApps:
			v = &ConnectionOptionsGoogleApps{}
		default:
			v = make(map[string]interface{})
		}

		if w.RawOptions != nil {
			err = json.Unmarshal(w.RawOptions, &v)
			if err != nil {
				return err
			}
		}

		c.Options = v
	}

	return nil
}

// ConnectionOptions is used to configure Connections.
type ConnectionOptions struct {
	// Options for multifactor authentication. Can be used to set active and
	// return_enroll_settings.
	MFA map[string]interface{} `json:"mfa,omitempty"`

	// Options for validation.
	Validation map[string]interface{} `json:"validation,omitempty"`

	// Password strength level, can be one of:
	// "none", "low", "fair", "good", "excellent" or null.
	PasswordPolicy *string `json:"passwordPolicy,omitempty"`

	// Options for password history policy.
	PasswordHistory map[string]interface{} `json:"password_history,omitempty"`

	// Options for password expiration policy.
	PasswordNoPersonalInfo map[string]interface{} `json:"password_no_personal_info,omitempty"`

	// Options for password dictionary policy.
	PasswordDictionary map[string]interface{} `json:"password_dictionary,omitempty"`

	// Options for password complexity options.
	PasswordComplexityOptions map[string]interface{} `json:"password_complexity_options,omitempty"`

	// Set to true to inject context into custom DB scripts (warning: cannot be disabled once enabled).
	EnableScriptContext *bool `json:"enable_script_context,omitempty"`

	// Set to true to use a legacy user store.
	EnabledDatabaseCustomization *bool `json:"enabledDatabaseCustomization,omitempty"`

	BruteForceProtection *bool `json:"brute_force_protection,omitempty"`

	ImportMode *bool `json:"import_mode,omitempty"`

	DisableSignup *bool `json:"disable_signup,omitempty"`

	RequiresUsername *bool `json:"requires_username,omitempty"`

	// Scripts for the connection.
	// Allowed keys are: "get_user", "login", "create", "verify", "change_password", "delete"".
	CustomScripts *map[string]string `json:"customScripts,omitempty"`

	// Configuration variables that can be used in custom scripts.
	Configuration *map[string]string `json:"configuration,omitempty"`

	StrategyVersion *int `json:"strategy_version,omitempty"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`

	// Set to true to stop the "Forgot Password" being displayed on login pages
	DisableSelfServiceChangePassword *bool `json:"disable_self_service_change_password,omitempty"`

	// Options for enabling authentication methods.
	AuthenticationMethods *AuthenticationMethods `json:"authentication_methods,omitempty"`

	// Options for the passkey authentication method.
	PasskeyOptions *PasskeyOptions `json:"passkey_options,omitempty"`
}

// AuthenticationMethods represents the options for enabling authentication methods for the connection.
type AuthenticationMethods struct {
	Password *PasswordAuthenticationMethod `json:"password,omitempty"`
	Passkey  *PasskeyAuthenticationMethod  `json:"passkey,omitempty"`
}

// PasswordAuthenticationMethod represents password authentication enablement for the connection.
type PasswordAuthenticationMethod struct {
	// Determines whether passwords are enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// PasskeyAuthenticationMethod represents passkey authentication enablement for the connection.
type PasskeyAuthenticationMethod struct {
	// Determines whether passkeys are enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// PasskeyOptions contains Passkey configuration for the connection.
type PasskeyOptions struct {
	// Controls the UI used to challenge the user for their passkey. Should be one of "both", "autofill", or "button".
	ChallengeUI *string `json:"challenge_ui,omitempty"`
	// Enables or disables progressive enrollment of passkeys for the connection.
	ProgressiveEnrollmentEnabled *bool `json:"progressive_enrollment_enabled,omitempty"`
	// Enables or disables enrollment prompt for local passkey when user authenticates using a cross-device passkey for the connection.
	LocalEnrollmentEnabled *bool `json:"local_enrollment_enabled,omitempty"`
}

// ConnectionOptionsOkta is used to configure an Okta Workforce Connection.
type ConnectionOptionsOkta struct {
	ClientID              *string                `json:"client_id,omitempty"`
	ClientSecret          *string                `json:"client_secret,omitempty"`
	Domain                *string                `json:"domain,omitempty"`
	DomainAliases         *[]string              `json:"domain_aliases,omitempty"`
	LogoURL               *string                `json:"icon_url,omitempty"`
	AuthorizationEndpoint *string                `json:"authorization_endpoint"`
	Issuer                *string                `json:"issuer"`
	JWKSURI               *string                `json:"jwks_uri"`
	UserInfoEndpoint      *string                `json:"userinfo_endpoint"`
	TokenEndpoint         *string                `json:"token_endpoint"`
	Scope                 *string                `json:"scope,omitempty"`
	SetUserAttributes     *string                `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs    *[]string              `json:"non_persistent_attrs,omitempty"`
	UpstreamParams        map[string]interface{} `json:"upstream_params,omitempty"`

	ConnectionSettings *ConnectionOptionsOIDCConnectionSettings `json:"connection_settings,omitempty"`
	AttributeMap       *ConnectionOptionsOIDCAttributeMap       `json:"attribute_map,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsOkta.
func (c *ConnectionOptionsOkta) Scopes() []string {
	return strings.Fields(c.GetScope())
}

// SetScopes sets the scopes for ConnectionOptionsOkta.
func (c *ConnectionOptionsOkta) SetScopes(enable bool, scopes ...string) {
	scopeMap := make(map[string]bool)
	for _, scope := range c.Scopes() {
		scopeMap[scope] = true
	}
	for _, scope := range scopes {
		scopeMap[scope] = enable
	}
	scopeSlice := make([]string, 0, len(scopeMap))
	for scope, enabled := range scopeMap {
		if enabled {
			scopeSlice = append(scopeSlice, scope)
		}
	}
	sort.Strings(scopeSlice)
	scope := strings.Join(scopeSlice, " ")
	c.Scope = &scope
}

// ConnectionOptionsGoogleOAuth2 is used to configure a GoogleOAuth2 Connection.
type ConnectionOptionsGoogleOAuth2 struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	AllowedAudiences *[]string `json:"-"`

	Email                  *bool         `json:"email,omitempty" scope:"email"`
	Profile                *bool         `json:"profile,omitempty" scope:"profile"`
	Contacts               *bool         `json:"contacts,omitempty" scope:"contacts"`
	Blogger                *bool         `json:"blogger,omitempty" scope:"blogger"`
	Calendar               *bool         `json:"calendar,omitempty" scope:"calendar"`
	Gmail                  *bool         `json:"gmail,omitempty" scope:"gmail"`
	GooglePlus             *bool         `json:"google_plus,omitempty" scope:"google_plus"`
	Orkut                  *bool         `json:"orkut,omitempty" scope:"orkut"`
	PicasaWeb              *bool         `json:"picasa_web,omitempty" scope:"picasa_web"`
	Tasks                  *bool         `json:"tasks,omitempty" scope:"tasks"`
	Youtube                *bool         `json:"youtube,omitempty" scope:"youtube"`
	AdsenseManagement      *bool         `json:"adsense_management,omitempty" scope:"adsense_management"`
	GoogleAffiliateNetwork *bool         `json:"google_affiliate_network,omitempty" scope:"google_affiliate_network"`
	Analytics              *bool         `json:"analytics,omitempty" scope:"analytics"`
	GoogleBooks            *bool         `json:"google_books,omitempty" scope:"google_books"`
	GoogleCloudStorage     *bool         `json:"google_cloud_storage,omitempty" scope:"google_cloud_storage"`
	ContentAPIForShopping  *bool         `json:"content_api_for_shopping,omitempty" scope:"content_api_for_shopping"`
	ChromeWebStore         *bool         `json:"chrome_web_store,omitempty" scope:"chrome_web_store"`
	DocumentList           *bool         `json:"document_list,omitempty" scope:"document_list"`
	GoogleDrive            *bool         `json:"google_drive,omitempty" scope:"google_drive"`
	GoogleDriveFiles       *bool         `json:"google_drive_files,omitempty" scope:"google_drive_files"`
	LatitudeBest           *bool         `json:"latitude_best,omitempty" scope:"latitude_best"`
	LatitudeCity           *bool         `json:"latitude_city,omitempty" scope:"latitude_city"`
	Moderator              *bool         `json:"moderator,omitempty" scope:"moderator"`
	Sites                  *bool         `json:"sites,omitempty" scope:"sites"`
	Spreadsheets           *bool         `json:"spreadsheets,omitempty" scope:"spreadsheets"`
	URLShortener           *bool         `json:"url_shortener,omitempty" scope:"url_shortener"`
	WebmasterTools         *bool         `json:"webmaster_tools,omitempty" scope:"webmaster_tools"`
	Coordinate             *bool         `json:"coordinate,omitempty" scope:"coordinate"`
	CoordinateReadonly     *bool         `json:"coordinate_readonly,omitempty" scope:"coordinate_readonly"`
	SetUserAttributes      *string       `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs     *[]string     `json:"non_persistent_attrs,omitempty"`
	Scope                  []interface{} `json:"scope,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsGoogleOAuth2.
func (c *ConnectionOptionsGoogleOAuth2) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsGoogleOAuth2.
func (c *ConnectionOptionsGoogleOAuth2) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle differences in the allowed_audiences field, which can
// be an array of strings or a single string.
func (c *ConnectionOptionsGoogleOAuth2) UnmarshalJSON(data []byte) error {
	type connectionOptionsGoogleOAuth2 ConnectionOptionsGoogleOAuth2
	type connectionOptionsGoogleOAuth2Wrapper struct {
		*connectionOptionsGoogleOAuth2
		RawAllowedAudiences interface{} `json:"allowed_audiences,omitempty"`
	}

	alias := &connectionOptionsGoogleOAuth2Wrapper{(*connectionOptionsGoogleOAuth2)(c), nil}

	err := json.Unmarshal(data, alias)
	if err != nil {
		return err
	}

	if alias.RawAllowedAudiences != nil {
		switch rawAllowedAudiences := alias.RawAllowedAudiences.(type) {
		case []interface{}:
			audiences := make([]string, len(rawAllowedAudiences))
			for i, v := range rawAllowedAudiences {
				audiences[i] = v.(string)
			}

			c.AllowedAudiences = &audiences
		case string:
			c.AllowedAudiences = &[]string{}
		default:
			return fmt.Errorf("unexpected type for field allowed_audiences: %T", alias.RawAllowedAudiences)
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (c *ConnectionOptionsGoogleOAuth2) MarshalJSON() ([]byte, error) {
	type connectionOptionsGoogleOAuth2 ConnectionOptionsGoogleOAuth2
	type connectionOptionsGoogleOAuth2Wrapper struct {
		*connectionOptionsGoogleOAuth2
		RawAllowedAudiences interface{} `json:"allowed_audiences,omitempty"`
	}

	alias := &connectionOptionsGoogleOAuth2Wrapper{(*connectionOptionsGoogleOAuth2)(c), nil}
	if c.AllowedAudiences != nil {
		alias.RawAllowedAudiences = c.AllowedAudiences
	}

	return json.Marshal(alias)
}

// ConnectionOptionsFacebook is used to configure a Facebook Connection.
type ConnectionOptionsFacebook struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	AllowContextProfileField    *bool     `json:"allow_context_profile_field,omitempty"`
	Email                       *bool     `json:"email,omitempty" scope:"email"`
	GroupsAccessMemberInfo      *bool     `json:"groups_access_member_info,omitempty" scope:"groups_access_member_info"`
	PublishToGroups             *bool     `json:"publish_to_groups,omitempty" scope:"publish_to_groups"`
	UserAgeRange                *bool     `json:"user_age_range,omitempty" scope:"user_age_range"`
	UserBirthday                *bool     `json:"user_birthday,omitempty" scope:"user_birthday"`
	AdsManagement               *bool     `json:"ads_management,omitempty" scope:"ads_management"`
	AdsRead                     *bool     `json:"ads_read,omitempty" scope:"ads_read"`
	ReadAudienceNetworkInsights *bool     `json:"read_audience_network_insights,omitempty" scope:"read_audience_network_insights"`
	ReadInsights                *bool     `json:"read_insights,omitempty" scope:"read_insights"`
	ManageNotifications         *bool     `json:"manage_notifications,omitempty" scope:"manage_notifications"`
	PublishActions              *bool     `json:"publish_actions,omitempty" scope:"publish_actions"`
	ReadMailbox                 *bool     `json:"read_mailbox,omitempty" scope:"read_mailbox"`
	PublicProfile               *bool     `json:"public_profile,omitempty" scope:"public_profile"`
	UserEvents                  *bool     `json:"user_events,omitempty" scope:"user_events"`
	UserFriends                 *bool     `json:"user_friends,omitempty" scope:"user_friends"`
	UserGender                  *bool     `json:"user_gender,omitempty" scope:"user_gender"`
	UserHometown                *bool     `json:"user_hometown,omitempty" scope:"user_hometown"`
	UserLikes                   *bool     `json:"user_likes,omitempty" scope:"user_likes"`
	UserLink                    *bool     `json:"user_link,omitempty" scope:"user_link"`
	UserLocation                *bool     `json:"user_location,omitempty" scope:"user_location"`
	UserPhotos                  *bool     `json:"user_photos,omitempty" scope:"user_photos"`
	UserPosts                   *bool     `json:"user_posts,omitempty" scope:"user_posts"`
	UserTaggedPlaces            *bool     `json:"user_tagged_places,omitempty" scope:"user_tagged_places"`
	UserVideos                  *bool     `json:"user_videos,omitempty" scope:"user_videos"`
	BusinessManagement          *bool     `json:"business_management,omitempty" scope:"business_management"`
	LeadsRetrieval              *bool     `json:"leads_retrieval,omitempty" scope:"leads_retrieval"`
	ManagePages                 *bool     `json:"manage_pages,omitempty" scope:"manage_pages"`
	PagesManageCTA              *bool     `json:"pages_manage_cta,omitempty" scope:"pages_manage_cta"`
	PagesManageInstantArticles  *bool     `json:"pages_manage_instant_articles,omitempty" scope:"pages_manage_instant_articles"`
	PagesShowList               *bool     `json:"pages_show_list,omitempty" scope:"pages_show_list"`
	PagesMessaging              *bool     `json:"pages_messaging,omitempty" scope:"pages_messaging"`
	PagesMessagingPhoneNumber   *bool     `json:"pages_messaging_phone_number,omitempty" scope:"pages_messaging_phone_number"`
	PagesMessagingSubscriptions *bool     `json:"pages_messaging_subscriptions,omitempty" scope:"pages_messaging_subscriptions"`
	PublishPages                *bool     `json:"publish_pages,omitempty" scope:"publish_pages"`
	PublishVideo                *bool     `json:"publish_video,omitempty" scope:"publish_video"`
	ReadPageMailboxes           *bool     `json:"read_page_mailboxes,omitempty" scope:"read_page_mailboxes"`
	ReadStream                  *bool     `json:"read_stream,omitempty" scope:"read_stream"`
	UserGroups                  *bool     `json:"user_groups,omitempty" scope:"user_groups"`
	UserManagedGroups           *bool     `json:"user_managed_groups,omitempty" scope:"user_managed_groups"`
	UserStatus                  *bool     `json:"user_status,omitempty" scope:"user_status"`
	SetUserAttributes           *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs          *[]string `json:"non_persistent_attrs,omitempty"`

	// Scope is a comma separated list of scopes.
	Scope *string `json:"scope,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsFacebook.
func (c *ConnectionOptionsFacebook) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsFacebook.
func (c *ConnectionOptionsFacebook) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsApple is used to configure an Apple Connection.
type ConnectionOptionsApple struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"app_secret,omitempty"`

	TeamID *string `json:"team_id,omitempty"`
	KeyID  *string `json:"kid,omitempty"`

	Name  *bool `json:"name,omitempty" scope:"name"`
	Email *bool `json:"email,omitempty" scope:"email"`

	Scope              *string   `json:"scope,omitempty"`
	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsApple.
func (c *ConnectionOptionsApple) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsApple.
func (c *ConnectionOptionsApple) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsLinkedin is used to configure a Linkedin Connection.
type ConnectionOptionsLinkedin struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	StrategyVersion *int `json:"strategy_version,omitempty"`

	Email        *bool `json:"email,omitempty" scope:"email"`
	Profile      *bool `json:"profile,omitempty" scope:"profile"`
	BasicProfile *bool `json:"basic_profile,omitempty" scope:"basic_profile"`

	Scope []interface{} `json:"scope,omitempty"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsLinkedin.
func (c *ConnectionOptionsLinkedin) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsLinkedin.
func (c *ConnectionOptionsLinkedin) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsGitHub is used to configure a GitHub Connection.
type ConnectionOptionsGitHub struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	Email          *bool `json:"email,omitempty" scope:"email"`
	ReadUser       *bool `json:"read_user,omitempty" scope:"read_user"`
	Follow         *bool `json:"follow,omitempty" scope:"follow"`
	PublicRepo     *bool `json:"public_repo,omitempty" scope:"public_repo"`
	Repo           *bool `json:"repo,omitempty" scope:"repo"`
	RepoDeployment *bool `json:"repo_deployment,omitempty" scope:"repo_deployment"`
	RepoStatus     *bool `json:"repo_status,omitempty" scope:"repo_status"`
	DeleteRepo     *bool `json:"delete_repo,omitempty" scope:"delete_repo"`
	Notifications  *bool `json:"notifications,omitempty" scope:"notifications"`
	Gist           *bool `json:"gist,omitempty" scope:"gist"`
	ReadRepoHook   *bool `json:"read_repo_hook,omitempty" scope:"read_repo_hook"`
	WriteRepoHook  *bool `json:"write_repo_hook,omitempty" scope:"write_repo_hook"`
	AdminRepoHook  *bool `json:"admin_repo_hook,omitempty" scope:"admin_repo_hook"`
	ReadOrg        *bool `json:"read_org,omitempty" scope:"read_org"`
	AdminOrg       *bool `json:"admin_org,omitempty" scope:"admin_org"`
	ReadPublicKey  *bool `json:"read_public_key,omitempty" scope:"read_public_key"`
	WritePublicKey *bool `json:"write_public_key,omitempty" scope:"write_public_key"`
	AdminPublicKey *bool `json:"admin_public_key,omitempty" scope:"admin_public_key"`
	WriteOrg       *bool `json:"write_org,omitempty" scope:"write_org"`
	Profile        *bool `json:"profile,omitempty" scope:"profile"`

	Scope []interface{} `json:"scope,omitempty"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsGitHub.
func (c *ConnectionOptionsGitHub) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsGitHub.
func (c *ConnectionOptionsGitHub) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsEmail is used to configure an Email Connection.
type ConnectionOptionsEmail struct {
	Name  *string                         `json:"name,omitempty"`
	Email *ConnectionOptionsEmailSettings `json:"email,omitempty"`

	OTP *ConnectionOptionsOTP `json:"totp,omitempty"`

	AuthParams interface{} `json:"authParams,omitempty"`

	DisableSignup        *bool     `json:"disable_signup,omitempty"`
	BruteForceProtection *bool     `json:"brute_force_protection,omitempty"`
	SetUserAttributes    *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs   *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// ConnectionOptionsEmailSettings is used to configure
// the email settings on a ConnectionOptionsEmail.
type ConnectionOptionsEmailSettings struct {
	Syntax  *string `json:"syntax,omitempty"`
	From    *string `json:"from,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Body    *string `json:"body,omitempty"`
}

// ConnectionOptionsOTP is used to configure the
// OTP settings on a ConnectionOptionsEmail.
type ConnectionOptionsOTP struct {
	TimeStep *int `json:"time_step,omitempty"`
	Length   *int `json:"length,omitempty"`
}

// ConnectionGatewayAuthentication is used to configure the
// GatewayAuthentication settings on a ConnectionOptionsSMS.
type ConnectionGatewayAuthentication struct {
	Method              *string `json:"method,omitempty"`
	Subject             *string `json:"subject,omitempty"`
	Audience            *string `json:"audience,omitempty"`
	Secret              *string `json:"secret,omitempty"`
	SecretBase64Encoded *bool   `json:"secret_base64_encoded,omitempty"`
}

// ConnectionOptionsSMS is used to configure an SMS Connection.
type ConnectionOptionsSMS struct {
	Name     *string `json:"name,omitempty"`
	From     *string `json:"from"`
	Syntax   *string `json:"syntax,omitempty"`
	Template *string `json:"template,omitempty"`

	OTP *ConnectionOptionsOTP `json:"totp,omitempty"`

	AuthParams interface{} `json:"authParams,omitempty"`

	TwilioSID           *string `json:"twilio_sid,omitempty"`
	TwilioToken         *string `json:"twilio_token,omitempty"`
	MessagingServiceSID *string `json:"messaging_service_sid"`

	Provider              *string                          `json:"provider,omitempty"`
	GatewayURL            *string                          `json:"gateway_url,omitempty"`
	GatewayAuthentication *ConnectionGatewayAuthentication `json:"gateway_authentication,omitempty"`
	ForwardRequestInfo    *bool                            `json:"forward_req_info,omitempty"`

	DisableSignup        *bool `json:"disable_signup,omitempty"`
	BruteForceProtection *bool `json:"brute_force_protection,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// ConnectionOptionsWindowsLive is used to configure a WindowsLive Connection.
type ConnectionOptionsWindowsLive struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	StrategyVersion *int `json:"strategy_version,omitempty"`

	OfflineAccess   *bool `json:"offline_access,omitempty" scope:"offline_access"`
	UserUpdate      *bool `json:"graph_user_update,omitempty" scope:"graph_user_update"`
	UserActivity    *bool `json:"graph_user_activity,omitempty" scope:"graph_user_activity"`
	Device          *bool `json:"graph_device,omitempty" scope:"graph_device"`
	Emails          *bool `json:"graph_emails,omitempty" scope:"graph_emails"`
	NotesUpdate     *bool `json:"graph_notes_update,omitempty" scope:"graph_notes_update"`
	User            *bool `json:"graph_user,omitempty" scope:"graph_user"`
	DeviceCommand   *bool `json:"graph_device_command,omitempty" scope:"graph_device_command"`
	EmailsUpdate    *bool `json:"graph_emails_update,omitempty" scope:"graph_emails_update"`
	Calendars       *bool `json:"graph_calendars,omitempty" scope:"graph_calendars"`
	CalendarsUpdate *bool `json:"graph_calendars_update,omitempty" scope:"graph_calendars_update"`
	Contacts        *bool `json:"graph_contacts,omitempty" scope:"graph_contacts"`
	ContactsUpdate  *bool `json:"graph_contacts_update,omitempty" scope:"graph_contacts_update"`
	Files           *bool `json:"graph_files,omitempty" scope:"graph_files"`
	FilesAll        *bool `json:"graph_files_all,omitempty" scope:"graph_files_all"`
	FilesUpdate     *bool `json:"graph_files_update,omitempty" scope:"graph_files_update"`
	FilesAllUpdate  *bool `json:"graph_files_all_update,omitempty" scope:"graph_files_all_update"`
	Notes           *bool `json:"graph_notes,omitempty" scope:"graph_notes"`
	NotesCreate     *bool `json:"graph_notes_create,omitempty" scope:"graph_notes_create"`
	Tasks           *bool `json:"graph_tasks,omitempty" scope:"graph_tasks"`
	TasksUpdate     *bool `json:"graph_tasks_update,omitempty" scope:"graph_tasks_update"`
	Signin          *bool `json:"signin,omitempty" scope:"signin"`

	Scope []interface{} `json:"scope,omitempty"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsWindowsLive.
func (c *ConnectionOptionsWindowsLive) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsWindowsLive.
func (c *ConnectionOptionsWindowsLive) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsSalesforce is used to configure a SalesForce Connection.
type ConnectionOptionsSalesforce struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	Profile *bool `json:"profile,omitempty" scope:"profile"`

	Scope []interface{} `json:"scope,omitempty"`

	CommunityBaseURL   *string   `json:"community_base_url,omitempty"`
	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsSalesforce.
func (c *ConnectionOptionsSalesforce) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsSalesforce.
func (c *ConnectionOptionsSalesforce) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsOIDC is used to configure an OIDC Connection.
type ConnectionOptionsOIDC struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	TenantDomain  *string   `json:"tenant_domain,omitempty"`
	DomainAliases *[]string `json:"domain_aliases,omitempty"`
	LogoURL       *string   `json:"icon_url,omitempty"`

	DiscoveryURL          *string `json:"discovery_url"`
	AuthorizationEndpoint *string `json:"authorization_endpoint"`
	Issuer                *string `json:"issuer"`
	JWKSURI               *string `json:"jwks_uri"`
	Type                  *string `json:"type"`
	UserInfoEndpoint      *string `json:"userinfo_endpoint"`
	TokenEndpoint         *string `json:"token_endpoint"`
	Scope                 *string `json:"scope,omitempty"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`

	ConnectionSettings *ConnectionOptionsOIDCConnectionSettings `json:"connection_settings,omitempty"`
	AttributeMap       *ConnectionOptionsOIDCAttributeMap       `json:"attribute_map,omitempty"`
}

// ConnectionOptionsOIDCConnectionSettings contains PKCE configuration for the connection.
//
// PKCE possible values:
//
//	auto     - Uses the strongest algorithm available.
//	S256     - Uses the SHA-256 algorithm. Auth0 does not currently support RS512 tokens.
//	plain    - Uses plaintext as described in the PKCE specification.
//	disabled - Disables support for PKCE.
//
// Setting the PKCE property to a value other than auto may prevent a connection from
// working properly if the selected value is not supported by the identity provider.
type ConnectionOptionsOIDCConnectionSettings struct {
	PKCE *string `json:"pkce,omitempty"`
}

// ConnectionOptionsOIDCAttributeMap contains the mapping of claims received from the identity provider (IdP).
type ConnectionOptionsOIDCAttributeMap struct {
	// Scopes to send to the IdP's Userinfo endpoint.
	UserInfoScope *string `json:"userinfo_scope,omitempty"`
	// Method used to map incoming claims.
	// Possible values: `use_map`, `bind_all`, `basic_profile`.
	MappingMode *string `json:"mapping_mode,omitempty"`
	// Object containing mapping details for incoming claims.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsOIDC.
func (c *ConnectionOptionsOIDC) Scopes() []string {
	return strings.Fields(c.GetScope())
}

// SetScopes sets the scopes for ConnectionOptionsOIDC.
func (c *ConnectionOptionsOIDC) SetScopes(enable bool, scopes ...string) {
	scopeMap := make(map[string]bool)
	for _, scope := range c.Scopes() {
		scopeMap[scope] = true
	}
	for _, scope := range scopes {
		scopeMap[scope] = enable
	}
	scopeSlice := make([]string, 0, len(scopeMap))
	for scope, enabled := range scopeMap {
		if enabled {
			scopeSlice = append(scopeSlice, scope)
		}
	}
	sort.Strings(scopeSlice)
	scope := strings.Join(scopeSlice, " ")
	c.Scope = &scope
}

// ConnectionOptionsOAuth2 is used to configure an OAuth2 Connection.
type ConnectionOptionsOAuth2 struct {
	ClientID           *string   `json:"client_id,omitempty"`
	ClientSecret       *string   `json:"client_secret,omitempty"`
	AuthorizationURL   *string   `json:"authorizationURL"`
	TokenURL           *string   `json:"tokenURL"`
	Scope              *string   `json:"scope,omitempty"`
	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`
	LogoURL            *string   `json:"icon_url,omitempty"`
	PKCEEnabled        *bool     `json:"pkce_enabled,omitempty"`
	// Scripts for the connection
	// Allowed keys are: "fetchUserProfile"
	Scripts *map[string]string `json:"scripts,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsOAuth2.
func (c *ConnectionOptionsOAuth2) Scopes() []string {
	return strings.Fields(c.GetScope())
}

// SetScopes sets the scopes for ConnectionOptionsOAuth2.
func (c *ConnectionOptionsOAuth2) SetScopes(enable bool, scopes ...string) {
	scopeMap := make(map[string]bool)
	for _, scope := range c.Scopes() {
		scopeMap[scope] = true
	}
	for _, scope := range scopes {
		scopeMap[scope] = enable
	}
	scopeSlice := make([]string, 0, len(scopeMap))
	for scope, enabled := range scopeMap {
		if enabled {
			scopeSlice = append(scopeSlice, scope)
		}
	}
	sort.Strings(scopeSlice)
	scope := strings.Join(scopeSlice, " ")
	c.Scope = &scope
}

// ConnectionOptionsAD is used to configure an AD Connection.
type ConnectionOptionsAD struct {
	TenantDomain         *string                `json:"tenant_domain,omitempty"`
	DomainAliases        *[]string              `json:"domain_aliases,omitempty"`
	LogoURL              *string                `json:"icon_url,omitempty"`
	IPs                  *[]string              `json:"ips,omitempty"`
	CertAuth             *bool                  `json:"certAuth,omitempty"`
	Kerberos             *bool                  `json:"kerberos,omitempty"`
	DisableCache         *bool                  `json:"disable_cache,omitempty"`
	BruteForceProtection *bool                  `json:"brute_force_protection,omitempty"`
	SetUserAttributes    *string                `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs   *[]string              `json:"non_persistent_attrs,omitempty"`
	UpstreamParams       map[string]interface{} `json:"upstream_params,omitempty"`
	Thumbprints          *[]string              `json:"thumbprints,omitempty"`
	Certs                *[]string              `json:"certs,omitempty"`
	AgentIP              *string                `json:"agentIP,omitempty"`
	AgentVersion         *string                `json:"agentVersion,omitempty"`
	AgentMode            *bool                  `json:"agentMode,omitempty"`

	// Set to true to stop the "Forgot Password" being displayed on login pages.
	DisableSelfServiceChangePassword *bool `json:"disable_self_service_change_password,omitempty"`
}

// ConnectionOptionsAzureAD is used to configure an AzureAD Connection.
type ConnectionOptionsAzureAD struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	AppID         *string   `json:"app_id,omitempty"`
	TenantDomain  *string   `json:"tenant_domain,omitempty"`
	Domain        *string   `json:"domain,omitempty"`
	DomainAliases *[]string `json:"domain_aliases,omitempty"`
	LogoURL       *string   `json:"icon_url,omitempty"`

	IdentityAPI *string `json:"identity_api,omitempty"`

	WAADProtocol       *string `json:"waad_protocol,omitempty"`
	WAADCommonEndpoint *bool   `json:"waad_common_endpoint,omitempty"`

	UseWSFederation     *bool   `json:"use_wsfed,omitempty"`
	UseCommonEndpoint   *bool   `json:"useCommonEndpoint,omitempty"`
	EnableUsersAPI      *bool   `json:"api_enable_users,omitempty"`
	MaxGroupsToRetrieve *string `json:"max_groups_to_retrieve,omitempty"`

	BasicProfile    *bool `json:"basic_profile,omitempty" scope:"basic_profile"`
	ExtendedProfile *bool `json:"ext_profile,omitempty" scope:"ext_profile"`
	Groups          *bool `json:"ext_groups,omitempty" scope:"ext_groups"`
	NestedGroups    *bool `json:"ext_nested_groups,omitempty" scope:"ext_nested_groups"`
	Admin           *bool `json:"ext_admin,omitempty" scope:"ext_admin"`
	IsSuspended     *bool `json:"ext_is_suspended,omitempty" scope:"ext_is_suspended"`
	AgreedTerms     *bool `json:"ext_agreed_terms,omitempty" scope:"ext_agreed_terms"`
	AssignedPlans   *bool `json:"ext_assigned_plans,omitempty" scope:"ext_assigned_plans"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	TrustEmailVerified *string   `json:"should_trust_email_verified_connection,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`

	AppDomain                *string   `json:"app_domain,omitempty"`
	Thumbprints              *[]string `json:"thumbprints,omitempty"`
	CertRolloverNotification *string   `json:"cert_rollover_notification,omitempty"`
	Granted                  *bool     `json:"granted,omitempty"`
	TenantID                 *string   `json:"tenantId,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsAzureAD.
func (c *ConnectionOptionsAzureAD) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsAzureAD.
func (c *ConnectionOptionsAzureAD) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsADFS is used to configure an ADFS Connection.
type ConnectionOptionsADFS struct {
	TenantDomain       *string                `json:"tenant_domain,omitempty"`
	DomainAliases      *[]string              `json:"domain_aliases,omitempty"`
	LogoURL            *string                `json:"icon_url,omitempty"`
	ADFSServer         *string                `json:"adfs_server,omitempty"`
	FedMetadataXML     *string                `json:"fedMetadataXml,omitempty"`
	EnableUsersAPI     *bool                  `json:"api_enable_users,omitempty"`
	NonPersistentAttrs *[]string              `json:"non_persistent_attrs,omitempty"`
	UpstreamParams     map[string]interface{} `json:"upstream_params,omitempty"`
	Thumbprints        *[]string              `json:"thumbprints,omitempty"`
	SignInEndpoint     *string                `json:"signInEndpoint,omitempty"`
	TrustEmailVerified *string                `json:"should_trust_email_verified_connection,omitempty"`

	// Set to on_first_login to avoid setting user attributes at each login.
	SetUserAttributes *string `json:"set_user_root_attributes,omitempty"`

	EntityID                 *string   `json:"entityID,omitempty"`
	CertRolloverNotification *string   `json:"cert_rollover_notification,omitempty"`
	PreviousThumbprints      *[]string `json:"prev_thumbprints,omitempty"`
}

// ConnectionOptionsPingFederate is used to configure a Ping Federate Connection.
type ConnectionOptionsPingFederate struct {
	// SigningCert should be used when creating or updating the public key for the Ping Federate server, it will not be
	// present when reading a connection, and instead you should use the Cert field to check the value.
	SigningCert *string `json:"signingCert,omitempty"`

	// Cert should only be used when reading the connection. It should not be set on creation or update of a connection,
	// instead SigningCert should be used to update the public key for the Ping Federate server.
	Cert *string `json:"cert,omitempty"`

	LogoURL             *string                            `json:"icon_url,omitempty"`
	IdpInitiated        *ConnectionOptionsSAMLIdpInitiated `json:"idpinitiated,omitempty"`
	TenantDomain        *string                            `json:"tenant_domain,omitempty"`
	DomainAliases       *[]string                          `json:"domain_aliases,omitempty"`
	SignInEndpoint      *string                            `json:"signInEndpoint,omitempty"`
	DigestAlgorithm     *string                            `json:"digestAlgorithm,omitempty"`
	SignSAMLRequest     *bool                              `json:"signSAMLRequest,omitempty"`
	SignatureAlgorithm  *string                            `json:"signatureAlgorithm,omitempty"`
	PingFederateBaseURL *string                            `json:"pingFederateBaseUrl,omitempty"`
	NonPersistentAttrs  *[]string                          `json:"non_persistent_attrs,omitempty"`
	UpstreamParams      map[string]interface{}             `json:"upstream_params,omitempty"`
	SetUserAttributes   *string                            `json:"set_user_root_attributes,omitempty"`

	APIEnableUsers           *bool                               `json:"api_enable_users,omitempty"`
	SignOutEndpoint          *string                             `json:"signOuEndpoint,omitempty"`
	Subject                  map[string]interface{}              `json:"subject,omitempty"`
	DisableSignout           *bool                               `json:"disableSignout,omitempty"`
	UserIDAttribute          *string                             `json:"user_id_attribute,omitempty"`
	Debug                    *bool                               `json:"debug,omitempty"`
	ProtocolBinding          *string                             `json:"protocolBinding,omitempty"`
	RequestTemplate          *string                             `json:"requestTemplate,omitempty"`
	BindingMethod            *string                             `json:"bindingMethod,omitempty"`
	Thumbprints              *[]string                           `json:"thumbprints,omitempty"`
	Expires                  *string                             `json:"expires,omitempty"`
	MetadataURL              *string                             `json:"metadataUrl,omitempty"`
	FieldsMap                map[string]interface{}              `json:"fields_map,omitempty"`
	MetadataXML              *string                             `json:"metadataXml,omitempty"`
	EntityID                 *string                             `json:"entityId,omitempty"`
	CertRolloverNotification *string                             `json:"cert_rollover_notification,omitempty"`
	SigningKey               *ConnectionOptionsSAMLSigningKey    `json:"signing_key,omitempty"`
	DecryptionKey            *ConnectionOptionsSAMLDecryptionKey `json:"decryption_key,omitempty"`
	AgentIP                  *string                             `json:"agentIP,omitempty"`
	AgentVersion             *string                             `json:"agentVersion,omitempty"`
	AgentMode                *bool                               `json:"agentMode,omitempty"`
	ExtGroups                *bool                               `json:"ext_groups,omitempty" scope:"ext_groups"`
	ExtProfile               *bool                               `json:"ext_profile,omitempty" scope:"ext_profile"`
}

// Scopes returns the scopes for ConnectionOptionsPingFederate.
func (c *ConnectionOptionsPingFederate) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsPingFederate.
func (c *ConnectionOptionsPingFederate) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsSAML is used to configure a SAML Connection.
type ConnectionOptionsSAML struct {
	Cert               *string                             `json:"cert,omitempty"`
	Debug              *bool                               `json:"debug,omitempty"`
	Expires            *string                             `json:"expires,omitempty"`
	IdpInitiated       *ConnectionOptionsSAMLIdpInitiated  `json:"idpinitiated,omitempty"`
	SigningKey         *ConnectionOptionsSAMLSigningKey    `json:"signing_key,omitempty"`
	DecryptionKey      *ConnectionOptionsSAMLDecryptionKey `json:"decryptionKey,omitempty"`
	SigningCert        *string                             `json:"signingCert,omitempty"`
	Thumbprints        []interface{}                       `json:"thumbprints,omitempty"`
	ProtocolBinding    *string                             `json:"protocolBinding,omitempty"`
	TenantDomain       *string                             `json:"tenant_domain,omitempty"`
	DomainAliases      *[]string                           `json:"domain_aliases,omitempty"`
	SignInEndpoint     *string                             `json:"signInEndpoint,omitempty"`
	SignOutEndpoint    *string                             `json:"signOutEndpoint,omitempty"`
	DisableSignOut     *bool                               `json:"disableSignout,omitempty"`
	SignatureAlgorithm *string                             `json:"signatureAlgorithm,omitempty"`
	DigestAglorithm    *string                             `json:"digestAlgorithm,omitempty"`
	MetadataXML        *string                             `json:"metadataXml,omitempty"`
	MetadataURL        *string                             `json:"metadataUrl,omitempty"`
	FieldsMap          map[string]interface{}              `json:"fieldsMap,omitempty"`
	Subject            map[string]interface{}              `json:"subject,omitempty"`
	SignSAMLRequest    *bool                               `json:"signSAMLRequest,omitempty"`
	RequestTemplate    *string                             `json:"requestTemplate,omitempty"`
	UserIDAttribute    *string                             `json:"user_id_attribute,omitempty"`
	LogoURL            *string                             `json:"icon_url,omitempty"`
	EntityID           *string                             `json:"entityId,omitempty"`

	BindingMethod            *string `json:"binding_method,omitempty"`
	CertRolloverNotification *string `json:"cert_rollover_notification,omitempty"`
	AgentIP                  *string `json:"agentIP,omitempty"`
	AgentVersion             *string `json:"agentVersion,omitempty"`
	AgentMode                *bool   `json:"agentMode,omitempty"`
	ExtGroups                *bool   `json:"ext_groups,omitempty" scope:"ext_groups"`
	ExtProfile               *bool   `json:"ext_profile,omitempty" scope:"ext_profile"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`

	UpstreamParams map[string]interface{} `json:"upstream_params,omitempty"`
}

// ConnectionOptionsSAMLIdpInitiated is used to configure the
// IdpInitiated settings on a ConnectionOptionsSAML.
type ConnectionOptionsSAMLIdpInitiated struct {
	Enabled              *bool   `json:"enabled,omitempty"`
	ClientID             *string `json:"client_id,omitempty"`
	ClientProtocol       *string `json:"client_protocol,omitempty"`
	ClientAuthorizeQuery *string `json:"client_authorizequery,omitempty"`

	SetUserAttributes  *string   `json:"set_user_root_attributes,omitempty"`
	NonPersistentAttrs *[]string `json:"non_persistent_attrs,omitempty"`
}

// ConnectionOptionsSAMLSigningKey is used to configure the
// SigningKey settings on a ConnectionOptionsSAML.
type ConnectionOptionsSAMLSigningKey struct {
	Key  *string `json:"key,omitempty"`
	Cert *string `json:"cert,omitempty"`
}

// ConnectionOptionsSAMLDecryptionKey is used to configure the
// DecryptionKey settings on a ConnectionOptionsSAML.
type ConnectionOptionsSAMLDecryptionKey struct {
	Key  *string `json:"key,omitempty"`
	Cert *string `json:"cert,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsSAML.
func (c *ConnectionOptionsSAML) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsSAML.
func (c *ConnectionOptionsSAML) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionOptionsGoogleApps is used to configure a GoogleApps Connection.
type ConnectionOptionsGoogleApps struct {
	ClientID                  *string                `json:"client_id,omitempty"`
	ClientSecret              *string                `json:"client_secret,omitempty"`
	Domain                    *string                `json:"domain,omitempty"`
	TenantDomain              *string                `json:"tenant_domain,omitempty"`
	BasicProfile              *bool                  `json:"basic_profile,omitempty" scope:"basic_profile"`
	ExtendedProfile           *bool                  `json:"ext_profile,omitempty" scope:"ext_profile"`
	Groups                    *bool                  `json:"ext_groups,omitempty" scope:"ext_groups"`
	Admin                     *bool                  `json:"ext_is_admin,omitempty" scope:"ext_is_admin"`
	IsSuspended               *bool                  `json:"ext_is_suspended,omitempty" scope:"ext_is_suspended"`
	AgreedTerms               *bool                  `json:"ext_agreed_terms,omitempty" scope:"ext_agreed_terms"`
	EnableUsersAPI            *bool                  `json:"api_enable_users,omitempty"`
	SetUserAttributes         *string                `json:"set_user_root_attributes,omitempty"`
	MapUserIDtoID             *bool                  `json:"map_user_id_to_id,omitempty"`
	HandleLoginFromSocial     *bool                  `json:"handle_login_from_social,omitempty"`
	DomainAliases             *[]string              `json:"domain_aliases,omitempty"`
	LogoURL                   *string                `json:"icon_url,omitempty"`
	AdminAccessTokenExpiresIn *string                `json:"admin_access_token_expiresin,omitempty"`
	AdminAccessToken          *string                `json:"admin_access_token,omitempty"`
	AdminRefreshToken         *string                `json:"admin_refresh_token,omitempty"`
	UpstreamParams            map[string]interface{} `json:"upstream_params,omitempty"`
	NonPersistentAttrs        *[]string              `json:"non_persistent_attrs,omitempty"`
}

// Scopes returns the scopes for ConnectionOptionsGoogleApps.
func (c *ConnectionOptionsGoogleApps) Scopes() []string {
	return tag.Scopes(c)
}

// SetScopes sets the scopes for ConnectionOptionsGoogleApps.
func (c *ConnectionOptionsGoogleApps) SetScopes(enable bool, scopes ...string) {
	tag.SetScopes(c, enable, scopes...)
}

// ConnectionManager manages Auth0 Connection resources.
type ConnectionManager manager

// ConnectionList is a list of Connections.
type ConnectionList struct {
	List
	Connections []*Connection `json:"connections"`
}

// Create a new connection.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/post_connections
func (m *ConnectionManager) Create(ctx context.Context, c *Connection, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("connections"), c, opts...)
}

// Read retrieves a connection by its id.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/get_connections_by_id
func (m *ConnectionManager) Read(ctx context.Context, id string, opts ...RequestOption) (c *Connection, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("connections", id), &c, opts...)
	return
}

// List connections.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/get_connections
func (m *ConnectionManager) List(ctx context.Context, opts ...RequestOption) (c *ConnectionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("connections"), &c, applyListDefaults(opts))
	return
}

// Update a connection.
//
// Note: if you use the options' parameter, the whole options object will be
// overridden, so ensure that all parameters are present.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/patch_connections_by_id
func (m *ConnectionManager) Update(ctx context.Context, id string, c *Connection, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("connections", id), c, opts...)
}

// Delete a connection and all its users.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/delete_connections_by_id
func (m *ConnectionManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("connections", id), nil, opts...)
}

// ReadByName retrieves a connection by its name. This is a helper method when a
// connection id is not readily available.
func (m *ConnectionManager) ReadByName(ctx context.Context, name string, opts ...RequestOption) (*Connection, error) {
	if name == "" {
		return nil, &managementError{400, "Bad Request", "Name cannot be empty"}
	}
	c, err := m.List(ctx, append(opts, Parameter("name", name))...)
	if err != nil {
		return nil, err
	}
	if len(c.Connections) > 0 {
		return c.Connections[0], nil
	}
	return nil, &managementError{404, "Not Found", "Connection not found"}
}
