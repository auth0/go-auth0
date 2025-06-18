package management

import (
	"context"
	"encoding/json"
	"fmt"
)

// Branding is used to customize the look and feel of Auth0 to align
// with an organization's brand requirements and user expectations.
//
// See: https://auth0.com/docs/customize
type Branding struct {
	// Change login page colors.
	Colors *BrandingColors `json:"colors,omitempty"`

	// URL for the favicon. Must use HTTPS.
	FaviconURL *string `json:"favicon_url,omitempty"`

	// URL for the logo. Must use HTTPS.
	LogoURL *string `json:"logo_url,omitempty"`

	Font *BrandingFont `json:"font,omitempty"`
}

// BrandingColors are used to customize the Universal Login Page.
type BrandingColors struct {
	// Accent color.
	Primary *string `json:"primary,omitempty"`

	// Page background color.
	//
	// Only one of PageBackground and PageBackgroundGradient should be set. If
	// both fields are set, PageBackground takes priority.
	PageBackground *string `json:"-"`

	// Page background gradient.
	//
	// Only one of PageBackground and PageBackgroundGradient should be set. If
	// both fields are set, PageBackground takes priority.
	PageBackgroundGradient *BrandingPageBackgroundGradient `json:"-"`
}

// BrandingPageBackgroundGradient is used to customize
// the background color of the Universal Login Page.
type BrandingPageBackgroundGradient struct {
	Type        *string `json:"type,omitempty"`
	Start       *string `json:"start,omitempty"`
	End         *string `json:"end,omitempty"`
	AngleDegree *int    `json:"angle_deg,omitempty"`
}

// BrandingPhoneProviderList is list of BrandingPhoneProvider.
type BrandingPhoneProviderList struct {
	Providers []*BrandingPhoneProvider `json:"providers,omitempty"`
}

// BrandingPhoneProvider is used to configure phone providers.
type BrandingPhoneProvider struct {
	ID            *string                             `json:"id,omitempty"`
	Name          *string                             `json:"name,omitempty"`
	Disabled      *bool                               `json:"disabled,omitempty"`
	Channel       *string                             `json:"channel,omitempty"`
	Tenant        *string                             `json:"tenant,omitempty"`
	Configuration *BrandingPhoneProviderConfiguration `json:"configuration,omitempty"`
	Credentials   *BrandingPhoneProviderCredential    `json:"credentials,omitempty"`
}

// BrandingTryPhoneProvider is used to test a phone provider.
type BrandingTryPhoneProvider struct {
	To             *string `json:"to,omitempty"`
	DeliveryMethod *string `json:"delivery_method,omitempty"`
	Code           *int    `json:"code,omitempty"`
	Message        *string `json:"message,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
//
// It is required to handle the json field credentials, which can either
// be a JSON object, or null.
func (b *BrandingTryPhoneProvider) MarshalJSON() ([]byte, error) {
	type BrandingTryPhoneProviderSubset struct {
		To             *string `json:"to,omitempty"`
		DeliveryMethod *string `json:"delivery_method,omitempty"`
	}

	return json.Marshal(&BrandingTryPhoneProviderSubset{
		To:             b.To,
		DeliveryMethod: b.DeliveryMethod,
	})
}

// BrandingTryPhoneNotificationTemplate is used to test a phone notification template.
type BrandingTryPhoneNotificationTemplate struct {
	To             *string `json:"to,omitempty"`
	DeliveryMethod *string `json:"delivery_method,omitempty"`
	Message        *string `json:"message,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
//
// It is required to handle the json field credentials, which can either
// be a JSON object, or null.
func (b *BrandingTryPhoneNotificationTemplate) MarshalJSON() ([]byte, error) {
	type BrandingTryPhoneNotificationTemplateSubset struct {
		To             *string `json:"to,omitempty"`
		DeliveryMethod *string `json:"delivery_method,omitempty"`
	}

	return json.Marshal(&BrandingTryPhoneNotificationTemplateSubset{
		To:             b.To,
		DeliveryMethod: b.DeliveryMethod,
	})
}

// BrandingPhoneProviderCredential represents the credentials for a phone provider.
type BrandingPhoneProviderCredential struct {
	AuthToken *string `json:"auth_token,omitempty"`
}

// BrandingPhoneProviderConfiguration is used to configure phone providers.
type BrandingPhoneProviderConfiguration struct {
	DefaultFrom     *string   `json:"default_from,omitempty"`
	MSSID           *string   `json:"mssid,omitempty"`
	SID             *string   `json:"sid,omitempty"`
	DeliveryMethods *[]string `json:"delivery_methods,omitempty"`
}

// BrandingPhoneNotificationTemplateList is list of BrandingPhoneNotificationTemplate.
type BrandingPhoneNotificationTemplateList struct {
	Templates []*BrandingPhoneNotificationTemplate `json:"templates,omitempty"`
}

// BrandingPhoneNotificationTemplate is used to customize the phone notification template.
type BrandingPhoneNotificationTemplate struct {
	ID           *string                                   `json:"id,omitempty"`
	Channel      *string                                   `json:"channel,omitempty"`
	Tenant       *string                                   `json:"tenant,omitempty"`
	Customizable *bool                                     `json:"customizable,omitempty"`
	Content      *BrandingPhoneNotificationTemplateContent `json:"content,omitempty"`
	Type         *string                                   `json:"type,omitempty"`
	Disabled     *bool                                     `json:"disabled,omitempty"`
}

// BrandingPhoneNotificationTemplateContent is used to customize the phone notification template content.
type BrandingPhoneNotificationTemplateContent struct {
	Syntax *string                                       `json:"syntax,omitempty"`
	From   *string                                       `json:"from,omitempty"`
	Body   *BrandingPhoneNotificationTemplateContentBody `json:"body,omitempty"`
}

// BrandingPhoneNotificationTemplateContentBody is used to customize the phone notification template content body.
type BrandingPhoneNotificationTemplateContentBody struct {
	Text  *string `json:"text,omitempty"`
	Voice *string `json:"voice,omitempty"`
}

func (bpnt *BrandingPhoneNotificationTemplate) reset(method string) {
	bpnt.ID = nil
	bpnt.Channel = nil
	bpnt.Tenant = nil
	bpnt.Customizable = nil

	switch method {
	case "update":
		bpnt.Type = nil
		bpnt.Content.Syntax = nil
	default:
	}
}

// reset resets the BrandingPhoneProvider fields.
func (b *BrandingPhoneProvider) reset() {
	b.ID = nil
	b.Channel = nil
	b.Tenant = nil
}

// MarshalJSON implements the json.Marshaler interface.
//
// It is required to handle the json field page_background, which can either
// be a hex color string, or an object describing a gradient.
func (bc *BrandingColors) MarshalJSON() ([]byte, error) {
	type brandingColors BrandingColors

	type brandingColorsWrapper struct {
		*brandingColors
		RawPageBackground interface{} `json:"page_background,omitempty"`
	}

	alias := &brandingColorsWrapper{(*brandingColors)(bc), nil}

	switch {
	case bc.PageBackground != nil && bc.PageBackgroundGradient != nil:
		return nil, fmt.Errorf("only one of PageBackground and PageBackgroundGradient is allowed")
	case bc.PageBackground != nil:
		alias.RawPageBackground = bc.PageBackground
	case bc.PageBackgroundGradient != nil:
		alias.RawPageBackground = bc.PageBackgroundGradient
	}

	return json.Marshal(alias)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the json field page_background, which can either
// be a hex color string, or an object describing a gradient.
func (bc *BrandingColors) UnmarshalJSON(data []byte) error {
	type brandingColors BrandingColors

	type brandingColorsWrapper struct {
		*brandingColors
		RawPageBackground json.RawMessage `json:"page_background,omitempty"`
	}

	alias := &brandingColorsWrapper{(*brandingColors)(bc), nil}

	err := json.Unmarshal(data, alias)
	if err != nil {
		return err
	}

	if alias.RawPageBackground != nil {
		var v interface{}

		err = json.Unmarshal(alias.RawPageBackground, &v)
		if err != nil {
			return err
		}

		switch rawPageBackground := v.(type) {
		case string:
			bc.PageBackground = &rawPageBackground

		case map[string]interface{}:
			var gradient BrandingPageBackgroundGradient

			err = json.Unmarshal(alias.RawPageBackground, &gradient)
			if err != nil {
				return err
			}

			bc.PageBackgroundGradient = &gradient

		default:
			return fmt.Errorf("unexpected type for field page_background")
		}
	}

	return nil
}

// BrandingFont is used to customize the
// font on the Universal Login Page.
type BrandingFont struct {
	// URL for the custom font. Must use HTTPS.
	URL *string `json:"url,omitempty"`
}

// BrandingUniversalLogin is used to customize
// the body of the Universal Login Page.
type BrandingUniversalLogin struct {
	Body *string `json:"body,omitempty"`
}

// BrandingManager manages Auth0 Branding resources.
type BrandingManager manager

// Read retrieves various settings related to branding.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get_branding
func (m *BrandingManager) Read(ctx context.Context, opts ...RequestOption) (b *Branding, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("branding"), &b, opts...)
	return
}

// Update various fields related to branding.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/patch_branding
func (m *BrandingManager) Update(ctx context.Context, t *Branding, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("branding"), t, opts...)
}

// UniversalLogin retrieves the template for the New Universal Login Experience.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get_universal_login
func (m *BrandingManager) UniversalLogin(ctx context.Context, opts ...RequestOption) (ul *BrandingUniversalLogin, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("branding", "templates", "universal-login"), &ul, opts...)
	return
}

// SetUniversalLogin sets the template for the New Universal Login Experience.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/put_universal_login
func (m *BrandingManager) SetUniversalLogin(ctx context.Context, ul *BrandingUniversalLogin, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PUT", m.management.URI("branding", "templates", "universal-login"), ul.Body, opts...)
}

// DeleteUniversalLogin deletes the template for the New Universal Login Experience.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/delete_universal_login
func (m *BrandingManager) DeleteUniversalLogin(ctx context.Context, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("branding", "templates", "universal-login"), nil, opts...)
}

// ListPhoneProviders retrieves the list of phone providers for a Tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get-branding-phone-providers
func (m *BrandingManager) ListPhoneProviders(ctx context.Context, opts ...RequestOption) (pps *BrandingPhoneProviderList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("branding", "phone", "providers"), &pps, opts...)
	return
}

// ReadPhoneProvider retrieves a phone provider for a Tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get-phone-provider
func (m *BrandingManager) ReadPhoneProvider(ctx context.Context, id string, opts ...RequestOption) (pp *BrandingPhoneProvider, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("branding", "phone", "providers", id), &pp, opts...)
	return
}

// CreatePhoneProvider creates a phone provider for a Tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/create-phone-provider
func (m *BrandingManager) CreatePhoneProvider(ctx context.Context, pp *BrandingPhoneProvider, opts ...RequestOption) (err error) {
	pp.reset()
	return m.management.Request(ctx, "POST", m.management.URI("branding", "phone", "providers"), pp, opts...)
}

// DeletePhoneProvider deletes a phone provider for a Tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/delete-phone-provider
func (m *BrandingManager) DeletePhoneProvider(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("branding", "phone", "providers", id), nil, opts...)
}

// UpdatePhoneProvider updates a phone provider for a Tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/update-phone-provider
func (m *BrandingManager) UpdatePhoneProvider(ctx context.Context, id string, pp *BrandingPhoneProvider, opts ...RequestOption) (err error) {
	pp.reset()
	return m.management.Request(ctx, "PATCH", m.management.URI("branding", "phone", "providers", id), pp, opts...)
}

// TryPhoneProvider sends a test message to a phone provider for a Tenant.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/try-phone-providerto
func (m *BrandingManager) TryPhoneProvider(ctx context.Context, id string, provider *BrandingTryPhoneProvider, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("branding", "phone", "providers", id, "try"), provider, opts...)
}

// ListPhoneNotificationTemplate retrieves a list of phone notification templates.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get-phone-templates
func (m *BrandingManager) ListPhoneNotificationTemplate(ctx context.Context, opts ...RequestOption) (pnts *BrandingPhoneNotificationTemplateList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("branding", "phone", "templates"), &pnts, opts...)
	return
}

// ReadPhoneNotificationTemplate retrieves a phone notification template.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get-phone-template
func (m *BrandingManager) ReadPhoneNotificationTemplate(ctx context.Context, id string, opts ...RequestOption) (pnt *BrandingPhoneNotificationTemplate, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("branding", "phone", "templates", id), &pnt, opts...)
	return
}

// UpdatePhoneNotificationTemplate updates a phone notification template.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/update-phone-template
func (m *BrandingManager) UpdatePhoneNotificationTemplate(ctx context.Context, id string, pp *BrandingPhoneNotificationTemplate, opts ...RequestOption) (err error) {
	pp.reset("update")
	return m.management.Request(ctx, "PATCH", m.management.URI("branding", "phone", "templates", id), pp, opts...)
}

// CreatePhoneNotificationTemplate creates a phone notification template.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/create-phone-template
func (m *BrandingManager) CreatePhoneNotificationTemplate(ctx context.Context, pp *BrandingPhoneNotificationTemplate, opts ...RequestOption) (err error) {
	pp.reset("create")
	return m.management.Request(ctx, "POST", m.management.URI("branding", "phone", "templates"), pp, opts...)
}

// DeletePhoneNotificationTemplate deletes a phone notification template.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/delete-phone-template
func (m *BrandingManager) DeletePhoneNotificationTemplate(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("branding", "phone", "templates", id), nil, opts...)
}

// ResetPhoneNotificationTemplate resets a phone notification template.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/reset-phone-template
func (m *BrandingManager) ResetPhoneNotificationTemplate(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("branding", "phone", "templates", id, "reset"), nil, opts...)
}

// TryPhoneNotificationTemplate sends a test message to a phone notification template.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/try-phone-template
func (m *BrandingManager) TryPhoneNotificationTemplate(ctx context.Context, id string, template *BrandingTryPhoneNotificationTemplate, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("branding", "phone", "templates", id, "try"), template, opts...)
}
