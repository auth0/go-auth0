package management

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestTokenExchangeProfileManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)
	action := givenAnCustomTokenExchangeAction(t)
	tokenExchangeProfile := &TokenExchangeProfile{
		Name:             auth0.String("Test-Token-Exchange-Profile"),
		SubjectTokenType: auth0.String("https://acme.com/cis-token"),
		ActionID:         auth0.String(action.GetID()),
		Type:             auth0.String("custom_authentication"),
	}
	err := api.TokenExchangeProfile.Create(context.Background(), tokenExchangeProfile)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenExchangeProfile.GetID())

	tokenExchangeProfileRetrieved, err := api.TokenExchangeProfile.Read(context.Background(), tokenExchangeProfile.GetID())
	assert.NoError(t, err)
	assert.Equal(t, tokenExchangeProfileRetrieved, tokenExchangeProfile)

	t.Cleanup(func() {
		cleanupTokenExchangeProfile(t, tokenExchangeProfile.GetID())
		cleanupAction(t, action.GetID())
	})
}

func TestTokenExchangeProfileManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	tokenExchangeProfile := givenAnTokenExchangeProfile(t)

	tokenExchangeProfileList, err := api.TokenExchangeProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(tokenExchangeProfileList.TokenExchangeProfiles), 0)
	assert.Contains(t, tokenExchangeProfileList.TokenExchangeProfiles, tokenExchangeProfile)
}

func TestTokenExchangeProfileManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	tokenExchangeProfile := givenAnTokenExchangeProfile(t)

	tokenExchangeProfileRetrieved, err := api.TokenExchangeProfile.Read(context.Background(), tokenExchangeProfile.GetID())
	assert.NoError(t, err)
	assert.Equal(t, tokenExchangeProfileRetrieved, tokenExchangeProfile)
}

func TestTokenExchangeProfileManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	tokenExchangeProfile := givenAnTokenExchangeProfile(t)

	tokenExchangeProfileUpdated := &TokenExchangeProfile{
		Name:             auth0.Stringf("%d-Test-Token-Exchange-Profile-Updated", time.Now().UTC().Unix()),
		SubjectTokenType: auth0.String("https://acme.com/ping-token"),
		ActionID:         auth0.String(tokenExchangeProfile.GetActionID()),
		Type:             auth0.String("custom_authentication"),
	}
	err := api.TokenExchangeProfile.Update(context.Background(), tokenExchangeProfile.GetID(), tokenExchangeProfileUpdated)
	assert.NoError(t, err)

	tokenExchangeProfileRetrieved, err := api.TokenExchangeProfile.Read(context.Background(), tokenExchangeProfile.GetID())
	assert.NoError(t, err)
	assert.Equal(t, tokenExchangeProfileRetrieved, tokenExchangeProfileUpdated)
}

func TestTokenExchangeProfileManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	tokenExchangeProfile := givenAnTokenExchangeProfile(t)

	err := api.TokenExchangeProfile.Delete(context.Background(), tokenExchangeProfile.GetID())
	assert.NoError(t, err)

	_, err = api.TokenExchangeProfile.Read(context.Background(), tokenExchangeProfile.GetID())
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenAnTokenExchangeProfile(t *testing.T) *TokenExchangeProfile {
	action := givenAnCustomTokenExchangeAction(t)
	tokenExchangeProfile := &TokenExchangeProfile{
		Name:             auth0.Stringf("%d-Test-Token-Exchange-Profile", time.Now().UTC().Unix()),
		SubjectTokenType: auth0.String("https://acme.com/cis-token"),
		ActionID:         auth0.String(action.GetID()),
		Type:             auth0.String("custom_authentication"),
	}
	err := api.TokenExchangeProfile.Create(context.Background(), tokenExchangeProfile)
	assert.NoError(t, err)

	t.Cleanup(func() {
		cleanupTokenExchangeProfile(t, tokenExchangeProfile.GetID())
		cleanupAction(t, action.GetID())
	})

	return tokenExchangeProfile
}

func givenAnCustomTokenExchangeAction(t *testing.T) *Action {
	action := &Action{
		Name:              auth0.Stringf("%d-Test-Action", time.Now().UTC().Unix()),
		SupportedTriggers: []ActionTrigger{{ID: auth0.String("custom-token-exchange"), Version: auth0.String("v1")}},
		Code:              auth0.String("exports.onExecutePostLogin = async (event, api) => {}"),
		Runtime:           auth0.String("node18"),
	}
	err := api.Action.Create(context.Background(), action)
	assert.NoError(t, err)
	ensureActionBuilt(t, action.GetID())

	_, err = api.Action.Deploy(context.Background(), action.GetID())
	assert.NoError(t, err)
	assert.NotEmpty(t, action.GetID())

	return action
}

func cleanupTokenExchangeProfile(t *testing.T, id string) {
	t.Helper()

	err := api.TokenExchangeProfile.Delete(context.Background(), id)
	assert.NoError(t, err)
}
