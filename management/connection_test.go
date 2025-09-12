package management

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

var connectionTestCases = []connectionTestCase{
	{
		name: "Auth0 Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			StrategyVersion: auth0.Int(2),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Wordpress Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Wordpress-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("wordpress"),
		},
		options: &ConnectionOptionsOAuth2{
			StrategyVersion: auth0.Int(2),
			Scope:           auth0.String("email profile openid"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
			CustomHeaders: &map[string]string{
				"X-Auth0-Client":        "my-client",
				"X-Auth0-Client-Secret": "my-secret",
			},
		},
	},
	{
		name: "GoogleOAuth2 Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-GoogleOAuth2-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("google-oauth2"),
		},
		options: &ConnectionOptionsGoogleOAuth2{
			AllowedAudiences: &[]string{
				"example.com",
				"api.example.com",
			},
			Profile:  auth0.Bool(true),
			Calendar: auth0.Bool(true),
			Youtube:  auth0.Bool(false),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "GoogleApps Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-GoogleApps-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("google-apps"),
		},
		options: &ConnectionOptionsGoogleApps{
			Domain:          auth0.String("example.com"),
			TenantDomain:    auth0.String("example.com"),
			BasicProfile:    auth0.Bool(true),
			ExtendedProfile: auth0.Bool(true),
			Groups:          auth0.Bool(true),
			Admin:           auth0.Bool(true),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Email Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Email-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("email"),
		},
		options: &ConnectionOptionsEmail{
			Email: &ConnectionOptionsEmailSettings{
				Syntax:  auth0.String("liquid"),
				From:    auth0.String("{{application.name}} <test@example.com>"),
				Subject: auth0.String("Email Login - {{application.name}}"),
				Body:    auth0.String("<html><body>email contents</body></html>"),
			},
			OTP: &ConnectionOptionsOTP{
				TimeStep: auth0.Int(100),
				Length:   auth0.Int(4),
			},
			AuthParams: map[string]string{
				"scope": "openid profile",
			},
			BruteForceProtection: auth0.Bool(true),
			DisableSignup:        auth0.Bool(true),
			Name:                 auth0.String("Test-Connection-Email"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "SMS Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-SMS-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("sms"),
		},
		options: &ConnectionOptionsSMS{
			From:     auth0.String("+17777777777"),
			Template: auth0.String("Your verification code is { code }}"),
			Syntax:   auth0.String("liquid"),
			OTP: &ConnectionOptionsOTP{
				TimeStep: auth0.Int(110),
				Length:   auth0.Int(5),
			},
			AuthParams: map[string]string{
				"scope": "openid profile",
			},
			BruteForceProtection: auth0.Bool(true),
			DisableSignup:        auth0.Bool(true),
			Name:                 auth0.String("Test-Connection-SMS"),
			TwilioSID:            auth0.String("abc132asdfasdf56"),
			TwilioToken:          auth0.String("234127asdfsada23"),
			MessagingServiceSID:  auth0.String("273248090982390423"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Custom SMS Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Custom-SMS-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("sms"),
		},
		options: &ConnectionOptionsSMS{
			From:     auth0.String("+17777777777"),
			Template: auth0.String("Your verification code is { code }}"),
			Syntax:   auth0.String("liquid"),
			OTP: &ConnectionOptionsOTP{
				TimeStep: auth0.Int(110),
				Length:   auth0.Int(5),
			},
			BruteForceProtection: auth0.Bool(true),
			DisableSignup:        auth0.Bool(true),
			Name:                 auth0.String("Test-Connection-Custom-SMS"),
			Provider:             auth0.String("sms_gateway"),
			GatewayURL:           auth0.String("https://test.com/sms-gateway"),
			GatewayAuthentication: &ConnectionGatewayAuthentication{
				Method:              auth0.String("bearer"),
				Subject:             auth0.String("test.us.auth0.com:sms"),
				Audience:            auth0.String("test.com/sms-gateway"),
				Secret:              auth0.String("my-secret"),
				SecretBase64Encoded: auth0.Bool(false),
			},
			ForwardRequestInfo: auth0.Bool(true),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "SAML Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-SAML-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("samlp"),
		},
		options: &ConnectionOptionsSAML{
			GlobalTokenRevocationJWTSub: auth0.String("user123"),
			GlobalTokenRevocationJWTIss: auth0.String("issuer.example.com"),
			StrategyVersion:             auth0.Int(2),
			SignInEndpoint:              auth0.String("https://saml.identity/provider"),
			SigningCert: auth0.String(`-----BEGIN CERTIFICATE-----
MIID6TCCA1ICAQEwDQYJKoZIhvcNAQEFBQAwgYsxCzAJBgNVBAYTAlVTMRMwEQYD
VQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQK
EwtHb29nbGUgSW5jLjEMMAoGA1UECxMDRW5nMQwwCgYDVQQDEwNhZ2wxHTAbBgkq
hkiG9w0BCQEWDmFnbEBnb29nbGUuY29tMB4XDTA5MDkwOTIyMDU0M1oXDTEwMDkw
OTIyMDU0M1owajELMAkGA1UEBhMCQVUxEzARBgNVBAgTClNvbWUtU3RhdGUxITAf
BgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEjMCEGA1UEAxMaZXVyb3Bh
LnNmby5jb3JwLmdvb2dsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQC6pgYt7/EibBDumASF+S0qvqdL/f+nouJw2T1Qc8GmXF/iiUcrsgzh/Fd8
pDhz/T96Qg9IyR4ztuc2MXrmPra+zAuSf5bevFReSqvpIt8Duv0HbDbcqs/XKPfB
uMDe+of7a9GCywvAZ4ZUJcp0thqD9fKTTjUWOBzHY1uNE4RitrhmJCrbBGXbJ249
bvgmb7jgdInH2PU7PT55hujvOoIsQW2osXBFRur4pF1wmVh4W4lTLD6pjfIMUcML
ICHEXEN73PDic8KS3EtNYCwoIld+tpIBjE1QOb1KOyuJBNW6Esw9ALZn7stWdYcE
qAwvv20egN2tEXqj7Q4/1ccyPZc3PQgC3FJ8Be2mtllM+80qf4dAaQ/fWvCtOrQ5
pnfe9juQvCo8Y0VGlFcrSys/MzSg9LJ/24jZVgzQved/Qupsp89wVidwIzjt+WdS
fyWfH0/v1aQLvu5cMYuW//C0W2nlYziL5blETntM8My2ybNARy3ICHxCBv2RNtPI
WQVm+E9/W5rwh2IJR4DHn2LHwUVmT/hHNTdBLl5Uhwr4Wc7JhE7AVqb14pVNz1lr
5jxsp//ncIwftb7mZQ3DF03Yna+jJhpzx8CQoeLT6aQCHyzmH68MrHHT4MALPyUs
Pomjn71GNTtDeWAXibjCgdL6iHACCF6Htbl0zGlG0OAK+bdn0QIDAQABMA0GCSqG
SIb3DQEBBQUAA4GBAOKnQDtqBV24vVqvesL5dnmyFpFPXBn3WdFfwD6DzEb21UVG
5krmJiu+ViipORJPGMkgoL6BjU21XI95VQbun5P8vvg8Z+FnFsvRFY3e1CCzAVQY
ZsUkLw2I7zI/dNlWdB8Xp7v+3w9sX5N3J/WuJ1KOO5m26kRlHQo7EzT3974g
-----END CERTIFICATE-----`),
			TenantDomain: auth0.String("example.com"),
			FieldsMap: map[string]interface{}{
				"email":       "EmailAddress",
				"given_name":  "FirstName",
				"family_name": "LastName",
			},
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
			DecryptionKey: &ConnectionOptionsSAMLDecryptionKey{
				Key:  auth0.String(`-----BEGIN PRIVATE KEY-----\n...{your private key here}...\n-----END PRIVATE KEY-----`),
				Cert: auth0.String(`-----BEGIN CERTIFICATE-----\n...{your public key cert here}...\n-----END CERTIFICATE-----`),
			},
		},
	},
	{
		name: "Azure-AD Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-AzureAD-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("waad"),
		},
		options: &ConnectionOptionsAzureAD{
			StrategyVersion: auth0.Int(2),
			Domain:          auth0.String("example.onmicrosoft.com"),
			TenantDomain:    auth0.String("example.onmicrosoft.com"),
			ClientID:        auth0.String("123456"),
			ClientSecret:    auth0.String("123456"),
			UserIDAttribute: auth0.String("sub"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "AD Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-AD-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("ad"),
		},
		options: &ConnectionOptionsAD{
			StrategyVersion: auth0.Int(2),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "ADFS Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-ADFS-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("adfs"),
		},
		options: &ConnectionOptionsADFS{
			StrategyVersion: auth0.Int(2),
			FedMetadataXML: auth0.String(`<?xml version="1.0" encoding="utf-8"?>
<EntityDescriptor entityID="https://example.com"
                  xmlns="urn:oasis:names:tc:SAML:2.0:metadata">
    <RoleDescriptor xsi:type="fed:ApplicationServiceType"
                    protocolSupportEnumeration="http://docs.oasis-open.org/wsfed/federation/200706"
                    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                    xmlns:fed="http://docs.oasis-open.org/wsfed/federation/200706">
        <fed:TargetScopes>
            <wsa:EndpointReference xmlns:wsa="http://www.w3.org/2005/08/addressing">
                <wsa:Address>https://adfs.provider/</wsa:Address>
            </wsa:EndpointReference>
        </fed:TargetScopes>
        <fed:ApplicationServiceEndpoint>
            <wsa:EndpointReference xmlns:wsa="http://www.w3.org/2005/08/addressing">
                <wsa:Address>https://adfs.provider/wsfed</wsa:Address>
            </wsa:EndpointReference>
        </fed:ApplicationServiceEndpoint>
        <fed:PassiveRequestorEndpoint>
            <wsa:EndpointReference xmlns:wsa="http://www.w3.org/2005/08/addressing">
                <wsa:Address>https://adfs.provider/wsfed</wsa:Address>
            </wsa:EndpointReference>
        </fed:PassiveRequestorEndpoint>
    </RoleDescriptor>
    <IDPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
        <SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
                             Location="https://adfs.provider/sign_out"/>
        <SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
                             Location="https://adfs.provider/sign_in"/>
    </IDPSSODescriptor>
</EntityDescriptor>
`),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Facebook Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Facebook-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("facebook"),
		},
		options: &ConnectionOptionsFacebook{
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Apple Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Apple-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("apple"),
		},
		options: &ConnectionOptionsApple{
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "LinkedIn Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-LinkedIn-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("linkedin"),
		},
		options: &ConnectionOptionsLinkedin{
			StrategyVersion: auth0.Int(2),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "GitHub Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-GitHub-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("github"),
		},
		options: &ConnectionOptionsGitHub{
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "WindowsLive Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-WindowsLive-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("windowslive"),
		},
		options: &ConnectionOptionsWindowsLive{
			StrategyVersion: auth0.Int(2),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "SalesForce Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-SalesForce-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("salesforce"),
		},
		options: &ConnectionOptionsSalesforce{
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "OIDC Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-OIDC-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("oidc"),
		},
		options: &ConnectionOptionsOIDC{
			ClientID:              auth0.String("4ef8d976-71bd-4473-a7ce-087d3f0fafd8"),
			Scope:                 auth0.String("openid profile email"),
			Issuer:                auth0.String("https://example.com"),
			AuthorizationEndpoint: auth0.String("https://example.com/authorize"),
			JWKSURI:               auth0.String("https://example.com/.well-known/jwks.json"),
			Type:                  auth0.String("front_channel"),
			DiscoveryURL:          auth0.String("https://example.com/.well-known/openid-configuration"),
			OIDCMetadata: map[string]interface{}{
				"issuer":                                "https://example.com",
				"authorization_endpoint":                "https://example.com/authorize",
				"token_endpoint":                        "https://example.com/oauth/token",
				"userinfo_endpoint":                     "https://example.com/userinfo",
				"jwks_uri":                              "https://example.com/.well-known/jwks.json",
				"response_types_supported":              []string{"code", "id_token", "token", "code id_token"},
				"subject_types_supported":               []string{"public"},
				"id_token_signing_alg_values_supported": []string{"RS256"},
				"token_endpoint_auth_methods_supported": []string{"client_secret_basic", "client_secret_post"},
				"claims_supported":                      []string{"sub", "iss", "aud", "exp", "iat", "name", "email"},
				"scopes_supported":                      []string{"openid", "profile", "email"},
				"grant_types_supported":                 []string{"authorization_code", "implicit", "refresh_token"},
			},
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Okta Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Okta-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("okta"),
		},
		options: &ConnectionOptionsOkta{
			ClientID:              auth0.String("4ef8d976-71bd-4473-a7ce-087d3f0fafd8"),
			ClientSecret:          auth0.String("mySecret"),
			Scope:                 auth0.String("openid"),
			Domain:                auth0.String("domain.okta.com"),
			Issuer:                auth0.String("https://example.com"),
			AuthorizationEndpoint: auth0.String("https://example.com"),
			JWKSURI:               auth0.String("https://example.com/jwks"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	},
	{
		name: "Ping Federate Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-Ping-Federate-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("pingfederate"),
		},
		options: &ConnectionOptionsPingFederate{
			PingFederateBaseURL: auth0.String("https://ping.example.com"),
			SigningCert: auth0.String(`-----BEGIN CERTIFICATE-----
MIID6TCCA1ICAQEwDQYJKoZIhvcNAQEFBQAwgYsxCzAJBgNVBAYTAlVTMRMwEQYD
VQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQK
EwtHb29nbGUgSW5jLjEMMAoGA1UECxMDRW5nMQwwCgYDVQQDEwNhZ2wxHTAbBgkq
hkiG9w0BCQEWDmFnbEBnb29nbGUuY29tMB4XDTA5MDkwOTIyMDU0M1oXDTEwMDkw
OTIyMDU0M1owajELMAkGA1UEBhMCQVUxEzARBgNVBAgTClNvbWUtU3RhdGUxITAf
BgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEjMCEGA1UEAxMaZXVyb3Bh
LnNmby5jb3JwLmdvb2dsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQC6pgYt7/EibBDumASF+S0qvqdL/f+nouJw2T1Qc8GmXF/iiUcrsgzh/Fd8
pDhz/T96Qg9IyR4ztuc2MXrmPra+zAuSf5bevFReSqvpIt8Duv0HbDbcqs/XKPfB
uMDe+of7a9GCywvAZ4ZUJcp0thqD9fKTTjUWOBzHY1uNE4RitrhmJCrbBGXbJ249
bvgmb7jgdInH2PU7PT55hujvOoIsQW2osXBFRur4pF1wmVh4W4lTLD6pjfIMUcML
ICHEXEN73PDic8KS3EtNYCwoIld+tpIBjE1QOb1KOyuJBNW6Esw9ALZn7stWdYcE
qAwvv20egN2tEXqj7Q4/1ccyPZc3PQgC3FJ8Be2mtllM+80qf4dAaQ/fWvCtOrQ5
pnfe9juQvCo8Y0VGlFcrSys/MzSg9LJ/24jZVgzQved/Qupsp89wVidwIzjt+WdS
fyWfH0/v1aQLvu5cMYuW//C0W2nlYziL5blETntM8My2ybNARy3ICHxCBv2RNtPI
WQVm+E9/W5rwh2IJR4DHn2LHwUVmT/hHNTdBLl5Uhwr4Wc7JhE7AVqb14pVNz1lr
5jxsp//ncIwftb7mZQ3DF03Yna+jJhpzx8CQoeLT6aQCHyzmH68MrHHT4MALPyUs
Pomjn71GNTtDeWAXibjCgdL6iHACCF6Htbl0zGlG0OAK+bdn0QIDAQABMA0GCSqG
SIb3DQEBBQUAA4GBAOKnQDtqBV24vVqvesL5dnmyFpFPXBn3WdFfwD6DzEb21UVG
5krmJiu+ViipORJPGMkgoL6BjU21XI95VQbun5P8vvg8Z+FnFsvRFY3e1CCzAVQY
ZsUkLw2I7zI/dNlWdB8Xp7v+3w9sX5N3J/WuJ1KOO5m26kRlHQo7EzT3974g
-----END CERTIFICATE-----`),
			SignSAMLRequest:    auth0.Bool(true),
			SignatureAlgorithm: auth0.String("rsa-sha256"),
			DigestAlgorithm:    auth0.String("sha256"),
		},
	},
}

type connectionTestCase struct {
	name       string
	connection Connection
	options    interface{}
}

func TestConnectionManager_Create(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully create a "+testCase.name, func(t *testing.T) {
			configureHTTPTestRecordings(t)

			expectedConnection := testCase.connection
			expectedConnection.Options = testCase.options

			err := api.Connection.Create(context.Background(), &expectedConnection)

			assert.NoError(t, err)
			assert.NotEmpty(t, expectedConnection.GetID())
			assert.IsType(t, testCase.options, expectedConnection.Options)

			t.Cleanup(func() {
				cleanupConnection(t, expectedConnection.GetID())
			})
		})
	}
}

var Auth0ConnectionTestCase = []connectionTestCase{
	{
		name: "Auth0 Connection With RequireUsername",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-RequireUsername-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence:       &[]string{"username", "email", "phone_number"},
			RequiresUsername: auth0.Bool(true),
		},
	},
	{
		name: "Auth0 Connection with PhoneNumber as Identifier",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Phone-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				PhoneNumber: &ConnectionOptionsPhoneNumberAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection with Email as Identifier",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Email-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				Email: &ConnectionOptionsEmailAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection with Email as Identifier With VerificationMethod OTP",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Email-VerificationMethod-OTP%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				Email: &ConnectionOptionsEmailAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					VerificationMethod: &ConnectionOptionsEmailAttributeVerificationMethodOtp,
					ProfileRequired:    auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection with Email as Identifier With VerificationMethod Link",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Email-VerificationMethod-Link%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				Email: &ConnectionOptionsEmailAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					VerificationMethod: &ConnectionOptionsEmailAttributeVerificationMethodLink,
					ProfileRequired:    auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection with Username as Identifier",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Username-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				Username: &ConnectionOptionsUsernameAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
					},
					Validation: &ConnectionOptionsAttributeValidation{
						MinLength: auth0.Int(1),
						MaxLength: auth0.Int(5),
						AllowedTypes: &ConnectionOptionsAttributeAllowedTypes{
							Email:       auth0.Bool(true),
							PhoneNumber: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection Cannot set both requires_username and attributes together",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Invalid-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence:       &[]string{"username", "email", "phone_number"},
			RequiresUsername: auth0.Bool(true),
			Attributes: &ConnectionOptionsAttributes{
				Email: &ConnectionOptionsEmailAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection With No attribute is active",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-No-Active-Attributes-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				PhoneNumber: &ConnectionOptionsPhoneNumberAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(false),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
				Email: &ConnectionOptionsEmailAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(false),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
				Username: &ConnectionOptionsUsernameAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(false),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
					},
				},
			},
		},
	},
	{
		name: "Auth0 Connection Cannot set both validation and attributes together",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Attributes-And-Validation-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				Email: &ConnectionOptionsEmailAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
				Username: &ConnectionOptionsUsernameAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("required"),
					},
				},
			},
			Validation: map[string]interface{}{
				"username": map[string]interface{}{
					"min": 1,
					"max": 5,
				},
			},
		},
	},
	{
		name: "Auth0 Connection Attribute required in profile but inactive on signup",
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-Attribute-Inactive-On-Signup-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
		options: &ConnectionOptions{
			Precedence: &[]string{"username", "email", "phone_number"},
			Attributes: &ConnectionOptionsAttributes{
				PhoneNumber: &ConnectionOptionsPhoneNumberAttribute{
					Identifier: &ConnectionOptionsAttributeIdentifier{
						Active: auth0.Bool(true),
					},
					ProfileRequired: auth0.Bool(true),
					Signup: &ConnectionOptionsAttributeSignup{
						Status: auth0.String("inactive"),
						Verification: &ConnectionOptionsAttributeVerification{
							Active: auth0.Bool(false),
						},
					},
				},
			},
		},
	},
}

func TestConnectionManager_CreateDBConnectionWithDifferentOptions(t *testing.T) {
	for _, testCase := range Auth0ConnectionTestCase {
		t.Run("It handles "+testCase.name, func(t *testing.T) {
			configureHTTPTestRecordings(t)

			expectedConnection := testCase.connection
			expectedConnection.Options = testCase.options

			err := api.Connection.Create(context.Background(), &expectedConnection)

			switch testCase.name {
			case "Auth0 Connection Cannot set both requires_username and attributes together":
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "Cannot set both options.attributes and options.requires_username")
			case "Auth0 Connection With No attribute is active":
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "attributes must contain one active identifier")
			case "Auth0 Connection Cannot set both validation and attributes together":
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "Cannot set both options.attributes and options.validation")
			case "Auth0 Connection Attribute required in profile but inactive on signup":
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "attribute phone_number must also be required on signup")
			default:
				assert.NoError(t, err)
				assert.NotEmpty(t, expectedConnection.Name)
				assert.NotEmpty(t, expectedConnection.Strategy)
			}

			if err == nil {
				t.Cleanup(func() {
					cleanupConnection(t, expectedConnection.GetID())
				})
			}
		})
	}
}

func TestConnectionManager_Read(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully read a "+testCase.name, func(t *testing.T) {
			configureHTTPTestRecordings(t)

			expectedConnection := givenAConnection(t, testCase)

			actualConnection, err := api.Connection.Read(context.Background(), expectedConnection.GetID())

			assert.NoError(t, err)
			assert.Equal(t, expectedConnection.GetID(), actualConnection.GetID())
			assert.Equal(t, expectedConnection.GetName(), actualConnection.GetName())
			assert.Equal(t, expectedConnection.GetStrategy(), actualConnection.GetStrategy())
			assert.IsType(t, testCase.options, actualConnection.Options)

			switch testCase.connection.GetStrategy() {
			case "ad", "adfs", "auth0", "samlp", "waad", "windowslive", "wordpress":
				assert.ObjectsAreEqualValues(getStrategyVersion(testCase.connection.GetStrategy(), testCase.options), getStrategyVersion(actualConnection.GetStrategy(), actualConnection.Options))
			}

			t.Cleanup(func() {
				cleanupConnection(t, expectedConnection.GetID())
			})
		})
	}
}

func TestConnectionManager_ReadByName(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully find a "+testCase.name+" by its name", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			expectedConnection := givenAConnection(t, testCase)

			actualConnection, err := api.Connection.ReadByName(context.Background(), expectedConnection.GetName())

			assert.NoError(t, err)
			assert.Equal(t, expectedConnection.GetID(), actualConnection.GetID())
			assert.Equal(t, expectedConnection.GetName(), actualConnection.GetName())
			assert.Equal(t, expectedConnection.GetStrategy(), actualConnection.GetStrategy())
			assert.IsType(t, testCase.options, actualConnection.Options)

			t.Cleanup(func() {
				cleanupConnection(t, expectedConnection.GetID())
			})
		})
	}

	t.Run("throw an error when connection name is empty", func(t *testing.T) {
		actualConnection, err := api.Connection.ReadByName(context.Background(), "")

		assert.EqualError(t, err, "400 Bad Request: Name cannot be empty")
		assert.Empty(t, actualConnection)
	})
}

func TestConnectionManager_EnabledClients(t *testing.T) {
	configureHTTPTestRecordings(t)

	client := givenAClient(t)

	connection := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-%d", time.Now().UnixNano()),
			Strategy: auth0.String("auth0"),
		},
	})

	t.Run("add client to connection", func(t *testing.T) {
		payload := []ConnectionEnabledClient{
			{
				ClientID: client.ClientID,
				Status:   auth0.Bool(true),
			},
		}

		err := api.Connection.UpdateEnabledClients(context.Background(), connection.GetID(), payload)
		assert.NoError(t, err)

		clientIDs := getEnabledClientIDs(t, connection.GetID())
		assert.Contains(t, clientIDs, client.GetClientID(), "client should be enabled")
	})

	t.Run("remove client from connection", func(t *testing.T) {
		payload := []ConnectionEnabledClient{
			{
				ClientID: client.ClientID,
				Status:   auth0.Bool(false),
			},
		}

		err := api.Connection.UpdateEnabledClients(context.Background(), connection.GetID(), payload)
		assert.NoError(t, err)

		clientIDs := getEnabledClientIDs(t, connection.GetID())
		assert.NotContains(t, clientIDs, client.GetClientID(), "client should be removed")
	})
}

func TestConnectionManager_Update(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully update a "+testCase.name, func(t *testing.T) {
			if testCase.connection.GetStrategy() == "oidc" ||
				testCase.connection.GetStrategy() == "samlp" ||
				testCase.connection.GetStrategy() == "okta" ||
				testCase.connection.GetStrategy() == "adfs" ||
				testCase.connection.GetStrategy() == "waad" ||
				testCase.connection.GetStrategy() == "pingfederate" {
				t.Skip("Skipping because we can't create an oidc, okta, samlp, adfs, waad, or pingfederate connection with no options")
			}

			configureHTTPTestRecordings(t)

			connection := givenAConnection(t, connectionTestCase{connection: testCase.connection})

			connectionWithUpdatedOptions := &Connection{
				Options: testCase.options,
			}

			err := api.Connection.Update(context.Background(), connection.GetID(), connectionWithUpdatedOptions)
			assert.NoError(t, err)

			actualConnection, err := api.Connection.Read(context.Background(), connection.GetID())
			assert.NoError(t, err)
			assert.ObjectsAreEqualValues(testCase.options, actualConnection.Options)
		})
	}
}

func TestConnectionManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
	})

	err := api.Connection.Delete(context.Background(), expectedConnection.GetID())
	assert.NoError(t, err)

	actualConnection, err := api.Connection.Read(context.Background(), expectedConnection.GetID())
	assert.Nil(t, actualConnection)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestConnectionManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-List-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
	})

	needle := &Connection{
		ID:                 expectedConnection.ID,
		IsDomainConnection: auth0.Bool(false),
	}
	connectionList, err := api.Connection.List(context.Background(), IncludeFields("id"))
	assert.NoError(t, err)
	assert.Contains(t, connectionList.Connections, needle)
}

func TestConnectionOptionsScopes(t *testing.T) {
	t.Run("It can successfully set the scopes on the options of a OIDC connection", func(t *testing.T) {
		options := &ConnectionOptionsOIDC{}

		options.SetScopes(true, "foo", "bar", "baz")
		assert.Equal(t, []string{"bar", "baz", "foo"}, options.Scopes())

		options.SetScopes(false, "foo", "baz")
		assert.Equal(t, []string{"bar"}, options.Scopes())
	})

	t.Run("It can successfully set the scopes on the options of an OAuth2 connection", func(t *testing.T) {
		options := &ConnectionOptionsOAuth2{}

		options.SetScopes(true, "foo", "bar", "baz")
		assert.Equal(t, []string{"bar", "baz", "foo"}, options.Scopes())

		options.SetScopes(false, "foo", "baz")
		assert.Equal(t, []string{"bar"}, options.Scopes())
	})

	t.Run("It can successfully set the scopes on the options of an Okta connection", func(t *testing.T) {
		options := &ConnectionOptionsOkta{}

		options.SetScopes(true, "foo", "bar", "baz")
		assert.Equal(t, []string{"bar", "baz", "foo"}, options.Scopes())

		options.SetScopes(false, "foo", "baz")
		assert.Equal(t, []string{"bar"}, options.Scopes())
	})
}

func TestConnectionManager_CreateSCIMConfiguration(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)
	expectedSCIMConfig := &SCIMConfiguration{
		Mapping: &[]SCIMConfigurationMapping{
			{SCIM: auth0.String("userName"), Auth0: auth0.String("username")},
			{SCIM: auth0.String("email"), Auth0: auth0.String("email")},
		},
		UserIDAttribute: auth0.String("userName"),
	}
	err := api.Connection.CreateSCIMConfiguration(context.Background(), expectedConnection.GetID(), expectedSCIMConfig)
	assert.NoError(t, err)

	actualSCIMConfiguration, err := api.Connection.ReadSCIMConfiguration(context.Background(), expectedConnection.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedConnection.GetID(), actualSCIMConfiguration.GetConnectionID())
	assert.IsType(t, &SCIMConfiguration{}, actualSCIMConfiguration)
	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_CreateSCIMConfigurationWithoutBody(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)
	expectedSCIMConfig := &SCIMConfiguration{}
	err := api.Connection.CreateSCIMConfiguration(context.Background(), expectedConnection.GetID(), expectedSCIMConfig)
	assert.NoError(t, err)

	actualSCIMConfiguration, err := api.Connection.ReadSCIMConfiguration(context.Background(), expectedConnection.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedConnection.GetID(), actualSCIMConfiguration.GetConnectionID())
	assert.IsType(t, &SCIMConfiguration{}, actualSCIMConfiguration)
	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_UpdateSCIMConfiguration(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)
	expectedSCIMConfig := givenASCIMConfiguration(t, expectedConnection.GetID())
	assert.Equal(t, expectedConnection.GetID(), expectedSCIMConfig.GetConnectionID())
	expectedSCIMConfig = &SCIMConfiguration{
		Mapping: &[]SCIMConfigurationMapping{
			{SCIM: auth0.String("userName"), Auth0: auth0.String("username")},
			{SCIM: auth0.String("email"), Auth0: auth0.String("email")},
		},
		UserIDAttribute: auth0.String("userName"),
	}

	err := api.Connection.UpdateSCIMConfiguration(context.Background(), expectedConnection.GetID(), expectedSCIMConfig)
	assert.NoError(t, err)

	actualSCIMConfiguration, err := api.Connection.ReadSCIMConfiguration(context.Background(), expectedConnection.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedSCIMConfig, actualSCIMConfiguration)
	assert.Equal(t, expectedConnection.GetID(), actualSCIMConfiguration.GetConnectionID())
	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_DeleteSCIMConfiguration(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)

	expectedSCIMConfiguration := givenASCIMConfiguration(t, expectedConnection.GetID())

	err := api.Connection.DeleteSCIMConfiguration(context.Background(), expectedSCIMConfiguration.GetConnectionID())
	assert.NoError(t, err)

	actualSCIMConfiguration, err := api.Connection.ReadSCIMConfiguration(context.Background(), expectedSCIMConfiguration.GetConnectionID())
	assert.Nil(t, actualSCIMConfiguration)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestConnectionManager_ReadSCIMConfiguration(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)

	expectedSCIMConfig := &SCIMConfiguration{
		Mapping: &[]SCIMConfigurationMapping{
			{SCIM: auth0.String("userName"), Auth0: auth0.String("username")},
			{SCIM: auth0.String("email"), Auth0: auth0.String("email")},
		},
		UserIDAttribute: auth0.String("userName"),
	}
	err := api.Connection.CreateSCIMConfiguration(context.Background(), expectedConnection.GetID(), expectedSCIMConfig)
	assert.NoError(t, err)

	actualSCIMConfiguration, err := api.Connection.ReadSCIMConfiguration(context.Background(), expectedSCIMConfig.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, expectedConnection.GetID(), actualSCIMConfiguration.GetConnectionID())
	assert.Equal(t, expectedSCIMConfig, actualSCIMConfiguration)

	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_ReadSCIMDefaultConfiguration(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)

	expectedSCIMConfig := &SCIMConfiguration{}
	err := api.Connection.CreateSCIMConfiguration(context.Background(), expectedConnection.GetID(), expectedSCIMConfig)
	assert.NoError(t, err)

	actualSCIMConfiguration, err := api.Connection.ReadSCIMDefaultConfiguration(context.Background(), expectedSCIMConfig.GetConnectionID())
	assert.NoError(t, err)
	assert.Equal(t, expectedSCIMConfig.GetMapping(), actualSCIMConfiguration.GetMapping())

	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_CreateSCIMToken(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)
	expectedSCIMConfig := givenASCIMConfiguration(t, expectedConnection.GetID())

	SCIMTokenPayload := &SCIMToken{
		Scopes: &[]string{"get:users", "post:users", "put:users", "patch:users"},
	}

	err := api.Connection.CreateSCIMToken(context.Background(), expectedSCIMConfig.GetConnectionID(), SCIMTokenPayload)
	assert.NoError(t, err)

	assert.NotEmpty(t, SCIMTokenPayload.GetToken())

	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_ListSCIMTokens(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)

	expectedSCIMConfig := givenASCIMConfiguration(t, expectedConnection.GetID())

	SCIMTokenPayload := &SCIMToken{
		Scopes: &[]string{"get:users", "post:users", "put:users", "patch:users"},
	}

	err := api.Connection.CreateSCIMToken(context.Background(), expectedSCIMConfig.GetConnectionID(), SCIMTokenPayload)
	assert.NoError(t, err)

	SCIMTokenPayload.Token = nil

	actualSCIMTokens, err := api.Connection.ListSCIMToken(context.Background(), expectedConnection.GetID())
	assert.NoError(t, err)

	assert.Contains(t, actualSCIMTokens, SCIMTokenPayload)

	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_DeleteSCIMToken(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedConnection := givenAOktaConnection(t)
	expectedSCIMConfig := givenASCIMConfiguration(t, expectedConnection.GetID())

	expectedSCIMToken := &SCIMToken{
		Scopes: &[]string{"get:users", "post:users", "put:users", "patch:users"},
	}

	err := api.Connection.CreateSCIMToken(context.Background(), expectedSCIMConfig.GetConnectionID(), expectedSCIMToken)
	assert.NoError(t, err)

	expectedSCIMToken.Token = nil

	actualSCIMTokens, err := api.Connection.ListSCIMToken(context.Background(), expectedSCIMConfig.GetConnectionID())
	assert.NoError(t, err)

	assert.Contains(t, actualSCIMTokens, expectedSCIMToken)

	err = api.Connection.DeleteSCIMToken(context.Background(), expectedSCIMConfig.GetConnectionID(), expectedSCIMToken.GetTokenID())
	assert.NoError(t, err)

	actualSCIMTokens, err = api.Connection.ListSCIMToken(context.Background(), expectedSCIMConfig.GetConnectionID())
	assert.NoError(t, err)
	assert.Empty(t, actualSCIMTokens)

	t.Cleanup(func() {
		cleanupSCIMConfig(t, expectedConnection.GetID())
	})
}

func TestConnectionManager_ReadKeys(t *testing.T) {
	configureHTTPTestRecordings(t)

	tests := []struct {
		name     string
		strategy string
		options  interface{}
	}{
		{
			name:     "OIDC Connection",
			strategy: "oidc",
			options: &ConnectionOptionsOIDC{
				ClientID:              auth0.String("4ef8d976-71bd-4473-a7ce-087d3f0fafd8"),
				Scope:                 auth0.String("openid"),
				Issuer:                auth0.String("https://example.com"),
				AuthorizationEndpoint: auth0.String("https://example.com"),
				JWKSURI:               auth0.String("https://example.com/jwks"),
				Type:                  auth0.String("front_channel"),
				DiscoveryURL:          auth0.String("https://www.paypalobjects.com/.well-known/openid-configuration"),
				UpstreamParams: map[string]interface{}{
					"screen_name": map[string]interface{}{
						"alias": "login_hint",
					},
				},
				TokenEndpointAuthMethod:     auth0.String("private_key_jwt"),
				TokenEndpointAuthSigningAlg: auth0.String("RS256"),
			},
		},
		{
			name:     "Okta Connection",
			strategy: "okta",
			options: &ConnectionOptionsOkta{
				ClientID:              auth0.String("4ef8d976-71bd-4473-a7ce-087d3f0fafd8"),
				ClientSecret:          auth0.String("mySecret"),
				Scope:                 auth0.String("openid"),
				Domain:                auth0.String("domain.okta.com"),
				Issuer:                auth0.String("https://example.com"),
				AuthorizationEndpoint: auth0.String("https://example.com"),
				JWKSURI:               auth0.String("https://example.com/jwks"),
				UpstreamParams: map[string]interface{}{
					"screen_name": map[string]interface{}{
						"alias": "login_hint",
					},
				},
				TokenEndpointAuthMethod:     auth0.String("private_key_jwt"),
				TokenEndpointAuthSigningAlg: auth0.String("RS256"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			conn := givenAConnection(t, connectionTestCase{
				connection: Connection{
					Name:     auth0.Stringf("Test-%s-Connection-%d", strings.ToUpper(tc.strategy), time.Now().Unix()),
					Strategy: auth0.String(tc.strategy),
				},
				options: tc.options,
			})

			keys, err := api.Connection.ReadKeys(context.Background(), conn.GetID())
			assert.NoError(t, err)
			assert.NotEmpty(t, keys)
		})
	}
}

func TestConnectionManager_RotateKeys(t *testing.T) {
	configureHTTPTestRecordings(t)

	conn := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:     auth0.Stringf("Test-Connection-Rotate-Keys-%d", time.Now().Unix()),
			Strategy: auth0.String("oidc"),
		},
		options: &ConnectionOptionsOIDC{
			ClientID:              auth0.String("4ef8d976-71bd-4473-a7ce-087d3f0fafd8"),
			Scope:                 auth0.String("openid"),
			Issuer:                auth0.String("https://example.com"),
			AuthorizationEndpoint: auth0.String("https://example.com"),
			JWKSURI:               auth0.String("https://example.com/jwks"),
			Type:                  auth0.String("front_channel"),
			DiscoveryURL:          auth0.String("https://www.paypalobjects.com/.well-known/openid-configuration"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
			TokenEndpointAuthMethod:     auth0.String("private_key_jwt"),
			TokenEndpointAuthSigningAlg: auth0.String("RS256"),
		},
	})

	ctx := context.Background()

	beforeKeys, err := api.Connection.ReadKeys(ctx, conn.GetID())
	assert.NoError(t, err)
	assert.NotEmpty(t, beforeKeys)

	rotatedKey, err := api.Connection.RotateKeys(ctx, conn.GetID())
	assert.NoError(t, err)
	assert.NotNil(t, rotatedKey)

	afterKeys, err := api.Connection.ReadKeys(ctx, conn.GetID())
	assert.NoError(t, err)
	assert.NotEmpty(t, afterKeys)

	foundRotated := false

	for _, k := range afterKeys {
		if k.GetKID() == rotatedKey.GetKID() {
			foundRotated = true
			break
		}
	}

	assert.True(t, foundRotated, "Rotated key should be present in afterKeys")

	assert.NotEqual(t, len(beforeKeys), len(afterKeys), "Keys should be different after rotation")

	t.Cleanup(func() {
		cleanupConnection(t, conn.GetID())
	})
}

func TestConnectionOptionsUsernameAttribute_MarshalJSON(t *testing.T) {
	for attribute, expected := range map[*ConnectionOptionsUsernameAttribute]string{
		{
			Identifier: &ConnectionOptionsAttributeIdentifier{
				Active: auth0.Bool(true),
			},
			ProfileRequired: auth0.Bool(true),
			Signup: &ConnectionOptionsAttributeSignup{
				Status:       auth0.String("required"),
				Verification: &ConnectionOptionsAttributeVerification{Active: auth0.Bool(false)},
			},
		}: `{"identifier":{"active":true},"profile_required":true,"signup":{"status":"required"}}`,
		{
			ProfileRequired: auth0.Bool(true),
			Signup: &ConnectionOptionsAttributeSignup{
				Status: auth0.String("required"),
			},
		}: `{"profile_required":true,"signup":{"status":"required"}}`,

		{
			Identifier: &ConnectionOptionsAttributeIdentifier{
				Active: auth0.Bool(true),
			},
			ProfileRequired: auth0.Bool(true),
		}: `{"identifier":{"active":true},"profile_required":true}`,
	} {
		payload, err := json.Marshal(attribute)
		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(payload))
	}
}

func TestOAuth2Connection_MarshalJSON(t *testing.T) {
	for connection, expected := range map[*ConnectionOptionsOAuth2]string{
		{Scope: auth0.String("foo bar baz")}: `{"authorizationURL":null,"tokenURL":null,"scope":["foo","bar","baz"]}`,
		{Scope: auth0.String("")}:            `{"authorizationURL":null,"tokenURL":null,"scope":[]}`,
		{Scope: nil}:                         `{"authorizationURL":null,"tokenURL":null}`,
	} {
		payload, err := json.Marshal(connection)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func TestOAuth2Connection_UnmarshalJSON(t *testing.T) {
	for expectedAsString, expected := range map[string]*ConnectionOptionsOAuth2{
		`{"scope":["foo","bar","baz"]}`: {Scope: auth0.String("foo bar baz")},
		`{"scope":null}`:                {Scope: nil},
		`{}`:                            {},
		`{"scope":[]}`:                  {Scope: auth0.String("")},
	} {
		var actual *ConnectionOptionsOAuth2
		err := json.Unmarshal([]byte(expectedAsString), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestGoogleOauth2Connection_MarshalJSON(t *testing.T) {
	var emptySlice []string
	for connection, expected := range map[*ConnectionOptionsGoogleOAuth2]string{
		{AllowedAudiences: nil}:                                              `{}`,
		{AllowedAudiences: &emptySlice}:                                      `{"allowed_audiences":null}`,
		{AllowedAudiences: &[]string{}}:                                      `{"allowed_audiences":[]}`,
		{AllowedAudiences: &[]string{"foo", "bar"}}:                          `{"allowed_audiences":["foo","bar"]}`,
		{AllowedAudiences: &[]string{"foo", "bar"}, Email: auth0.Bool(true)}: `{"email":true,"allowed_audiences":["foo","bar"]}`,
	} {
		payload, err := json.Marshal(connection)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func TestGoogleOauth2Connection_UnmarshalJSON(t *testing.T) {
	for expectedAsString, expected := range map[string]*ConnectionOptionsGoogleOAuth2{
		`{}`:                          {},
		`{"allowed_audiences": null}`: {},
		`{"allowed_audiences": ""}`:   {AllowedAudiences: &[]string{}},
		`{"allowed_audiences": []}`:   {AllowedAudiences: &[]string{}},
		`{"allowed_audiences": ["foo", "bar"], "scope": ["email"] }`: {AllowedAudiences: &[]string{"foo", "bar"}, Scope: []interface{}{"email"}},
	} {
		var actual *ConnectionOptionsGoogleOAuth2
		err := json.Unmarshal([]byte(expectedAsString), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}

	t.Run("Throws an unexpected type error", func(t *testing.T) {
		var actual *ConnectionOptionsGoogleOAuth2
		err := json.Unmarshal([]byte(`{"allowed_audiences": 1}`), &actual)
		assert.EqualError(t, err, "unexpected type for field allowed_audiences: float64")
	})
}

func cleanupConnection(t *testing.T, connectionID string) {
	t.Helper()

	err := api.Connection.Delete(context.Background(), connectionID)
	require.NoError(t, err)
}

func cleanupSCIMConfig(t *testing.T, connectionID string) {
	t.Helper()

	err := api.Connection.DeleteSCIMConfiguration(context.Background(), connectionID)
	require.NoError(t, err)
}

func givenAConnection(t *testing.T, testCase connectionTestCase) *Connection {
	t.Helper()

	connection := testCase.connection
	connection.Options = testCase.options

	err := api.Connection.Create(context.Background(), &connection)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupConnection(t, connection.GetID())
	})

	return &connection
}

func givenASCIMConfiguration(t *testing.T, connectionID string) *SCIMConfiguration {
	t.Helper()

	expectedSCIMConfig := &SCIMConfiguration{}

	err := api.Connection.CreateSCIMConfiguration(context.Background(), connectionID, expectedSCIMConfig)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupSCIMConfig(t, connectionID)
	})

	return expectedSCIMConfig
}

func givenAOktaConnection(t *testing.T) *Connection {
	t.Helper()

	return givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:     auth0.Stringf("Test-Okta-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("okta"),
		},
		options: &ConnectionOptionsOkta{
			ClientID:              auth0.String("4ef8d976-71bd-4473-a7ce-087d3f0fafd8"),
			ClientSecret:          auth0.String("mySecret"),
			Scope:                 auth0.String("openid"),
			Domain:                auth0.String("domain.okta.com"),
			Issuer:                auth0.String("https://example.com"),
			AuthorizationEndpoint: auth0.String("https://example.com"),
			JWKSURI:               auth0.String("https://example.com/jwks"),
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
			},
		},
	})
}

func getStrategyVersion(strategy string, options interface{}) int {
	switch strategy {
	case "ad":
		return options.(*ConnectionOptionsAD).GetStrategyVersion()
	case "adfs":
		return options.(*ConnectionOptionsADFS).GetStrategyVersion()
	case "auth0":
		return options.(*ConnectionOptions).GetStrategyVersion()
	case "samlp":
		return options.(*ConnectionOptionsSAML).GetStrategyVersion()
	case "waad":
		return options.(*ConnectionOptionsAzureAD).GetStrategyVersion()
	case "windowslive":
		return options.(*ConnectionOptionsWindowsLive).GetStrategyVersion()
	case "wordpress":
		return options.(*ConnectionOptionsOAuth2).GetStrategyVersion()
	default:
		return -1
	}
}

func getEnabledClientIDs(t *testing.T, connectionID string) []string {
	t.Helper()

	resp, err := api.Connection.ReadEnabledClients(context.Background(), connectionID)
	assert.NoError(t, err, "failed to read enabled clients")
	assert.NotNil(t, resp.GetClients(), "clients list should not be nil")

	var clientIDs []string
	for _, c := range resp.GetClients() {
		clientIDs = append(clientIDs, c.GetClientID())
	}

	return clientIDs
}
