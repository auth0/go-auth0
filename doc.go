/*
Package auth0 provides a client for using the Auth0 Authentication and Management APIs.

# Authentication

# Usage

	import (
		github.com/auth0/go-auth0
		github.com/auth0/go-auth0/authentication
		github.com/auth0/go-auth0/authentication/database
		github.com/auth0/go-auth0/authentication/oauth
	)

Initialize a new client using a context, domain, client ID, and client secret if required.

	authAPI, err := authentication.New(
		context.Background(),
		domain,
		authentication.WithClientID(id),
		authentication.WithClientSecret(secret) // Optional depending on the grants used
	)
	if err != nil {
		// handle err
	}

Now we have an authentication client, we can interact with the Auth0 Authentication API.

	// Sign up a user
	userData := database.SignupRequest{
		Connection: "Username-Password-Authentication",
		Username:   "mytestaccount",
		Password:   "mypassword",
		Email:      "mytestaccount@example.com",
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	if err != nil {
		// handle err
	}

	// Login using OAuth grants
	tokenSet, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
		Code:         "test-code",
		CodeVerifier: "test-code-verifier",
	}, oauth.IDTokenValidationOptionalVerification{})
	if err != nil {
		// handle err
	}

# Management

Usage

	import (
		github.com/auth0/go-auth0
		github.com/auth0/go-auth0/management
	)

Initialize a new client using a domain, client ID and secret.

	m, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), id, secret)
	)
	if err != nil {
		// handle err
	}

Or using a static token.

	m, err := management.New(domain, management.WithStaticToken(token))
	if err != nil {
		// handle err
	}

With a management client we can then interact with the Auth0 Management API.

	c := &management.Client{
		Name:        auth0.String("Client Name"),
		Description: auth0.String("Long description of client"),
	}

	err = m.Client.Create(context.Background(), c)
	if err != nil {
		// handle err
	}

## Authentication

The auth0 management package handles authentication by exchanging the client ID and secret
supplied when creating a new management client.

This is handled internally using the https://godoc.org/golang.org/x/oauth2
package.

## Rate Limiting

The auth0 package also handles rate limiting by respecting the `X-Rate-Limit-*`
headers sent by the server.

The amount of time the client waits for the rate limit to be reset is taken from
the `X-Rate-Limit-Reset` header as the amount of seconds to wait.

# Configuration

There are several other options that can be specified during the creation of a
new client.

	m, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), id, secret),
		management.WithDebug(true)
	)

## Request Options

As with the global client configuration, fine-grained configuration can be done
on a request basis.

	c, err := m.Connection.List(
		context.Background(),
		management.Page(2),
		management.PerPage(10),
		management.IncludeFields("id", "name", "options")
		management.Parameter("strategy", "auth0"),
	)
*/
package auth0
