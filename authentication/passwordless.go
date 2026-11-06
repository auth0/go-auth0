package authentication

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/auth0/go-auth0/v3/authentication/oauth"
	"github.com/auth0/go-auth0/v3/authentication/passwordless"
	"github.com/auth0/go-auth0/v3/internal/idtokenvalidator"
)

const passwordlessOTPGrantType = "http://auth0.com/oauth/grant-type/passwordless/otp" //nolint:gosec // Grant type identifier, not a credential.

// e164PhoneNumberRegexp loosely matches E.164 phone numbers.
var e164PhoneNumberRegexp = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

// Passwordless exposes logging in using the passwordless APIs.
type Passwordless manager

// SendEmail starts a passwordless flow by sending a link or code via email.
//
// In order to set the `x-request-language` header when sending this request, use the `Header` RequestOption
// helper.
//
// See: https://auth0.com/docs/api/authentication?http#get-code-or-link
func (p *Passwordless) SendEmail(ctx context.Context, params passwordless.SendEmailRequest, opts ...RequestOption) (r *passwordless.SendEmailResponse, err error) {
	err = p.authentication.addClientAuthenticationToClientAuthStruct(&params.ClientAuthentication, false)
	if err != nil {
		return nil, err
	}

	params.Connection = "email"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("passwordless", "start"), params, &r, opts...)

	return
}

// LoginWithEmail completes the passwordless flow started in `SendEmail` by exchanging the code for a token.
//
// See: https://auth0.com/docs/api/authentication?http#authenticate-user
func (p *Passwordless) LoginWithEmail(ctx context.Context, params passwordless.LoginWithEmailRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	err = p.authentication.addClientAuthenticationToClientAuthStruct(&params.ClientAuthentication, false)
	if err != nil {
		return nil, err
	}

	params.GrantType = "http://auth0.com/oauth/grant-type/passwordless/otp"
	params.Realm = "email"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("oauth", "token"), params, &t, opts...)

	if t != nil && t.IDToken != "" {
		err = p.authentication.idTokenValidator.Validate(t.IDToken, idtokenvalidator.ValidationOptions{
			MaxAge:       validationOptions.MaxAge,
			Nonce:        validationOptions.Nonce,
			Organization: validationOptions.Organization,
		})
		if err != nil {
			return nil, err
		}
	}

	return
}

// SendSMS starts a passwordless flow by sending a code via SMS.
//
// In order to set the `x-request-language` header when sending this request, use the `Header` RequestOption
// helper.
//
// See: https://auth0.com/docs/api/authentication?http#get-code-or-link
func (p *Passwordless) SendSMS(ctx context.Context, params passwordless.SendSMSRequest, opts ...RequestOption) (r *passwordless.SendSMSResponse, err error) {
	err = p.authentication.addClientAuthenticationToClientAuthStruct(&params.ClientAuthentication, false)
	if err != nil {
		return nil, err
	}

	params.Connection = "sms"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("passwordless", "start"), params, &r, opts...)

	return
}

// LoginWithSMS completes the passwordless flow started in `SendSMS` by exchanging the code for a token.
//
// See: https://auth0.com/docs/api/authentication?http#authenticate-user
func (p *Passwordless) LoginWithSMS(ctx context.Context, params passwordless.LoginWithSMSRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	err = p.authentication.addClientAuthenticationToClientAuthStruct(&params.ClientAuthentication, false)
	if err != nil {
		return nil, err
	}

	params.GrantType = "http://auth0.com/oauth/grant-type/passwordless/otp"
	params.Realm = "sms"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("oauth", "token"), params, &t, opts...)

	if t != nil && t.IDToken != "" {
		err = p.authentication.idTokenValidator.Validate(t.IDToken, idtokenvalidator.ValidationOptions{
			MaxAge:       validationOptions.MaxAge,
			Nonce:        validationOptions.Nonce,
			Organization: validationOptions.Organization,
		})
		if err != nil {
			return nil, err
		}
	}

	return
}

// ChallengeWithEmail starts a passwordless OTP flow on a database connection by sending a code to the user's email.
//
// This is the first step of a two-step flow. Pass the returned AuthSession, which is opaque and must not be
// parsed, together with the code the user received to `LoginWithOTPChallenge`.
//
// To prevent user enumeration, the challenge succeeds even if the user does not exist.
func (p *Passwordless) ChallengeWithEmail(ctx context.Context, params passwordless.ChallengeWithEmailRequest, opts ...RequestOption) (r *passwordless.ChallengeResponse, err error) {
	err = p.authentication.addClientAuthenticationToClientAuthStruct(&params.ClientAuthentication, false)
	if err != nil {
		return nil, err
	}

	missing := []string{}
	check(&missing, "ClientID", params.ClientID != "")
	check(&missing, "Connection", params.Connection != "")
	check(&missing, "Email", params.Email != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("otp", "challenge"), params, &r, opts...)

	return
}

// ChallengeWithPhoneNumber starts a passwordless OTP flow on a database connection by sending a code to the
// user's phone number by text or voice.
//
// This is the first step of a two-step flow. Pass the returned AuthSession, which is opaque and must not be
// parsed, together with the code the user received to `LoginWithOTPChallenge`.
//
// To prevent user enumeration, the challenge succeeds even if the user does not exist.
func (p *Passwordless) ChallengeWithPhoneNumber(ctx context.Context, params passwordless.ChallengeWithPhoneNumberRequest, opts ...RequestOption) (r *passwordless.ChallengeResponse, err error) {
	err = p.authentication.addClientAuthenticationToClientAuthStruct(&params.ClientAuthentication, false)
	if err != nil {
		return nil, err
	}

	missing := []string{}
	check(&missing, "ClientID", params.ClientID != "")
	check(&missing, "Connection", params.Connection != "")
	check(&missing, "PhoneNumber", params.PhoneNumber != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	if !e164PhoneNumberRegexp.MatchString(params.PhoneNumber) {
		return nil, errors.New("PhoneNumber must be in E.164 format (e.g. +14155550100)")
	}

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("otp", "challenge"), params, &r, opts...)

	return
}

// LoginWithOTPChallenge completes the passwordless OTP flow started by `ChallengeWithEmail` or
// `ChallengeWithPhoneNumber` by exchanging the auth session and the code for a token.
//
// If multi-factor authentication is required, the returned error is an `*authentication.Error`
// whose `GetMFAToken()` is set.
func (p *Passwordless) LoginWithOTPChallenge(ctx context.Context, params passwordless.LoginWithOTPChallengeRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	missing := []string{}
	check(&missing, "AuthSession", params.AuthSession != "")
	check(&missing, "OTP", params.OTP != "")
	check(&missing, "ClientID", params.ClientID != "" || p.authentication.clientID != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	data := url.Values{
		"auth_session": []string{params.AuthSession},
		"otp":          []string{params.OTP},
	}

	addIfNotEmpty("scope", params.Scope, data)
	addIfNotEmpty("audience", params.Audience, data)

	for k, v := range params.ExtraParameters {
		data.Set(k, v)
	}

	err = p.authentication.addClientAuthenticationToURLValues(params.ClientAuthentication, data, false)
	if err != nil {
		return nil, err
	}

	return (*OAuth)(p).LoginWithGrant(ctx, passwordlessOTPGrantType, data, validationOptions, opts...)
}
