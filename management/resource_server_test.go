package management

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestResourceServer_Create(t *testing.T) {
	setupHTTPRecordings(t)

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

	t.Cleanup(func() {
		cleanupResourceServer(t, expectedResourceServer.GetID())
	})
}

func TestResourceServer_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	actualResourceServer, err := m.ResourceServer.Read(expectedResourceServer.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetName(), actualResourceServer.GetName())
}

func TestResourceServer_Update(t *testing.T) {
	setupHTTPRecordings(t)

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
	setupHTTPRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	err := m.ResourceServer.Delete(expectedResourceServer.GetID())
	assert.NoError(t, err)

	actualResourceServer, err := m.ResourceServer.Read(expectedResourceServer.GetID())
	assert.Empty(t, actualResourceServer)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestResourceServer_List(t *testing.T) {
	setupHTTPRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	resourceServerList, err := m.ResourceServer.List(IncludeFields("id"))

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
		Scopes: []*ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
	}

	err := m.ResourceServer.Create(resourceServer)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupResourceServer(t, resourceServer.GetID())
	})

	return resourceServer
}

func cleanupResourceServer(t *testing.T, resourceServerID string) {
	t.Helper()

	err := m.ResourceServer.Delete(resourceServerID)
	require.NoError(t, err)
}
