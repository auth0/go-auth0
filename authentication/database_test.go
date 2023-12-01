package authentication

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/authentication/database"
	"github.com/auth0/go-auth0/management"
)

func TestDatabaseSignUp(t *testing.T) {
	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t)

	userData := database.SignupRequest{
		Connection: details.connection,
		Username:   details.username,
		Password:   details.password,
		Email:      details.email,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)
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

type userDetails struct {
	username   string
	password   string
	email      string
	connection string
}

func givenSignUpDetails(t *testing.T) *userDetails {
	t.Helper()
	// If we're running from recordings then we want to return the default
	if usingRecordingResponses(t) {
		return &userDetails{
			username:   "mytestaccount",
			password:   "mypassword",
			email:      "mytestaccount@example.com",
			connection: "Username-Password-Authentication",
		}
	}

	conn := givenAConnection(t)

	return &userDetails{
		username:   fmt.Sprintf("chuck%d", rand.Intn(999)),
		password:   "Passwords hide their chuck",
		email:      fmt.Sprintf("chuck%d@example.com", rand.Intn(999)),
		connection: conn.GetName(),
	}
}

func givenAConnection(t *testing.T) management.Connection {
	conn := &management.Connection{
		Name:           auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
		Strategy:       auth0.String("auth0"),
		EnabledClients: &[]string{clientID, mgmtClientID},
		Options: &management.ConnectionOptions{
			RequiresUsername: auth0.Bool(true),
		},
	}

	err := mgmtAPI.Connection.Create(context.Background(), conn)
	require.NoError(t, err)

	t.Cleanup(func() {
		err := mgmtAPI.Connection.Delete(context.Background(), conn.GetID())
		require.NoError(t, err)
	})

	return *conn
}
