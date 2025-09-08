# Contributing

We appreciate feedback and contribution to this repo.
Before you submit a pull request, there are a couple requirements to satisfy.

## Prerequisites

Ensure you have the following prerequisites met:

- **Go Version**: You must have Go version 1.24 or newer installed on your system. You can check your Go version by running `go version` in your terminal.
- **GOPATH Set Up**: Make sure you have set up your GOPATH.
- **Export GOPATH**: Ensure you have exported your GOPATH to your system's environment variables. This allows tools and libraries to locate your Go workspace.
- **Modify bashrc or zshrc**: To automatically set your GOPATH and export it every time you open a terminal, you can add the following lines to your `~/.bashrc` or `~/.zshrc` file on Linux and macOS systems:
- **Environment File**: Ensure you have a `.env` file in the root folder of the repository or that the environment variables have been exported.

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

After adding these lines, run source `~/.bashrc` or source `~/.zshrc` to apply the changes.

## About This SDK

This SDK is auto-generated from Auth0's OpenAPI specifications. Most of the management API code is automatically generated and should **not** be manually modified.

**Manual modifications are only allowed in the `authentication` package**, which contains hand-written code for authentication flows.

## Safe Getters

The SDK offers safe getters to access pointer fields and avoid panics on nil pointers.

```go
// Example
client, err := mgmt.Client.Read(context.Background(), "EXAMPLE_16L9d34h0qe4NVE6SaHxZEid")
if err != nil {
    return err
}

// GetLogoURI() is safe to use as it will return "" if client.LogoURI is nil.
fmt.Printf("Logo URI: %v", client.GetLogoURI())
```

The getter methods are automatically generated as part of the SDK generation process.

## Code Quality and Security

### Linting

Before submitting your pull request, ensure your code follows the project's coding standards by running the linting task:

```sh
make lint
```

This will run golangci-lint across the codebase, identifying potential stylistic issues and suspicious constructs.

### Vulnerability Checking

To check for any known vulnerabilities in the dependencies, run:

```sh
make check-vuln
```

This uses govulncheck to scan the codebase for known security issues in the Go ecosystem.

### Code Validation

For the auto-generated management API code, validation is handled during the generation process. For manual changes to the authentication package, ensure your code follows Go best practices and doesn't break existing functionality by running the standard Go tools and tests.

## Running the Tests

There are two ways of running the tests:

- `make test` - runs the tests with http recordings. To run a specific test pass the `FILTER` var. Usage `make test FILTER="TestResourceServer_Read"`.
- `make test-e2e` - runs the tests against a real Auth0 tenant. To run a specific test pass the `FILTER` var. Usage `make test-record FILTER="TestResourceServer_Read"`.

### Running against an Auth0 tenant

To run the tests against an Auth0 tenant start by creating an M2M app using `auth0 apps create --name go-auth0-mgmt-tests --description "App used for go-auth0 management tests" --type m2m`, then
run `auth0 apps open <CLIENT ID>`. Authorize the Management API in the `APIs` tab and enable all permissions.

Then create a `.env` file in the root folder of the repository with the following details :

- `AUTH0_DOMAIN`: The **Domain** of the Auth0 tenant
- `AUTH0_CLIENT_ID`: The **Client ID** of the M2M app
- `AUTH0_CLIENT_SECRET`: The **Client Secret** of the M2M app
- `AUTH0_DEBUG`: Set to `true` to call the Management API in debug mode, which dumps the HTTP requests and responses to the output

Now for the Authentication tests create another M2M app using `auth0 apps create --name go-auth0-auth-tests --description "App used for go-auth0 authentication tests" --type m2m`, then run
`auth0 apps open <CLIENT ID>`. Ensure all `Grant Types` except `Client Credentials` are enabled in `Advanced Settings`, then set the `Authentication Method` to `None` in the `Credentials` tab.

Then create a `.env` file in the root folder of the repository with the following details :

- `AUTH0_DOMAIN`: The **Domain** of the Auth0 tenant
- `AUTH0_CLIENT_ID`: The **Client ID** of the management M2M app
- `AUTH0_CLIENT_SECRET`: The **Client Secret** of the management M2M app
- `AUTH0_AUTH_CLIENT_ID`: The **Client ID** of the authentication M2M app
- `AUTH0_AUTH_CLIENT_SECRET`: The **Client Secret** of the authentication M2M app
- `AUTH0_DEBUG`: Set to `true` to call the Management API in debug mode, which dumps the HTTP requests and responses to the output

> **Note**
> The http test recordings can be found in the [recordings](./test/data/recordings) folder.

## Adding New HTTP Test Recordings

To add new http test recordings or to regenerate old ones you can make use of the `make test-record` command.

> **Warning**
> If you need to regenerate an old recording, make sure to delete the corresponding recording file first.

To add only one specific http test recording pass the `FILTER` var, e.g. `make test-record FILTER="TestResourceServer_Read"`.

> **Warning**
> Recording a new http test interaction will make use of a real Auth0 test tenant.
