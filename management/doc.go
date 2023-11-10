// Package management provides a client for using the Auth0 Management API.
//
// # Usage
//
// A management client can be instantiated a domain, client ID and secret.
//
//	import (
//		"github.com/auth0/go-auth0"
//		"github.com/auth0/go-auth0/management"
//	)
//	// Initialize a new client using a domain, client ID and secret.
//	m, err := management.New(
//		domain,
//		management.WithClientCredentials(context.TODO(), id, secret),
//	)
//	if err != nil {
//		// handle err
//	}
//
// Or a domain, and a static token.
//
//	import (
//		"github.com/auth0/go-auth0"
//		"github.com/auth0/go-auth0/management"
//	)
//	// Initialize a new client using a domain and static token.
//	m, err := management.New(domain, management.WithStaticToken(token))
//	if err != nil {
//		// handle err
//	}
//
// With a management client we can then interact with the Auth0 Management API.
//
//	c := &management.Client{
//		Name:        auth0.String("Client Name"),
//		Description: auth0.String("Long description of client"),
//	}
//	err = m.Client.Create(context.TODO(), c)
//	if err != nil {
//		// handle err
//	}
//
// # Authentication
//
// The auth0 management package handles authentication by exchanging the client ID and secret
// supplied when creating a new management client.
//
// This is handled internally using the https://godoc.org/golang.org/x/oauth2
// package.
//
// # Rate Limiting.
//
// The auth0 package also handles rate limiting by respecting the `X-Rate-Limit-*`
// headers sent by the server.
//
// The amount of time the client waits for the rate limit to be reset is taken from
// the `X-Rate-Limit-Reset` header as the amount of seconds to wait.
//
// The rate limit configuration cane be configured using the [WithRetries] helper that allows
// configuring the maximum number of retries and the HTTP status codes to retry on.
// If you wish to disable retries the [WithNoRetries] helper can be used.
//
// # Configuration
//
// There are several options that can be specified during the creation of a
// new client. For a complete list see [Option].
//
//	m, err := management.New(
//		domain,
//		management.WithClientCredentials(context.TODO(), id, secret),
//		management.WithDebug(true),
//	)
//
// # Request Options
//
// As with the global client configuration, fine-grained configuration can be done
// on a request basis. For a complete list see [RequestOption].
//
//	c, err := m.Connection.List(
//		context.TODO(),
//		management.Page(2),
//		management.PerPage(10),
//		management.IncludeFields("id", "name", "options"),
//		management.Parameter("strategy", "auth0"),
//	)
//
// # Page Based Pagination
//
// To use page based pagination use the [Page] and [PerPage] RequestOptions, then the
// HasNext method of the return value can be used to check if there is remaining data.
//
//	for {
//		clients, err := m.Client.List(
//			context.TODO(),
//			management.Page(page),
//			management.PerPage(100),
//		)
//		if err != nil {
//			return err
//		}
//		// Accumulate here the results or check for a specific client.
//
//		// The `HasNext` helper func checks whether
//		// the API has informed us that there is
//		// more data to retrieve or not.
//		if !clients.HasNext() {
//			break
//		}
//		page++
//	}
//
// # Checkpoint Pagination
//
// Checkpoint pagination should be used when you wish to retrieve more than 1000 results. The APIs
// that support it are:
//   - `Log.List` (`/api/v2/logs`)
//   - `Organization.List` (`/api/v2/organizations`)
//   - `Organization.Members` (`/api/v2/organizations/{id}/members`)
//   - `Role.Users` (`/api/v2/roles/{id}/users`)
//
// To use checkpoint pagination, for the first call, only pass the [Take] query parameter,
// the API will then return a `Next` value that can be used for future requests with the [From]
// RequestOption.//
//
//	orgList, err := m.Organization.List(context.TODO(), management.Take(100))
//	if err != nil {
//		log.Fatalf("err: %+v", err)
//	}
//	if !orgList.HasNext() {
//		// No need to continue we can stop here.
//		return
//	}
//	for {
//		// Pass the `next` and `take` query parameters now so
//		// that we can correctly paginate the organizations.
//		orgList, err = m.Organization.List(
//			context.TODO(),
//			management.From(orgList.Next),
//			management.Take(100),
//		)
//		if err != nil {
//			log.Fatalf("err :%+v", err)
//		}
//		// Accumulate here the results or check for a specific client.
//
//		// The `HasNext` helper func checks whether
//		// the API has informed us that there is
//		// more data to retrieve or not.
//		if !orgList.HasNext() {
//			break
//		}
//	}
package management
