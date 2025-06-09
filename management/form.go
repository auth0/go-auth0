package management

import (
	"context"
	"encoding/json"
	"time"
)

// Form represents an Auth0 form resource.
//
// See: https://auth0.com/docs/customize/forms
type Form struct {
	// ID is the unique identifier for the form.
	ID *string `json:"id,omitempty"`

	// Name is the name of the form.
	Name *string `json:"name,omitempty"`

	// Messages contains custom and error messages for the form.
	Messages *FormMessages `json:"messages,omitempty"`
	// Languages contains the languages of the form.
	Languages *FormLanguages `json:"languages,omitempty"`
	// Translations holds the translations for the form.
	Translations *map[string]interface{} `json:"translations,omitempty"`

	// Start defines the starting point of the form.
	Start *map[string]interface{} `json:"start,omitempty"`
	// Nodes represents the nodes in the form.
	Nodes []interface{} `json:"nodes,omitempty"`
	// Ending defines the ending point of the form.
	Ending *map[string]interface{} `json:"ending,omitempty"`

	// Style contains the style of the form.
	Style *map[string]interface{} `json:"style,omitempty"`

	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	EmbeddedAt  *time.Time `json:"embedded_at,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
}

// FormLanguages represents the languages of the form.
type FormLanguages struct {
	Primary *string `json:"primary,omitempty"`
	Default *string `json:"default,omitempty"`
}

// FormMessages represents custom and error messages of the form.
type FormMessages struct {
	Custom *map[string]interface{} `json:"custom,omitempty"`
	Errors *map[string]interface{} `json:"errors,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface.
func (f *Form) MarshalJSON() ([]byte, error) {
	type FormSubset struct {
		Name         *string                 `json:"name,omitempty"`
		Languages    *FormLanguages          `json:"languages,omitempty"`
		Style        *map[string]interface{} `json:"style,omitempty"`
		Messages     *FormMessages           `json:"messages,omitempty"`
		Translations *map[string]interface{} `json:"translations,omitempty"`
		Start        *map[string]interface{} `json:"start,omitempty"`
		Nodes        []interface{}           `json:"nodes,omitempty"`
		Ending       *map[string]interface{} `json:"ending,omitempty"`
	}

	return json.Marshal(&FormSubset{
		Name:         f.Name,
		Languages:    f.Languages,
		Style:        f.Style,
		Messages:     f.Messages,
		Translations: f.Translations,
		Start:        f.Start,
		Nodes:        f.Nodes,
		Ending:       f.Ending,
	})
}

// FormList holds a list of Forms.
type FormList struct {
	List
	Forms []*Form `json:"forms"`
}

// FormManager manages Auth0 Form resources.
type FormManager manager

// Create a new form.
func (m *FormManager) Create(ctx context.Context, r *Form, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("forms"), r, opts...)
}

// Read retrieves form details.
func (m *FormManager) Read(ctx context.Context, id string, opts ...RequestOption) (r *Form, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("forms", id), &r, opts...)
	return
}

// Update an existing form.
func (m *FormManager) Update(ctx context.Context, id string, r *Form, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("forms", id), r, opts...)
}

// Delete a form.
func (m *FormManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("forms", id), nil, opts...)
}

// List form.
func (m *FormManager) List(ctx context.Context, opts ...RequestOption) (r *FormList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("forms"), &r, applyListDefaults(opts))
	return
}
