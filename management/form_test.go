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
		Name: auth0.String("test-form"),
		Languages: &FormLanguages{
			Primary: auth0.String("en"),
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
		Name: auth0.String("updated-test-form"),
	}
	err := api.Form.Update(context.Background(), expectedForm.GetID(), updatedForm)

	assert.NoError(t, err)
	assert.Equal(t, "updated-test-form", updatedForm.GetName())
	assert.Equal(t, expectedForm.GetLanguages(), updatedForm.GetLanguages())
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
	form.Ending = nil
	form.Messages = nil
	form.Languages = nil
	form.Nodes = nil
	form.Style = nil
	form.Start = nil
	form.Translations = nil

	formList, err := api.Form.List(context.Background())
	assert.NoError(t, err)
	assert.Greater(t, len(formList.Forms), 0)
	assert.Contains(t, formList.Forms, form)
}

func TestFormManager_MarshalJSON(t *testing.T) {
	for form, expected := range map[*Form]string{
		{}: `{}`,
		{
			Name: auth0.String("test-form"),
			Languages: &FormLanguages{
				Primary: auth0.String("en"),
			},
		}: `{"name":"test-form","languages":{"primary":"en"}}`,
		{
			Messages: &FormMessages{
				Custom: &map[string]interface{}{
					"welcome": "Welcome to the form",
				},
				Errors: &map[string]interface{}{
					"required": "This field is required",
				},
			},
		}: `{"messages":{"custom":{"welcome":"Welcome to the form"},"errors":{"required":"This field is required"}}}`,
		{
			ID:        auth0.String("some-id"),
			CreatedAt: auth0.Time(time.Now()),
			UpdatedAt: auth0.Time(time.Now()),
		}: `{}`,
	} {
		payload, err := json.Marshal(form)
		assert.NoError(t, err)
		assert.Equal(t, expected, string(payload))
	}
}

func givenAForm(t *testing.T) *Form {
	t.Helper()

	form := &Form{
		Name: auth0.String("test-form"),
		Languages: &FormLanguages{
			Primary: auth0.String("en"),
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
