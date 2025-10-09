package authentication

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/auth0/go-auth0/v2/authentication/mfa"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
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
		//nolint ST1005 Keeping message unchanged for backward compatibility
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
		//nolint ST1005 Keeping message unchanged for backward compatibility
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
		//nolint ST1005 Keeping message unchanged for backward compatibility
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
		//nolint ST1005 Keeping message unchanged for backward compatibility
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

// AddAuthenticator Associates or adds a new authenticator for multi-factor authentication (MFA).
//
// See: https://auth0.com/docs/api/authentication#add-an-authenticator
func (m *MFA) AddAuthenticator(ctx context.Context, accessOrMfaToken string, body mfa.AddAuthenticatorRequest, opts ...RequestOption) (a *mfa.AddAuthenticatorResponse, err error) {
	opts = append(opts, Header("Authorization", "Bearer "+accessOrMfaToken))
	missing := []string{}
	check(&missing, "ClientID", (body.ClientID != "" || m.authentication.clientID != ""))
	check(&missing, "AuthenticatorTypes", len(body.AuthenticatorTypes) > 0)

	if len(missing) > 0 {
		//nolint ST1005 Keeping message unchanged for backward compatibility
		return nil, fmt.Errorf("Missing required fields: %s", strings.Join(missing, ", "))
	}

	if err = m.authentication.addClientAuthenticationToClientAuthStruct(&body.ClientAuthentication, true); err != nil {
		return
	}

	err = m.authentication.Request(ctx, "POST", m.authentication.URI("mfa", "associate"), body, &a, opts...)

	return
}

// ListAuthenticators Returns a list of authenticators associated with your application.
//
// See: https://auth0.com/docs/api/authentication#list-authenticators
func (m *MFA) ListAuthenticators(ctx context.Context, accessOrMfaToken string, opts ...RequestOption) (a []mfa.ListAuthenticatorsResponse, err error) {
	opts = append(opts, Header("Authorization", "Bearer "+accessOrMfaToken))
	err = m.authentication.Request(ctx, "GET", m.authentication.URI("mfa", "authenticators"), nil, &a, opts...)

	return
}

// DeleteAuthenticator Deletes an associated authenticator using its ID.
//
// See: https://auth0.com/docs/api/authentication#delete-an-authenticator
func (m *MFA) DeleteAuthenticator(ctx context.Context, accessToken string, authenticatorID string, opts ...RequestOption) (err error) {
	opts = append(opts, Header("Authorization", "Bearer "+accessToken))
	err = m.authentication.Request(ctx, "DELETE", m.authentication.URI("mfa", "authenticators", authenticatorID), nil, nil, opts...)

	return
}
