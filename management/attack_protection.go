package management

import (
	"net/http"
)

// AttackProtectionManager manages Auth0 Attack Protection settings.
//
// See: https://auth0.com/docs/secure/attack-protection
type AttackProtectionManager struct {
	*Management
}

func newAttackProtectionManager(m *Management) *AttackProtectionManager {
	return &AttackProtectionManager{m}
}

// BreachedPasswordDetection protects applications from
// bad actors logging in with stolen credentials.
//
// See: https://auth0.com/docs/secure/attack-protection/breached-password-detection
type BreachedPasswordDetection struct {
	Enabled                    *bool     `json:"enabled,omitempty"`
	Shields                    *[]string `json:"shields,omitempty"`
	AdminNotificationFrequency *[]string `json:"admin_notification_frequency,omitempty"`
	Method                     *string   `json:"method,omitempty"`
}

// GetBreachedPasswordDetection retrieves breached password detection settings.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/api/management/v2#!/Attack_Protection/get_breached_password_detection
func (m *AttackProtectionManager) GetBreachedPasswordDetection(
	opts ...RequestOption,
) (*BreachedPasswordDetection, error) {
	var breachedPasswordDetection BreachedPasswordDetection
	err := m.Request(
		http.MethodGet,
		m.URI("attack-protection", "breached-password-detection"),
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
	breachedPasswordDetection *BreachedPasswordDetection,
	opts ...RequestOption,
) error {
	return m.Request(
		http.MethodPatch,
		m.URI("attack-protection", "breached-password-detection"),
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
	opts ...RequestOption,
) (*BruteForceProtection, error) {
	var bruteForceProtection BruteForceProtection
	err := m.Request(
		http.MethodGet,
		m.URI("attack-protection", "brute-force-protection"),
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
	bruteForceProtection *BruteForceProtection,
	opts ...RequestOption,
) error {
	return m.Request(
		http.MethodPatch,
		m.URI("attack-protection", "brute-force-protection"),
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
	opts ...RequestOption,
) (*SuspiciousIPThrottling, error) {
	var suspiciousIPThrottling SuspiciousIPThrottling
	err := m.Request(
		http.MethodGet,
		m.URI("attack-protection", "suspicious-ip-throttling"),
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
	suspiciousIPThrottling *SuspiciousIPThrottling,
	opts ...RequestOption,
) error {
	return m.Request(
		http.MethodPatch,
		m.URI("attack-protection", "suspicious-ip-throttling"),
		suspiciousIPThrottling,
		opts...,
	)
}

// BotDetection mitigates scripted attacks by detecting
// when a request is likely to be coming from a bot.
//
// See: https://auth0.com/docs/secure/attack-protection/bot-detection
type BotDetection struct {
	Response *BotDetectionResponse
}

// BotDetectionResponse used to block suspected bot traffic by
// requiring a CAPTCHA during the login process.
type BotDetectionResponse struct {
	Policy    *string           `json:"policy,omitempty"`
	Selected  *string           `json:"selected,omitempty"`
	Providers *CaptchaProviders `json:"providers,omitempty"`
}

// CaptchaProviders holds all the available captcha providers.
type CaptchaProviders struct {
	Auth0               *CaptchaProviderAuth0               `json:"auth0,omitempty"`
	RecaptchaV2         *CaptchaProviderRecaptchaV2         `json:"recaptcha_v2,omitempty"`
	RecaptchaEnterprise *CaptchaProviderRecaptchaEnterprise `json:"recaptcha_enterprise,omitempty"`
}

// CaptchaProviderAuth0 currently has no settings.
type CaptchaProviderAuth0 struct{}

// CaptchaProviderRecaptchaV2 holds settings for RecaptchaV2.
type CaptchaProviderRecaptchaV2 struct {
	Secret  *string `json:"secret,omitempty"`
	SiteKey *string `json:"siteKey,omitempty"`
}

// CaptchaProviderRecaptchaEnterprise holds settings for RecaptchaEnterprise.
type CaptchaProviderRecaptchaEnterprise struct {
	APIKey    *string `json:"apiKey,omitempty"`
	ProjectID *string `json:"projectId,omitempty"`
	SiteKey   *string `json:"siteKey,omitempty"`
}

// GetBotDetection retrieves the bot detection settings.
//
// Required scope: `read:attack_protection`
//
// See: https://auth0.com/docs/secure/attack-protection/bot-detection#configure-bot-detection
func (m *AttackProtectionManager) GetBotDetection(
	opts ...RequestOption,
) (*BotDetection, error) {
	var botDetection BotDetection
	err := m.Request(http.MethodGet, m.URI("anomaly", "captchas"), &botDetection.Response, opts...)
	return &botDetection, err
}

// UpdateBotDetection updates the bot detection settings.
//
// Required scope: `read:attack_protection`
//
// https://auth0.com/docs/secure/attack-protection/bot-detection#configure-bot-detection
func (m *AttackProtectionManager) UpdateBotDetection(
	botDetection *BotDetection,
	opts ...RequestOption,
) error {
	return m.Request(http.MethodPost, m.URI("anomaly", "captchas"), &botDetection.Response, opts...)
}
