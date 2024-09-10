package management

import (
	"encoding/json"
	"time"
)

// Flow represents an Auth0 flow for form resource.
//
// See: https://auth0.com/docs/customize/forms/intro-to-flows
type Flow struct {
	// Flow identifier
	ID *string `json:"id"`

	// Flow name
	Name *string `json:"name"`

	// Flow description
	Description *string `json:"description,omitempty"`

	// Sync or Async flow type
	Synchronous *bool `json:"synchronous"`

	// Flow actions
	Actions []interface{} `json:"actions"`

	// Flow triggers to execute forms
	Triggers *map[string]interface{} `json:"triggers"`

	// Security configuration for flow executions
	Security *map[string]interface{} `json:"security"`

	// Number of forms linked to this flow
	FormCount *int `json:"form_count"`

	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	ExecutedAt *time.Time `json:"executed_at"`
}

// MarshalJSON implements the json.Marshaller interface.
func (f *Flow) MarshalJSON() ([]byte, error) {
	type FlowSubset struct {
		Name        *string                 `json:"name"`
		Description *string                 `json:"description,omitempty"`
		Synchronous *bool                   `json:"synchronous"`
		Actions     []interface{}           `json:"actions"`
		Triggers    *map[string]interface{} `json:"triggers"`
		Security    *map[string]interface{} `json:"security"`
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
