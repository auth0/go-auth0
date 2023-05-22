# Go-Auth0 v1 Migration Guide


**Please review this guide thoroughly to understand the changes required to migrate your usage of go-auth0 to v1**

- [Public API Changes](#public-api-changes)
  - [Improve Typing Of Client Addons](#improve-typing-of-client-addons)
  - [Require Passing Context To APIs](#require-passing-context-to-apis)

## Public API Changes

### Improve Typing Of Client Addons

The `Addons` property of a `Client` was previously simply typed as `map[string]interface{}`, now it is explicitly typed to allow better support of addons.

However, with this there are now only two explicitly supported Addon types, `SAML2` and `WS-FED` which are the 2 addons supported in the Management UI. If you require a specific Addon, please open a [feature request](https://github.com/auth0/go-auth0/issues/new?assignees=&labels=feature&projects=&template=feature_request.yml)

<table>
<tr>
<th>Before (v0.17.1)</th>
<th>After (v1.0.0)</th>
</tr>
<tr>
<td>

```go
client := &management.Client{ 
    Name: auth0.String("My Client"),
    Addons: map[string]interface{}{
        "samlp": map[string]interface{}{
            "audience": "my-audience",
        },
        "wsfed": map[string]interface{}{},
    },
}
```
</td>
<td>

```go
client := &management.Client{
    Name: auth0.String("My Client"),
    Addons: &management.ClientAddons{
        SAML2: &management.SAML2ClientAddon{
            Audience: auth0.String("my-audience"),
        },
        WSFED: &management.WSFEDClientAddon{},
    },
}
```
</td>
</tr>
</table>

### Require Passing Context To APIs

All relevant methods have now been updated to require a [`context.Context`](https://pkg.go.dev/context?utm_source=godoc) argument as the first parameter. This applies to all management APIs and the helpers used to configure the `ClientCredential` used.

Subsequently, the `management.Context` and `management.WithContext` helpers has been removed.

<table>
<tr>
<th>Before (v0.17.1)</th>
<th>After (v1.0.0)</th>
</tr>
<tr>
<td>

```go
// Context could be set on the Management Client for the ClientCredential methods to use.
m, err := management.New(
    domain,
    management.WithClientCredentials(id, secret),
    management.WithContext(context.Background())
)

// Context could be passed to individual requests with the `Context` method.
err := auth0API.Client.Create(
    client,
    management.Context(context.Background()),
)
```
</td>
<td>

```go
// Context should be passed directly to the `WithClientCredentials` or `WithClientCredentialsAndAudience` methods.
m, err := management.New(
    domain,
    management.WithClientCredentials(context.Background(), id, secret),
)

// Context should be passed to to the method as the first argument.
err := auth0API.Client.Create(
    context.Background(),
    client
)
```
</td>
</tr>
</table>