package management

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestAttackProtection(t *testing.T) {
	t.Run("Get breached password detection settings", func(t *testing.T) {
		setupHTTPRecordings(t)

		breachedPasswordDetection, err := m.AttackProtection.GetBreachedPasswordDetection()
		assert.NoError(t, err)
		assert.IsType(t, &BreachedPasswordDetection{}, breachedPasswordDetection)
	})

	t.Run("Update breached password detection settings", func(t *testing.T) {
		setupHTTPRecordings(t)

		// Save initial settings.
		preTestBPDSettings, err := m.AttackProtection.GetBreachedPasswordDetection()
		assert.NoError(t, err)

		expected := &BreachedPasswordDetection{
			Enabled: auth0.Bool(true),
			Method:  auth0.String("standard"),
		}

		err = m.AttackProtection.UpdateBreachedPasswordDetection(expected)
		assert.NoError(t, err)

		actual, err := m.AttackProtection.GetBreachedPasswordDetection()
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEnabled(), actual.GetEnabled())
		assert.Equal(t, expected.GetMethod(), actual.GetMethod())

		// Restore initial settings.
		err = m.AttackProtection.UpdateBreachedPasswordDetection(preTestBPDSettings)
		assert.NoError(t, err)
	})

	t.Run("Get the brute force configuration", func(t *testing.T) {
		setupHTTPRecordings(t)

		bruteForceProtection, err := m.AttackProtection.GetBruteForceProtection()
		assert.NoError(t, err)
		assert.IsType(t, &BruteForceProtection{}, bruteForceProtection)
	})

	t.Run("Update the brute force configuration", func(t *testing.T) {
		setupHTTPRecordings(t)

		// Save initial settings.
		preTestBFPSettings, err := m.AttackProtection.GetBruteForceProtection()
		assert.NoError(t, err)

		expected := &BruteForceProtection{
			Enabled:     auth0.Bool(true),
			MaxAttempts: auth0.Int(10),
		}

		err = m.AttackProtection.UpdateBruteForceProtection(expected)
		assert.NoError(t, err)

		actual, err := m.AttackProtection.GetBruteForceProtection()
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEnabled(), actual.GetEnabled())
		assert.Equal(t, expected.GetMaxAttempts(), actual.GetMaxAttempts())

		// Restore initial settings.
		err = m.AttackProtection.UpdateBruteForceProtection(preTestBFPSettings)
		assert.NoError(t, err)
	})

	t.Run("Get the suspicious IP throttling configuration", func(t *testing.T) {
		setupHTTPRecordings(t)

		suspiciousIPThrottling, err := m.AttackProtection.GetSuspiciousIPThrottling()
		assert.NoError(t, err)
		assert.IsType(t, &SuspiciousIPThrottling{}, suspiciousIPThrottling)
	})

	t.Run("Update the suspicious IP throttling configuration", func(t *testing.T) {
		setupHTTPRecordings(t)

		// Save initial settings.
		preTestSIPSettings, err := m.AttackProtection.GetSuspiciousIPThrottling()
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

		err = m.AttackProtection.UpdateSuspiciousIPThrottling(expected)
		assert.NoError(t, err)

		actual, err := m.AttackProtection.GetSuspiciousIPThrottling()
		assert.NoError(t, err)
		assert.Equal(t, expected.GetEnabled(), actual.GetEnabled())
		assert.Equal(t, expected.GetStage().GetPreLogin().GetRate(), actual.GetStage().GetPreLogin().GetRate())
		assert.Equal(t, expected.GetStage().GetPreLogin().GetMaxAttempts(), actual.GetStage().GetPreLogin().GetMaxAttempts())
		assert.Equal(t, expected.GetStage().GetPreUserRegistration().GetRate(), actual.GetStage().GetPreUserRegistration().GetRate())
		assert.Equal(t, expected.GetStage().GetPreUserRegistration().GetMaxAttempts(), actual.GetStage().GetPreUserRegistration().GetMaxAttempts())

		// Restore initial settings.
		err = m.AttackProtection.UpdateSuspiciousIPThrottling(preTestSIPSettings)
		assert.NoError(t, err)
	})
}
