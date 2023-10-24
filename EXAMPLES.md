# Examples

- [Request Options](#request-options)
- [Pagination](#pagination)
  - [Page based pagination](#page-based-pagination)
  - [Checkpoint pagination](#checkpoint-pagination)
- [Custom User Structs](#providing-a-custom-user-struct)

## Request Options

Fine-grained configuration can be provided on a per-request basis to enhance the request with specific query params, headers, or to pass it a custom context.

> **Note**
> Not all of the API endpoints support the query parameters added by these funcs.
> Review the [API docs](https://auth0.com/docs/api/management/v2) for the full documentation of the supported parameters per endpoint.

```go
// Example
userGrants, err := auth0API.Grant.List(
    context.Background(),
    management.Header("MySpecialHeader","MySpecialHeaderValue"),
    management.Parameter("user_id", "someUserID"),
    management.Parameter("client", "someClientID"),
    management.Parameter("audience", "someAPIAudience"),
)

// Other helper funcs
management.Query()
management.ExcludeFields()
management.IncludeFields()
management.Page()
management.PerPage()
management.IncludeTotals()
management.Take()
management.From()
management.Sort()
```

## Pagination

This SDK supports both offset and checkpoint pagination.

### Page based pagination

When retrieving lists of resources, if no query parameters are set using the `management.PerPage` and `Management.IncludeTotals` helper funcs, then the SDK will default to sending `per_page=50` and `include_totals=true`. 

> **Note**
> The maximum value of the `per_page` query parameter is 100.

<details>
  <summary>Page based pagination example</summary>

```go
var page int
for {
    clients, err := auth0API.Client.List(
        context.Background(),
        management.Page(page),
        management.PerPage(100),
    )
    if err != nil {
        return err
    }

    // Accumulate here the results or check for a specific client.

    if !clients.HasNext() {
        break
    }

    page++
}
```
</details>

### Checkpoint pagination

Checkpoint pagination can be used when you wish to retrieve more than 1000 results from certain APIs. The APIs that support checkpoint based pagination are:

* `Log.List` (`/api/v2/logs`)
* `Organization.List` (`/api/v2/organizations`)
* `Organization.Members` (`/api/v2/organizations/{id}/members`)
* `Role.Users` (`/api/v2/roles/{id}/users`)

<details>
  <summary>Checkpoint pagination example</summary>

```go
// For the first call, only pass the `take` query parameter, the API will
// then return a `Next` value that can be used for future requests.
orgList, err := auth0API.Organization.List(context.Background(), management.Take(100))
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
    orgList, err = auth0API.Organization.List(
        context.Background(),
        management.From(orgList.Next),
        management.Take(100),
    )
    if err != nil {
        log.Fatalf("err :%+v", err)
    }
    
    for _, org := range orgList.Organizations {
        log.Printf("org %s", org.GetID())
    }
    
    // The `HasNext` helper func checks whether
    // the API has informed us that there is
    // more data to retrieve or not.
    if !orgList.HasNext() {
        break
    }
}
```
</details>

However, for `Log.List`, the `Next` value is not returned via the API but instead is an ID of a log entry. Determining if there are more logs to retrieved must also be done manually.

<details>
  <summary>Checkpoint pagination example for <code>Log.List</code></summary>

```go
var logs []*management.Log
initialLogId := "LOGID"
for {
    // Retrieve 100 logs after the specified log
    logs, err = auth0API.Log.List(
        context.Background(),
        management.From(logFromId),
        management.Take(100),
    )

    if err != nil {
        log.Fatalf("err: %+v", err)
    }

    for _, logData := range logs {
        log.Printf("ID %s", logData.GetID())
        log.Printf("Type %s", logData.GetType())

    }

    // The HasNext helper cannot be used with `Log.List` so instead we check the length of the
    // returned logs array. When it reaches 0 there are no more logs left to process.
    if len(logs) == 0 {
        break
    }

    logFromId = logs[len(logs)-1].GetID()
}
```
</details>

## Providing a custom User struct

The `management.User` struct within the SDK only contains the properties supported by Auth0. Therefore, any extra properties added by an external identity provider will not be included within the struct returned from the SDK APIs. To expose these custom properties, we recommend creating a custom struct and then manually calling the API via the lower level request functionality exposed by the SDK, as shown below.

First, define a custom struct that embeds the `management.User` struct exposed by the SDK, and add any helper funcs required to safely access your custom values.

```go
// Define a custom struct that embeds the `management.User` struct exposed by the SDK.
type CustomUser struct {
	management.User
	OurCustomID *string `json:"custom_id,omitempty"`
}

// Create a helper func that will safely retrieve the `OurCustomId` value from CustomUser in the
// cases where it may be nil.
func (u *CustomUser) GetOurCustomID() string {
	if u == nil || u.OurCustomID == nil {
		return ""
	}
	return *u.OurCustomID
}
```

Then, create a request using the lower level request functionality exposed by the SDK.

```go
var user CustomUser
err := auth0API.Request(context.Background(), http.MethodGet, auth0API.URI("users", "auth0|63cfb8ca89c31c3f33f1dffd"), &user)
if err != nil {
    log.Fatalf("error was %+v", err)
}
log.Printf("User %s", user.GetOurCustomID())
```