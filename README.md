![Go SDK for Auth0](https://cdn.auth0.com/website/sdks/banners/go-auth0-banner.png)

<div align="center">

[![GoDoc](https://pkg.go.dev/badge/github.com/auth0/go-auth0/v2.svg)](https://pkg.go.dev/github.com/auth0/go-auth0/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/auth0/go-auth0/v2?style=flat-square)](https://goreportcard.com/report/github.com/auth0/go-auth0/v2)
[![Release](https://img.shields.io/github/v/release/auth0/go-auth0?include_prereleases&style=flat-square)](https://github.com/auth0/go-auth0/releases)
[![License](https://img.shields.io/github/license/auth0/go-auth0.svg?style=flat-square)](https://github.com/auth0/go-auth0/blob/v2/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/auth0/go-auth0/main.yml?branch=v2&style=flat-square)](https://github.com/auth0/go-auth0/actions?query=branch%3Av2)
[![Codecov](https://img.shields.io/codecov/c/github/auth0/go-auth0/v2?style=flat-square)](https://codecov.io/gh/auth0/go-auth0/tree/v2)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fauth0%2Fgo-auth0.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fauth0%2Fgo-auth0?ref=badge_shield)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/auth0/go-auth0)

📚 [Documentation](#documentation) • 🚀 [Getting started](#getting-started) • 💬 [Feedback](#feedback)

</div>

---

## Documentation

- [Godoc](https://pkg.go.dev/github.com/auth0/go-auth0/v2) - explore the Go SDK documentation.
- [Docs site](https://www.auth0.com/docs) — explore our docs site and learn more about Auth0.
- [Examples](./EXAMPLES.md) - Further examples around usage of the SDK.
- [API Reference](./reference.md) - Complete API reference documentation.

## Getting started

### Requirements

This library follows the [same support policy as Go](https://go.dev/doc/devel/release#policy). The last two major Go releases are actively supported and compatibility issues will be fixed. While you may find that older versions of Go may work, we will not actively test and fix compatibility issues with these versions.

- Go 1.24+

### Installation

```shell
go get github.com/auth0/go-auth0/v2
```

### Usage

#### Authentication API Client

This client can be used to access Auth0's [Authentication API](https://auth0.com/docs/api/authentication).

```go
package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0/v2/authentication"
	"github.com/auth0/go-auth0/v2/authentication/database"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
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

The Auth0 Management API is meant to be used by back-end servers or trusted parties performing administrative tasks. Generally speaking, anything that can be done through the Auth0 dashboard (and more) can also be done through this API.

Initialize your client class with a domain and token:

```go
package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0/v2/management/option"
	management "github.com/auth0/go-auth0/v2/management/client"
)

func main() {
	mgmt, err := management.New(
		"{YOUR_TENANT_AND REGION}.auth0.com",
		option.WithToken("{YOUR_API_V2_TOKEN}"),  // Replace with a Context that better suits your usage
	)
}
```

Or use client credentials:

```go
package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0/v2/management/option"
	management "github.com/auth0/go-auth0/v2/management/client"
)

func main() {
	mgmt, err := management.New(
		"{YOUR_TENANT_AND REGION}.auth0.com",
		option.WithClientCredentials(context.TODO(), clientID, clientSecret),  // Replace with a Context that better suits your usage
	)
}
```

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