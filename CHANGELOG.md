# Change Log

## [v1.26.0](https://github.com/auth0/go-auth0/tree/v1.26.0) (2025-08-11)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.25.0...v1.26.0)

**Added**
- feat: Add attribute `startFrom`  to Log Streams [\#584](https://github.com/auth0/go-auth0/pull/584) ([duedares-rvj](https://github.com/duedares-rvj))
- feat: add missing mgmt api fields to guardian enrollment ticket structure [\#587](https://github.com/auth0/go-auth0/pull/587) ([JohnRoesler](https://github.com/JohnRoesler))
- feat(risk-assessment): add `RiskAssessmentManager` for managing risk assessment settings and new device settings [\#582](https://github.com/auth0/go-auth0/pull/582) ([developerkunal](https://github.com/developerkunal))
- feat: support online refresh token and cascade revocation in ClientManager’s N2W session transfer [\#576](https://github.com/auth0/go-auth0/pull/576) ([nelsonmaia](https://github.com/nelsonmaia))

**Fixed**
- fix(mfa): deprecate OOBChannels and add OOBChannel in authenticator responses; improve test validation [\#583](https://github.com/auth0/go-auth0/pull/583) ([developerkunal](https://github.com/developerkunal))

## [v1.25.0](https://github.com/auth0/go-auth0/tree/v1.25.0) (2025-07-15)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.24.0...v1.25.0)

**Added**
- Add support for `domain metadata` in `CustomDomain` struct, implement `ListWithPagination` method, and update Management options with `WithCustomDomainHeader` [\#551](https://github.com/auth0/go-auth0/pull/551) ([developerkunal](https://github.com/developerkunal))
- feat(logstreams): Add PIIConfig to LogStream struct for PII masking support [\#575](https://github.com/auth0/go-auth0/pull/575) ([developerkunal](https://github.com/developerkunal))

## [v1.24.0](https://github.com/auth0/go-auth0/tree/v1.24.0) (2025-06-30)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.23.0...v1.24.0)

**Added**
- feat(prompt): add Filters and UsePageTemplate support to PromptRendering [\#573](https://github.com/auth0/go-auth0/pull/573) ([developerkunal](https://github.com/developerkunal))

## [v1.23.0](https://github.com/auth0/go-auth0/tree/v1.23.0) (2025-06-19)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.22.2...v1.23.0)

**Added**
- docs: update comment to reflect DPoP support in `ResourceServerManager` [\#571](https://github.com/auth0/go-auth0/pull/571) ([developerkunal](https://github.com/developerkunal))
- Add support for `ReadEnabledConnections` in ClientsManager [\#569](https://github.com/auth0/go-auth0/pull/569) ([developerkunal](https://github.com/developerkunal))
- Add Patch method and corresponding tests for Network ACL [\#568](https://github.com/auth0/go-auth0/pull/568) ([developerkunal](https://github.com/developerkunal))
- Add `ReadKeys` and `RotateKeys` methods and extend `ConnectionOptions` with `TokenEndpointAuthMethod` and `TokenEndpointAuthSigningAlg` in `ConnectionManager` [\#559](https://github.com/auth0/go-auth0/pull/559) ([developerkunal](https://github.com/developerkunal))

## [v1.22.2](https://github.com/auth0/go-auth0/tree/v1.22.2) (2025-06-09)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.22.1...v1.22.2)

**Fixed**
- Revert `DELETE` body restriction and fix `CleanForPatch` nil checks in Update methods [\#565](https://github.com/auth0/go-auth0/pull/565) ([developerkunal](https://github.com/developerkunal))

## [v1.22.1](https://github.com/auth0/go-auth0/tree/v1.22.1) (2025-06-03)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.22.0...v1.22.1)

**Fixed**
- Fix: Set Content-Type header to application/json for non-empty request bodies [\#562](https://github.com/auth0/go-auth0/pull/562) ([developerkunal](https://github.com/developerkunal))

## [v1.22.0](https://github.com/auth0/go-auth0/tree/v1.22.0) (2025-05-30)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.21.0...v1.22.0)

**Added**
- Add Support for Limiting M2M Usage via `Tenant-Wide` Defaults and `Client/Organization` Overrides [\#537](https://github.com/auth0/go-auth0/pull/537) ([developerkunal](https://github.com/developerkunal))

**Fixed**
- Fix: Prevent sending `{}` as body in requests for methods that don’t allow or expect a payload [\#546](https://github.com/auth0/go-auth0/pull/546) ([jeffmay](https://github.com/jeffmay))

## [v1.21.0](https://github.com/auth0/go-auth0/tree/v1.21.0) (2025-05-20)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.20.0...v1.21.0)

**⚠️ BREAKING CHANGES**
- Fix: Correct `CustomHeaders` type and JSON tag in `ConnectionOptionsOAuth2` within `ConnectionManager` [\#554](https://github.com/auth0/go-auth0/pull/554) ([developerkunal](https://github.com/developerkunal))

**Added**
- Add support for `FlexibleMappings` SAML2 mappings in SAML2ClientAddon and enhance serialization tests [\#555](https://github.com/auth0/go-auth0/pull/555) ([developerkunal](https://github.com/developerkunal))
- Add `ReadEnabledClients` and `UpdateEnabledClients` methods; deprecate `EnabledClients` field in `Connection` struct of `ConnectionManager` [\#556](https://github.com/auth0/go-auth0/pull/556) ([ErwinSteffens](https://github.com/ErwinSteffens))
- Add Support for `ListRendering` method and associated tests for `PromptManager` [\#552](https://github.com/auth0/go-auth0/pull/552) ([developerkunal](https://github.com/developerkunal))
- Add `AllowRefreshToken` support to `SessionTransfer`  in `ClientManager` [\#557](https://github.com/auth0/go-auth0/pull/557) ([nelsonmaia](https://github.com/nelsonmaia))

## [v1.20.0](https://github.com/auth0/go-auth0/tree/v1.20.0) (2025-05-05)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.19.0...v1.20.0)

**Added**
- Add support for setting per-attempt-timeout [\#539](https://github.com/auth0/go-auth0/pull/539) ([pete-woods](https://github.com/pete-woods))
- Add 'mgmt_api_read' log type to enhance log categorization in `LogManager` [\#545](https://github.com/auth0/go-auth0/pull/545) ([developerkunal](https://github.com/developerkunal))
- Adding support for Native to Web SSO - Session Transfer [\#536](https://github.com/auth0/go-auth0/pull/536) ([nelsonmaia](https://github.com/nelsonmaia))
- Add support to set RealmFallback to Connection Options [\#541](https://github.com/auth0/go-auth0/pull/541) ([duedares-rvj](https://github.com/duedares-rvj))
- Add Support Multi-Resource Refresh Token `Policies` in `ClientRefreshToken` within `Client Manager` [\#532](https://github.com/auth0/go-auth0/pull/532) ([developerkunal](https://github.com/developerkunal))
- Add Private Key JWT support for client credentials in Management API [\#528](https://github.com/auth0/go-auth0/pull/528) ([ErwinSteffens](https://github.com/ErwinSteffens))

## [v1.19.0](https://github.com/auth0/go-auth0/tree/v1.19.0) (2025-03-28)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.18.0...v1.19.0)

**⚠️ BREAKING CHANGES**
- Breaking Change: Migrate `custom_headers` from `map[string]string` to `*[]map[string]string` in `ConnectionOptionsOAuth2` [\#534](https://github.com/auth0/go-auth0/pull/534) ([developerkunal](https://github.com/developerkunal))

**Added**
- Add Support For Tenant ACL Endpoints to `NetworkACLManager` [\#504](https://github.com/auth0/go-auth0/pull/504) ([developerkunal](https://github.com/developerkunal))
- Add Extended Support for `DomainAliasesConfig` and `ConnectionConfig` Enhancements in `SelfServiceProfileTicket` in `SelfServiceProfileManager` [\#509](https://github.com/auth0/go-auth0/pull/509) ([developerkunal](https://github.com/developerkunal))
- Add Support for Breached Password Detection `(PreChangePassword)` Stage in `AttackProtectionManager` [\#499](https://github.com/auth0/go-auth0/pull/499) ([developerkunal](https://github.com/developerkunal))
- Add support for `GetUserLogs` method to fetch user log events in `UserManager` [\#529](https://github.com/auth0/go-auth0/pull/529) ([developerkunal](https://github.com/developerkunal))

## [v1.18.0](https://github.com/auth0/go-auth0/tree/v1.18.0) (2025-03-11)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.17.0...v1.18.0)

**Added**
- Add tenant OIDC logout configuration [\#517](https://github.com/auth0/go-auth0/pull/517) ([ErwinSteffens](https://github.com/ErwinSteffens))
- Add `custom_headers` Attribute to `ConnectionOptionsOAuth2` in `ConnectionManager` [\#522](https://github.com/auth0/go-auth0/pull/522) ([developerkunal](https://github.com/developerkunal))
- Add Support for `FCMv1 Server Credentials` in `MultiFactorPush` Sub-Manager of `GuardianManager` [\#512](https://github.com/auth0/go-auth0/pull/512) ([chrisnellis](https://github.com/chrisnellis))

**Changed**
- Bump Go version to 1.23 and upgrade golang.org/x/oauth2 to v0.28.0 [\#518](https://github.com/auth0/go-auth0/pull/518) ([developerkunal](https://github.com/developerkunal))
- refactor: use a single client assertion audience [\#513](https://github.com/auth0/go-auth0/pull/513) ([panva](https://github.com/panva))

**Fixed**
- [GH-501] Fix Retry Mechanism to Handle Burst Limit Due to Clock Skew Issue [\#523](https://github.com/auth0/go-auth0/pull/523) ([developerkunal](https://github.com/developerkunal))

## [v1.17.0](https://github.com/auth0/go-auth0/tree/v1.17.0) (2025-02-14)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.16.0...v1.17.0)

**Added**
- Add Support For Extensibility as Custom Provider in `BrandingManager` [\#495](https://github.com/auth0/go-auth0/pull/495) ([developerkunal](https://github.com/developerkunal))
- Add Support for `form-content` Insertion Point For Prompt Partials in `PromptManager` [\#503](https://github.com/auth0/go-auth0/pull/503) ([developerkunal](https://github.com/developerkunal))

## [v1.16.0](https://github.com/auth0/go-auth0/tree/v1.16.0) (2025-02-06)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.15.0...v1.16.0)

**Added**
- Add Support For Event Streams with `EventStreamManager` [\#496](https://github.com/auth0/go-auth0/pull/496) ([duedares-rvj](https://github.com/duedares-rvj))
- Add Support For Global Token Revocation To SAML Connections in `ConnectionManager` [\#491](https://github.com/auth0/go-auth0/pull/491) ([developerkunal](https://github.com/developerkunal))
- Add Support For `Sign in With Google For Native Apps` in `ClientManager` [\#493](https://github.com/auth0/go-auth0/pull/493) ([developerkunal](https://github.com/developerkunal))
- Add support for setting `captcha_widget_theme` in `BrandingTheme` Struct in `BrandingThemeManager` [\#492](https://github.com/auth0/go-auth0/pull/492) ([duedares-rvj](https://github.com/duedares-rvj))

## [v1.15.0](https://github.com/auth0/go-auth0/tree/v1.15.0) (2025-01-29)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.14.0...v1.15.0)

**Added**
- Add Support for Token Exchange Profile with `TokenExchangeProfileManager` [\#478](https://github.com/auth0/go-auth0/pull/478) ([developerkunal](https://github.com/developerkunal))
- Add Support for Refresh Tokens and Sessions Endpoints [\#484](https://github.com/auth0/go-auth0/pull/484) ([developerkunal](https://github.com/developerkunal))

## [v1.14.0](https://github.com/auth0/go-auth0/tree/v1.14.0) (2025-01-08)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.13.1...v1.14.0)

**Added**
- Add Support for User sessions Endpoints in `UserManager` [\#482](https://github.com/auth0/go-auth0/pull/482) ([developerkunal](https://github.com/developerkunal))
- Support for CIBA  [\#473](https://github.com/auth0/go-auth0/pull/473) ([duedares-rvj](https://github.com/duedares-rvj))
- Add Support for `VerificationMethod` Attribute for Email-based Auth0 ConnectionOptions [\#481](https://github.com/auth0/go-auth0/pull/481) ([developerkunal](https://github.com/developerkunal))

## [v1.13.1](https://github.com/auth0/go-auth0/tree/v1.13.1) (2024-12-20)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.13.0...v1.13.1)

**Fixed**
- PATCH: Fix logic to handle the PATCH payload for `renderingMode` as standard [\#476](https://github.com/auth0/go-auth0/pull/476) ([ramya18101](https://github.com/ramya18101))

## [v1.13.0](https://github.com/auth0/go-auth0/tree/v1.13.0) (2024-12-06)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.12.0...v1.13.0)

**Added**
- Add support for ACUL Endpoints in `PromptManager` [\#458](https://github.com/auth0/go-auth0/pull/458) ([developerkunal](https://github.com/developerkunal))

## [v1.12.0](https://github.com/auth0/go-auth0/tree/v1.12.0) (2024-11-28)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.11.2...v1.12.0)

**Added**
- Add isPriority parameter to Log Streams [\#467](https://github.com/auth0/go-auth0/pull/467) ([duedares-rvj](https://github.com/duedares-rvj))
- SSO v2 updates [\#457](https://github.com/auth0/go-auth0/pull/457) ([duedares-rvj](https://github.com/duedares-rvj))
- Add `DeviceCredentialsManager` to manage device credentials [\#369](https://github.com/auth0/go-auth0/pull/369) ([Zarux](https://github.com/Zarux))

## [v1.11.2](https://github.com/auth0/go-auth0/tree/v1.11.2) (2024-10-14)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.11.1...v1.11.2)

**Fixed**
- PATCH: Rename `user_id_attribute` to `userid_attribute` in Azure Connection Options [\#454](https://github.com/auth0/go-auth0/pull/454) ([duedares-rvj](https://github.com/duedares-rvj))

## [v1.11.1](https://github.com/auth0/go-auth0/tree/v1.11.1) (2024-10-07)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.11.0...v1.11.1)

**Added**
- Add missing support for `Custom` email provider in EmailProviderManager [\#452](https://github.com/auth0/go-auth0/pull/452) ([duedares-rvj](https://github.com/duedares-rvj))

## [v1.11.0](https://github.com/auth0/go-auth0/tree/v1.11.0) (2024-09-27)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.10.0...v1.11.0)

**Added**
- Add support for `Form`, `Flow`, and `Flow Vault Connection` Managers [\#444](https://github.com/auth0/go-auth0/pull/444) ([kushalshit27](https://github.com/kushalshit27))
- Add support for MFA Authentication Endpoints: Add, List, and Delete Authenticators [\#447](https://github.com/auth0/go-auth0/pull/447) ([developerkunal](https://github.com/developerkunal))
- Add `user_id_attribute` support to AzureAD connection options in Connection Manager [\#445](https://github.com/auth0/go-auth0/pull/445) ([acwest](https://github.com/acwest))
- Add `strategy_version` support to required connections in Connection Manager [\#443](https://github.com/auth0/go-auth0/pull/443) ([acwest](https://github.com/acwest))

## [v1.10.0](https://github.com/auth0/go-auth0/tree/v1.10.0) (2024-09-03)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.9.0...v1.10.0)

**Added**
- Add support for new Prompts and Screens for Prompt Partial [\#437](https://github.com/auth0/go-auth0/pull/437) ([developerkunal](https://github.com/developerkunal))
- Add Organizations for Client Credentials [\#432](https://github.com/auth0/go-auth0/pull/432) ([duedares-rvj](https://github.com/duedares-rvj))
- Add Support for Control Your Own Key (CYOK) and Bring Your Own Key (BYOK) Features with New `EncryptionKeyManager` [\#435](https://github.com/auth0/go-auth0/pull/435) ([developerkunal](https://github.com/developerkunal))

> **Note**
> The methods `ReadPartials`, `CreatePartials`, and `UpdatePartials` are deprecated. Please use `GetPartials` and `SetPartials` instead for managing Prompt Partials.

## [v1.9.0](https://github.com/auth0/go-auth0/tree/v1.9.0) (2024-08-16)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.8.0...v1.9.0)

**Added**
- Add Support for Self Service Profiles [\#431](https://github.com/auth0/go-auth0/pull/431) ([duedares-rvj](https://github.com/duedares-rvj))
- Add Support for HRI Features [\#429](https://github.com/auth0/go-auth0/pull/429) ([developerkunal](https://github.com/developerkunal))
- Add Support for Flexible Identifiers on ConnectionOptions & `phone_number` to SignupRequest Struct [\#421](https://github.com/auth0/go-auth0/pull/421) ([developerkunal](https://github.com/developerkunal))

## [v1.8.0](https://github.com/auth0/go-auth0/tree/v1.8.0) (2024-07-09)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.7.0...v1.8.0)

**Added**
- Add support for managing a connection's SCIM (System for Cross-domain Identity Management) configuration [\#419](https://github.com/auth0/go-auth0/pull/419) ([developerkunal](https://github.com/developerkunal))

## [v1.7.0](https://github.com/auth0/go-auth0/tree/v1.7.0) (2024-06-14)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.6.0...v1.7.0)

**Added**
- Add is_signup_enabled field to OrganizationConnection [\#413](https://github.com/auth0/go-auth0/pull/413) ([developerkunal](https://github.com/developerkunal))

## [v1.6.0](https://github.com/auth0/go-auth0/tree/v1.6.0) (2024-05-09)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.5.0...v1.6.0)

**Added**
- [GH-380] - Add User Refresh Token Management Functionality to SDK [\#404](https://github.com/auth0/go-auth0/pull/404) ([developerkunal](https://github.com/developerkunal))
- [GH-366] - Add support for enabled_connections field in CreateOrganizationConnections [\#394](https://github.com/auth0/go-auth0/pull/394) ([developerkunal](https://github.com/developerkunal))

**Fixed**
- [GH-372] Fix unmarshaling issue with Wordpress strategy connections [\#398](https://github.com/auth0/go-auth0/pull/398) ([developerkunal](https://github.com/developerkunal))
- [GH-336] - Fix Unmarshal JSON error as string [\#393](https://github.com/auth0/go-auth0/pull/393) ([developerkunal](https://github.com/developerkunal))

## [v1.5.0](https://github.com/auth0/go-auth0/tree/v1.5.0) (2024-04-23)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.4.1...v1.5.0)

**Added**
- Send extra parameters with client credentials request [\#354](https://github.com/auth0/go-auth0/pull/354) ([weirdian2k3](https://github.com/weirdian2k3))
- Add support for `oidc_logout` parameters [\#384](https://github.com/auth0/go-auth0/pull/384) ([developerkunal](https://github.com/developerkunal))
- Add `show_as_button` field to Organization Enabled Connection [\#386](https://github.com/auth0/go-auth0/pull/386) ([developerkunal](https://github.com/developerkunal))

**Fixed**
- Fix sending unnecessary `null` body in requests [\#387](https://github.com/auth0/go-auth0/pull/387) ([developerkunal](https://github.com/developerkunal))

## [v1.4.1](https://github.com/auth0/go-auth0/tree/v1.4.1) (2024-02-28)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.4.0...v1.4.1)

**Added**
- feat: Add Prompt Partials Support [\#341](https://github.com/auth0/go-auth0/pull/341) ([m3talsmith](https://github.com/m3talsmith)) [#360](https://github.com/auth0/go-auth0/pull/360) ([sergiught](https://github.com/sergiught))

## [v1.4.0](https://github.com/auth0/go-auth0/tree/v1.4.0) (2023-12-14)
[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.3.1...v1.4.0)

**Added**
- [SDK-4478] Support organization id when resetting a password [\#333](https://github.com/auth0/go-auth0/pull/333) ([ewanharris](https://github.com/ewanharris))
- Implement MFA authentication APIs [\#331](https://github.com/auth0/go-auth0/pull/331) ([ewanharris](https://github.com/ewanharris))
- Export an Authentication Error to allow type assertions [\#330](https://github.com/auth0/go-auth0/pull/330) ([ewanharris](https://github.com/ewanharris))
- [SDK-4738] Add support for performing Pushed Authorization Requests [\#327](https://github.com/auth0/go-auth0/pull/327) ([ewanharris](https://github.com/ewanharris))

**Removed**
- [DXEX-3404] chore: revert actions log sessions [\#325](https://github.com/auth0/go-auth0/pull/325) ([johneke-auth0](https://github.com/johneke-auth0))

<a name="v1.3.1"></a>

## [v1.3.1](https://github.com/auth0/go-auth0/tree/v1.3.1) (2023-11-28)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.3.0..v1.3.1)

### Changed

- Allow `Scopes` in `ClientGrant` to be sent as an empty array ([#318](https://github.com/auth0/go-auth0/pull/318))

<a name="v1.3.0"></a>

## [v1.3.0](https://github.com/auth0/go-auth0/tree/v1.3.0) (2023-11-10)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.2.0..v1.3.0)

### Added

- Add `DisableSelfServiceChangePassword` to `ConnectionOptionsAD` ([#308](https://github.com/auth0/go-auth0/pull/308))
- Add support for using Organizations with Client Grants ([#309](https://github.com/auth0/go-auth0/pull/309))

<a name="v1.2.0"></a>

## [v1.2.0](https://github.com/auth0/go-auth0/tree/v1.2.0) (2023-10-25)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.1.0..v1.2.0)

### Added

- Add `Roles` field to the `OrganizationMember` struct ([#293](https://github.com/auth0/go-auth0/pull/293))
- Add `CustomizeMFAInPostLoginAction` tenant setting ([#294](https://github.com/auth0/go-auth0/pull/294))
- Add Passkey data to `AuthenticationMethod` and `ConnectionOptions` ([#296](https://github.com/auth0/go-auth0/pull/296))
- Add a `Sort` RequestOption helper to sort fields ([#298](https://github.com/auth0/go-auth0/pull/298))

<a name="v1.1.0"></a>

## [v1.1.0](https://github.com/auth0/go-auth0/tree/v1.1.0) (2023-09-14)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.0.2..v1.1.0)

### Added

- Add `ConnectionSettings` and `AttributeMap` to `ConnectionOptionsOIDC` and `ConnectionOptionsOkta` ([#277](https://github.com/auth0/go-auth0/pull/277))
- Add properties to `ConnectionOptionsGoogleApps` ([#278](https://github.com/auth0/go-auth0/pull/278))
- Add `Scopes` and `SetScopes` for `ConnectionOptionsPingFederate` and `ConnectionOptionsSAML` ([#274](https://github.com/auth0/go-auth0/pull/274))
- Add support for `allow_organization_name_in_authentication_api` tenant setting ([#280](https://github.com/auth0/go-auth0/pull/280)/[#281](https://github.com/auth0/go-auth0/pull/281))

### Changed

- Bump Go version to 1.20 ([#279](https://github.com/auth0/go-auth0/pull/279))

### Fixed

- Allow sending null values for `From` and `MessagingServiceSID` in `ConnectionOptionsSMS` ([#275](https://github.com/auth0/go-auth0/pull/275))

<a name="v1.0.2"></a>

## [v1.0.2](https://github.com/auth0/go-auth0/tree/v1.0.2) (2023-08-30)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.0.1..v1.0.2)

### Added

- Add line connection strategy ([#271](https://github.com/auth0/go-auth0/pull/271))

### Fixed

- Add missing scope tags to connection options ([#269](https://github.com/auth0/go-auth0/pull/269))
- Fix incorrect json tag for `PreviousThumbprints` on `ConnectionOptionsADFS` ([#270](https://github.com/auth0/go-auth0/pull/270))

<a name="v1.0.1"></a>

## [v1.0.1](https://github.com/auth0/go-auth0/tree/v1.0.1) (2023-08-22)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v1.0.0..v1.0.1)

### Added

- Support Client Assertion authentication on Authentication Client ([#260](https://github.com/auth0/go-auth0/pull/260))
- Add more known properties onto `ad`, `adfs`, `pingfederate`, `saml`, and `waad` connection options ([#263](https://github.com/auth0/go-auth0/pull/263))

<a name="v1.0.0"></a>

## [v1.0.0](https://github.com/auth0/go-auth0/tree/v1.0.0) (2023-07-27)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.17.2..v1.0.0)

### Added

- Generate safe getters for `map[string]interface{}` types ([#205](https://github.com/auth0/go-auth0/pull/205))
- Expose configuration of retry strategy ([#216](https://github.com/auth0/go-auth0/pull/216))
- Authentication client ([#222](https://github.com/auth0/go-auth0/pull/222), [#227](https://github.com/auth0/go-auth0/pull/227), [#226](https://github.com/auth0/go-auth0/pull/226), [#229](https://github.com/auth0/go-auth0/pull/229), [#232](https://github.com/auth0/go-auth0/pull/232))
- Validate `org_name` claim in ID token ([#246](https://github.com/auth0/go-auth0/pull/246))
- Support `azure_cs` and `ms365` email providers ([#247](https://github.com/auth0/go-auth0/pull/247))
- Support `OIDCLogoutPrompt` setting on tenant ([#249](https://github.com/auth0/go-auth0/pull/249))
- Support `DecryptionKey` on SAML Connection Options ([#251](https://github.com/auth0/go-auth0/pull/251))

### Breaking Changes

- Improve typing of Client Addons ([#208](https://github.com/auth0/go-auth0/pull/208), [#228](https://github.com/auth0/go-auth0/pull/228))
- Accept a `context.Context` as the first argument for all methods ([#212](https://github.com/auth0/go-auth0/pull/212))
- Refactor Management struct to avoid chaining issue ([#214](https://github.com/auth0/go-auth0/pull/214))
- Removed deprecated `EmailManager` ([#218](https://github.com/auth0/go-auth0/pull/218))
- Removed `ResourceServer.Stream` ([#224](https://github.com/auth0/go-auth0/pull/224))
- Removed `DigitalOcean`, `Discord`, `Figma`, `Imgur`, `Slack`, `Spotify`, `Twitch`, and `Vimeo` connection strategies ([#245](https://github.com/auth0/go-auth0/pull/245))

Please review the [migration guide](./MIGRATION_GUIDE.md) to understand the changes required to update your application.

<a name="v1.0.0-beta.0"></a>

## [v1.0.0-beta.0](https://github.com/auth0/go-auth0/tree/v1.0.0-beta.0) (2023-06-26)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.17.2..v1.0.0-beta.0)

### Added

- Generate safe getters for `map[string]interface{}` types ([#205](https://github.com/auth0/go-auth0/pull/205))
- Expose configuration of retry strategy ([#216](https://github.com/auth0/go-auth0/pull/216))
- Add Authentication Client ([#222](https://github.com/auth0/go-auth0/pull/222), [#227](https://github.com/auth0/go-auth0/pull/227), [#226](https://github.com/auth0/go-auth0/pull/226), [#229](https://github.com/auth0/go-auth0/pull/229), [#232](https://github.com/auth0/go-auth0/pull/232))

### Breaking Changes

- Improve typing of Client Addons ([#208](https://github.com/auth0/go-auth0/pull/208), [#228](https://github.com/auth0/go-auth0/pull/228))
- Accept a `context.Context` as the first argument for all methods ([#212](https://github.com/auth0/go-auth0/pull/212))
- Refactor Management struct to avoid chaining issue ([#214](https://github.com/auth0/go-auth0/pull/214))
- Remove deprecated `EmailManager` ([#218](https://github.com/auth0/go-auth0/pull/218))
- Remove `ResourceServer.Stream` ([#224](https://github.com/auth0/go-auth0/pull/224))

Please review the [migration guide](./MIGRATION_GUIDE.md) to understand the changes required to update your application.

<a name="v0.17.2"></a>

## [v0.17.2](https://github.com/auth0/go-auth0/tree/v0.17.2) (2023-05-22)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.17.1...v0.17.2)

### Added

- Added support for updating the `Expiry` of a `Credential` ([#206](https://github.com/auth0/go-auth0/pull/206))/([#209](https://github.com/auth0/go-auth0/pull/209))

### Fixed

- Correct the `ForwardReqInfo` property on `ConnectionOptionsSMS` ([#207](https://github.com/auth0/go-auth0/pull/207))

<a name="v0.17.1"></a>

## [v0.17.1](https://github.com/auth0/go-auth0/tree/v0.17.1) (2023-05-15)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.17.0...v0.17.1)

### Added

- Added support for configuring `RequirePushedAuthorizationRequests` on Clients and Tenants ([#201](https://github.com/auth0/go-auth0/pull/201))
- Added support for settings the `OIDCBackchannelLogout` configuration for a Client ([#202](https://github.com/auth0/go-auth0/pull/202))

> **Note**
> At the time of release, both of these features require enablement for a tenant

<a name="v0.17.0"></a>

## [v0.17.0](https://github.com/auth0/go-auth0/tree/v0.17.0) (2023-04-19)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.16.0...v0.17.0)

### Added

- Added support for managing `APNS` and `FCM` provider configurations ([#184](https://github.com/auth0/go-auth0/pull/184))
- Added support for the `MFAShowFactorListOnEnrollment` tenant flag ([#187](https://github.com/auth0/go-auth0/pull/187))
- Added support for setting `OrganizationID` on a `Ticket` ([#195](https://github.com/auth0/go-auth0/pull/195))
- Added support for managing Client Credentials ([#196](https://github.com/auth0/go-auth0/pull/196))
- Added support for setting `Identity` and `OrganizationID` on `VerifyEmail` calls ([#197](https://github.com/auth0/go-auth0/pull/197))

### Fixed

- Fixed handling of `Log` entries where `Scope` is returned as an array of strings ([#194](https://github.com/auth0/go-auth0/pull/194))

<a name="v0.16.0"></a>

## [v0.16.0](https://github.com/auth0/go-auth0/tree/v0.16.0) (2023-03-15)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.15.1...v0.16.0)

### Added

- Introduce sending client information with requests ([#164](https://github.com/auth0/go-auth0/pull/164))
- Added support for `Ping Federate` connection ([#175](https://github.com/auth0/go-auth0/pull/175))
- Added support for Factor Management endpoints ([#176](https://github.com/auth0/go-auth0/pull/176))
- Added support for setting `disable_self_service_change_password` on database connections ([#178](https://github.com/auth0/go-auth0/pull/178))

### Fixed

- Fixed issue when decoding `ConnectionOptionsGoogleOAuth2` with `allowed_audiences` set as an empty string ([#174](https://github.com/auth0/go-auth0/pull/174))
- Fixed support for checkpoint pagination ([#179](https://github.com/auth0/go-auth0/pull/179))


<a name="v0.15.1"></a>

## [v0.15.1](https://github.com/auth0/go-auth0/tree/v0.15.1) (2023-01-30)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.15.0...v0.15.1)

### Added

- Added `EnableScriptContext` to `ConnectionOptions` ([#158](https://github.com/auth0/go-auth0/pull/158))
- Added `TrustEmailVerified`, `SignInEndpoint`, `Thumbprints`, `FedMetadataXML` to `ConnectionOptionsADFS` ([#161](https://github.com/auth0/go-auth0/pull/161))

### Fixed

- Fixed the `Connection.UnmarshalJSON()` for `ConnectionStrategyADFS`  ([#160](https://github.com/auth0/go-auth0/pull/160))

### Changed

- Changed the `CrossOriginAuth`'s json tag `cross_origin_auth` to `cross_origin_authentication` ([#159](https://github.com/auth0/go-auth0/pull/159))


<a name="v0.15.0"></a>

## [v0.15.0](https://github.com/auth0/go-auth0/tree/v0.15.0) (2023-01-26)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.14.0...v0.15.0)

### Added

- Added support for segment log stream type ([#151](https://github.com/auth0/go-auth0/pull/151))
- Added `Stage` field to `BreachedPasswordDetection` ([#148](https://github.com/auth0/go-auth0/pull/148))
- Added `LogoURL` field to `ConnectionOptionsOkta` ([#153](https://github.com/auth0/go-auth0/pull/153))


<a name="v0.14.0"></a>

## [v0.14.0](https://github.com/auth0/go-auth0/tree/v0.14.0) (2022-12-19)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.13.1...v0.14.0)

### Added

- Added ability to retrieve job errors ([#141](https://github.com/auth0/go-auth0/pull/141))
- Added summary to jobs ([#140](https://github.com/auth0/go-auth0/pull/140))

### Fixed

- Fixed import users job ([#139](https://github.com/auth0/go-auth0/pull/139))


<a name="v0.13.1"></a>

## [v0.13.1](https://github.com/auth0/go-auth0/tree/v0.13.1) (2022-12-14)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.13.0...v0.13.1)

### Added

- Added `Provider()` and `UpdateProvider()` to MFA Push ([#136](https://github.com/auth0/go-auth0/pull/136))
- Added Mixpanel `LogStream` ([#133](https://github.com/auth0/go-auth0/pull/133))


<a name="v0.13.0"></a>

## [v0.13.0](https://github.com/auth0/go-auth0/tree/v0.13.0) (2022-11-07)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.12.0...v0.13.0)

### Added

- Added support for creating action log sessions ([#127](https://github.com/auth0/go-auth0/pull/127))
- Added new email provider manager ([#129](https://github.com/auth0/go-auth0/pull/129))
- Added support for Okta connection ([#130](https://github.com/auth0/go-auth0/pull/130))

### Deprecated

- Deprecated email manager ([#129](https://github.com/auth0/go-auth0/pull/129))


<a name="v0.12.0"></a>

## [v0.12.0](https://github.com/auth0/go-auth0/tree/v0.12.0) (2022-10-18)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.11.0...v0.12.0)

### Added

- Added `omitempty` json tag to `IdentityAPI` on `ConnectionOptionsAzureAD` ([#119](https://github.com/auth0/go-auth0/pull/119))
- Added ability to retrieve the default `BrandingTheme` ([#122](https://github.com/auth0/go-auth0/pull/122))
- Added missing fields to `Log` ([#115](https://github.com/auth0/go-auth0/pull/115))
- Added `LastPasswordReset` to `User` ([#125](https://github.com/auth0/go-auth0/pull/125))

### ⚠️Breaking Changes

- Changed `ClientMetadata` to `*map[string]interface{}` in `Client` ([#120](https://github.com/auth0/go-auth0/pull/120))
- Changed `BrandingTheme` fields from `int` to `float64` to increase precision ([#121](https://github.com/auth0/go-auth0/pull/121))


<a name="v0.11.0"></a>

## [v0.11.0](https://github.com/auth0/go-auth0/tree/v0.11.0) (2022-10-07)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.10.1...v0.11.0)

### ⚠️ Notice of Breaking Changes

Please note that this version introduces many minor breaking changes. The intention with this release was to capture as many breaking changes into a single release to minimize future disruption.

These changes were introduced for three primary reasons: 
- Decoupling from [Auth0 Terraform Provider](https://github.com/auth0/terraform-provider-auth0) implementation
- More precise typing of struct properties for better DX
- Enabling Auth0 Terraform provider to explicitly set empty values (see: [related Github issue](https://github.com/auth0/terraform-provider-auth0/issues/14))

Review the complete list below before upgrading.

### Changed

- `Action` struct fields
  - `SupportedTriggers` to `[]ActionTrigger` type
  - `Dependencies` to ` *[]ActionDependency` type
  - `Secrets` to `*[]ActionSecret` type
- `Client` struct fields
  - `Callbacks` to `*[]string` type
  - `AllowedOrigins` to `*[]string` type
  - `WebOrigins` to `*[]string` type
  - `ClientAliases` to `*[]string` type
  - `AllowedClients` to `*[]string` type
  - `AllowedLogoutURLs` to `*[]string` type
  - `EncryptionKey` to `*map[string]string` type
  - `GrantTypes` to `*[]string` type
  - `ClientMetadata` to `*map[string]string` type
  - `Mobile` to dedicated `*ClientMobile` type
  - `Scopes` to `*map[string]string` type
- `ClientNativeSocialLogin` struct
  - `Apple` to dedicated `*ClientNativeSocialLoginSupportEnabled` type
  - `Facebook` to dedicated `*ClientNativeSocialLoginSupportEnabled` type
- `ClientGrant` struct field
  - `Scope` to `[]string` type
- `Connection` struct fields
  - `EnabledClients` to `*[]string` type
  - `Realms` to `*[]string` type
  - `Metadata` to `*map[string]string` type
- `ConnectionOptions` struct fields
 - `CustomScripts` to `*map[string]string` type
 - `Configuration` to `*map[string]string` type
- `ConnectionOptionsGoogleOAuth2` struct field
  - `AllowedAudiences` to `*[]string` type
- `ConnectionOptionsOIDC` struct field
  - `DomainAliases` to `*[]string` type
- `ConnectionOptionsOAuth2` struct field
  - `Scripts` to `*map[string]string` type
- `ConnectionOptionsAD` struct fields
  - `DomainAliases` to `*[]string` type
  - `IPs` to `*[]string` type
  - `IPs` added `omitempty` JSON struct tag
- `ConnectionOptionsAzureAD` struct field
  - `DomainAliases` to `*[]string` type
- `ConnectionOptionsADFS` struct field
  - `DomainAliases` to `*[]string` type
- `ConnectionOptionsSAML` struct field
  - `DomainAliases` to `*[]string` type
- `ConnectionOptionsGoogleApps` struct field
  - `DomainAliases` to `*[]string` type
- `Hook` struct field
  - `Dependencies` to `*map[string]string` type
- `LogStream` struct field
  - `Filters` to `*[]map[string]string` type
- `LogStreamSinkHTTP` struct field
  - `CustomHeaders` to `*[]map[string]string` type
- `Organization` struct field
  - `Metadata` to `*map[string]string` type
- `OrganizationBranding` struct field
  - `Colors` to `*map[string]string` type
- `ResourceServer` struct fields
  - `Scopes` to `*[]ResourceServerScope` type
  - `Options` to `*map[string]string` type
- `Tenant` struct fields
  - `AllowedLogoutURLs` to `*[]string` type
  - `SandboxVersionAvailable` to `*[]string` type
  - `EnabledLocales` to `*[]string` type

<a name="v0.10.1"></a>

## [v0.10.1](https://github.com/auth0/go-auth0/tree/v0.10.1) (2022-08-30)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.10.0...v0.10.1)

### Changed

- Enhance reliability of how we make requests and parse errors ([#108](https://github.com/auth0/go-auth0/pull/108)

<a name="v0.10.0"></a>

## [v0.10.0](https://github.com/auth0/go-auth0/tree/v0.10.0) (2022-08-23)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.9.3...v0.10.0)

### Added

- Added support for branding themes ([#103](https://github.com/auth0/go-auth0/pull/103), [#104](https://github.com/auth0/go-auth0/pull/104))
- Added ability to pass a custom audience when using client credentials ([#106](https://github.com/auth0/go-auth0/pull/106))

<a name="v0.9.3"></a>

## [v0.9.3](https://github.com/auth0/go-auth0/tree/v0.9.3) (2022-07-26)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.9.2...v0.9.3)

### Added

- Added support for checkpoint based pagination ([#97](https://github.com/auth0/go-auth0/pull/97))
- Added Multifactor field to User struct ([#97](https://github.com/auth0/go-auth0/pull/97))
- Added ClientID field to User struct ([#97](https://github.com/auth0/go-auth0/pull/97))

<a name="v0.9.2"></a>

## [v0.9.2](https://github.com/auth0/go-auth0/tree/v0.9.2) (2022-07-18)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.9.1...v0.9.2)

### Changed

- **[Breaking]** Change Metadata to `*map[string]interface{}` in Organization ([#95](https://github.com/auth0/go-auth0/pull/95))
- **[Breaking]** Change Metadata fields to `*map[string]interface{}` in User ([#96s](https://github.com/auth0/go-auth0/pull/95))

<a name="v0.9.1"></a>

## [v0.9.1](https://github.com/auth0/go-auth0/tree/v0.9.1) (2022-07-15)

[Full Changelog](https://github.com/auth0/go-auth0/compare/v0.9.0...v0.9.1)

### Added

- Added more connection strategies that default to OAuth2 ([#93](https://github.com/auth0/go-auth0/pull/93))

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
