package management

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestCustomDomainManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	expected := &CustomDomain{
		Domain:         auth0.Stringf("%d.tempdomain.com", time.Now().UTC().Unix()),
		Type:           auth0.String("auth0_managed_certs"),
		TLSPolicy:      auth0.String("recommended"),
		DomainMetadata: &map[string]interface{}{"key": "value"},
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

	t.Run("Verify initial metadata", func(t *testing.T) {
		initial, err := api.CustomDomain.Read(context.Background(), customDomain.GetID())
		assertNoCustomDomainErr(t, err)
		assert.Equal(t, map[string]interface{}{"key": "value"}, initial.GetDomainMetadata())
	})

	t.Run("Append a new metadata key", func(t *testing.T) {
		err := api.CustomDomain.Update(context.Background(), customDomain.GetID(),
			&CustomDomain{
				TLSPolicy:      auth0.String("recommended"),
				DomainMetadata: &map[string]interface{}{"key2": "value2"},
			},
		)
		assertNoCustomDomainErr(t, err)

		updated, err := api.CustomDomain.Read(context.Background(), customDomain.GetID())
		assertNoCustomDomainErr(t, err)

		expected := map[string]interface{}{
			"key":  "value",  // original preserved
			"key2": "value2", // newly added
		}

		assert.Equal(t, "recommended", updated.GetTLSPolicy())
		assert.Equal(t, expected, updated.GetDomainMetadata())
	})

	t.Run("Remove all metadata keys explicitly", func(t *testing.T) {
		err := api.CustomDomain.Update(context.Background(), customDomain.GetID(),
			&CustomDomain{
				DomainMetadata: &map[string]interface{}{
					"key":  nil,
					"key2": nil,
				},
			},
		)
		assertNoCustomDomainErr(t, err)

		cleared, err := api.CustomDomain.Read(context.Background(), customDomain.GetID())
		assertNoCustomDomainErr(t, err)
		assert.Empty(t, cleared.GetDomainMetadata())
	})
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
	customDomain2 := givenACustomDomain(t)

	customDomainList, err := api.CustomDomain.List(context.Background())

	assertNoCustomDomainErr(t, err)
	assert.Greater(t, len(customDomainList), 2)
	// Create a map to check existence regardless of order
	domainMap := make(map[string]bool)
	for _, domain := range customDomainList {
		domainMap[domain.GetID()] = true
	}

	// Check both domains are in the list
	assert.True(t, domainMap[customDomain.GetID()], "First custom domain should be in the list")
	assert.True(t, domainMap[customDomain2.GetID()], "Second custom domain should be in the list")
}

func TestCustomDomainManager_ListWithPagination(t *testing.T) {
	configureHTTPTestRecordings(t)

	// Create 3 custom domains to ensure pagination spans multiple pages
	domain1 := givenACustomDomain(t)
	domain2 := givenACustomDomain(t)
	domain3 := givenACustomDomain(t)

	// Request the first page (limit 2)
	firstPageOpts := []RequestOption{Take(2)}
	firstPage, err := api.CustomDomain.ListWithPagination(context.Background(), firstPageOpts...)
	assertNoCustomDomainErr(t, err)
	assert.Len(t, firstPage.CustomDomains, 2, "First page should return 2 domains")

	// Request the second page using the checkpoint token
	secondPageOpts := []RequestOption{
		From(firstPage.Next),
		Take(2),
	}
	secondPage, err := api.CustomDomain.ListWithPagination(context.Background(), secondPageOpts...)
	assertNoCustomDomainErr(t, err)
	assert.Greater(t, len(secondPage.CustomDomains), 1, "Second page should return 1 domain")

	// Combine all domains and verify each created one exists
	allDomains := make([]*CustomDomain, 0, len(firstPage.CustomDomains)+len(secondPage.CustomDomains))
	allDomains = append(allDomains, firstPage.CustomDomains...)
	allDomains = append(allDomains, secondPage.CustomDomains...)
	found := map[string]bool{}

	for _, d := range allDomains {
		found[d.GetID()] = true
	}

	assert.True(t, found[domain1.GetID()], "Expected domain1 to be present in paginated results")
	assert.True(t, found[domain2.GetID()], "Expected domain2 to be present in paginated results")
	assert.True(t, found[domain3.GetID()], "Expected domain3 to be present in paginated results")
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
		Domain:         auth0.Stringf("%d.tempdomain.com", time.Now().UTC().UnixNano()),
		Type:           auth0.String("auth0_managed_certs"),
		TLSPolicy:      auth0.String("recommended"),
		DomainMetadata: &map[string]interface{}{"key": "value"},
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
		var mErr Error
		if errors.As(err, &mErr) && mErr.Status() == http.StatusForbidden {
			t.Skip(err)
		}
	}
}
