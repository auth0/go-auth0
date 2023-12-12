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
			MFAToken:     "Fe26.2*SERVER_1694679013*6327d15b8d0275969292f0a04bdfccfbfa628fea3ad63416a21d2ebabbe86ef2*V2p3Qn58h0z6J0oKt6bP8Q*m-Og0UIoY0kJaevfoqxYmxkfq8txoP_g6Sx2bDG_Rxc0dMsm9SryIQdGoz-lgJfMB-x_y1QJp9HDf7AehZlP5YmOo7BreIET5qaHnJuYYLM8uHyst58hXtiB0ib0k7oxdlerGvQYRf-xxavdd_ffm0CJfso6F9NJ4ts5H8sYXzVcs5uvfBqFESu9w4Kz6euV4SKxuF78dFMB-OIHeeJR7x61jdHikax8V9MivdFYoe5Hi3W8zrJywvrlowYIcSzPu_LEaDmkKcS82tr6ezrjr-4I6VCgka-qC2OUqPiUxvffYGbbAhI1poNSKbF66d0rraMPVS7C3ejFAmA4rDT9Y7BB5wtG_og5KWlcQdCPb_5HhUWz0kjIKmZHYVzlB2_TYZDnzcM5GWv_zag2CmVlDExLhVISjpzN5DoXM4-TDJdPMNKhnAQ5d5InYro9Fwum5CP8vkY3IWlb-1U8Ool25AUC9roFGQ-0_TS9L4pmyQHMGLP3Oxsk3YPfD_XPhqoopNJg2AagshYy8yypHc9916nokiNm9fjJk3ux8h4t6LJRXYnpxxVR0GC9VLc-hCjIssopMFXktk9BQi7yZKcRf8TKXuH4tyYjzOm8j7D7eziMBZszPIEppHzM5q_wcBvJvtYFT2FMnK3myd71QiF2KlaxwGghjLfwSEhQR1xW8E_dnok5zgLFQD-5rsEBIsD7_xwUdHzxf724EYY3G5HlkqJBHG5V5AdT8QcA7cxpHVRC8TvxZt932a5VS1Reay32l9YeqRsSrfbWYIETSzuApvpXnqSt2RFp_WYrfln4PcuQE1ctl6oKKBZhxN3agLp5pFutt6NQ1e-p5tIWDco5fEowI6zhfCMoeCbmDdpSh2v4bPmMWJcS1taK8F9ygIk1XtbFHIrY6kBIINL_WsgoO8OG8_w_HlCqNU-yyvPIFr6nSzDsAqc0hwBsCy7lI_ju1E0ZbWJK6a4rjbdZXzUG87bmuzrtWD7MsUZ5387nmL_2P6I06ckrqiu32mDrtJ3h_q67JS1obima0r86OHn8Zy8RKrPDkosjwXxVQL9eFePOOpTzfyL5I60QrKSxUxBNmldeZIsDl3B1YRUjlTVXyXWTYio6Wi1fCQzeZ3bBzBqpe_m-SIlRcQzL0B7laGeUqFroZPrt93Q4Pz9sYXdnCw2SwrvprIfbEfDpz0RfkaWH8_EFPso72Ii96LBXGoa2OdspoXfYSJpVDUTis4BsGJl01hnZ3SgNkvZNUnst0DvoWYPfODgQMT27ejHn9G7-R_q8p7Upy2j3DaoOOclZh1XaX44-Ogp6t57r0R5BmkdzSCpW8w35nPLGkFmoZiec8iuViDRa9LhuBHRNVVjl_y-kafasIWofJAC6StQ6pntwwm8sbcC8ALhtcjeRp5n-61fwVrtX5kwMyn-pUDli1XpwFtTa_zielhczadzrYOYD9NbhKf24pFtJj72SXiX3wpvhk6Pehy_fuHzsi5Ly8HjuuaRb5iMnL0Rj_Q37LELisy6OZ7Qcc7dpWhyU0kUGDJWJIlMzeU3iSBjH0zthmML7T-a4mtNSj2uIqaXTagt59RsDOMdpOAjDLv-IQ2HfRBMJEdXIczqWx3ROX7XFHHDOBho7D75domhT40xLtNBq9LCE6xDhcH_EDw1ifw_b_RWj0zSjiOozy_2VR5hubFaN-HJI-QJvCUplYPcGxfb8kUWYVjXvCmr8_MIaRa-_lv-xn-ABeLdPB3B-6RFXEQ*1702415848792*bf09812e2469b3346ecfb6cf22902b3484bb0e6f9e3704428970bcb195031317*CC1VR1vWgeGTehkNlBaW0Z1C6wboHpGuiizLGsnJomE",
			RecoveryCode: "M67LWV6ZJGNDUGH43BADW46N",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, tokenset)
		assert.NotEmpty(t, tokenset.RecoveryCode)
	})
}
