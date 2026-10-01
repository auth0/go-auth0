package authentication

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/v3/internal/client"

	"github.com/auth0/go-auth0/v3/authentication/oauth"
	"github.com/auth0/go-auth0/v3/authentication/passwordless"
)

func TestSendEmail(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	r, err := authAPI.Passwordless.SendEmail(context.Background(), passwordless.SendEmailRequest{
		Email: "test-email@example.com",
		Send:  "code",
	})

	assert.NoError(t, err)
	assert.Equal(t, "test-email@example.com", r.Email)
	assert.Equal(t, true, r.EmailVerified)
}

func TestLoginWithEmail(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	token, err := authAPI.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
		Code:     "123456",
		Email:    "test-email@example.com",
		Scope:    "openid profile email offline_access",
		Audience: "https://api.example.com",
	}, oauth.IDTokenValidationOptions{})

	assert.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func TestSendSMS(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	r, err := authAPI.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
		PhoneNumber: "+123456789",
	})

	assert.NoError(t, err)
	assert.Equal(t, "+123456789", r.PhoneNumber)
	assert.Equal(t, true, r.PhoneVerified)
}

func TestLoginWithSMS(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	token, err := authAPI.Passwordless.LoginWithSMS(context.Background(), passwordless.LoginWithSMSRequest{
		PhoneNumber: "+123456789",
		Code:        "123456",
		Scope:       "openid profile email offline_access",
		Audience:    "https://api.example.com",
	}, oauth.IDTokenValidationOptions{})

	assert.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func TestPasswordlessWithIDTokenVerification(t *testing.T) {
	t.Run("error for an invalid organization when using org_id", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"org_id": "org_124",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{Organization: "org_456"})

		assert.ErrorContains(t, err, "org_id claim value mismatch in the ID token")
	})

	t.Run("error for an invalid organization when using org_name", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"org_name": "wrong-org",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{Organization: "right-org"})

		assert.ErrorContains(t, err, "org_name claim value mismatch in the ID token")
	})

	t.Run("error for an invalid nonce", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"nonce": "wrong-nonce",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithEmail(context.Background(), passwordless.LoginWithEmailRequest{
			Code:     "123456",
			Email:    "test-email@example.com",
			Scope:    "openid profile email offline_access",
			Audience: "https://api.example.com",
		}, oauth.IDTokenValidationOptions{Nonce: "test-nonce"})

		assert.ErrorContains(t, err, "nonce claim value mismatch in the ID token; expected")
	})

	t.Run("error for an invalid maxage", func(t *testing.T) {
		skipE2E(t)

		extras := map[string]interface{}{
			"auth_time": time.Now().Add(-500 * time.Second).Unix(),
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.Passwordless.LoginWithSMS(context.Background(), passwordless.LoginWithSMSRequest{
			PhoneNumber: "+123456789",
			Code:        "123456",
			Scope:       "openid profile email offline_access",
			Audience:    "https://api.example.com",
		}, oauth.IDTokenValidationOptions{MaxAge: 100 * time.Second})

		assert.ErrorContains(t, err, "auth_time claim in the ID token indicates that too much time has passed")
	})
}

func TestPasswordlessWithClientAssertion(t *testing.T) {
	t.Run("Should support using private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
			WithClientAssertion(jwtPrivateKey, "RS256"),
		)

		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		r, err := api.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
			PhoneNumber: "+123456789",
		})

		assert.NoError(t, err)
		assert.Equal(t, "+123456789", r.PhoneNumber)
		assert.True(t, r.PhoneVerified)
	})

	t.Run("Should support passing private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
		)
		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		auth, err := client.CreateClientAssertion("RS256", jwtPrivateKey, clientID, "https://"+domain+"/")
		require.NoError(t, err)

		r, err := api.Passwordless.SendSMS(context.Background(), passwordless.SendSMSRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientAssertion:     auth,
				ClientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			},
			PhoneNumber: "+123456789",
		})

		assert.NoError(t, err)
		assert.Equal(t, "+123456789", r.PhoneNumber)
		assert.True(t, r.PhoneVerified)
	})
}

const testDBConnection = "Username-Password-Authentication"

func TestPasswordlessChallengeWithEmail(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	r, err := authAPI.Passwordless.ChallengeWithEmail(context.Background(), passwordless.ChallengeWithEmailRequest{
		Connection: testDBConnection,
		Email:      "test-email@example.com",
	})

	require.NoError(t, err)
	assert.NotEmpty(t, r.AuthSession)
	assert.Equal(t, "test-auth_session", r.AuthSession)
}

func TestPasswordlessChallengeWithEmailValidationErrors(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	_, err := authAPI.Passwordless.ChallengeWithEmail(context.Background(), passwordless.ChallengeWithEmailRequest{
		Connection: testDBConnection,
		Email:      "not-an-email",
	})

	require.Error(t, err)

	var authErr *Error
	require.True(t, errors.As(err, &authErr))
	assert.Equal(t, 400, authErr.StatusCode)
	assert.Equal(t, "invalid_request", authErr.Err)
	assert.Equal(t, []ValidationError{{Field: "email", Message: "Invalid email format"}}, authErr.GetValidationErrors())
}

func TestPasswordlessChallengeWithPhoneNumber(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	r, err := authAPI.Passwordless.ChallengeWithPhoneNumber(context.Background(), passwordless.ChallengeWithPhoneNumberRequest{
		Connection:     testDBConnection,
		PhoneNumber:    "+14155550100",
		DeliveryMethod: "voice",
	})

	require.NoError(t, err)
	assert.Equal(t, "test-auth_session", r.AuthSession)
}

func TestPasswordlessChallengeAllowSignup(t *testing.T) {
	t.Run("sends allow_signup true when set", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		r, err := authAPI.Passwordless.ChallengeWithEmail(context.Background(), passwordless.ChallengeWithEmailRequest{
			Connection:  testDBConnection,
			Email:       "test-email@example.com",
			AllowSignup: true,
		})

		require.NoError(t, err)
		assert.NotEmpty(t, r.AuthSession)
	})

	t.Run("sends allow_signup false by default", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		r, err := authAPI.Passwordless.ChallengeWithEmail(context.Background(), passwordless.ChallengeWithEmailRequest{
			Connection: testDBConnection,
			Email:      "test-email@example.com",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, r.AuthSession)
	})
}

func TestPasswordlessLoginWithOTPChallenge(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	token, err := authAPI.Passwordless.LoginWithOTPChallenge(context.Background(), passwordless.LoginWithOTPChallengeRequest{
		AuthSession: "test-auth_session",
		OTP:         "test-otp",
		Scope:       "openid profile email offline_access",
		Audience:    "https://api.example.com",
	}, oauth.IDTokenValidationOptions{})

	require.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func TestPasswordlessLoginWithOTPChallengeExtraParameters(t *testing.T) {
	skipE2E(t)
	configureHTTPTestRecordings(t, authAPI)

	token, err := authAPI.Passwordless.LoginWithOTPChallenge(context.Background(), passwordless.LoginWithOTPChallengeRequest{
		AuthSession: "test-auth_session",
		OTP:         "test-otp",
		Scope:       "openid profile email offline_access",
		Audience:    "https://api.example.com",
		ExtraParameters: map[string]string{
			"extra_param": "extra_value",
		},
	}, oauth.IDTokenValidationOptions{})

	require.NoError(t, err)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func TestPasswordlessLoginWithOTPChallengeErrors(t *testing.T) {
	login := func() error {
		_, err := authAPI.Passwordless.LoginWithOTPChallenge(context.Background(), passwordless.LoginWithOTPChallengeRequest{
			AuthSession: "test-auth_session",
			OTP:         "test-otp",
		}, oauth.IDTokenValidationOptions{})

		return err
	}

	t.Run("returns an error for a wrong otp", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		err := login()

		var authErr *Error
		require.True(t, errors.As(err, &authErr))
		assert.Equal(t, 400, authErr.StatusCode)
	})

	t.Run("returns an error for an expired session", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		err := login()

		var authErr *Error
		require.True(t, errors.As(err, &authErr))
		assert.Equal(t, 400, authErr.StatusCode)
	})

	t.Run("returns the mfa token when mfa is required", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		err := login()

		var authErr *Error
		require.True(t, errors.As(err, &authErr))
		assert.Equal(t, 403, authErr.StatusCode)
		assert.Equal(t, "mfa_required", authErr.Err)
		assert.NotEmpty(t, authErr.GetMFAToken())
		require.NotNil(t, authErr.GetMFARequirements())
		require.Len(t, authErr.GetMFARequirements().Challenge, 1)
		assert.Equal(t, "otp", authErr.GetMFARequirements().Challenge[0].Type)
	})
}

func TestPasswordlessOTPChallengeMissingFields(t *testing.T) {
	noClientAPI, err := New(context.Background(), domain)
	require.NoError(t, err)

	t.Run("ChallengeWithEmail", func(t *testing.T) {
		tests := []struct {
			name     string
			api      *Authentication
			req      passwordless.ChallengeWithEmailRequest
			expected string
		}{
			{"missing client id", noClientAPI, passwordless.ChallengeWithEmailRequest{Connection: testDBConnection, Email: "a@example.com"}, "missing required fields: ClientID"},
			{"missing connection", authAPI, passwordless.ChallengeWithEmailRequest{Email: "a@example.com"}, "missing required fields: Connection"},
			{"missing email", authAPI, passwordless.ChallengeWithEmailRequest{Connection: testDBConnection}, "missing required fields: Email"},
			{"missing all", noClientAPI, passwordless.ChallengeWithEmailRequest{}, "missing required fields: ClientID, Connection, Email"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				_, err := tc.api.Passwordless.ChallengeWithEmail(context.Background(), tc.req)
				assert.EqualError(t, err, tc.expected)
			})
		}
	})

	t.Run("ChallengeWithPhoneNumber", func(t *testing.T) {
		tests := []struct {
			name     string
			api      *Authentication
			req      passwordless.ChallengeWithPhoneNumberRequest
			expected string
		}{
			{"missing client id", noClientAPI, passwordless.ChallengeWithPhoneNumberRequest{Connection: testDBConnection, PhoneNumber: "+14155550100"}, "missing required fields: ClientID"},
			{"missing connection", authAPI, passwordless.ChallengeWithPhoneNumberRequest{PhoneNumber: "+14155550100"}, "missing required fields: Connection"},
			{"missing phone number", authAPI, passwordless.ChallengeWithPhoneNumberRequest{Connection: testDBConnection}, "missing required fields: PhoneNumber"},
			{"missing all", noClientAPI, passwordless.ChallengeWithPhoneNumberRequest{}, "missing required fields: ClientID, Connection, PhoneNumber"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				_, err := tc.api.Passwordless.ChallengeWithPhoneNumber(context.Background(), tc.req)
				assert.EqualError(t, err, tc.expected)
			})
		}
	})

	t.Run("LoginWithOTPChallenge", func(t *testing.T) {
		tests := []struct {
			name     string
			api      *Authentication
			req      passwordless.LoginWithOTPChallengeRequest
			expected string
		}{
			{"missing auth session", authAPI, passwordless.LoginWithOTPChallengeRequest{OTP: "test-otp"}, "missing required fields: AuthSession"},
			{"missing otp", authAPI, passwordless.LoginWithOTPChallengeRequest{AuthSession: "test-auth_session"}, "missing required fields: OTP"},
			{"missing client id", noClientAPI, passwordless.LoginWithOTPChallengeRequest{AuthSession: "test-auth_session", OTP: "test-otp"}, "missing required fields: ClientID"},
			{"missing all", noClientAPI, passwordless.LoginWithOTPChallengeRequest{}, "missing required fields: AuthSession, OTP, ClientID"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				_, err := tc.api.Passwordless.LoginWithOTPChallenge(context.Background(), tc.req, oauth.IDTokenValidationOptions{})
				assert.EqualError(t, err, tc.expected)
			})
		}
	})
}

func TestPasswordlessChallengeWithPhoneNumberInvalidE164(t *testing.T) {
	for _, number := range []string{"5555550123", "+0123", "+1 555 555 0123"} {
		t.Run(number, func(t *testing.T) {
			_, err := authAPI.Passwordless.ChallengeWithPhoneNumber(context.Background(), passwordless.ChallengeWithPhoneNumberRequest{
				Connection:  testDBConnection,
				PhoneNumber: number,
			})

			assert.EqualError(t, err, "PhoneNumber must be in E.164 format (e.g. +14155550100)")
		})
	}
}

func TestPasswordlessOTPChallengeWithClientAssertion(t *testing.T) {
	t.Run("challenge supports using private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
			WithClientAssertion(jwtPrivateKey, "RS256"),
		)
		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		r, err := api.Passwordless.ChallengeWithEmail(context.Background(), passwordless.ChallengeWithEmailRequest{
			Connection: testDBConnection,
			Email:      "test-email@example.com",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, r.AuthSession)
	})

	t.Run("login supports using private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
			WithClientAssertion(jwtPrivateKey, "RS256"),
		)
		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		token, err := api.Passwordless.LoginWithOTPChallenge(context.Background(), passwordless.LoginWithOTPChallengeRequest{
			AuthSession: "test-auth_session",
			OTP:         "test-otp",
		}, oauth.IDTokenValidationOptions{})

		require.NoError(t, err)
		assert.NotEmpty(t, token.AccessToken)
	})
}

func TestPasswordlessLoginWithOTPChallengeIDTokenVerification(t *testing.T) {
	login := func(api *Authentication, opts oauth.IDTokenValidationOptions) error {
		_, err := api.Passwordless.LoginWithOTPChallenge(context.Background(), passwordless.LoginWithOTPChallengeRequest{
			AuthSession: "test-auth_session",
			OTP:         "test-otp",
			Scope:       "openid profile email offline_access",
		}, opts)

		return err
	}

	t.Run("error for an invalid nonce", func(t *testing.T) {
		skipE2E(t)

		api, err := withIDToken(t, map[string]interface{}{"nonce": "wrong-nonce"})
		require.NoError(t, err)

		err = login(api, oauth.IDTokenValidationOptions{Nonce: "test-nonce"})

		assert.ErrorContains(t, err, "nonce claim value mismatch in the ID token; expected")
	})

	t.Run("error for an invalid organization", func(t *testing.T) {
		skipE2E(t)

		api, err := withIDToken(t, map[string]interface{}{"org_id": "org_124"})
		require.NoError(t, err)

		err = login(api, oauth.IDTokenValidationOptions{Organization: "org_456"})

		assert.ErrorContains(t, err, "org_id claim value mismatch in the ID token")
	})
}
