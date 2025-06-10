package management

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
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
	assert.Greater(t, len(networkACLList), 0)
	assert.Contains(t, networkACLList, expectedNetworkACL)
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
	networkACL.Priority = auth0.Int(2)
	networkACL.Active = auth0.Bool(false)
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

// TestNetworkACLManager_Patch tests the ability to patch a Network ACL.
func TestNetworkACLManager_Patch(t *testing.T) {
	configureHTTPTestRecordings(t)
	networkACL := givenANetworkACL(t)

	patch := &NetworkACL{
		Description: auth0.String("patched-description"),
		Priority:    auth0.Int(2),
	}

	err := api.NetworkACL.Patch(context.Background(), networkACL.GetID(), patch)
	assert.NoError(t, err)

	actualNetworkACL, err := api.NetworkACL.Read(context.Background(), networkACL.GetID())
	assert.NoError(t, err)
	assert.Equal(t, "patched-description", actualNetworkACL.GetDescription())
	assert.Equal(t, 2, actualNetworkACL.GetPriority())
}
func TestNetworkACLManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	networkACL := givenANetworkACL(t)
	err := api.NetworkACL.Delete(context.Background(), networkACL.GetID())
	assert.NoError(t, err)
	_, err = api.NetworkACL.Read(context.Background(), networkACL.GetID())
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
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
