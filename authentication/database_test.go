package authentication

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/v3"

	"github.com/auth0/go-auth0/v3/authentication/database"
	"github.com/auth0/go-auth0/v3/management"
)

// TestDatabaseSignUp_RequiresUsername tests the Database.Signup method with a connection that requires a username.
func TestDatabaseSignUp_RequiresUsername(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		RequiresUsername: auth0.Bool(true),
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

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

// TestDatabaseSignUp_WithEmailIdentifier tests the Database.Signup method with a connection that requires an email.
func TestDatabaseSignUp_WithEmailIdentifier(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		Attributes: &management.ConnectionAttributes{
			Email: &management.EmailAttribute{
				Identifier: &management.ConnectionEmailAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupVerified{
					Status: management.SignupStatusEnumRequired.Ptr(),
					Verification: &management.SignupVerification{
						Active: auth0.Bool(false),
					},
				},
			},
		},
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

	userData := database.SignupRequest{
		Connection: details.connection,
		Password:   details.password,
		Email:      details.email,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)

	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.Email, createdUser.Email)
}

// TestDatabaseSignUp_WithUsernameIdentifier tests the Database.Signup method with a connection that requires a username.
func TestDatabaseSignUp_WithUsernameIdentifier(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		Attributes: &management.ConnectionAttributes{
			Username: &management.UsernameAttribute{
				Identifier: &management.ConnectionUsernameAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupSchema{
					Status: management.SignupStatusEnumRequired.Ptr(),
				},
			},
		},
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

	userData := database.SignupRequest{
		Connection: details.connection,
		Password:   details.password,
		Username:   details.username,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)

	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.Username, createdUser.Username)
}

// TestDatabaseSignUp_WithUsernameAndEmailIdentifiers tests the Database.Signup method with a connection that requires both username and email.
func TestDatabaseSignUp_WithUsernameAndEmailIdentifiers(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		Attributes: &management.ConnectionAttributes{
			Username: &management.UsernameAttribute{
				Identifier: &management.ConnectionUsernameAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupSchema{
					Status: management.SignupStatusEnumRequired.Ptr(),
				},
			},
			Email: &management.EmailAttribute{
				Identifier: &management.ConnectionEmailAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupVerified{
					Status: management.SignupStatusEnumRequired.Ptr(),
					Verification: &management.SignupVerification{
						Active: auth0.Bool(false),
					},
				},
			},
		},
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

	userData := database.SignupRequest{
		Connection: details.connection,
		Password:   details.password,
		Username:   details.username,
		Email:      details.email,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)

	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.Username, createdUser.Username)
	assert.Equal(t, userData.Email, createdUser.Email)
}

// TestDatabaseSignUp_WithPhoneNumberIdentifier tests the Database.Signup method with a connection that requires a phone number.
func TestDatabaseSignUp_WithPhoneNumberIdentifier(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		Attributes: &management.ConnectionAttributes{
			PhoneNumber: &management.PhoneAttribute{
				Identifier: &management.ConnectionPhoneAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupVerified{
					Status: management.SignupStatusEnumRequired.Ptr(),
					Verification: &management.SignupVerification{
						Active: auth0.Bool(false),
					},
				},
			},
		},
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

	userData := database.SignupRequest{
		Connection:  details.connection,
		Password:    details.password,
		PhoneNumber: details.phoneNumber,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)

	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.PhoneNumber, createdUser.PhoneNumber)
}

// TestDatabaseSignUp_WithUsernameAndPhoneNumberIdentifiers tests the Database.Signup method with a connection that requires both username and phone number.
func TestDatabaseSignUp_WithUsernameAndPhoneNumberIdentifiers(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		Attributes: &management.ConnectionAttributes{
			Username: &management.UsernameAttribute{
				Identifier: &management.ConnectionUsernameAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupSchema{
					Status: management.SignupStatusEnumRequired.Ptr(),
				},
			},
			PhoneNumber: &management.PhoneAttribute{
				Identifier: &management.ConnectionPhoneAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupVerified{
					Status: management.SignupStatusEnumRequired.Ptr(),
					Verification: &management.SignupVerification{
						Active: auth0.Bool(false),
					},
				},
			},
		},
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

	userData := database.SignupRequest{
		Connection:  details.connection,
		Password:    details.password,
		Username:    details.username,
		PhoneNumber: details.phoneNumber,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)

	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.Username, createdUser.Username)
	assert.Equal(t, userData.PhoneNumber, createdUser.PhoneNumber)
}

// TestDatabaseSignUp_WithUsernameEmailAndPhoneNumberIdentifiers tests the Database.Signup method with a connection that requires username, email, and phone number.
func TestDatabaseSignUp_WithUsernameEmailAndPhoneNumberIdentifiers(t *testing.T) {
	connectionOptions := &management.ConnectionOptionsAuth0{
		Attributes: &management.ConnectionAttributes{
			Username: &management.UsernameAttribute{
				Identifier: &management.ConnectionUsernameAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupSchema{
					Status: management.SignupStatusEnumRequired.Ptr(),
				},
			},
			Email: &management.EmailAttribute{
				Identifier: &management.ConnectionEmailAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupVerified{
					Status: management.SignupStatusEnumRequired.Ptr(),
					Verification: &management.SignupVerification{
						Active: auth0.Bool(false),
					},
				},
			},
			PhoneNumber: &management.PhoneAttribute{
				Identifier: &management.ConnectionPhoneAttributeIdentifier{
					Active: auth0.Bool(true),
				},
				ProfileRequired: auth0.Bool(true),
				Signup: &management.SignupVerified{
					Status: management.SignupStatusEnumRequired.Ptr(),
					Verification: &management.SignupVerification{
						Active: auth0.Bool(false),
					},
				},
			},
		},
	}

	configureHTTPTestRecordings(t, authAPI)

	details := givenSignUpDetails(t, connectionOptions)

	userData := database.SignupRequest{
		Connection:  details.connection,
		Password:    details.password,
		Username:    details.username,
		Email:       details.email,
		PhoneNumber: details.phoneNumber,
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	require.NoError(t, err)

	assert.NotEmpty(t, createdUser.ID)
	assert.Equal(t, userData.Username, createdUser.Username)
	assert.Equal(t, userData.Email, createdUser.Email)
	assert.Equal(t, userData.PhoneNumber, createdUser.PhoneNumber)
}

// TestDatabaseChangePassword tests the Database.ChangePassword method.
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
	username    string
	password    string
	email       string
	connection  string
	phoneNumber string
}

func givenSignUpDetails(t *testing.T, options *management.ConnectionOptionsAuth0) *userDetails {
	t.Helper()
	// If we're running from recordings then we want to return the default
	if usingRecordingResponses(t) {
		return &userDetails{
			username:    "mytestaccount",
			password:    "mypassword",
			email:       "mytestaccount@example.com",
			phoneNumber: "+12345678900",
			connection:  "Username-Password-Authentication",
		}
	}

	conn := givenAConnection(t, options)

	return &userDetails{
		username:    fmt.Sprintf("chuck%d", rand.Intn(999)),
		password:    "Passwords hide their chuck",
		email:       fmt.Sprintf("chuck%d@example.com", rand.Intn(999)),
		phoneNumber: fmt.Sprintf("+1234567890%d", rand.Intn(9)),
		connection:  conn.GetName(),
	}
}

func givenAConnection(t *testing.T, options *management.ConnectionOptionsAuth0) *management.ConnectionResponseContentAuth0 {
	enabledClients := management.ConnectionEnabledClients{clientID, mgmtClientID}
	conn := &management.CreateConnectionRequestContent{
		CreateConnectionRequestContentAuth0: &management.CreateConnectionRequestContentAuth0{
			Name:           auth0.String(fmt.Sprintf("Test-Connection-%d", time.Now().Unix())),
			EnabledClients: &enabledClients,
			Options:        options,
		},
	}

	connectionCreated, err := mgmtAPI.Connections.Create(context.Background(), conn)
	require.NoError(t, err)

	auth0Conn := connectionCreated.GetConnectionResponseContentAuth0()

	t.Cleanup(func() {
		err := mgmtAPI.Connections.Delete(context.Background(), auth0Conn.GetID())
		require.NoError(t, err)
	})

	return auth0Conn
}
