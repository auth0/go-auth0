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

func TestPromptManager_ReadPartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	_ = givenAUniversalLogin(t)
	prompt := PartialsPromptSignup
	expected := &PartialsPrompt{Segment: prompt}
	got, err := api.Prompt.ReadPartials(context.Background(), prompt)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	t.Cleanup(func() {
		cleanupUniversalLogin(t)
		cleanupCustomDomain(t, customDomain.GetID())
	})
}

func TestPromptManager_CreatePartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	_ = givenAUniversalLogin(t)
	prompt := PartialsPromptSignup
	original := &PartialsPrompt{Segment: prompt}
	expected := &PartialsPrompt{Segment: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.Prompt.CreatePartials(context.Background(), expected)
	assert.NoError(t, err)

	got, err := api.Prompt.ReadPartials(context.Background(), prompt)
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, got)

	t.Cleanup(func() {
		cleanupPartialsPrompt(t, customDomain.GetID(), prompt)
	})
}

func TestPromptManager_UpdatePartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	_ = givenAUniversalLogin(t)
	prompt := PartialsPromptSignup
	original := &PartialsPrompt{Segment: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.Prompt.CreatePartials(context.Background(), original)
	assert.NoError(t, err)

	expected := &PartialsPrompt{Segment: prompt, FormContentStart: `<div>Updated Test Content</div>`}
	err = api.Prompt.UpdatePartials(context.Background(), expected)
	assert.NoError(t, err)

	got, err := api.Prompt.ReadPartials(context.Background(), prompt)
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, expected)

	t.Cleanup(func() {
		cleanupPartialsPrompt(t, customDomain.GetID(), prompt)
	})
}

func TestPromptManager_DeletePartials(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	_ = givenAUniversalLogin(t)
	prompt := PartialsPromptSignup
	original := &PartialsPrompt{Segment: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.Prompt.CreatePartials(context.Background(), original)
	assert.NoError(t, err)

	expected := &PartialsPrompt{Segment: prompt}
	err = api.Prompt.DeletePartials(context.Background(), expected)
	assert.NoError(t, err)

	got, err := api.Prompt.ReadPartials(context.Background(), prompt)
	assert.NoError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, expected)

	t.Cleanup(func() {
		cleanupPartialsPrompt(t, customDomain.GetID(), prompt)
	})
}

func givenAUniversalLogin(t *testing.T) *BrandingUniversalLogin {
	t.Helper()

	body := `<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>`
	ul := &BrandingUniversalLogin{
		Body: auth0.String(body),
	}

	err := api.Branding.SetUniversalLogin(context.Background(), ul)
	assert.NoError(t, err)

	return ul
}

func cleanupPartialsPrompt(t *testing.T, customDomainID string, prompt PartialsPromptSegment) {
	t.Helper()

	c := &PartialsPrompt{Segment: prompt}
	err := api.Prompt.DeletePartials(context.Background(), c)
	assert.NoError(t, err)

	cleanupUniversalLogin(t)
	cleanupCustomDomain(t, customDomainID)
}

func cleanupUniversalLogin(t *testing.T) {
	t.Helper()

	err := api.Branding.DeleteUniversalLogin(context.Background())
	assert.NoError(t, err)
}
