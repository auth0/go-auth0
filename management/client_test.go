package management

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestClient_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := &Client{
		Name:        auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
	}

	err := api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())

	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestClient_CreateWithClientRefreshToken(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Create a Resource Server
	resourceServer := &ResourceServer{
		Name:       auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier: auth0.String("https://mrrt"),
		Scopes: &[]ResourceServerScope{
			{
				Description: auth0.String("This is just a test client."),
				Value:       auth0.String("create:bar"),
			},
			{
				Description: auth0.String("This is just a test client."),
				Value:       auth0.String("read:bar"),
			},
		},
		SkipConsentForVerifiableFirstPartyClients: auth0.Bool(true),
		AllowOfflineAccess:                        auth0.Bool(true),
	}

	err := api.ResourceServer.Create(context.Background(), resourceServer)
	assert.NoError(t, err)
	assert.NotEmpty(t, resourceServer.GetID())
	t.Cleanup(func() {
		cleanupResourceServer(t, resourceServer.GetID())
	})

	// Create a Client with Refresh Token
	expectedClient := &Client{
		Name:         auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description:  auth0.String("This is just a test client."),
		IsFirstParty: auth0.Bool(true),
		RefreshToken: &ClientRefreshToken{
			ExpirationType: auth0.String("expiring"),
			RotationType:   auth0.String("non-rotating"),
			Policies: &[]ClientRefreshTokenPolicy{
				{
					Audience: auth0.String(resourceServer.GetIdentifier()),
					Scope:    &[]string{"create:bar", "read:bar"},
				},
			},
		},
	}
	err = api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())
	actualClient, err := api.Client.Read(context.Background(), expectedClient.GetClientID())
	assert.NoError(t, err)
	assert.Equal(t, expectedClient.GetRefreshToken(), actualClient.GetRefreshToken())
	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestClientManager_ReadEnabledConnections(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)
	actualConnections, err := api.Client.ReadEnabledConnections(context.Background(), expectedClient.GetClientID())
	assert.NoError(t, err)
	assert.Equal(t, 0, len(actualConnections.Connections))

	expectedConnection := givenAOktaConnection(t)
	payload := []ConnectionEnabledClient{
		{
			ClientID: expectedClient.ClientID,
			Status:   auth0.Bool(true),
		},
	}
	err = api.Connection.UpdateEnabledClients(context.Background(), expectedConnection.GetID(), payload)

	assert.NoError(t, err)
	actualConnections, err = api.Client.ReadEnabledConnections(context.Background(), expectedClient.GetClientID())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actualConnections.Connections))
	assert.Equal(t, expectedConnection.GetID(), actualConnections.Connections[0].GetID())
}

func TestClient_CreateWithTokenExchange(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := &Client{
		Name:        auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
		TokenExchange: &ClientTokenExchange{
			AllowAnyProfileOfType: &[]string{"custom_authentication"},
		},
		OIDCConformant: auth0.Bool(true),
	}

	err := api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())
	actualClient, err := api.Client.Read(context.Background(), expectedClient.GetClientID())
	assert.NoError(t, err)
	assert.Equal(t, expectedClient.GetTokenExchange(), actualClient.GetTokenExchange())
	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestClient_SessionTransfer(t *testing.T) {
	configureHTTPTestRecordings(t)

	ctx := context.Background()

	clientName := auth0.Stringf("Test Client SessionTransfer (%s)", time.Now().Format(time.StampMilli))
	expectedClient := &Client{
		Name:        clientName,
		Description: auth0.String("This is a test client with Session Transfer."),
		SessionTransfer: &SessionTransfer{
			CanCreateSessionTransferToken: auth0.Bool(true),
			AllowedAuthenticationMethods:  &[]string{"cookie", "query"},
			EnforceDeviceBinding:          auth0.String("ip"),
			AllowRefreshToken:             auth0.Bool(true),
			EnforceCascadeRevocation:      auth0.Bool(true),
			EnforceOnlineRefreshTokens:    auth0.Bool(true),
		},
	}

	// Create client
	require.NoError(t, api.Client.Create(ctx, expectedClient))
	require.NotEmpty(t, expectedClient.GetClientID())

	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})

	// Verify creation
	created, err := api.Client.Read(ctx, expectedClient.GetClientID())
	require.NoError(t, err)
	require.NotNil(t, created.SessionTransfer)
	assert.Equal(t, expectedClient.GetSessionTransfer(), created.GetSessionTransfer())

	// Update session transfer
	created.SessionTransfer = &SessionTransfer{
		CanCreateSessionTransferToken: auth0.Bool(false),
		AllowedAuthenticationMethods:  &[]string{"cookie"},
		EnforceDeviceBinding:          auth0.String("none"),
		AllowRefreshToken:             auth0.Bool(false),
		EnforceCascadeRevocation:      auth0.Bool(false),
		EnforceOnlineRefreshTokens:    auth0.Bool(false),
	}

	// Strip fields not allowed on update
	created.ClientID = nil
	created.SigningKeys = nil

	if created.JWTConfiguration != nil {
		created.JWTConfiguration.SecretEncoded = nil
	}

	require.NoError(t, api.Client.Update(ctx, expectedClient.GetClientID(), created))

	// Verify update
	updated, err := api.Client.Read(ctx, expectedClient.GetClientID())
	require.NoError(t, err)
	require.NotNil(t, updated.SessionTransfer)
	assert.Equal(t, created.GetSessionTransfer(), updated.GetSessionTransfer())

	// Remove session transfer via PATCH
	type clientPatch struct {
		SessionTransfer *SessionTransfer `json:"session_transfer"`
	}

	patch := &clientPatch{SessionTransfer: nil}
	require.NoError(t, api.Request(ctx, http.MethodPatch, api.URI("clients", expectedClient.GetClientID()), patch))

	// Verify removal
	final, err := api.Client.Read(ctx, expectedClient.GetClientID())
	require.NoError(t, err)
	assert.Nil(t, final.GetSessionTransfer())
}

func TestClient_CreateWithDefaultOrg(t *testing.T) {
	configureHTTPTestRecordings(t)

	org := givenAnOrganization(t)

	expectedClient := &Client{
		Name:              auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description:       auth0.String("This is just a test client."),
		OrganizationUsage: auth0.String("allow"),
		DefaultOrganization: &ClientDefaultOrganization{
			Flows:          &[]string{"client_credentials"},
			OrganizationID: auth0.String(org.GetID()),
		},
	}

	err := api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())

	retrievedClient, err := api.Client.Read(context.Background(), expectedClient.GetClientID())
	assert.NoError(t, err)
	assert.NotEmpty(t, retrievedClient.DefaultOrganization.GetOrganizationID())
	assert.NotEmpty(t, retrievedClient.DefaultOrganization.GetFlows())

	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestClientSignedRequestObject(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := &Client{
		Name:        auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
		SignedRequestObject: &ClientSignedRequestObject{
			Required: auth0.Bool(true),
			Credentials: &[]Credential{
				{
					Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
					CredentialType: auth0.String("public_key"),
					PEM: auth0.String(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAua6LXMfgDE/tDdkOL1Oe
3oWUwg1r4dSTg9L7RCcI5hItUzmkVofHtWN0H4CH2lm2ANmaJUsnhzctYowYW2+R
tHvU9afTmtbdhpy993972hUqZSYLsE3iGziphYkOKVsqq38+VRH3TNg93zSLoRao
JnTTkMXseVqiyqYRmFN8+gQQoEclHSGPUWQG5XMZ+hhuXeFyo+Yw/qbZWca/6/2I
3rsca9jXR1alhxhHrXrg8N4Dm3gBgGbmiht6YYYT2Tyl1OqB9+iOI/9D7dfoCF6X
AWJXRE454cmC8k8oucpjZVpflA+ocKshwPDR6YTLQYbXYiaWxEoaz0QGUErNQBnG
I+sr9jDY3ua/s6HF6h0qyi/HVZH4wx+m4CtOfJoYTjrGBbaRszzUxhtSN2/MhXDu
+a35q9/2zcu/3fjkkfVvGUt+NyyiYOKQ9vsJC1g/xxdUWtowjNwjfZE2zcG4usi8
r38Bp0lmiipAsMLduZM/D5dFXkRdWCBNDfULmmg/4nv2wwjbjQuLemAMh7mmrztW
i/85WMnjKQZT8NqS43pmgyIzg1gK1neMqdS90YmQ/PvJ36qALxCs245w1JpN9BAL
JbwxCg/dbmKT7PalfWrksx9hGcJxtGqebldaOpw+5GVIPxxtC1C0gVr9BKeiDS3f
aibASY5pIRiKENmbZELDtucCAwEAAQ==
-----END PUBLIC KEY-----`),
				},
			},
		},
		JWTConfiguration:                   &ClientJWTConfiguration{Algorithm: auth0.String("PS256")},
		RequirePushedAuthorizationRequests: auth0.Bool(true),
		ComplianceLevel:                    auth0.String("fapi1_adv_pkj_par"),
		RequireProofOfPossession:           auth0.Bool(true),
	}

	err := api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())
	assert.Equal(t, true, expectedClient.GetSignedRequestObject().GetRequired())
	assert.Equal(t, "fapi1_adv_pkj_par", expectedClient.GetComplianceLevel())
	assert.Equal(t, "PS256", expectedClient.GetJWTConfiguration().GetAlgorithm())
	assert.Equal(t, true, expectedClient.GetRequirePushedAuthorizationRequests())
	assert.Equal(t, true, expectedClient.GetRequireProofOfPossession())

	clientID := expectedClient.GetClientID()
	expectedClient.ClientID = nil                       // Read-Only: Additional properties not allowed.
	expectedClient.SigningKeys = nil                    // Read-Only: Additional properties not allowed.
	expectedClient.JWTConfiguration.SecretEncoded = nil // Read-Only: Additional properties not allowed.

	updatedClient := expectedClient
	updatedClient.SignedRequestObject.Required = auth0.Bool(false)
	updatedClient.ComplianceLevel = auth0.String("fapi1_adv_mtls_par")
	updatedClient.RequirePushedAuthorizationRequests = auth0.Bool(false)
	updatedClient.JWTConfiguration.Algorithm = auth0.String("RS256")
	updatedClient.RequireProofOfPossession = auth0.Bool(false)

	err = api.Client.Update(context.Background(), clientID, updatedClient)
	assert.NoError(t, err)

	assert.Equal(t, false, updatedClient.GetSignedRequestObject().GetRequired())
	assert.Equal(t, "fapi1_adv_mtls_par", updatedClient.GetComplianceLevel())
	assert.Equal(t, false, updatedClient.GetRequirePushedAuthorizationRequests())
	assert.Equal(t, "RS256", updatedClient.GetJWTConfiguration().GetAlgorithm())
	assert.Equal(t, false, updatedClient.GetRequireProofOfPossession())
	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestClientAuthenticationMethods(t *testing.T) {
	updateAndVerifyClient := func(t *testing.T, clientID string, updatedClient *Client) {
		err := api.Client.Update(context.Background(), clientID, updatedClient)
		assert.NoError(t, err)
		assert.Equal(t, "fapi1_adv_mtls_par", updatedClient.GetComplianceLevel())
		assert.Equal(t, false, updatedClient.GetRequirePushedAuthorizationRequests())
		assert.Equal(t, "RS256", updatedClient.GetJWTConfiguration().GetAlgorithm())
	}

	cleanupTestClient := func(t *testing.T, clientID string) {
		t.Cleanup(func() {
			cleanupClient(t, clientID)
		})
	}

	t.Run("GetTLSClientAuth", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		client := givenAClientAuthenticationMethodsClient(t, &TLSClientAuth{
			Credentials: &[]Credential{
				{
					Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
					CredentialType: auth0.String("cert_subject_dn"),
					PEM: auth0.String(`-----BEGIN CERTIFICATE-----
MIIDPDCCAiQCCQDWNMOIuzwDfzANBgkqhkiG9w0BAQUFADBgMQswCQYDVQQGEwJK
UDEOMAwGA1UECAwFVG9reW8xEzARBgNVBAcMCkNoaXlvZGEta3UxDzANBgNVBAoM
BkNsaWVudDEbMBkGA1UEAwwSY2xpZW50LmV4YW1wbGUub3JnMB4XDTE5MTAyODA3
MjczMFoXDTIwMTAyNzA3MjczMFowYDELMAkGA1UEBhMCSlAxDjAMBgNVBAgMBVRv
a3lvMRMwEQYDVQQHDApDaGl5b2RhLWt1MQ8wDQYDVQQKDAZDbGllbnQxGzAZBgNV
BAMMEmNsaWVudC5leGFtcGxlLm9yZzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAK2Oyc+BV4N5pYcp47opUwsb2NaJq4X+d5Itq8whpFlZ9uCCHzF5TWSF
XrpYscOp95veGPF42eT1grfxYyvjFotE76caHhBLCkIbBh6Vf222IGMwwBbSZfO9
J3eURtEADBvsZ117HkPVdjYqvt3Pr4RxdR12zG1TcBAoTLGchyr8nBqRADFhUTCL
msYaz1ADiQ/xbJN7VUNQpKhzRWHCdYS03HpbGjYCtAbl9dJnH2EepNF0emGiSPFq
df6taToyCr7oZjM7ufmKPjiiEDbeSYTf6kbPNmmjtoPNNLeejHjP9p0IYx7l0Gkj
mx4kSMLp4vSDftrFgGfcxzaMmKBsosMCAwEAATANBgkqhkiG9w0BAQUFAAOCAQEA
qzdDYbntFLPBlbwAQlpwIjvmvwzvkQt6qgZ9Y0oMAf7pxq3i9q7W1bDol0UF4pIM
z3urEJCHO8w18JRlfOnOENkcLLLntrjOUXuNkaCDLrnv8pnp0yeTQHkSpsyMtJi9
R6r6JT9V57EJ/pWQBgKlN6qMiBkIvX7U2hEMmhZ00h/E5xMmiKbySBiJV9fBzDRf
mAy1p9YEgLsEMLnGjKHTok+hd0BLvcmXVejdUsKCg84F0zqtXEDXLCiKcpXCeeWv
lmmXxC5PH/GEMkSPiGSR7+b1i0sSotsq+M3hbdwabpJ6nQLLbKkFSGcsQ87yL+gr
So6zun26vAUJTu1o9CIjxw==
-----END CERTIFICATE-----`),
				},
			},
		})

		clientID := client.GetClientID()
		client.ClientID = nil
		client.SigningKeys = nil
		client.JWTConfiguration.SecretEncoded = nil

		updatedClient := client
		updatedClient.ComplianceLevel = auth0.String("fapi1_adv_mtls_par")
		updatedClient.RequirePushedAuthorizationRequests = auth0.Bool(false)
		updatedClient.JWTConfiguration.Algorithm = auth0.String("RS256")

		updateAndVerifyClient(t, clientID, updatedClient)
		cleanupTestClient(t, client.GetClientID())
	})

	t.Run("GetSelfSignedTLSClientAuth", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		client := givenAClientAuthenticationMethodsClient(t, &SelfSignedTLSClientAuth{
			Credentials: &[]Credential{
				{
					Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
					CredentialType: auth0.String("x509_cert"),
					PEM: auth0.String(`-----BEGIN CERTIFICATE-----
MIIDwTCCAyqgAwIBAgICDh4wDQYJKoZIhvcNAQEFBQAwgZsxCzAJBgNVBAYTAkpQ
MQ4wDAYDVQQIEwVUb2t5bzEQMA4GA1UEBxMHQ2h1by1rdTERMA8GA1UEChMIRnJh
bms0REQxGDAWBgNVBAsTD1dlYkNlcnQgU3VwcG9ydDEYMBYGA1UEAxMPRnJhbms0
REQgV2ViIENBMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGZyYW5rNGRkLmNvbTAi
GA8wMDAwMDEwMTAwMDAwMVoYDzk5OTkxMjMxMjM1OTU5WjCBgTELMAkGA1UEBhMC
SlAxDjAMBgNVBAgTBVRva3lvMREwDwYDVQQKEwhGcmFuazRERDEQMA4GA1UECxMH
U3VwcG9ydDEiMCAGCSqGSIb3DQEJARYTcHVibGljQGZyYW5rNGRkLmNvbTEZMBcG
A1UEAxMQd3d3LmZyYW5rNGRkLmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkC
gYEA4rkBL30FzR2ZHZ1vpF9kGBO0DMwhu2pcrkcLJ0SEuf52ggo+md0tPis8f1KN
Tchxj6DtxWT3c7ECW0c1ALpu6mNVE+GaM94KsckSDehoPfbLjT9Apcc/F0mqvDsC
N6fPdDixWrjx6xKT7xXi3lCy1yIKRMHA6Ha+T4qPyyCyMPECAwEAAaOCASYwggEi
MAwGA1UdEwEB/wQCMAAwCwYDVR0PBAQDAgWgMB0GA1UdDgQWBBRWKE5tXPIyS0pC
fE5taGO5Q84gyTCB0AYDVR0jBIHIMIHFgBRi83vtBtSx1Zx/SOXvxckVYf3ZEaGB
oaSBnjCBmzELMAkGA1UEBhMCSlAxDjAMBgNVBAgTBVRva3lvMRAwDgYDVQQHEwdD
aHVvLWt1MREwDwYDVQQKEwhGcmFuazRERDEYMBYGA1UECxMPV2ViQ2VydCBTdXBw
b3J0MRgwFgYDVQQDEw9GcmFuazRERCBXZWIgQ0ExIzAhBgkqhkiG9w0BCQEWFHN1
cHBvcnRAZnJhbms0ZGQuY29tggkAxscECbwiW6AwEwYDVR0lBAwwCgYIKwYBBQUH
AwEwDQYJKoZIhvcNAQEFBQADgYEAfXCfXcePJwnMKc06qLa336cEPpXEsPed1bw4
xiIXfgZ39duBnN+Nv4a49Yl2kbh4JO8tcr5h8WYAI/a/69w8qBFQBUAjTEY/+lcw
9/6wU7UA3kh7yexeqDiNTRflnPUv3sfiVdLDTjqLWWAxGS8L26PjVaCUFfJLNiYJ
jerREgM=
-----END CERTIFICATE-----`),
				},
			},
		})

		clientID := client.GetClientID()
		client.ClientID = nil
		client.SigningKeys = nil
		client.JWTConfiguration.SecretEncoded = nil

		updatedClient := client
		updatedClient.ComplianceLevel = auth0.String("fapi1_adv_mtls_par")
		updatedClient.RequirePushedAuthorizationRequests = auth0.Bool(false)
		updatedClient.JWTConfiguration.Algorithm = auth0.String("RS256")

		updateAndVerifyClient(t, clientID, updatedClient)
		cleanupTestClient(t, client.GetClientID())
	})

	t.Run("GetPrivateKeyJWT", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		client := givenAClientAuthenticationMethodsClient(t, &PrivateKeyJWT{
			Credentials: &[]Credential{
				{
					Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
					CredentialType: auth0.String("public_key"),
					PEM: auth0.String(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA3njxXJoHnuN4hByBhSUo
0kIbXkJTA0wP0fig87MyVz5KgohPrPJgbRSZ7yz/MmXa4qRNHkWiClJybMS2a98M
6ELOFG8pfDb6J7JaJqx0Kvqn6xsGInbpwsth3K582Cxrp+Y+GBNja++8wDY5IqAi
TSKSZRNies0GO0grzQ7kj2p0+R7a0c86mdLO4JnGrHoBqEY1HcsfnJvkJkqETlGi
yMzDQw8Wkux7P59N/3wuroAI83+HMYl1fV39ek3L/GrsLjECrNe5/CVFtblNltyb
/va9+pAP7Ye5p6tTW2oj3fzUvdX3dYzENWEtRB7DBHXnfEHMjTaBiQeWb2yDHBCw
++Uh1OCKw9ZLYzoE6gcDQspYf+fFU3F0kuU4c//gSoNuj/iEjaNmOEK6S3xGy8fE
TjsC+0oF6YaokDZO9+NreL/sGxFfOAysybrKWrMoaYwa81RlpcmBGZM7H1M00zLH
PPfCYVhGhFs5X3Qzzt6MQE+msgMt9zeGH7liJbOSW2NGSJwbmn7q35YYIfJEoXRF
1iefT/9fJB9vhQhtYfCOe3AEpTQq6Yz5ViLhToBdsVDBbz2gmRLALs9/D91SE9T4
XzvXjHGyxWVu0jdvS9hyhJzP4165k1cYDgx8mmg0VxR7j79LmCUDsFcvvSrAOf6y
0zY7r4pmNyQQ0r4in/gs/wkCAwEAAQ==
-----END PUBLIC KEY-----`),
				},
			},
		})

		clientID := client.GetClientID()
		client.ClientID = nil
		client.SigningKeys = nil
		client.JWTConfiguration.SecretEncoded = nil

		updatedClient := client
		updatedClient.ComplianceLevel = auth0.String("fapi1_adv_mtls_par")
		updatedClient.RequirePushedAuthorizationRequests = auth0.Bool(false)
		updatedClient.JWTConfiguration.Algorithm = auth0.String("RS256")

		updateAndVerifyClient(t, clientID, updatedClient)
		cleanupTestClient(t, client.GetClientID())
	})
}

func TestClient_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)

	actualClient, err := api.Client.Read(context.Background(), expectedClient.GetClientID())

	assert.NoError(t, err)
	assert.Equal(t, expectedClient.GetName(), actualClient.GetName())
}

func TestClient_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenASimpleClient(t)

	expectedDescription := "This is more than just a test client."
	expectedClient.Description = &expectedDescription

	clientID := expectedClient.GetClientID()
	expectedClient.ClientID = nil                       // Read-Only: Additional properties not allowed.
	expectedClient.SigningKeys = nil                    // Read-Only: Additional properties not allowed.
	expectedClient.JWTConfiguration.SecretEncoded = nil // Read-Only: Additional properties not allowed.

	err := api.Client.Update(context.Background(), clientID, expectedClient)

	assert.NoError(t, err)
	assert.Equal(t, expectedDescription, *expectedClient.Description)
}

func TestClient_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)

	err := api.Client.Delete(context.Background(), expectedClient.GetClientID())
	assert.NoError(t, err)

	actualClient, err := api.Client.Read(context.Background(), expectedClient.GetClientID())

	assert.Empty(t, actualClient)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestClient_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)

	clientList, err := api.Client.List(context.Background(), IncludeFields("client_id"))

	assert.NoError(t, err)
	assert.Contains(t, clientList.Clients, &Client{ClientID: expectedClient.ClientID})
}

func TestClient_RotateSecret(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenASimpleClient(t)

	oldSecret := expectedClient.GetClientSecret()
	actualClient, err := api.Client.RotateSecret(context.Background(), expectedClient.GetClientID())

	assert.NoError(t, err)
	assert.NotEqual(t, oldSecret, actualClient.GetClientSecret())
}

func TestClient_CreateWithClientAddons(t *testing.T) {
	configureHTTPTestRecordings(t)

	t.Run("legacy mappings", func(t *testing.T) {
		mappings := map[string]string{
			"email": "User.Email",
			"name":  "User.Name",
		}
		client := &Client{
			Name:        auth0.Stringf("Test Client Legacy Mappings (%s)", time.Now().Format(time.StampMilli)),
			Description: auth0.String("Client with legacy mappings."),
			Addons: &ClientAddons{
				SAML2: &SAML2ClientAddon{
					Audience: auth0.String("legacy-audience"),
					Mappings: &mappings,
				},
			},
		}

		err := api.Client.Create(context.Background(), client)
		require.NoError(t, err)
		require.NotEmpty(t, client.GetClientID())
		t.Cleanup(func() { cleanupClient(t, client.GetClientID()) })

		// Assert immediately after creation
		addons := client.GetAddons()
		assert.Equal(t, "legacy-audience", addons.GetSAML2().GetAudience())
		assert.Equal(t, mappings, addons.GetSAML2().GetMappings())
		assert.NotNil(t, addons.GetSAML2().GetFlexibleMappings())

		// Now read back the client from the API and verify the same fields
		clientRead, err := api.Client.Read(context.Background(), client.GetClientID())
		require.NoError(t, err)

		addonsRead := clientRead.GetAddons()
		assert.Equal(t, "legacy-audience", addonsRead.GetSAML2().GetAudience())
		assert.Equal(t, mappings, addonsRead.GetSAML2().GetMappings())
		assert.NotNil(t, addonsRead.GetSAML2().GetFlexibleMappings())
	})

	t.Run("flexible mappings", func(t *testing.T) {
		flexible := map[string]interface{}{
			"email": []interface{}{"User.Email", "User.AltEmail"},
			"name":  "User.Name",
		}
		client := &Client{
			Name:        auth0.Stringf("Test Client Flexible Mappings (%s)", time.Now().Format(time.StampMilli)),
			Description: auth0.String("Client with flexible mappings."),
			Addons: &ClientAddons{
				SAML2: &SAML2ClientAddon{
					Audience:         auth0.String("flexible-audience"),
					FlexibleMappings: flexible,
				},
			},
		}

		err := api.Client.Create(context.Background(), client)
		require.NoError(t, err)
		require.NotEmpty(t, client.GetClientID())
		t.Cleanup(func() { cleanupClient(t, client.GetClientID()) })

		// Assert immediately after creation
		addons := client.GetAddons()
		assert.Equal(t, "flexible-audience", addons.GetSAML2().GetAudience())
		assert.Equal(t, flexible, addons.GetSAML2().GetFlexibleMappings())
		assert.Empty(t, addons.GetSAML2().GetMappings())

		// Now read back the client from the API and verify the same fields
		clientRead, err := api.Client.Read(context.Background(), client.GetClientID())
		require.NoError(t, err)

		addonsRead := clientRead.GetAddons()
		assert.Equal(t, "flexible-audience", addonsRead.GetSAML2().GetAudience())
		assert.Equal(t, flexible, addonsRead.GetSAML2().GetFlexibleMappings())
		assert.Empty(t, addonsRead.GetSAML2().GetMappings())
	})
	t.Run("both legacy and flexible mappings set", func(t *testing.T) {
		mappings := map[string]string{
			"email": "User.Email",
		}
		flexible := map[string]interface{}{
			"email": []interface{}{"User.Email", "User.AltEmail"},
		}
		client := &Client{
			Name:        auth0.Stringf("Test Client Both Mappings (%s)", time.Now().Format(time.StampMilli)),
			Description: auth0.String("Client with both legacy and flexible mappings."),
			Addons: &ClientAddons{
				SAML2: &SAML2ClientAddon{
					Audience:         auth0.String("hybrid-audience"),
					Mappings:         &mappings,
					FlexibleMappings: flexible,
				},
			},
		}
		err := api.Client.Create(context.Background(), client)
		require.NoError(t, err)
		require.NotEmpty(t, client.GetClientID())
		t.Cleanup(func() { cleanupClient(t, client.GetClientID()) })

		client, err = api.Client.Read(context.Background(), client.GetClientID())
		require.NoError(t, err)

		addons := client.GetAddons()
		assert.Equal(t, "hybrid-audience", addons.GetSAML2().GetAudience())
		assert.Equal(t, flexible, addons.GetSAML2().GetFlexibleMappings())
		assert.Empty(t, addons.GetSAML2().GetMappings())
	})
}

func TestClient_SAML2ClientAddonMappings(t *testing.T) {
	t.Run("Legacy mappings", func(t *testing.T) {
		legacyMappings := map[string]string{
			"email": "User.Email",
			"name":  "User.Name",
		}
		saml2 := &SAML2ClientAddon{
			Mappings: &legacyMappings,
		}

		// Verify initial getters
		assert.Equal(t, legacyMappings, saml2.GetMappings())
		assert.Empty(t, saml2.GetFlexibleMappings())

		// Marshal & unmarshal to check JSON round-trip
		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		var decoded SAML2ClientAddon
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)

		assert.Equal(t, legacyMappings, decoded.GetMappings())
		assert.Equal(t, map[string]interface{}{"email": "User.Email", "name": "User.Name"}, decoded.GetFlexibleMappings())
	})

	t.Run("Flexible mappings", func(t *testing.T) {
		flexibleMappings := map[string]interface{}{
			"email": []interface{}{"User.Email", "User.AltEmail"},
			"name":  "User.Name",
		}
		saml2 := &SAML2ClientAddon{
			FlexibleMappings: flexibleMappings,
		}

		// Initial getters
		assert.Empty(t, saml2.GetMappings())
		assert.Equal(t, flexibleMappings, saml2.GetFlexibleMappings())

		// Marshal & unmarshal
		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		var decoded SAML2ClientAddon
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)

		assert.Empty(t, decoded.GetMappings())
		assert.Equal(t, flexibleMappings, decoded.GetFlexibleMappings())
	})

	t.Run("Flexible mappings with string slice", func(t *testing.T) {
		flexibleMappings := map[string]interface{}{
			"roles": []interface{}{"admin", "user"},
		}
		saml2 := &SAML2ClientAddon{
			FlexibleMappings: flexibleMappings,
		}
		assert.Empty(t, saml2.GetMappings())
		assert.Equal(t, flexibleMappings, saml2.GetFlexibleMappings())

		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		var decoded SAML2ClientAddon
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)

		assert.Empty(t, decoded.GetMappings())
		assert.Equal(t, flexibleMappings, decoded.GetFlexibleMappings())
	})

	t.Run("Both mappings set - prefer flexible", func(t *testing.T) {
		legacyMappings := map[string]string{
			"email": "User.Email",
		}
		flexibleMappings := map[string]interface{}{
			"email": []interface{}{"User.Email", "User.AltEmail"},
		}
		saml2 := &SAML2ClientAddon{
			Mappings:         &legacyMappings,
			FlexibleMappings: flexibleMappings,
		}

		// Check initial state
		assert.Equal(t, flexibleMappings, saml2.GetFlexibleMappings())
		assert.Equal(t, legacyMappings, saml2.GetMappings())

		// Marshal & unmarshal
		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		var decoded SAML2ClientAddon
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)

		assert.Empty(t, decoded.GetMappings()) // Legacy mappings should be ignored
		assert.Equal(t, flexibleMappings, decoded.GetFlexibleMappings())
	})
	t.Run("Invalid mappings type", func(t *testing.T) {
		input := `{"mappings": "not-an-object"}`

		var decoded SAML2ClientAddon
		err := json.Unmarshal([]byte(input), &decoded)
		assert.Error(t, err)
	})
	t.Run("Marshal with legacy mappings", func(t *testing.T) {
		legacyMappings := map[string]string{
			"email": "User.Email",
			"name":  "User.Name",
		}
		saml2 := &SAML2ClientAddon{
			Mappings: &legacyMappings,
		}

		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		expectedJSON := `{
		"mappings": {
			"email": "User.Email",
			"name": "User.Name"
		}}`

		assert.JSONEq(t, expectedJSON, string(data))
	})
	t.Run("Marshal with flexible mappings", func(t *testing.T) {
		flexibleMappings := map[string]interface{}{
			"email": []interface{}{"User.Email", "User.AltEmail"},
			"name":  "User.Name",
		}
		saml2 := &SAML2ClientAddon{
			FlexibleMappings: flexibleMappings,
		}

		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		expectedJSON := `{
		"mappings": {
			"email": ["User.Email", "User.AltEmail"],
			"name": "User.Name"
		}
	}`

		assert.JSONEq(t, expectedJSON, string(data))
	})

	t.Run("Marshal with both mappings - prefer flexible", func(t *testing.T) {
		legacy := map[string]string{"email": "User.Email"}
		flexible := map[string]interface{}{
			"email": []interface{}{"User.Email", "User.AltEmail"},
		}
		saml2 := &SAML2ClientAddon{
			Mappings:         &legacy,
			FlexibleMappings: flexible,
		}

		data, err := json.Marshal(saml2)
		assert.NoError(t, err)

		expectedJSON := `{
		"mappings": {
			"email": ["User.Email", "User.AltEmail"]
		}
	}`

		assert.JSONEq(t, expectedJSON, string(data))
	})
}

func TestClient_CreateWithOIDCLogout(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := &Client{
		Name: auth0.Stringf("Test Client OIDC Logout (%s)", time.Now().Format(time.StampMilli)),
		OIDCLogout: &OIDCLogout{
			BackChannelLogoutURLs: &[]string{"https://example.com/logout"},
			BackChannelLogoutInitiators: &BackChannelLogoutInitiators{
				Mode: auth0.String("custom"),
				SelectedInitiators: &[]string{
					"rp-logout",
					"idp-logout",
				},
			},
		},
	}

	err := api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())

	oidcLogout := expectedClient.GetOIDCLogout()
	assert.Equal(t, oidcLogout.GetBackChannelLogoutURLs(), []string{"https://example.com/logout"})
	assert.Equal(t, oidcLogout.GetBackChannelLogoutInitiators().GetMode(), "custom")
	assert.Equal(t, oidcLogout.GetBackChannelLogoutInitiators().GetSelectedInitiators(), []string{"rp-logout", "idp-logout"})

	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
	})
}

func TestJWTConfiguration(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		for clientJWTConfiguration, expected := range map[*ClientJWTConfiguration]string{
			{}:                                   `{}`,
			{LifetimeInSeconds: auth0.Int(1000)}: `{"lifetime_in_seconds":1000}`,
		} {
			jsonBody, err := json.Marshal(clientJWTConfiguration)
			assert.NoError(t, err)
			assert.Equal(t, string(jsonBody), expected)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		for jsonBody, expected := range map[string]*ClientJWTConfiguration{
			`{}`:                             {LifetimeInSeconds: nil},
			`{"lifetime_in_seconds":1000}`:   {LifetimeInSeconds: auth0.Int(1000)},
			`{"lifetime_in_seconds":"1000"}`: {LifetimeInSeconds: auth0.Int(1000)},
		} {
			var actual ClientJWTConfiguration
			err := json.Unmarshal([]byte(jsonBody), &actual)

			assert.NoError(t, err)
			assert.Equal(t, &actual, expected)
		}

		t.Run("Should error if unexpected type", func(t *testing.T) {
			var actual ClientJWTConfiguration
			err := json.Unmarshal([]byte(`{"lifetime_in_seconds":true}`), &actual)
			assert.Contains(t, err.Error(), "unexpected type for field lifetime_in_seconds")

			err = json.Unmarshal([]byte(`{"lifetime_in_seconds":"fooo"}`), &actual)
			assert.Contains(t, err.Error(), "unexpected type for field lifetime_in_seconds")
		})
	})
}

func TestClient_CreateCredential(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)

	credential := &Credential{
		Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
		CredentialType: auth0.String("public_key"),
		PEM: auth0.String(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA3njxXJoHnuN4hByBhSUo
0kIbXkJTA0wP0fig87MyVz5KgohPrPJgbRSZ7yz/MmXa4qRNHkWiClJybMS2a98M
6ELOFG8pfDb6J7JaJqx0Kvqn6xsGInbpwsth3K582Cxrp+Y+GBNja++8wDY5IqAi
TSKSZRNies0GO0grzQ7kj2p0+R7a0c86mdLO4JnGrHoBqEY1HcsfnJvkJkqETlGi
yMzDQw8Wkux7P59N/3wuroAI83+HMYl1fV39ek3L/GrsLjECrNe5/CVFtblNltyb
/va9+pAP7Ye5p6tTW2oj3fzUvdX3dYzENWEtRB7DBHXnfEHMjTaBiQeWb2yDHBCw
++Uh1OCKw9ZLYzoE6gcDQspYf+fFU3F0kuU4c//gSoNuj/iEjaNmOEK6S3xGy8fE
TjsC+0oF6YaokDZO9+NreL/sGxFfOAysybrKWrMoaYwa81RlpcmBGZM7H1M00zLH
PPfCYVhGhFs5X3Qzzt6MQE+msgMt9zeGH7liJbOSW2NGSJwbmn7q35YYIfJEoXRF
1iefT/9fJB9vhQhtYfCOe3AEpTQq6Yz5ViLhToBdsVDBbz2gmRLALs9/D91SE9T4
XzvXjHGyxWVu0jdvS9hyhJzP4165k1cYDgx8mmg0VxR7j79LmCUDsFcvvSrAOf6y
0zY7r4pmNyQQ0r4in/gs/wkCAwEAAQ==
-----END PUBLIC KEY-----`),
	}

	err := api.Client.CreateCredential(context.Background(), expectedClient.GetClientID(), credential)
	assert.NoError(t, err)
	assert.NotEmpty(t, credential.GetID())

	t.Cleanup(func() {
		cleanupCredential(t, expectedClient.GetClientID(), credential.GetID())
	})
}

func TestClient_ListCredentials(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)
	createdCredentials := expectedClient.GetClientAuthenticationMethods().GetPrivateKeyJWT().GetCredentials()
	createdCredential := createdCredentials[0]
	expectedCredential := givenACredential(t, expectedClient)

	credentials, err := api.Client.ListCredentials(context.Background(), expectedClient.GetClientID())

	assert.NoError(t, err)
	assert.Equal(t, createdCredential.GetID(), credentials[0].GetID())
	assert.Equal(t, expectedCredential.GetID(), credentials[1].GetID())
}

func TestClient_GetCredentials(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)
	expectedCredential := givenACredential(t, expectedClient)

	credential, err := api.Client.GetCredential(context.Background(), expectedClient.GetClientID(), expectedCredential.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedCredential.GetID(), credential.GetID())
}

func TestClient_UpdateCredential(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)
	expectedCredential := givenACredential(t, expectedClient)

	expiresAt := time.Now().Add(time.Minute * 10)
	expectedCredential.ExpiresAt = &expiresAt

	pem := expectedCredential.GetPEM()
	credentialID := expectedCredential.GetID()
	expectedCredential.ID = nil

	err := api.Client.UpdateCredential(context.Background(), expectedClient.GetClientID(), credentialID, expectedCredential)

	assert.NoError(t, err)
	assert.Equal(t, expectedCredential.GetExpiresAt(), expiresAt)
	assert.Equal(t, expectedCredential.GetID(), credentialID) // Check that we unmarshall the result into this struct.
	assert.Equal(t, expectedCredential.GetPEM(), pem)
}

func TestClient_DeleteCredential(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedClient := givenAClient(t)
	expectedCredential := givenACredential(t, expectedClient)

	err := api.Client.DeleteCredential(context.Background(), expectedClient.GetClientID(), expectedCredential.GetID())
	assert.NoError(t, err)

	actualCredential, err := api.Client.Read(context.Background(), expectedCredential.GetID())

	assert.Empty(t, actualCredential)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}
func TestClient_CreateAllCredential(t *testing.T) {
	t.Run("Should create PrivateJWT Credential", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		expectedClient := givenAClient(t)

		credential := &Credential{
			Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
			CredentialType: auth0.String("public_key"),
			PEM: auth0.String(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA3njxXJoHnuN4hByBhSUo
0kIbXkJTA0wP0fig87MyVz5KgohPrPJgbRSZ7yz/MmXa4qRNHkWiClJybMS2a98M
6ELOFG8pfDb6J7JaJqx0Kvqn6xsGInbpwsth3K582Cxrp+Y+GBNja++8wDY5IqAi
TSKSZRNies0GO0grzQ7kj2p0+R7a0c86mdLO4JnGrHoBqEY1HcsfnJvkJkqETlGi
yMzDQw8Wkux7P59N/3wuroAI83+HMYl1fV39ek3L/GrsLjECrNe5/CVFtblNltyb
/va9+pAP7Ye5p6tTW2oj3fzUvdX3dYzENWEtRB7DBHXnfEHMjTaBiQeWb2yDHBCw
++Uh1OCKw9ZLYzoE6gcDQspYf+fFU3F0kuU4c//gSoNuj/iEjaNmOEK6S3xGy8fE
TjsC+0oF6YaokDZO9+NreL/sGxFfOAysybrKWrMoaYwa81RlpcmBGZM7H1M00zLH
PPfCYVhGhFs5X3Qzzt6MQE+msgMt9zeGH7liJbOSW2NGSJwbmn7q35YYIfJEoXRF
1iefT/9fJB9vhQhtYfCOe3AEpTQq6Yz5ViLhToBdsVDBbz2gmRLALs9/D91SE9T4
XzvXjHGyxWVu0jdvS9hyhJzP4165k1cYDgx8mmg0VxR7j79LmCUDsFcvvSrAOf6y
0zY7r4pmNyQQ0r4in/gs/wkCAwEAAQ==
-----END PUBLIC KEY-----`),
		}

		err := api.Client.CreateCredential(context.Background(), expectedClient.GetClientID(), credential)
		assert.NoError(t, err)
		assert.NotEmpty(t, credential.GetID())

		t.Cleanup(func() {
			cleanupCredential(t, expectedClient.GetClientID(), credential.GetID())
		})
	})
	t.Run("Should create TLSClientAuth Credential", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		expectedClient := givenAClient(t)

		credential := &Credential{
			Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
			CredentialType: auth0.String("cert_subject_dn"),
			PEM: auth0.String(`-----BEGIN CERTIFICATE-----
MIIDPDCCAiQCCQDWNMOIuzwDfzANBgkqhkiG9w0BAQUFADBgMQswCQYDVQQGEwJK
UDEOMAwGA1UECAwFVG9reW8xEzARBgNVBAcMCkNoaXlvZGEta3UxDzANBgNVBAoM
BkNsaWVudDEbMBkGA1UEAwwSY2xpZW50LmV4YW1wbGUub3JnMB4XDTE5MTAyODA3
MjczMFoXDTIwMTAyNzA3MjczMFowYDELMAkGA1UEBhMCSlAxDjAMBgNVBAgMBVRv
a3lvMRMwEQYDVQQHDApDaGl5b2RhLWt1MQ8wDQYDVQQKDAZDbGllbnQxGzAZBgNV
BAMMEmNsaWVudC5leGFtcGxlLm9yZzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAK2Oyc+BV4N5pYcp47opUwsb2NaJq4X+d5Itq8whpFlZ9uCCHzF5TWSF
XrpYscOp95veGPF42eT1grfxYyvjFotE76caHhBLCkIbBh6Vf222IGMwwBbSZfO9
J3eURtEADBvsZ117HkPVdjYqvt3Pr4RxdR12zG1TcBAoTLGchyr8nBqRADFhUTCL
msYaz1ADiQ/xbJN7VUNQpKhzRWHCdYS03HpbGjYCtAbl9dJnH2EepNF0emGiSPFq
df6taToyCr7oZjM7ufmKPjiiEDbeSYTf6kbPNmmjtoPNNLeejHjP9p0IYx7l0Gkj
mx4kSMLp4vSDftrFgGfcxzaMmKBsosMCAwEAATANBgkqhkiG9w0BAQUFAAOCAQEA
qzdDYbntFLPBlbwAQlpwIjvmvwzvkQt6qgZ9Y0oMAf7pxq3i9q7W1bDol0UF4pIM
z3urEJCHO8w18JRlfOnOENkcLLLntrjOUXuNkaCDLrnv8pnp0yeTQHkSpsyMtJi9
R6r6JT9V57EJ/pWQBgKlN6qMiBkIvX7U2hEMmhZ00h/E5xMmiKbySBiJV9fBzDRf
mAy1p9YEgLsEMLnGjKHTok+hd0BLvcmXVejdUsKCg84F0zqtXEDXLCiKcpXCeeWv
lmmXxC5PH/GEMkSPiGSR7+b1i0sSotsq+M3hbdwabpJ6nQLLbKkFSGcsQ87yL+gr
So6zun26vAUJTu1o9CIjxw==
-----END CERTIFICATE-----`),
		}

		err := api.Client.CreateCredential(context.Background(), expectedClient.GetClientID(), credential)
		assert.NoError(t, err)
		assert.NotEmpty(t, credential.GetID())

		t.Cleanup(func() {
			cleanupCredential(t, expectedClient.GetClientID(), credential.GetID())
		})
	})
	t.Run("Should create SelfSignedTLSClientAuth Credential", func(t *testing.T) {
		configureHTTPTestRecordings(t)

		expectedClient := givenAClient(t)

		credential := &Credential{
			Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
			CredentialType: auth0.String("x509_cert"),
			PEM: auth0.String(`-----BEGIN CERTIFICATE-----
MIIDwTCCAyqgAwIBAgICDh4wDQYJKoZIhvcNAQEFBQAwgZsxCzAJBgNVBAYTAkpQ
MQ4wDAYDVQQIEwVUb2t5bzEQMA4GA1UEBxMHQ2h1by1rdTERMA8GA1UEChMIRnJh
bms0REQxGDAWBgNVBAsTD1dlYkNlcnQgU3VwcG9ydDEYMBYGA1UEAxMPRnJhbms0
REQgV2ViIENBMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGZyYW5rNGRkLmNvbTAi
GA8wMDAwMDEwMTAwMDAwMVoYDzk5OTkxMjMxMjM1OTU5WjCBgTELMAkGA1UEBhMC
SlAxDjAMBgNVBAgTBVRva3lvMREwDwYDVQQKEwhGcmFuazRERDEQMA4GA1UECxMH
U3VwcG9ydDEiMCAGCSqGSIb3DQEJARYTcHVibGljQGZyYW5rNGRkLmNvbTEZMBcG
A1UEAxMQd3d3LmZyYW5rNGRkLmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkC
gYEA4rkBL30FzR2ZHZ1vpF9kGBO0DMwhu2pcrkcLJ0SEuf52ggo+md0tPis8f1KN
Tchxj6DtxWT3c7ECW0c1ALpu6mNVE+GaM94KsckSDehoPfbLjT9Apcc/F0mqvDsC
N6fPdDixWrjx6xKT7xXi3lCy1yIKRMHA6Ha+T4qPyyCyMPECAwEAAaOCASYwggEi
MAwGA1UdEwEB/wQCMAAwCwYDVR0PBAQDAgWgMB0GA1UdDgQWBBRWKE5tXPIyS0pC
fE5taGO5Q84gyTCB0AYDVR0jBIHIMIHFgBRi83vtBtSx1Zx/SOXvxckVYf3ZEaGB
oaSBnjCBmzELMAkGA1UEBhMCSlAxDjAMBgNVBAgTBVRva3lvMRAwDgYDVQQHEwdD
aHVvLWt1MREwDwYDVQQKEwhGcmFuazRERDEYMBYGA1UECxMPV2ViQ2VydCBTdXBw
b3J0MRgwFgYDVQQDEw9GcmFuazRERCBXZWIgQ0ExIzAhBgkqhkiG9w0BCQEWFHN1
cHBvcnRAZnJhbms0ZGQuY29tggkAxscECbwiW6AwEwYDVR0lBAwwCgYIKwYBBQUH
AwEwDQYJKoZIhvcNAQEFBQADgYEAfXCfXcePJwnMKc06qLa336cEPpXEsPed1bw4
xiIXfgZ39duBnN+Nv4a49Yl2kbh4JO8tcr5h8WYAI/a/69w8qBFQBUAjTEY/+lcw
9/6wU7UA3kh7yexeqDiNTRflnPUv3sfiVdLDTjqLWWAxGS8L26PjVaCUFfJLNiYJ
jerREgM=
-----END CERTIFICATE-----`),
		}

		err := api.Client.CreateCredential(context.Background(), expectedClient.GetClientID(), credential)
		assert.NoError(t, err)
		assert.NotEmpty(t, credential.GetID())

		t.Cleanup(func() {
			cleanupCredential(t, expectedClient.GetClientID(), credential.GetID())
		})
	})
}

func givenASimpleClient(t *testing.T) *Client {
	t.Helper()

	client := &Client{
		Name:              auth0.Stringf("Test Simple Client (%s)", time.Now().Format(time.StampMilli)),
		Description:       auth0.String("This is just a simple test client."),
		OrganizationUsage: auth0.String("allow"),
	}

	err := api.Client.Create(context.Background(), client)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupClient(t, client.GetClientID())
	})

	return client
}

func givenAClient(t *testing.T) *Client {
	t.Helper()

	client := &Client{
		Name:              auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description:       auth0.String("This is just a test client."),
		OrganizationUsage: auth0.String("allow"),
		ClientAuthenticationMethods: &ClientAuthenticationMethods{
			PrivateKeyJWT: &PrivateKeyJWT{
				Credentials: &[]Credential{
					{
						Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
						CredentialType: auth0.String("public_key"),
						PEM: auth0.String(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAua6LXMfgDE/tDdkOL1Oe
3oWUwg1r4dSTg9L7RCcI5hItUzmkVofHtWN0H4CH2lm2ANmaJUsnhzctYowYW2+R
tHvU9afTmtbdhpy993972hUqZSYLsE3iGziphYkOKVsqq38+VRH3TNg93zSLoRao
JnTTkMXseVqiyqYRmFN8+gQQoEclHSGPUWQG5XMZ+hhuXeFyo+Yw/qbZWca/6/2I
3rsca9jXR1alhxhHrXrg8N4Dm3gBgGbmiht6YYYT2Tyl1OqB9+iOI/9D7dfoCF6X
AWJXRE454cmC8k8oucpjZVpflA+ocKshwPDR6YTLQYbXYiaWxEoaz0QGUErNQBnG
I+sr9jDY3ua/s6HF6h0qyi/HVZH4wx+m4CtOfJoYTjrGBbaRszzUxhtSN2/MhXDu
+a35q9/2zcu/3fjkkfVvGUt+NyyiYOKQ9vsJC1g/xxdUWtowjNwjfZE2zcG4usi8
r38Bp0lmiipAsMLduZM/D5dFXkRdWCBNDfULmmg/4nv2wwjbjQuLemAMh7mmrztW
i/85WMnjKQZT8NqS43pmgyIzg1gK1neMqdS90YmQ/PvJ36qALxCs245w1JpN9BAL
JbwxCg/dbmKT7PalfWrksx9hGcJxtGqebldaOpw+5GVIPxxtC1C0gVr9BKeiDS3f
aibASY5pIRiKENmbZELDtucCAwEAAQ==
-----END PUBLIC KEY-----`),
					},
				},
			},
		},
		JWTConfiguration: &ClientJWTConfiguration{
			Algorithm: auth0.String("RS256"),
		},
	}

	err := api.Client.Create(context.Background(), client)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupClient(t, client.GetClientID())
	})

	return client
}

func givenAClientAuthenticationMethodsClient(t *testing.T, authMethod interface{}) *Client {
	client := &Client{
		Name:        auth0.Stringf("Test Client (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client."),
		ClientAuthenticationMethods: &ClientAuthenticationMethods{
			TLSClientAuth:           nil,
			SelfSignedTLSClientAuth: nil,
			PrivateKeyJWT:           nil,
		},
		JWTConfiguration:                   &ClientJWTConfiguration{Algorithm: auth0.String("PS256")},
		RequirePushedAuthorizationRequests: auth0.Bool(true),
		ComplianceLevel:                    auth0.String("fapi1_adv_pkj_par"),
	}

	switch v := authMethod.(type) {
	case *TLSClientAuth:
		client.ClientAuthenticationMethods.TLSClientAuth = v
	case *SelfSignedTLSClientAuth:
		client.ClientAuthenticationMethods.SelfSignedTLSClientAuth = v
	case *PrivateKeyJWT:
		client.ClientAuthenticationMethods.PrivateKeyJWT = v
	default:
		t.Fatalf("Unsupported authentication method")
	}

	err := api.Client.Create(context.Background(), client)
	assert.NoError(t, err)
	assert.NotEmpty(t, client.GetClientID())
	assert.Equal(t, "fapi1_adv_pkj_par", client.GetComplianceLevel())
	assert.Equal(t, "PS256", client.GetJWTConfiguration().GetAlgorithm())
	assert.Equal(t, true, client.GetRequirePushedAuthorizationRequests())

	return client
}

func cleanupClient(t *testing.T, clientID string) {
	t.Helper()

	err := api.Client.Delete(context.Background(), clientID)
	require.NoError(t, err)
}

func givenACredential(t *testing.T, client *Client) *Credential {
	t.Helper()

	credential := &Credential{
		Name:           auth0.Stringf("Test Credential (%s)", time.Now().Format(time.StampMilli)),
		CredentialType: auth0.String("public_key"),
		PEM: auth0.String(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA3njxXJoHnuN4hByBhSUo
0kIbXkJTA0wP0fig87MyVz5KgohPrPJgbRSZ7yz/MmXa4qRNHkWiClJybMS2a98M
6ELOFG8pfDb6J7JaJqx0Kvqn6xsGInbpwsth3K582Cxrp+Y+GBNja++8wDY5IqAi
TSKSZRNies0GO0grzQ7kj2p0+R7a0c86mdLO4JnGrHoBqEY1HcsfnJvkJkqETlGi
yMzDQw8Wkux7P59N/3wuroAI83+HMYl1fV39ek3L/GrsLjECrNe5/CVFtblNltyb
/va9+pAP7Ye5p6tTW2oj3fzUvdX3dYzENWEtRB7DBHXnfEHMjTaBiQeWb2yDHBCw
++Uh1OCKw9ZLYzoE6gcDQspYf+fFU3F0kuU4c//gSoNuj/iEjaNmOEK6S3xGy8fE
TjsC+0oF6YaokDZO9+NreL/sGxFfOAysybrKWrMoaYwa81RlpcmBGZM7H1M00zLH
PPfCYVhGhFs5X3Qzzt6MQE+msgMt9zeGH7liJbOSW2NGSJwbmn7q35YYIfJEoXRF
1iefT/9fJB9vhQhtYfCOe3AEpTQq6Yz5ViLhToBdsVDBbz2gmRLALs9/D91SE9T4
XzvXjHGyxWVu0jdvS9hyhJzP4165k1cYDgx8mmg0VxR7j79LmCUDsFcvvSrAOf6y
0zY7r4pmNyQQ0r4in/gs/wkCAwEAAQ==
-----END PUBLIC KEY-----`),
	}

	err := api.Client.CreateCredential(context.Background(), client.GetClientID(), credential)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupCredential(t, client.GetClientID(), credential.GetID())
	})

	return credential
}

func cleanupCredential(t *testing.T, clientID string, credentialID string) {
	t.Helper()

	err := api.Client.DeleteCredential(context.Background(), clientID, credentialID)
	require.NoError(t, err)
}
