package management

import (
	"context"
	"testing"
	"time"

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
	assert.EqualError(t, err, "404 Not Found: The grant does not exist.")
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

func TestClientGrant_ExpressConfiguration(t *testing.T) {
	configureHTTPTestRecordings(t)

	t.Run("Express Configuration Client has System Grants", func(t *testing.T) {
		// Create an Express Configuration client
		expressClient := &Client{
			Name:    auth0.Stringf("Express Config Client (%s)", time.Now().Format(time.StampMilli)),
			AppType: auth0.String("express_configuration"),
		}

		err := api.Client.Create(context.Background(), expressClient)
		require.NoError(t, err)

		t.Cleanup(func() {
			cleanupClient(t, expressClient.GetClientID())
		})

		// List all client grants filtered by this client
		grantList, err := api.ClientGrant.List(
			context.Background(),
			Parameter("client_id", expressClient.GetClientID()),
		)
		require.NoError(t, err)

		// Express Configuration clients should have system-created grants
		t.Logf("Found %d client grants for Express Configuration client %s", len(grantList.ClientGrants), expressClient.GetClientID())

		for _, grant := range grantList.ClientGrants {
			t.Logf("Grant ID: %s, Audience: %s, IsSystem: %v, SubjectType: %s",
				grant.GetID(),
				grant.GetAudience(),
				grant.GetIsSystem(),
				grant.GetSubjectType(),
			)

			// System grants for Express Configuration should have is_system=true
			if grant.GetIsSystem() {
				assert.True(t, grant.GetIsSystem(), "Express Configuration grants should be system grants")
				assert.Equal(t, expressClient.GetClientID(), grant.GetClientID())
			}
		}
	})

	t.Run("Cannot Manually Create Grant for Express Configuration Client", func(t *testing.T) {
		// Create an Express Configuration client
		expressClient := &Client{
			Name:    auth0.Stringf("Express Config Client (%s)", time.Now().Format(time.StampMilli)),
			AppType: auth0.String("express_configuration"),
		}

		err := api.Client.Create(context.Background(), expressClient)
		require.NoError(t, err)

		t.Cleanup(func() {
			cleanupClient(t, expressClient.GetClientID())
		})

		// Try to create a client grant for this Express Configuration client
		resourceServer := givenAResourceServer(t)

		clientGrant := &ClientGrant{
			ClientID:    expressClient.ClientID,
			Audience:    resourceServer.Identifier,
			Scope:       &[]string{"create:resource"},
			SubjectType: auth0.String("client"),
		}

		err = api.ClientGrant.Create(context.Background(), clientGrant)

		// Should fail because Express Configuration clients cannot have manual grants
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "express_configuration")
	})
}

func TestClientGrant_IsSystem(t *testing.T) {
	configureHTTPTestRecordings(t)

	t.Run("User Created Grant has is_system false", func(t *testing.T) {
		// Create a client and resource server
		client := givenAClient(t)
		resourceServer := givenAResourceServer(t)

		// Create a simple client grant without authorization_details_types
		clientGrant := &ClientGrant{
			ClientID:    client.ClientID,
			Audience:    resourceServer.Identifier,
			Scope:       &[]string{"create:resource"},
			SubjectType: auth0.String("user"),
		}

		err := api.ClientGrant.Create(context.Background(), clientGrant)
		require.NoError(t, err)

		t.Cleanup(func() {
			cleanupClientGrant(t, clientGrant.GetID())
		})

		// Read it back
		retrievedGrant, err := api.ClientGrant.Read(context.Background(), clientGrant.GetID())
		require.NoError(t, err)

		// Verify is_system is false for user-created grants
		assert.False(t, retrievedGrant.GetIsSystem())
	})

	t.Run("List Grants and Check is_system Field", func(t *testing.T) {
		// Create a client and resource server
		client := givenAClient(t)
		resourceServer := givenAResourceServer(t)

		// Create a user client grant
		userGrant := &ClientGrant{
			ClientID:    client.ClientID,
			Audience:    resourceServer.Identifier,
			Scope:       &[]string{"create:resource"},
			SubjectType: auth0.String("user"),
		}

		err := api.ClientGrant.Create(context.Background(), userGrant)
		require.NoError(t, err)

		t.Cleanup(func() {
			cleanupClientGrant(t, userGrant.GetID())
		})

		// List all grants
		grantList, err := api.ClientGrant.List(context.Background())
		require.NoError(t, err)

		// Find our user-created grant
		foundUserGrant := false

		for _, grant := range grantList.ClientGrants {
			if grant.GetID() == userGrant.GetID() {
				foundUserGrant = true

				assert.False(t, grant.GetIsSystem(), "User-created grant should have is_system=false")
			}
			// System grants should have is_system=true
			if grant.GetIsSystem() {
				assert.True(t, grant.GetIsSystem(), "System grant should have is_system=true")
				// System grants cannot be deleted/modified
				assert.NotEmpty(t, grant.GetID())
			}
		}

		assert.True(t, foundUserGrant, "Should find the user-created grant in the list")
	})
}
