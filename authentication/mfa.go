package authentication

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/auth0/go-auth0/authentication/mfa"
	"github.com/auth0/go-auth0/authentication/oauth"
)

// MFA exposes requesting an MFA challenge and verifying MFA methods.
type MFA manager

// Challenge requests a challenge for multi-factor authentication (MFA) based on the challenge types supported by the application and user.
//
// See: https://auth0.com/docs/api/authentication#challenge-request
func (m *MFA) Challenge(ctx context.Context, body mfa.ChallengeRequest, opts ...RequestOption) (c *mfa.ChallengeResponse, err error) {
	missing := []string{}
	check(&missing, "ClientID", (body.ClientID != "" || m.authentication.clientID != ""))
	check(&missing, "MFAToken", body.MFAToken != "")
	check(&missing, "ChallengeType", body.ChallengeType != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("Missing required fields: %s", strings.Join(missing, ", "))
	}

	err = m.authentication.addClientAuthenticationToClientAuthStruct(&body.ClientAuthentication, false)

	if err != nil {
		return nil, err
	}

	err = m.authentication.Request(ctx, "POST", m.authentication.URI("mfa", "challenge"), body, &c, opts...)

	if err != nil {
		return nil, err
	}

	return
}

// VerifyWithOTP verifies an MFA challenge using a one-time password (OTP).
//
// See: https://auth0.com/docs/api/authentication#verify-with-one-time-password-otp-
func (m *MFA) VerifyWithOTP(ctx context.Context, body mfa.VerifyWithOTPRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	missing := []string{}
	check(&missing, "ClientID", (body.ClientID != "" || m.authentication.clientID != ""))
	check(&missing, "MFAToken", body.MFAToken != "")
	check(&missing, "OTP", body.OTP != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("Missing required fields: %s", strings.Join(missing, ", "))
	}

	data := url.Values{
		"mfa_token":  []string{body.MFAToken},
		"grant_type": []string{"http://auth0.com/oauth/grant-type/mfa-otp"},
		"otp":        []string{body.OTP},
	}

	err = m.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, true)

	if err != nil {
		return nil, err
	}

	err = m.authentication.Request(ctx, "POST", m.authentication.URI("oauth", "token"), data, &t, opts...)

	return
}

// VerifyWithOOB verifies an MFA challenge using an out-of-band challenge (OOB), either push notification,
// SMS, or voice.
//
// See: https://auth0.com/docs/api/authentication#verify-with-out-of-band-oob-
func (m *MFA) VerifyWithOOB(ctx context.Context, body mfa.VerifyWithOOBRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	missing := []string{}
	check(&missing, "ClientID", (body.ClientID != "" || m.authentication.clientID != ""))
	check(&missing, "MFAToken", body.MFAToken != "")
	check(&missing, "OOBCode", body.OOBCode != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("Missing required fields: %s", strings.Join(missing, ", "))
	}

	data := url.Values{
		"mfa_token":  []string{body.MFAToken},
		"grant_type": []string{"http://auth0.com/oauth/grant-type/mfa-oob"},
		"oob_code":   []string{body.OOBCode},
	}

	if body.BindingCode != "" {
		data.Set("binding_code", body.BindingCode)
	}

	err = m.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, true)

	if err != nil {
		return nil, err
	}

	err = m.authentication.Request(ctx, "POST", m.authentication.URI("oauth", "token"), data, &t, opts...)

	return
}

// VerifyWithRecoveryCode verifies an MFA challenge using a recovery code.
//
// See: https://auth0.com/docs/api/authentication#verify-with-recovery-code
func (m *MFA) VerifyWithRecoveryCode(ctx context.Context, body mfa.VerifyWithRecoveryCodeRequest, opts ...RequestOption) (t *mfa.VerifyWithRecoveryCodeResponse, err error) {
	missing := []string{}
	check(&missing, "ClientID", (body.ClientID != "" || m.authentication.clientID != ""))
	check(&missing, "MFAToken", body.MFAToken != "")
	check(&missing, "RecoveryCode", body.RecoveryCode != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("Missing required fields: %s", strings.Join(missing, ", "))
	}

	data := url.Values{
		"mfa_token":     []string{body.MFAToken},
		"grant_type":    []string{"http://auth0.com/oauth/grant-type/mfa-recovery-code"},
		"recovery_code": []string{body.RecoveryCode},
	}

	err = m.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, true)

	if err != nil {
		return nil, err
	}

	err = m.authentication.Request(ctx, "POST", m.authentication.URI("oauth", "token"), data, &t, opts...)

	return
}
