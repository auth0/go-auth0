package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestPrompt(t *testing.T) {
	configureHTTPTestRecordings(t)

	t.Cleanup(func() {
		err := api.Prompt.Update(context.Background(), &Prompt{
			UniversalLoginExperience:    "classic",
			IdentifierFirst:             auth0.Bool(false),
			WebAuthnPlatformFirstFactor: auth0.Bool(false),
		})
		require.NoError(t, err)
	})

	t.Run("Update to the new identifier first experience with biometrics", func(t *testing.T) {
		err := api.Prompt.Update(context.Background(), &Prompt{
			UniversalLoginExperience:    "new",
			IdentifierFirst:             auth0.Bool(true),
			WebAuthnPlatformFirstFactor: auth0.Bool(true),
		})
		assert.NoError(t, err)

		ps, err := api.Prompt.Read(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "new", ps.UniversalLoginExperience)
		assert.True(t, ps.GetIdentifierFirst())
		assert.True(t, ps.GetWebAuthnPlatformFirstFactor())
	})

	t.Run("Update to the classic non identifier first experience without biometrics", func(t *testing.T) {
		err := api.Prompt.Update(context.Background(), &Prompt{
			UniversalLoginExperience:    "classic",
			IdentifierFirst:             auth0.Bool(false),
			WebAuthnPlatformFirstFactor: auth0.Bool(false),
		})
		assert.NoError(t, err)

		ps, err := api.Prompt.Read(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "classic", ps.UniversalLoginExperience)
		assert.False(t, ps.GetIdentifierFirst())
		assert.False(t, ps.GetWebAuthnPlatformFirstFactor())
	})
}

func TestPromptCustomText(t *testing.T) {
	configureHTTPTestRecordings(t)

	const prompt = "login"
	const lang = "en"

	t.Cleanup(func() {
		body := make(map[string]interface{})
		err := api.Prompt.SetCustomText(context.Background(), prompt, lang, body)
		require.NoError(t, err)
	})

	body := map[string]interface{}{
		"login": map[string]interface{}{
			"title": "Welcome",
		},
	}

	err := api.Prompt.SetCustomText(context.Background(), prompt, lang, body)
	assert.NoError(t, err)

	texts, err := api.Prompt.CustomText(context.Background(), prompt, lang)
	assert.NoError(t, err)
	assert.Equal(t, "Welcome", texts["login"].(map[string]interface{})["title"])
}

func TestPromptManager_GetPartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	prompt := PromptLogin

	expected := givenAPartialPrompt(t, prompt)

	actual, err := api.Prompt.GetPartials(context.Background(), prompt)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestPromptManager_SetPartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	prompt := PromptLoginPasswordLess

	expected := &PromptScreenPartials{
		ScreenLoginPasswordlessSMSOTP: {
			InsertionPointFormContentStart: `<div>Form Content Start</div>`,
			InsertionPointFormContentEnd:   `<div>Form Content Start</div>`,
		},
		ScreenLoginPasswordlessEmailCode: {
			InsertionPointFormContentStart: `<div>Form Content Start</div>`,
			InsertionPointFormContentEnd:   `<div>Form Content Start</div>`,
		},
	}

	err := api.Prompt.SetPartials(context.Background(), prompt, expected)
	assert.NoError(t, err)

	t.Cleanup(func() {
		cleanupPromptPartials(t, prompt)
	})

	actual, err := api.Prompt.GetPartials(context.Background(), prompt)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestPromptManager_ReadACULSettings(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)
	expected := givenAACULSettings(t)
	actual, err := api.Prompt.ReadACULSettings(context.Background(), PromptSignup, ScreenSignup)
	assert.NoError(t, err)
	assert.Equal(t, expected.RenderingMode, actual.RenderingMode)
	assert.Equal(t, expected.ContextConfiguration, actual.ContextConfiguration)
	assert.Equal(t, expected.DefaultHeadTagsDisabled, actual.DefaultHeadTagsDisabled)
	assert.Equal(t, expected.HeadTags, actual.HeadTags)
	assert.Equal(t, string(PromptSignup), *actual.Prompt)
	assert.Equal(t, string(ScreenSignup), *actual.Screen)
}

func TestPromptManager_UpdateACULSettings(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)
	expected := givenAACULSettings(t)
	expected.RenderingMode = auth0.String("standard")
	expected.ContextConfiguration = &[]string{"branding.settings", "branding.themes.default", "client.logo_uri"}
	expected.DefaultHeadTagsDisabled = auth0.Bool(true)

	err := api.Prompt.UpdateACULSettings(context.Background(), PromptSignup, ScreenSignup, expected)
	assert.NoError(t, err)

	actual, err := api.Prompt.ReadACULSettings(context.Background(), PromptSignup, ScreenSignup)
	assert.NoError(t, err)
	assert.Equal(t, expected.RenderingMode, actual.RenderingMode)
	assert.Equal(t, expected.ContextConfiguration, actual.ContextConfiguration)
	assert.Equal(t, expected.DefaultHeadTagsDisabled, actual.DefaultHeadTagsDisabled)
	assert.Equal(t, expected.HeadTags, actual.HeadTags)
	assert.Equal(t, string(PromptSignup), *actual.Prompt)
	assert.Equal(t, string(ScreenSignup), *actual.Screen)
}

func TestPromptManager_GetPartialsGuardGuardError(t *testing.T) {
	configureHTTPTestRecordings(t)

	prompt := PromptType("OtherPrompt")

	_, err := api.Prompt.GetPartials(context.Background(), prompt)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "cannot customize partials for prompt")
}

func TestPromptManager_SetPartialsGuardGuardError(t *testing.T) {
	configureHTTPTestRecordings(t)

	prompt := PromptType("OtherPrompt")

	expected := &PromptScreenPartials{}

	err := api.Prompt.SetPartials(context.Background(), prompt, expected)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "cannot customize partials for prompt")
}

func TestPromptManager_ReadPartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	prompt := PromptSignup
	expected := &PromptPartials{Prompt: prompt}
	actual, err := api.Prompt.ReadPartials(context.Background(), prompt)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestPromptManager_CreatePartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	prompt := PromptSignup
	original := &PromptPartials{Prompt: prompt}
	expected := &PromptPartials{Prompt: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.Prompt.CreatePartials(context.Background(), expected)
	assert.NoError(t, err)

	got, err := api.Prompt.ReadPartials(context.Background(), prompt)
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, got)

	t.Cleanup(func() {
		deletePromptPartials(t, prompt)
	})
}

func TestPromptManager_UpdatePartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	prompt := PromptSignup
	original := &PromptPartials{Prompt: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.Prompt.CreatePartials(context.Background(), original)
	assert.NoError(t, err)

	expected := &PromptPartials{Prompt: prompt, FormContentStart: `<div>Updated Test Content</div>`}
	err = api.Prompt.UpdatePartials(context.Background(), expected)
	assert.NoError(t, err)

	got, err := api.Prompt.ReadPartials(context.Background(), prompt)
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, expected)

	t.Cleanup(func() {
		deletePromptPartials(t, prompt)
	})
}

func TestPromptManager_DeletePartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	prompt := PromptSignup
	original := &PromptPartials{Prompt: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.Prompt.CreatePartials(context.Background(), original)
	assert.NoError(t, err)

	expected := &PromptPartials{Prompt: prompt}
	err = api.Prompt.DeletePartials(context.Background(), prompt)
	assert.NoError(t, err)

	got, err := api.Prompt.ReadPartials(context.Background(), prompt)
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, expected)

	t.Cleanup(func() {
		deletePromptPartials(t, prompt)
	})
}

func givenAUniversalLoginTemplate(t *testing.T) *BrandingUniversalLogin {
	t.Helper()

	body := `<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>`
	ul := &BrandingUniversalLogin{
		Body: auth0.String(body),
	}

	err := api.Branding.SetUniversalLogin(context.Background(), ul)
	assert.NoError(t, err)

	t.Cleanup(func() {
		deleteUniversalLoginTemplate(t)
	})

	return ul
}

func deletePromptPartials(t *testing.T, prompt PromptType) {
	t.Helper()

	err := api.Prompt.DeletePartials(context.Background(), prompt)
	assert.NoError(t, err)
}

func deleteUniversalLoginTemplate(t *testing.T) {
	t.Helper()

	err := api.Branding.DeleteUniversalLogin(context.Background())
	assert.NoError(t, err)
}

func givenAPartialPrompt(t *testing.T, prompt PromptType) *PromptScreenPartials {
	t.Helper()

	partials := &PromptScreenPartials{
		ScreenLogin: {
			InsertionPointFormContentStart: `<div>Form Content Start</div>`,
			InsertionPointFormContentEnd:   `<div>Form Content Start</div>`,
		},
	}

	err := api.Prompt.SetPartials(context.Background(), prompt, partials)
	assert.NoError(t, err)

	t.Cleanup(func() {
		cleanupPromptPartials(t, prompt)
	})

	return partials
}

func givenAACULSettings(t *testing.T) *PromptACULSettings {
	t.Helper()

	settings := &PromptACULSettings{
		RenderingMode:           auth0.String("advanced"),
		ContextConfiguration:    &[]string{"branding.settings", "branding.themes.default"},
		DefaultHeadTagsDisabled: auth0.Bool(false),
		HeadTags: []interface{}{
			map[string]interface{}{
				"tag":     "script",
				"content": "",
				"attributes": map[string]interface{}{
					"defer": true,
					"src":   "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.js",
					"async": true,
					"integrity": []string{
						"sha512-v2CJ7UaYy4JwqLDIrZUI/4hqeoQieOmAZNXBeQyjo21dadnwR+8ZaIJVT8EE2iyI61OV8e6M8PP2/4hpQINQ/g==",
					},
				},
			},
		},
	}

	err := api.Prompt.UpdateACULSettings(context.Background(), PromptSignup, ScreenSignup, settings)
	assert.NoError(t, err)

	return settings
}

func cleanupPromptPartials(t *testing.T, prompt PromptType) {
	t.Helper()

	err := api.Prompt.DeletePartials(context.Background(), prompt)
	assert.NoError(t, err)
}
