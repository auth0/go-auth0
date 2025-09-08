package management

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestResourceServer_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := &ResourceServer{
		Name:                auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:          auth0.String("https://api.example.com/"),
		SigningAlgorithm:    auth0.String("PS256"),
		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),
		Scopes: &[]ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
		},
		EnforcePolicies: auth0.Bool(true),
		TokenDialect:    auth0.String("rfc9068_profile_authz"),
		ConsentPolicy:   auth0.String("transactional-authorization-with-mfa"),
		AuthorizationDetails: &[]ResourceServerAuthorizationDetails{
			{
				Type: auth0.String("payment"),
			},
			{
				Type: auth0.String("my custom type"),
			},
		},
		TokenEncryption: &ResourceServerTokenEncryption{
			Format: auth0.String("compact-nested-jwe"),
			EncryptionKey: &ResourceServerTokenEncryptionKey{
				Name: auth0.String("my JWE public key"),
				Alg:  auth0.String("RSA-OAEP-256"),
				Kid:  auth0.String("my-key-id"),
				Pem: auth0.String(`-----BEGIN PUBLIC KEY-----
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
		ProofOfPossession: &ResourceServerProofOfPossession{
			Mechanism: auth0.String("mtls"),
			Required:  auth0.Bool(true),
		},
		SubjectTypeAuthorization: &ResourceServerSubjectTypeAuthorization{
			User: &ResourceServerSubjectTypeAuthorizationUser{
				Policy: auth0.String("allow_all"),
			},
			Client: &ResourceServerSubjectTypeAuthorizationClient{
				Policy: auth0.String("require_client_grant"),
			},
		},
	}

	err := api.ResourceServer.Create(context.Background(), expectedResourceServer)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedResourceServer.GetID())

	t.Cleanup(func() {
		cleanupResourceServer(t, expectedResourceServer.GetID())
	})
}

func TestResourceServer_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	actualResourceServer, err := api.ResourceServer.Read(context.Background(), expectedResourceServer.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetName(), actualResourceServer.GetName())
}

func TestResourceServer_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	resourceServerID := expectedResourceServer.GetID()

	expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
	expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
	expectedResourceServer.SigningSecret = nil

	expectedResourceServer.AllowOfflineAccess = auth0.Bool(true)
	expectedResourceServer.SigningAlgorithm = auth0.String("RS256")
	expectedResourceServer.SkipConsentForVerifiableFirstPartyClients = auth0.Bool(true)
	expectedResourceServer.TokenLifetime = auth0.Int(7200)
	expectedResourceServer.TokenLifetimeForWeb = auth0.Int(5400)
	scopes := append(expectedResourceServer.GetScopes(), ResourceServerScope{
		Value:       auth0.String("update:resource"),
		Description: auth0.String("Update Resource"),
	})
	expectedResourceServer.Scopes = &scopes
	expectedResourceServer.EnforcePolicies = auth0.Bool(true)
	expectedResourceServer.TokenDialect = auth0.String("access_token_authz")
	expectedResourceServer.SubjectTypeAuthorization = &ResourceServerSubjectTypeAuthorization{
		User: &ResourceServerSubjectTypeAuthorizationUser{
			Policy: auth0.String("deny_all"),
		},
		Client: &ResourceServerSubjectTypeAuthorizationClient{
			Policy: auth0.String("deny_all"),
		},
	}
	err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)

	assert.NoError(t, err)
	assert.Equal(t, expectedResourceServer.GetAllowOfflineAccess(), true)
	assert.Equal(t, expectedResourceServer.GetSigningAlgorithm(), "RS256")
	assert.Equal(t, expectedResourceServer.GetSkipConsentForVerifiableFirstPartyClients(), true)
	assert.Equal(t, expectedResourceServer.GetTokenLifetime(), 7200)
	assert.Equal(t, expectedResourceServer.GetTokenLifetimeForWeb(), 5400)
	assert.Equal(t, len(expectedResourceServer.GetScopes()), 3)
	assert.Equal(t, expectedResourceServer.GetTokenDialect(), "access_token_authz")
	assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), true)
	assert.Equal(t, expectedResourceServer.GetSubjectTypeAuthorization(), &ResourceServerSubjectTypeAuthorization{
		User: &ResourceServerSubjectTypeAuthorizationUser{
			Policy: auth0.String("deny_all"),
		},
		Client: &ResourceServerSubjectTypeAuthorizationClient{
			Policy: auth0.String("deny_all"),
		},
	})
}

func TestResourceServer_TokenDialect(t *testing.T) {
	t.Run("When_TokenDialect_is_rfc9068_profile_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("rfc9068_profile")
		expectedResourceServer.EnforcePolicies = auth0.Bool(false)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "rfc9068_profile")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), false)
	})

	t.Run("When_TokenDialect_is_access_token_authz_and_RBAC_enabled_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("access_token_authz")
		expectedResourceServer.EnforcePolicies = auth0.Bool(true)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "access_token_authz")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), true)
	})

	t.Run("When_TokenDialect_is_rfc9068_profile_authz_and_RBAC_enabled_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("rfc9068_profile_authz")
		expectedResourceServer.EnforcePolicies = auth0.Bool(true)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "rfc9068_profile_authz")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), true)
	})

	t.Run("When_TokenDialect_is_access_token_should_succeed", func(t *testing.T) {
		configureHTTPTestRecordings(t)
		expectedResourceServer := givenAResourceServer(t)

		resourceServerID := expectedResourceServer.GetID()

		expectedResourceServer.ID = nil         // Read-Only: Additional properties not allowed.
		expectedResourceServer.Identifier = nil // Read-Only: Additional properties not allowed.
		expectedResourceServer.SigningSecret = nil

		expectedResourceServer.TokenDialect = auth0.String("access_token")
		expectedResourceServer.EnforcePolicies = auth0.Bool(false)

		err := api.ResourceServer.Update(context.Background(), resourceServerID, expectedResourceServer)
		assert.NoError(t, err)
		assert.Equal(t, expectedResourceServer.GetTokenDialect(), "access_token")
		assert.Equal(t, expectedResourceServer.GetEnforcePolicies(), false)
	})
}

func TestResourceServer_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)

	err := api.ResourceServer.Delete(context.Background(), expectedResourceServer.GetID())
	assert.NoError(t, err)

	actualResourceServer, err := api.ResourceServer.Read(context.Background(), expectedResourceServer.GetID())
	assert.Empty(t, actualResourceServer)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestResourceServer_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedResourceServer := givenAResourceServer(t)
	resourceServerList, err := api.ResourceServer.List(context.Background(), IncludeFields("id", "identifier"), Parameter("identifiers", expectedResourceServer.GetIdentifier()))
	require.NoError(t, err)
	assert.NotEqual(t, len(resourceServerList.ResourceServers), 0)
	assert.NoError(t, err)
	assert.Contains(t, resourceServerList.ResourceServers, &ResourceServer{ID: expectedResourceServer.ID, Identifier: expectedResourceServer.Identifier})
}

func givenAResourceServer(t *testing.T) *ResourceServer {
	t.Helper()

	resourceServer := &ResourceServer{
		Name:                auth0.Stringf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:          auth0.String("https://api.example.com/"),
		SigningAlgorithm:    auth0.String("HS256"),
		TokenLifetime:       auth0.Int(7200),
		TokenLifetimeForWeb: auth0.Int(3600),
		TokenDialect:        auth0.String("access_token"),
		EnforcePolicies:     auth0.Bool(false),
		Scopes: &[]ResourceServerScope{
			{
				Value:       auth0.String("create:resource"),
				Description: auth0.String("Create Resource"),
			},
			{
				Value:       auth0.String("create:organization_client_grants"),
				Description: auth0.String("Create Org Client Grants"),
			},
		},
		SubjectTypeAuthorization: &ResourceServerSubjectTypeAuthorization{
			User: &ResourceServerSubjectTypeAuthorizationUser{
				Policy: auth0.String("allow_all"),
			},
			Client: &ResourceServerSubjectTypeAuthorizationClient{
				Policy: auth0.String("require_client_grant"),
			},
		},
	}

	err := api.ResourceServer.Create(context.Background(), resourceServer)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupResourceServer(t, resourceServer.GetID())
	})

	return resourceServer
}

func cleanupResourceServer(t *testing.T, resourceServerID string) {
	t.Helper()

	err := api.ResourceServer.Delete(context.Background(), resourceServerID)
	require.NoError(t, err)
}
