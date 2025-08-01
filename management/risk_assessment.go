package management

import "context"

// RiskAssessmentSettings : Risk Assessment Settings.
type RiskAssessmentSettings struct {
	// Enabled indicates whether the risk assessment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// RiskAssessmentSettingsNewDevice : Risk Assessment Settings for New Device.
type RiskAssessmentSettingsNewDevice struct {
	// RememberFor is the Length of time to remember devices for, in days.
	RememberFor *int `json:"remember_for,omitempty"`
}

// RiskAssessmentManager manages Auth0 Risk Assessment resources.
type RiskAssessmentManager manager

// ReadSettings retrieves the risk assessment settings.
func (m *RiskAssessmentManager) ReadSettings(ctx context.Context, opts ...RequestOption) (setting *RiskAssessmentSettings, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("risk-assessments", "settings"), &setting, opts...)
	return
}

// UpdateSettings updates the risk assessment settings.
func (m *RiskAssessmentManager) UpdateSettings(ctx context.Context, setting *RiskAssessmentSettings, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PATCH", m.management.URI("risk-assessments", "settings"), setting, opts...)
	return
}

// ReadNewDeviceSettings retrieves the risk assessment settings for new devices.
func (m *RiskAssessmentManager) ReadNewDeviceSettings(ctx context.Context, opts ...RequestOption) (setting *RiskAssessmentSettingsNewDevice, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("risk-assessments", "settings", "new-device"), &setting, opts...)
	return
}

// UpdateNewDeviceSettings updates the risk assessment settings for new devices.
func (m *RiskAssessmentManager) UpdateNewDeviceSettings(ctx context.Context, setting *RiskAssessmentSettingsNewDevice, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "PATCH", m.management.URI("risk-assessments", "settings", "new-device"), setting, opts...)
	return
}
