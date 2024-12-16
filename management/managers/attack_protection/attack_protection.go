package attackProtection

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

// GetBreachedPasswordDetection Get breached password detection settings
//
// https://auth0.com/docs/api/management/v2/#!/AttackProtection/get_breached_password_detection
func (m *Manager) GetBreachedPasswordDetection(ctx context.Context, opts ...management.RequestOption) (*models.GetBreachedPasswordDetection200Response, error) {
	var localVarReturnValue *models.GetBreachedPasswordDetection200Response
	err := m.management.Request(ctx, "GET", m.management.URI("attack-protection", "breached-password-detection"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetBruteForceProtection Get the brute force configuration
//
// https://auth0.com/docs/api/management/v2/#!/AttackProtection/get_brute_force_protection
func (m *Manager) GetBruteForceProtection(ctx context.Context, opts ...management.RequestOption) (*models.GetBruteForceProtection200Response, error) {
	var localVarReturnValue *models.GetBruteForceProtection200Response
	err := m.management.Request(ctx, "GET", m.management.URI("attack-protection", "brute-force-protection"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSuspiciousIpThrottling Get the suspicious IP throttling configuration
//
// https://auth0.com/docs/api/management/v2/#!/AttackProtection/get_suspicious_ip_throttling
func (m *Manager) GetSuspiciousIpThrottling(ctx context.Context, opts ...management.RequestOption) (*models.GetSuspiciousIpThrottling200Response, error) {
	var localVarReturnValue *models.GetSuspiciousIpThrottling200Response
	err := m.management.Request(ctx, "GET", m.management.URI("attack-protection", "suspicious-ip-throttling"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchBreachedPasswordDetection Update breached password detection settings
//
// https://auth0.com/docs/api/management/v2/#!/AttackProtection/patch_breached_password_detection
func (m *Manager) PatchBreachedPasswordDetection(ctx context.Context, patchBreachedPasswordDetectionRequest *models.PatchBreachedPasswordDetectionRequest, opts ...management.RequestOption) (*models.GetBreachedPasswordDetection200Response, error) {
	var localVarReturnValue *models.GetBreachedPasswordDetection200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("attack-protection", "breached-password-detection"), patchBreachedPasswordDetectionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchBruteForceProtection Update the brute force configuration
//
// https://auth0.com/docs/api/management/v2/#!/AttackProtection/patch_brute_force_protection
func (m *Manager) PatchBruteForceProtection(ctx context.Context, patchBruteForceProtectionRequest *models.PatchBruteForceProtectionRequest, opts ...management.RequestOption) (*models.GetBruteForceProtection200Response, error) {
	var localVarReturnValue *models.GetBruteForceProtection200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("attack-protection", "brute-force-protection"), patchBruteForceProtectionRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchSuspiciousIpThrottling Update the suspicious IP throttling configuration
//
// https://auth0.com/docs/api/management/v2/#!/AttackProtection/patch_suspicious_ip_throttling
func (m *Manager) PatchSuspiciousIpThrottling(ctx context.Context, patchSuspiciousIpThrottlingRequest *models.PatchSuspiciousIpThrottlingRequest, opts ...management.RequestOption) (*models.GetSuspiciousIpThrottling200Response, error) {
	var localVarReturnValue *models.GetSuspiciousIpThrottling200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("attack-protection", "suspicious-ip-throttling"), patchSuspiciousIpThrottlingRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
