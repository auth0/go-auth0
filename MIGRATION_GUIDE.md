# Auth0 Go SDK Migration Guide

This document covers the changes required to move between major versions of go-auth0.

- [Migrating from v2 to v3](#migrating-from-v2-to-v3)
- [Migrating from v1 to v2](#migrating-from-v1-to-v2)

---

# Migrating from v2 to v3

**Please review this section thoroughly to understand the changes required to migrate from go-auth0 v2 to go-auth0 v3.**

## Overview

v3 is a compatible evolution of the v2 client. The client initialization, package layout, and option pattern introduced in v2 all stay the same. The breaking changes are focused on a small number of request and response types where field types were tightened for correctness. Most applications will only need to update the specific call sites that touch the types listed below.

## v3 Breaking Changes

- [Connection Attribute Identifier Types](#connection-attribute-identifier-types)
- [Role Pagination Field Types](#role-pagination-field-types)
- [Phone Provider Protection Backoff Strategy Enum](#phone-provider-protection-backoff-strategy-enum)
- [Federated Connections Tokensets Removed](#federated-connections-tokensets-removed)
- [Federated Connections Access Tokens Removed](#federated-connections-access-tokens-removed)
- [Session Transfer Delegation Device Binding Enum](#session-transfer-delegation-device-binding-enum)

### Connection Attribute Identifier Types

The single `ConnectionAttributeIdentifier` type has been replaced with dedicated per-attribute identifier types. Each connection attribute now has its own identifier type, which allows the SDK to model the attribute-specific `default_method` values correctly.

| Attribute | v2 identifier type | v3 identifier type |
| --- | --- | --- |
| `EmailAttribute` | `ConnectionAttributeIdentifier` | `EmailAttributeIdentifier` |
| `PhoneAttribute` | `ConnectionAttributeIdentifier` | `PhoneAttributeIdentifier` |
| `UsernameAttribute` | `ConnectionAttributeIdentifier` | `UsernameAttributeIdentifier` |

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
attributes := &management.ConnectionAttributes{
    Email: &management.EmailAttribute{
        Identifier: &management.ConnectionAttributeIdentifier{
            Active: auth0.Bool(true),
        },
    },
    Username: &management.UsernameAttribute{
        Identifier: &management.ConnectionAttributeIdentifier{
            Active: auth0.Bool(true),
        },
    },
}
```

</td>
<td>

```go
attributes := &management.ConnectionAttributes{
    Email: &management.EmailAttribute{
        Identifier: &management.EmailAttributeIdentifier{
            Active: auth0.Bool(true),
        },
    },
    Username: &management.UsernameAttribute{
        Identifier: &management.UsernameAttributeIdentifier{
            Active: auth0.Bool(true),
        },
    },
}
```

</td>
</tr>
</table>

The `DefaultMethod` field is now typed per attribute as well. In v2 every attribute shared the `DefaultMethodEmailIdentifierEnum` type on its `DefaultMethod` field, so the phone attribute was incorrectly limited to the email enum values. In v3 each attribute uses the enum that matches it:

| Attribute identifier | v2 `DefaultMethod` type | v3 `DefaultMethod` type | v3 enum values |
| --- | --- | --- | --- |
| `EmailAttributeIdentifier` | `*DefaultMethodEmailIdentifierEnum` | `*DefaultMethodEmailIdentifierEnum` | `password`, `email_otp` |
| `PhoneAttributeIdentifier` | `*DefaultMethodEmailIdentifierEnum` | `*DefaultMethodPhoneNumberIdentifierEnum` | `password`, `phone_otp` |
| `UsernameAttributeIdentifier` | n/a | n/a (only `Active`) | n/a |

The new `DefaultMethodPhoneNumberIdentifierEnum` exposes `DefaultMethodPhoneNumberIdentifierEnumPassword` (`"password"`) and `DefaultMethodPhoneNumberIdentifierEnumPhoneOtp` (`"phone_otp"`), along with the usual `NewDefaultMethodPhoneNumberIdentifierEnumFromString` constructor and `Ptr` helper. If you set a default method on a phone identifier, switch from the email enum to the phone enum.

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
Identifier: &management.ConnectionAttributeIdentifier{
    Active:        auth0.Bool(true),
    DefaultMethod: management.DefaultMethodEmailIdentifierEnumPassword.Ptr(),
}
```

</td>
<td>

```go
Identifier: &management.PhoneAttributeIdentifier{
    Active:        auth0.Bool(true),
    DefaultMethod: management.DefaultMethodPhoneNumberIdentifierEnumPassword.Ptr(),
}
```

</td>
</tr>
</table>

### Role Pagination Field Types

On `ListRolesOffsetPaginatedResponseContent`, the pagination fields changed from optional pointers to non-pointer values, since the API always returns them.

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
Start *float64
Limit *float64
Total *float64

func (l *ListRolesOffsetPaginatedResponseContent) SetStart(start *float64)
func (l *ListRolesOffsetPaginatedResponseContent) SetLimit(limit *float64)
func (l *ListRolesOffsetPaginatedResponseContent) SetTotal(total *float64)
```

</td>
<td>

```go
Start float64
Limit float64
Total float64

func (l *ListRolesOffsetPaginatedResponseContent) SetStart(start float64)
func (l *ListRolesOffsetPaginatedResponseContent) SetLimit(limit float64)
func (l *ListRolesOffsetPaginatedResponseContent) SetTotal(total float64)
```

</td>
</tr>
</table>

If you read these fields directly, drop the pointer dereference. The `GetStart`, `GetLimit`, and `GetTotal` accessors continue to return `float64`, so code using the accessors needs no change.

### Phone Provider Protection Backoff Strategy Enum

On `PhoneProviderProtectionBackoffStrategyEnum`, the `None` value (`"none"`) has been removed and replaced with `Default` (`"default"`). The `Exponential` value is unchanged.

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
strategy := management.PhoneProviderProtectionBackoffStrategyEnumNone // "none"
```

</td>
<td>

```go
strategy := management.PhoneProviderProtectionBackoffStrategyEnumDefault // "default"
```

</td>
</tr>
</table>

Replace any use of `PhoneProviderProtectionBackoffStrategyEnumNone` with `PhoneProviderProtectionBackoffStrategyEnumDefault`. `NewPhoneProviderProtectionBackoffStrategyEnumFromString` no longer accepts `"none"`.

### Federated Connections Tokensets Removed

The `Users.FederatedConnectionsTokensets` sub-client has been removed, along with its `List` and `Delete` methods and the `FederatedConnectionTokenSet` type. The underlying `/users/{id}/federated-connections-tokensets` endpoints are no longer part of the SDK.

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
tokensets, err := client.Users.FederatedConnectionsTokensets.List(context.TODO(), userID)
err = client.Users.FederatedConnectionsTokensets.Delete(context.TODO(), userID, tokensetID)
```

</td>
<td>

```go
// No replacement. Remove any calls to Users.FederatedConnectionsTokensets.
```

</td>
</tr>
</table>

### Federated Connections Access Tokens Removed

The `ConnectionFederatedConnectionsAccessTokens` type has been removed, along with the `FederatedConnectionsAccessTokens` field (and its getter and setter) on the connection options types, including `ConnectionPropertiesOptions` (used by `CreateConnectionRequestContent`) and `UpdateConnectionOptions` (used by `UpdateConnectionRequestContent`).

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
req := &management.CreateConnectionRequestContent{
    Name:     "my-connection",
    Strategy: management.ConnectionIdentityProviderEnumOidc,
    Options: &management.ConnectionPropertiesOptions{
        ImportMode: auth0.Bool(false),
        FederatedConnectionsAccessTokens: &management.ConnectionFederatedConnectionsAccessTokens{
            Active: auth0.Bool(true),
        },
    },
}
```

</td>
<td>

```go
// Drop the FederatedConnectionsAccessTokens field; the rest is unchanged.
req := &management.CreateConnectionRequestContent{
    Name:     "my-connection",
    Strategy: management.ConnectionIdentityProviderEnumOidc,
    Options: &management.ConnectionPropertiesOptions{
        ImportMode: auth0.Bool(false),
    },
}
```

</td>
</tr>
</table>

### Session Transfer Delegation Device Binding Enum

On `ClientSessionTransferDelegationDeviceBindingEnum`, the `Asn` value (`"asn"`) has been removed. The only supported value is now `IP` (`"ip"`), which enforces device binding by IP, meaning the Session Transfer Token must be consumed from the same IP as the issuer.

<table>
<tr>
<th>v2</th>
<th>v3</th>
</tr>
<tr>
<td>

```go
binding := management.ClientSessionTransferDelegationDeviceBindingEnumAsn // "asn"
```

</td>
<td>

```go
binding := management.ClientSessionTransferDelegationDeviceBindingEnumIP // "ip"
```

</td>
</tr>
</table>

Replace any use of `ClientSessionTransferDelegationDeviceBindingEnumAsn` with `ClientSessionTransferDelegationDeviceBindingEnumIP`. `NewClientSessionTransferDelegationDeviceBindingEnumFromString` no longer accepts `"asn"`. This applies only to the delegation (impersonation) enum; the unrelated `ClientSessionTransferDeviceBindingEnum` still supports `ip`, `asn`, and `none`.

## v3 Migration Steps

1. Update the dependency to the v3 major and update your import paths if you pin to a specific version.
2. Search your codebase for `ConnectionAttributeIdentifier` and replace each occurrence with the identifier type that matches the enclosing attribute (`EmailAttributeIdentifier`, `PhoneAttributeIdentifier`, or `UsernameAttributeIdentifier`).
3. Remove pointer dereferences on the `Start`, `Limit`, and `Total` fields of `ListRolesOffsetPaginatedResponseContent` if you read them directly.
4. Replace `PhoneProviderProtectionBackoffStrategyEnumNone` with `PhoneProviderProtectionBackoffStrategyEnumDefault`.
5. Remove any calls to `Users.FederatedConnectionsTokensets` and any use of the `FederatedConnectionTokenSet` type.
6. Remove the `FederatedConnectionsAccessTokens` field from connection options and drop any use of the `ConnectionFederatedConnectionsAccessTokens` type.
7. Replace `ClientSessionTransferDelegationDeviceBindingEnumAsn` with `ClientSessionTransferDelegationDeviceBindingEnumIP`.
8. Build and run your tests to catch any remaining type mismatches.

---

# Migrating from v1 to v2

**Please review this section thoroughly to understand the changes required to migrate from go-auth0 v1 to go-auth0 v2**

## Table of Contents

- [Overview](#overview)
- [Breaking Changes](#breaking-changes)
  - [Module Name Change](#module-name-change)
  - [Package Structure Reorganization](#package-structure-reorganization)
  - [Generated Code and Type Safety](#generated-code-and-type-safety)
  - [Management Client Initialization](#management-client-initialization)
  - [API Method Signatures](#api-method-signatures)
  - [Request and Response Types](#request-and-response-types)
  - [Error Handling](#error-handling)
- [Migration Steps](#migration-steps)
- [Examples](#examples)
  - [Client Management](#client-management)
  - [User Management](#user-management)
  - [Connection Management](#connection-management)

## Overview

The Auth0 Go SDK v2 represents a complete rewrite of the SDK with the following major improvements:

- **Generated from OpenAPI**: v2 is generated from Auth0's OpenAPI specifications, ensuring consistency and accuracy
- **Improved Type Safety**: Strongly typed request/response structures with proper enums and validation
- **Better Organization**: Management APIs are organized into dedicated subpackages
- **Enhanced Developer Experience**: Better IntelliSense support and clear method signatures
- **Future-Proof**: Easier to maintain and update as Auth0's API evolves

## Breaking Changes

### Module Name Change

The import path has changed from `github.com/auth0/go-auth0` to `github.com/auth0/go-auth0/v2`.

<table>
<tr>
<th>v1 (go-auth0)</th>
<th>v2 (go-auth0)</th>
</tr>
<tr>
<td>

```go
import (
    "github.com/auth0/go-auth0"
    "github.com/auth0/go-auth0/management"
)
```

</td>
<td>

```go
import (
    "github.com/auth0/go-auth0/v2"
    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
)
```

</td>
</tr>
</table>

### Package Structure Reorganization

v2 organizes management APIs into dedicated subpackages, replacing the monolithic structure of v1.

<table>
<tr>
<th>v1 Structure</th>
<th>v2 Structure</th>
</tr>
<tr>
<td>

```
management/
├── client.go
├── user.go
├── connection.go
├── role.go
└── ... (all in root)
```

</td>
<td>

```
management/
├── clients/
│   └── client/
├── users/
│   └── client/
├── connections/
│   └── client/
├── roles/
│   └── client/
└── ... (organized by feature)
```

</td>
</tr>
</table>

### Generated Code and Type Safety

v2 uses generated code with strictly typed structures and enums, replacing the loosely typed approach of v1.

<table>
<tr>
<th>v1 (Loose Typing)</th>
<th>v2 (Strict Typing)</th>
</tr>
<tr>
<td>

```go
client := &management.Client{
    Name:    auth0.String("My App"),
    AppType: auth0.String("spa"),
    // Many fields are *string
}
```

</td>
<td>

```go
client := &management.CreateClientRequestContent{
    Name:    "My App",
    AppType: &management.ClientAppTypeEnumSpa,
    // Proper enums and required fields
}
```

</td>
</tr>
</table>

### Management Client Initialization

The way you initialize and access management clients has changed significantly.

<table>
<tr>
<th>v1 Initialization</th>
<th>v2 Initialization</th>
</tr>
<tr>
<td>

```go
import "github.com/auth0/go-auth0/management"

mgmt, err := management.New(
    domain,
    management.WithClientCredentials(
        context.Background(),
        clientID,
        clientSecret,
    ),
)

// Access managers directly
client, err := mgmt.Client.Create(ctx, clientData)
```

</td>
<td>

```go
import (
    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
    "github.com/auth0/go-auth0/v2/management/option"
)

mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentials(
        context.Background(),
        clientID,
        clientSecret,
    ),
)
if err != nil {
    // handle error
}

// Access through specific client managers
client, err := mgmt.Clients.Create(ctx, clientData)
```

</td>
</tr>
</table>

### API Method Signatures

Method signatures have been updated with more specific types and clearer parameter structures.

<table>
<tr>
<th>v1 Method Signature</th>
<th>v2 Method Signature</th>
</tr>
<tr>
<td>

```go
// Create a client
func (m *ClientManager) Create(
    ctx context.Context,
    c *Client,
    opts ...RequestOption,
) error

// List clients
func (m *ClientManager) List(
    ctx context.Context,
    opts ...RequestOption,
) (*ClientList, error)
```

</td>
<td>

```go
// Create a client
func (c *Client) Create(
    ctx context.Context,
    request *management.CreateClientRequestContent,
    opts ...option.RequestOption,
) (*management.CreateClientResponseContent, error)

// List clients
func (c *Client) List(
    ctx context.Context,
    request *management.ListClientsRequest,
    opts ...option.RequestOption,
) (*management.ListClientsResponse, error)
```

</td>
</tr>
</table>

### Request and Response Types

v2 introduces specific request and response types for each operation.

<table>
<tr>
<th>v1 Approach</th>
<th>v2 Approach</th>
</tr>
<tr>
<td>

```go
// Single type for all operations
client := &management.Client{
    Name: auth0.String("My App"),
    // ... other fields
}

err := mgmt.Client.Create(ctx, client)
```

</td>
<td>

```go
// Separate request and response types
createRequest := &management.CreateClientRequestContent{
    Name: "My App",
    // ... other fields
}

response, err := mgmt.Clients.Create(ctx, createRequest)
// response is *management.CreateClientResponseContent
```

</td>
</tr>
</table>

## Migration Steps

### Step 1: Update Dependencies

Update your `go.mod` file:

```bash
# Remove old dependency
go mod edit -droprequire github.com/auth0/go-auth0

# Add new dependency
go get github.com/auth0/go-auth0/v2
```

### Step 2: Update Imports

Replace all imports:

```go
// Old
import "github.com/auth0/go-auth0/management"

// New
import (
    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
    "github.com/auth0/go-auth0/v2/option"
)
```

### Step 3: Update Client Initialization

```go
// Old
mgmt, err := management.New(
    domain,
    management.WithClientCredentials(ctx, clientID, clientSecret),
)

// New
mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentials(
        context.Background(),
        clientID,
        clientSecret,
    ),
)
```

### Step 4: Update API Calls

Update your API calls to use the new structure and types:

```go
// Old
client := &management.Client{
    Name: auth0.String("My App"),
}
err := mgmt.Client.Create(ctx, client)

// New
request := &management.CreateClientRequestContent{
    Name: "My App",
}
response, err := mgmt.Clients.Create(ctx, request)
```

## Examples

### Client Management

<table>
<tr>
<th>v1 Example</th>
<th>v2 Example</th>
</tr>
<tr>
<td>

```go
package main

import (
    "context"
    "github.com/auth0/go-auth0"
    "github.com/auth0/go-auth0/management"
)

func main() {
    mgmt, err := management.New(
        "your-domain.auth0.com",
        management.WithClientCredentials(
            context.Background(),
            "client-id",
            "client-secret",
        ),
    )
    if err != nil {
        panic(err)
    }

    // Create client
    client := &management.Client{
        Name:    auth0.String("My App"),
        AppType: auth0.String("spa"),
        Callbacks: &[]string{
            "http://localhost:3000/callback",
        },
    }

    err = mgmt.Client.Create(context.Background(), client)
    if err != nil {
        panic(err)
    }

    // List clients
    clients, err := mgmt.Client.List(context.Background())
    if err != nil {
        panic(err)
    }
}
```

</td>
<td>

```go
package main

import (
    "context"
    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
    "github.com/auth0/go-auth0/v2/option"
)

func main() {
    mgmt, err := client.New(
        "your-domain.auth0.com",
        option.WithClientCredentials(
            context.Background(),
            "client-id",
            "client-secret",
        ),
    )
    if err != nil {
        panic(err)
    }

    // Create client
    request := &management.CreateClientRequestContent{
        Name:    "My App",
        AppType: &management.ClientAppTypeEnumSpa,
        Callbacks: []string{
            "http://localhost:3000/callback",
        },
    }

    response, err := mgmt.Clients.Create(context.Background(), request)
    if err != nil {
        panic(err)
    }

    // List clients
    listRequest := &management.ListClientsRequest{}
    clients, err := mgmt.Clients.List(context.Background(), listRequest)
    if err != nil {
        panic(err)
    }
}
```

</td>
</tr>
</table>

### User Management

<table>
<tr>
<th>v1 Example</th>
<th>v2 Example</th>
</tr>
<tr>
<td>

```go
// Create user
user := &management.User{
    Email:      auth0.String("user@example.com"),
    Connection: auth0.String("Username-Password-Authentication"),
    Password:   auth0.String("password123"),
}

err := mgmt.User.Create(ctx, user)

// List users
users, err := mgmt.User.List(ctx)
```

</td>
<td>

```go
// Create user
request := &management.CreateUserRequestContent{
    Email:      "user@example.com",
    Connection: "Username-Password-Authentication",
    Password:   management.String("password123"),
}

response, err := mgmt.Users.Create(ctx, request)

// List users
listRequest := &management.ListUsersRequest{}
users, err := mgmt.Users.List(ctx, listRequest)
```

</td>
</tr>
</table>

### Connection Management

<table>
<tr>
<th>v1 Example</th>
<th>v2 Example</th>
</tr>
<tr>
<td>

```go
// Create connection
connection := &management.Connection{
    Name:     auth0.String("my-database"),
    Strategy: auth0.String("auth0"),
    Options: &management.ConnectionOptions{
        BruteForceProtection: auth0.Bool(true),
    },
}

err := mgmt.Connection.Create(ctx, connection)
```

</td>
<td>

```go
// Create connection
request := &management.CreateConnectionRequestContent{
    Name:     "my-database",
    Strategy: management.ConnectionStrategyEnumAuth0,
    Options: &management.ConnectionOptions{
        BruteForceProtection: management.Bool(true),
    },
}

response, err := mgmt.Connections.Create(ctx, request)
```

</td>
</tr>
</table>

## Authentication Package

The authentication package remains largely unchanged between v1 and v2. You can continue using it with minimal modifications:

```go
// Both v1 and v2 support similar authentication patterns
import "github.com/auth0/go-auth0/v2/authentication"

auth, err := authentication.New(
    context.Background(),
    "your-domain.auth0.com",
    authentication.WithClientID("client-id"),
    authentication.WithClientSecret("client-secret"),
)
```

## Additional Notes

1. **Context Requirement**: All v2 methods require a `context.Context` as the first parameter
2. **Option Pattern**: v2 uses a consistent option pattern for configuration
3. **Error Handling**: Error types may have changed - review your error handling logic
4. **Pagination**: List methods now return specific response types with pagination information
5. **Type Safety**: Take advantage of the improved type safety by using the provided enums and constants

## Troubleshooting

### Common Issues

1. **Import Errors**: Make sure to update all import statements to use the new module path
2. **Type Mismatches**: Use the new strongly-typed structures instead of generic map types
3. **Method Not Found**: Check the new package structure - methods may have moved to subpackages
4. **Initialization Errors**: Use the new client initialization pattern with options

### Getting Help

- Check the [API Reference Documentation](reference.md)
- Review the examples in the repository
- Open an issue on GitHub for specific migration problems

This migration guide covers the major changes needed to upgrade from go-auth0 v1 to go-auth0 v2. While the changes are significant, the improved type safety and organization make the SDK more robust and easier to use.
