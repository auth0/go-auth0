package management

import (
	"testing"

	"github.com/auth0/go-auth0"
)

func TestAttackProtection(t *testing.T) {
	t.Run("Get breached password detection settings", func(t *testing.T) {
		breachedPasswordDetection, err := m.AttackProtection.GetBreachedPasswordDetection()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", breachedPasswordDetection)
	})

	t.Run("Update breached password detection settings", func(t *testing.T) {
		breachedPasswordDetection := &BreachedPasswordDetection{
			Enabled: auth0.Bool(true),
			Method:  auth0.String("standard"),
		}

		err := m.AttackProtection.UpdateBreachedPasswordDetection(breachedPasswordDetection)
		if err != nil {
			t.Error(err)
		}

		breachedPasswordDetection, err = m.AttackProtection.GetBreachedPasswordDetection()
		if err != nil {
			t.Error(err)
		}

		t.Logf("%+v\n", breachedPasswordDetection)
	})

	t.Run("Get the brute force configuration", func(t *testing.T) {
		bruteForceProtection, err := m.AttackProtection.GetBruteForceProtection()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", bruteForceProtection)
	})

	t.Run("Update the brute force configuration", func(t *testing.T) {
		bruteForceProtection := &BruteForceProtection{
			Enabled:     auth0.Bool(true),
			MaxAttempts: auth0.Int(10),
		}

		err := m.AttackProtection.UpdateBruteForceProtection(bruteForceProtection)
		if err != nil {
			t.Error(err)
		}

		bruteForceProtection, err = m.AttackProtection.GetBruteForceProtection()
		if err != nil {
			t.Error(err)
		}

		t.Logf("%+v\n", bruteForceProtection)
	})

	t.Run("Get the suspicious IP throttling configuration", func(t *testing.T) {
		suspiciousIPThrottling, err := m.AttackProtection.GetSuspiciousIPThrottling()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", suspiciousIPThrottling)
	})

	t.Run("Update the suspicious IP throttling configuration", func(t *testing.T) {
		suspiciousIPThrottling := &SuspiciousIPThrottling{
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

		err := m.AttackProtection.UpdateSuspiciousIPThrottling(suspiciousIPThrottling)
		if err != nil {
			t.Error(err)
		}

		suspiciousIPThrottling, err = m.AttackProtection.GetSuspiciousIPThrottling()
		if err != nil {
			t.Error(err)
		}

		t.Logf("%+v\n", suspiciousIPThrottling)
	})
}
