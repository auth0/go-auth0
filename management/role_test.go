package management

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestRoleManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	role := &Role{
		Name:        auth0.String("test-role"),
		Description: auth0.String("Test Role"),
	}

	err := api.Role.Create(context.Background(), role)

	assert.NoError(t, err)
	assert.NotEmpty(t, role.GetID())

	t.Cleanup(func() {
		cleanupRole(t, role.GetID())
	})
}

func TestRoleManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedRole := givenARole(t)

	actualRole, err := api.Role.Read(context.Background(), expectedRole.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, actualRole)
}

func TestRoleManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedRole := givenARole(t)

	updatedRole := &Role{
		Description: auth0.String("The Administrator"),
	}
	err := api.Role.Update(context.Background(), expectedRole.GetID(), updatedRole)

	assert.NoError(t, err)
	assert.Equal(t, "The Administrator", updatedRole.GetDescription())
	assert.Equal(t, expectedRole.GetName(), updatedRole.GetName())
}

func TestRoleManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedRole := givenARole(t)

	err := api.Role.Delete(context.Background(), expectedRole.GetID())
	assert.NoError(t, err)

	actualRole, err := api.Role.Read(context.Background(), expectedRole.GetID())
	assert.Empty(t, actualRole)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestRoleManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	role := givenARole(t)

	roleList, err := api.Role.List(context.Background(), Parameter("name_filter", role.GetName()))

	assert.NoError(t, err)
	assert.Len(t, roleList.Roles, 1)
	assert.Equal(t, role.GetID(), roleList.Roles[0].GetID())
}

func TestRoleManager_Users(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	role := givenARole(t)

	err := api.Role.AssignUsers(context.Background(), role.GetID(), []*User{user})
	assert.NoError(t, err)

	roleUsers, err := api.Role.Users(context.Background(), role.GetID())
	assert.NoError(t, err)
	assert.Len(t, roleUsers.Users, 1)
	assert.Equal(t, user.GetID(), roleUsers.Users[0].GetID())
}

func TestRoleManager_UsersCheckpointPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	users := make([]*User, 0)

	for i := 0; i < 3; i++ {
		user := givenAUser(t)
		users = append(users, user)
	}

	role := givenARole(t)

	err := api.Role.AssignUsers(context.Background(), role.GetID(), users)
	assert.NoError(t, err)

	roleUsers, err := api.Role.Users(context.Background(), role.GetID(), Take(2))
	assert.NoError(t, err)
	assert.Len(t, roleUsers.Users, 2)
	assert.True(t, roleUsers.HasNext())

	roleUsers, err = api.Role.Users(context.Background(), role.GetID(), Take(2), From(roleUsers.Next))
	assert.NoError(t, err)
	assert.Len(t, roleUsers.Users, 1)
	assert.False(t, roleUsers.HasNext())
}

func TestRoleManager_Permissions(t *testing.T) {
	configureHTTPTestRecordings(t)

	role := givenARole(t)
	resourceServer := givenAResourceServer(t)
	permissions := []*Permission{
		{
			Name:                     resourceServer.GetScopes()[0].Value,
			ResourceServerIdentifier: resourceServer.Identifier,
		},
	}

	err := api.Role.AssociatePermissions(context.Background(), role.GetID(), permissions)
	assert.NoError(t, err)

	permissionList, err := api.Role.Permissions(context.Background(), role.GetID())
	assert.NoError(t, err)
	assert.Len(t, permissionList.Permissions, 1)
	assert.Equal(t, permissions[0].GetName(), permissionList.Permissions[0].GetName())
	assert.Equal(t, permissions[0].GetResourceServerIdentifier(), permissionList.Permissions[0].GetResourceServerIdentifier())

	err = api.Role.RemovePermissions(context.Background(), role.GetID(), permissions)
	assert.NoError(t, err)

	permissionList, err = api.Role.Permissions(context.Background(), role.GetID())
	assert.NoError(t, err)
	assert.Len(t, permissionList.Permissions, 0)
}

func givenARole(t *testing.T) *Role {
	t.Helper()

	role := &Role{
		Name:        auth0.String(fmt.Sprintf("test-role%d", rand.Intn(999))),
		Description: auth0.String("Test Role"),
	}

	err := api.Role.Create(context.Background(), role)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupRole(t, role.GetID())
	})

	return role
}

func cleanupRole(t *testing.T, roleID string) {
	t.Helper()

	err := api.Role.Delete(context.Background(), roleID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
