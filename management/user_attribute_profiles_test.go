package management

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestUserAttributeProfileManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	t.Run("all fields", func(t *testing.T) {
		expectedUserAttributeProfile := &UserAttributeProfile{
			Name: auth0.Stringf("Test User Attribute Profile (%s)", time.Now().Format(time.StampMilli)),
			UserID: &UserAttributeProfileUserID{
				OIDCMapping: auth0.String("sub"),
				SAMLMapping: &[]string{"urn:oid:0.9.2342.19200300.100.1.1"},
				SCIMMapping: auth0.String("userName"),
			},
			UserAttributes: map[string]*UserAttributeProfileUserAttributes{
				"email": {
					Description:     auth0.String("User's email address"),
					Label:           auth0.String("Email"),
					ProfileRequired: auth0.Bool(true),
					Auth0Mapping:    auth0.String("email"),
					OIDCMapping: &UserAttributeProfileOIDCMapping{
						Mapping: auth0.String("email"),
					},
					SAMLMapping: &[]string{"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"},
					SCIMMapping: auth0.String("emails[primary eq true].value"),
				},
			},
		}

		err := api.UserAttributeProfile.Create(context.Background(), expectedUserAttributeProfile)
		assert.NoError(t, err)
		assert.NotEmpty(t, expectedUserAttributeProfile.GetID())

		t.Cleanup(func() {
			cleanupUserAttributeProfile(t, expectedUserAttributeProfile.GetID())
		})
	})

	t.Run("only required fields", func(t *testing.T) {
		expectedUserAttributeProfile := &UserAttributeProfile{
			Name: auth0.Stringf("Test User Attribute Profile (%s)", time.Now().Format(time.StampMilli)),
			UserAttributes: map[string]*UserAttributeProfileUserAttributes{
				"email": {
					Description:     auth0.String("User's email address"),
					Label:           auth0.String("Email"),
					ProfileRequired: auth0.Bool(true),
					Auth0Mapping:    auth0.String("email"),
					OIDCMapping: &UserAttributeProfileOIDCMapping{
						Mapping: auth0.String("email"),
					},
					SAMLMapping: &[]string{"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"},
					SCIMMapping: auth0.String("emails[primary eq true].value"),
				},
			},
		}

		err := api.UserAttributeProfile.Create(context.Background(), expectedUserAttributeProfile)
		assert.NoError(t, err)
		assert.NotEmpty(t, expectedUserAttributeProfile.GetID())

		t.Cleanup(func() {
			cleanupUserAttributeProfile(t, expectedUserAttributeProfile.GetID())
		})
	})
}

func TestUserAttributeProfileManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUserAttributeProfile := givenAUserAttributeProfile(t)
	actualUserAttributeProfile, err := api.UserAttributeProfile.Read(context.Background(), expectedUserAttributeProfile.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedUserAttributeProfile.GetID(), actualUserAttributeProfile.GetID())
	assert.Equal(t, expectedUserAttributeProfile.GetName(), actualUserAttributeProfile.GetName())
}

func TestUserAttributeProfileManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	t.Run("supports updating user_id", func(t *testing.T) {
		expectedUserAttributeProfile := givenAUserAttributeProfile(t)

		userAttributeProfileID := expectedUserAttributeProfile.GetID()

		expectedName := "Updated User Attribute Profile Name"
		expectedUserAttributeProfile.Name = &expectedName
		// SAMLMapping should remain the default
		expectedUserAttributeProfile.UserID = &UserAttributeProfileUserID{
			OIDCMapping: auth0.String("sub"),
			SCIMMapping: auth0.String("userName"),
		}

		err := api.UserAttributeProfile.Update(context.Background(), userAttributeProfileID, expectedUserAttributeProfile)
		assert.NoError(t, err)
		assert.Equal(t, expectedName, expectedUserAttributeProfile.GetName())

		// Verify the update was applied
		updatedProfile, err := api.UserAttributeProfile.Read(context.Background(), userAttributeProfileID)
		assert.NoError(t, err)
		assert.Equal(t, expectedName, updatedProfile.GetName())

		userID := updatedProfile.GetUserID()
		assert.NotNil(t, userID)
		assert.Equal(t, "sub", userID.GetOIDCMapping())
		assert.Equal(t, "userName", userID.GetSCIMMapping())
		assert.Len(t, userID.GetSAMLMapping(), 3)
	})

	t.Run("supports setting user_id to defaults", func(t *testing.T) {
		expectedUserAttributeProfile := givenAUserAttributeProfile(t)

		userAttributeProfileID := expectedUserAttributeProfile.GetID()

		expectedName := "Updated User Attribute Profile Name"
		expectedUserAttributeProfile.Name = &expectedName
		// doesn't clear the value but instead resets to defaults
		expectedUserAttributeProfile.UserID = nil

		err := api.UserAttributeProfile.Update(context.Background(), userAttributeProfileID, expectedUserAttributeProfile)
		assert.NoError(t, err)
		assert.Equal(t, expectedName, expectedUserAttributeProfile.GetName())

		// Verify the update was applied
		updatedProfile, err := api.UserAttributeProfile.Read(context.Background(), userAttributeProfileID)
		assert.NoError(t, err)
		assert.Equal(t, expectedName, updatedProfile.GetName())

		userID := updatedProfile.GetUserID()
		assert.NotNil(t, userID)
		assert.Equal(t, "sub", userID.GetOIDCMapping())
		assert.Equal(t, "externalId", userID.GetSCIMMapping())
		assert.Len(t, userID.GetSAMLMapping(), 3)
	})
}

func TestUserAttributeProfileManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUserAttributeProfile := givenAUserAttributeProfile(t)
	userAttributeProfileID := expectedUserAttributeProfile.GetID()

	err := api.UserAttributeProfile.Delete(context.Background(), userAttributeProfileID)
	assert.NoError(t, err)

	// Verify the profile has been deleted
	deletedProfile, err := api.UserAttributeProfile.Read(context.Background(), userAttributeProfileID)
	assert.Error(t, err)
	assert.Nil(t, deletedProfile)
}

func TestUserAttributeProfileManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedUserAttributeProfile := givenAUserAttributeProfile(t)

	userAttributeProfileList, err := api.UserAttributeProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(userAttributeProfileList.UserAttributeProfiles), 0)

	// Find our created profile in the list
	found := false

	for _, profile := range userAttributeProfileList.UserAttributeProfiles {
		if profile.GetID() == expectedUserAttributeProfile.GetID() {
			found = true

			assert.Equal(t, expectedUserAttributeProfile.GetName(), profile.GetName())

			break
		}
	}

	assert.True(t, found, "Created user attribute profile should be found in the list")
}

func TestUserAttributeProfileManager_ListTemplates(t *testing.T) {
	configureHTTPTestRecordings(t)

	templateList, err := api.UserAttributeProfile.ListTemplates(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, templateList)

	for _, template := range templateList.UserAttributeProfileTemplates {
		assert.NotEmpty(t, template.GetID())
		assert.NotEmpty(t, template.GetDisplayName())
		assert.NotNil(t, template.Template)
	}
}

func TestUserAttributeProfileManager_GetTemplate(t *testing.T) {
	configureHTTPTestRecordings(t)

	// First, get the list of templates to find a valid template ID
	templateList, err := api.UserAttributeProfile.ListTemplates(context.Background())
	assert.NoError(t, err)
	require.Greater(t, len(templateList.UserAttributeProfileTemplates), 0)

	templateID := templateList.UserAttributeProfileTemplates[0].GetID()

	template, err := api.UserAttributeProfile.GetTemplate(context.Background(), templateID)
	assert.NoError(t, err)
	assert.NotNil(t, template)
	assert.NotEmpty(t, template.GetDisplayName())
}

// Helper function to create a user attribute profile for testing.
func givenAUserAttributeProfile(t *testing.T) *UserAttributeProfile {
	t.Helper()

	userAttributeProfile := &UserAttributeProfile{
		Name: auth0.Stringf("Test User Attribute Profile (%s)", time.Now().Format(time.StampMilli)),
		UserID: &UserAttributeProfileUserID{
			OIDCMapping: auth0.String("sub"),
			SAMLMapping: &[]string{"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"},
			SCIMMapping: auth0.String("userName"),
		},
		UserAttributes: map[string]*UserAttributeProfileUserAttributes{
			"email": {
				Description:     auth0.String("User's email address"),
				Label:           auth0.String("Email"),
				ProfileRequired: auth0.Bool(true),
				Auth0Mapping:    auth0.String("email"),
				OIDCMapping: &UserAttributeProfileOIDCMapping{
					Mapping: auth0.String("email"),
				},
				SAMLMapping: &[]string{"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"},
				SCIMMapping: auth0.String("emails[primary eq true].value"),
			},
			"name": {
				Description:     auth0.String("User's full name"),
				Label:           auth0.String("Name"),
				ProfileRequired: auth0.Bool(false),
				Auth0Mapping:    auth0.String("name"),
				OIDCMapping: &UserAttributeProfileOIDCMapping{
					Mapping: auth0.String("email"),
				},
				SAMLMapping: &[]string{"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"},
				SCIMMapping: auth0.String("displayName"),
				StrategyOverrides: map[string]*UserAttributesStrategyOverride{
					"oidc": {
						SCIMMapping: auth0.String("name.givenName"),
					},
				},
			},
		},
	}

	err := api.UserAttributeProfile.Create(context.Background(), userAttributeProfile)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupUserAttributeProfile(t, userAttributeProfile.GetID())
	})

	return userAttributeProfile
}

// Helper function to clean up user attribute profiles.
func cleanupUserAttributeProfile(t *testing.T, userAttributeProfileID string) {
	t.Helper()

	err := api.UserAttributeProfile.Delete(context.Background(), userAttributeProfileID)
	require.NoError(t, err)
}
