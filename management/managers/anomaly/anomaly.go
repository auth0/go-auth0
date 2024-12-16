package anomaly

import (
	"context"

	"github.com/auth0/go-auth0/management"
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

// DeleteIpsById Remove the blocked IP address
//
// https://auth0.com/docs/api/management/v2/#!/Anomaly/delete_ips_by_id
func (m *Manager) DeleteIpsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("anomaly", "blocks", "ips", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetIpsById Check if an IP address is blocked
//
// https://auth0.com/docs/api/management/v2/#!/Anomaly/get_ips_by_id
func (m *Manager) GetIpsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "GET", m.management.URI("anomaly", "blocks", "ips", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
