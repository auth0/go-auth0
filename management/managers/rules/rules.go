package rules

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

// DeleteRulesById Delete a rule
//
// https://auth0.com/docs/api/management/v2/#!/Rules/delete_rules_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("rules", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetRules Get rules
//
// https://auth0.com/docs/api/management/v2/#!/Rules/get_rules
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) (*models.GetRules200Response, error) {
	var localVarReturnValue *models.GetRules200Response
	err := m.management.Request(ctx, "GET", m.management.URI("rules"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetRulesById Get a rule
//
// https://auth0.com/docs/api/management/v2/#!/Rules/get_rules_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.Rule, error) {
	var localVarReturnValue *models.Rule
	err := m.management.Request(ctx, "GET", m.management.URI("rules", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchRulesById Update a rule
//
// https://auth0.com/docs/api/management/v2/#!/Rules/patch_rules_by_id
func (m *Manager) Update(ctx context.Context, id string, ruleUpdate *models.RuleUpdate, opts ...management.RequestOption) (*models.Rule, error) {
	var localVarReturnValue *models.Rule
	err := m.management.Request(ctx, "PATCH", m.management.URI("rules", string(id)), ruleUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostRules Create a rule
//
// https://auth0.com/docs/api/management/v2/#!/Rules/post_rules
func (m *Manager) Create(ctx context.Context, ruleCreate *models.RuleCreate, opts ...management.RequestOption) (*models.Rule, error) {
	var localVarReturnValue *models.Rule
	err := m.management.Request(ctx, "POST", m.management.URI("rules"), ruleCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
