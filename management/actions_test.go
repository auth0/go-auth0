package management

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestActionManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := &Action{
		Name: auth0.Stringf("Test Action (%s)", time.Now().Format(time.StampMilli)),
		Code: auth0.String("exports.onExecutePostLogin = async (event, api) =\u003e {}"),
		SupportedTriggers: []ActionTrigger{
			{
				ID:      auth0.String(ActionTriggerPostLogin),
				Version: auth0.String("v2"),
			},
		},
		Dependencies: &[]ActionDependency{
			{
				Name:        auth0.String("lodash"),
				Version:     auth0.String("4.0.0"),
				RegistryURL: auth0.String("https://www.npmjs.com/package/lodash"),
			},
		},
		Secrets: &[]ActionSecret{
			{
				Name:  auth0.String("mySecretName"),
				Value: auth0.String("mySecretValue"),
			},
		},
	}

	err := api.Action.Create(context.Background(), expectedAction)

	assert.NoError(t, err)
	assert.NotEmpty(t, expectedAction.GetID())

	t.Cleanup(func() {
		cleanupAction(t, expectedAction.GetID())
	})
}

func TestActionManager_CreateWithDeploy(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := &Action{
		Name: auth0.Stringf("Test Action (%s)", time.Now().Format(time.StampMilli)),
		Code: auth0.String("exports.onExecutePostLogin = async (event, api) =\u003e {}"),
		SupportedTriggers: []ActionTrigger{
			{
				ID:      auth0.String(ActionTriggerCustomTokenExchange),
				Version: auth0.String("v1"),
			},
		},
		Secrets: &[]ActionSecret{
			{
				Name:  auth0.String("mySecretName"),
				Value: auth0.String("mySecretValue"),
			},
		},
		Deploy:  auth0.Bool(true),
		Runtime: auth0.String("node18"),
	}

	err := api.Action.Create(context.Background(), expectedAction)

	assert.NoError(t, err)
	assert.NotEmpty(t, expectedAction.GetID())

	ensureActionBuilt(t, expectedAction.GetID())

	actualAction, err := api.Action.Read(context.Background(), expectedAction.GetID())
	assert.NoError(t, err)
	assert.Equal(t, ActionStatusBuilt, actualAction.GetStatus())
	assert.NotEmpty(t, actualAction.GetDeployedVersion())
	t.Cleanup(func() {
		cleanupAction(t, expectedAction.GetID())
	})
}

func TestActionManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := givenAnAction(t)
	actualAction, err := api.Action.Read(context.Background(), expectedAction.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedAction.GetID(), actualAction.GetID())
}

func TestActionManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := givenAnAction(t)

	actionID := expectedAction.GetID()

	expectedAction.ID = nil        // Read-Only: Additional properties not allowed.
	expectedAction.UpdatedAt = nil // Read-Only: Additional properties not allowed.
	expectedAction.CreatedAt = nil // Read-Only: Additional properties not allowed.
	expectedAction.Status = nil    // Read-Only: Additional properties not allowed.

	expectedCode := "exports.onExecutePostLogin = async (event, api) => { api.user.setUserMetadata('myParam', 'foo'); };"
	expectedAction.Code = &expectedCode

	err := api.Action.Update(context.Background(), actionID, expectedAction)

	assert.NoError(t, err)
	assert.Equal(t, expectedCode, *expectedAction.Code)
}

func TestActionManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := givenAnAction(t)

	err := api.Action.Delete(context.Background(), expectedAction.GetID())
	assert.NoError(t, err)

	actualAction, err := api.Action.Read(context.Background(), expectedAction.GetID())

	assert.Empty(t, actualAction)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestActionManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := givenAnAction(t)

	actionList, err := api.Action.List(context.Background(), Parameter("actionName", expectedAction.GetName()))

	assert.NoError(t, err)
	assert.Equal(t, expectedAction.GetID(), actionList.Actions[0].GetID())
}

func TestActionManager_Triggers(t *testing.T) {
	configureHTTPTestRecordings(t)

	actionTriggerList, err := api.Action.Triggers(context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, actionTriggerList)
}

func TestActionManager_Deploy(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedAction := givenAnAction(t)

	ensureActionBuilt(t, expectedAction.GetID())

	actualActionVersion, err := api.Action.Deploy(context.Background(), expectedAction.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedAction.GetID(), actualActionVersion.Action.GetID())
}

func TestActionManager_DeployVersion(t *testing.T) {
	configureHTTPTestRecordings(t)

	action := givenAnAction(t)
	ensureActionBuilt(t, action.GetID())

	version, err := api.Action.Deploy(context.Background(), action.GetID())
	require.NoError(t, err)

	_, err = api.Action.DeployVersion(context.Background(), action.GetID(), version.GetID())

	assert.NoError(t, err)
}

func TestActionManager_Version(t *testing.T) {
	configureHTTPTestRecordings(t)

	action := givenAnAction(t)
	ensureActionBuilt(t, action.GetID())

	deployedVersion, err := api.Action.Deploy(context.Background(), action.GetID())
	require.NoError(t, err)

	actualVersion, err := api.Action.Version(context.Background(), action.GetID(), deployedVersion.GetID())

	assert.NoError(t, err)
	assert.Equal(t, deployedVersion.GetID(), actualVersion.GetID())
}

func TestActionManager_Versions(t *testing.T) {
	configureHTTPTestRecordings(t)

	action := givenAnAction(t)
	ensureActionBuilt(t, action.GetID())

	deployedVersion, err := api.Action.Deploy(context.Background(), action.GetID())
	require.NoError(t, err)

	actualVersions, err := api.Action.Versions(context.Background(), action.GetID())

	assert.NoError(t, err)
	assert.Equal(t, deployedVersion.GetID(), actualVersions.Versions[0].GetID())
}

func TestActionManager_Bindings(t *testing.T) {
	configureHTTPTestRecordings(t)

	action := givenAnAction(t)
	ensureActionBuilt(t, action.GetID())

	_, err := api.Action.Deploy(context.Background(), action.GetID())
	require.NoError(t, err)

	emptyBinding := make([]*ActionBinding, 0)
	err = api.Action.UpdateBindings(context.Background(), ActionTriggerPostLogin, emptyBinding)
	assert.NoError(t, err)

	binding := []*ActionBinding{
		{
			Ref: &ActionBindingReference{
				Type:  auth0.String(ActionBindingReferenceByName),
				Value: action.Name,
			},
			DisplayName: auth0.String("My test action Binding"),
		},
	}

	err = api.Action.UpdateBindings(context.Background(), ActionTriggerPostLogin, binding)
	assert.NoError(t, err)

	bindingList, err := api.Action.Bindings(context.Background(), ActionTriggerPostLogin)

	assert.NoError(t, err)
	assert.Len(t, bindingList.Bindings, 1)

	t.Cleanup(func() {
		err = api.Action.UpdateBindings(context.Background(), ActionTriggerPostLogin, emptyBinding)
		assert.NoError(t, err)
	})
}

func TestActionManager_Test(t *testing.T) {
	configureHTTPTestRecordings(t)

	action := givenAnAction(t)
	ensureActionBuilt(t, action.GetID())

	test := &ActionTestPayload{
		"event": ActionTestPayload{
			"user": ActionTestPayload{
				"email":         "j+smith@example.com",
				"emailVerified": true,
				"id":            "auth0|5f7c8ec7c33c6c004bbafe82",
			},
		},
	}
	err := api.Action.Test(context.Background(), action.GetID(), test)
	assert.NoError(t, err)
}

func TestActionManager_Execution(t *testing.T) {
	configureHTTPTestRecordings(t)

	_, err := api.Action.Execution(context.Background(), "M9IqRp9wQLaYNrSwz6YPTTIwMjEwNDA0")
	// Expect a 404 as we can't get execution ID via API
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func cleanupAction(t *testing.T, actionID string) {
	t.Helper()

	err := api.Action.Delete(context.Background(), actionID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}

func givenAnAction(t *testing.T) *Action {
	t.Helper()

	action := &Action{
		Name: auth0.Stringf("Test Action (%s)", time.Now().Format(time.StampMilli)),
		Code: auth0.String("exports.onExecutePostLogin = async (event, api) =\u003e {}"),
		SupportedTriggers: []ActionTrigger{
			{
				ID:      auth0.String(ActionTriggerPostLogin),
				Version: auth0.String("v2"),
			},
		},
		Dependencies: &[]ActionDependency{
			{
				Name:        auth0.String("lodash"),
				Version:     auth0.String("4.0.0"),
				RegistryURL: auth0.String("https://www.npmjs.com/package/lodash"),
			},
		},
		Secrets: &[]ActionSecret{
			{
				Name:  auth0.String("mySecretName"),
				Value: auth0.String("mySecretValue"),
			},
		},
	}

	err := api.Action.Create(context.Background(), action)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupAction(t, action.GetID())
	})

	return action
}

func ensureActionBuilt(t *testing.T, actionID string) {
	t.Helper()

	var actionBuilt bool

	for i := 0; i < 60; i++ {
		action, err := api.Action.Read(context.Background(), actionID)
		assert.NoError(t, err)

		if action.GetStatus() == ActionStatusBuilt {
			actionBuilt = true
			break
		}

		time.Sleep(time.Second)
	}

	if actionBuilt {
		return
	}

	t.Fatalf("action with ID: %s, failed to build", actionID)
}
