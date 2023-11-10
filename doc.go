// Package auth0 is the Auth0 SDK for Go.
//
// The SDK provides clients that interact with the Auth0 Authentication and Management APIs
// and is split into separate packages as such.
//
//   - auth0 - Provides helpers for providing values as pointers.
//   - authentication - Provides an Authentication Client for use when interacting with the
//     [Authentication API].
//   - management - Provides a Management Client for use when interacting with the Auth0
//     [Management API].
//
// # Getting Started
//
// Install the SDK using `go get`
//
//	go get github.com/auth0/go-auth0
//
// # Authentication
//
// Below is an example of using the Authentication client, for full documentation visit the
// [authentication client docs].
//
//	authAPI, err := authentication.New(
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
//	createdUser, err := authAPI.Database.Signup(context.TODO(), userData)
//	if err != nil {
//		// handle err
//	}
//
// # Management
//
// Below is an example of using the Management client, for full documentation visit the
// [management client docs].
//
//	import (
//		"github.com/auth0/go-auth0"
//		"github.com/auth0/go-auth0/management"
//	)
//	// Initialize a new client using a domain, client ID and secret.
//	m, err := management.New(
//		domain,
//		management.WithClientCredentials(context.TODO(), id, secret),
//	)
//	if err != nil {
//		// handle err
//	}
//	c := &management.Client{
//		Name:        auth0.String("Client Name"),
//		Description: auth0.String("Long description of client"),
//	}
//	err = m.Client.Create(context.TODO(), c)
//	if err != nil {
//		// handle err
//	}
//
// [management client docs]: https://pkg.go.dev/github.com/auth0/go-auth0/management
// [Authentication API]: https://auth0.com/docs/api/authentication
// [Management API]: https://auth0.com/docs/api/management/v2
// [authentication client docs]: https://pkg.go.dev/github.com/auth0/go-auth0/authentication
package auth0
