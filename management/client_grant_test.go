package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGrantManager_Create(t *testing.T) {
	client := givenAClient(t)
	resourceServer := givenAResourceServer(t)
	expectedClientGrant := &ClientGrant{
		ClientID: client.ClientID,
		Audience: resourceServer.Identifier,
		Scope:    []interface{}{"create:resource"},
	}

	err := m.ClientGrant.Create(expectedClientGrant)

	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClientGrant.GetID())

	defer cleanupClientGrant(
		t,
		expectedClientGrant.GetID(),
		client.GetClientID(),
		resourceServer.GetID(),
	)
}

func TestClientGrantManager_Read(t *testing.T) {
	expectedClientGrant, clientID, resourceServerID := givenAClientGrant(t)
	defer cleanupClientGrant(t, expectedClientGrant.GetID(), clientID, resourceServerID)

	actualClientGrant, err := m.ClientGrant.Read(expectedClientGrant.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedClientGrant.GetID(), actualClientGrant.GetID())
	assert.Equal(t, expectedClientGrant.GetClientID(), actualClientGrant.GetClientID())
	assert.Equal(t, expectedClientGrant.GetAudience(), actualClientGrant.GetAudience())
}

func TestClientGrantManager_Update(t *testing.T) {
	expectedClientGrant, clientID, resourceServerID := givenAClientGrant(t)
	defer cleanupClientGrant(t, expectedClientGrant.GetID(), clientID, resourceServerID)

	clientGrantID := expectedClientGrant.GetID()

	expectedClientGrant.ID = nil       // Read-Only: Additional properties not allowed.
	expectedClientGrant.Audience = nil // Read-Only: Additional properties not allowed.
	expectedClientGrant.ClientID = nil // Read-Only: Additional properties not allowed.

	expectedClientGrant.Scope = append(expectedClientGrant.Scope, "update:resource")

	err := m.ClientGrant.Update(clientGrantID, expectedClientGrant)

	assert.NoError(t, err)
	assert.Equal(t, len(expectedClientGrant.Scope), 2)
}

func TestClientGrantManager_Delete(t *testing.T) {
	expectedClientGrant, clientID, resourceServerID := givenAClientGrant(t)
	defer cleanupClient(t, clientID)
	defer cleanupResourceServer(t, resourceServerID)

	err := m.ClientGrant.Delete(expectedClientGrant.GetID())

	assert.NoError(t, err)
}

func TestClientGrantManager_List(t *testing.T) {
	expectedClientGrant, clientID, resourceServerID := givenAClientGrant(t)
	defer cleanupClientGrant(t, expectedClientGrant.GetID(), clientID, resourceServerID)

	clientGrantList, err := m.ClientGrant.List(
		Parameter("client_id", expectedClientGrant.GetClientID()),
	)

	assert.NoError(t, err)
	assert.Equal(t, len(clientGrantList.ClientGrants), 1)
}

func givenAClientGrant(t *testing.T) (clientGrant *ClientGrant, clientID string, resourceServerID string) {
	client := givenAClient(t)
	resourceServer := givenAResourceServer(t)

	clientGrant = &ClientGrant{
		ClientID: client.ClientID,
		Audience: resourceServer.Identifier,
		Scope:    []interface{}{"create:resource"},
	}

	err := m.ClientGrant.Create(clientGrant)
	require.NoError(t, err)

	return clientGrant, client.GetClientID(), resourceServer.GetID()
}

func cleanupClientGrant(t *testing.T, clientGrantID, clientID, resourceServerID string) {
	err := m.ClientGrant.Delete(clientGrantID)
	require.NoError(t, err)

	cleanupClient(t, clientID)
	cleanupResourceServer(t, resourceServerID)
}
