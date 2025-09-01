# Auth0 Go SDK v1 to v2 Migration Guide

**Please review this guide thoroughly to understand the changes required to migrate from go-auth0 v1 to go-auth0 v2**

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
