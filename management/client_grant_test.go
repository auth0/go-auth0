package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestClientGrantManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	client := givenAClient(t)
	resourceServer := givenAResourceServer(t)
	expectedClientGrant := &ClientGrant{
		ClientID:    client.ClientID,
		Audience:    resourceServer.Identifier,
		Scope:       &[]string{"create:resource"},
		SubjectType: auth0.String("user"),
		AuthorizationDetailsTypes: &[]string{
			"payment",
			"shipping",
		},
	}

	err := api.ClientGrant.Create(context.Background(), expectedClientGrant)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClientGrant.GetID())

	t.Cleanup(func() {
		cleanupClientGrant(t, expectedClientGrant.GetID())
	})
}

func TestClientGrantManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClientGrant := givenAClientGrant(t, false)

	actualClientGrant, err := api.ClientGrant.Read(context.Background(), expectedClientGrant.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedClientGrant.GetID(), actualClientGrant.GetID())
	assert.Equal(t, expectedClientGrant.GetClientID(), actualClientGrant.GetClientID())
	assert.Equal(t, expectedClientGrant.GetAudience(), actualClientGrant.GetAudience())
	assert.Equal(t, expectedClientGrant.GetSubjectType(), actualClientGrant.GetSubjectType())
	assert.Equal(t, expectedClientGrant.GetScope(), actualClientGrant.GetScope())
	assert.Equal(t, expectedClientGrant.GetAuthorizationDetailsTypes(), actualClientGrant.GetAuthorizationDetailsTypes())
}

func TestClientGrantManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClientGrant := givenAClientGrant(t, false)

	clientGrantID := expectedClientGrant.GetID()

	expectedClientGrant.ID = nil          // Read-Only: Additional properties not allowed.
	expectedClientGrant.Audience = nil    // Read-Only: Additional properties not allowed.
	expectedClientGrant.ClientID = nil    // Read-Only: Additional properties not allowed.
	expectedClientGrant.SubjectType = nil // Read-Only: Additional properties not allowed.

	scope := expectedClientGrant.GetScope()
	scope = append(scope, "update:resource")
	expectedClientGrant.Scope = &scope
	expectedClientGrant.AuthorizationDetailsTypes = &[]string{}

	err := api.ClientGrant.Update(context.Background(), clientGrantID, expectedClientGrant)
	assert.NoError(t, err)
	assert.Equal(t, len(expectedClientGrant.GetScope()), 2)
	assert.Empty(t, expectedClientGrant.GetAuthorizationDetailsTypes())
}

func TestClientGrantManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClientGrant := givenAClientGrant(t, false)

	err := api.ClientGrant.Delete(context.Background(), expectedClientGrant.GetID())
	assert.NoError(t, err)

	actualClientGrant, err := api.ClientGrant.Read(context.Background(), expectedClientGrant.GetID())
	assert.Empty(t, actualClientGrant)
	assert.EqualError(t, err, "404 Not Found: Client grant not found")
}

func TestClientGrantManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClientGrant := givenAClientGrant(t, false)

	clientGrantList, err := api.ClientGrant.List(
		context.Background(),
		Parameter("client_id", expectedClientGrant.GetClientID()),
	)

	assert.NoError(t, err)
	assert.Equal(t, len(clientGrantList.ClientGrants), 1)
}

func TestClientGrantManager_Organizations(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	clientGrant := givenAClientGrant(t, true)

	err := api.Organization.AssociateClientGrant(context.Background(), org.GetID(), clientGrant.GetID())
	require.NoError(t, err)

	associatedOrg, err := api.ClientGrant.Organizations(context.Background(), clientGrant.GetID())
	require.NoError(t, err)
	assert.Equal(t, org.GetID(), associatedOrg.Organizations[0].GetID())
}

func givenAClientGrant(t *testing.T, allowOrganizations bool) (clientGrant *ClientGrant) {
	t.Helper()

	client := givenAClient(t)
	resourceServer := givenAResourceServer(t)

	clientGrant = &ClientGrant{
		ClientID:    client.ClientID,
		Audience:    resourceServer.Identifier,
		Scope:       &[]string{"create:resource"},
		SubjectType: auth0.String("user"),
		AuthorizationDetailsTypes: &[]string{
			"payment",
			"shipping",
		},
	}

	if allowOrganizations {
		clientGrant.SubjectType = auth0.String("client")
		clientGrant.AllowAnyOrganization = auth0.Bool(true)
		clientGrant.OrganizationUsage = auth0.String("allow")
		clientGrant.AuthorizationDetailsTypes = nil // Authorization details are not supported with client subject type.
	}

	err := api.ClientGrant.Create(context.Background(), clientGrant)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupClientGrant(t, clientGrant.GetID())
	})

	return clientGrant
}

func cleanupClientGrant(t *testing.T, clientGrantID string) {
	t.Helper()

	err := api.ClientGrant.Delete(context.Background(), clientGrantID)
	require.NoError(t, err)
}
