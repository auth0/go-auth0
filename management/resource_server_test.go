package management

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestResourceServer_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := &ResourceServer{
		Name:                auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:          auth0.String("https://api.example.com/"),
		SigningAlgorithm:    auth0.String("HS256"),
		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),
		Scopes: &[]ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
		EnforcePolicies: auth0.Bool(true),
		TokenDialect:    auth0.String("rfc9068_profile_authz"),
	}

	err := api.ResourceServer.Create(context.Background(), expectedResourceServer)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedResourceServer.GetID())

	t.Cleanup(func() {
		cleanupResourceServer(t, expectedResourceServer.GetID())
	})
}

func TestResourceServer_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	actualResourceServer, err := api.ResourceServer.Read(context.Background(), expectedResourceServer.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetName(), actualResourceServer.GetName())
}

func TestResourceServer_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	resourceServerID := expectedResourceServer.GetID()

	expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
	expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
	expectedResourceServer.SigningSecret = nil

	expectedResourceServer.AllowOfflineAccess = auth0.Bool(true)
	expectedResourceServer.SigningAlgorithm = auth0.String("RS256")
	expectedResourceServer.SkipConsentForVerifiableFirstPartyClients = auth0.Bool(true)
	expectedResourceServer.TokenLifetime = auth0.Int(7200)
	expectedResourceServer.TokenLifetimeForWeb = auth0.Int(5400)
	scopes := append(expectedResourceServer.GetScopes(), ResourceServerScope{
		Value:       auth0.String("update:resource"),
		Description: auth0.String("Update Resource"),
	})
	expectedResourceServer.Scopes = &scopes
	expectedResourceServer.EnforcePolicies = auth0.Bool(true)
	expectedResourceServer.TokenDialect = auth0.String("access_token_authz")

	err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetAllowOfflineAccess(), true)
	assert.Equal(t, expectedResourceServer.GetSigningAlgorithm(), "RS256")
	assert.Equal(t, expectedResourceServer.GetSkipConsentForVerifiableFirstPartyClients(), true)
	assert.Equal(t, expectedResourceServer.GetTokenLifetime(), 7200)
	assert.Equal(t, expectedResourceServer.GetTokenLifetimeForWeb(), 5400)
	assert.Equal(t, len(expectedResourceServer.GetScopes()), 2)
	assert.Equal(t, expectedResourceServer.GetTokenDialect(), "access_token_authz")
	assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), true)
}

func TestResourceServer_TokenDialect(t *testing.T) {

	t.Run("When_TokenDialect_is_rfc9068_profile_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("rfc9068_profile")
		expectedResourceServer.EnforcePolicies = auth0.Bool(false)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "rfc9068_profile")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), false)
	})

	t.Run("When_TokenDialect_is_access_token_authz_and_RBAC_enabled_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("access_token_authz")
		expectedResourceServer.EnforcePolicies = auth0.Bool(true)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "access_token_authz")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), true)
	})

	t.Run("When_TokenDialect_is_rfc9068_profile_authz_and_RBAC_enabled_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("rfc9068_profile_authz")
		expectedResourceServer.EnforcePolicies = auth0.Bool(true)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "rfc9068_profile_authz")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), true)
	})

	t.Run("When_TokenDialect_is_access_token_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("access_token")
		expectedResourceServer.EnforcePolicies = auth0.Bool(false)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "access_token")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), false)
	})
}

func TestResourceServer_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	err := api.ResourceServer.Delete(context.Background(), expectedResourceServer.GetID())
	assert.NoError(t, err)

	actualResourceServer, err := api.ResourceServer.Read(context.Background(), expectedResourceServer.GetID())
	assert.Empty(t, actualResourceServer)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestResourceServer_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	resourceServerList, err := api.ResourceServer.List(context.Background(), IncludeFields("id"))

	assert.NoError(t, err)
	assert.Contains(t, resourceServerList.ResourceServers, &ResourceServer{ID: expectedResourceServer.ID})
}

func givenAResourceServer(t *testing.T) *ResourceServer {
	t.Helper()

	resourceServer := &ResourceServer{
		Name:                auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:          auth0.String("https://api.example.com/"),
		SigningAlgorithm:    auth0.String("HS256"),
		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),
		TokenDialect:        auth0.String("access_token"),
		EnforcePolicies:     auth0.Bool(false),
		Scopes: &[]ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
	}

	err := api.ResourceServer.Create(context.Background(), resourceServer)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupResourceServer(t, resourceServer.GetID())
	})

	return resourceServer
}

func cleanupResourceServer(t *testing.T, resourceServerID string) {
	t.Helper()

	err := api.ResourceServer.Delete(context.Background(), resourceServerID)
	require.NoError(t, err)
}
