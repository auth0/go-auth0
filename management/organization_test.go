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
	org := &Organization{
		Name:        auth0.String(fmt.Sprintf("test-organization%v", rand.Intn(999))),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoURL: auth0.String("https://example.com/logo.gif"),
		},
	}

	err := m.Organization.Create(org)
	assert.NoError(t, err)
	assert.NotEmpty(t, org.GetID())

	t.Cleanup(func() {
		cleanupOrganization(t, org.GetID())
	})
}

func TestOrganizationManager_Read(t *testing.T) {
	org := givenAnOrganization(t)

	actualOrg, err := m.Organization.Read(org.GetID())

	assert.NoError(t, err)
	assert.Equal(t, org, actualOrg)
}

func TestOrganizationManager_ReadByName(t *testing.T) {
	org := givenAnOrganization(t)

	actualOrg, err := m.Organization.ReadByName(org.GetName())

	assert.NoError(t, err)
	assert.Equal(t, org, actualOrg)
}

func TestOrganizationManager_Update(t *testing.T) {
	org := givenAnOrganization(t)

	err := m.Organization.Update(org.GetID(), &Organization{Name: auth0.String("new-org-name")})
	assert.NoError(t, err)

	actualOrg, err := m.Organization.Read(org.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "new-org-name", actualOrg.GetName())
}

func TestOrganizationManager_Delete(t *testing.T) {
	org := givenAnOrganization(t)

	err := m.Organization.Delete(org.GetID())
	assert.NoError(t, err)

	actualOrg, err := m.Organization.Read(org.GetID())
	assert.Empty(t, actualOrg)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_List(t *testing.T) {
	givenAnOrganization(t)

	orgList, err := m.Organization.List()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(orgList.Organizations), 1)
}

func TestOrganizationManager_AddConnection(t *testing.T) {
	org := givenAnOrganization(t)
	client := givenAClient(t)
	conn := givenAConnection(t, connectionTestCase{connection: Connection{
		Name:        auth0.String(fmt.Sprintf("test-conn%v", rand.Intn(999))),
		DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", rand.Intn(999))),
		Strategy:    auth0.String(ConnectionStrategyAuth0),
		EnabledClients: []interface{}{
			os.Getenv("AUTH0_CLIENT_ID"),
			client.ClientID,
		},
	}})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
	}

	err := m.Organization.AddConnection(org.GetID(), orgConn)
	assert.NoError(t, err)
	assert.NotEmpty(t, orgConn.GetConnection())
}

func TestOrganizationManager_Connection(t *testing.T) {
	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	actualOrgConn, err := m.Organization.Connection(org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, orgConn, actualOrgConn)
}

func TestOrganizationManager_UpdateConnection(t *testing.T) {
	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	err := m.Organization.UpdateConnection(
		org.GetID(),
		orgConn.GetConnectionID(),
		&OrganizationConnection{
			AssignMembershipOnLogin: auth0.Bool(false),
		},
	)
	assert.NoError(t, err)

	actualOrgConn, err := m.Organization.Connection(org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, false, actualOrgConn.GetAssignMembershipOnLogin())
}

func TestOrganizationManager_DeleteConnection(t *testing.T) {
	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	err := m.Organization.DeleteConnection(org.GetID(), orgConn.GetConnectionID())
	assert.NoError(t, err)

	actualOrgConn, err := m.Organization.Connection(org.GetID(), orgConn.GetConnectionID())
	assert.Error(t, err)
	assert.Empty(t, actualOrgConn)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_Connections(t *testing.T) {
	org := givenAnOrganization(t)
	orgConn := givenAnOrganizationConnection(t, org.GetID())

	orgConnList, err := m.Organization.Connections(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, orgConnList.OrganizationConnections, 1)
	assert.Equal(t, orgConn.GetConnectionID(), orgConnList.OrganizationConnections[0].GetConnectionID())
}

func TestOrganizationManager_CreateInvitation(t *testing.T) {
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

	res, err := m.Organization.CreateInvitation(org.GetID(), orgInvite)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, orgInvite.GetID())
}

func TestOrganizationManager_Invitation(t *testing.T) {
	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	actualOrgInvite, err := m.Organization.Invitation(org.GetID(), orgInvite.GetID())
	assert.NoError(t, err)
	assert.Equal(t, orgInvite, actualOrgInvite)
}

func TestOrganizationManager_DeleteInvitation(t *testing.T) {
	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	err := m.Organization.DeleteInvitation(org.GetID(), orgInvite.GetID())
	assert.NoError(t, err)

	actualOrgInvite, err := m.Organization.Invitation(org.GetID(), orgInvite.GetID())
	assert.Error(t, err)
	assert.Empty(t, actualOrgInvite)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestOrganizationManager_Invitations(t *testing.T) {
	org := givenAnOrganization(t)
	orgInvite := givenAnOrganizationInvitation(t, org.GetID())

	invitations, err := m.Organization.Invitations(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, invitations.OrganizationInvitations, 1)
	assert.Equal(t, orgInvite.GetID(), invitations.OrganizationInvitations[0].GetID())
}

func TestOrganizationManager_AddMembers(t *testing.T) {
	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := m.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)
}

func TestOrganizationManager_DeleteMembers(t *testing.T) {
	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := m.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = m.Organization.DeleteMember(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	members, err := m.Organization.Members(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, members.Members, 0)
}

func TestOrganizationManager_Members(t *testing.T) {
	org := givenAnOrganization(t)
	user := givenAUser(t)

	err := m.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	members, err := m.Organization.Members(org.GetID())
	assert.NoError(t, err)
	assert.Len(t, members.Members, 1)
	assert.Equal(t, user.GetID(), members.Members[0].GetUserID())
}

func TestOrganizationManager_MemberRoles(t *testing.T) {
	org := givenAnOrganization(t)
	user := givenAUser(t)
	role := givenARole(t)

	err := m.Organization.AddMembers(org.GetID(), []string{user.GetID()})
	assert.NoError(t, err)

	err = m.Organization.AssignMemberRoles(org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	roles, err := m.Organization.MemberRoles(org.GetID(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 1)
	assert.Equal(t, role.GetID(), roles.Roles[0].GetID())

	err = m.Organization.DeleteMemberRoles(org.GetID(), user.GetID(), []string{role.GetID()})
	assert.NoError(t, err)

	roles, err = m.Organization.MemberRoles(org.GetID(), user.GetID())
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

	err := m.Organization.Create(org)
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
			EnabledClients: []interface{}{
				os.Getenv("AUTH0_CLIENT_ID"),
				client.ClientID,
			},
		},
	})
	orgConn := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
	}

	err := m.Organization.AddConnection(orgID, orgConn)
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

	res, err := m.Organization.CreateInvitation(orgID, orgInvite)
	require.NoError(t, err)
	assert.NotEmpty(t, res)
	return orgInvite
}

func cleanupOrganization(t *testing.T, orgID string) {
	t.Helper()

	err := m.Organization.Delete(orgID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
