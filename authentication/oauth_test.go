package authentication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0/authentication/oauth"
)

func TestOAuthLoginWithPassword(t *testing.T) {
	configureHTTPTestRecordings(t)

	tokenSet, err := authAPI.OAuth.LoginWithPassword(context.Background(), oauth.LoginWithPasswordRequest{
		Username: "testuser",
		Password: "testuser123",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenSet.AccessToken)
	assert.NotEmpty(t, tokenSet.IDToken)
	assert.Equal(t, "Bearer", tokenSet.TokenType)
}

func TestLoginWithAuthCode(t *testing.T) {
	t.Run("Should require client_secret", func(t *testing.T) {
		_, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			Code: "my-code",
		})

		assert.Error(t, err, "client_secret is required but not provided")
	})

	t.Run("Should throw for an invalid code", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		_, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Code: "my-invalid-code",
		})

		assert.Error(t, err, "Invalid authorization code")
	})

	t.Run("Should return tokens", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		tokenSet, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Code: "my-code",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.NotEmpty(t, tokenSet.IDToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestLoginWithAuthCodeWithPKCE(t *testing.T) {
	t.Run("Should throw for an invalid code", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		_, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "my-invalid-code",
			CodeVerifier: "test-code-verifier",
		})

		assert.Error(t, err, "Invalid authorization code")
	})

	t.Run("Should throw for an invalid code verifier", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		_, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "test-code",
			CodeVerifier: "test-invalid-code-verifier",
		})

		assert.Error(t, err, "Failed to verify code verifier")
	})

	t.Run("Should return tokens", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		tokenSet, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "test-code",
			CodeVerifier: "test-code-verifier",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.NotEmpty(t, tokenSet.IDToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestLoginWithClientCredentials(t *testing.T) {
	t.Run("Should require client_secret", func(t *testing.T) {
		_, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			Audience: "test-audience",
		})

		assert.Error(t, err, "client_secret is required but not provided")
	})

	t.Run("Should return tokens", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		tokenSet, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Audience: "test-audience",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("Should return tokens", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		tokenSet, err := authAPI.OAuth.RefreshToken(context.Background(), oauth.RefreshTokenRequest{
			RefreshToken: "test-refresh-token",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.NotEmpty(t, tokenSet.IDToken)
		assert.NotEmpty(t, tokenSet.RefreshToken)
		assert.NotEmpty(t, tokenSet.Scope)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should return tokens with reduced scopes", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		tokenSet, err := authAPI.OAuth.RefreshToken(context.Background(), oauth.RefreshTokenRequest{
			RefreshToken: "test-refresh-token",
			Scope:        "openid profile offline_access",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.NotEmpty(t, tokenSet.IDToken)
		assert.NotEmpty(t, tokenSet.RefreshToken)
		assert.Equal(t, "openid profile offline_access", tokenSet.Scope)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestRevokeRefreshToken(t *testing.T) {
	t.Run("Should revoke token", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		err := authAPI.OAuth.RevokeRefreshToken(context.Background(), oauth.RevokeRefreshTokenRequest{
			Token: "test-refresh-token",
		})

		assert.NoError(t, err)
	})
}
