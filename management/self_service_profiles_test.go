package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

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
		UserAttributes: []*UserAttributes{
			{
				Name:        auth0.String("some-name-here"),
				Description: auth0.String("some-description"),
				IsOptional:  auth0.Bool(true),
			},
		},
	}

	err := api.SelfServiceProfile.Create(context.Background(), ssop)
	assert.Equal(t, ssop.UserAttributes[0].GetName(), "some-name-here")
	assert.NotEmpty(t, ssop.ID)
	assert.NoError(t, err)

	ssops, err := api.SelfServiceProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(ssops), 0)
	assert.Equal(t, ssops[0].UserAttributes[0].GetName(), ssop.UserAttributes[0].GetName())
	assert.Equal(t, ssops[0].UserAttributes[0].GetDescription(), ssop.UserAttributes[0].GetDescription())
	assert.Equal(t, ssops[0].UserAttributes[0].GetIsOptional(), ssop.UserAttributes[0].GetIsOptional())
	assert.Equal(t, ssops[0].GetBranding().GetColors().GetPrimary(), ssop.GetBranding().GetColors().GetPrimary())
	assert.Equal(t, ssops[0].GetBranding().GetLogoURL(), ssop.GetBranding().GetLogoURL())

	cleanSelfServiceProfile(t, ssop.GetID())
}

func TestSelfServiceProfileManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssopList, err := api.SelfServiceProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(ssopList), 0)
	assert.Equal(t, ssopList[0].GetBranding().GetColors().GetPrimary(), ssop.GetBranding().GetColors().GetPrimary())
}

func TestSelfServiceProfileManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssopRetrieved, err := api.SelfServiceProfile.Read(context.Background(), ssop.GetID())
	assert.NoError(t, err)
	assert.Equal(t, ssopRetrieved.UserAttributes[0].GetName(), ssop.UserAttributes[0].GetName())
}

func TestSelfServiceProfileManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	oldName := ssop.UserAttributes[0].GetName()
	ssop.UserAttributes[0].Name = auth0.String("some-new-name-here")
	id := ssop.GetID()
	ssop.CreatedAt, ssop.UpdatedAt, ssop.ID = nil, nil, nil
	err := api.SelfServiceProfile.Update(context.Background(), id, ssop)
	assert.NoError(t, err)
	assert.NotEqual(t, oldName, ssop.UserAttributes[0].GetName())
}

func TestSelfServiceProfileManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	err := api.SelfServiceProfile.Delete(context.Background(), ssop.GetID())
	assert.NoError(t, err)
}

func TestSelfServiceProfileManager_CreateTicket(t *testing.T) {
	configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	client := givenAClient(t)
	org := givenAnOrganization(t)

	ticket := &SSOTicket{
		ConnectionConfig: &ConnectionConfig{
			Name: "sso-generated-ticket",
		},
		EnabledClients: []*string{auth0.String(client.GetClientID())},
		EnabledOrganizations: []*EnabledOrganizations{
			{
				org.GetID(),
			},
		},
	}
	err := api.SelfServiceProfile.CreateTicket(context.Background(), ssop.GetID(), ticket)
	assert.NoError(t, err)
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
		UserAttributes: []*UserAttributes{
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
