package authentication

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/v3/authentication/ciba"
)

func TestCIBA_Initiate(t *testing.T) {
	configureHTTPTestRecordings(t, authAPI)

	// Call the Initiate method of the CIBA manager
	resp, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		Scope:          "openid",
		LoginHint:      cibaLoginHint(t, ""),
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
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			Scope:          "openid",
			LoginHint:      cibaLoginHint(t, "auth0|Random-ID"),
			BindingMessage: "TEST-BINDING-MESSAGE",
		})

		assert.ErrorContains(t, err, "User ID is malformed or unknown")
	})

	t.Run("Should throw error if scope is not openid", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			Scope:          "tempID",
			LoginHint:      cibaLoginHint(t, ""),
			BindingMessage: "TEST-BINDING-MESSAGE",
		})

		assert.ErrorContains(t, err, "openid scope must be requested")
	})
}

// cibaLoginHint returns the login hint used by the CIBA tests. When replaying recordings it uses
// placeholders, otherwise the issuer and user ID come from AUTH0_CIBA_ISSUER and AUTH0_CIBA_USER_ID.
// A non-empty userID overrides the user ID.
func cibaLoginHint(t *testing.T, userID string) map[string]string {
	t.Helper()

	issuer, sub := recordingsCIBAIssuer, recordingsCIBAUserID

	if !usingRecordingResponses(t) {
		if v := os.Getenv("AUTH0_CIBA_ISSUER"); v != "" {
			issuer = v
		}

		if v := os.Getenv("AUTH0_CIBA_USER_ID"); v != "" {
			sub = v
		}
	}

	if userID != "" {
		sub = userID
	}

	return map[string]string{
		"format": "iss_sub",
		"iss":    issuer,
		"sub":    sub,
	}
}
