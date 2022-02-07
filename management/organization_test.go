package management

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/auth0/go-auth0"
)

func TestOrganization(t *testing.T) {
	var err error

	ts := time.Now().Format("20060102150405")

	client := &Client{
		Name:              auth0.Stringf("testclient%v", ts),
		AppType:           auth0.String("regular_web"),
		GrantTypes:        []interface{}{"client_credentials"},
		OrganizationUsage: auth0.String("allow"),
	}
	err = m.Client.Create(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.Client.Delete(auth0.StringValue(client.ClientID))
	})

	connection := &Connection{
		Name:        auth0.String(fmt.Sprintf("testconn%v", ts)),
		DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", ts)),
		Strategy:    auth0.String(ConnectionStrategyAuth0),
		EnabledClients: []interface{}{
			os.Getenv("AUTH0_CLIENT_ID"),
			client.ClientID,
		},
	}
	err = m.Connection.Create(connection)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.Connection.Delete(connection.GetID())
	})

	user := &User{
		Connection: connection.Name,
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("Passwords hide their Chuck"),
		GivenName:  auth0.String("Chuck"),
		FamilyName: auth0.String("Norris"),
	}
	err = m.User.Create(user)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.User.Delete(user.GetID())
	})

	role := &Role{
		Name:        auth0.String("admin"),
		Description: auth0.String("Administrator"),
	}
	err = m.Role.Create(role)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.Role.Delete(role.GetID())
	})

	o := &Organization{
		Name:        auth0.String(fmt.Sprintf("testorganization%v", ts)),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoURL: auth0.String("https://example.com/logo.gif"),
		},
	}

	oi := &OrganizationInvitation{
		Inviter: &OrganizationInvitationInviter{
			Name: auth0.String("Test Inviter"),
		},
		Invitee: &OrganizationInvitationInvitee{
			Email: auth0.String("test@example.com"),
		},
		ClientID: client.ClientID,
	}

	oc := &OrganizationConnection{
		ConnectionID:            connection.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
	}

	t.Run("Create", func(t *testing.T) {
		err = m.Organization.Create(o)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("List", func(t *testing.T) {
		var ol *OrganizationList
		ol, err = m.Organization.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ol)
	})

	t.Run("Read", func(t *testing.T) {
		o, err = m.Organization.Read(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("Update", func(t *testing.T) {
		id := o.GetID()
		o.ID = nil
		o.Name = auth0.Stringf("testorg%v", ts)
		err = m.Organization.Update(id, o)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("ReadByName", func(t *testing.T) {
		o, err = m.Organization.ReadByName(o.GetName())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("AddConnection", func(t *testing.T) {
		err = m.Organization.AddConnection(o.GetID(), oc)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("Connections", func(t *testing.T) {
		l, err := m.Organization.Connections(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("Connection", func(t *testing.T) {
		oc, err = m.Organization.Connection(o.GetID(), oc.GetConnectionID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("UpdateConnection", func(t *testing.T) {
		connectionID := oc.GetConnectionID()
		err = m.Organization.UpdateConnection(o.GetID(), connectionID, &OrganizationConnection{
			AssignMembershipOnLogin: auth0.Bool(false),
		})
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("CreateInvitation", func(t *testing.T) {
		err = m.Organization.CreateInvitation(o.GetID(), oi)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oi)
	})

	t.Run("Invitations", func(t *testing.T) {
		l, err := m.Organization.Invitations(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("Invitation", func(t *testing.T) {
		var i *OrganizationInvitation
		i, err = m.Organization.Invitation(o.GetID(), oi.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", i)
	})

	t.Run("AddMembers", func(t *testing.T) {
		err = m.Organization.AddMembers(o.GetID(), []string{user.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Members", func(t *testing.T) {
		l, err := m.Organization.Members(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("MemberRoles", func(t *testing.T) {
		l, err := m.Organization.MemberRoles(o.GetID(), user.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("AssignMemberRoles", func(t *testing.T) {
		err = m.Organization.AssignMemberRoles(o.GetID(), user.GetID(), []string{role.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteMemberRoles", func(t *testing.T) {
		err = m.Organization.DeleteMemberRoles(o.GetID(), user.GetID(), []string{role.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteMember", func(t *testing.T) {
		err = m.Organization.DeleteMember(o.GetID(), []string{user.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteConnection", func(t *testing.T) {
		err = m.Organization.DeleteConnection(o.GetID(), oc.GetConnectionID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("DeleteInvitation", func(t *testing.T) {
		err = m.Organization.DeleteInvitation(o.GetID(), oi.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oi)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Organization.Delete(o.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}
