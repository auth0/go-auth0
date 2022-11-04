package management

import (
	"encoding/json"
	"net/http"
)

const (
	// EmailProviderMandrill constant.
	EmailProviderMandrill = "mandrill"

	// EmailProviderSES constant.
	EmailProviderSES = "ses"

	// EmailProviderSendGrid constant.
	EmailProviderSendGrid = "sendgrid"

	// EmailProviderSparkPost constant.
	EmailProviderSparkPost = "sparkpost"

	// EmailProviderMailgun constant.
	EmailProviderMailgun = "mailgun"

	// EmailProviderSMTP constant.
	EmailProviderSMTP = "smtp"
)

// EmailProvider is used to configure Email Providers.
//
// See: https://auth0.com/docs/customize/email
type EmailProvider struct {
	// The name of the email provider.
	// Can be one of "mandrill", "ses", "sendgrid", "sparkpost", "mailgun"  or "smtp".
	Name *string `json:"name,omitempty"`

	// Indicates whether the email provider is enabled or not.
	// Defaults to true.
	Enabled *bool `json:"enabled,omitempty"`

	// The default FROM address.
	DefaultFromAddress *string `json:"default_from_address,omitempty"`

	// Credentials required to use the provider.
	Credentials interface{} `json:"-"`

	// Specific provider settings.
	Settings interface{} `json:"-"`
}

// EmailProviderCredentialsMandrill represent the
// credentials required to use the mandrill provider.
type EmailProviderCredentialsMandrill struct {
	APIKey *string `json:"api_key,omitempty"`
}

// EmailProviderCredentialsSES represent the
// credentials required to use the ses provider.
type EmailProviderCredentialsSES struct {
	AccessKeyID     *string `json:"accessKeyId,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	Region          *string `json:"region,omitempty"`
}

// EmailProviderCredentialsSendGrid represent the
// credentials required to use the sendgrid provider.
type EmailProviderCredentialsSendGrid struct {
	APIKey *string `json:"api_key,omitempty"`
}

// EmailProviderCredentialsSparkPost represent the
// credentials required to use the sparkpost provider.
type EmailProviderCredentialsSparkPost struct {
	APIKey *string `json:"api_key,omitempty"`
	Region *string `json:"region,omitempty"`
}

// EmailProviderCredentialsMailgun represent the
// credentials required to use the mailgun provider.
type EmailProviderCredentialsMailgun struct {
	APIKey *string `json:"api_key,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Region *string `json:"region,omitempty"`
}

// EmailProviderCredentialsSMTP represent the
// credentials required to use the smtp provider.
type EmailProviderCredentialsSMTP struct {
	SMTPHost *string `json:"smtp_host,omitempty"`
	SMTPPort *int    `json:"smtp_port,omitempty"`
	SMTPUser *string `json:"smtp_user,omitempty"`
	SMTPPass *string `json:"smtp_pass,omitempty"`
}

// EmailProviderSettingsMandrill are the provider
// specific settings used by the mandrill provider.
type EmailProviderSettingsMandrill struct {
	Message *EmailProviderSettingsMandrillMessage `json:"message,omitempty"`
}

// EmailProviderSettingsMandrillMessage contains the
// message setting content for the mandrill provider.
type EmailProviderSettingsMandrillMessage struct {
	ViewContentLink *bool `json:"view_content_link,omitempty"`
}

// EmailProviderSettingsSES are the provider
// specific settings used by the ses provider.
type EmailProviderSettingsSES struct {
	Message *EmailProviderSettingsSESMessage `json:"message,omitempty"`
}

// EmailProviderSettingsSESMessage contains the
// message setting content for the ses provider.
type EmailProviderSettingsSESMessage struct {
	ConfigurationSetName *string `json:"configuration_set_name,omitempty"`
}

// EmailProviderSettingsSMTP are the provider
// specific settings used by the smtp provider.
type EmailProviderSettingsSMTP struct {
	Headers *EmailProviderSettingsSMTPHeaders `json:"headers,omitempty"`
}

// EmailProviderSettingsSMTPHeaders contains the
// headers settings content for the smtp provider.
type EmailProviderSettingsSMTPHeaders struct {
	XMCViewContentLink   *string `json:"X-MC-ViewContentLink,omitempty"`
	XSESConfigurationSet *string `json:"X-SES-Configuration-Set,omitempty"`
}

// MarshalJSON is a custom serializer for the EmailProvider type.
func (ep *EmailProvider) MarshalJSON() ([]byte, error) {
	type emailProvider EmailProvider
	type emailProviderWrapper struct {
		*emailProvider
		RawCredentials json.RawMessage `json:"credentials,omitempty"`
		RawSettings    json.RawMessage `json:"settings,omitempty"`
	}

	wrapper := &emailProviderWrapper{(*emailProvider)(ep), nil, nil}

	if ep.Credentials != nil {
		credentialsJSON, err := json.Marshal(ep.Credentials)
		if err != nil {
			return nil, err
		}
		wrapper.RawCredentials = credentialsJSON
	}

	if ep.Settings != nil {
		settingsJSON, err := json.Marshal(ep.Settings)
		if err != nil {
			return nil, err
		}
		wrapper.RawSettings = settingsJSON
	}

	return json.Marshal(wrapper)
}

// UnmarshalJSON is a custom deserializer for the EmailProvider type.
func (ep *EmailProvider) UnmarshalJSON(b []byte) error {
	type emailProvider EmailProvider
	type emailProviderWrapper struct {
		*emailProvider
		RawCredentials json.RawMessage `json:"credentials,omitempty"`
		RawSettings    json.RawMessage `json:"settings,omitempty"`
	}

	wrapper := &emailProviderWrapper{(*emailProvider)(ep), nil, nil}

	if err := json.Unmarshal(b, wrapper); err != nil {
		return err
	}

	var credentials, settings interface{}

	switch ep.GetName() {
	case EmailProviderMandrill:
		credentials = &EmailProviderCredentialsMandrill{}
		settings = &EmailProviderSettingsMandrill{}
	case EmailProviderSES:
		credentials = &EmailProviderCredentialsSES{}
		settings = &EmailProviderSettingsSES{}
	case EmailProviderSendGrid:
		credentials = &EmailProviderCredentialsSendGrid{}
		// No settings for sendgrid.
		settings = nil
	case EmailProviderSparkPost:
		credentials = &EmailProviderCredentialsSparkPost{}
		// No settings for sparkpost.
		settings = nil
	case EmailProviderMailgun:
		credentials = &EmailProviderCredentialsMailgun{}
		// No settings for mailgun.
		settings = nil
	case EmailProviderSMTP:
		credentials = &EmailProviderCredentialsSMTP{}
		settings = &EmailProviderSettingsSMTP{}
	case "":
		credentials = nil
		settings = nil
	default:
		// Just making sure we're covered if
		// new email providers are introduced.
		credentials = make(map[string]interface{})
		settings = make(map[string]interface{})
	}

	if wrapper.RawCredentials != nil {
		if err := json.Unmarshal(wrapper.RawCredentials, &credentials); err != nil {
			return err
		}
	}

	if wrapper.RawSettings != nil {
		if err := json.Unmarshal(wrapper.RawSettings, &settings); err != nil {
			return err
		}
	}

	ep.Credentials = credentials
	ep.Settings = settings

	return nil
}

// EmailProviderManager manages the Auth0 EmailProvider.
type EmailProviderManager struct {
	*Management
}

func newEmailProviderManager(m *Management) *EmailProviderManager {
	return &EmailProviderManager{m}
}

// Create an email provider.
//
// The credentials object requires different properties depending on the email
// provider (which is specified using the name property):
//
// - `mandrill` requires `api_key`
// - `sendgrid` requires `api_key`
// - `sparkpost` requires `api_key`. Optionally, set `region` to `eu` to use the
// SparkPost service hosted in Western Europe; set to `null` to use the
// SparkPost service hosted in North America. `eu` or `null` are the only valid
// values for `region`.
// - ses requires accessKeyId, secretAccessKey, and region
// - smtp requires smtp_host, smtp_port, smtp_user, and smtp_pass
// - `mailgun` requires `api_key` and `domain`. Optionally, set region to eu to
// use the Mailgun service hosted in Europe; set to null otherwise. "eu" or null
// are the only valid values for region.
//
// Depending on the type of provider it is possible to specify settings object
// with different configuration options, which will be used when sending an
// email:
//
// - `smtp` provider, `settings` may contain `headers` object. When using AWS
// SES SMTP host, you may provide a name of configuration set in an
// `X-SES-Configuration-Set` header. The value must be a string.
//
// See: https://auth0.com/docs/api/management/v2#!/Emails/post_provider
func (m *EmailProviderManager) Create(ep *EmailProvider, opts ...RequestOption) error {
	return m.Request(http.MethodPost, m.URI("emails", "provider"), ep, opts...)
}

// Read email provider details.
//
// See: https://auth0.com/docs/api/management/v2#!/Emails/get_provider
func (m *EmailProviderManager) Read(opts ...RequestOption) (ep *EmailProvider, err error) {
	opts = append(opts, IncludeFields("name", "enabled", "default_from_address", "credentials", "settings"))
	err = m.Request(http.MethodGet, m.URI("emails", "provider"), &ep, opts...)
	return
}

// Update an email provider.
//
// See: https://auth0.com/docs/api/management/v2#!/Emails/patch_provider
func (m *EmailProviderManager) Update(ep *EmailProvider, opts ...RequestOption) (err error) {
	return m.Request(http.MethodPatch, m.URI("emails", "provider"), ep, opts...)
}

// Delete the email provider.
//
// See: https://auth0.com/docs/api/management/v2#!/Emails/delete_provider
func (m *EmailProviderManager) Delete(opts ...RequestOption) (err error) {
	return m.Request(http.MethodDelete, m.URI("emails", "provider"), nil, opts...)
}
