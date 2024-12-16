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

// DeleteLogStreamsById Delete log stream
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/delete_log_streams_by_id
func (m *Manager) DeleteLogStreamsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("log-streams", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetLogStreams Get log streams
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/get_log_streams
func (m *Manager) GetLogStreams(ctx context.Context, opts ...management.RequestOption) ([]*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue []*models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("log-streams"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetLogStreamsById Get log stream by ID
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/get_log_streams_by_id
func (m *Manager) GetLogStreamsById(ctx context.Context, id string, opts ...management.RequestOption) (*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue *models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("log-streams", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchLogStreamsById Update a log stream
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/patch_log_streams_by_id
func (m *Manager) PatchLogStreamsById(ctx context.Context, id string, patchLogStreamsByIdRequest *models.PatchLogStreamsByIdRequest, opts ...management.RequestOption) (*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue *models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "PATCH", m.management.URI("log-streams", string(id)), patchLogStreamsByIdRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostLogStreams Create a log stream
//
// https://auth0.com/docs/api/management/v2/#!/LogStreams/post_log_streams
func (m *Manager) PostLogStreams(ctx context.Context, postLogStreamsRequest *models.PostLogStreamsRequest, opts ...management.RequestOption) (*models.GetLogStreams200ResponseInner, error) {
	var localVarReturnValue *models.GetLogStreams200ResponseInner
	err := m.management.Request(ctx, "POST", m.management.URI("log-streams"), postLogStreamsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
