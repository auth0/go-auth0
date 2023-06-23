package authentication

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/auth0/go-auth0/authentication/passwordless"
)

func TestSendEmail(t *testing.T) {
	configureHTTPTestRecordings(t)

	r, err := authAPI.Passwordless.SendEmail(context.Background(), passwordless.SendEmailRequest{
		Email: "test-email@example.com",
		Send:  "code",
	})

	assert.NoError(t, err)
	assert.Equal(t, "test-email@example.com", r.Email)
	assert.Equal(t, true, r.EmailVerified)
}

func TestLoginWithEmail(t *testing.T) {
	configureHTTPTestRecordings(t)

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
	configureHTTPTestRecordings(t)

	r, err := authAPI.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
		PhoneNumber: "+123456789",
	})

	assert.NoError(t, err)
	assert.Equal(t, "+123456789", r.PhoneNumber)
	assert.Equal(t, true, r.PhoneVerified)
}

func TestLoginWithSMS(t *testing.T) {
	configureHTTPTestRecordings(t)

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
	t.Run("error for an invalid organization", func(t *testing.T) {
		extras := map[string]interface{}{
			"org_id": "wrong-org",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{OrganizationID: "right-org"})

		assert.ErrorContains(t, err, "org_id claim value mismatch in the ID token")
	})

	t.Run("error for an invalid nonce", func(t *testing.T) {
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
		extras := map[string]interface{}{
			"auth_time": time.Now().Add(-500 * time.Second).Unix(),
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{MaxAge: 100 * time.Second})

		assert.ErrorContains(t, err, "auth_time claim in the ID token indicates that too much time has passed")
	})
}
