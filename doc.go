/*
Package auth0 provides a client for using the Auth0 Management API.

Usage

	import (
	    github.com/auth0/go-auth0
	    github.com/auth0/go-auth0/management
	)

Initialize a new client using a domain, client ID and secret.

	m, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), id, secret)
	)
	if err != nil {
	    // handle err
	}

Or using a static token.

	m, err := management.New(domain, management.WithStaticToken(token))
	if err != nil {
	    // handle err
	}

With a management client we can then interact with the Auth0 Management API.

	c := &management.Client{
	    Name:        auth0.String("Client Name"),
	    Description: auth0.String("Long description of client"),
	}

	err = m.Client.Create(context.Background(), c)
	if err != nil {
	    // handle err
	}

# Authentication

The auth0 package handles authentication by exchanging the client id and secret
supplied when creating a new management client.

This is handled internally using the https://godoc.org/golang.org/x/oauth2
package.

# Rate Limiting

The auth0 package also handles rate limiting by respecting the `X-Rate-Limit-*`
headers sent by the server.

The amount of time the client waits for the rate limit to be reset is taken from
the `X-Rate-Limit-Reset` header as the amount of seconds to wait.

# Configuration

There are several other options that can be specified during the creation of a
new client.

	m, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), id, secret),
		management.WithDebug(true)
	)

# Request Options

As with the global client configuration, fine-grained configuration can be done
on a request basis.

	c, err := m.Connection.List(
		context.Background(),
		management.Page(2),
		management.PerPage(10),
		management.IncludeFields("id", "name", "options")
		management.Parameter("strategy", "auth0"),
	)
*/
package auth0
