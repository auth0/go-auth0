package management

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestBrandingThemeManager_Default(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedTheme := givenABrandingTheme(t)

	actualTheme, err := api.BrandingTheme.Default(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedTheme.GetDisplayName(), actualTheme.GetDisplayName())
}

func TestBrandingThemeManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedTheme := BrandingTheme{
		Borders: BrandingThemeBorders{
			ButtonBorderRadius: 2,
			ButtonBorderWeight: 2,
			ButtonsStyle:       "pill",
			InputBorderRadius:  2,
			InputBorderWeight:  2,
			InputsStyle:        "pill",
			ShowWidgetShadow:   true,
			WidgetBorderWeight: 2,
			WidgetCornerRadius: 2,
		},
		Colors: BrandingThemeColors{
			BodyText:                "#00FF00",
			Error:                   "#00FF00",
			CaptchaWidgetTheme:      "auto",
			Header:                  "#00FF00",
			Icons:                   "#00FF00",
			InputBackground:         "#00FF00",
			InputBorder:             "#00FF00",
			InputFilledText:         "#00FF00",
			InputLabelsPlaceholders: "#00FF00",
			LinksFocusedComponents:  "#00FF00",
			PrimaryButton:           "#00FF00",
			PrimaryButtonLabel:      "#00FF00",
			SecondaryButtonBorder:   "#00FF00",
			SecondaryButtonLabel:    "#00FF00",
			Success:                 "#00FF00",
			WidgetBackground:        "#00FF00",
			WidgetBorder:            "#00FF00",
		},
		Fonts: BrandingThemeFonts{
			BodyText: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			ButtonsText: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			FontURL: "https://google.com/font.woff",
			InputLabels: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			Links: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			LinksStyle:        "normal",
			ReferenceTextSize: 12,
			Subtitle: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			Title: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
		},
		PageBackground: BrandingThemePageBackground{
			BackgroundColor:    "#000000",
			BackgroundImageURL: "https://google.com/background.png",
			PageLayout:         "center",
		},
		Widget: BrandingThemeWidget{
			HeaderTextAlignment: "center",
			LogoHeight:          55,
			LogoPosition:        "center",
			LogoURL:             "https://google.com/logo.png",
			SocialButtonsLayout: "top",
		},
	}

	err := api.BrandingTheme.Create(context.Background(), &expectedTheme)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedTheme.GetID())

	t.Cleanup(func() {
		cleanupBrandingTheme(t, expectedTheme.GetID())
	})
}

func TestBrandingThemeManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedTheme := givenABrandingTheme(t)

	actualTheme, err := api.BrandingTheme.Read(context.Background(), expectedTheme.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedTheme.GetDisplayName(), actualTheme.GetDisplayName())
}

func TestBrandingThemeManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedTheme := givenABrandingTheme(t)

	actualTheme := &BrandingTheme{
		DisplayName: auth0.String("newTheme"),
		Borders: BrandingThemeBorders{
			ButtonBorderRadius: 2,
			ButtonBorderWeight: 2,
			ButtonsStyle:       "pill",
			InputBorderRadius:  2,
			InputBorderWeight:  2,
			InputsStyle:        "pill",
			ShowWidgetShadow:   true,
			WidgetBorderWeight: 2,
			WidgetCornerRadius: 2,
		},
		Colors: BrandingThemeColors{
			BodyText:                "#00FF00",
			Error:                   "#00FF00",
			CaptchaWidgetTheme:      "auto",
			Header:                  "#00FF00",
			Icons:                   "#00FF00",
			InputBackground:         "#00FF00",
			InputBorder:             "#00FF00",
			InputFilledText:         "#00FF00",
			InputLabelsPlaceholders: "#00FF00",
			LinksFocusedComponents:  "#00FF00",
			PrimaryButton:           "#00FF00",
			PrimaryButtonLabel:      "#00FF00",
			SecondaryButtonBorder:   "#00FF00",
			SecondaryButtonLabel:    "#00FF00",
			Success:                 "#00FF00",
			WidgetBackground:        "#00FF00",
			WidgetBorder:            "#00FF00",
		},
		Fonts: BrandingThemeFonts{
			BodyText: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			ButtonsText: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			FontURL: "https://google.com/font.woff",
			InputLabels: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			Links: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			LinksStyle:        "normal",
			ReferenceTextSize: 12,
			Subtitle: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			Title: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
		},
		PageBackground: BrandingThemePageBackground{
			BackgroundColor:    "#000000",
			BackgroundImageURL: "https://google.com/background.png",
			PageLayout:         "center",
		},
		Widget: BrandingThemeWidget{
			HeaderTextAlignment: "center",
			LogoHeight:          55,
			LogoPosition:        "center",
			LogoURL:             "https://google.com/logo.png",
			SocialButtonsLayout: "top",
		},
	}

	err := api.BrandingTheme.Update(context.Background(), expectedTheme.GetID(), actualTheme)
	assert.NoError(t, err)

	actualTheme, err = api.BrandingTheme.Read(context.Background(), expectedTheme.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "newTheme", actualTheme.GetDisplayName())
}

func TestBrandingThemeManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	theme := givenABrandingTheme(t)

	err := api.BrandingTheme.Delete(context.Background(), theme.GetID())
	assert.NoError(t, err)

	actualTheme, err := api.BrandingTheme.Read(context.Background(), theme.GetID())
	assert.Empty(t, actualTheme)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenABrandingTheme(t *testing.T) *BrandingTheme {
	t.Helper()

	theme := &BrandingTheme{
		Borders: BrandingThemeBorders{
			ButtonBorderRadius: 2,
			ButtonBorderWeight: 2,
			ButtonsStyle:       "pill",
			InputBorderRadius:  2,
			InputBorderWeight:  2,
			InputsStyle:        "pill",
			ShowWidgetShadow:   true,
			WidgetBorderWeight: 2,
			WidgetCornerRadius: 2,
		},
		Colors: BrandingThemeColors{
			BodyText:                "#00FF00",
			Error:                   "#00FF00",
			CaptchaWidgetTheme:      "auto",
			Header:                  "#00FF00",
			Icons:                   "#00FF00",
			InputBackground:         "#00FF00",
			InputBorder:             "#00FF00",
			InputFilledText:         "#00FF00",
			InputLabelsPlaceholders: "#00FF00",
			LinksFocusedComponents:  "#00FF00",
			PrimaryButton:           "#00FF00",
			PrimaryButtonLabel:      "#00FF00",
			SecondaryButtonBorder:   "#00FF00",
			SecondaryButtonLabel:    "#00FF00",
			Success:                 "#00FF00",
			WidgetBackground:        "#00FF00",
			WidgetBorder:            "#00FF00",
		},
		Fonts: BrandingThemeFonts{
			BodyText: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			ButtonsText: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			FontURL: "https://google.com/font.woff",
			InputLabels: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			Links: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			LinksStyle:        "normal",
			ReferenceTextSize: 12,
			Subtitle: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
			Title: BrandingThemeText{
				Bold: true,
				Size: 100,
			},
		},
		PageBackground: BrandingThemePageBackground{
			BackgroundColor:    "#000000",
			BackgroundImageURL: "https://google.com/background.png",
			PageLayout:         "center",
		},
		Widget: BrandingThemeWidget{
			HeaderTextAlignment: "center",
			LogoHeight:          55,
			LogoPosition:        "center",
			LogoURL:             "https://google.com/logo.png",
			SocialButtonsLayout: "top",
		},
	}

	err := api.BrandingTheme.Create(context.Background(), theme)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupBrandingTheme(t, theme.GetID())
	})

	return theme
}

func cleanupBrandingTheme(t *testing.T, themeID string) {
	t.Helper()

	err := api.BrandingTheme.Delete(context.Background(), themeID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
