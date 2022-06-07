package management

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestCustomDomainManager_Create(t *testing.T) {
	setupHTTPRecordings(t)

	expected := &CustomDomain{
		Domain:    auth0.Stringf("%d.auth.uat.auth0.com", time.Now().UTC().Unix()),
		Type:      auth0.String("auth0_managed_certs"),
		TLSPolicy: auth0.String("recommended"),
	}

	err := m.CustomDomain.Create(expected)
	assertNoCustomDomainErr(t, err)
	assert.NotEmpty(t, expected.GetID())

	t.Cleanup(func() {
		cleanupCustomDomain(t, expected.GetID())
	})
}

func TestCustomDomainManager_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expected := givenACustomDomain(t)

	actual, err := m.CustomDomain.Read(expected.GetID())

	assertNoCustomDomainErr(t, err)
	assert.Equal(t, expected.GetDomain(), actual.GetDomain())
}

func TestCustomDomainManager_Update(t *testing.T) {
	setupHTTPRecordings(t)

	customDomain := givenACustomDomain(t)

	err := m.CustomDomain.Update(customDomain.GetID(), &CustomDomain{TLSPolicy: auth0.String("recommended")})
	assertNoCustomDomainErr(t, err)

	actual, err := m.CustomDomain.Read(customDomain.GetID())

	assertNoCustomDomainErr(t, err)
	assert.Equal(t, "recommended", actual.GetTLSPolicy())
}

func TestCustomDomainManager_Delete(t *testing.T) {
	setupHTTPRecordings(t)

	customDomain := givenACustomDomain(t)

	err := m.CustomDomain.Delete(customDomain.GetID())
	assertNoCustomDomainErr(t, err)

	_, err = m.CustomDomain.Read(customDomain.GetID())

	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func TestCustomDomainManager_List(t *testing.T) {
	setupHTTPRecordings(t)

	customDomain := givenACustomDomain(t)

	customDomainList, err := m.CustomDomain.List()

	assertNoCustomDomainErr(t, err)
	assert.Len(t, customDomainList, 1)
	assert.Equal(t, customDomainList[0].GetID(), customDomain.GetID())
}

func TestCustomDomainManager_Verify(t *testing.T) {
	setupHTTPRecordings(t)

	customDomain := givenACustomDomain(t)

	actualDomain, err := m.CustomDomain.Verify(customDomain.GetID())

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

	err := m.CustomDomain.Create(customDomain)
	assertNoCustomDomainErr(t, err)

	t.Cleanup(func() {
		cleanupCustomDomain(t, customDomain.GetID())
	})

	return customDomain
}

func cleanupCustomDomain(t *testing.T, customDomainID string) {
	t.Helper()

	err := m.CustomDomain.Delete(customDomainID)
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
