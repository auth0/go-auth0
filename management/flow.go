package management

import (
	"context"
	"encoding/json"
	"time"
)

// Flow represents an Auth0 flow for flow resource.
//
// See: https://auth0.com/docs/customize/forms/intro-to-flows
type Flow struct {
	// Flow identifier
	ID *string `json:"id,omitempty"`

	// Flow name
	Name *string `json:"name,omitempty"`
	// Flow actions
	Actions []interface{} `json:"actions,omitempty"`

	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	ExecutedAt *time.Time `json:"executed_at,omitempty"`
}

// FlowList holds a list of Flow.
type FlowList struct {
	List
	Flows []*Flow `json:"flows"`
}

// MarshalJSON implements the json.Marshaller interface.
func (f *Flow) MarshalJSON() ([]byte, error) {
	type FlowSubset struct {
		Name    *string       `json:"name,omitempty"`
		Actions []interface{} `json:"actions,omitempty"`
	}

	return json.Marshal(&FlowSubset{
		Name:    f.Name,
		Actions: f.Actions,
	})
}

// FlowVaultConnection represents an Auth0 flow vault connection resource.
//
// See: https://auth0.com/docs/customize/forms/vault
type FlowVaultConnection struct {
	// Flow vault connection identifier
	ID *string `json:"id,omitempty"`

	// Flow vault connection app identifier
	AppID *string `json:"app_id,omitempty"`

	// Flow vault connection environment
	Environment *string `json:"environment,omitempty"`

	// Flow vault connection name
	Name *string `json:"name,omitempty"`

	// Flow vault connection configuration
	Setup *map[string]interface{} `json:"setup,omitempty"`

	// Flow vault connection custom account name
	AccountName *string `json:"account_name,omitempty"`

	// When Flow vault connection is configured
	Ready *bool `json:"ready,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty"`

	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	RefreshedAt *time.Time `json:"refreshed_at,omitempty"`
}

// FlowVaultConnectionList holds a list of flow vault connections.
type FlowVaultConnectionList struct {
	List
	Connections []*FlowVaultConnection `json:"connections"`
}

// MarshalJSON implements the json.Marshaller interface.
func (f *FlowVaultConnection) MarshalJSON() ([]byte, error) {
	type FlowVaultConnectionSubset struct {
		AppID *string                 `json:"app_id,omitempty"`
		Name  *string                 `json:"name,omitempty"`
		Setup *map[string]interface{} `json:"setup,omitempty"`
	}

	return json.Marshal(&FlowVaultConnectionSubset{
		AppID: f.AppID,
		Name:  f.Name,
		Setup: f.Setup,
	})
}

// FlowManager manages Auth0 Flow resources.
type FlowManager struct {
	management *Management

	// FlowVaultConnection manages flow vault connection resources.
	Vault *flowVaultConnectionManager
}

// Create a new flow.
func (m *FlowManager) Create(ctx context.Context, r *Flow, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("flows"), r, opts...)
}

// Read retrieves flow details.
func (m *FlowManager) Read(ctx context.Context, id string, opts ...RequestOption) (r *Flow, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("flows", id), &r, opts...)
	return
}

// Update an existing flow.
func (m *FlowManager) Update(ctx context.Context, id string, r *Flow, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("flows", id), r, opts...)
}

// Delete a flow.
func (m *FlowManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("flows", id), nil, opts...)
}

// List flow.
func (m *FlowManager) List(ctx context.Context, opts ...RequestOption) (r *FlowList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("flows"), &r, applyListDefaults(opts))
	return
}

// FlowVaultConnectionManager manages flow vault connection resources.
type flowVaultConnectionManager manager

// CreateConnection Create a new flow vault connection.
func (m *flowVaultConnectionManager) CreateConnection(ctx context.Context, r *FlowVaultConnection, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("flows", "vault", "connections"), r, opts...)
}

// GetConnection Retrieve flow vault connection details.
func (m *flowVaultConnectionManager) GetConnection(ctx context.Context, id string, opts ...RequestOption) (r *FlowVaultConnection, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("flows", "vault", "connections", id), &r, opts...)
	return
}

// UpdateConnection Update an existing flow vault connection.
func (m *flowVaultConnectionManager) UpdateConnection(ctx context.Context, id string, r *FlowVaultConnection, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("flows", "vault", "connections", id), r, opts...)
}

// DeleteConnection Delete a flow vault connection.
func (m *flowVaultConnectionManager) DeleteConnection(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("flows", "vault", "connections", id), nil, opts...)
}

// GetConnectionList List flow vault connections.
func (m *flowVaultConnectionManager) GetConnectionList(ctx context.Context, opts ...RequestOption) (r *FlowVaultConnectionList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("flows", "vault", "connections"), &r, applyListDefaults(opts))
	return
}
