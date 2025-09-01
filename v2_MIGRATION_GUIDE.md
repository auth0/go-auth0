# V2 Migration Guide

A guide to migrating the Auth0 Go SDK from `1.x` to `2.x`.

- [Overall changes](#overall-changes)
    - [Accepted Node versions](#accepted-node-versions)
    - [Authentication API](#authentication-api)
    - [Management API](#management-api)
- [Specific changes to the Management API](#specific-changes-to-the-management-api)
    - [Method name changes](#method-name-changes)
    - [Pagination and Response Changes](#pagination-and-response-changes)
        - [Accessing Response Data](#accessing-response-data)
        - [Migrating from V1 to V2 Pagination](#migrating-from-v1-to-v2-pagination)
        - [Non-Paginated Responses (Create, Update, etc.)](#non-paginated-responses-create-update-etc)
        - [Advanced Pagination](#advanced-pagination)
    - [Type name changes](#type-name-changes)
    - [Import users takes a Reader](#import-users-takes-a-reader)

## Overall changes

### Accepted GO versions

This SDK actively supports Go versions `1.23` or higher. While you may find that older versions of Go may work, we will not actively test and fix compatibility issues with these versions.

### Authentication API

This major version change does not affect the Authentication API. Any code written for the Authentication API in the `1.x` version should work in the `2.x` version.

### Management API

V2 introduces significant improvements to the Management API SDK by migrating to [Fern](https://github.com/fern-api/fern) as our code generation tool. While the SDK was previously generated, v2 benefits from Fern's enhanced capabilities including better resource grouping, sub-client organization, and customization options. Additionally, v2 leverages a substantially improved and optimized OpenAPI specification that provides more accurate type definitions and better API structure representation. This combination introduces a number of benefits and changes to the SDK, which we'll outline below.

## Specific changes to the Management API

### Method name changes

V2 introduces a more consistent and intuitive API structure. We moved subresources into sub-clients and use consistent method naming: `list`, `create`, `update`, `delete`, `set`, `get`, and various actions like `test`, `deploy`, `reset`, etc.

Most method names are now in tighter lock-step with their respective endpoints (see the [Management API documentation](https://auth0.com/docs/api/management/v2) for information on all available endpoints). For instance, to get action versions, the v1 method `auth0API.Action.Versions()` becomes `actions.versions.list()` in v2:

```go
import (
    "context"
    "log"

    auth0 "github.com/fern-demo/auth0-go-sdk/management"
    "github.com/fern-demo/auth0-go-sdk/management/actions"
    "github.com/fern-demo/auth0-go-sdk/management/client"
    "github.com/fern-demo/auth0-go-sdk/management/option"
)

ctx := context.TODO()
m, err := client.New(
    "your-tenant.auth0.com",
    option.WithClientCredentials(
        ctx,
        "YOUR_CLIENT_ID",
        "YOUR_CLIENT_SECRET"
    ),
)
if err != nil {
    log.Fatalf("❌ Failed to create client: %v", err)
}

// V2: Subresources moved to sub-clients with consistent naming
versions, err := m.Actions.Versions.List(
    ctx,
    "action_id",
    &actions.ListActionVersionsRequestParameters{}
)
if err != nil {
    log.Fatalf("❌ Failed to list action versions: %v", err)
}

// V1: Direct method on main resource
// versions, err := auth0API.Action.Versions(
//     context.TODO(),
//     client.GetClientID()
// )
// if err != nil {
//    log.Fatalf("❌ Failed to list action versions: %v", err)
// }
```

In the common situation where there are two similar `GET` methods, one that "gets one" and another that "gets all" of a particular resource, the "gets all" method is consistently named `list()`, as in `client.users.list()` for the endpoint `GET /v2/users`.

### Pagination and Response Changes

All iterable responses, such as those returned by `*.list()` methods, are auto-paginated. This means that code can directly iterate over them without the need for manual pagination logic.

#### Accessing Response Data

**Important:** no longer returns a `data` property by default for endpoints that do not return a paginated response (e.g., `create`, `update`). To retrieve the same `data` property and be able to access headers, you can use `.withRawResponse()`.

#### Migrating from V1 to V2 Pagination

Here's how to migrate your pagination code from v1 to v2:

```go
// V2: Simple pagination with .Results property
list, err := m.Users.List(
    ctx,
    &auth0.ListUsersRequestParameters{
        Page: auth0.Int(1),
        PerPage: auth0.Int(5),
    },
)
if err != nil {
    log.Fatalf("❌ Failed to list users: %v", err)
}

for _, user := range list.Results {
    log.Printf("User ID: %s", *user.GetUserID())
}

// V1: Similar structure but different namespace
// users, err := auth0API.User.List(
//     context.TODO(),
//     management.Page(1),
//     management.PerPage(5),
// )
// if err != nil {
//     log.Fatalf("❌ Failed to list users: %v", err)
// }

// for _, user := range users.Users {
//     log.Printf("User ID: %s", user.GetID())
// }
```

#### Non-Paginated Responses (Create, Update, etc.)

For non-paginated responses, V2 returns the data directly, but you can use `.withRawResponse()` to access headers and the full response:

```go
// V2: Direct data access (no .Body wrapper)
appType := auth0.ClientAppTypeEnumRegularWeb
newClient, err := m.Clients.Create(
    ctx,
    &auth0.CreateClientRequestContent{
        Name: "New Client",
        AppType: &appType,
    },
)
log.Printf("Client ID: %s, Name: %s", newClient.ClientID, newClient.Name)

// V2: Using withRawResponse() to access headers and response metadata
appType := auth0.ClientAppTypeEnumRegularWeb
newClientWithResponse, err := m.Clients.WithRawResponse.Create(
    ctx,
    &auth0.CreateClientRequestContent{
        Name: "New Client",
        AppType: &appType,
    },
)
log.Printf("Client ID: %s, Name: %s", newClientWithResponse.Body.ClientID, newClientWithResponse.Body.Name)
// Access headers: newClientWithResponse.Header
// Access status: newClientWithResponse.StatusCode
```

#### Advanced Pagination

For more complex pagination scenarios, V2 provides enhanced pagination support:

```go
// V1: Manual pagination
orgList, err := m.Organization.List(context.TODO(), management.Take(100))
if err != nil {
    log.Fatalf("err: %+v", err)
}
if !orgList.HasNext() {
    // No need to continue we can stop here.
    return
}
for {
    // Pass the `next` and `take` query parameters now so
    // that we can correctly paginate the organizations.
    orgList, err = m.Organization.List(
        context.TODO(),
        management.From(orgList.Next),
        management.Take(100),
    )
    if err != nil {
        log.Fatalf("err :%+v", err)
    }
    // Accumulate here the results or check for a specific client.

    // The `HasNext` helper func checks whether
    // the API has informed us that there is
    // more data to retrieve or not.
    if !orgList.HasNext() {
        break
    }
}
```

In V2, `client.actions.list()` returns a response of type `Page<Action>`, over which the following pagination code is valid:

```go
page, err := m.Organizations.List(
    ctx,
    &auth0.ListOrganizationsRequestParameters{
        Take: auth0.Int(100),
    },
)
if err != nil {
    log.Fatalf("err: %+v", err)
}

for {
    log.Printf("page's contents: %v", page.Results)

    nextPage, err := page.GetNextPage(ctx);
    if err == nil {
        // out of pages, so exit loop
        if err == core.ErrNoPages {
            break
        }
        // otherwise, it's an actual error
        log.Fatalf("error while fetching next page: %v", err)
    }
    page = nextPage
}
```

### Type name changes

**Important:** Type names have changed drastically in V2, but the structure should remain unchanged for the most part. In V1, type names were auto-generated and often didn't have proper naming according to their actual purpose. V2 introduces properly named types that are more intuitive and follow consistent naming conventions.

If you're using TypeScript types in your code, you'll need to update the type names when migrating from V1 to V2. The new type names are more descriptive and follow a consistent pattern that aligns with the API structure.

```go
// V2: Properly named types
import (
    auth0 "github.com/auth0/go-auth0/management"
)

// Clear, descriptive type names
createRequest := &auth0.CreateClientRequestContent{
    Name: "My Application",
    AppType: auth0.ClientAppTypeEnumRegularWeb.Ptr(),
}

updateRequest := &auth0.UpdateUserRequestContent{
    UserMetadata: &auth0.UserMetadata{
        "role": "admin",
    },
}
```

### Import users takes a Reader

The method corresponding to the endpoint `POST /jobs/users-imports`, through which one can import users in bulk, is labeled `m.Jobs.UsersImports.Create()` (in v1, it was `auth0API.Job.ImportUsers()`).

In v1, this method accepted a `[]map[string]interface{}`:

```go
j := &management.Job{
    Users: []map[string]interface{}{
        {
            "UserMetadata": map[string]interface{}{
                "email": "user@example.com",
            },
        },
    },
}
auth0API.Job.ImportUsers(
    ctx,
    j
)
```

In V2, this method accepts an `io.Reader`, whose definition is copied below. This provides increased flexibility for reading users from a file:

```go
// Treat users like a file stream ...
file, err := os.Open("./users.json")
if err != nil {
    log.Fatalf("err: %+v", err)
}
m.Jobs.UsersImports.Create(
    ctx,
    &jobs.CreateImportUsersRequestContent{
        Users: file,
        ConnectionID: "con_123",
    },
)

// ... or read it in and treat it as a BLOB
fileContent, err := os.ReadFile("./myusers.json")
if err != nil {
    panic(err)
}
m.Jobs.UsersImports.Create(
    ctx,
    &jobs.CreateImportUsersRequestContent{
        Users: bytes.NewReader(fileContent),
        ConnectionID: "con_123",
    },
)
```
