# Go-Auth0 v1 Migration Guide


**Please review this guide thoroughly to understand the changes required to migrate your usage of go-auth0 to v1**

- [Public API Changes](#public-api-changes)
  - [Improve Typing Of Client Addons](#improve-typing-of-client-addons)
  - [Require Passing Context To APIs](#require-passing-context-to-apis)
  - [Removal Of Manager Chaining](#removal-of-manager-chaining)
  - [Changes To Retry Strategy](#changes-to-retry-strategy)

## Public API Changes

### Improve Typing Of Client Addons

The `Addons` property of a `Client` was previously simply typed as `map[string]interface{}`, now it is explicitly typed to allow better support of addons.

All addons that are listed in the [API2 documentation](https://auth0.com/docs/api/management/v2/clients/get-clients) have been typed to allow usage.

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
// Use the `EmailProvider` manager and use the provider specific configuration for `Credential` and `Settings`.
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

### Changes To Retry Strategy

Previously, when retrying requests the SDK would retry only for 429 status codes, waiting until the time indicated by the `X-RateLimit-Reset` header and would not have a maximum number of requests (i.e would retry infinitely).

Now, by default the SDK will retry 429 status codes and request errors that are deemed to be recoverable, and will retry a maximum of 2 times. The SDK will instead use an exponential backoff with a minimum time of 250 milliseconds, and a maximum delay time of 10 seconds and consult the `X-RateLimit-Reset` header to ensure that the next request will pass.

This logic can be customised by using the `management.WithRetries` helper, and can be disabled using the `management.WithNoRetries` method.

```go
// Enable a retry strategy with 3 retries that will retry on 429 and 503 status codes
m, err := management.New(
    domain,
    management.WithClientCredentials(context.Background(), id, secret),
    management.WithRetries(3, []int{http.StatusTooManyRequests, http.StatusBadGateway}),
)

// Disable retries
m, err := management.New(
    domain,
    management.WithClientCredentials(context.Background(), id, secret),
    management.WithNoRetries(),
)
```

### API Removals

Some APIs have been renamed or removed in order to ensure consistency in the SDK. Here is the list of removed/renamed APIs and the recommended replacements 

|Removed API|Reasoning|Resolution|
|-----------|--------------|--------------|
|`Organization.DeleteMember`|Renamed to align with other APIs|`Organization.DeleteMembers`|
|`ResourceServer.Stream`|Removed to allow potential for consistent pagination functionality across the SDK| Implement pagination manually similar to other APIs|
|`ConnectionStrategyDigitalOcean`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategyDiscord`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategyFigma`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategyImgur`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategySlack`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategySpotify`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategyTwitch`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|
|`ConnectionStrategyVimeo`|No longer supported as a strategy type|Use the `ConnectionStrategyOAuth2` strategy|

### Changes To Default User Agent

The default user agent string has changed from `Go-Auth0-SDK/<version>` to `Go-Auth0/<version>`, please be aware however that this is not part of the supported public API so could change in any future release and should not be relied upon.
