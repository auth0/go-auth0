package authentication

import (
	"context"
	"fmt"
	"github.com/auth0/go-auth0/authentication/ciba"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCIBA_Initiate(t *testing.T) {
	configureHTTPTestRecordings(t, authAPI)

	user := givenAUser(t)

	// Call the Initiate method of the CIBA manager
	resp, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
		ClientID:       mgmtClientID,
		ClientSecret:   mgmtClientSecret,
		Scope:          "openid",
		LoginHint:      fmt.Sprintf(`{"format":"iss_sub","iss":"https://witty-silver-sailfish-sus1-staging-20240704.sus.auth0.com/","sub":"%s"}`, user.GetID()),
		BindingMessage: "TEST-BINDING-MESSAGE",
	})

	// Validate the response
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.AuthReqID, "auth_req_id should not be empty")
	assert.Greater(t, resp.ExpiresIn, int64(0), "expires_in should be greater than 0")
	assert.Greater(t, resp.Interval, int64(0), "interval should be greater than 0")
}

//func givenAUser(t *testing.T) *management.User {
//	t.Helper()
//
//	userMetadata := map[string]interface{}{
//		"favourite_attack": "roundhouse_kick",
//	}
//	appMetadata := map[string]interface{}{
//		"facts": []string{
//			"count_to_infinity_twice",
//			"kill_two_stones_with_one_bird",
//			"can_hear_sign_language",
//		},
//	}
//	user := &management.User{
//		Connection:    auth0.String("Username-Password-Authentication"),
//		Email:         auth0.String(fmt.Sprintf("chuck%d@example.com", rand.Intn(999))),
//		Password:      auth0.String("Passwords hide their chuck"),
//		Username:      auth0.String(fmt.Sprintf("test-user%d", rand.Intn(999))),
//		GivenName:     auth0.String("Chuck"),
//		FamilyName:    auth0.String("Sanchez"),
//		Nickname:      auth0.String("Chucky"),
//		UserMetadata:  &userMetadata,
//		EmailVerified: auth0.Bool(true),
//		VerifyEmail:   auth0.Bool(false),
//		AppMetadata:   &appMetadata,
//		Picture:       auth0.String("https://example-picture-url.jpg"),
//		Blocked:       auth0.Bool(false),
//	}
//
//	err := mgmtAPI.User.Create(context.Background(), user)
//	require.NoError(t, err)
//
//	t.Cleanup(func() {
//		err := mgmtAPI.User.Delete(context.Background(), user.GetID())
//		require.NoError(t, err)
//	})
//
//	return user
//}
