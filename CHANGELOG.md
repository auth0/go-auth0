# Change Log

<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/auth0/go-auth0/tree/v0.6.0) (2022-02-22)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.5.0...v0.6.0)

### Added 

- [DXCDT-44] Add attack protection endpoints ([#27](https://github.com/auth0/go-auth0/pull/27))
- Added missing field `ClientID` to `Ticket` ([#25](https://github.com/auth0/go-auth0/pull/25/))
- Clarify intended usage of roles in app metadata vs RBAC ([#24](https://github.com/auth0/go-auth0/pull/24))

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
- Removed deprecated `ActionManager.ListTriggers()`
- Removed deprecated `ActionManager.ReadVersion()`
- Removed deprecated `ActionManager.ListVersions()`
- Removed deprecated `ActionManager.ListBindings()`
- Removed deprecated `ActionManager.ReadExecution()`
- Removed deprecated `ClientRefreshToken.Type`
- Removed deprecated `WithFields()`
- Removed deprecated `WithoutFields()`
- Removed deprecated `TenantFlags.ChangePasswordFlowV1`
