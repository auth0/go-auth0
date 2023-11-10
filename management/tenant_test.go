package management

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestTenantManager(t *testing.T) {
	configureHTTPTestRecordings(t)

	initialSettings, err := api.Tenant.Read(context.Background())
	assert.NoError(t, err)

	t.Cleanup(func() {
		initialSettings.SandboxVersionAvailable = nil
		initialSettings.UniversalLogin = nil
		initialSettings.Flags = nil
		err := api.Tenant.Update(context.Background(), initialSettings)
		require.NoError(t, err)
	})

	newTenantSettings := &Tenant{
		FriendlyName:          auth0.String("My Example Tenant"),
		SupportURL:            auth0.String("https://support.example.com"),
		SupportEmail:          auth0.String("support@example.com"),
		DefaultRedirectionURI: auth0.String("https://example.com/login"),
		SessionLifetime:       auth0.Float64(1080),
		IdleSessionLifetime:   auth0.Float64(720.2), // will be rounded off
		SessionCookie: &TenantSessionCookie{
			Mode: auth0.String("non-persistent"),
		},
		AllowedLogoutURLs:       &[]string{"https://app.com/logout", "http://localhost/logout"},
		EnabledLocales:          &[]string{"fr", "en", "es"},
		SandboxVersionAvailable: nil,
		Sessions: &TenantSessions{
			OIDCLogoutPromptEnabled: auth0.Bool(false),
		},
	}
	err = api.Tenant.Update(context.Background(), newTenantSettings)
	assert.NoError(t, err)

	actualTenantSettings, err := api.Tenant.Read(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, newTenantSettings.GetFriendlyName(), actualTenantSettings.GetFriendlyName())
	assert.Equal(t, newTenantSettings.GetIdleSessionLifetime(), actualTenantSettings.GetIdleSessionLifetime())
	assert.Equal(t, newTenantSettings.GetIdleSessionLifetime(), 720.0) // it got rounded off
	assert.Equal(t, newTenantSettings.GetSessionLifetime(), actualTenantSettings.GetSessionLifetime())
	assert.Equal(t, newTenantSettings.GetSupportEmail(), actualTenantSettings.GetSupportEmail())
	assert.Equal(t, newTenantSettings.GetSupportURL(), actualTenantSettings.GetSupportURL())
	assert.Equal(t, newTenantSettings.GetSessionCookie().GetMode(), actualTenantSettings.GetSessionCookie().GetMode())
	assert.Equal(t, newTenantSettings.GetAllowedLogoutURLs(), actualTenantSettings.GetAllowedLogoutURLs())
	assert.Equal(t, newTenantSettings.GetEnabledLocales(), actualTenantSettings.GetEnabledLocales())
	assert.Equal(t, newTenantSettings.GetSandboxVersion(), actualTenantSettings.GetSandboxVersion())
	assert.Equal(t, newTenantSettings.GetSessions().GetOIDCLogoutPromptEnabled(), actualTenantSettings.GetSessions().GetOIDCLogoutPromptEnabled())
}

func TestTenant_MarshalJSON(t *testing.T) {
	for tenant, expected := range map[*Tenant]string{
		{}:                                         `{}`,
		{SessionLifetime: auth0.Float64(1.2)}:      `{"session_lifetime":1}`,
		{SessionLifetime: auth0.Float64(1.19)}:     `{"session_lifetime":1}`,
		{SessionLifetime: auth0.Float64(1)}:        `{"session_lifetime":1}`,
		{SessionLifetime: auth0.Float64(720)}:      `{"session_lifetime":720}`,
		{IdleSessionLifetime: auth0.Float64(1)}:    `{"idle_session_lifetime":1}`,
		{IdleSessionLifetime: auth0.Float64(1.2)}:  `{"idle_session_lifetime":1}`,
		{SessionLifetime: auth0.Float64(0.25)}:     `{"session_lifetime_in_minutes":15}`,
		{SessionLifetime: auth0.Float64(0.5)}:      `{"session_lifetime_in_minutes":30}`,
		{SessionLifetime: auth0.Float64(0.99)}:     `{"session_lifetime_in_minutes":59}`,
		{IdleSessionLifetime: auth0.Float64(0.25)}: `{"idle_session_lifetime_in_minutes":15}`,
		{AllowedLogoutURLs: nil}:                   `{}`,
		{AllowedLogoutURLs: &[]string{}}:           `{"allowed_logout_urls":[]}`,
	} {
		payload, err := json.Marshal(tenant)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func TestTenantUniversalLoginColors_MarshalJSON(t *testing.T) {
	for tenantUniversalLoginColors, expected := range map[*TenantUniversalLoginColors]string{
		{}: `{}`,
		{
			PageBackground: auth0.String("#ffffff"),
		}: `{"page_background":"#ffffff"}`,
		{
			PageBackgroundGradient: &BrandingPageBackgroundGradient{
				Type:        auth0.String("linear-gradient"),
				Start:       auth0.String("#ffffff"),
				End:         auth0.String("#000000"),
				AngleDegree: auth0.Int(3),
			},
		}: `{"page_background":{"type":"linear-gradient","start":"#ffffff","end":"#000000","angle_deg":3}}`,
	} {
		payload, err := json.Marshal(tenantUniversalLoginColors)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}

	t.Run("Should disallow setting PageBackground and PageBackgroundGradient", func(t *testing.T) {
		_, err := json.Marshal(&TenantUniversalLoginColors{
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

func TestTenantUniversalLoginColors_UnmarshalJSON(t *testing.T) {
	for jsonBody, expected := range map[string]*TenantUniversalLoginColors{
		`{}`: {},
		`{"page_background":"#ffffff"}`: {
			PageBackground: auth0.String("#ffffff"),
		},
		`{"page_background":{"type":"linear-gradient","start":"#ffffff","end":"#000000","angle_deg":3}}`: {
			PageBackgroundGradient: &BrandingPageBackgroundGradient{
				Type:        auth0.String("linear-gradient"),
				Start:       auth0.String("#ffffff"),
				End:         auth0.String("#000000"),
				AngleDegree: auth0.Int(3),
			},
		},
	} {
		var actual TenantUniversalLoginColors
		err := json.Unmarshal([]byte(jsonBody), &actual)
		assert.NoError(t, err)
		assert.Equal(t, &actual, expected)
	}

	t.Run("Should error is not expected type", func(t *testing.T) {
		var actual TenantUniversalLoginColors
		err := json.Unmarshal([]byte(`{"page_background":123}`), &actual)
		assert.Contains(t, err.Error(), "unexpected type for field page_background")
	})
}
