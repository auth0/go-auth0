package rulesConfigs

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

// DeleteRulesConfigsByKey Delete rules config for a given key
//
// https://auth0.com/docs/api/management/v2/#!/RulesConfigs/delete_rules_configs_by_key
func (m *Manager) DeleteRulesConfigsByKey(ctx context.Context, key string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("rules-configs", string(key)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetRulesConfigs Retrieve config variable keys for rules (get_rules-configs)
//
// https://auth0.com/docs/api/management/v2/#!/RulesConfigs/get_rules_configs
func (m *Manager) GetRulesConfigs(ctx context.Context, opts ...management.RequestOption) ([]*models.GetRulesConfigs200ResponseInner, error) {
	var localVarReturnValue []*models.GetRulesConfigs200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("rules-configs"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutRulesConfigsByKey Set rules config for a given key
//
// https://auth0.com/docs/api/management/v2/#!/RulesConfigs/put_rules_configs_by_key
func (m *Manager) PutRulesConfigsByKey(ctx context.Context, key string, putRulesConfigsByKeyRequest *models.PutRulesConfigsByKeyRequest, opts ...management.RequestOption) (*models.PutRulesConfigsByKey200Response, error) {
	var localVarReturnValue *models.PutRulesConfigsByKey200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("rules-configs", string(key)), putRulesConfigsByKeyRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
