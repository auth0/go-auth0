# Change Log

<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/auth0/go-auth0/tree/v0.5.0) (2022-02-14)

This project is a continuation of [go-auth0/auth0](https://github.com/go-auth0/auth0).
To view previous release notes please check [CHANGELOG.md](https://github.com/go-auth0/auth0/blob/master/CHANGELOG.md).

**Breaking Changes**

- Renamed `client.ClientCredentials` to `client.OAuth2ClientCredentials`
- Renamed `ActionVersionError.Url` to `ActionVersionError.URL`
- Renamed `ActionBindingReferenceById` to `ActionBindingReferenceByID`
- Renamed `ConnectionOptionsSMS.GatewayUrl` to `ConnectionOptionsSMS.GatewayURL`
- Renamed `Connection.ProvisioningTicketUrl` to `Connection.ProvisioningTicketURL`

**Deprecations**

- Deprecated `ActionManager.ListTriggers()`
- Deprecated `ActionManager.ReadVersion()`
- Deprecated `ActionManager.ListVersions()`
- Deprecated `ActionManager.ListBindings()`
- Deprecated `ActionManager.ReadExecution()`
- Deprecated `ClientRefreshToken.Type`
- Deprecated `WithFields()`
- Deprecated `WithoutFields()`
- Deprecated `TenantFlags.ChangePasswordFlowV1`
