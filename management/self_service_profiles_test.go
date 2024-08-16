package management

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/auth0/go-auth0"
)

func TestSelfServiceProfileManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := &SelfServiceProfile{
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
	assert.Greater(t, len(ssopList), 0)
	assert.Contains(t, ssopList, ssop)
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

func TestSelfServiceProfileManager_CreateTicket(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	client := givenAClient(t)
	org := givenAnOrganization(t)

	ticket := &SelfServiceProfileTicket{
		ConnectionConfig: &SelfServiceProfileTicketConnectionConfig{
			Name: "sso-generated-ticket",
		},
		EnabledClients: []*string{auth0.String(client.GetClientID())},
		EnabledOrganizations: []*SelfServiceProfileTicketEnabledOrganizations{
			{
				org.GetID(),
			},
		},
	}
	err := api.SelfServiceProfile.CreateTicket(context.Background(), ssop.GetID(), ticket)
	assert.NoError(t, err)
	assert.NotEmpty(t, ticket.GetTicket())

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
	} {
		payload, err := json.Marshal(selfServiceProfile)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func givenASelfServiceProfile(t *testing.T) *SelfServiceProfile {
	t.Helper()
	ssop := &SelfServiceProfile{
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
