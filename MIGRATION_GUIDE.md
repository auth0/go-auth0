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

All relevant methods have now been updated to require a [`context.Context`](https://pkg.go.dev/context?utm_source=godoc) argument as the first parameter. This applies to all management APIs and the helpers used to configure the `ClientCredential`.

Subsequently, the `management.Context` and `management.WithContext` helpers have been removed.

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

// Context should be passed to the method as the first argument.
err := auth0API.Client.Create(
    context.Background(),
    client
)
```
</td>
</tr>
</table>

### Removal Of Manager Chaining

Previously, the individual Manager instances in the Management Client could be chained when calling them, this was unintended and has now been removed.

This should improve code completion features such as IntelliSense to only provide the relevant methods rather than all the individual Managers.

<table>
<tr>
<th>Before (v0.17.1)</th>
<th>After (v1.0.0)</th>
</tr>
<tr>
<td>

```go
// Managers could be chained before calling the desired method.
api.Action.Client.Action.Client.Create()
```
</td>
<td>

```go
// Chaining is no longer supported.
api.Client.Create()
```
</td>
</tr>
</table>

### Removal Of Email

In version `0.13.0` the `Email` manager was deprecated in favour of the `EmailProvider` manager to improve the configurations provided. `Email` has now been removed so the `EmailProvider` manager should be used instead.

<table>
<tr>
<th>Before (v0.17.1)</th>
<th>After (v1.0.0)</th>
</tr>
<tr>
<td>

```go
// `Email` manager was used and there was only one type for `Credentials` and `Settings`.
emailConfig := &management.Email{
    Name:               auth0.String("smtp"),
    Enabled:            auth0.Bool(true),
    DefaultFromAddress: auth0.String("no-reply@example.com"),
    Credentials: &management.EmailCredentials{
        SMTPHost: auth0.String("smtp.example.com"),
        SMTPPort: auth0.Int(587),
        SMTPUser: auth0.String("user"),
        SMTPPass: auth0.String("pass"),
    },
    Settings: map[string]interface{}{
        "Headers": &map[string]interface{}{
            "XSESConfigurationSet": auth0.String("example"),
        },
    },
}
```
</td>
<td>

```go
// Use the `EmailProvider` manager and use the provider specific configuration for `Credential` and `Settings`
emailProviderConfig := &management.EmailProvider{
    Name:               auth0.String("mandrill"),
    Enabled:            auth0.Bool(true),
    DefaultFromAddress: auth0.String("no-reply@example.com"),
    Credentials: &management.EmailProviderCredentialsSMTP{
        SMTPHost: auth0.String("smtp.example.com"),
        SMTPPort: auth0.Int(587),
        SMTPUser: auth0.String("user"),
        SMTPPass: auth0.String("pass"),
    },
    Settings: &management.EmailProviderSettingsSMTP{
        Headers: &management.EmailProviderSettingsSMTPHeaders{
            XSESConfigurationSet: auth0.String("example"),
        },
    },
}
```
</td>
</tr>
</table>
