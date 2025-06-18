package management

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestOrganizationManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	orgConn := givenAnOrganizationConnectionWithoutOrgID(t)
	orgConn2 := givenAnOrganizationDBConnectionWithoutOrgID(t)

	org := &Organization{
		Name:        auth0.String(fmt.Sprintf("test-organization%v", rand.Intn(999))),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoURL: auth0.String("https://example.com/logo.gif"),
		},
		EnabledConnections: []*OrganizationConnection{orgConn, orgConn2},
	}

	err := api.Organization.Create(context.Background(), org)
	assert.NoError(t, err)
	assert.NotEmpty(t, org.GetID())

	t.Cleanup(func() {
		cleanupOrganization(t, org.GetID())
	})
}

func TestOrganizationManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	actualOrg, err := api.Organization.Read(context.Background(), org.GetID())

	assert.NoError(t, err)
	assert.Equal(t, org, actualOrg)
}

func TestOrganizationManager_ReadByName(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	actualOrg, err := api.Organization.ReadByName(context.Background(), org.GetName())

	assert.NoError(t, err)
	assert.Equal(t, org, actualOrg)
}

func TestOrganizationManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	err := api.Organization.Update(context.Background(), org.GetID(), &Organization{Name: auth0.String("new-org-name")})
	assert.NoError(t, err)

	actualOrg, err := api.Organization.Read(context.Background(), org.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "new-org-name", actualOrg.GetName())
}

func TestOrganizationManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	err := api.Organization.Delete(context.Background(), org.GetID())
	assert.NoError(t, err)

	actualOrg, err := api.Organization.Read(context.Background(), org.GetID())
	assert.Empty(t, actualOrg)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenAnOrganization(t)

	orgList, err := api.Organization.List(context.Background())
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(orgList.Organizations), 1)
}

func TestOrganizationManager_ListCheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	for i := 0; i < 3; i++ {
		givenAnOrganization(t)
	}

	orgList, err := api.Organization.List(context.Background(), Take(2))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(orgList.Organizations), 2)
	assert.True(t, orgList.HasNext())

	orgList, err = api.Organization.List(context.Background(), Take(2), From(orgList.Next))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(orgList.Organizations), 1)
	assert.False(t, orgList.HasNext())
}

func TestOrganizationManager_AddConnection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	client := givenAClient(t)
	conn := givenAConnection(t, connectionTestCase{connection: Connection{
		Name:        auth0.String(fmt.Sprintf("test-conn%v", rand.Intn(999))),
		DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", rand.Intn(999))),
		Strategy:    auth0.String(ConnectionStrategyAD),
		EnabledClients: &[]string{
			os.Getenv("AUTH0_CLIENT_ID"),
			client.GetClientID(),
		},
	}})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
		ShowAsButton:            auth0.Bool(true),
	}

	err := api.Organization.AddConnection(context.Background(), org.GetID(), orgConn)
	assert.NoError(t, err)
	assert.NotEmpty(t, orgConn.GetConnection())
}

func TestOrganizationManager_Connection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	actualOrgConn, err := api.Organization.Connection(context.Background(), org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, orgConn, actualOrgConn)
}

func TestOrganizationManager_DBConnection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationDBConnection(t, org.GetID())

	actualOrgConn, err := api.Organization.Connection(context.Background(), org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, orgConn, actualOrgConn)
}

func TestOrganizationManager_UpdateConnection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	err := api.Organization.UpdateConnection(
		context.Background(),
		org.GetID(),
		orgConn.GetConnectionID(),
		&OrganizationConnection{
			AssignMembershipOnLogin: auth0.Bool(false),
			ShowAsButton:            auth0.Bool(false),
		},
	)
	assert.NoError(t, err)

	actualOrgConn, err := api.Organization.Connection(context.Background(), org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, false, actualOrgConn.GetAssignMembershipOnLogin())
	assert.Equal(t, false, actualOrgConn.GetShowAsButton())
}

// TestOrganizationManager_UpdateDBConnection tests the UpdateConnection method of OrganizationManager for Database Connection.
func TestOrganizationManager_UpdateDBConnection(t *testing.T) {
	// Test when IsSignupEnabled true with AssignMembershipOnLogin false should fail
	t.Run("When_signup_enabled_with_assign_membership_false_should_fail", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		org := givenAnOrganization(t)
		orgConn := givenAnOrganizationDBConnection(t, org.GetID())

		err := api.Organization.UpdateConnection(
			context.Background(),
			org.GetID(),
			orgConn.GetConnectionID(),
			&OrganizationConnection{
				AssignMembershipOnLogin: auth0.Bool(false),
				IsSignupEnabled:         auth0.Bool(true),
			},
		)
		assert.Error(t, err, "Expected error when is_signup_enabled is true and assign_membership_on_login is false")
		assert.Contains(t, err.Error(), "Only database connections with assign_membership_on_login = true support is_signup_enabled = true.")
	})

	// Test when IsSignupEnabled and AssignMembershipOnLogin are false, should succeed
	t.Run("When_signup_and_assign_membership_are_false_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		org := givenAnOrganization(t)
		orgConn := givenAnOrganizationDBConnection(t, org.GetID())

		err := api.Organization.UpdateConnection(
			context.Background(),
			org.GetID(),
			orgConn.GetConnectionID(),
			&OrganizationConnection{
				AssignMembershipOnLogin: auth0.Bool(false),
				IsSignupEnabled:         auth0.Bool(false),
			},
		)
		assert.NoError(t, err)

		actualOrgConn, err := api.Organization.Connection(context.Background(), org.GetID(), orgConn.GetConnectionID())
		assert.NoError(t, err)
		assert.Equal(t, false, actualOrgConn.GetAssignMembershipOnLogin())
		assert.Equal(t, false, actualOrgConn.GetIsSignupEnabled())
	})

	// Test when IsSignupEnabled with AssignMembershipOnLogin true should succeed
	t.Run("When_signup_enabled_with_assign_membership_true_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		org := givenAnOrganization(t)
		orgConn := givenAnOrganizationDBConnection(t, org.GetID())

		err := api.Organization.UpdateConnection(
			context.Background(),
			org.GetID(),
			orgConn.GetConnectionID(),
			&OrganizationConnection{
				AssignMembershipOnLogin: auth0.Bool(true),
				IsSignupEnabled:         auth0.Bool(true),
			},
		)
		assert.NoError(t, err)

		actualOrgConn, err := api.Organization.Connection(context.Background(), org.GetID(), orgConn.GetConnectionID())
		assert.NoError(t, err)
		assert.Equal(t, true, actualOrgConn.GetAssignMembershipOnLogin())
		assert.Equal(t, true, actualOrgConn.GetIsSignupEnabled())
	})
}

func TestOrganizationManager_DeleteConnection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	err := api.Organization.DeleteConnection(context.Background(), org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)

	actualOrgConn, err := api.Organization.Connection(context.Background(), org.GetID(), orgConn.GetConnectionID())
	assert.Error(t, err)
	assert.Empty(t, actualOrgConn)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_Connections(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	orgConnList, err := api.Organization.Connections(context.Background(), org.GetID())
	assert.NoError(t, err)
	assert.Len(t, orgConnList.OrganizationConnections, 1)
	assert.Equal(t, orgConn.GetConnectionID(), orgConnList.OrganizationConnections[0].GetConnectionID())
}

func TestOrganizationManager_CreateInvitation(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	client := givenAClient(t)
	orgInvite := &OrganizationInvitation{
		Inviter: &OrganizationInvitationInviter{
			Name: auth0.String("Test Inviter"),
		},
		Invitee: &OrganizationInvitationInvitee{
			Email: auth0.String("test@example.com"),
		},
		ClientID: client.ClientID,
	}

	err := api.Organization.CreateInvitation(context.Background(), org.GetID(), orgInvite)
	assert.NoError(t, err)
	assert.NotEmpty(t, orgInvite.GetID())
}

func TestOrganizationManager_Invitation(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	actualOrgInvite, err := api.Organization.Invitation(context.Background(), org.GetID(), orgInvite.GetID())
	assert.NoError(t, err)
	assert.Equal(t, orgInvite, actualOrgInvite)
}

func TestOrganizationManager_DeleteInvitation(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	err := api.Organization.DeleteInvitation(context.Background(), org.GetID(), orgInvite.GetID())
	assert.NoError(t, err)

	actualOrgInvite, err := api.Organization.Invitation(context.Background(), org.GetID(), orgInvite.GetID())
	assert.Error(t, err)
	assert.Empty(t, actualOrgInvite)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_Invitations(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	invitations, err := api.Organization.Invitations(context.Background(), org.GetID())
	assert.NoError(t, err)
	assert.Len(t, invitations.OrganizationInvitations, 1)
	assert.Equal(t, orgInvite.GetID(), invitations.OrganizationInvitations[0].GetID())
}

func TestOrganizationManager_AddMembers(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := api.Organization.AddMembers(context.Background(), org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)
}

func TestOrganizationManager_DeleteMembers(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := api.Organization.AddMembers(context.Background(), org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = api.Organization.DeleteMembers(context.Background(), org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	members, err := api.Organization.Members(context.Background(), org.GetID())
	assert.NoError(t, err)
	assert.Len(t, members.Members, 0)
}

func TestOrganizationManager_Members(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)
	role := givenARole(t)

	err := api.Organization.AddMembers(context.Background(), org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = api.Organization.AssignMemberRoles(context.Background(), org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	members, err := api.Organization.Members(
		context.Background(), org.GetID(),
		IncludeFields("user_id", "roles"),
	)

	assert.NoError(t, err)
	assert.Len(t, members.Members, 1)
	assert.Equal(t, user.GetID(), members.Members[0].GetUserID())

	assert.Len(t, members.Members[0].Roles, 1)
	assert.Equal(t, role.GetID(), members.Members[0].Roles[0].GetID())
}

func TestOrganizationManager_MembersCheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	users := make([]string, 0)

	for i := 0; i < 3; i++ {
		user := givenAUser(t)
		users = append(users, user.GetID())
	}

	err := api.Organization.AddMembers(context.Background(), org.GetID(), users)
	assert.NoError(t, err)

	members, err := api.Organization.Members(context.Background(), org.GetID(), Take(2))
	assert.NoError(t, err)
	assert.Len(t, members.Members, 2)
	assert.True(t, members.HasNext())

	members, err = api.Organization.Members(context.Background(), org.GetID(), Take(2), From(members.Next))
	assert.NoError(t, err)
	assert.Len(t, members.Members, 1)
	assert.True(t, members.HasNext())

	// Org members pagination will only return an empty `Next` value when the number of members
	// returned is 0 unlike other pagination APIs
	members, err = api.Organization.Members(context.Background(), org.GetID(), Take(2), From(members.Next))
	assert.NoError(t, err)
	assert.Len(t, members.Members, 0)
	assert.False(t, members.HasNext())
}

func TestOrganizationManager_MemberRoles(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)
	role := givenARole(t)

	err := api.Organization.AddMembers(context.Background(), org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = api.Organization.AssignMemberRoles(context.Background(), org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	roles, err := api.Organization.MemberRoles(context.Background(), org.GetID(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 1)
	assert.Equal(t, role.GetID(), roles.Roles[0].GetID())

	err = api.Organization.DeleteMemberRoles(context.Background(), org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	roles, err = api.Organization.MemberRoles(context.Background(), org.GetID(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 0)
}

func TestOrganizationManager_ClientGrants(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	clientGrant := givenAClientGrant(t, true)

	err := api.Organization.AssociateClientGrant(context.Background(), org.GetID(), clientGrant.GetID())
	require.NoError(t, err)

	associatedGrants, err := api.Organization.ClientGrants(context.Background(), org.GetID())
	require.NoError(t, err)
	assert.Len(t, associatedGrants.ClientGrants, 1)
	assert.Equal(t, clientGrant.GetID(), associatedGrants.ClientGrants[0].GetID())

	err = api.Organization.RemoveClientGrant(context.Background(), org.GetID(), clientGrant.GetID())
	require.NoError(t, err)

	associatedGrants, err = api.Organization.ClientGrants(context.Background(), org.GetID())
	require.NoError(t, err)
	assert.Len(t, associatedGrants.ClientGrants, 0)
}

func TestOrganizationManager_ClientGrantsWithOrg(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	resourceServer := givenAResourceServer(t)

	client := &Client{
		Name:              auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description:       auth0.String("This is just a test client."),
		OrganizationUsage: auth0.String("allow"),
		DefaultOrganization: &ClientDefaultOrganization{
			Flows:          &[]string{"client_credentials"},
			OrganizationID: auth0.String(org.GetID()),
		},
	}
	// Create a client that shall be used for testing.
	err := api.Client.Create(context.Background(), client)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupClient(t, client.GetClientID())
	})

	clientGrant := &ClientGrant{
		ClientID:             client.ClientID,
		Audience:             resourceServer.Identifier,
		Scope:                &[]string{"create:resource", "create:organization_client_grants"},
		AllowAnyOrganization: auth0.Bool(true),
		OrganizationUsage:    auth0.String("allow"),
	}

	// Create a clientGrant and associate with the client created above.
	err = api.ClientGrant.Create(context.Background(), clientGrant)
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupClientGrant(t, clientGrant.GetID())
	})

	// Associates the grant with an organization.
	err = api.Organization.AssociateClientGrant(context.Background(), org.GetID(), clientGrant.GetID())
	require.NoError(t, err)

	// List all clients associated with a ClientGrant given an organizationID as query param
	clients, err := api.Client.List(context.Background(), Parameter("q", fmt.Sprintf("client_grant.organization_id:%s", org.GetID())))
	require.NoError(t, err)

	for _, c := range clients.Clients {
		assert.Equal(t, org.GetID(), c.DefaultOrganization.GetOrganizationID())
	}

	// List all ClientGrants given a list of grant_ids as query param
	associatedGrants, err := api.Organization.ClientGrants(context.Background(), org.GetID(), Parameter("grant_ids", clientGrant.GetID()))
	require.NoError(t, err)
	assert.Greater(t, len(associatedGrants.ClientGrants), 0)
	assert.Contains(t, associatedGrants.ClientGrants, clientGrant)

	// Remove the associated ClientGrants
	err = api.Organization.RemoveClientGrant(context.Background(), org.GetID(), clientGrant.GetID())
	require.NoError(t, err)

	// List all ClientGrants which should be an empty list since grant has been removed from the organization.
	associatedGrants, err = api.Organization.ClientGrants(context.Background(), org.GetID(), Parameter("grant_ids", clientGrant.GetID()))
	require.NoError(t, err)
	assert.Len(t, associatedGrants.ClientGrants, 0)

	// Delete the ClientGrant.
	err = api.ClientGrant.Delete(context.Background(), clientGrant.GetID())
	require.NoError(t, err)

	// Retrieve the ClientGrant and ensure error is return since grant has been deleted.
	retrievedGrant, err := api.ClientGrant.Read(context.Background(), clientGrant.GetID())
	assert.Nil(t, retrievedGrant)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenAnOrganization(t *testing.T) *Organization {
	org := &Organization{
		Name:        auth0.String(fmt.Sprintf("test-organization%v", rand.Intn(999))),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoURL: auth0.String("https://example.com/logo.gif"),
		},
	}

	err := api.Organization.Create(context.Background(), org)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupOrganization(t, org.GetID())
	})

	return org
}

func givenAnOrganizationConnection(t *testing.T, orgID string) *OrganizationConnection {
	client := givenAClient(t)
	conn := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:        auth0.String(fmt.Sprintf("test-conn%v", rand.Intn(999))),
			DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", rand.Intn(999))),
			Strategy:    auth0.String(ConnectionStrategyAD),
			EnabledClients: &[]string{
				os.Getenv("AUTH0_CLIENT_ID"),
				client.GetClientID(),
			},
		},
	})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
		ShowAsButton:            auth0.Bool(true),
	}

	err := api.Organization.AddConnection(context.Background(), orgID, orgConn)
	require.NoError(t, err)

	return orgConn
}

func givenAnOrganizationDBConnection(t *testing.T, orgID string) *OrganizationConnection {
	client := givenAClient(t)
	conn := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:        auth0.String(fmt.Sprintf("test-conn%v", rand.Intn(999))),
			DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", rand.Intn(999))),
			Strategy:    auth0.String(ConnectionStrategyAuth0),
			EnabledClients: &[]string{
				os.Getenv("AUTH0_CLIENT_ID"),
				client.GetClientID(),
			},
		},
	})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
		IsSignupEnabled:         auth0.Bool(true),
	}

	err := api.Organization.AddConnection(context.Background(), orgID, orgConn)
	require.NoError(t, err)

	return orgConn
}

func givenAnOrganizationConnectionWithoutOrgID(t *testing.T) *OrganizationConnection {
	client := givenAClient(t)
	conn := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:        auth0.String(fmt.Sprintf("test-conn%v", rand.Intn(999))),
			DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", rand.Intn(999))),
			Strategy:    auth0.String(ConnectionStrategyAD),
			EnabledClients: &[]string{
				os.Getenv("AUTH0_CLIENT_ID"),
				client.GetClientID(),
			},
		},
	})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
		ShowAsButton:            auth0.Bool(true),
	}

	return orgConn
}

func givenAnOrganizationDBConnectionWithoutOrgID(t *testing.T) *OrganizationConnection {
	client := givenAClient(t)
	conn := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:        auth0.String(fmt.Sprintf("test-conn%v", rand.Intn(999))),
			DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", rand.Intn(999))),
			Strategy:    auth0.String(ConnectionStrategyAuth0),
			EnabledClients: &[]string{
				os.Getenv("AUTH0_CLIENT_ID"),
				client.GetClientID(),
			},
		},
	})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
		IsSignupEnabled:         auth0.Bool(true),
	}

	return orgConn
}

func givenAnOrganizationInvitation(t *testing.T, orgID string) *OrganizationInvitation {
	t.Helper()

	client := givenAClient(t)
	orgInvite := &OrganizationInvitation{
		Inviter: &OrganizationInvitationInviter{
			Name: auth0.String("Test Inviter"),
		},
		Invitee: &OrganizationInvitationInvitee{
			Email: auth0.String("test@example.com"),
		},
		ClientID: client.ClientID,
	}

	err := api.Organization.CreateInvitation(context.Background(), orgID, orgInvite)
	require.NoError(t, err)

	return orgInvite
}

func cleanupOrganization(t *testing.T, orgID string) {
	t.Helper()

	err := api.Organization.Delete(context.Background(), orgID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
