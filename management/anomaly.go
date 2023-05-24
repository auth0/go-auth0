package management

import (
	"context"
	"net/http"
)

// AnomalyManager manages Auth0 Anomaly resources.
type AnomalyManager manager

// CheckIP checks if a given IP address is blocked via the multiple
// user accounts trigger due to multiple failed logins.
//
// See: https://auth0.com/docs/api/management/v2#!/Anomaly/get_ips_by_id
func (m *AnomalyManager) CheckIP(ctx context.Context, ip string, opts ...RequestOption) (isBlocked bool, err error) {
	request, err := m.management.NewRequest(ctx, "GET", m.management.URI("anomaly", "blocks", "ips", ip), nil, opts...)
	if err != nil {
		return false, err
	}

	response, err := m.management.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	// 200: IP address specified is currently blocked.
	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	// 404: IP address specified is not currently blocked.
	if response.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, newError(response)
}

// UnblockIP unblocks an IP address currently blocked by the multiple
// user accounts trigger due to multiple failed logins.
//
// See: https://auth0.com/docs/api/management/v2#!/Anomaly/delete_ips_by_id
func (m *AnomalyManager) UnblockIP(ctx context.Context, ip string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("anomaly", "blocks", "ips", ip), nil, opts...)
}
