package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGrantManager_Create(t *testing.T) {
	setupHTTPRecordings(t)

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

	t.Cleanup(func() {
		cleanupClientGrant(t, expectedClientGrant.GetID())
	})
}

func TestClientGrantManager_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClientGrant := givenAClientGrant(t)

	actualClientGrant, err := m.ClientGrant.Read(expectedClientGrant.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedClientGrant.GetID(), actualClientGrant.GetID())
	assert.Equal(t, expectedClientGrant.GetClientID(), actualClientGrant.GetClientID())
	assert.Equal(t, expectedClientGrant.GetAudience(), actualClientGrant.GetAudience())
}

func TestClientGrantManager_Update(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClientGrant := givenAClientGrant(t)

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
	setupHTTPRecordings(t)

	expectedClientGrant := givenAClientGrant(t)

	err := m.ClientGrant.Delete(expectedClientGrant.GetID())
	assert.NoError(t, err)

	actualClientGrant, err := m.ClientGrant.Read(expectedClientGrant.GetID())
	assert.Empty(t, actualClientGrant)
	assert.EqualError(t, err, "404 Not Found: Client grant not found")
}

func TestClientGrantManager_List(t *testing.T) {
	setupHTTPRecordings(t)

	expectedClientGrant := givenAClientGrant(t)

	clientGrantList, err := m.ClientGrant.List(
		Parameter("client_id", expectedClientGrant.GetClientID()),
	)

	assert.NoError(t, err)
	assert.Equal(t, len(clientGrantList.ClientGrants), 1)
}

func givenAClientGrant(t *testing.T) (clientGrant *ClientGrant) {
	t.Helper()

	client := givenAClient(t)
	resourceServer := givenAResourceServer(t)

	clientGrant = &ClientGrant{
		ClientID: client.ClientID,
		Audience: resourceServer.Identifier,
		Scope:    []interface{}{"create:resource"},
	}

	err := m.ClientGrant.Create(clientGrant)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupClientGrant(t, clientGrant.GetID())
	})

	return clientGrant
}

func cleanupClientGrant(t *testing.T, clientGrantID string) {
	t.Helper()

	err := m.ClientGrant.Delete(clientGrantID)
	require.NoError(t, err)
}
