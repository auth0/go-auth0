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

// DeleteAction Delete an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/delete_action
func (m *Manager) DeleteAction(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("actions", "actions", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAction Get an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_action
func (m *Manager) GetAction(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetActions200ResponseActionsInner, error) {
	var localVarReturnValue *models.GetActions200ResponseActionsInner
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetActionVersion Get a specific version of an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_action_version
func (m *Manager) GetActionVersion(ctx context.Context, actionId string, id string, opts ...management.RequestOption) (*models.GetActionVersions200ResponseVersionsInner, error) {
	var localVarReturnValue *models.GetActionVersions200ResponseVersionsInner
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions", string(actionId), "versions", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetActionVersions Get an action&#39;s versions
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_action_versions
func (m *Manager) GetActionVersions(ctx context.Context, actionId string, opts ...management.RequestOption) (*models.GetActionVersions200Response, error) {
	var localVarReturnValue *models.GetActionVersions200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions", string(actionId), "versions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetActions Get actions
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_actions
func (m *Manager) GetActions(ctx context.Context, opts ...management.RequestOption) (*models.GetActions200Response, error) {
	var localVarReturnValue *models.GetActions200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "actions"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetBindings Get trigger bindings
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_bindings
func (m *Manager) GetBindings(ctx context.Context, triggerId string, opts ...management.RequestOption) (*models.GetBindings200Response, error) {
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

// GetTriggers Get triggers
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_triggers
func (m *Manager) GetTriggers(ctx context.Context, opts ...management.RequestOption) (*models.GetTriggers200Response, error) {
	var localVarReturnValue *models.GetTriggers200Response
	err := m.management.Request(ctx, "GET", m.management.URI("actions", "triggers"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchAction Update an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/patch_action
func (m *Manager) PatchAction(ctx context.Context, id string, patchActionRequest *models.PatchActionRequest, opts ...management.RequestOption) (*models.GetActions200ResponseActionsInner, error) {
	var localVarReturnValue *models.GetActions200ResponseActionsInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("actions", "actions", string(id)), patchActionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchBindings Update trigger bindings
//
// https://auth0.com/docs/api/management/v2/#!/Actions/patch_bindings
func (m *Manager) PatchBindings(ctx context.Context, triggerId string, patchBindingsRequest *models.PatchBindingsRequest, opts ...management.RequestOption) (*models.PatchBindings200Response, error) {
	var localVarReturnValue *models.PatchBindings200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("actions", "triggers", string(triggerId), "bindings"), patchBindingsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostAction Create an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_action
func (m *Manager) PostAction(ctx context.Context, postActionRequest *models.PostActionRequest, opts ...management.RequestOption) (*models.GetActions200ResponseActionsInner, error) {
	var localVarReturnValue *models.GetActions200ResponseActionsInner
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions"), postActionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostDeployAction Deploy an action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_action
func (m *Manager) PostDeployAction(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetActionVersions200ResponseVersionsInner, error) {
	var localVarReturnValue *models.GetActionVersions200ResponseVersionsInner
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions", string(id), "deploy"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostDeployDraftVersion Roll back to a previous action version
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_draft_version
func (m *Manager) PostDeployDraftVersion(ctx context.Context, id string, actionId string, postDeployDraftVersionRequest *models.PostDeployDraftVersionRequest, opts ...management.RequestOption) (*models.GetActionVersions200ResponseVersionsInner, error) {
	var localVarReturnValue *models.GetActionVersions200ResponseVersionsInner
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions", string(actionId), "versions", string(id), "deploy"), postDeployDraftVersionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostTestAction Test an Action
//
// https://auth0.com/docs/api/management/v2/#!/Actions/post_test_action
func (m *Manager) PostTestAction(ctx context.Context, id string, postTestActionRequest *models.PostTestActionRequest, opts ...management.RequestOption) (*models.PostTestAction200Response, error) {
	var localVarReturnValue *models.PostTestAction200Response
	err := m.management.Request(ctx, "POST", m.management.URI("actions", "actions", string(id), "test"), postTestActionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
