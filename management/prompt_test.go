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

func TestPromptManager_ReadRendering(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)

	client := givenAClient(t)
	expected := givenAPromptRendering(t, RenderingModeAdvanced, client.GetClientID())
	actual, err := api.Prompt.ReadRendering(context.Background(), PromptSignup, ScreenSignup)
	assert.NoError(t, err)
	assert.Equal(t, expected.GetRenderingMode(), actual.GetRenderingMode())
	assert.Equal(t, expected.GetContextConfiguration(), actual.GetContextConfiguration())
	assert.Equal(t, expected.GetDefaultHeadTagsDisabled(), actual.GetDefaultHeadTagsDisabled())
	assert.Equal(t, expected.GetFilters(), actual.GetFilters())
	assert.Equal(t, expected.GetUsePageTemplate(), actual.GetUsePageTemplate())
	assert.Equal(t, expected.HeadTags, actual.HeadTags)
	assert.Equal(t, PromptSignup, *actual.GetPrompt())
	assert.Equal(t, ScreenSignup, *actual.GetScreen())
}

func TestPromptManager_ListRendering(t *testing.T) {
	configureHTTPTestRecordings(t)
	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)
	client := givenAClient(t)
	expected := givenAPromptRendering(t, RenderingModeAdvanced, client.GetClientID())

	actual, err := api.Prompt.ListRendering(context.Background())
	assert.NoError(t, err)

	found := false

	for _, r := range actual.PromptRenderings {
		if r.RenderingMode != nil && expected.RenderingMode != nil &&
			*r.GetRenderingMode() == *expected.GetRenderingMode() &&
			assert.Equal(t, expected.GetContextConfiguration(), r.GetContextConfiguration()) &&
			assert.Equal(t, expected.GetDefaultHeadTagsDisabled(), r.GetDefaultHeadTagsDisabled()) &&
			assert.Equal(t, expected.GetFilters(), r.GetFilters()) &&
			assert.Equal(t, expected.GetUsePageTemplate(), r.GetUsePageTemplate()) &&
			assert.Equal(t, expected.HeadTags, r.HeadTags) {
			found = true
			break
		}
	}

	assert.True(t, found, "expected PromptRendering not found in actual.PromptRendering")
	assert.Greater(t, len(actual.PromptRenderings), 0)
}

// Able to update the renderingMode to advanced and the setting configs when parsing the advanced renderingMode in payload.
func TestPromptManager_UpdateRenderingWithAdvancedMode(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)
	client := givenAClient(t)
	_ = givenAPromptRendering(t, RenderingModeAdvanced, client.GetClientID())

	updateData := &PromptRendering{
		RenderingMode:           &RenderingModeAdvanced,
		ContextConfiguration:    &[]string{"branding.settings", "branding.themes.default", "client.logo_uri"},
		DefaultHeadTagsDisabled: auth0.Bool(true),
		Filters: &PromptRenderingFilters{
			MatchType: auth0.String("excludes_any"),
			Clients: &[]PromptRenderingFilter{
				{
					ID: auth0.String(client.GetClientID()),
					Metadata: &map[string]interface{}{
						"key2": "value2",
					},
				},
			},
		},
		UsePageTemplate: auth0.Bool(false),
	}

	err := api.Prompt.UpdateRendering(context.Background(), PromptSignup, ScreenSignup, updateData)
	assert.NoError(t, err)

	actual, err := api.Prompt.ReadRendering(context.Background(), PromptSignup, ScreenSignup)
	assert.NoError(t, err)
	assert.Equal(t, updateData.GetContextConfiguration(), actual.GetContextConfiguration())
	assert.Equal(t, updateData.GetDefaultHeadTagsDisabled(), actual.GetDefaultHeadTagsDisabled())
	assert.Equal(t, updateData.GetFilters(), actual.GetFilters())
	assert.Equal(t, updateData.GetUsePageTemplate(), actual.GetUsePageTemplate())
	assert.Equal(t, PromptSignup, *actual.GetPrompt())
	assert.Equal(t, ScreenSignup, *actual.GetScreen())
}

// Unable to update the setting configs and only able to update the renderingMode to standard when parsing the standard renderingMode in payload.
func TestPromptManager_UpdateRenderingWithStandardMode(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)
	client := givenAClient(t)
	expected := givenAPromptRendering(t, RenderingModeAdvanced, client.GetClientID())

	updateData := &PromptRendering{
		RenderingMode:           &RenderingModeStandard,
		ContextConfiguration:    &[]string{"branding.settings", "branding.themes.default", "client.logo_uri"},
		DefaultHeadTagsDisabled: auth0.Bool(true),
		Filters: &PromptRenderingFilters{
			MatchType: auth0.String("excludes_any"),
			Clients: &[]PromptRenderingFilter{
				{
					ID: auth0.String(client.GetClientID()),
					Metadata: &map[string]interface{}{
						"key2": "value2",
					},
				},
			},
		},
		UsePageTemplate: auth0.Bool(false),
	}

	err := api.Prompt.UpdateRendering(context.Background(), PromptSignup, ScreenSignup, updateData)
	assert.NoError(t, err)

	actual, err := api.Prompt.ReadRendering(context.Background(), PromptSignup, ScreenSignup)
	assert.NoError(t, err)

	assert.Equal(t, updateData.GetRenderingMode(), actual.GetRenderingMode())
	assert.NotEqual(t, updateData.GetContextConfiguration(), actual.GetContextConfiguration())
	assert.NotEqual(t, updateData.GetDefaultHeadTagsDisabled(), actual.GetDefaultHeadTagsDisabled())
	assert.NotEqual(t, updateData.GetFilters(), actual.GetFilters())
	assert.NotEqual(t, updateData.GetUsePageTemplate(), actual.GetUsePageTemplate())
	assert.Equal(t, expected.HeadTags, actual.HeadTags)
	assert.Equal(t, PromptSignup, *actual.GetPrompt())
	assert.Equal(t, ScreenSignup, *actual.GetScreen())
}

// Able to update the setting's configs even the existing renderingMode is standard since renderingMode is not parsed in payload(updateData).
func TestPromptManager_UpdateRendering(t *testing.T) {
	configureHTTPTestRecordings(t)

	_ = givenACustomDomain(t)
	_ = givenAUniversalLoginTemplate(t)
	client := givenAClient(t)
	_ = givenAPromptRendering(t, RenderingModeAdvanced, client.GetClientID())

	updateData := &PromptRendering{
		ContextConfiguration:    &[]string{"branding.settings", "branding.themes.default", "client.logo_uri"},
		DefaultHeadTagsDisabled: auth0.Bool(true),
		Filters: &PromptRenderingFilters{
			MatchType: auth0.String("excludes_any"),
			Clients: &[]PromptRenderingFilter{
				{
					ID: auth0.String(client.GetClientID()),
					Metadata: &map[string]interface{}{
						"key2": "value2",
					},
				},
			},
		},
		UsePageTemplate: auth0.Bool(false),
	}

	err := api.Prompt.UpdateRendering(context.Background(), PromptSignup, ScreenSignup, updateData)
	assert.NoError(t, err)

	actual, err := api.Prompt.ReadRendering(context.Background(), PromptSignup, ScreenSignup)
	assert.NoError(t, err)
	assert.Equal(t, updateData.GetContextConfiguration(), actual.GetContextConfiguration())
	assert.Equal(t, updateData.GetDefaultHeadTagsDisabled(), actual.GetDefaultHeadTagsDisabled())
	assert.Equal(t, updateData.GetFilters(), actual.GetFilters())
	assert.Equal(t, updateData.GetUsePageTemplate(), actual.GetUsePageTemplate())
	assert.Equal(t, PromptSignup, *actual.GetPrompt())
	assert.Equal(t, ScreenSignup, *actual.GetScreen())
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

func givenAPromptRendering(t *testing.T, mode RenderingMode, clientID string) *PromptRendering {
	t.Helper()

	settings := &PromptRendering{
		RenderingMode:           &mode,
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
					"integrity": []interface{}{
						"sha512-v2CJ7UaYy4JwqLDIrZUI/4hqeoQieOmAZNXBeQyjo21dadnwR+8ZaIJVT8EE2iyI61OV8e6M8PP2/4hpQINQ/g==",
					},
				},
			},
		},
		Filters: &PromptRenderingFilters{
			MatchType: auth0.String("includes_any"),
			Clients: &[]PromptRenderingFilter{
				{
					ID: auth0.String(clientID),
					Metadata: &map[string]interface{}{
						"key1": "value1",
					},
				},
			},
		},
		UsePageTemplate: auth0.Bool(true),
	}

	err := api.Prompt.UpdateRendering(context.Background(), PromptSignup, ScreenSignup, settings)
	assert.NoError(t, err)

	return settings
}

func cleanupPromptPartials(t *testing.T, prompt PromptType) {
	t.Helper()

	err := api.Prompt.DeletePartials(context.Background(), prompt)
	assert.NoError(t, err)
}
