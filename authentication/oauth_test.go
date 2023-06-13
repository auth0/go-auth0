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
