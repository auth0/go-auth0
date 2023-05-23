package management

import (
	"context"
	"time"
)

// StatManager manages Auth0 DailyStat resources.
type StatManager manager

// ActiveUsers retrieves the number of active users that logged in during the
// last 30 days.
//
// See: https://auth0.com/docs/api/management/v2#!/Stats/get_active_users
func (m *StatManager) ActiveUsers(ctx context.Context, opts ...RequestOption) (i int, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("stats", "active-users"), &i, opts...)
	return
}

// DailyStat for an Auth0 Tenant.
type DailyStat struct {
	Date            *time.Time `json:"date"`
	Logins          *int       `json:"logins"`
	Signups         *int       `json:"signups"`
	LeakedPasswords *int       `json:"leaked_passwords"`
	UpdatedAt       *time.Time `json:"updated_at"`
	CreatedAt       *time.Time `json:"created_at"`
}

// Daily retrieves the number of logins, signups and breached-password
// detections (subscription required) that occurred each day within a specified
// date range.
//
// See: https://auth0.com/docs/api/management/v2#!/Stats/get_daily
func (m *StatManager) Daily(ctx context.Context, opts ...RequestOption) (ds []*DailyStat, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("stats", "daily"), &ds, opts...)
	return
}
