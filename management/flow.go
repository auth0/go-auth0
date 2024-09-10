package management

import (
	"context"
	"encoding/json"
	"time"
)

// Flow represents an Auth0 flow for flow resource.
//
// See: https://auth0.com/docs/customize/flows/intro-to-flows
type Flow struct {
	// Flow identifier
	ID *string `json:"id,omitempty"`

	// Flow name
	Name *string `json:"name,omitempty"`

	// Flow description
	Description *string `json:"description,omitempty"`

	// Sync or Async flow type
	Synchronous *bool `json:"synchronous,omitempty"`

	// Flow actions
	Actions []interface{} `json:"actions,omitempty"`

	// Flow triggers to execute flows
	Triggers *map[string]interface{} `json:"triggers,omitempty"`

	// Security configuration for flow executions
	Security *map[string]interface{} `json:"security,omitempty"`

	// Number of flows linked to this flow
	FlowCount *int `json:"flow_count,omitempty"`

	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	ExecutedAt *time.Time `json:"executed_at,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface.
func (f *Flow) MarshalJSON() ([]byte, error) {
	type FlowSubset struct {
		Name        *string                 `json:"name,omitempty"`
		Description *string                 `json:"description,omitempty"`
		Synchronous *bool                   `json:"synchronous,omitempty"`
		Actions     []interface{}           `json:"actions,omitempty"`
		Triggers    *map[string]interface{} `json:"triggers,omitempty"`
		Security    *map[string]interface{} `json:"security,omitempty"`
	}

	return json.Marshal(&FlowSubset{
		Name:        f.Name,
		Description: f.Description,
		Synchronous: f.Synchronous,
		Actions:     f.Actions,
		Triggers:    f.Triggers,
		Security:    f.Security,
	})
}

// FlowList holds a list of Flow.
type FlowList struct {
	List
	Flows []*Flow `json:"flows"`
}

// FlowManager manages Auth0 Flow resources.
type FlowManager manager

// Create a new flow.
func (m *FlowManager) Create(ctx context.Context, r *Flow, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("flows"), r, opts...)
}

// Retrieve flow details.
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
