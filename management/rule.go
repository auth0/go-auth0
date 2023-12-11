package management

import "context"

// Rule is used as part of the authentication pipeline.
type Rule struct {
	// The rule's identifier.
	ID *string `json:"id,omitempty"`

	// The name of the rule. Can only contain alphanumeric characters, spaces
	// and '-'. Can neither start nor end with '-' or spaces.
	Name *string `json:"name,omitempty"`

	// A script that contains the rule's code.
	Script *string `json:"script,omitempty"`

	// The rule's order in relation to other rules. A rule with a lower order
	// than another rule executes first. If no order is provided it will
	// automatically be one greater than the current maximum.
	Order *int `json:"order,omitempty"`

	// Enabled should be set to true if the rule is enabled, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
}

// RuleList holds a list of Rules.
type RuleList struct {
	List
	Rules []*Rule `json:"rules"`
}

// RuleManager manages Auth0 Rule resources.
type RuleManager manager

// Create a new rule.
//
// Note: Changing a rule's stage of execution from the default `login_success`
// can change the rule's function signature to have user omitted.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/post_rules
func (m *RuleManager) Create(ctx context.Context, r *Rule, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("rules"), r, opts...)
}

// Retrieve rule details. Accepts a list of fields to include or exclude in the result.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/get_rules_by_id
func (m *RuleManager) Read(ctx context.Context, id string, opts ...RequestOption) (r *Rule, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("rules", id), &r, opts...)
	return
}

// Update an existing rule.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/patch_rules_by_id
func (m *RuleManager) Update(ctx context.Context, id string, r *Rule, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("rules", id), r, opts...)
}

// Delete a rule.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/delete_rules_by_id
func (m *RuleManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("rules", id), nil, opts...)
}

// List rules.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/get_rules
func (m *RuleManager) List(ctx context.Context, opts ...RequestOption) (r *RuleList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("rules"), &r, applyListDefaults(opts))
	return
}
