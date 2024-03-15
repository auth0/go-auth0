![Go SDK for Auth0](https://cdn.auth0.com/website/sdks/banners/go-auth0-banner.png)

<div align="center">

[![GoDoc](https://pkg.go.dev/badge/github.com/auth0/go-auth0.svg)](https://pkg.go.dev/github.com/auth0/go-auth0)
[![Go Report Card](https://goreportcard.com/badge/github.com/auth0/go-auth0?style=flat-square)](https://goreportcard.com/report/github.com/auth0/go-auth0)
[![Release](https://img.shields.io/github/v/release/auth0/go-auth0?include_prereleases&style=flat-square)](https://github.com/auth0/go-auth0/releases)
[![License](https://img.shields.io/github/license/auth0/go-auth0.svg?style=flat-square)](https://github.com/auth0/go-auth0/blob/main/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/auth0/go-auth0/main.yml?branch=main&style=flat-square)](https://github.com/auth0/go-auth0/actions?query=branch%3Amain)
[![Codecov](https://img.shields.io/codecov/c/github/auth0/go-auth0?style=flat-square)](https://codecov.io/gh/auth0/go-auth0)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fauth0%2Fgo-auth0.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fauth0%2Fgo-auth0?ref=badge_shield)

📚 [Documentation](#documentation) • 🚀 [Getting Started](#getting-started) • 💬 [Feedback](#feedback)

</div>



-------------------------------------

## Documentation

- [Godoc](https://pkg.go.dev/github.com/auth0/go-auth0) - explore the Go SDK documentation.
- [Docs site](https://www.auth0.com/docs) — explore our docs site and learn more about Auth0.
- [Examples](./EXAMPLES.md) - Further examples around usage of the SDK.

## Getting started

### Requirements

This library follows the [same support policy as Go](https://go.dev/doc/devel/release#policy). The last two major Go releases are actively supported and compatibility issues will be fixed. While you may find that older versions of Go may work, we will not actively test and fix compatibility issues with these versions.

- Go 1.21+

### Installation

```shell
go get github.com/auth0/go-auth0
```

### Usage

#### Authentication API Client

The Authentication API client is based on the [Authentication API docs](https://auth0.com/docs/api/authentication).

Create an Authentication API client by providing the details of your Auth0 Application.

```go
package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0/authentication"
	"github.com/auth0/go-auth0/authentication/database"
	"github.com/auth0/go-auth0/authentication/oauth"
)

func main() {
	// Get these from your Auth0 Application Dashboard.
	domain := "example.us.auth0.com"
	clientID := "EXAMPLE_16L9d34h0qe4NVE6SaHxZEid"
	clientSecret := "EXAMPLE_XSQGmnt8JdXs23407hrK6XXXXXXX"

	// Initialize a new client using a domain, client ID and client secret.
	authAPI, err := authentication.New(
		context.TODO(), // Replace with a Context that better suits your usage
		domain,
		authentication.WithClientID(clientID),
		authentication.WithClientSecret(clientSecret), // Optional depending on the grants used
	)
	if err != nil {
		log.Fatalf("failed to initialize the auth0 authentication API client: %+v", err)
	}

	// Now we can interact with the Auth0 Authentication API.
	// Sign up a user
	userData := database.SignupRequest{
		Connection: "Username-Password-Authentication",
		Username:   "mytestaccount",
		Password:   "mypassword",
		Email:      "mytestaccount@example.com",
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	if err != nil {
		log.Fatalf("failed to sign user up: %+v", err)
	}

	// Login using OAuth grants
	tokenSet, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
		Code:         "test-code",
		CodeVerifier: "test-code-verifier",
	}, oauth.IDTokenValidationOptions{})
	if err != nil {
		log.Fatalf("failed to retrieve tokens: %+v", err)
	}
}
```
> **Note**
> The [context](https://pkg.go.dev/context?utm_source=godoc) package can be used to pass cancellation signals and deadlines to the Client for handling a request. If there is no context available then `context.Background()` can be used.

#### Management API Client

The Management API client is based on the [Management API docs](https://auth0.com/docs/api/management/v2).

```go
package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/management"
)

func main() {
	// Get these from your Auth0 Application Dashboard.
	// The application needs to be a Machine To Machine authorized
	// to request access tokens for the Auth0 Management API,
	// with the desired permissions (scopes).
	domain := "example.auth0.com"
	clientID := "EXAMPLE_16L9d34h0qe4NVE6SaHxZEid"
	clientSecret := "EXAMPLE_XSQGmnt8JdXs23407hrK6XXXXXXX"

	// Initialize a new client using a domain, client ID and client secret.
	// Alternatively you can specify an access token:
	// `management.WithStaticToken("token")`
	auth0API, err := management.New(
		domain,
		management.WithClientCredentials(context.TODO(), clientID, clientSecret),  // Replace with a Context that better suits your usage
	)
	if err != nil {
		log.Fatalf("failed to initialize the auth0 management API client: %+v", err)
	}

	// Now we can interact with the Auth0 Management API.
	// Example: Creating a new client.
	client := &management.Client{
		Name:        auth0.String("My Client"),
		Description: auth0.String("Client created through the Go SDK"),
	}

	// The passed in client will get hydrated with the response.
	// This means that after this request, we will have access
	// to the client ID on the same client object.
	err = auth0API.Client.Create(context.TODO(), client)  // Replace with a Context that better suits your usage
	if err != nil {
		log.Fatalf("failed to create a new client: %+v", err)
	}

	// Make use of the getter functions to safely access
	// fields without causing a panic due nil pointers.
	log.Printf(
		"Created an auth0 client successfully. The ID is: %q",
		client.GetClientID(),
	)
}
```
> **Note**
> The [context](https://pkg.go.dev/context?utm_source=godoc) package can be used to pass cancellation signals and deadlines to the Client for handling a request. If there is no context available then `context.Background()` can be used.

### Rate Limiting

The Auth0 Management API imposes a rate limit on all API clients. When the limit is reached, the SDK will handle it in
the background by retrying the API request when the limit is lifted.

> **Note**
> The SDK does not prevent `http.StatusTooManyRequests` errors, instead it waits for the rate limit to be reset based on
> the value of the `X-Rate-Limit-Reset` header as the amount of seconds to wait.

## Feedback

### Contributing

We appreciate feedback and contribution to this repo! Before you get started, please see the following:

- [Contribution Guide](./CONTRIBUTING.md)
- [Auth0's General Contribution Guidelines](https://github.com/auth0/open-source-template/blob/master/GENERAL-CONTRIBUTING.md)
- [Auth0's Code of Conduct Guidelines](https://github.com/auth0/open-source-template/blob/master/CODE-OF-CONDUCT.md)

### Raise an issue

To provide feedback or report a bug, [please raise an issue on our issue tracker](https://github.com/auth0/go-auth0/issues).

### Vulnerability Reporting

Please do not report security vulnerabilities on the public Github issue tracker. The [Responsible Disclosure Program](https://auth0.com/responsible-disclosure-policy) details the procedure for disclosing security issues.

---

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: light)" srcset="https://cdn.auth0.com/website/sdks/logos/auth0_light_mode.png" width="150">
    <source media="(prefers-color-scheme: dark)" srcset="https://cdn.auth0.com/website/sdks/logos/auth0_dark_mode.png" width="150">
    <img alt="Auth0 Logo" src="https://cdn.auth0.com/website/sdks/logos/auth0_light_mode.png" width="150">
  </picture>
</p>

<p align="center">Auth0 is an easy to implement, adaptable authentication and authorization platform.<br />To learn more checkout <a href="https://auth0.com/why-auth0">Why Auth0?</a></p>

<p align="center">This project is licensed under the MIT license. See the <a href="./LICENSE"> LICENSE</a> file for more info.</p>
