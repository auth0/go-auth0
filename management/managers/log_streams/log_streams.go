package logStreams

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

// Delete Delete log stream
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/delete_log_streams_by_id
func (m *Manager) Delete(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("log-streams", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAll Get log streams
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/get_log_streams
func (m *Manager) GetAll(ctx context.Context, opts ...management.RequestOption) ([]*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue []*models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("log-streams"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Get Get log stream by ID
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/get_log_streams_by_id
func (m *Manager) Get(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue *models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("log-streams", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Update Update a log stream
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/patch_log_streams_by_id
func (m *Manager) Update(ctx context.Context, id string, patchLogStreamsByIdRequest *models.PatchLogStreamsByIdRequest, opts ...management.RequestOption) (*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue *models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("log-streams", string(id)), patchLogStreamsByIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Create Create a log stream
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/post_log_streams
func (m *Manager) Create(ctx context.Context, postLogStreamsRequest *models.PostLogStreamsRequest, opts ...management.RequestOption) (*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue *models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "POST", m.management.URI("log-streams"), postLogStreamsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
