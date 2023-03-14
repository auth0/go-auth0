package management

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestOrganizationManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := &Organization{
		Name:        auth0.String(fmt.Sprintf("test-organization%v", rand.Intn(999))),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoURL: auth0.String("https://example.com/logo.gif"),
		},
	}

	err := api.Organization.Create(org)
	assert.NoError(t, err)
	assert.NotEmpty(t, org.GetID())

	t.Cleanup(func() {
		cleanupOrganization(t, org.GetID())
	})
}

func TestOrganizationManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	actualOrg, err := api.Organization.Read(org.GetID())

	assert.NoError(t, err)
	assert.Equal(t, org, actualOrg)
}

func TestOrganizationManager_ReadByName(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	actualOrg, err := api.Organization.ReadByName(org.GetName())

	assert.NoError(t, err)
	assert.Equal(t, org, actualOrg)
}

func TestOrganizationManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	err := api.Organization.Update(org.GetID(), &Organization{Name: auth0.String("new-org-name")})
	assert.NoError(t, err)

	actualOrg, err := api.Organization.Read(org.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "new-org-name", actualOrg.GetName())
}

func TestOrganizationManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	err := api.Organization.Delete(org.GetID())
	assert.NoError(t, err)

	actualOrg, err := api.Organization.Read(org.GetID())
	assert.Empty(t, actualOrg)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenAnOrganization(t)

	orgList, err := api.Organization.List()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(orgList.Organizations), 1)
}

func TestOrganizationManager_ListCheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	for i := 0; i < 3; i++ {
		givenAnOrganization(t)
	}

	orgList, err := api.Organization.List(Take(2))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(orgList.Organizations), 2)
	assert.True(t, orgList.HasNext())

	orgList, err = api.Organization.List(Take(2), From(orgList.Next))
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
		Strategy:    auth0.String(ConnectionStrategyAuth0),
		EnabledClients: &[]string{
			os.Getenv("AUTH0_CLIENT_ID"),
			client.GetClientID(),
		},
	}})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
	}

	err := api.Organization.AddConnection(org.GetID(), orgConn)
	assert.NoError(t, err)
	assert.NotEmpty(t, orgConn.GetConnection())
}

func TestOrganizationManager_Connection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	actualOrgConn, err := api.Organization.Connection(org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, orgConn, actualOrgConn)
}

func TestOrganizationManager_UpdateConnection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	err := api.Organization.UpdateConnection(
		org.GetID(),
		orgConn.GetConnectionID(),
		&OrganizationConnection{
			AssignMembershipOnLogin: auth0.Bool(false),
		},
	)
	assert.NoError(t, err)

	actualOrgConn, err := api.Organization.Connection(org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, false, actualOrgConn.GetAssignMembershipOnLogin())
}

func TestOrganizationManager_DeleteConnection(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	err := api.Organization.DeleteConnection(org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)

	actualOrgConn, err := api.Organization.Connection(org.GetID(), orgConn.GetConnectionID())
	assert.Error(t, err)
	assert.Empty(t, actualOrgConn)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_Connections(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	orgConnList, err := api.Organization.Connections(org.GetID())
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

	err := api.Organization.CreateInvitation(org.GetID(), orgInvite)
	assert.NoError(t, err)
	assert.NotEmpty(t, orgInvite.GetID())
}

func TestOrganizationManager_Invitation(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	actualOrgInvite, err := api.Organization.Invitation(org.GetID(), orgInvite.GetID())
	assert.NoError(t, err)
	assert.Equal(t, orgInvite, actualOrgInvite)
}

func TestOrganizationManager_DeleteInvitation(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	err := api.Organization.DeleteInvitation(org.GetID(), orgInvite.GetID())
	assert.NoError(t, err)

	actualOrgInvite, err := api.Organization.Invitation(org.GetID(), orgInvite.GetID())
	assert.Error(t, err)
	assert.Empty(t, actualOrgInvite)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_Invitations(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	invitations, err := api.Organization.Invitations(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, invitations.OrganizationInvitations, 1)
	assert.Equal(t, orgInvite.GetID(), invitations.OrganizationInvitations[0].GetID())
}

func TestOrganizationManager_AddMembers(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := api.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)
}

func TestOrganizationManager_DeleteMembers(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := api.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = api.Organization.DeleteMember(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	members, err := api.Organization.Members(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, members.Members, 0)
}

func TestOrganizationManager_Members(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := api.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	members, err := api.Organization.Members(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, members.Members, 1)
	assert.Equal(t, user.GetID(), members.Members[0].GetUserID())
}

func TestOrganizationManager_MembersCheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	users := make([]string, 0)

	for i := 0; i < 3; i++ {
		user := givenAUser(t)
		users = append(users, user.GetID())
	}

	err := api.Organization.AddMembers(org.GetID(), users)
	assert.NoError(t, err)

	members, err := api.Organization.Members(org.GetID(), Take(2))
	assert.NoError(t, err)
	assert.Len(t, members.Members, 2)
	assert.True(t, members.HasNext())

	members, err = api.Organization.Members(org.GetID(), Take(2), From(members.Next))
	assert.NoError(t, err)
	assert.Len(t, members.Members, 1)
	assert.True(t, members.HasNext())

	// Org members pagination will only return an empty `Next` value when the number of members
	// returned is 0 unlike other pagination APIs
	members, err = api.Organization.Members(org.GetID(), Take(2), From(members.Next))
	assert.NoError(t, err)
	assert.Len(t, members.Members, 0)
	assert.False(t, members.HasNext())
}

func TestOrganizationManager_MemberRoles(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)
	user := givenAUser(t)
	role := givenARole(t)

	err := api.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = api.Organization.AssignMemberRoles(org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	roles, err := api.Organization.MemberRoles(org.GetID(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 1)
	assert.Equal(t, role.GetID(), roles.Roles[0].GetID())

	err = api.Organization.DeleteMemberRoles(org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	roles, err = api.Organization.MemberRoles(org.GetID(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 0)
}

func givenAnOrganization(t *testing.T) *Organization {
	org := &Organization{
		Name:        auth0.String(fmt.Sprintf("test-organization%v", rand.Intn(999))),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoURL: auth0.String("https://example.com/logo.gif"),
		},
	}

	err := api.Organization.Create(org)
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
	}

	err := api.Organization.AddConnection(orgID, orgConn)
	require.NoError(t, err)

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

	err := api.Organization.CreateInvitation(orgID, orgInvite)
	require.NoError(t, err)

	return orgInvite
}

func cleanupOrganization(t *testing.T, orgID string) {
	t.Helper()

	err := api.Organization.Delete(orgID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
