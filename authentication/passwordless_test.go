package authentication

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/internal/client"

	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/auth0/go-auth0/authentication/passwordless"
)

func TestSendEmail(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	r, err := authAPI.Passwordless.SendEmail(context.Background(), passwordless.SendEmailRequest{
		Email: "test-email@example.com",
		Send:  "code",
	})

	assert.NoError(t, err)
	assert.Equal(t, "test-email@example.com", r.Email)
	assert.Equal(t, true, r.EmailVerified)
}

func TestLoginWithEmail(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	token, err := authAPI.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
		Code:     "123456",
		Email:    "test-email@example.com",
		Scope:    "openid profile email offline_access",
		Audience: "https://api.example.com",
	}, oauth.IDTokenValidationOptions{})

	assert.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func TestSendSMS(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	r, err := authAPI.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
		PhoneNumber: "+123456789",
	})

	assert.NoError(t, err)
	assert.Equal(t, "+123456789", r.PhoneNumber)
	assert.Equal(t, true, r.PhoneVerified)
}

func TestLoginWithSMS(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	token, err := authAPI.Passwordless.LoginWithSMS(context.Background(), passwordless.LoginWithSMSRequest{
		PhoneNumber: "+123456789",
		Code:        "123456",
		Scope:       "openid profile email offline_access",
		Audience:    "https://api.example.com",
	}, oauth.IDTokenValidationOptions{})

	assert.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func TestPasswordlessWithIDTokenVerification(t *testing.T) {
	t.Run("error for an invalid organization when using org_id", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"org_id": "org_124",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{Organization: "org_456"})

		assert.ErrorContains(t, err, "org_id claim value mismatch in the ID token")
	})

	t.Run("error for an invalid organization when using org_name", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"org_name": "wrong-org",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{Organization: "right-org"})

		assert.ErrorContains(t, err, "org_name claim value mismatch in the ID token")
	})

	t.Run("error for an invalid nonce", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"nonce": "wrong-nonce",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{Nonce: "test-nonce"})

		assert.ErrorContains(t, err, "nonce claim value mismatch in the ID token; expected")
	})

	t.Run("error for an invalid maxage", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"auth_time": time.Now().Add(-500 * time.Second).Unix(),
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithSMS(context.Background(), passwordless.LoginWithSMSRequest{
			PhoneNumber: "+123456789",
			Code:        "123456",
			Scope:       "openid profile email offline_access",
			Audience:    "https://api.example.com",
		}, oauth.IDTokenValidationOptions{MaxAge: 100 * time.Second})

		assert.ErrorContains(t, err, "auth_time claim in the ID token indicates that too much time has passed")
	})
}

func TestPasswordlessWithClientAssertion(t *testing.T) {
	t.Run("Should support using private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
			WithClientAssertion(jwtPrivateKey, "RS256"),
		)

		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		r, err := api.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
			PhoneNumber: "+123456789",
		})

		assert.NoError(t, err)
		assert.Equal(t, "+123456789", r.PhoneNumber)
		assert.True(t, r.PhoneVerified)
	})

	t.Run("Should support passing private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
		)
		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		auth, err := client.CreateClientAssertion("RS256", jwtPrivateKey, clientID, "https://"+domain+"/")
		require.NoError(t, err)

		r, err := api.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientAssertion:     auth,
				ClientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			},
			PhoneNumber: "+123456789",
		})

		assert.NoError(t, err)
		assert.Equal(t, "+123456789", r.PhoneNumber)
		assert.True(t, r.PhoneVerified)
	})
}
