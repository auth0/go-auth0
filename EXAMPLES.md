# Examples

This guide provides comprehensive examples for using the Auth0 Go SDK v2, covering authentication, management operations, pagination, and advanced usage patterns.

## Table of Contents

- [Client Credentials Authentication](#client-credentials-authentication)
- [Request Options](#request-options)
- [Pagination](#pagination)
  - [Page-based Pagination](#page-based-pagination)
  - [Checkpoint Pagination](#checkpoint-pagination)
- [Custom Request Handling](#custom-request-handling)
- [Error Handling](#error-handling)
- [Complete Examples](#complete-examples)

## Client Credentials Authentication

The SDK v2 provides several ways to authenticate with Auth0 using OAuth client credentials through the options pattern.

### Management API Initialization

```go
import (
    "context"
    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
    "github.com/auth0/go-auth0/v2/management/option"
)

// Standard client credentials with client secret
mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentials(
        context.Background(),
        "YOUR_CLIENT_ID",
        "YOUR_CLIENT_SECRET",
    ),
)

// Client credentials with custom audience
mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentialsAndAudience(
        context.Background(),
        "YOUR_CLIENT_ID",
        "YOUR_CLIENT_SECRET",
        "https://custom-api.example.com",
    ),
)

// Client credentials with Private Key JWT
mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentialsPrivateKeyJwt(
        context.Background(),
        "YOUR_CLIENT_ID",
        privateKeyPEM, // PEM-encoded private key
        "RS256",
    ),
)

// Private Key JWT with custom audience
mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentialsPrivateKeyJwtAndAudience(
        context.Background(),
        "YOUR_CLIENT_ID",
        privateKeyPEM,
        "RS256",
        "https://custom-api.example.com",
    ),
)

// Using a pre-acquired static token
mgmt, err := client.New(
    "your-tenant.auth0.com",
    option.WithToken("YOUR_ACCESS_TOKEN"),
)
```

### Authentication API Initialization

```go
import (
    "context"
    "github.com/auth0/go-auth0/v2/authentication"
)

// With client secret authentication
auth, err := authentication.New(
    context.Background(),
    "your-tenant.auth0.com",
    authentication.WithClientID("YOUR_CLIENT_ID"),
    authentication.WithClientSecret("YOUR_CLIENT_SECRET"),
)

// With private key JWT authentication
auth, err := authentication.New(
    context.Background(),
    "your-tenant.auth0.com",
    authentication.WithClientID("YOUR_CLIENT_ID"),
    authentication.WithClientAssertion(privateKeyPEM, "RS256"),
)
```

## Request Options

v2 provides fine-grained configuration through request options that can be applied to individual requests.

### Basic Request Options

```go
import (
    "net/http"
    "net/url"
)

// Add custom headers
clients, err := mgmt.Clients.List(
    ctx,
    &management.ListClientsRequestParameters{},
    option.WithHTTPHeader(http.Header{
        "X-Custom-Header": []string{"custom-value"},
        "X-Request-ID":    []string{"req-123"},
    }),
)

// Add query parameters
queryParams := url.Values{}
queryParams.Set("include_fields", "client_id,name,app_type")
queryParams.Set("exclude_fields", "client_secret")

clients, err := mgmt.Clients.List(
    ctx,
    &management.ListClientsRequestParameters{},
    option.WithQueryParameters(queryParams),
)

// Add body properties (for requests with body)
bodyProps := map[string]interface{}{
    "custom_field": "custom_value",
}

response, err := mgmt.Clients.Create(
    ctx,
    createRequest,
    option.WithBodyProperties(bodyProps),
)

// Configure retry attempts
users, err := mgmt.Users.List(
    ctx,
    &management.ListUsersRequestParameters{},
    option.WithMaxAttempts(3),
)
```

### Common Request Parameter Patterns

```go
// List clients with filtering and pagination
listRequest := &management.ListClientsRequestParameters{
    AppType:    management.String("spa"),
    Page:       management.Int(0),
    PerPage:    management.Int(25),
    Fields:     management.String("client_id,name,app_type"),
    IncludeFields: management.Bool(true),
}

clients, err := mgmt.Clients.List(ctx, listRequest)

// List users with search and sorting
userListRequest := &management.ListUsersRequestParameters{
    Search:     management.String("email:\"*@example.com\""),
    Sort:       management.String("created_at:1"),
    Page:       management.Int(0),
    PerPage:    management.Int(50),
}

users, err := mgmt.Users.List(ctx, userListRequest)

// List connections with specific fields
connectionListRequest := &management.ListConnectionsRequestParameters{
    Strategy:   management.String("auth0"),
    Fields:     management.String("id,name,strategy,enabled_clients"),
    IncludeFields: management.Bool(true),
}

connections, err := mgmt.Connections.List(ctx, connectionListRequest)
```

## Pagination

The SDK v2 uses a `Page` type for pagination with built-in iterator support.

### Page-based Pagination

```go
// Manual pagination using GetNextPage
ctx := context.Background()
page := 0

for {
    listRequest := &management.ListClientsRequestParameters{
        Page:    management.Int(page),
        PerPage: management.Int(25),
    }

    clientsPage, err := mgmt.Clients.List(ctx, listRequest)
    if err != nil {
        return err
    }

    // Process current page results
    for _, client := range clientsPage.Results {
        fmt.Printf("Client: %s (%s)\n",
            *client.GetName(),
            *client.GetClientID(),
        )
    }

    // Try to get next page
    nextPage, err := clientsPage.GetNextPage(ctx)
    if err != nil {
        if errors.Is(err, core.ErrNoPages) {
            // No more pages, we're done
            break
        }
        return err
    }

    // Continue with next page
    clientsPage = nextPage
    page++
}
```

### Using Page Iterator

```go
// Using the built-in iterator for automatic pagination
listRequest := &management.ListClientsRequestParameters{
    PerPage: management.Int(50),
}

clientsPage, err := mgmt.Clients.List(ctx, listRequest)
if err != nil {
    return err
}

iterator := clientsPage.Iterator()
for iterator.Next(ctx) {
    client := iterator.Current()
    fmt.Printf("Client: %s (%s)\n",
        *client.GetName(),
        *client.GetClientID(),
    )
}

if iterator.Err() != nil {
    return iterator.Err()
}
```

### Checkpoint Pagination (For Logs and Organizations)

For APIs that support checkpoint pagination (like logs, organizations), you can use either iterators or manual pagination:

```go
// Option 1: Using iterator (recommended)
logListRequest := &management.ListLogsRequestParameters{
    Take: management.Int(100),
}

logsPage, err := mgmt.Logs.List(ctx, logListRequest)
if err != nil {
    return err
}

iterator := logsPage.Iterator()
for iterator.Next(ctx) {
    log := iterator.Current()
    fmt.Printf("Log: %s - %s\n",
        *log.GetID(),
        *log.GetType(),
    )
}

if iterator.Err() != nil {
    return iterator.Err()
}
```

```go
// Option 2: Manual checkpoint pagination (for advanced use cases)
var fromLogID *string

for {
    logListRequest := &management.ListLogsRequestParameters{
        Take: management.Int(100),
    }

    if fromLogID != nil {
        logListRequest.From = fromLogID
    }

    logsPage, err := mgmt.Logs.List(ctx, logListRequest)
    if err != nil {
        return err
    }

    // Process logs
    for _, log := range logsPage.Results {
        fmt.Printf("Log: %s - %s\n",
            *log.GetID(),
            *log.GetType(),
        )
    }

    // Check if we have more logs
    if len(logsPage.Results) == 0 {
        break // No more logs
    }

    // Use the ID of the last log as the checkpoint for the next request
    lastLog := logsPage.Results[len(logsPage.Results)-1]
    fromLogID = lastLog.ID
}
```

## Custom Request Handling

For scenarios requiring custom request handling or working with fields not included in the generated types:

### Using WithBodyProperties for Additional Fields

The `WithBodyProperties` option allows you to send additional fields in request bodies that aren't part of the generated SDK types. This is useful for:

- Custom metadata fields
- Advanced configuration options
- Fields that may be added to the Auth0 API but not yet in the SDK

```go
// Example: Creating a client with additional custom properties
customBodyProps := map[string]interface{}{
    "custom_domain_verified": true,
    "integration_metadata": map[string]interface{}{
        "source":      "terraform",
        "environment": "production",
        "team":        "platform",
    },
    "advanced_settings": map[string]interface{}{
        "custom_login_page":         true,
        "universal_login_branding":  true,
    },
}

createRequest := &management.CreateClientRequestContent{
    Name:    "My Custom App",
    AppType: &management.ClientAppTypeEnumSpa,
    Callbacks: []string{"https://myapp.com/callback"},
}

client, err := mgmt.Clients.Create(
    ctx,
    createRequest,
    option.WithBodyProperties(customBodyProps),
)
```

### Working with Nullable Fields

```go
// Use WithBodyProperties to explicitly set fields to null
_, err := mgmt.Clients.Update(
    ctx,
    clientID,
    &management.UpdateClientRequestContent{
        Name: management.String("Updated App Name"),
    },
    option.WithBodyProperties(map[string]interface{}{
        "description":     nil, // This will set description to null
        "logo_uri":        "https://example.com/new-logo.png",
        "custom_metadata": nil, // This will set custom_metadata to null
    }),
)
```

### Working with ExtraProperties

Some SDK types include an `ExtraProperties` field that can be used in both directions:

1. **In Requests**: Set extra properties to send additional fields to Auth0
2. **In Responses**: Capture additional fields returned by Auth0 that aren't in the standard schema

#### Setting ExtraProperties in Requests

```go
// Example: Creating a connection with extra properties in the request struct
createRequest := &management.CreateConnectionRequestContent{
    Name:     "My Custom DB",
    Strategy: management.ConnectionIdentityProviderEnumAuth0,
    Options: &management.ConnectionPropertiesOptions{
        ImportMode: management.Bool(true),
        // Set extra properties directly on the struct
        ExtraProperties: map[string]interface{}{
            "custom_timeout":     30,
            "retry_policy":       "exponential_backoff",
            "webhook_endpoints": []string{
                "https://api.example.com/auth-webhook",
                "https://backup.example.com/auth-webhook",
            },
            "advanced_config": map[string]interface{}{
                "enable_debug_logs": true,
                "api_version":       "v2.1",
                "custom_headers": map[string]string{
                    "X-Custom-Source": "auth0-sdk",
                    "X-Environment":   "production",
                },
            },
        },
    },
}

newConn, err := mgmt.Connections.Create(ctx, createRequest)
if err != nil {
    return err
}
```

```go
// Example: Updating connection with extra properties
updateRequest := &management.UpdateConnectionRequestContent{
    Options: &management.ConnectionPropertiesOptions{
        ExtraProperties: map[string]interface{}{
            "maintenance_mode":    false,
            "last_updated_by":     "admin-user-123",
            "migration_settings": map[string]interface{}{
                "batch_size":     1000,
                "rate_limit_ms":  100,
                "error_handling": "continue",
            },
        },
    },
}

updatedConn, err := mgmt.Connections.Update(ctx, "con_123", updateRequest)
```

## Error Handling

```go
import (
    "errors"
    "fmt"
    "net/http"
)

// Basic error handling
clients, err := mgmt.Clients.List(ctx, &management.ListClientsRequestParameters{})
if err != nil {
    // Check for specific HTTP status codes if needed
    fmt.Printf("Error listing clients: %v\n", err)
    return err
}

// Handle pagination errors
clientsPage, err := mgmt.Clients.List(ctx, &management.ListClientsRequestParameters{})
if err != nil {
    return err
}

for {
    // Process current page
    for _, client := range clientsPage.Results {
        fmt.Printf("Client: %s\n", *client.GetName())
    }

    // Try to get next page
    nextPage, err := clientsPage.GetNextPage(ctx)
    if err != nil {
        if errors.Is(err, core.ErrNoPages) {
            // This is expected when we reach the end
            break
        }
        // This is an actual error
        return fmt.Errorf("error getting next page: %w", err)
    }

    clientsPage = nextPage
}
```

## Complete Examples

### Client Management Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
    "github.com/auth0/go-auth0/v2/management/option"
)

func main() {
    // Initialize management client
    mgmt, err := client.New(
        "your-tenant.auth0.com",
        option.WithClientCredentials(
            context.Background(),
            "YOUR_CLIENT_ID",
            "YOUR_CLIENT_SECRET",
        ),
    )
    if err != nil {
        log.Fatalf("Error creating management client: %v", err)
    }

    ctx := context.Background()

    // Create a new client application
    createRequest := &management.CreateClientRequestContent{
        Name:    "My Go SDK v2 App",
        AppType: &management.ClientAppTypeEnumSpa,
        Callbacks: []string{
            "http://localhost:3000/callback",
            "https://myapp.com/callback",
        },
        AllowedOrigins: []string{
            "http://localhost:3000",
            "https://myapp.com",
        },
        OidcConformant: management.Bool(true),
        JwtConfiguration: &management.ClientJwtConfiguration{
            Algorithm:         management.String("RS256"),
            LifetimeInSeconds: management.Int(3600),
        },
    }

    clientResponse, err := mgmt.Clients.Create(ctx, createRequest)
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    fmt.Printf("Created client: %s (ID: %s)\n",
        *clientResponse.GetName(),
        *clientResponse.GetClientID(),
    )

    // List all SPA clients
    listRequest := &management.ListClientsRequestParameters{
        AppType: management.String("spa"),
        Fields:  management.String("client_id,name,app_type,callbacks"),
        IncludeFields: management.Bool(true),
        PerPage: management.Int(10),
    }

    clientsPage, err := mgmt.Clients.List(ctx, listRequest)
    if err != nil {
        log.Fatalf("Error listing clients: %v", err)
    }

    fmt.Println("\nSPA Clients:")
    iterator := clientsPage.Iterator()
    for iterator.Next(ctx) {
        client := iterator.Current()
        fmt.Printf("- %s (%s)\n", *client.GetName(), *client.GetClientID())
    }

    if iterator.Err() != nil {
        log.Fatalf("Error iterating clients: %v", iterator.Err())
    }

    // Update the client
    updateRequest := &management.UpdateClientRequestContent{
        Description: management.String("Updated description for my app"),
        LogoURI:     management.String("https://example.com/logo.png"),
    }

    updatedClient, err := mgmt.Clients.Update(
        ctx,
        *clientResponse.GetClientID(),
        updateRequest,
    )
    if err != nil {
        log.Fatalf("Error updating client: %v", err)
    }

    fmt.Printf("Updated client description: %s\n",
        *updatedClient.GetDescription(),
    )
}
```

### User Management Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/auth0/go-auth0/v2/management"
    "github.com/auth0/go-auth0/v2/management/client"
    "github.com/auth0/go-auth0/v2/management/option"
)

func main() {
    mgmt, err := client.New(
        "your-tenant.auth0.com",
        option.WithClientCredentials(
            context.Background(),
            "YOUR_CLIENT_ID",
            "YOUR_CLIENT_SECRET",
        ),
    )
    if err != nil {
        log.Fatalf("Error creating management client: %v", err)
    }

    ctx := context.Background()

    // Create a new user
    createUserRequest := &management.CreateUserRequestContent{
        Email:      "newuser@example.com",
        Connection: "Username-Password-Authentication",
        Password:   management.String("TempPassword123!"),
        UserMetadata: map[string]interface{}{
            "preference": "email",
            "plan":       "premium",
        },
        AppMetadata: map[string]interface{}{
            "role":  "user",
            "scope": "read:profile",
        },
    }

    userResponse, err := mgmt.Users.Create(ctx, createUserRequest)
    if err != nil {
        log.Fatalf("Error creating user: %v", err)
    }

    fmt.Printf("Created user: %s (ID: %s)\n",
        *userResponse.GetEmail(),
        *userResponse.GetUserID(),
    )

    // Search for users by email domain
    searchRequest := &management.ListUsersRequestParameters{
        Search:  management.String("email:\"*@example.com\""),
        Sort:    management.String("created_at:1"),
        PerPage: management.Int(25),
    }

    usersPage, err := mgmt.Users.List(ctx, searchRequest)
    if err != nil {
        log.Fatalf("Error searching users: %v", err)
    }

    fmt.Println("\nUsers with @example.com domain:")
    for _, user := range usersPage.Results {
        fmt.Printf("- %s (Created: %s)\n",
            *user.GetEmail(),
            *user.GetCreatedAt(),
        )
    }

    // Update user metadata
    updateUserRequest := &management.UpdateUserRequestContent{
        UserMetadata: map[string]interface{}{
            "preference":    "sms",
            "plan":          "premium",
            "last_updated":  "2024-01-01",
        },
    }

    updatedUser, err := mgmt.Users.Update(
        ctx,
        *userResponse.GetUserID(),
        updateUserRequest,
    )
    if err != nil {
        log.Fatalf("Error updating user: %v", err)
    }

    fmt.Printf("Updated user metadata: %v\n",
        updatedUser.UserMetadata,
    )
}
```

These examples demonstrate the key patterns for using the Auth0 Go SDK v2 effectively. The SDK's type-safe approach and built-in pagination support make it easier to work with Auth0's Management API while providing flexibility for advanced use cases.
