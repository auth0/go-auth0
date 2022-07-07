package management

import (
	"net/http"
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
			UpstreamParams: map[string]interface{}{
				"screen_name": map[string]interface{}{
					"alias": "login_hint",
				},
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
			AllowedAudiences: []interface{}{
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
			SignInEndpoint: auth0.String("https://saml.identity/provider"),
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
		},
	},
	{
		name: "AD Connection",
		connection: Connection{
			Name:     auth0.Stringf("Test-AD-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("ad"),
		},
		options: &ConnectionOptionsAD{
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
			Scope:                 auth0.String("openid"),
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
}

type connectionTestCase struct {
	name       string
	connection Connection
	options    interface{}
}

func TestConnectionManager_Create(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully create a "+testCase.name, func(t *testing.T) {
			setupHTTPRecordings(t)

			expectedConnection := testCase.connection
			expectedConnection.Options = testCase.options

			err := m.Connection.Create(&expectedConnection)

			assert.NoError(t, err)
			assert.NotEmpty(t, expectedConnection.GetID())
			assert.IsType(t, testCase.options, expectedConnection.Options)

			t.Cleanup(func() {
				cleanupConnection(t, expectedConnection.GetID())
			})
		})
	}
}

func TestConnectionManager_Read(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully read a "+testCase.name, func(t *testing.T) {
			setupHTTPRecordings(t)

			expectedConnection := givenAConnection(t, testCase)

			actualConnection, err := m.Connection.Read(expectedConnection.GetID())

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
}

func TestConnectionManager_ReadByName(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully find a "+testCase.name+" by its name", func(t *testing.T) {
			setupHTTPRecordings(t)

			expectedConnection := givenAConnection(t, testCase)

			actualConnection, err := m.Connection.ReadByName(expectedConnection.GetName())

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
		setupHTTPRecordings(t)

		actualConnection, err := m.Connection.ReadByName("")

		assert.EqualError(t, err, "400 Bad Request: Name cannot be empty")
		assert.Empty(t, actualConnection)
	})
}

func TestConnectionManager_Update(t *testing.T) {
	for _, testCase := range connectionTestCases {
		t.Run("It can successfully update a "+testCase.name, func(t *testing.T) {
			if testCase.connection.GetStrategy() == "oidc" || testCase.connection.GetStrategy() == "samlp" {
				t.Skip("Skipping because we can't create an oidc or samlp connection with no options")
			}

			setupHTTPRecordings(t)

			connection := givenAConnection(t, connectionTestCase{connection: testCase.connection})

			connectionWithUpdatedOptions := &Connection{
				Options: testCase.options,
			}

			err := m.Connection.Update(connection.GetID(), connectionWithUpdatedOptions)
			assert.NoError(t, err)

			actualConnection, err := m.Connection.Read(connection.GetID())
			assert.NoError(t, err)
			assert.ObjectsAreEqualValues(testCase.options, actualConnection.Options)
		})
	}
}

func TestConnectionManager_Delete(t *testing.T) {
	setupHTTPRecordings(t)

	expectedConnection := givenAConnection(t, connectionTestCase{
		connection: Connection{
			Name:     auth0.Stringf("Test-Auth0-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("auth0"),
		},
	})

	err := m.Connection.Delete(expectedConnection.GetID())
	assert.NoError(t, err)

	actualConnection, err := m.Connection.Read(expectedConnection.GetID())
	assert.Nil(t, actualConnection)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestConnectionManager_List(t *testing.T) {
	setupHTTPRecordings(t)

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
	connectionList, err := m.Connection.List(IncludeFields("id"))
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

	t.Run("It can successfully set the scopes on the options of a OAuth2 connection", func(t *testing.T) {
		options := &ConnectionOptionsOAuth2{}

		options.SetScopes(true, "foo", "bar", "baz")
		assert.Equal(t, []string{"bar", "baz", "foo"}, options.Scopes())

		options.SetScopes(false, "foo", "baz")
		assert.Equal(t, []string{"bar"}, options.Scopes())
	})
}

func cleanupConnection(t *testing.T, connectionID string) {
	t.Helper()

	err := m.Connection.Delete(connectionID)
	require.NoError(t, err)
}

func givenAConnection(t *testing.T, testCase connectionTestCase) *Connection {
	t.Helper()

	connection := testCase.connection
	connection.Options = testCase.options

	err := m.Connection.Create(&connection)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupConnection(t, connection.GetID())
	})

	return &connection
}
