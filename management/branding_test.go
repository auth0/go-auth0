package management

import (
	"context"
	"encoding/json"
	"net/http"
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

func TestBrandingManager_ListPhoneProviders(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedProvider := givenAnBrandingPhoneProvider(t)
	actualProviders, err := api.Branding.ListPhoneProviders(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, actualProviders)
	assert.Equal(t, actualProviders.Providers[0].GetID(), expectedProvider.GetID())
}

func TestBrandingManager_CreatePhoneProvider(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedProvider := &BrandingPhoneProvider{
		Name:     auth0.String("custom"),
		Disabled: auth0.Bool(false),
		Configuration: &BrandingPhoneProviderConfiguration{
			DeliveryMethods: &[]string{"text"},
		},
		Credentials: &BrandingPhoneProviderCredential{},
	}

	err := api.Branding.CreatePhoneProvider(context.Background(), expectedProvider)
	assert.NoError(t, err)

	actualProvider, err := api.Branding.ReadPhoneProvider(context.Background(), expectedProvider.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedProvider, actualProvider)

	t.Cleanup(func() {
		cleanupBrandingPhoneProvider(t, expectedProvider.GetID())
	})
}

func TestBrandingManager_ReadPhoneProvider(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedProvider := givenAnBrandingPhoneProvider(t)

	actualProvider, err := api.Branding.ReadPhoneProvider(context.Background(), expectedProvider.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedProvider, actualProvider)
}

func TestBrandingManager_UpdatePhoneProvider(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedProvider := givenAnBrandingPhoneProvider(t)

	updatedProvider := &BrandingPhoneProvider{
		Name:     auth0.String("custom"),
		Disabled: auth0.Bool(false),
		Configuration: &BrandingPhoneProviderConfiguration{
			DeliveryMethods: &[]string{"text"},
		},
		Credentials: &BrandingPhoneProviderCredential{},
	}

	err := api.Branding.UpdatePhoneProvider(context.Background(), expectedProvider.GetID(), updatedProvider)
	assert.NoError(t, err)

	actualProvider, err := api.Branding.ReadPhoneProvider(context.Background(), expectedProvider.GetID())
	assert.NoError(t, err)
	assert.Equal(t, updatedProvider, actualProvider)
}

func TestBrandingManager_DeletePhoneProvider(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedProvider := givenAnBrandingPhoneProvider(t)

	err := api.Branding.DeletePhoneProvider(context.Background(), expectedProvider.GetID())
	assert.NoError(t, err)
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

func TestBrandingPhoneProvider_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    *BrandingPhoneProvider
		wantErr bool
	}{
		{
			name:    "Valid JSON with credentials",
			jsonStr: `{"name":"twilio","disabled":false,"configuration":{"delivery_methods":["text"],"default_from":"1234567890","sid":"sid"},"credentials":{"auth_token":"auth_token"}}`,
			want: &BrandingPhoneProvider{
				Name:     auth0.String("twilio"),
				Disabled: auth0.Bool(false),
				Configuration: &BrandingPhoneProviderConfiguration{
					DeliveryMethods: &[]string{"text"},
					DefaultFrom:     auth0.String("1234567890"),
					SID:             auth0.String("sid"),
				},
				Credentials: &BrandingPhoneProviderCredential{
					AuthToken: auth0.String("auth_token"),
				},
			},
			wantErr: false,
		},
		{
			name:    "Valid JSON without credentials",
			jsonStr: `{"name":"twilio","disabled":false,"configuration":{"delivery_methods":["text"],"default_from":"1234567890","sid":"sid"}}`,
			want: &BrandingPhoneProvider{
				Name:     auth0.String("twilio"),
				Disabled: auth0.Bool(false),
				Configuration: &BrandingPhoneProviderConfiguration{
					DeliveryMethods: &[]string{"text"},
					DefaultFrom:     auth0.String("1234567890"),
					SID:             auth0.String("sid"),
				},
				Credentials: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got BrandingPhoneProvider
			err := json.Unmarshal([]byte(tt.jsonStr), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, &got)
		})
	}
}

func givenAnBrandingPhoneProvider(t *testing.T) *BrandingPhoneProvider {
	t.Helper()

	provider := &BrandingPhoneProvider{
		Name:     auth0.String("twilio"),
		Disabled: auth0.Bool(false),
		Configuration: &BrandingPhoneProviderConfiguration{
			DeliveryMethods: &[]string{"text"},
			DefaultFrom:     auth0.String("1234567890"),
			SID:             auth0.String("sid"),
		},
		Credentials: &BrandingPhoneProviderCredential{
			AuthToken: auth0.String("auth_token"),
		},
	}

	err := api.Branding.CreatePhoneProvider(context.Background(), provider)
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		cleanupBrandingPhoneProvider(t, provider.GetID())
	})
	return provider
}

func cleanupBrandingPhoneProvider(t *testing.T, providerID string) {
	t.Helper()

	err := api.Branding.DeletePhoneProvider(context.Background(), providerID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
