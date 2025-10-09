package authentication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/v2/authentication/ciba"
)

func TestCIBA_Initiate(t *testing.T) {
	configureHTTPTestRecordings(t, authAPI)

	// Call the Initiate method of the CIBA manager
	resp, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        "openid",
		LoginHint: map[string]string{
			"format": "iss_sub",
			"iss":    "https://witty-silver-sailfish-sus1-staging-20240704.sus.auth0.com/",
			"sub":    "auth0|6707939cad3d8bec47ecfa2e",
		},
		BindingMessage: "TEST-BINDING-MESSAGE",
	})

	// Validate the response
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.AuthReqID, "auth_req_id should not be empty")
	assert.Greater(t, resp.ExpiresIn, int64(0), "expires_in should be greater than 0")
	assert.Greater(t, resp.Interval, int64(0), "interval should be greater than 0")
}

func TestCIBANegative_Initiate(t *testing.T) {
	t.Run("Should throw error for missing LoginHint and BindingMessage", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scope:        "openid",
		})

		assert.ErrorContains(t, err, "missing required fields: LoginHint, BindingMessage")
	})

	t.Run("Should throw error for invalid User ID", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scope:        "openid",
			LoginHint: map[string]string{
				"format": "iss_sub",
				"iss":    "https://witty-silver-sailfish-sus1-staging-20240704.sus.auth0.com/",
				"sub":    "auth0|Random-ID",
			},
			BindingMessage: "TEST-BINDING-MESSAGE",
		})

		assert.ErrorContains(t, err, "User ID is malformed or unknown")
	})

	t.Run("Should throw error if scope is not openid", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scope:        "tempID",
			LoginHint: map[string]string{
				"format": "iss_sub",
				"iss":    "https://witty-silver-sailfish-sus1-staging-20240704.sus.auth0.com/",
				"sub":    "auth0|6707939cad3d8bec47ecfa2e",
			},
			BindingMessage: "TEST-BINDING-MESSAGE",
		})

		assert.ErrorContains(t, err, "openid scope must be requested")
	})
}
