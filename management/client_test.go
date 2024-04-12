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

	expectedClient := &Client{
		Name:        auth0.Stringf("Test Client Addons (%s)", time.Now().Format(time.StampMilli)),
		Description: auth0.String("This is just a test client with addons."),
		Addons: &ClientAddons{
			SAML2: &SAML2ClientAddon{
				Audience: auth0.String("my-audience"),
			},
			WSFED: &WSFEDClientAddon{},
		},
	}

	err := api.Client.Create(context.Background(), expectedClient)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedClient.GetClientID())

	addons := expectedClient.GetAddons()
	assert.Equal(t, "my-audience", addons.GetSAML2().GetAudience())
	assert.NotNil(t, addons.GetWSFED())

	t.Cleanup(func() {
		cleanupClient(t, expectedClient.GetClientID())
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
