package management

import (
	"context"
	"net/http"
)

// BrandingTheme is used to customize the login experience
// by selecting colors, fonts, and more.
type BrandingTheme struct {
	ID             *string                     `json:"themeId,omitempty"`
	DisplayName    *string                     `json:"displayName,omitempty"`
	Borders        BrandingThemeBorders        `json:"borders"`
	Colors         BrandingThemeColors         `json:"colors"`
	Fonts          BrandingThemeFonts          `json:"fonts"`
	PageBackground BrandingThemePageBackground `json:"page_background"`
	Widget         BrandingThemeWidget         `json:"widget"`
}

// BrandingThemeBorders contains borders settings for the BrandingTheme.
type BrandingThemeBorders struct {
	ButtonBorderRadius float64 `json:"button_border_radius"`
	ButtonBorderWeight float64 `json:"button_border_weight"`
	ButtonsStyle       string  `json:"buttons_style"`
	InputBorderRadius  float64 `json:"input_border_radius"`
	InputBorderWeight  float64 `json:"input_border_weight"`
	InputsStyle        string  `json:"inputs_style"`
	ShowWidgetShadow   bool    `json:"show_widget_shadow"`
	WidgetBorderWeight float64 `json:"widget_border_weight"`
	WidgetCornerRadius float64 `json:"widget_corner_radius"`
}

// BrandingThemeColors contains colors settings for the BrandingTheme.
type BrandingThemeColors struct {
	BaseFocusColor          *string `json:"base_focus_color,omitempty"`
	BaseHoverColor          *string `json:"base_hover_color,omitempty"`
	BodyText                string  `json:"body_text"`
	Error                   string  `json:"error"`
	Header                  string  `json:"header"`
	Icons                   string  `json:"icons"`
	InputBackground         string  `json:"input_background"`
	InputBorder             string  `json:"input_border"`
	InputFilledText         string  `json:"input_filled_text"`
	InputLabelsPlaceholders string  `json:"input_labels_placeholders"`
	LinksFocusedComponents  string  `json:"links_focused_components"`
	PrimaryButton           string  `json:"primary_button"`
	PrimaryButtonLabel      string  `json:"primary_button_label"`
	SecondaryButtonBorder   string  `json:"secondary_button_border"`
	SecondaryButtonLabel    string  `json:"secondary_button_label"`
	Success                 string  `json:"success"`
	WidgetBackground        string  `json:"widget_background"`
	WidgetBorder            string  `json:"widget_border"`
}

// BrandingThemeFonts contains fonts settings for the BrandingTheme.
type BrandingThemeFonts struct {
	BodyText          BrandingThemeText `json:"body_text"`
	ButtonsText       BrandingThemeText `json:"buttons_text"`
	FontURL           string            `json:"font_url"`
	InputLabels       BrandingThemeText `json:"input_labels"`
	Links             BrandingThemeText `json:"links"`
	LinksStyle        string            `json:"links_style"`
	ReferenceTextSize float64           `json:"reference_text_size"`
	Subtitle          BrandingThemeText `json:"subtitle"`
	Title             BrandingThemeText `json:"title"`
}

// BrandingThemeText contains text settings
// for the BrandingThemeFonts.
type BrandingThemeText struct {
	Bold bool    `json:"bold"`
	Size float64 `json:"size"`
}

// BrandingThemePageBackground contains page
// background settings for the BrandingTheme.
type BrandingThemePageBackground struct {
	BackgroundColor    string `json:"background_color"`
	BackgroundImageURL string `json:"background_image_url"`
	PageLayout         string `json:"page_layout"`
}

// BrandingThemeWidget contains widget settings for the BrandingTheme.
type BrandingThemeWidget struct {
	HeaderTextAlignment string  `json:"header_text_alignment"`
	LogoHeight          float64 `json:"logo_height"`
	LogoPosition        string  `json:"logo_position"`
	LogoURL             string  `json:"logo_url"`
	SocialButtonsLayout string  `json:"social_buttons_layout"`
}

// BrandingThemeManager manages Auth0 BrandingTheme resources.
type BrandingThemeManager manager

// Default retrieves the default BrandingTheme.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get_default_branding_theme
func (m *BrandingThemeManager) Default(ctx context.Context, opts ...RequestOption) (theme *BrandingTheme, err error) {
	err = m.management.Request(ctx, http.MethodGet, m.management.URI("branding", "themes", "default"), &theme, opts...)
	return
}

// Create a new BrandingTheme.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/post_branding_theme
func (m *BrandingThemeManager) Create(ctx context.Context, theme *BrandingTheme, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, http.MethodPost, m.management.URI("branding", "themes"), theme, opts...)
}

// Read retrieves a BrandingTheme.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get_branding_theme
func (m *BrandingThemeManager) Read(ctx context.Context, id string, opts ...RequestOption) (theme *BrandingTheme, err error) {
	err = m.management.Request(ctx, http.MethodGet, m.management.URI("branding", "themes", id), &theme, opts...)
	return
}

// Update a BrandingTheme.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/patch_branding_theme
func (m *BrandingThemeManager) Update(ctx context.Context, id string, theme *BrandingTheme, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, http.MethodPatch, m.management.URI("branding", "themes", id), theme, opts...)
}

// Delete a BrandingTheme.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/delete_branding_theme
func (m *BrandingThemeManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, http.MethodDelete, m.management.URI("branding", "themes", id), nil, opts...)
}
