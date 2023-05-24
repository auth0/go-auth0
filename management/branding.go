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
