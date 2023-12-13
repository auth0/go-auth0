package authentication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/authentication/mfa"
	"github.com/auth0/go-auth0/authentication/oauth"
)

func TestMFAChallenge(t *testing.T) {
	t.Run("Should require ClientID, MFAToken, and ChallengeType", func(t *testing.T) {
		auth, err := New(
			context.Background(),
			domain,
		)
		require.NoError(t, err)

		_, err = auth.MFA.Challenge(context.Background(), mfa.ChallengeRequest{})
		assert.ErrorContains(t, err, "Missing required fields: ClientID, MFAToken, ChallengeType")
	})

	t.Run("Should make a challenge request using OTP", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		r, err := authAPI.MFA.Challenge(context.Background(), mfa.ChallengeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken:        "mfa-token",
			ChallengeType:   "otp",
			AuthenticatorID: "totp|dev_id",
		})

		require.NoError(t, err)
		assert.Equal(t, "otp", r.ChallengeType)
	})

	t.Run("Should make a challenge request using OOB", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		r, err := authAPI.MFA.Challenge(context.Background(), mfa.ChallengeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken:        "mfa-token",
			ChallengeType:   "oob",
			AuthenticatorID: "push|dev_ADbB4Y2ozSOwynKu",
		})

		require.NoError(t, err)
		assert.Equal(t, "oob", r.ChallengeType)
	})
}

func TestMFAVerifyWithOTP(t *testing.T) {
	t.Run("Should require ClientID, MFAToken, and OTP", func(t *testing.T) {
		auth, err := New(
			context.Background(),
			domain,
		)
		require.NoError(t, err)

		_, err = auth.MFA.VerifyWithOTP(context.Background(), mfa.VerifyWithOTPRequest{})
		assert.ErrorContains(t, err, "Missing required fields: ClientID, MFAToken, OTP")
	})

	t.Run("Should return tokens for a valid request", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenset, err := authAPI.MFA.VerifyWithOTP(context.Background(), mfa.VerifyWithOTPRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken: "mfa-token",
			OTP:      "853860",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, tokenset)
	})
}

func TestMFAVerifyWithOOB(t *testing.T) {
	t.Run("Should require ClientID, MFAToken, and OOB", func(t *testing.T) {
		auth, err := New(
			context.Background(),
			domain,
		)
		require.NoError(t, err)

		_, err = auth.MFA.VerifyWithOOB(context.Background(), mfa.VerifyWithOOBRequest{})
		assert.ErrorContains(t, err, "Missing required fields: ClientID, MFAToken, OOB")
	})

	t.Run("Should return an error when requesting before authorizing", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.MFA.VerifyWithOOB(context.Background(), mfa.VerifyWithOOBRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken: "mfa-token",
			OOBCode:  "oob-token",
		})

		assert.ErrorContains(t, err, "Authorization pending: please repeat the request in a few seconds.")
	})

	t.Run("Should return an error when requesting when authorization has been denied", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.MFA.VerifyWithOOB(context.Background(), mfa.VerifyWithOOBRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken: "mfa-token",
			OOBCode:  "oob-token",
		})

		assert.ErrorContains(t, err, "MFA Authorization rejected")
	})

	t.Run("Should return tokens when authorization is approved", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenset, err := authAPI.MFA.VerifyWithOOB(context.Background(), mfa.VerifyWithOOBRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken: "mfa-token",
			OOBCode:  "oob-token",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, tokenset)
	})
}

func TestMFAVerifyWithRecoveryCode(t *testing.T) {
	t.Run("Should require ClientID, MFAToken, and OOB", func(t *testing.T) {
		auth, err := New(
			context.Background(),
			domain,
		)
		require.NoError(t, err)

		_, err = auth.MFA.VerifyWithRecoveryCode(context.Background(), mfa.VerifyWithRecoveryCodeRequest{})
		assert.ErrorContains(t, err, "Missing required fields: ClientID, MFAToken, RecoveryCode")
	})

	t.Run("Should return tokens for a valid request", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenset, err := authAPI.MFA.VerifyWithRecoveryCode(context.Background(), mfa.VerifyWithRecoveryCodeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			MFAToken:     "mfa-token",
			RecoveryCode: "M67LWV6ZJGNDUGH43BADW46N",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, tokenset)
		assert.NotEmpty(t, tokenset.RecoveryCode)
	})
}
