// Package authentication provides a client for using the Auth0 Authentication API.
//
// # Usage
//
//	import (
//		"github.com/auth0/go-auth0/v2"
//		"github.com/auth0/go-auth0/v2/authentication"
//		"github.com/auth0/go-auth0/v2/authentication/database"
//		"github.com/auth0/go-auth0/v2/authentication/oauth"
//	)
//	a, err := authentication.New(
//		context.TODO(),
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
//	createdUser, err := a.Database.Signup(context.TODO(), userData)
//	if err != nil {
//		// handle err
//	}
//	// Login using OAuth grants
//	tokenSet, err := a.OAuth.LoginWithAuthCodeWithPKCE(context.TODO(), oauth.LoginWithAuthCodeWithPKCERequest{
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
//		context.TODO(),
//		domain,
//		authentication.WithClientID(id),
//		authentication.WithClientSecret(secret), // Optional depending on the grants used
//		authentication.WithClockTolerance(10 * time.Second),
//	)
//
// # Handling Errors
//
// This package exports an [authentication.Error] type that can be used to check errors
// returned from the Authentication API and handle them as necessary, for example
//
//	tokens, err := auth.OAuth.LoginWithPassword(context.Background(), oauth.LoginWithPasswordRequest{
//		Username: "test@example.com",
//		Password: "hunter2",
//	}, oauth.IDTokenValidationOptions{})
//
//	if err != nil {
//		if aerr, ok := err.(*authentication.Error); ok {
//			if aerr.Err == "mfa_required" {
//				// Handle prompting for MFA usage
//			}
//		}
//	}
package authentication
