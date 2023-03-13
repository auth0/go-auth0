package management

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestUserManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := &User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("chuck@example.com"),
		Username:   auth0.String("chuck"),
		Password:   auth0.String("I have a password and its a secret"),
	}

	err := api.User.Create(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, user.GetID())

	defer cleanupUser(t, user.GetID())
}

func TestUserManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	actualUser, err := api.User.Read(expectedUser.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.GetID(), actualUser.GetID())
	assert.Equal(t, expectedUser.GetName(), actualUser.GetName())
}

func TestUserManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	appMetadata := map[string]interface{}{"foo": "bar"}
	actualUser := &User{
		Connection:  auth0.String("Username-Password-Authentication"),
		Password:    auth0.String("I don't need one"),
		AppMetadata: &appMetadata,
	}
	err := api.User.Update(expectedUser.GetID(), actualUser)

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"foo": "bar",
		"facts": []interface{}{
			"count_to_infinity_twice",
			"kill_two_stones_with_one_bird",
			"can_hear_sign_language",
		},
	}, *actualUser.AppMetadata)
	assert.Equal(t, "Username-Password-Authentication", actualUser.GetConnection())
}

func TestUserManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	err := api.User.Delete(expectedUser.GetID())

	assert.NoError(t, err)

	actualUser, err := api.User.Read(expectedUser.GetID())

	assert.Empty(t, actualUser)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestUserManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	// The List() endpoint is slow to pick up the newly created user,
	// so we wait a second before executing the request.
	time.Sleep(time.Second)

	userQuery := fmt.Sprintf("username:%q", expectedUser.GetUsername())
	userList, err := api.User.List(Query(userQuery))

	assert.NoError(t, err)
	assert.Len(t, userList.Users, 1)
}

func TestUserManager_Search(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	// Giving the search enough time to be
	// able to pick up the created user.
	time.Sleep(time.Second)

	userList, err := api.User.Search(Query(fmt.Sprintf("email:%q", expectedUser.GetEmail())))

	assert.NoError(t, err)
	assert.Len(t, userList.Users, 1)
}

func TestUserManager_ListByEmail(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	// Giving the search enough time to be
	// able to pick up the created user.
	time.Sleep(time.Second)

	users, err := api.User.ListByEmail(expectedUser.GetEmail())

	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestUserManager_Roles(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	role := givenARole(t)

	err := api.User.AssignRoles(user.GetID(), []*Role{role})
	assert.NoError(t, err)

	roles, err := api.User.Roles(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 1)
	assert.Equal(t, role.GetName(), roles.Roles[0].GetName())

	err = api.User.RemoveRoles(user.GetID(), []*Role{role})
	assert.NoError(t, err)

	roles, err = api.User.Roles(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 0)
}

func TestUserManager_Permissions(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	resourceServer := givenAResourceServer(t)
	permissions := []*Permission{
		{
			Name:                     resourceServer.GetScopes()[0].Value,
			ResourceServerIdentifier: resourceServer.Identifier,
		},
	}

	err := api.User.AssignPermissions(user.GetID(), permissions)
	assert.NoError(t, err)

	permissionList, err := api.User.Permissions(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, permissionList.Permissions, 1)
	assert.Equal(t, permissions[0].GetName(), permissionList.Permissions[0].GetName())
	assert.Equal(t, permissions[0].GetResourceServerIdentifier(), permissionList.Permissions[0].GetResourceServerIdentifier())

	err = api.User.RemovePermissions(user.GetID(), permissions)
	assert.NoError(t, err)

	permissionList, err = api.User.Permissions(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, permissionList.Permissions, 0)
}

func TestUserManager_Blocks(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	blockedIPs, err := api.User.Blocks(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, blockedIPs, 0)
}

func TestUserManager_BlocksByIdentifier(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	blockedIPs, err := api.User.BlocksByIdentifier(user.GetUsername())
	assert.NoError(t, err)
	assert.Len(t, blockedIPs, 0)
}

func TestUserManager_Unblock(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.Unblock(user.GetID())
	assert.NoError(t, err)
}

func TestUserManager_UnblockByIdentifier(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.UnblockByIdentifier(user.GetUsername())
	assert.NoError(t, err)
}

func TestUserManager_Enrollments(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	userEnrollments, err := api.User.Enrollments(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, userEnrollments, 0)
}

func TestUserManager_RegenerateRecoveryCode(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	recoveryCode, err := api.User.RegenerateRecoveryCode(user.GetID())
	assert.NoError(t, err)
	assert.NotEmpty(t, recoveryCode)
}

func TestUserManager_InvalidateRememberBrowser(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.InvalidateRememberBrowser(user.GetID())
	assert.NoError(t, err)
}

func TestUserManager_Link(t *testing.T) {
	configureHTTPTestRecordings(t)

	mainUser := givenAUser(t)
	secondaryUser := givenAUser(t)
	conn, err := api.Connection.ReadByName("Username-Password-Authentication")
	assert.NoError(t, err)

	mainUserIdentities, err := api.User.Link(
		mainUser.GetID(),
		&UserIdentityLink{
			Provider:     auth0.String("auth0"),
			UserID:       secondaryUser.ID,
			ConnectionID: conn.ID,
		},
	)
	assert.NoError(t, err)
	assert.Len(t, mainUserIdentities, 2)
	assert.Equal(t, mainUser.GetID(), "auth0|"+mainUserIdentities[0].GetUserID())
	assert.Equal(t, secondaryUser.GetID(), "auth0|"+mainUserIdentities[1].GetUserID())
}

func TestUserManager_Unlink(t *testing.T) {
	configureHTTPTestRecordings(t)

	provider := "auth0"
	mainUser := givenAUser(t)
	secondaryUser := givenAUser(t)
	conn, err := api.Connection.ReadByName("Username-Password-Authentication")
	assert.NoError(t, err)

	_, err = api.User.Link(
		mainUser.GetID(),
		&UserIdentityLink{
			Provider:     &provider,
			UserID:       secondaryUser.ID,
			ConnectionID: conn.ID,
		},
	)
	assert.NoError(t, err)

	unlinkedIdentities, err := api.User.Unlink(
		mainUser.GetID(),
		provider,
		strings.TrimPrefix(secondaryUser.GetID(), "auth0|"),
	)
	assert.NoError(t, err)
	assert.Len(t, unlinkedIdentities, 1)
	assert.Equal(t, mainUser.GetID(), "auth0|"+unlinkedIdentities[0].GetUserID())
}

func TestUser_MarshalJSON(t *testing.T) {
	for user, expected := range map[*User]string{
		{}:                                 `{}`,
		{EmailVerified: auth0.Bool(true)}:  `{"email_verified":true}`,
		{EmailVerified: auth0.Bool(false)}: `{"email_verified":false}`,
	} {
		payload, err := json.Marshal(user)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func TestUser_UnmarshalJSON(t *testing.T) {
	for payload, expected := range map[string]*User{
		`{}`:                         {EmailVerified: nil},
		`{"email_verified":true}`:    {EmailVerified: auth0.Bool(true)},
		`{"email_verified":false}`:   {EmailVerified: auth0.Bool(false)},
		`{"email_verified":"true"}`:  {EmailVerified: auth0.Bool(true)},
		`{"email_verified":"false"}`: {EmailVerified: auth0.Bool(false)},
	} {
		var actual User
		err := json.Unmarshal([]byte(payload), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEmailVerified(), actual.GetEmailVerified())
	}
}

func TestUserIdentity_MarshalJSON(t *testing.T) {
	for userIdentity, expected := range map[*UserIdentity]string{
		{}:                            `{}`,
		{UserID: auth0.String("1")}:   `{"user_id":"1"}`,
		{UserID: auth0.String("foo")}: `{"user_id":"foo"}`,
	} {
		payload, err := json.Marshal(userIdentity)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func TestUserIdentity_UnmarshalJSON(t *testing.T) {
	for expectedAsString, expected := range map[string]*UserIdentity{
		`{}`:                {UserID: nil},
		`{"user_id":1}`:     {UserID: auth0.String("1")},
		`{"user_id":"1"}`:   {UserID: auth0.String("1")},
		`{"user_id":"foo"}`: {UserID: auth0.String("foo")},
		`{"profileData": {"picture": "some-picture.jpeg"}}`: {
			ProfileData: &map[string]interface{}{
				"picture": "some-picture.jpeg",
			},
		},
	} {
		var actual *UserIdentity
		err := json.Unmarshal([]byte(expectedAsString), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestUserManager_AuthenticationMethods(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	method := AuthenticationMethod{
		Name:  auth0.String("Test"),
		Type:  auth0.String("email"),
		Email: auth0.String(user.GetEmail()),
	}

	err := api.User.CreateAuthenticationMethod(user.GetID(), &method)
	assert.NoError(t, err)

	methods, err := api.User.ListAuthenticationMethods(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 1)
	assert.Equal(t, method.GetID(), methods.Authenticators[0].GetID())

	methodInfo, err := api.User.GetAuthenticationMethodByID(user.GetID(), method.GetID())
	assert.NoError(t, err)
	assert.Equal(t, method.GetID(), methodInfo.GetID())

	err = api.User.UpdateAuthenticationMethod(user.GetID(), methodInfo.GetID(), &AuthenticationMethod{
		Name: auth0.String("Test2"),
	})
	assert.NoError(t, err)

	methodInfo, err = api.User.GetAuthenticationMethodByID(user.GetID(), method.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "Test2", methodInfo.GetName())

	err = api.User.DeleteAuthenticationMethod(user.GetID(), method.GetID())
	assert.NoError(t, err)

	methods, err = api.User.ListAuthenticationMethods(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 0)

	updateMethods := []AuthenticationMethod{
		{
			Type:  auth0.String("email"),
			Name:  auth0.String("MyEmail"),
			Email: auth0.String("foo@example.com"),
		},
		{
			Type:                          auth0.String("phone"),
			Name:                          auth0.String("MyPhone"),
			PhoneNumber:                   auth0.String("+44207183875044"),
			PreferredAuthenticationMethod: auth0.String("sms"),
		},
		{
			Type:       auth0.String("totp"),
			Name:       auth0.String("MyTotp"),
			TOTPSecret: auth0.String("5HCWXIP2MJNSUBGYVUZFLRB2HWIGXR4SYJQXNBQ"),
		},
	}

	err = api.User.UpdateAllAuthenticationMethods(user.GetID(), &updateMethods)
	assert.NoError(t, err)

	methods, err = api.User.ListAuthenticationMethods(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 3)

	err = api.User.DeleteAllAuthenticationMethods(user.GetID())
	assert.NoError(t, err)

	methods, err = api.User.ListAuthenticationMethods(user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 0)
}

func givenAUser(t *testing.T) *User {
	t.Helper()

	userMetadata := map[string]interface{}{
		"favourite_attack": "roundhouse_kick",
	}
	appMetadata := map[string]interface{}{
		"facts": []string{
			"count_to_infinity_twice",
			"kill_two_stones_with_one_bird",
			"can_hear_sign_language",
		},
	}
	user := &User{
		Connection:    auth0.String("Username-Password-Authentication"),
		Email:         auth0.String(fmt.Sprintf("chuck%d@example.com", rand.Intn(999))),
		Password:      auth0.String("Passwords hide their chuck"),
		Username:      auth0.String(fmt.Sprintf("test-user%d", rand.Intn(999))),
		GivenName:     auth0.String("Chuck"),
		FamilyName:    auth0.String("Sanchez"),
		Nickname:      auth0.String("Chucky"),
		UserMetadata:  &userMetadata,
		EmailVerified: auth0.Bool(true),
		VerifyEmail:   auth0.Bool(false),
		AppMetadata:   &appMetadata,
		Picture:       auth0.String("https://example-picture-url.jpg"),
		Blocked:       auth0.Bool(false),
	}

	err := api.User.Create(user)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupUser(t, user.GetID())
	})

	return user
}

func cleanupUser(t *testing.T, userID string) {
	t.Helper()

	err := api.User.Delete(userID)
	require.NoError(t, err)
}
