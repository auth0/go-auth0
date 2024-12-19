package actions

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/auth0/go-auth0/management/models"
)

// Manager defines the base manager interface
type Manager struct {
	management *management.Management
}

// NewManager creates a new manager for  operations
func NewManager(mgmt *management.Management) *Manager {
	return &Manager{
		management: mgmt,
	}
}

// Delete Delete an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/delete_action
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("actions", "actions", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Get Get an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_action
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetActions200ResponseActionsInner, error) {
	var localVarReturnValue *models.GetActions200ResponseActionsInner
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetVersion Get a specific version of an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_action_version
func (m *Manager) GetVersion(ctx context.Context, actionId string, id string, opts ...management.RequestOption) (*models.GetActionVersions200ResponseVersionsInner, error) {
	var localVarReturnValue *models.GetActionVersions200ResponseVersionsInner
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions", string(actionId), "versions", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetVersions Get an action&#39;s versions
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_action_versions
func (m *Manager) GetVersions(ctx context.Context, actionId string, opts ...management.RequestOption) (*models.GetActionVersions200Response, error) {
	var localVarReturnValue *models.GetActionVersions200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions", string(actionId), "versions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetAll Get actions
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_actions
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetActions200Response, error) {
	var localVarReturnValue *models.GetActions200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetTriggerBindings Get trigger bindings
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_bindings
func (m *Manager) GetTriggerBindings(ctx context.Context, triggerId string, opts ...management.RequestOption) (*models.GetBindings200Response, error) {
	var localVarReturnValue *models.GetBindings200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "triggers", string(triggerId), "bindings"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetExecution Get an execution
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_execution
func (m *Manager) GetExecution(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetExecution200Response, error) {
	var localVarReturnValue *models.GetExecution200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "executions", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetAllTriggers Get triggers
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_triggers
func (m *Manager) GetAllTriggers(ctx context.Context, opts ...management.RequestOption) (*models.GetTriggers200Response, error) {
	var localVarReturnValue *models.GetTriggers200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "triggers"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Update Update an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/patch_action
func (m *Manager) Update(ctx context.Context, id string, patchActionRequest *models.PatchActionRequest, opts ...management.RequestOption) (*models.GetActions200ResponseActionsInner, error) {
	var localVarReturnValue *models.GetActions200ResponseActionsInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("actions", "actions", string(id)), patchActionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateTriggerBindings Update trigger bindings
//
// https://auth0.com/docs/api/management/v2/#!/Actions/patch_bindings
func (m *Manager) UpdateTriggerBindings(ctx context.Context, triggerId string, patchBindingsRequest *models.PatchBindingsRequest, opts ...management.RequestOption) (*models.PatchBindings200Response, error) {
	var localVarReturnValue *models.PatchBindings200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("actions", "triggers", string(triggerId), "bindings"), patchBindingsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Create Create an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_action
func (m *Manager) Create(ctx context.Context, postActionRequest *models.PostActionRequest, opts ...management.RequestOption) (*models.GetActions200ResponseActionsInner, error) {
	var localVarReturnValue *models.GetActions200ResponseActionsInner
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions"), postActionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Deploy Deploy an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_action
func (m *Manager) Deploy(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetActionVersions200ResponseVersionsInner, error) {
	var localVarReturnValue *models.GetActionVersions200ResponseVersionsInner
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions", string(id), "deploy"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// DeployVersion Roll back to a previous action version
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_draft_version
func (m *Manager) DeployVersion(ctx context.Context, id string, actionId string, postDeployDraftVersionRequest *models.PostDeployDraftVersionRequest, opts ...management.RequestOption) (*models.GetActionVersions200ResponseVersionsInner, error) {
	var localVarReturnValue *models.GetActionVersions200ResponseVersionsInner
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions", string(actionId), "versions", string(id), "deploy"), postDeployDraftVersionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Test Test an Action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_test_action
func (m *Manager) Test(ctx context.Context, id string, postTestActionRequest *models.PostTestActionRequest, opts ...management.RequestOption) (*models.PostTestAction200Response, error) {
	var localVarReturnValue *models.PostTestAction200Response
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions", string(id), "test"), postTestActionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
