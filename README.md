# Auth0 Go SDK

[![GoDoc](https://pkg.go.dev/badge/github.com/auth0/go-auth0.svg)](https://pkg.go.dev/github.com/auth0/go-auth0)
[![License](https://img.shields.io/github/license/auth0/go-auth0.svg?style=flat-square)](https://github.com/auth0/go-auth0/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/auth0/go-auth0?include_prereleases&style=flat-square)](https://github.com/auth0/go-auth0/releases)
[![Tests](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fauth0%2Fgo-auth0%2Fbadge%3Fref%3Dmain&style=flat-square)](https://github.com/auth0/go-auth0/actions?query=branch%3Amain)
[![Codecov](https://img.shields.io/codecov/c/github/auth0/go-auth0?style=flat-square)](https://codecov.io/gh/auth0/go-auth0)

---

Go client library for the [Auth0](https://auth0.com/) platform.

_Note: This SDK was previously maintained under [go-auth0/auth0](https://github.com/go-auth0/auth0)._

-------------------------------------

## Table of Contents

- [Installation](#installation)
- [Documentation](#documentation)
- [Usage](#usage)
- [Testing](#testing)
- [What is Auth0?](#what-is-auth0)
- [Create a free Auth0 Account](#create-a-free-auth0-account)
- [Issue Reporting](#issue-reporting)
- [Author](#author)
- [License](#license)

-------------------------------------

## Installation

```shell
go get github.com/auth0/go-auth0
```

[[table of contents]](#table-of-contents)

## Documentation

Reference documentation can be found at [pkg.go.dev](https://pkg.go.dev/gopkg.in/auth0.v5).
For more information about [Auth0](http://auth0.com/) please visit the [Auth0 Docs](http://docs.auth0.com/) page and the
[Auth0 Management API Docs](https://auth0.com/docs/api/management/v2).

[[table of contents]](#table-of-contents)

## Usage

```go
import (
	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/management"
)
```

Initialize a new client using a domain, client ID and secret.

```go
m, err := management.New(domain, management.WithClientCredentials(id, secret))
if err != nil {
	// handle err
}
```

With the management client we can now interact with the Auth0 Management API.

```go
c := &management.Client{
	Name:        auth0.String("Client Name"),
	Description: auth0.String("Long description of client"),
}

err = m.Client.Create(c)
if err != nil {
	// handle err
}

fmt.Printf("Created client %s\n", c.ClientID)
```

The following Auth0 resources are supported:

- [x] [Actions](https://auth0.com/docs/api/management/v2/#!/Actions/get_actions)
- [x] [Branding](https://auth0.com/docs/api/management/v2/#!/Branding/get_branding)
- [x] [Clients (Applications)](https://auth0.com/docs/api/management/v2#!/Clients/get_clients)
- [x] [Client Grants](https://auth0.com/docs/api/management/v2#!/Client_Grants/get_client_grants)
- [x] [Connections](https://auth0.com/docs/api/management/v2#!/Connections/get_connections)
- [x] [Custom Domains](https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains)
- [ ] [Device Credentials](https://auth0.com/docs/api/management/v2#!/Device_Credentials/get_device_credentials)
- [x] [Grants](https://auth0.com/docs/api/management/v2#!/Grants/get_grants)
- [x] [Hooks](https://auth0.com/docs/api/management/v2#!/Hooks/get_hooks)
- [x] [Hook Secrets](https://auth0.com/docs/api/management/v2/#!/Hooks/get_secrets)
- [x] [Log Streams](https://auth0.com/docs/api/management/v2#!/Log_Streams/get_log_streams)
- [x] [Logs](https://auth0.com/docs/api/management/v2#!/Logs/get_logs)
- [x] [Organizations](https://auth0.com/docs/api/management/v2#!/Organizations/get_organizations)
- [x] [Prompts](https://auth0.com/docs/api/management/v2#!/Prompts/get_prompts)
- [x] [Resource Servers (APIs)](https://auth0.com/docs/api/management/v2#!/Resource_Servers/get_resource_servers)
- [x] [Roles](https://auth0.com/docs/api/management/v2#!/Roles)
- [x] [Rules](https://auth0.com/docs/api/management/v2#!/Rules/get_rules)
- [x] [Rules Configs](https://auth0.com/docs/api/management/v2#!/Rules_Configs/get_rules_configs)
- [x] [User Blocks](https://auth0.com/docs/api/management/v2#!/User_Blocks/get_user_blocks)
- [x] [Users](https://auth0.com/docs/api/management/v2#!/Users/get_users)
- [x] [Users By Email](https://auth0.com/docs/api/management/v2#!/Users_By_Email/get_users_by_email)
- [x] [Blacklists](https://auth0.com/docs/api/management/v2#!/Blacklists/get_tokens)
- [x] [Email Templates](https://auth0.com/docs/api/management/v2#!/Email_Templates/get_email_templates_by_templateName)
- [x] [Emails](https://auth0.com/docs/api/management/v2#!/Emails/get_provider)
- [x] [Guardian](https://auth0.com/docs/api/management/v2#!/Guardian/get_factors)
- [x] [Jobs](https://auth0.com/docs/api/management/v2#!/Jobs/get_jobs_by_id)
- [x] [Stats](https://auth0.com/docs/api/management/v2#!/Stats/get_active_users)
- [x] [Tenants](https://auth0.com/docs/api/management/v2#!/Tenants/get_settings)
- [X] [Anomaly](https://auth0.com/docs/api/management/v2#!/Anomaly/get_ips_by_id)
- [x] [Tickets](https://auth0.com/docs/api/management/v2#!/Tickets/post_email_verification)
- [x] [Signing Keys](https://auth0.com/docs/api/management/v2#!/Keys/get_signing_keys)

[[table of contents]](#table-of-contents)

## Testing

The tests must run against an Auth0 tenant. They also need an
[M2M app](https://auth0.com/docs/applications/set-up-an-application/register-machine-to-machine-applications) in that
tenant that has been authorized to call the Management API. You can easily set one of these up by creating an
[API Explorer Application](https://auth0.com/docs/tokens/management-api-access-tokens/create-and-authorize-a-machine-to-machine-application)
in your tenant.

Then simply create a local `.env` file with the following settings:

* `AUTH0_DOMAIN`: The **Domain** of the M2M app
* `AUTH0_CLIENT_ID`: The **Client ID** of the M2M app
* `AUTH0_CLIENT_SECRET`: The **Client Secret** of the M2M app
* `AUTH0_DEBUG`: Set to `true` to call the Management API in debug mode, which dumps the HTTP requests and responses to the output

[[table of contents]](#table-of-contents)

## What is Auth0?

Auth0 helps you to:

- Add authentication with [multiple authentication sources](https://docs.auth0.com/identityproviders), either social like **Google, Facebook, Microsoft Account, LinkedIn, GitHub, Twitter, Box, Salesforce, amont others**, or enterprise identity systems like **Windows Azure AD, Google Apps, Active Directory, ADFS or any SAML Identity Provider**.
- Add authentication through more traditional **[username/password databases](https://docs.auth0.com/mysql-connection-tutorial)**.
- Add support for **[linking different user accounts](https://docs.auth0.com/link-accounts)** with the same user.
- Support for generating signed [Json Web Tokens](https://docs.auth0.com/jwt) to call your APIs and **flow the user identity** securely.
- Analytics of how, when and where users are logging in.
- Pull data from other sources and add it to the user profile, through [JavaScript rules](https://docs.auth0.com/rules).

[[table of contents]](#table-of-contents)

## Create a free Auth0 Account

1.  Go to [Auth0](https://auth0.com) and click "Try Auth0 for Free".
2.  Use Google, GitHub or Microsoft Account to login.

[[table of contents]](#table-of-contents)

## Issue Reporting

If you have found a bug or if you have a feature request, please report them at this repository issues section.
Please do not report security vulnerabilities on the public GitHub issue tracker.
The [Responsible Disclosure Program](https://auth0.com/whitehat) details the procedure for disclosing security issues.

[[table of contents]](#table-of-contents)

## Author

[Auth0](https://auth0.com/)

[[table of contents]](#table-of-contents)

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.

[[table of contents]](#table-of-contents)
