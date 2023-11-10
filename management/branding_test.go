package management

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestBrandingManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	branding, err := api.Branding.Read(context.Background())
	assert.NoError(t, err)
	assert.IsType(t, &Branding{}, branding)
}

func TestBrandingManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Save initial branding settings.
	preTestBrandingSettings, err := api.Branding.Read(context.Background())
	assert.NoError(t, err)

	expected := &Branding{
		Colors: &BrandingColors{
			Primary: auth0.String("#ea5323"),
			PageBackgroundGradient: &BrandingPageBackgroundGradient{
				Type:        auth0.String("linear-gradient"),
				Start:       auth0.String("#000000"),
				End:         auth0.String("#ffffff"),
				AngleDegree: auth0.Int(35),
			},
		},
		FaviconURL: auth0.String("https://mycompany.org/favicon.ico"),
		LogoURL:    auth0.String("https://mycompany.org/logo.png"),
		Font: &BrandingFont{
			URL: auth0.String("https://mycompany.org/font.otf"),
		},
	}

	err = api.Branding.Update(context.Background(), expected)
	assert.NoError(t, err)

	actual, err := api.Branding.Read(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expected.GetColors().GetPrimary(), actual.GetColors().GetPrimary())
	assert.Equal(t, expected.GetFont().GetURL(), actual.GetFont().GetURL())
	assert.Equal(t, expected.GetFaviconURL(), actual.GetFaviconURL())

	// Restore initial branding settings.
	err = api.Branding.Update(context.Background(), preTestBrandingSettings)
	assert.NoError(t, err)
}

func TestBrandingManager_UniversalLogin(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenACustomDomain(t)

	body := `<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>`
	expectedUL := &BrandingUniversalLogin{
		Body: auth0.String(body),
	}

	err := api.Branding.SetUniversalLogin(context.Background(), expectedUL)
	assert.NoError(t, err)

	actualUL, err := api.Branding.UniversalLogin(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedUL, actualUL)

	t.Cleanup(func() {
		err = api.Branding.DeleteUniversalLogin(context.Background())
		assert.NoError(t, err)
	})
}

func TestBrandingColors(t *testing.T) {
	var testCases = []struct {
		name   string
		colors *BrandingColors
		expect string
	}{
		{
			name: "PageBackground",
			colors: &BrandingColors{
				Primary:        auth0.String("#ea5323"),
				PageBackground: auth0.String("#000000"),
			},
			expect: `{"primary":"#ea5323","page_background":"#000000"}`,
		},
		{
			name: "PageBackgroundGradient",
			colors: &BrandingColors{
				Primary: auth0.String("#ea5323"),
				PageBackgroundGradient: &BrandingPageBackgroundGradient{
					Type:        auth0.String("linear-gradient"),
					Start:       auth0.String("#000000"),
					End:         auth0.String("#ffffff"),
					AngleDegree: auth0.Int(35),
				},
			},
			expect: `{"primary":"#ea5323","page_background":{"type":"linear-gradient","start":"#000000","end":"#ffffff","angle_deg":35}}`,
		},
		{
			name: "PageBackgroundNil",
			colors: &BrandingColors{
				Primary: auth0.String("#ea5323"),
			},
			expect: `{"primary":"#ea5323"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b, err := json.Marshal(testCase.colors)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expect, string(b))

			var colors BrandingColors
			err = json.Unmarshal([]byte(testCase.expect), &colors)
			assert.NoError(t, err)
			assert.Equal(t, testCase.colors, &colors)
		})
	}

	t.Run("Should error is not expected type", func(t *testing.T) {
		var actual BrandingColors
		err := json.Unmarshal([]byte(`{"page_background":123}`), &actual)
		assert.Contains(t, err.Error(), "unexpected type for field page_background")
	})

	t.Run("Should disallow setting PageBackground and PageBackgroundGradient", func(t *testing.T) {
		_, err := json.Marshal(&BrandingColors{
			PageBackground: auth0.String("#ffffff"),
			PageBackgroundGradient: &BrandingPageBackgroundGradient{
				Type:        auth0.String("linear-gradient"),
				Start:       auth0.String("#ffffff"),
				End:         auth0.String("#000000"),
				AngleDegree: auth0.Int(3),
			},
		})
		assert.Contains(t, err.Error(), "only one of PageBackground and PageBackgroundGradient is allowed")
	})
}
