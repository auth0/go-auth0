package management

import (
	"context"
	"net/http"
)

// AttackProtectionManager manages Auth0 Attack Protection settings.
//
// See: https://auth0.com/docs/secure/attack-protection
type AttackProtectionManager manager

// BreachedPasswordDetection protects applications from
// bad actors logging in with stolen credentials.
//
// See: https://auth0.com/docs/secure/attack-protection/breached-password-detection
type BreachedPasswordDetection struct {
	Enabled                    *bool                           `json:"enabled,omitempty"`
	Shields                    *[]string                       `json:"shields,omitempty"`
	AdminNotificationFrequency *[]string                       `json:"admin_notification_frequency,omitempty"`
	Method                     *string                         `json:"method,omitempty"`
	Stage                      *BreachedPasswordDetectionStage `json:"stage,omitempty"`
}

// BreachedPasswordDetectionStage is used to specify per-stage configuration options.
type BreachedPasswordDetectionStage struct {
	PreUserRegistration *BreachedPasswordDetectionPreUserRegistration `json:"pre-user-registration,omitempty"`
}

// BreachedPasswordDetectionPreUserRegistration is used to specify breached password detection
// configuration (shields) for the sign up flow.
type BreachedPasswordDetectionPreUserRegistration struct {
	// Action to take when a breached password is detected during a signup.
	// Possible values: block, admin_notification.
	Shields *[]string `json:"shields,omitempty"`
}

// GetBreachedPasswordDetection retrieves breached password detection settings.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/get_breached_password_detection
func (m *AttackProtectionManager) GetBreachedPasswordDetection(
	ctx context.Context,
	opts ...RequestOption,
) (*BreachedPasswordDetection, error) {
	var breachedPasswordDetection BreachedPasswordDetection
	err := m.management.Request(
		ctx,
		http.MethodGet,
		m.management.URI("attack-protection", "breached-password-detection"),
		&breachedPasswordDetection,
		opts...,
	)

	return &breachedPasswordDetection, err
}

// UpdateBreachedPasswordDetection updates the breached password detection settings.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/patch_breached_password_detection
func (m *AttackProtectionManager) UpdateBreachedPasswordDetection(
	ctx context.Context,
	breachedPasswordDetection *BreachedPasswordDetection,
	opts ...RequestOption,
) error {
	return m.management.Request(
		ctx,
		http.MethodPatch,
		m.management.URI("attack-protection", "breached-password-detection"),
		breachedPasswordDetection,
		opts...,
	)
}

// BruteForceProtection safeguards against a single
// IP address attacking a single user account.
//
// See: https://auth0.com/docs/secure/attack-protection/brute-force-protection
type BruteForceProtection struct {
	Enabled     *bool     `json:"enabled,omitempty"`
	Shields     *[]string `json:"shields,omitempty"`
	AllowList   *[]string `json:"allowlist,omitempty"`
	Mode        *string   `json:"mode,omitempty"`
	MaxAttempts *int      `json:"max_attempts,omitempty"`
}

// GetBruteForceProtection retrieves the brute force configuration.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/get_brute_force_protection
func (m *AttackProtectionManager) GetBruteForceProtection(
	ctx context.Context,
	opts ...RequestOption,
) (*BruteForceProtection, error) {
	var bruteForceProtection BruteForceProtection
	err := m.management.Request(
		ctx,
		http.MethodGet,
		m.management.URI("attack-protection", "brute-force-protection"),
		&bruteForceProtection,
		opts...,
	)

	return &bruteForceProtection, err
}

// UpdateBruteForceProtection updates the brute force configuration.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/patch_brute_force_protection
func (m *AttackProtectionManager) UpdateBruteForceProtection(
	ctx context.Context,
	bruteForceProtection *BruteForceProtection,
	opts ...RequestOption,
) error {
	return m.management.Request(
		ctx,
		http.MethodPatch,
		m.management.URI("attack-protection", "brute-force-protection"),
		bruteForceProtection,
		opts...,
	)
}

// SuspiciousIPThrottling blocks traffic from any IP address
// that rapidly attempts too many logins or signups.
//
// See: https://auth0.com/docs/secure/attack-protection/suspicious-ip-throttling
type SuspiciousIPThrottling struct {
	Enabled   *bool     `json:"enabled,omitempty"`
	Shields   *[]string `json:"shields,omitempty"`
	AllowList *[]string `json:"allowlist,omitempty"`
	Stage     *Stage    `json:"stage,omitempty"`
}

// Stage is used to customize thresholds for limiting
// suspicious traffic in login and sign up flows.
type Stage struct {
	PreLogin            *PreLogin            `json:"pre-login,omitempty"`
	PreUserRegistration *PreUserRegistration `json:"pre-user-registration,omitempty"`
}

// PreLogin is used to customize thresholds for login flow.
type PreLogin struct {
	MaxAttempts *int `json:"max_attempts,omitempty"`
	Rate        *int `json:"rate,omitempty"`
}

// PreUserRegistration is used to customize thresholds for sign up flow.
type PreUserRegistration struct {
	MaxAttempts *int `json:"max_attempts,omitempty"`
	Rate        *int `json:"rate,omitempty"`
}

// GetSuspiciousIPThrottling retrieves the suspicious IP throttling configuration.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/get_suspicious_ip_throttling
func (m *AttackProtectionManager) GetSuspiciousIPThrottling(
	ctx context.Context,
	opts ...RequestOption,
) (*SuspiciousIPThrottling, error) {
	var suspiciousIPThrottling SuspiciousIPThrottling
	err := m.management.Request(
		ctx,
		http.MethodGet,
		m.management.URI("attack-protection", "suspicious-ip-throttling"),
		&suspiciousIPThrottling,
		opts...,
	)

	return &suspiciousIPThrottling, err
}

// UpdateSuspiciousIPThrottling updates the suspicious IP throttling configuration.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/patch_suspicious_ip_throttling
func (m *AttackProtectionManager) UpdateSuspiciousIPThrottling(
	ctx context.Context,
	suspiciousIPThrottling *SuspiciousIPThrottling,
	opts ...RequestOption,
) error {
	return m.management.Request(
		ctx,
		http.MethodPatch,
		m.management.URI("attack-protection", "suspicious-ip-throttling"),
		suspiciousIPThrottling,
		opts...,
	)
}
