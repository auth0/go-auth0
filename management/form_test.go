package management

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestFormManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)
	form := &Form{
		Name:        auth0.String("test-form"),
		Description: auth0.String("A test form"),
		Languages: &FormLanguages{
			Primary: auth0.String("en"),
		},
		Style: &FormStyle{
			Theme:   auth0.String("SOFT"),
			Version: auth0.String("MODERN"),
		},
	}

	err := api.Form.Create(context.Background(), form)

	assert.NoError(t, err)
	assert.NotEmpty(t, form.GetID())

	t.Cleanup(func() {
		cleanupForm(t, form.GetID())
	})
}

func TestFormManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedForm := givenAForm(t)

	actualForm, err := api.Form.Read(context.Background(), expectedForm.GetID())

	assert.NoError(t, err)
	assert.Equal(t, expectedForm, actualForm)
}

func TestFormManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedForm := givenAForm(t)
	updatedForm := &Form{
		Description: auth0.String("Updated Description test"),
	}
	err := api.Form.Update(context.Background(), expectedForm.GetID(), updatedForm)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Description test", updatedForm.GetDescription())
	assert.Equal(t, expectedForm.GetName(), updatedForm.GetName())
}

func TestFormManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	expectedForm := givenAForm(t)

	err := api.Form.Delete(context.Background(), expectedForm.GetID())
	assert.NoError(t, err)
}

func TestFormManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	form := givenAForm(t)

	formList, err := api.Form.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(formList.Forms), 0)
	assert.Contains(t, getFormIDs(formList.Forms), form.GetID())
}

func TestFormManager_MarshalJSON(t *testing.T) {
	for form, expected := range map[*Form]string{
		{}: `{}`,
		{
			Name:        auth0.String("test-form"),
			Description: auth0.String("A test form"),
			Languages: &FormLanguages{
				Primary: auth0.String("en"),
			},
			Style: &FormStyle{
				Theme: auth0.String("SOFT"),
			},
		}: `{"name":"test-form","description":"A test form","languages":{"primary":"en"},"style":{"theme":"SOFT"}}`,
		{
			Messages: &FormMessages{
				Custom: map[string]interface{}{
					"welcome": "Welcome to the form",
				},
				Errors: map[string]interface{}{
					"required": "This field is required",
				},
			},
		}: `{"messages":{"custom":{"welcome":"Welcome to the form"},"errors":{"required":"This field is required"}}}`,
		{
			ID:        auth0.String("some-id"),
			FlowCount: auth0.Int(1),
			CreatedAt: auth0.Time(time.Now()),
			UpdatedAt: auth0.Time(time.Now()),
		}: `{}`,
	} {
		payload, err := json.Marshal(form)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func getFormIDs(forms []*Form) []string {
	ids := make([]string, len(forms))
	for i, f := range forms {
		ids[i] = f.GetID()
	}
	return ids
}

func givenAForm(t *testing.T) *Form {
	t.Helper()
	form := &Form{
		Name:        auth0.String("test-form"),
		Description: auth0.String("A test form"),
		Languages: &FormLanguages{
			Primary: auth0.String("en"),
		},
		Style: &FormStyle{
			Theme:   auth0.String("SOFT"),
			Version: auth0.String("MODERN"),
		},
	}

	err := api.Form.Create(context.Background(), form)
	assert.NoError(t, err)
	t.Cleanup(func() {
		cleanupForm(t, form.GetID())
	})
	return form
}

func cleanupForm(t *testing.T, formID string) {
	t.Helper()

	err := api.Form.Delete(context.Background(), formID)
	if err != nil {
		var managementErr Error
		ok := errors.As(err, &managementErr)
		// We don't want to fail the test if the resource is already deleted.
		// clean up, therefore we only raise non-404 errors.
		// If `err` doesn't cast to management.Error, we raise it immediately.
		if !ok || managementErr.Status() != http.StatusNotFound {
			t.Error(err)
		}
	}
}
