# Change Log

<a name="v0.9.0"></a>
## [v0.9.0](https://github.com/auth0/go-auth0/tree/v0.9.0) (2022-07-12)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.8.0...v0.9.0)

### Added

- Added `session_cookie` tenant property ([#88](https://github.com/auth0/go-auth0/pull/88))
- Added `upstream_params` connection property ([#89](https://github.com/auth0/go-auth0/pull/89))
- Added `include_email_in_redirect` email template property ([#90](https://github.com/auth0/go-auth0/pull/90))

<a name="v0.8.0"></a>
## [v0.8.0](https://github.com/auth0/go-auth0/tree/v0.8.0) (2022-07-06)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.7.0...v0.8.0)

### Added

- Added `DisableSignOut` field to SAMLP connection options ([#78](https://github.com/auth0/go-auth0/pull/78))
- Added several missing tenant flags ([#80](https://github.com/auth0/go-auth0/pull/80))
- Added `PKCEEnabled` field to Oauth2 connection options ([#82](https://github.com/auth0/go-auth0/pull/82))
- Added `Read()` and `Update()` to WebAuthn Roaming Settings ([#83](https://github.com/auth0/go-auth0/pull/83))
- Added `Read()` and `Update()` to WebAuthn Platform Settings ([#83](https://github.com/auth0/go-auth0/pull/83))
- Added `Read()` and `Update()` to Duo Settings ([#84](https://github.com/auth0/go-auth0/pull/84))
- Added `Read()` and `Update()` to Push CustomApp Settings ([#85](https://github.com/auth0/go-auth0/pull/85))
- Added `Enable()` Recovery Code MFA ([#86](https://github.com/auth0/go-auth0/pull/86))


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/auth0/go-auth0/tree/v0.7.0) (2022-06-22)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.6.4...v0.7.0)

### Added

- Added `icon_url` to OAuth2 connection options ([#74](https://github.com/auth0/go-auth0/pull/74))

### Changed

- **[Breaking]** Changed `AuthParams` to an `interface{}` in Email and SMS connection options ([#75](https://github.com/auth0/go-auth0/pull/75))


<a name="v0.6.4"></a>
## [v0.6.4](https://github.com/auth0/go-auth0/tree/v0.6.4) (2022-06-08)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.6.3...v0.6.4)

### Added

- Added support for webauthn_platform_first_factor_prompt flag in the prompt ([#59](https://github.com/auth0/go-auth0/pull/59))

### Changed

- Bumped Go version to 1.18 ([#71](https://github.com/auth0/go-auth0/pull/71))
- Ensured that all the tests can be run in any order 
- Ensured that all the tests clean up the test tenant afterwards of any created resources
- Enabled http recordings with go-vcr to be used within tests for more reliable testing


<a name="v0.6.3"></a>
## [v0.6.3](https://github.com/auth0/go-auth0/tree/v0.6.3) (2022-04-13)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.6.2...v0.6.3)

### Fixed

- Fixed uri path escaping for param IDs that have a forward slash in them ([#40](https://github.com/auth0/go-auth0/pull/40))


<a name="v0.6.2"></a>
## [v0.6.2](https://github.com/auth0/go-auth0/tree/v0.6.2) (2022-04-07)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.6.1...v0.6.2)

### Added

- Added a method to `Unlink` a `User`'s identity ([#35](https://github.com/auth0/go-auth0/pull/35))
- Added fields required to support self-managed certificates ([#37](https://github.com/auth0/go-auth0/pull/37))
- Added `profileData` key to `UserIdentity` ([#33](https://github.com/auth0/go-auth0/pull/33))
- [DXCDT-104] Added `Filters` to `LogStream` ([#38](https://github.com/auth0/go-auth0/pull/38))


<a name="v0.6.1"></a>
## [v0.6.1](https://github.com/auth0/go-auth0/tree/v0.6.1) (2022-03-03)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.6.0...v0.6.1)

### Added

- Added `ShowAsButton` option on `Connection` ([#29](https://github.com/auth0/go-auth0/pull/29))

### Fixed

- Fixed json tags on AllowList for Attack Protection structs ([#30](https://github.com/auth0/go-auth0/pull/30))

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
