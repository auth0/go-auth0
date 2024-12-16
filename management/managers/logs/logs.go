package logs

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

// GetLogs Search log events
//
// https://auth0.com/docs/api/management/v2/#!/Logs/get_logs
func (m *Manager) GetLogs(ctx context.Context, opts ...management.RequestOption) (*models.GetLogs200Response, error) {
	var localVarReturnValue *models.GetLogs200Response
	err := m.management.Request(ctx, "GET", m.management.URI("logs"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetLogsById Get a log event by id
//
// https://auth0.com/docs/api/management/v2/#!/Logs/get_logs_by_id
func (m *Manager) GetLogsById(ctx context.Context, id string, opts ...management.RequestOption) (*models.Log, error) {
	var localVarReturnValue *models.Log
	err := m.management.Request(ctx, "GET", m.management.URI("logs", management.SafeString(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
