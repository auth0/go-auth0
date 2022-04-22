package management

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestResourceServer_Create(t *testing.T) {
	expectedResourceServer := &ResourceServer{
		Name:                auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:          auth0.String("https://api.example.com/"),
		SigningAlgorithm:    auth0.String("HS256"),
		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),
		Scopes: []*ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
	}

	err := m.ResourceServer.Create(expectedResourceServer)

	assert.NoError(t, err)
	assert.NotEmpty(t, expectedResourceServer.GetID())

	defer cleanupResourceServer(t, expectedResourceServer.GetID())
}

func TestResourceServer_Read(t *testing.T) {
	expectedResourceServer := givenAResourceServer(t)
	defer cleanupResourceServer(t, expectedResourceServer.GetID())

	actualResourceServer, err := m.ResourceServer.Read(expectedResourceServer.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetName(), actualResourceServer.GetName())
}

func TestResourceServer_Update(t *testing.T) {
	expectedResourceServer := givenAResourceServer(t)
	defer cleanupResourceServer(t, expectedResourceServer.GetID())

	resourceServerID := expectedResourceServer.GetID()

	expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
	expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.

	expectedResourceServer.AllowOfflineAccess = auth0.Bool(true)
	expectedResourceServer.SigningAlgorithm = auth0.String("RS256")
	expectedResourceServer.SkipConsentForVerifiableFirstPartyClients = auth0.Bool(true)
	expectedResourceServer.TokenLifetime = auth0.Int(7200)
	expectedResourceServer.TokenLifetimeForWeb = auth0.Int(5400)
	expectedResourceServer.Scopes = append(expectedResourceServer.Scopes, &ResourceServerScope{
		Value:       auth0.String("update:resource"),
		Description: auth0.String("Update Resource"),
	})

	err := m.ResourceServer.Update(resourceServerID, expectedResourceServer)

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetAllowOfflineAccess(), true)
	assert.Equal(t, expectedResourceServer.GetSigningAlgorithm(), "RS256")
	assert.Equal(t, expectedResourceServer.GetSkipConsentForVerifiableFirstPartyClients(), true)
	assert.Equal(t, expectedResourceServer.GetTokenLifetime(), 7200)
	assert.Equal(t, expectedResourceServer.GetTokenLifetimeForWeb(), 5400)
	assert.Equal(t, len(expectedResourceServer.Scopes), 2)
}

func TestResourceServer_Delete(t *testing.T) {
	expectedResourceServer := givenAResourceServer(t)

	err := m.ResourceServer.Delete(expectedResourceServer.GetID())

	assert.NoError(t, err)

	actualResourceServer, err := m.ResourceServer.Read(expectedResourceServer.GetID())

	assert.Empty(t, actualResourceServer)
	assert.EqualError(t, err, "404 Not Found: The resource server does not exist")
}

func TestResourceServer_List(t *testing.T) {
	expectedResourceServer := givenAResourceServer(t)
	defer cleanupResourceServer(t, expectedResourceServer.GetID())

	resourceServerList, err := m.ResourceServer.List(IncludeFields("id"))

	assert.NoError(t, err)
	assert.Contains(t, resourceServerList.ResourceServers, &ResourceServer{ID: expectedResourceServer.ID})
}

func givenAResourceServer(t *testing.T) *ResourceServer {
	resourceServer := &ResourceServer{
		Name:                auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:          auth0.String("https://api.example.com/"),
		SigningAlgorithm:    auth0.String("HS256"),
		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),
		Scopes: []*ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
	}

	err := m.ResourceServer.Create(resourceServer)
	require.NoError(t, err)

	return resourceServer
}

func cleanupResourceServer(t *testing.T, resourceServerID string) {
	err := m.ResourceServer.Delete(resourceServerID)
	require.NoError(t, err)
}
