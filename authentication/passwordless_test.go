package authentication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

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
		Code:  "123456",
		Email: "test-email@example.com",
		Scope: "openid profile email offline_access",
	})

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
	assert.Equal(t, "en", r.RequestLanguage)
}

func TestLoginWithSMS(t *testing.T) {
	configureHTTPTestRecordings(t)

	token, err := authAPI.Passwordless.LoginWithSMS(context.Background(), passwordless.LoginWithSMSRequest{
		PhoneNumber: "+123456789",
		Code:        "123456",
		Scope:       "openid profile email offline_access",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}
