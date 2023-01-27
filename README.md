<div align="center">
  <h1>Auth0 Go SDK</h1>

[![GoDoc](https://pkg.go.dev/badge/github.com/auth0/go-auth0.svg)](https://pkg.go.dev/github.com/auth0/go-auth0)
[![Go Report Card](https://goreportcard.com/badge/github.com/auth0/go-auth0?style=flat-square)](https://goreportcard.com/report/github.com/auth0/go-auth0)
[![Release](https://img.shields.io/github/v/release/auth0/go-auth0?include_prereleases&style=flat-square)](https://github.com/auth0/go-auth0/releases)
[![License](https://img.shields.io/github/license/auth0/go-auth0.svg?style=flat-square)](https://github.com/auth0/go-auth0/blob/main/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/auth0/go-auth0/main.yml?branch=main&style=flat-square)](https://github.com/auth0/go-auth0/actions?query=branch%3Amain)
[![Codecov](https://img.shields.io/codecov/c/github/auth0/go-auth0?style=flat-square)](https://codecov.io/gh/auth0/go-auth0)

</div>

---

Go SDK for the [Auth0](https://auth0.com/) Management API.

📚 [Documentation](#documentation) • 🚀 [Getting Started](#getting-started) • 💬 [Feedback](#feedback)

-------------------------------------

## Documentation

- [Godoc](https://pkg.go.dev/github.com/auth0/go-auth0) - explore the Go SDK documentation.
- [Management API docs](https://auth0.com/docs/api/management/v2) - explore the Auth0 Management API that this SDK interacts with.
- [Docs site](https://www.auth0.com/docs) — explore our docs site and learn more about Auth0.

## Getting started

### Requirements

- Go 1.17+

**go-auth0** tracks [Go's version support policy](https://go.dev/doc/devel/release#policy). 

### Installation

```shell
go get github.com/auth0/go-auth0
```

### Usage

```go
package main

import (
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
		management.WithClientCredentials(clientID, clientSecret),
	)
	if err != nil {
		log.Fatal("failed to initialize the auth0 management API client: %w", err)
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
	err = auth0API.Client.Create(client)
	if err != nil {
		log.Fatal("failed to create a new client: %w", err)
	}

	// Make use of the getter functions to safely access
	// fields without causing a panic due nil pointers.
	log.Printf(
		"Created an auth0 client successfully. The ID is: %q",
		client.GetClientID(),
	)
}
```

### Rate Limiting

The Auth0 Management API imposes a rate limit on all API clients. When the limit is reached, the SDK will handle it in
the background by retrying the API request when the limit is lifted.

> **Note**
> The SDK does not prevent `http.StatusTooManyRequests` errors, instead it waits for the rate limit to be reset based on
> the value of the `X-Rate-Limit-Reset` header as the amount of seconds to wait.

### Request Options

Fine-grained configuration can be provided on a per-request basis to enhance the request with specific query params, headers
or to pass it a custom context.

```go
// Example
userGrants, err := auth0API.Grant.List(
	management.Context(ctx)
	management.Header("MySpecialHeader","MySpecialHeaderValue")
    management.Parameter("user_id", "someUserID"),
    management.Parameter("client", "someClientID"),
    management.Parameter("audience", "someAPIAudience"),
)

// Other helper funcs.
management.Query()
management.ExcludeFields()
management.IncludeFields()
management.Page()
management.PerPage()
management.Take()
```

### Pagination

When retrieving lists of resources, if no query parameters are passed,
the following query parameters will get added by default: `?per_page=50,include_totals=true`. 

> **Note**
> The maximum value of the per_page query param is 100.

To get more than 50 results (the default), iterate through the returned pages.

```go
// Example
var page int
for {
    clients, err := auth0API.Client.List(management.Page(page))
    if err != nil {
        return err
    }
    
    // Accumulate here the results or check for a specific client.
    
    if !clients.HasNext() {
        break
    }

    page++
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

<p align="center">This project is licensed under the MIT license. See the <a href="./LICENSE.md"> LICENSE</a> file for more info.</p>
