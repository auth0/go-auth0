package management

import (
	"context"
	"net/http"
	"time"
)

const (
	// ActionTriggerPostLogin constant.
	ActionTriggerPostLogin string = "post-login"
	// ActionTriggerClientCredentials constant.
	ActionTriggerClientCredentials string = "client-credentials"
)

// ActionTrigger is part of a Flow.
//
// See: https://auth0.com/docs/customize/actions/flows-and-triggers
type ActionTrigger struct {
	ID      *string `json:"id"`
	Version *string `json:"version"`
	Status  *string `json:"status,omitempty"`
}

// ActionTriggerList is a list of ActionTriggers.
type ActionTriggerList struct {
	Triggers []*ActionTrigger `json:"triggers"`
}

// ActionDependency is used to allow the use of packages from the npm registry.
//
// See: https://auth0.com/docs/customize/actions/flows-and-triggers
type ActionDependency struct {
	Name        *string `json:"name"`
	Version     *string `json:"version,omitempty"`
	RegistryURL *string `json:"registry_url,omitempty"`
}

// ActionSecret is used to hold Secret values within an Action.
//
// See: https://auth0.com/docs/customize/actions/write-your-first-action#add-a-secret
type ActionSecret struct {
	Name      *string    `json:"name"`
	Value     *string    `json:"value,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ActionVersionError is used to keep track of
// the errors of a specific ActionVersion.
type ActionVersionError struct {
	ID      *string `json:"id"`
	Message *string `json:"msg"`
	URL     *string `json:"url"`
}

const (
	// ActionStatusPending constant.
	ActionStatusPending string = "pending"
	// ActionStatusBuilding constant.
	ActionStatusBuilding string = "building"
	// ActionStatusPackaged constant.
	ActionStatusPackaged string = "packaged"
	// ActionStatusBuilt constant.
	ActionStatusBuilt string = "built"
	// ActionStatusRetrying constant.
	ActionStatusRetrying string = "retrying"
	// ActionStatusFailed constant.
	ActionStatusFailed string = "failed"
)

// Action represents an Auth0 Action.
//
// See: https://auth0.com/docs/customize/actions/actions-overview
type Action struct {
	// ID of the action
	ID *string `json:"id,omitempty"`
	// The name of an action
	Name *string `json:"name"`
	// List of triggers that this action supports. At this time, an action can
	// only target a single trigger at a time.
	SupportedTriggers []ActionTrigger `json:"supported_triggers"`
	// The source code of the action.
	Code *string `json:"code,omitempty"`
	// List of third party npm modules, and their versions, that this action
	// depends on.
	Dependencies *[]ActionDependency `json:"dependencies,omitempty"`
	// The Node runtime. For example `node16`, defaults to `node12`
	Runtime *string `json:"runtime,omitempty"`
	// List of secrets that are included in an action or a version of an action.
	Secrets *[]ActionSecret `json:"secrets,omitempty"`
	// Version of the action that is currently deployed.
	DeployedVersion *ActionVersion `json:"deployed_version,omitempty"`
	// The build status of this action.
	Status *string `json:"status,omitempty"`
	// True if all of an Action's contents have been deployed.
	AllChangesDeployed bool `json:"all_changes_deployed,omitempty"`
	// The time when this action was built successfully.
	BuiltAt *time.Time `json:"built_at,omitempty"`
	// The time when this action was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time when this action was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ActionList is a list of Actions.
type ActionList struct {
	List
	Actions []*Action `json:"actions"`
}

// ActionVersion is used to manage Actions version history.
//
// See: https://auth0.com/docs/customize/actions/manage-versions
type ActionVersion struct {
	ID           *string             `json:"id,omitempty"`
	Code         *string             `json:"code"`
	Dependencies []*ActionDependency `json:"dependencies,omitempty"`
	Deployed     bool                `json:"deployed"`
	Status       *string             `json:"status,omitempty"`
	Number       int                 `json:"number,omitempty"`

	Errors []*ActionVersionError `json:"errors,omitempty"`
	Action *Action               `json:"action,omitempty"`

	BuiltAt   *time.Time `json:"built_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ActionVersionList is a list of ActionVersions.
type ActionVersionList struct {
	List
	Versions []*ActionVersion `json:"versions"`
}

const (
	// ActionBindingReferenceByName constant.
	ActionBindingReferenceByName string = "action_name"
	// ActionBindingReferenceByID constant.
	ActionBindingReferenceByID string = "action_id"
)

// ActionBindingReference holds the reference
// of an Action attached to an ActionTrigger.
type ActionBindingReference struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

// ActionBinding is used to attach an Action to an ActionTrigger.
type ActionBinding struct {
	ID          *string `json:"id,omitempty"`
	TriggerID   *string `json:"trigger_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`

	Ref     *ActionBindingReference `json:"ref,omitempty"`
	Action  *Action                 `json:"action,omitempty"`
	Secrets []*ActionSecret         `json:"secrets,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ActionBindingList is a list of ActionBindings.
type ActionBindingList struct {
	List
	Bindings []*ActionBinding `json:"bindings"`
}

type actionBindingsPerTrigger struct {
	Bindings []*ActionBinding `json:"bindings"`
}

// ActionTestPayload is used for testing Actions prior to being deployed.
//
// See: https://auth0.com/docs/customize/actions/test-actions
type ActionTestPayload map[string]interface{}

type actionTestRequest struct {
	Payload *ActionTestPayload `json:"payload"`
}

// ActionExecutionResult holds the results of an ActionExecution.
type ActionExecutionResult struct {
	ActionName *string                `json:"action_name,omitempty"`
	Error      map[string]interface{} `json:"error,omitempty"`

	StartedAt *time.Time `json:"started_at,omitempty"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}

// ActionExecution is used to retrieve information
// about a specific execution of an ActionTrigger.
type ActionExecution struct {
	ID        *string                  `json:"id"`
	TriggerID *string                  `json:"trigger_id"`
	Status    *string                  `json:"status"`
	Results   []*ActionExecutionResult `json:"results"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// ActionManager manages Auth0 Action resources.
type ActionManager manager

func applyActionsListDefaults(options []RequestOption) RequestOption {
	return newRequestOption(func(r *http.Request) {
		PerPage(50).apply(r)
		for _, option := range options {
			option.apply(r)
		}
	})
}

// Triggers lists the available triggers.
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_triggers
func (m *ActionManager) Triggers(ctx context.Context, opts ...RequestOption) (l *ActionTriggerList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "triggers"), &l, opts...)
	return
}

// Create a new action.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/post_action
func (m *ActionManager) Create(ctx context.Context, a *Action, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("actions", "actions"), a, opts...)
}

// Retrieve action details.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/get_action
func (m *ActionManager) Read(ctx context.Context, id string, opts ...RequestOption) (a *Action, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "actions", id), &a, opts...)
	return
}

// Update an existing action.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/patch_action
func (m *ActionManager) Update(ctx context.Context, id string, a *Action, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("actions", "actions", id), &a, opts...)
}

// Delete an action
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/delete_action
func (m *ActionManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("actions", "actions", id), nil, opts...)
}

// List actions.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/get_actions
func (m *ActionManager) List(ctx context.Context, opts ...RequestOption) (l *ActionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "actions"), &l, applyActionsListDefaults(opts))
	return
}

// Version retrieves the version of an action.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_action_version
func (m *ActionManager) Version(ctx context.Context, id string, versionID string, opts ...RequestOption) (v *ActionVersion, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "actions", id, "versions", versionID), &v, opts...)
	return
}

// Versions lists versions of an action.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_action_versions
func (m *ActionManager) Versions(ctx context.Context, id string, opts ...RequestOption) (c *ActionVersionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "actions", id, "versions"), &c, applyActionsListDefaults(opts))
	return
}

// UpdateBindings of a trigger.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/patch_bindings
func (m *ActionManager) UpdateBindings(ctx context.Context, triggerID string, b []*ActionBinding, opts ...RequestOption) error {
	bl := &actionBindingsPerTrigger{
		Bindings: b,
	}
	return m.management.Request(ctx, "PATCH", m.management.URI("actions", "triggers", triggerID, "bindings"), &bl, opts...)
}

// Bindings lists the bindings of a trigger.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_bindings
func (m *ActionManager) Bindings(ctx context.Context, triggerID string, opts ...RequestOption) (bl *ActionBindingList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "triggers", triggerID, "bindings"), &bl, applyActionsListDefaults(opts))
	return
}

// Deploy an action
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_action
func (m *ActionManager) Deploy(ctx context.Context, id string, opts ...RequestOption) (v *ActionVersion, err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("actions", "actions", id, "deploy"), &v, opts...)
	return
}

// DeployVersion of an action
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_draft_version
func (m *ActionManager) DeployVersion(ctx context.Context, id string, versionID string, opts ...RequestOption) (v *ActionVersion, err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("actions", "actions", id, "versions", versionID, "deploy"), &v, opts...)
	return
}

// Test an action.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/post_test_action
func (m *ActionManager) Test(ctx context.Context, id string, payload *ActionTestPayload, opts ...RequestOption) (err error) {
	r := &actionTestRequest{
		Payload: payload,
	}
	err = m.management.Request(ctx, "POST", m.management.URI("actions", "actions", id, "test"), &r, opts...)
	return
}

// Execution retrieves the details of an action execution.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_execution
func (m *ActionManager) Execution(ctx context.Context, executionID string, opts ...RequestOption) (v *ActionExecution, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("actions", "executions", executionID), &v, opts...)
	return
}
