package management

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestSelfServiceProfileManager_Create(t *testing.T) {
	// configureHTTPTestRecordings(t)
	ssop := &SelfServiceProfile{
		//SAMLMappings: []*SAMLMappings{
		//	{
		//		Name:        auth0.String("mapping0-1719998496750"),
		//		Description: auth0.String("This is a mapping"),
		//		IsOptional:  auth0.Bool(true),
		//	},
		//},
		//OIDCScopes: []*string{auth0.String("sub"), auth0.String("profile")},
		Branding: &Branding{
			LogoURL: auth0.String("https://example.com/logo.png"),
			Colors: &BrandingColors{
				Primary: auth0.String("#334455"),
			},
		},
		UserAttributes: []*UserAttributes{
			{
				Name:        auth0.String("some-name-here"),
				DisplayName: auth0.String("some-display-name"),
				Description: auth0.String("some-description"),
				IsOptional:  auth0.Bool(true),
				DefaultMappings: &DefaultMappings{
					Saml: auth0.String("some-saml"),
					Oidc: auth0.String("some-oidc"),
				},
			},
		},
	}

	err := api.SelfServiceProfile.Create(context.Background(), ssop)
	assert.Equal(t, ssop.UserAttributes[0].GetName(), "some-name-here")
	assert.NotEmpty(t, ssop.ID)
	assert.NoError(t, err)

	ssops, err := api.SelfServiceProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ssops))
	assert.Equal(t, ssops[0].UserAttributes[0].GetName(), ssop.UserAttributes[0].GetName())
	assert.Equal(t, ssops[0].UserAttributes[0].GetDisplayName(), ssop.UserAttributes[0].GetDisplayName())
	assert.Equal(t, ssops[0].UserAttributes[0].GetDescription(), ssop.UserAttributes[0].GetDescription())
	assert.Equal(t, ssops[0].UserAttributes[0].GetIsOptional(), ssop.UserAttributes[0].GetIsOptional())
	assert.Equal(t, ssops[0].UserAttributes[0].GetDefaultMappings().GetSaml(), ssop.UserAttributes[0].GetDefaultMappings().GetSaml())
	assert.Equal(t, ssops[0].UserAttributes[0].GetDefaultMappings().GetOidc(), ssop.UserAttributes[0].GetDefaultMappings().GetOidc())
	assert.Equal(t, ssops[0].GetBranding().GetColors().GetPrimary(), ssop.GetBranding().GetColors().GetPrimary())
	assert.Equal(t, ssops[0].GetBranding().GetLogoURL(), ssop.GetBranding().GetLogoURL())

	cleanSelfServiceProfile(t, ssop.GetID())
}

func TestSelfServiceProfileManager_List(t *testing.T) {
	// configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssopList, err := api.SelfServiceProfile.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ssopList))
	assert.Equal(t, ssop.GetBranding().GetColors().GetPrimary(), ssopList[0].GetBranding().GetColors().GetPrimary())
}

func TestSelfServiceProfileManager_Read(t *testing.T) {
	// configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	ssop, err := api.SelfServiceProfile.Read(context.Background(), ssop.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "some-name-here", ssop.UserAttributes[0].GetName())
}

func TestSelfServiceProfileManager_Update(t *testing.T) {
	// configureHTTPTestRecordings(t)
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
	// configureHTTPTestRecordings(t)
	ssop := givenASelfServiceProfile(t)
	err := api.SelfServiceProfile.Delete(context.Background(), ssop.GetID())
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
				DisplayName: auth0.String("some-display-name"),
				Description: auth0.String("some-description"),
				IsOptional:  auth0.Bool(true),
				DefaultMappings: &DefaultMappings{
					Saml: auth0.String("string"),
					Oidc: auth0.String("string"),
				},
			},
		},
	}

	_ = api.SelfServiceProfile.Create(context.Background(), ssop)
	t.Cleanup(func() {
		cleanSelfServiceProfile(t, ssop.GetID())
	})

	return ssop
}

func cleanSelfServiceProfile(t *testing.T, id string) {
	t.Helper()
	_ = api.SelfServiceProfile.Delete(context.Background(), id)
}
