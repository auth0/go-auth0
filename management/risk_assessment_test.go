package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestRiskAssessmentManager_ReadSettings(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Read the risk assessment settings
	settings, err := api.RiskAssessment.ReadSettings(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, settings)
	assert.NotNil(t, settings.GetEnabled())
	assert.Equal(t, settings.GetEnabled(), false)
}

func TestRiskAssessmentManager_UpdateSettings(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Update the risk assessment settings
	settings := &RiskAssessmentSettings{
		Enabled: auth0.Bool(true),
	}
	err := api.RiskAssessment.UpdateSettings(context.Background(), settings)
	assert.NoError(t, err)

	// Verify the update
	updatedSettings, err := api.RiskAssessment.ReadSettings(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, updatedSettings)
	assert.Equal(t, updatedSettings.GetEnabled(), true)
	t.Cleanup(func() {
		cleanupRiskAssessmentSettings(t)
	})
}

func TestRiskAssessmentManager_ReadNewDeviceSettings(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Read the risk assessment settings for new devices
	newDeviceSettings, err := api.RiskAssessment.ReadNewDeviceSettings(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, newDeviceSettings)
	assert.NotNil(t, newDeviceSettings.GetRememberFor())
	assert.Equal(t, newDeviceSettings.GetRememberFor(), 30) // Default value is 30 days
}

func TestRiskAssessmentManager_UpdateNewDeviceSettings(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Update the risk assessment settings for new devices
	newDeviceSettings := &RiskAssessmentSettingsNewDevice{
		RememberFor: auth0.Int(60), // Set to 60 days
	}
	err := api.RiskAssessment.UpdateNewDeviceSettings(context.Background(), newDeviceSettings)
	assert.NoError(t, err)

	// Verify the update
	updatedNewDeviceSettings, err := api.RiskAssessment.ReadNewDeviceSettings(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, updatedNewDeviceSettings)
	assert.Equal(t, updatedNewDeviceSettings.GetRememberFor(), 60)
	t.Cleanup(func() {
		cleanupRiskAssessmentNewDeviceSettings(t)
	})
}

func cleanupRiskAssessmentSettings(t *testing.T) {
	// Reset the risk assessment settings to default
	defaultSettings := &RiskAssessmentSettings{
		Enabled: auth0.Bool(false),
	}
	err := api.RiskAssessment.UpdateSettings(context.Background(), defaultSettings)
	assert.NoError(t, err)
}

func cleanupRiskAssessmentNewDeviceSettings(t *testing.T) {
	// Reset the risk assessment settings for new devices to default
	defaultNewDeviceSettings := &RiskAssessmentSettingsNewDevice{
		RememberFor: auth0.Int(30), // Default value is 30 days
	}
	err := api.RiskAssessment.UpdateNewDeviceSettings(context.Background(), defaultNewDeviceSettings)
	assert.NoError(t, err)
}
