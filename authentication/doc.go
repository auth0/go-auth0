// Package authentication provides a client for using the Auth0 Authentication API.
//
// # Usage
//
//	import (
//		"github.com/auth0/go-auth0"
//		"github.com/auth0/go-auth0/authentication"
//		"github.com/auth0/go-auth0/authentication/database"
//		"github.com/auth0/go-auth0/authentication/oauth"
//	)
//	a, err := authentication.New(
//		context.Background(),
//		domain,
//		authentication.WithClientID(id),
//		authentication.WithClientSecret(secret), // Optional depending on the grants used
//	)
//	if err != nil {
//		// handle err
//	}
//	// Now we have an authentication client, we can interact with the Auth0 Authentication API.
//	// Sign up a user
//	userData := database.SignupRequest{
//		Connection: "Username-Password-Authentication",
//		Username:   "mytestaccount",
//		Password:   "mypassword",
//		Email:      "mytestaccount@example.com",
//	}
//	createdUser, err := a.Database.Signup(context.Background(), userData)
//	if err != nil {
//		// handle err
//	}
//	// Login using OAuth grants
//	tokenSet, err := a.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
//		Code:         "test-code",
//		CodeVerifier: "test-code-verifier",
//	}, oauth.IDTokenValidationOptionalVerification{})
//
//	if err != nil {
//		// handle err
//	}
//
// # Configuration
//
// There are several options that can be specified during the creation of a client.
// For a complete list see [Option].
//
//	a, err := authentication.New(
//		context.Background(),
//		domain,
//		authentication.WithClientID(id),
//		authentication.WithClientSecret(secret), // Optional depending on the grants used
//		authentication.WithClockTolerance(10 * time.Second),
//	)
package authentication
