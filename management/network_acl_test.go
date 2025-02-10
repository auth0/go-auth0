package management

import (
	"context"
	"github.com/auth0/go-auth0"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNetworkACLManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expectedNetworkACL := &NetworkACL{
		Description: auth0.String("some-description"),
		Active:      auth0.Bool(true),
		Priority:    auth0.Int(1),
		Rule: &NetworkACLRule{
			Action: &NetworkACLRuleAction{
				Block: auth0.Bool(true),
			},
			Match: &NetworkACLRuleMatch{
				IPv6Cidrs: &[]string{"2001:db8::/32"},
			},
			Scope: auth0.String("management"),
		},
	}

	err := api.NetworkACL.Create(context.Background(), expectedNetworkACL)
	assert.NoError(t, err)
	assert.NotEmpty(t, expectedNetworkACL.GetID())
	t.Cleanup(func() {
		cleanupNetworkACL(t, expectedNetworkACL.GetID())
	})
	actualNetworkACL, err := api.NetworkACL.Read(context.Background(), expectedNetworkACL.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedNetworkACL, actualNetworkACL)
}

func TestNetworkACLManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedNetworkACL := givenANetworkACL(t)
	networkACLList, err := api.NetworkACL.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(networkACLList.NetworkACLs), 0)
	assert.Contains(t, networkACLList.NetworkACLs, expectedNetworkACL)
}

func TestNetworkACLManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedNetworkACL := givenANetworkACL(t)
	actualNetworkACL, err := api.NetworkACL.Read(context.Background(), expectedNetworkACL.GetID())
	assert.NoError(t, err)
	assert.Equal(t, expectedNetworkACL, actualNetworkACL)
}

func TestNetworkACLManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	networkACL := givenANetworkACL(t)
	networkACL.Description = auth0.String("updated-description")
	networkACL.Rule = &NetworkACLRule{
		Action: &NetworkACLRuleAction{
			Block: auth0.Bool(true),
		},
		Match: &NetworkACLRuleMatch{
			GeoCountryCodes: &[]string{"US", "CA"},
		},
		Scope: auth0.String("authentication"),
	}
	err := api.NetworkACL.Update(context.Background(), networkACL.GetID(), networkACL)
	assert.NoError(t, err)
	actualNetworkACL, err := api.NetworkACL.Read(context.Background(), networkACL.GetID())
	assert.NoError(t, err)
	assert.Equal(t, networkACL, actualNetworkACL)
}

func givenANetworkACL(t *testing.T) *NetworkACL {
	networkACL := &NetworkACL{
		Description: auth0.String("some-description"),
		Active:      auth0.Bool(true),
		Priority:    auth0.Int(1),
		Rule: &NetworkACLRule{
			Action: &NetworkACLRuleAction{
				Redirect:    auth0.Bool(true),
				RedirectURI: auth0.String("https://example.com"),
			},
			Match: &NetworkACLRuleMatch{
				GeoCountryCodes: &[]string{"US"},
			},
			Scope: auth0.String("authentication"),
		},
	}

	err := api.NetworkACL.Create(context.Background(), networkACL)
	assert.NoError(t, err)
	assert.NotEmpty(t, networkACL.GetID())
	t.Cleanup(func() {
		cleanupNetworkACL(t, networkACL.GetID())
	})
	return networkACL
}

func cleanupNetworkACL(t *testing.T, networkACLID string) {
	t.Helper()

	err := api.NetworkACL.Delete(context.Background(), networkACLID)
	if err != nil {
		if err.(Error).Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
