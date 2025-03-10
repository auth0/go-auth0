package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestAttackProtection(t *testing.T) {
	t.Run("Get breached password detection settings", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		breachedPasswordDetection, err := api.AttackProtection.GetBreachedPasswordDetection(context.Background())
		assert.NoError(t, err)
		assert.IsType(t, &BreachedPasswordDetection{}, breachedPasswordDetection)
	})

	t.Run("Update breached password detection settings", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		// Save initial settings.
		preTestBPDSettings, err := api.AttackProtection.GetBreachedPasswordDetection(context.Background())
		assert.NoError(t, err)

		expected := &BreachedPasswordDetection{
			Enabled: auth0.Bool(true),
			Method:  auth0.String("standard"),
			Stage: &BreachedPasswordDetectionStage{
				PreUserRegistration: &BreachedPasswordDetectionPreUserRegistration{
					Shields: &[]string{"block"},
				},
				PreChangePassword: &BreachedPasswordDetectionPreChangePassword{
					Shields: &[]string{"block"},
				},
			},
		}

		err = api.AttackProtection.UpdateBreachedPasswordDetection(context.Background(), expected)
		assert.NoError(t, err)

		actual, err := api.AttackProtection.GetBreachedPasswordDetection(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEnabled(), actual.GetEnabled())
		assert.Equal(t, expected.GetMethod(), actual.GetMethod())
		assert.Equal(t, expected.GetStage().GetPreUserRegistration().GetShields(), actual.GetStage().GetPreUserRegistration().GetShields())
		assert.Equal(t, expected.GetStage().GetPreChangePassword().GetShields(), actual.GetStage().GetPreChangePassword().GetShields())
		// Restore initial settings.
		err = api.AttackProtection.UpdateBreachedPasswordDetection(context.Background(), preTestBPDSettings)
		assert.NoError(t, err)
	})

	t.Run("Get the brute force configuration", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		bruteForceProtection, err := api.AttackProtection.GetBruteForceProtection(context.Background())
		assert.NoError(t, err)
		assert.IsType(t, &BruteForceProtection{}, bruteForceProtection)
	})

	t.Run("Update the brute force configuration", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		// Save initial settings.
		preTestBFPSettings, err := api.AttackProtection.GetBruteForceProtection(context.Background())
		assert.NoError(t, err)

		expected := &BruteForceProtection{
			Enabled:     auth0.Bool(true),
			MaxAttempts: auth0.Int(10),
		}

		err = api.AttackProtection.UpdateBruteForceProtection(context.Background(), expected)
		assert.NoError(t, err)

		actual, err := api.AttackProtection.GetBruteForceProtection(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEnabled(), actual.GetEnabled())
		assert.Equal(t, expected.GetMaxAttempts(), actual.GetMaxAttempts())

		// Restore initial settings.
		err = api.AttackProtection.UpdateBruteForceProtection(context.Background(), preTestBFPSettings)
		assert.NoError(t, err)
	})

	t.Run("Get the suspicious IP throttling configuration", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		suspiciousIPThrottling, err := api.AttackProtection.GetSuspiciousIPThrottling(context.Background())
		assert.NoError(t, err)
		assert.IsType(t, &SuspiciousIPThrottling{}, suspiciousIPThrottling)
	})

	t.Run("Update the suspicious IP throttling configuration", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		// Save initial settings.
		preTestSIPSettings, err := api.AttackProtection.GetSuspiciousIPThrottling(context.Background())
		assert.NoError(t, err)

		expected := &SuspiciousIPThrottling{
			Enabled: auth0.Bool(true),
			Stage: &Stage{
				PreLogin: &PreLogin{
					MaxAttempts: auth0.Int(100),
					Rate:        auth0.Int(864000),
				},
				PreUserRegistration: &PreUserRegistration{
					MaxAttempts: auth0.Int(50),
					Rate:        auth0.Int(1200),
				},
			},
		}

		err = api.AttackProtection.UpdateSuspiciousIPThrottling(context.Background(), expected)
		assert.NoError(t, err)

		actual, err := api.AttackProtection.GetSuspiciousIPThrottling(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEnabled(), actual.GetEnabled())
		assert.Equal(t, expected.GetStage().GetPreLogin().GetRate(), actual.GetStage().GetPreLogin().GetRate())
		assert.Equal(t, expected.GetStage().GetPreLogin().GetMaxAttempts(), actual.GetStage().GetPreLogin().GetMaxAttempts())
		assert.Equal(t, expected.GetStage().GetPreUserRegistration().GetRate(), actual.GetStage().GetPreUserRegistration().GetRate())
		assert.Equal(t, expected.GetStage().GetPreUserRegistration().GetMaxAttempts(), actual.GetStage().GetPreUserRegistration().GetMaxAttempts())

		// Restore initial settings.
		err = api.AttackProtection.UpdateSuspiciousIPThrottling(context.Background(), preTestSIPSettings)
		assert.NoError(t, err)
	})
}
