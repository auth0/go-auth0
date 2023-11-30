package authentication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0/authentication/database"
)

func TestDatabaseSignUp(t *testing.T) {
	configureHTTPTestRecordings(t, authAPI)

	userData := database.SignupRequest{
		Connection: "Username-Password-Authentication",
		Username:   "mytestaccount",
		Password:   "mypassword",
		Email:      "mytestaccount@example.com",
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.Username, createdUser.Username)
}

func TestDatabaseChangePassword(t *testing.T) {
	configureHTTPTestRecordings(t, authAPI)

	resp, err := authAPI.Database.ChangePassword(context.Background(), database.ChangePasswordRequest{
		Connection: "Username-Password-Authentication",
		Email:      "mytestaccount@example.com",
	})

	assert.NoError(t, err)
	assert.Equal(t, "We've just sent you an email to reset your password.", resp)
}
