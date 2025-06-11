package management

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestSelfServiceProfileManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	ssop := &SelfServiceProfile{
		Name:              auth0.String("Sample Self Service Profile"),
		Description:       auth0.String("Sample Desc"),
		AllowedStrategies: &[]string{"oidc"},
		Branding: &Branding{
			LogoURL: auth0.String("https://example.com/logo.png"),
			Colors: &BrandingColors{
				Primary: auth0.String("#334455"),
			},
		},
		UserAttributes: []*SelfServiceProfileUserAttributes{
			{
				Name:        auth0.String("some-name-here"),
				Description: auth0.String("some-description"),
				IsOptional:  auth0.Bool(true),
			},
		},
	}

	err := api.SelfServiceProfile.Create(context.Background(), ssop)
	assert.NoError(t, err)
	assert.NotEmpty(t, ssop.GetID())

	ssopRetrieved, err := api.SelfServiceProfile.Read(context.Background(), ssop.GetID())
	assert.NoError(t, err)
	assert.Equal(t, ssopRetrieved, ssop)
	cleanSelfServiceProfile(t, ssop.GetID())
}

func TestSelfServiceProfileManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssopList, err := api.SelfServiceProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(ssopList.SelfServiceProfile), 0)
	assert.Contains(t, ssopList.SelfServiceProfile, ssop)
}

func TestSelfServiceProfileManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssopRetrieved, err := api.SelfServiceProfile.Read(context.Background(), ssop.GetID())
	assert.NoError(t, err)
	assert.Equal(t, ssopRetrieved, ssop)
}

func TestSelfServiceProfileManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssopUpdated := &SelfServiceProfile{
		UserAttributes: []*SelfServiceProfileUserAttributes{
			{
				Name:        auth0.String("some-new-name-here"),
				Description: auth0.String("some-description"),
				IsOptional:  auth0.Bool(true),
			},
		},
	}

	err := api.SelfServiceProfile.Update(context.Background(), ssop.GetID(), ssopUpdated)
	assert.NoError(t, err)
	assert.NotEqual(t, ssop, ssopUpdated)
	assert.Equal(t, ssop.GetID(), ssopUpdated.GetID())
	assert.Equal(t, ssop.GetBranding(), ssopUpdated.GetBranding())
	assert.Equal(t, ssop.UserAttributes[0].GetDescription(), ssopUpdated.UserAttributes[0].GetDescription())
	assert.Equal(t, ssop.UserAttributes[0].GetIsOptional(), ssopUpdated.UserAttributes[0].GetIsOptional())
}

func TestSelfServiceProfileManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	err := api.SelfServiceProfile.Delete(context.Background(), ssop.GetID())
	assert.NoError(t, err)

	retrievedProfile, err := api.SelfServiceProfile.Read(context.Background(), ssop.GetID())
	assert.Nil(t, retrievedProfile)
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestSelfServiceProfileManager_SetGetCustomText(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	payload := map[string]interface{}{
		"introduction": "\"Welcome! With only a few steps you'll be able to setup your new connection.\"\n",
	}
	err := api.SelfServiceProfile.SetCustomText(context.Background(), ssop.GetID(), "en", "get-started", payload)
	assert.Equal(t, err, nil)
	retrievedCustomText, err := api.SelfServiceProfile.GetCustomText(context.Background(), ssop.GetID(), "en", "get-started")
	assert.Equal(t, err, nil)
	assert.Equal(t, payload, retrievedCustomText)
}

func TestSelfServiceProfileManager_CreateAndRevokeTicket(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	client := givenAClient(t)
	org := givenAnOrganization(t)

	ticket := &SelfServiceProfileTicket{
		ConnectionConfig: &SelfServiceProfileTicketConnectionConfig{
			Name:               auth0.String("sso-generated-ticket"),
			DisplayName:        auth0.String("sso-generated-ticket-display-name"),
			IsDomainConnection: auth0.Bool(true),
			ShowAsButton:       auth0.Bool(true),
			Metadata: &map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			Options: &SelfServiceProfileTicketConnectionConfigOptions{
				IconURL:       auth0.String("https://metabox.com/my_icon.jpeg"),
				DomainAliases: &[]string{"okta.com"},
				IDPInitiated: &SelfServiceProfileTicketConnectionConfigOptionsIDPInitiated{
					Enabled:              auth0.Bool(true),
					ClientID:             auth0.String(client.GetClientID()),
					ClientProtocol:       auth0.String("oauth2"),
					ClientAuthorizeQuery: auth0.String("scope=openid,profile,email"),
				},
			},
		},
		EnabledClients: &[]string{client.GetClientID()},
		EnabledOrganizations: []*SelfServiceProfileTicketEnabledOrganizations{
			{
				AssignMembershipOnLogin: auth0.Bool(true),
				ShowAsButton:            auth0.Bool(true),
				OrganizationID:          auth0.String(org.GetID()),
			},
		},
		DomainAliasesConfig: &SelfServiceProfileTicketDomainAliasesConfig{
			DomainVerification: auth0.String("none"),
		},
	}
	err := api.SelfServiceProfile.CreateTicket(context.Background(), ssop.GetID(), ticket)
	assert.NoError(t, err)
	assert.NotEmpty(t, ticket.GetTicket())

	ticketURL := ticket.GetTicket()

	ticketIDMap, err := url.ParseQuery(ticketURL)
	if err != nil {
		ticketID := ticketIDMap["ticketId"][0]
		err = api.SelfServiceProfile.RevokeTicket(context.Background(), ssop.GetID(), ticketID)
	}

	assert.NoError(t, err)
}

func TestSelfServiceProfileManager_MarshalJSON(t *testing.T) {
	for selfServiceProfile, expected := range map[*SelfServiceProfile]string{
		{}: `{}`,
		{
			UserAttributes: []*SelfServiceProfileUserAttributes{
				{
					Name:        auth0.String("some-name"),
					Description: auth0.String("some-desc"),
					IsOptional:  auth0.Bool(true),
				},
			},
		}: `{"user_attributes":[{"name":"some-name","description":"some-desc","is_optional":true}]}`,
		{
			Branding: &Branding{
				LogoURL:    auth0.String("some-url"),
				FaviconURL: auth0.String("some-favicon-url"),
			},
		}: `{"branding":{"favicon_url":"some-favicon-url","logo_url":"some-url"}}`,
		{
			ID:        auth0.String("some-id"),
			CreatedAt: auth0.Time(time.Now()),
			UpdatedAt: auth0.Time(time.Now()),
		}: `{}`,
	} {
		payload, err := json.Marshal(selfServiceProfile)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func givenASelfServiceProfile(t *testing.T) *SelfServiceProfile {
	t.Helper()

	ssop := &SelfServiceProfile{
		Name:              auth0.String("Sample Self Service Profile"),
		Description:       auth0.String("Sample Desc"),
		AllowedStrategies: &[]string{"oidc"},
		Branding: &Branding{
			LogoURL: auth0.String("https://example.com/logo.png"),
			Colors: &BrandingColors{
				Primary: auth0.String("#334455"),
			},
		},
		UserAttributes: []*SelfServiceProfileUserAttributes{
			{
				Name:        auth0.String("some-name-here"),
				Description: auth0.String("some-description"),
				IsOptional:  auth0.Bool(true),
			},
		},
	}

	err := api.SelfServiceProfile.Create(context.Background(), ssop)
	assert.NoError(t, err)
	t.Cleanup(func() {
		cleanSelfServiceProfile(t, ssop.GetID())
	})

	return ssop
}

func cleanSelfServiceProfile(t *testing.T, id string) {
	t.Helper()

	err := api.SelfServiceProfile.Delete(context.Background(), id)
	assert.NoError(t, err)
}
