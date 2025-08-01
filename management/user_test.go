package management

import (
	"context"
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

	err := api.User.Create(context.Background(), user)

	assert.NoError(t, err)
	assert.NotEmpty(t, user.GetID())

	defer cleanupUser(t, user.GetID())
}

func TestUserManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	actualUser, err := api.User.Read(context.Background(), expectedUser.GetID())

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
	err := api.User.Update(context.Background(), expectedUser.GetID(), actualUser)

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

	err := api.User.Delete(context.Background(), expectedUser.GetID())

	assert.NoError(t, err)

	actualUser, err := api.User.Read(context.Background(), expectedUser.GetID())

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
	userList, err := api.User.List(context.Background(), Query(userQuery))

	assert.NoError(t, err)
	assert.Len(t, userList.Users, 1)
}

func TestUserManager_Search(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	// Giving the search enough time to be
	// able to pick up the created user.
	time.Sleep(time.Second)

	userList, err := api.User.Search(context.Background(), Query(fmt.Sprintf("email:%q", expectedUser.GetEmail())))

	assert.NoError(t, err)
	assert.Len(t, userList.Users, 1)
}

func TestUserManager_ListByEmail(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUser := givenAUser(t)

	// Giving the search enough time to be
	// able to pick up the created user.
	time.Sleep(time.Second)

	users, err := api.User.ListByEmail(context.Background(), expectedUser.GetEmail())

	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestUserManager_Roles(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	role := givenARole(t)

	err := api.User.AssignRoles(context.Background(), user.GetID(), []*Role{role})
	assert.NoError(t, err)

	roles, err := api.User.Roles(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 1)
	assert.Equal(t, role.GetName(), roles.Roles[0].GetName())

	err = api.User.RemoveRoles(context.Background(), user.GetID(), []*Role{role})
	assert.NoError(t, err)

	roles, err = api.User.Roles(context.Background(), user.GetID())
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

	err := api.User.AssignPermissions(context.Background(), user.GetID(), permissions)
	assert.NoError(t, err)

	permissionList, err := api.User.Permissions(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, permissionList.Permissions, 1)
	assert.Equal(t, permissions[0].GetName(), permissionList.Permissions[0].GetName())
	assert.Equal(t, permissions[0].GetResourceServerIdentifier(), permissionList.Permissions[0].GetResourceServerIdentifier())

	err = api.User.RemovePermissions(context.Background(), user.GetID(), permissions)
	assert.NoError(t, err)

	permissionList, err = api.User.Permissions(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, permissionList.Permissions, 0)
}

func TestUserManager_Blocks(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	blockedIPs, err := api.User.Blocks(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, blockedIPs, 0)
}

func TestUserManager_BlocksByIdentifier(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	blockedIPs, err := api.User.BlocksByIdentifier(context.Background(), user.GetUsername())
	assert.NoError(t, err)
	assert.Len(t, blockedIPs, 0)
}

func TestUserManager_Unblock(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.Unblock(context.Background(), user.GetID())
	assert.NoError(t, err)
}

func TestUserManager_UnblockByIdentifier(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.UnblockByIdentifier(context.Background(), user.GetUsername())
	assert.NoError(t, err)
}

func TestUserManager_Enrollments(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	userEnrollments, err := api.User.Enrollments(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, userEnrollments, 0)
}

func TestUserManager_RegenerateRecoveryCode(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	recoveryCode, err := api.User.RegenerateRecoveryCode(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.NotEmpty(t, recoveryCode)
}

func TestUserManager_InvalidateRememberBrowser(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.InvalidateRememberBrowser(context.Background(), user.GetID())
	assert.NoError(t, err)
}

func TestUserManager_Link(t *testing.T) {
	configureHTTPTestRecordings(t)

	mainUser := givenAUser(t)
	secondaryUser := givenAUser(t)
	conn, err := api.Connection.ReadByName(context.Background(), "Username-Password-Authentication")
	assert.NoError(t, err)

	mainUserIdentities, err := api.User.Link(
		context.Background(),
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
	conn, err := api.Connection.ReadByName(context.Background(), "Username-Password-Authentication")
	assert.NoError(t, err)

	_, err = api.User.Link(
		context.Background(),
		mainUser.GetID(),
		&UserIdentityLink{
			Provider:     &provider,
			UserID:       secondaryUser.ID,
			ConnectionID: conn.ID,
		},
	)
	assert.NoError(t, err)

	unlinkedIdentities, err := api.User.Unlink(
		context.Background(),
		mainUser.GetID(),
		provider,
		strings.TrimPrefix(secondaryUser.GetID(), "auth0|"),
	)
	assert.NoError(t, err)
	assert.Len(t, unlinkedIdentities, 1)
	assert.Equal(t, mainUser.GetID(), "auth0|"+unlinkedIdentities[0].GetUserID())
}

func TestUserManager_ListUserSessions(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	sessions, err := api.User.ListUserSessions(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, sessions.Sessions, 0)
}

func TestUserManager_DeleteUserSessions(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	err := api.User.DeleteUserSessions(context.Background(), user.GetID())
	assert.NoError(t, err)
}

func TestUserManager_GetUserLogs(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	logs, err := api.User.GetUserLogs(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, logs, 0)
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

	err := api.User.CreateAuthenticationMethod(context.Background(), user.GetID(), &method)
	assert.NoError(t, err)

	methods, err := api.User.ListAuthenticationMethods(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 1)
	assert.Equal(t, method.GetID(), methods.Authenticators[0].GetID())

	methodInfo, err := api.User.GetAuthenticationMethodByID(context.Background(), user.GetID(), method.GetID())
	assert.NoError(t, err)
	assert.Equal(t, method.GetID(), methodInfo.GetID())

	err = api.User.UpdateAuthenticationMethod(context.Background(), user.GetID(), methodInfo.GetID(), &AuthenticationMethod{
		Name: auth0.String("Test2"),
	})
	assert.NoError(t, err)

	methodInfo, err = api.User.GetAuthenticationMethodByID(context.Background(), user.GetID(), method.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "Test2", methodInfo.GetName())

	err = api.User.DeleteAuthenticationMethod(context.Background(), user.GetID(), method.GetID())
	assert.NoError(t, err)

	methods, err = api.User.ListAuthenticationMethods(context.Background(), user.GetID())
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

	err = api.User.UpdateAllAuthenticationMethods(context.Background(), user.GetID(), &updateMethods)
	assert.NoError(t, err)

	methods, err = api.User.ListAuthenticationMethods(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 3)

	err = api.User.DeleteAllAuthenticationMethods(context.Background(), user.GetID())
	assert.NoError(t, err)

	methods, err = api.User.ListAuthenticationMethods(context.Background(), user.GetID())
	assert.NoError(t, err)
	assert.Len(t, methods.Authenticators, 0)
}

func TestUserManager_Organizations(t *testing.T) {
	configureHTTPTestRecordings(t)

	user := givenAUser(t)
	org := givenAnOrganization(t)

	err := api.Organization.AddMembers(context.Background(), org.GetID(), []string{user.GetID()})
	require.NoError(t, err)

	orgs, err := api.User.Organizations(context.Background(), user.GetID())
	require.NoError(t, err)
	assert.Len(t, orgs.Organizations, 1)
	assert.Equal(t, org.GetID(), orgs.Organizations[0].GetID())
}

// TestUserManager_ListRefreshTokens tests the ListRefreshTokens method of UserManager.
// This E2E test is skipped because refresh tokens cannot be created without UI interaction.
func TestUserManager_ListRefreshTokens(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	// RecordingNote: This test recording was manually generated to match these details.
	// If any changes occur here, the test recording will need manual modification.
	user := &User{ID: auth0.String("UserID")}
	expectedToken1 := RefreshToken{
		ID:        auth0.String("RefreshTokenID"),
		UserID:    auth0.String("UserID"),
		CreatedAt: auth0.Time(time.Date(2024, 5, 1, 13, 0, 30, 38000000, time.UTC)),
		ClientID:  auth0.String("CLIENTID"),
		Rotating:  auth0.Bool(false),
		ResourceServer: []*RefreshTokenResourceServer{
			{
				Audience: auth0.String("https://go-auth0-dev.eu.auth0.com.us.auth0.com/api/v2/"),
				Scopes:   auth0.String("openid profile offline_access"),
			},
		},
	}
	expectedToken2 := RefreshToken{
		ID:        auth0.String("RefreshTokenID"),
		UserID:    auth0.String("UserID"),
		CreatedAt: auth0.Time(time.Date(2024, 5, 3, 11, 58, 27, 35000000, time.UTC)),
		ClientID:  auth0.String("CLIENTID"),
		Rotating:  auth0.Bool(false),
		ResourceServer: []*RefreshTokenResourceServer{
			{
				Audience: auth0.String("https://go-auth0-dev.eu.auth0.com.us.auth0.com/api/v2/"),
				Scopes:   auth0.String("openid profile email address phone delete:current_user_device_credentials create:current_user_device_credentials offline_access"),
			},
		},
	}

	expectedTokens := []*RefreshToken{&expectedToken1, &expectedToken2}

	tokens, err := api.User.ListRefreshTokens(context.Background(), user.GetID())
	require.NoError(t, err)
	assert.Equal(t, expectedTokens, tokens.Tokens)
	assert.Equal(t, "RefreshTokenID", tokens.Next)
}

// TestUserManager_DeleteRefreshTokens tests the DeleteRefreshTokens method of UserManager.
// This E2E test is skipped because refresh tokens cannot be created without UI interaction.
func TestUserManager_DeleteRefreshTokens(t *testing.T) {
	skipTestIfRunningE2E(t)
	configureHTTPTestRecordings(t)

	// RecordingNote: This test recording was manually generated to match these details.
	// If any changes occur here, the test recording will need manual modification.
	user := &User{ID: auth0.String("UserID")}
	expectedToken1 := RefreshToken{
		ID:        auth0.String("RefreshTokenID"),
		UserID:    auth0.String("UserID"),
		CreatedAt: auth0.Time(time.Date(2024, 5, 1, 13, 0, 30, 38000000, time.UTC)),
		ClientID:  auth0.String("CLIENTID"),
		Rotating:  auth0.Bool(false),
		ResourceServer: []*RefreshTokenResourceServer{
			{
				Audience: auth0.String("https://go-auth0-dev.eu.auth0.com.us.auth0.com/api/v2/"),
				Scopes:   auth0.String("openid profile offline_access"),
			},
		},
	}
	expectedToken2 := RefreshToken{
		ID:        auth0.String("RefreshTokenID"),
		UserID:    auth0.String("UserID"),
		CreatedAt: auth0.Time(time.Date(2024, 5, 3, 11, 58, 27, 35000000, time.UTC)),
		ClientID:  auth0.String("CLIENTID"),
		Rotating:  auth0.Bool(false),
		ResourceServer: []*RefreshTokenResourceServer{
			{
				Audience: auth0.String("https://go-auth0-dev.eu.auth0.com.us.auth0.com/api/v2/"),
				Scopes:   auth0.String("openid profile email address phone delete:current_user_device_credentials create:current_user_device_credentials offline_access"),
			},
		},
	}
	expectedTokens := []*RefreshToken{&expectedToken1, &expectedToken2}

	tokens := retrieveRefreshTokens(t)
	assert.Equal(t, expectedTokens, tokens.Tokens)

	err := api.User.DeleteRefreshTokens(context.Background(), user.GetID())
	require.NoError(t, err)

	tokensAfterDeletion := retrieveRefreshTokens(t)
	assert.Empty(t, tokensAfterDeletion.Tokens)
	assert.Empty(t, tokensAfterDeletion.Next)
}

func TestUserManager_ClearRiskAssessmentAssessors(t *testing.T) {
	configureHTTPTestRecordings(t)
	user := givenAUser(t)
	Assessor := &UserRiskAssessmentAssessor{
		Connection: auth0.String("Username-Password-Authentication"),
		Assessors:  []string{"new-device"},
	}

	err := api.User.ClearRiskAssessmentAssessors(context.Background(), user.GetID(), Assessor)

	require.NoError(t, err)
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

	err := api.User.Create(context.Background(), user)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupUser(t, user.GetID())
	})

	return user
}

// retrieveRefreshTokens retrieves refresh tokens associated with a user.
//
// This function is responsible for fetching refresh tokens from the user.
// It does not create new refresh tokens but rather retrieves existing ones.
func retrieveRefreshTokens(t *testing.T) *RefreshTokenList {
	t.Helper()

	user := &User{ID: auth0.String("UserID")}

	tokens, err := api.User.ListRefreshTokens(context.Background(), user.GetID())
	require.NoError(t, err)

	return tokens
}

func cleanupUser(t *testing.T, userID string) {
	t.Helper()

	err := api.User.Delete(context.Background(), userID)
	require.NoError(t, err)
}

// skipTestIfRunningE2E skips the test if running in an end-to-end (E2E) scenario.
//
// This function is used to skip a test if it's being executed in an end-to-end (E2E) scenario
// where HTTP recordings are not enabled.
func skipTestIfRunningE2E(t *testing.T) {
	t.Helper()

	if !httpRecordingsEnabled {
		t.Skip("Skipped due to inability of setting this up for an E2E scenario")
	}
}
