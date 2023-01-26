# Contributing

We appreciate feedback and contribution to this repo.
Before you get started, there are a few guidelines you need to follow.

## Generating the Getters

The SDK offers safe getters in order to access pointer fields to avoid panics on nil pointers.

```go
// Example
client, err := auth0API.Client.Read("EXAMPLE_16L9d34h0qe4NVE6SaHxZEid")
if err != nil{
    return err
}

// GetLogoURI() is safe to use as it will return "" if client.LogoURI is nil. 
fmt.Printf("Logo URI: %v", client.GetLogoURI())
```

When adding a new pointer field to a struct, run `make generate` to generate the getter for it.

## Running the Tests

There are two ways of running the tests:

- `make test` - runs the tests with http recordings. To run a specific test pass the `FILTER` var. Usage `make test FILTER="TestResourceServer_Read"`.
- `make test-e2e` - runs the tests against a real Auth0 tenant. To run a specific test pass the `FILTER` var. Usage `make test-record FILTER="TestResourceServer_Read"`.

To run the tests against an Auth0 tenant start by creating an
[M2M app](https://auth0.com/docs/applications/set-up-an-application/register-machine-to-machine-applications) in the
tenant, that has been authorized to request access tokens for the Management API and has all the required permissions.

Then create a local `.env` file in the `management` folder with the following settings:

* `AUTH0_DOMAIN`: The **Domain** of the M2M app
* `AUTH0_CLIENT_ID`: The **Client ID** of the M2M app
* `AUTH0_CLIENT_SECRET`: The **Client Secret** of the M2M app
* `AUTH0_DEBUG`: Set to `true` to call the Management API in debug mode, which dumps the HTTP requests and responses to the output

The http test recordings can be found in the [recordings](./test/data/recordings) folder.

## Adding New HTTP Test Recordings

To add new http test recordings or to regenerate old ones you can make use of the `make test-record` command.
To add only one specific http test recording pass the `FILTER` var, e.g. `make test-record FILTER="TestResourceServer_Read"`.

> **Warning**
> Recording a new http test interaction will make use of a real Auth0 test tenant. 
