# Go-Auth0 v1 Migration Guide

With the v1 release of Go-Auth0, we've

**Please review this guide thoroughly to understand the changes required to migrate your usage of go-auth0 to v1**

## Public API Changes

### Improve Typing Of Client Addons

The `Addons` property of a `Client` was previously simply typed as `map[string]interface{}`, now it is explicitly typed to allow better support of addons.

However, with this there are now only two explicitly supported Addon types, `SAML2` and `WS-FED` which are the 2 addons supported in the Management UI. If you require a specific Addon, please open a [feature request](https://github.com/auth0/go-auth0/issues/new?assignees=&labels=feature&projects=&template=feature_request.yml)