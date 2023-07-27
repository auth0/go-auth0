package management

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestCustomDomainManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expected := &CustomDomain{
		Domain:    auth0.Stringf("%d.auth.uat.auth0.com", time.Now().UTC().Unix()),
		Type:      auth0.String("auth0_managed_certs"),
		TLSPolicy: auth0.String("recommended"),
	}

	err := api.CustomDomain.Create(context.Background(), expected)
	assertNoCustomDomainErr(t, err)
	assert.NotEmpty(t, expected.GetID())

	t.Cleanup(func() {
		cleanupCustomDomain(t, expected.GetID())
	})
}

func TestCustomDomainManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	expected := givenACustomDomain(t)

	actual, err := api.CustomDomain.Read(context.Background(), expected.GetID())

	assertNoCustomDomainErr(t, err)
	assert.Equal(t, expected.GetDomain(), actual.GetDomain())
}

func TestCustomDomainManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)

	err := api.CustomDomain.Update(context.Background(), customDomain.GetID(), &CustomDomain{TLSPolicy: auth0.String("recommended")})
	assertNoCustomDomainErr(t, err)

	actual, err := api.CustomDomain.Read(context.Background(), customDomain.GetID())

	assertNoCustomDomainErr(t, err)
	assert.Equal(t, "recommended", actual.GetTLSPolicy())
}

func TestCustomDomainManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)

	err := api.CustomDomain.Delete(context.Background(), customDomain.GetID())
	assertNoCustomDomainErr(t, err)

	_, err = api.CustomDomain.Read(context.Background(), customDomain.GetID())

	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestCustomDomainManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)

	customDomainList, err := api.CustomDomain.List(context.Background())

	assertNoCustomDomainErr(t, err)
	assert.Len(t, customDomainList, 1)
	assert.Equal(t, customDomainList[0].GetID(), customDomain.GetID())
}

func TestCustomDomainManager_Verify(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)

	actualDomain, err := api.CustomDomain.Verify(context.Background(), customDomain.GetID())

	assertNoCustomDomainErr(t, err)
	assert.Equal(t, "pending_verification", actualDomain.GetStatus())
}

func givenACustomDomain(t *testing.T) *CustomDomain {
	t.Helper()

	customDomain := &CustomDomain{
		Domain:    auth0.Stringf("%d.auth.uat.auth0.com", time.Now().UTC().Unix()),
		Type:      auth0.String("auth0_managed_certs"),
		TLSPolicy: auth0.String("recommended"),
	}

	err := api.CustomDomain.Create(context.Background(), customDomain)
	assertNoCustomDomainErr(t, err)

	t.Cleanup(func() {
		cleanupCustomDomain(t, customDomain.GetID())
	})

	return customDomain
}

func cleanupCustomDomain(t *testing.T, customDomainID string) {
	t.Helper()

	err := api.CustomDomain.Delete(context.Background(), customDomainID)
	assertNoCustomDomainErr(t, err)
}

func assertNoCustomDomainErr(t *testing.T, err error) {
	if err != nil {
		if mErr, ok := err.(Error); ok && mErr.Status() == http.StatusForbidden {
			t.Skip(err)
		} else {
			t.Error(err)
		}
	}
}
