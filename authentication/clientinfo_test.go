package authentication

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/auth0/go-auth0/v3/internal/client"
)

// newTestAuthentication creates an Authentication client without any network
// interaction: the HS256 signing algorithm skips the JWKS fetch that the ID
// token validator performs for the RS256 default.
func newTestAuthentication(t *testing.T, entryValue string) *Authentication {
	t.Helper()

	a, err := New(
		context.Background(),
		"go-auth0-dev.eu.auth0.com",
		WithClientID("test-client-id"),
		WithClientSecret("test-client-secret"),
		WithIDTokenSigningAlg("HS256"),
		WithAuth0ClientEnvEntry("test-entry", entryValue),
	)
	if err != nil {
		t.Fatalf("cannot create authentication client: %v", err)
	}

	return a
}

func assertDefaultClientInfoIntact(t *testing.T) {
	t.Helper()

	if got := client.DefaultAuth0ClientInfo().Env["test-entry"]; got != "" {
		t.Fatalf("DefaultAuth0ClientInfo was mutated by WithAuth0ClientEnvEntry: test-entry=%q", got)
	}
}

// TestNewDoesNotShareClientInfo checks that each Authentication is seeded
// with a freshly built client.DefaultAuth0ClientInfo(), so
// WithAuth0ClientEnvEntry neither leaks env entries between clients nor
// affects the values handed to later clients.
func TestNewDoesNotShareClientInfo(t *testing.T) {
	a1 := newTestAuthentication(t, "one")
	a2 := newTestAuthentication(t, "two")

	if a1.auth0ClientInfo == a2.auth0ClientInfo {
		t.Fatal("two authentication clients share the same Auth0ClientInfo pointer")
	}
	a1.auth0ClientInfo.Env["probe"] = "a1"
	if _, leaked := a2.auth0ClientInfo.Env["probe"]; leaked {
		t.Fatal("mutating the Env map of one client leaked into another client's Auth0ClientInfo")
	}
	if _, leaked := client.DefaultAuth0ClientInfo().Env["probe"]; leaked {
		t.Fatal("mutating a client's Env map leaked into DefaultAuth0ClientInfo")
	}
	if got := a1.auth0ClientInfo.Env["test-entry"]; got != "one" {
		t.Fatalf("first client has test-entry=%q, want %q", got, "one")
	}
	if got := a2.auth0ClientInfo.Env["test-entry"]; got != "two" {
		t.Fatalf("second client has test-entry=%q, want %q", got, "two")
	}
	if got := a1.auth0ClientInfo.Env["go"]; got == "" {
		t.Fatal("default env entry \"go\" was lost from the client's copy")
	}
	assertDefaultClientInfoIntact(t)
}

// TestNewConcurrentClientInfo runs many New calls concurrently, as happens
// when a library user constructs clients for different tenants from several
// goroutines. Without a fresh client info per call this races on, and
// intermittently crashes with Go's fatal "concurrent map writes", inside
// WithAuth0ClientEnvEntry. Run with:
//
//	go test -race ./authentication/ -run TestNewConcurrentClientInfo
func TestNewConcurrentClientInfo(t *testing.T) {
	const clients = 64

	var wg sync.WaitGroup
	errCh := make(chan error, clients)
	for i := 0; i < clients; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			a, err := New(
				context.Background(),
				fmt.Sprintf("go-auth0-dev-%d.eu.auth0.com", i),
				WithClientID("test-client-id"),
				WithClientSecret("test-client-secret"),
				WithIDTokenSigningAlg("HS256"),
				WithAuth0ClientEnvEntry("test-entry", fmt.Sprintf("client-%d", i)),
			)
			if err != nil {
				errCh <- fmt.Errorf("client %d: %w", i, err)
				return
			}
			if got, want := a.auth0ClientInfo.Env["test-entry"], fmt.Sprintf("client-%d", i); got != want {
				errCh <- fmt.Errorf("client %d has test-entry=%q, want %q", i, got, want)
			}
		}(i)
	}
	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Error(err)
	}
	assertDefaultClientInfoIntact(t)
}
