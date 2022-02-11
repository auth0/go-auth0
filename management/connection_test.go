package management

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/internal/testing/expect"
)

func TestConnection(t *testing.T) {
	c := &Connection{
		Name:     auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
		Strategy: auth0.String("auth0"),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Connection.Create(c)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := c.Options.(*ConnectionOptions); !ok {
			t.Errorf("unexpected options type %T", c.Options)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.Connection.Read(c.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("List", func(t *testing.T) {
		cs, err := m.Connection.List()
		if err != nil {
			t.Error(err)
		}
		for _, c := range cs.Connections {
			var ok bool

			switch c.GetStrategy() {
			case ConnectionStrategyAuth0:
				_, ok = c.Options.(*ConnectionOptions)
			case ConnectionStrategyGoogleOAuth2:
				_, ok = c.Options.(*ConnectionOptionsGoogleOAuth2)
			case ConnectionStrategyFacebook:
				_, ok = c.Options.(*ConnectionOptionsFacebook)
			case ConnectionStrategyApple:
				_, ok = c.Options.(*ConnectionOptionsApple)
			case ConnectionStrategyLinkedin:
				_, ok = c.Options.(*ConnectionOptionsLinkedin)
			case ConnectionStrategyGitHub:
				_, ok = c.Options.(*ConnectionOptionsGitHub)
			case ConnectionStrategyWindowsLive:
				_, ok = c.Options.(*ConnectionOptionsWindowsLive)
			case ConnectionStrategySalesforce, ConnectionStrategySalesforceCommunity, ConnectionStrategySalesforceSandbox:
				_, ok = c.Options.(*ConnectionOptionsSalesforce)
			case ConnectionStrategyEmail:
				_, ok = c.Options.(*ConnectionOptionsEmail)
			case ConnectionStrategySMS:
				_, ok = c.Options.(*ConnectionOptionsSMS)
			case ConnectionStrategyOIDC:
				_, ok = c.Options.(*ConnectionOptionsOIDC)
			case ConnectionStrategyOAuth2:
				_, ok = c.Options.(*ConnectionOptionsOAuth2)
			case ConnectionStrategyAD:
				_, ok = c.Options.(*ConnectionOptionsAD)
			case ConnectionStrategyAzureAD:
				_, ok = c.Options.(*ConnectionOptionsAzureAD)
			case ConnectionStrategySAML:
				_, ok = c.Options.(*ConnectionOptionsSAML)
			case ConnectionStrategyGoogleApps:
				_, ok = c.Options.(*ConnectionOptionsGoogleApps)
			default:
				_, ok = c.Options.(map[string]interface{})
			}

			if !ok {
				t.Errorf("unexpected options type %T", c.Options)
			}

			t.Logf("%s %s %T\n", c.GetID(), c.GetName(), c.Options)
		}
	})

	t.Run("Update", func(t *testing.T) {
		id := c.GetID()

		c.ID = nil       // read-only
		c.Name = nil     // read-only
		c.Strategy = nil // read-only

		c.Options = &ConnectionOptions{

			BruteForceProtection: auth0.Bool(true),
			ImportMode:           auth0.Bool(false), // try some zero values
			DisableSignup:        auth0.Bool(true),
			RequiresUsername:     auth0.Bool(false),

			CustomScripts:      map[string]interface{}{"get_user": "function( { return callback(null) }"},
			Configuration:      map[string]interface{}{"foo": "bar"},
			NonPersistentAttrs: &[]string{"ethnicity", "gender"},
		}

		err = m.Connection.Update(id, c)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Connection.Delete(c.GetID())
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ReadByName", func(t *testing.T) {
		cs, err := m.Connection.ReadByName("Username-Password-Authentication")
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", cs)
	})

	t.Run("ReadByNameEmptyName", func(t *testing.T) {
		cs, err := m.Connection.ReadByName("")
		if err == nil {
			t.Fail()
		}
		mgmtError, ok := err.(*managementError)
		if !ok {
			t.Fail()
		}
		if mgmtError.StatusCode != 400 {
			t.Fail()
		}
		if cs != nil {
			t.Fail()
		}
		t.Logf("%v\n", cs)
	})
}

func TestConnectionOptions(t *testing.T) {
	t.Run("GoogleOAuth2", func(t *testing.T) {
		g := &Connection{
			Name:     auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("google-oauth2"),
			Options: &ConnectionOptionsGoogleOAuth2{
				AllowedAudiences: []interface{}{
					"example.com",
					"api.example.com",
				},
				Profile:  auth0.Bool(true),
				Calendar: auth0.Bool(true),
				Youtube:  auth0.Bool(false),
			},
		}

		defer func() {
			m.Connection.Delete(g.GetID())
			assertDeleted(t, g)
		}()

		err := m.Connection.Create(g)
		if err != nil {
			t.Fatal(err)
		}
		c, err := m.Connection.Read(g.GetID())
		if err != nil {
			t.Fatal(err)
		}
		o, ok := c.Options.(*ConnectionOptionsGoogleOAuth2)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, o.GetProfile(), true)
		expect.Expect(t, o.GetCalendar(), true)
		expect.Expect(t, o.GetYoutube(), false)
		expect.Expect(t, o.Scopes(), []string{"email", "profile", "calendar"})

		o.NonPersistentAttrs = &[]string{"gender", "ethnicity", "favorite_color"}
		err = m.Connection.Update(g.GetID(), &Connection{
			Options: o,
		})
		if err != nil {
			t.Fatal(err)
		}
		c, _ = m.Connection.Read(g.GetID())
		o, ok = c.Options.(*ConnectionOptionsGoogleOAuth2)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{"gender", "ethnicity", "favorite_color"})

		o.NonPersistentAttrs = &[]string{""}
		err = m.Connection.Update(g.GetID(), &Connection{
			Options: o,
		})
		if err != nil {
			log.Fatal(err)
		}

		c, err = m.Connection.Read(g.GetID())
		if err != nil {
			log.Fatal(err)
		}
		o, ok = c.Options.(*ConnectionOptionsGoogleOAuth2)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{""})

		t.Logf("%s\n", g)
	})

	t.Run("GoogleApps", func(t *testing.T) {
		g := &Connection{
			Name:     auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("google-apps"),
			Options: &ConnectionOptionsGoogleApps{
				Domain:          auth0.String("example.com"),
				TenantDomain:    auth0.String("example.com"),
				BasicProfile:    auth0.Bool(true),
				ExtendedProfile: auth0.Bool(true),
				Groups:          auth0.Bool(true),
				Admin:           auth0.Bool(true),
			},
		}

		defer func() {
			m.Connection.Delete(g.GetID())
			assertDeleted(t, g)
		}()

		err := m.Connection.Create(g)
		if err != nil {
			t.Fatal(err)
		}
		c, err := m.Connection.Read(g.GetID())
		if err != nil {
			t.Fatal(err)
		}
		o, ok := c.Options.(*ConnectionOptionsGoogleApps)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, o.GetBasicProfile(), true)
		expect.Expect(t, o.GetExtendedProfile(), true)
		expect.Expect(t, o.GetGroups(), true)
		expect.Expect(t, o.GetAdmin(), true)
		expect.Expect(t, o.Scopes(), []string{"basic_profile", "ext_profile", "ext_groups", "ext_is_admin"})

		o.NonPersistentAttrs = &[]string{"gender", "ethnicity", "favorite_color"}
		err = m.Connection.Update(g.GetID(), &Connection{
			Options: o,
		})
		if err != nil {
			t.Fatal(err)
		}
		c, _ = m.Connection.Read(g.GetID())
		o, ok = c.Options.(*ConnectionOptionsGoogleApps)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{"gender", "ethnicity", "favorite_color"})

		o.NonPersistentAttrs = &[]string{""}
		err = m.Connection.Update(g.GetID(), &Connection{
			Options: o,
		})
		if err != nil {
			log.Fatal(err)
		}

		c, err = m.Connection.Read(g.GetID())
		if err != nil {
			log.Fatal(err)
		}
		o, ok = c.Options.(*ConnectionOptionsGoogleApps)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{""})

		t.Logf("%s\n", g)
	})

	t.Run("OIDC", func(t *testing.T) {
		o := &ConnectionOptionsOIDC{}
		expect.Expect(t, len(o.Scopes()), 0)

		o.SetScopes(true, "foo", "bar", "baz")
		expect.Expect(t, len(o.Scopes()), 3)
		expect.Expect(t, o.Scopes(), []string{"bar", "baz", "foo"})

		o.SetScopes(false, "baz")
		expect.Expect(t, len(o.Scopes()), 2)
		expect.Expect(t, o.Scopes(), []string{"bar", "foo"})

		o.NonPersistentAttrs = &[]string{"gender", "ethnicity", "favorite_color"}
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{"gender", "ethnicity", "favorite_color"})
	})

	t.Run("OAuth2", func(t *testing.T) {
		o := &ConnectionOptionsOAuth2{
			Scripts:            map[string]interface{}{"fetchUserProfile": "function( { return callback(null) }"},
			TokenURL:           auth0.String("https://example.com/oauth2/token"),
			AuthorizationURL:   auth0.String("https://example.com/oauth2/authorize"),
			ClientID:           auth0.String("test-client"),
			ClientSecret:       auth0.String("superSecretKey"),
			NonPersistentAttrs: &[]string{"gender", "ethnicity", "favorite_color"},
		}
		expect.Expect(t, len(o.Scopes()), 0)

		o.SetScopes(true, "foo", "bar", "baz")
		expect.Expect(t, len(o.Scopes()), 3)
		expect.Expect(t, o.Scopes(), []string{"bar", "baz", "foo"})

		o.SetScopes(false, "baz")
		expect.Expect(t, len(o.Scopes()), 2)
		expect.Expect(t, o.Scopes(), []string{"bar", "foo"})
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{"gender", "ethnicity", "favorite_color"})

		o.NonPersistentAttrs = &[]string{"foo", "bar"}
		expect.Expect(t, o.GetNonPersistentAttrs(), []string{"foo", "bar"})
	})

	t.Run("Email", func(t *testing.T) {
		e := &Connection{
			Name:     auth0.Stringf("Test-Connection-Email-%d", time.Now().Unix()),
			Strategy: auth0.String("email"),
			Options: &ConnectionOptionsEmail{
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
			},
		}

		defer func() {
			m.Connection.Delete(e.GetID())
			assertDeleted(t, e)
		}()

		err := m.Connection.Create(e)
		if err != nil {
			t.Fatal(err)
		}
		o, ok := e.Options.(*ConnectionOptionsEmail)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetEmail().GetSyntax(), "liquid")
		expect.Expect(t, o.GetEmail().GetFrom(), "{{application.name}} <test@example.com>")
		expect.Expect(t, o.GetEmail().GetSubject(), "Email Login - {{application.name}}")
		expect.Expect(t, o.GetEmail().GetBody(), "<html><body>email contents</body></html>")
		expect.Expect(t, o.GetOTP().GetTimeStep(), 100)
		expect.Expect(t, o.GetOTP().GetLength(), 4)
		expect.Expect(t, o.AuthParams["scope"], "openid profile")
		expect.Expect(t, o.GetBruteForceProtection(), true)
		expect.Expect(t, o.GetDisableSignup(), true)
		expect.Expect(t, o.GetName(), "Test-Connection-Email")

		o.NonPersistentAttrs = &[]string{"gender", "ethnicity", "favorite_color"}
		err = m.Connection.Update(e.GetID(), &Connection{
			Options: o,
		})
		if err != nil {
			t.Fatal(err)
		}
		e, _ = m.Connection.Read(e.GetID())
		o, ok = e.Options.(*ConnectionOptionsEmail)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetNonPersistentAttrs(), []string{"gender", "ethnicity", "favorite_color"})
		t.Logf("%s\n", e)
	})

	t.Run("SMS", func(t *testing.T) {
		s := &Connection{
			Name:     auth0.Stringf("Test-Connection-SMS-%d", time.Now().Unix()),
			Strategy: auth0.String("sms"),
			Options: &ConnectionOptionsSMS{
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
			},
		}

		defer func() {
			m.Connection.Delete(s.GetID())
			assertDeleted(t, s)
		}()

		err := m.Connection.Create(s)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := s.Options.(*ConnectionOptionsSMS)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetTemplate(), "Your verification code is { code }}")
		expect.Expect(t, o.GetFrom(), "+17777777777")
		expect.Expect(t, o.GetSyntax(), "liquid")
		expect.Expect(t, o.GetOTP().GetTimeStep(), 110)
		expect.Expect(t, o.GetOTP().GetLength(), 5)
		expect.Expect(t, o.AuthParams["scope"], "openid profile")
		expect.Expect(t, o.GetBruteForceProtection(), true)
		expect.Expect(t, o.GetDisableSignup(), true)
		expect.Expect(t, o.GetName(), "Test-Connection-SMS")
		expect.Expect(t, o.GetTwilioSID(), "abc132asdfasdf56")
		expect.Expect(t, o.GetTwilioToken(), "234127asdfsada23")
		expect.Expect(t, o.GetMessagingServiceSID(), "273248090982390423")

		t.Logf("%s\n", s)
	})

	t.Run("CustomSMS", func(t *testing.T) {
		s := &Connection{
			Name:     auth0.Stringf("Test-Connection-Custom-SMS-%d", time.Now().Unix()),
			Strategy: auth0.String("sms"),
			Options: &ConnectionOptionsSMS{
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
			},
		}

		defer func() {
			m.Connection.Delete(s.GetID())
			assertDeleted(t, s)
		}()

		err := m.Connection.Create(s)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := s.Options.(*ConnectionOptionsSMS)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetTemplate(), "Your verification code is { code }}")
		expect.Expect(t, o.GetFrom(), "+17777777777")
		expect.Expect(t, o.GetSyntax(), "liquid")
		expect.Expect(t, o.GetOTP().GetTimeStep(), 110)
		expect.Expect(t, o.GetOTP().GetLength(), 5)
		expect.Expect(t, o.GetBruteForceProtection(), true)
		expect.Expect(t, o.GetDisableSignup(), true)
		expect.Expect(t, o.GetName(), "Test-Connection-Custom-SMS")
		expect.Expect(t, o.GetProvider(), "sms_gateway")
		expect.Expect(t, o.GetGatewayURL(), "https://test.com/sms-gateway")
		expect.Expect(t, o.GetGatewayAuthentication().GetMethod(), "bearer")
		expect.Expect(t, o.GetGatewayAuthentication().GetSubject(), "test.us.auth0.com:sms")
		expect.Expect(t, o.GetGatewayAuthentication().GetAudience(), "test.com/sms-gateway")
		expect.Expect(t, o.GetGatewayAuthentication().GetSecret(), "my-secret")
		expect.Expect(t, o.GetGatewayAuthentication().GetSecretBase64Encoded(), false)
		expect.Expect(t, o.GetForwardRequestInfo(), true)

		t.Logf("%s\n", s)
	})

	t.Run("SAML", func(t *testing.T) {
		g := &Connection{
			Name:     auth0.Stringf("Test-SAML-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("samlp"),
			Options: &ConnectionOptionsSAML{
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
			},
		}
		defer func() {
			m.Connection.Delete(g.GetID())
			assertDeleted(t, g)
		}()

		err := m.Connection.Create(g)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := g.Options.(*ConnectionOptionsSAML)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetSignInEndpoint(), "https://saml.identity/provider")
		expect.Expect(t, o.GetTenantDomain(), "example.com")
	})

	t.Run("AD", func(t *testing.T) {
		a := &Connection{
			Name:     auth0.Stringf("Test-Connection-%d", time.Now().Unix()),
			Strategy: auth0.String("ad"),
		}

		defer func() {
			m.Connection.Delete(a.GetID())
			assertDeleted(t, a)
		}()

		if err := m.Connection.Create(a); err != nil {
			t.Fatal(err)
		}

		if a.ProvisioningTicketURL == nil {
			t.Fatal("provisioning_ticket_url should be returned")
		}

		o, ok := a.Options.(*ConnectionOptionsAD)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}

		expect.Expect(t, o.GetCertAuth(), false)
		expect.Expect(t, o.GetKerberos(), false)
		expect.Expect(t, o.GetDisableCache(), false)

		o.DisableCache = auth0.Bool(true)
		if err := m.Connection.Update(
			a.GetID(),
			&Connection{
				Options: o,
			},
		); err != nil {
			t.Fatal(err)
		}

		o, _ = a.Options.(*ConnectionOptionsAD)
		expect.Expect(t, o.GetDisableCache(), true)

		// failed to update with changing provisioning_ticket_url
		err := m.Connection.Update(
			a.GetID(),
			&Connection{
				ProvisioningTicketURL: auth0.String("https://invalid-domain.com"),
				Options:               o,
			},
		)
		if err == nil {
			t.Fatal("err should be returned")
		}
	})
}

func assertDeleted(t *testing.T, c *Connection) {
	c, err := m.Connection.Read(c.GetID())
	assert.Nil(t, c)
	assert.EqualError(t, err, "404 Not Found: The connection does not exist")
}
