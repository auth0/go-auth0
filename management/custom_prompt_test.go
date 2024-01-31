package management

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCustomPromptManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	prompt := CustomPromptSignup
	expected := &CustomPrompt{Prompt: prompt}
	got, err := api.CustomPrompt.Read(context.Background(), prompt)

	assertNoCustomPromptError(t, err)
	assert.Equal(t, expected, got)

	cleanupCustomDomain(t, customDomain.GetID())
}

func TestCustomPromptManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	prompt := CustomPromptSignup
	original, err := api.CustomPrompt.Read(context.Background(), prompt)
	assertNoCustomPromptError(t, err)

	expected := &CustomPrompt{Prompt: prompt, FormContentStart: `<div>Test Content</div>`}

	err = api.CustomPrompt.Create(context.Background(), expected)
	assertNoCustomPromptError(t, err)

	got, err := api.CustomPrompt.Read(context.Background(), prompt)
	assertNoCustomPromptError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, got)

	t.Cleanup(func() {
		cleanupCustomPrompt(t, customDomain.GetID(), prompt)
	})
}

func TestCustomPromptManager_Update(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	prompt := CustomPromptSignup
	original := &CustomPrompt{Prompt: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.CustomPrompt.Create(context.Background(), original)
	assertNoCustomPromptError(t, err)

	expected := &CustomPrompt{Prompt: prompt, FormContentStart: `<div>Updated Test Content</div>`}
	err = api.CustomPrompt.Update(context.Background(), expected)
	assertNoCustomPromptError(t, err)

	got, err := api.CustomPrompt.Read(context.Background(), prompt)
	assertNoCustomPromptError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, expected)

	t.Cleanup(func() {
		cleanupCustomPrompt(t, customDomain.GetID(), prompt)
	})
}

func TestCustomPromptManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)

	customDomain := givenACustomDomain(t)
	prompt := CustomPromptSignup
	original := &CustomPrompt{Prompt: prompt, FormContentStart: `<div>Test Content</div>`}

	err := api.CustomPrompt.Create(context.Background(), original)
	assertNoCustomPromptError(t, err)

	expected := &CustomPrompt{Prompt: prompt}
	err = api.CustomPrompt.Delete(context.Background(), expected)
	assertNoCustomPromptError(t, err)

	got, err := api.CustomPrompt.Read(context.Background(), prompt)
	assertNoCustomPromptError(t, err)

	assert.Equal(t, expected, got)
	assert.NotEqual(t, original, expected)

	t.Cleanup(func() {
		cleanupCustomPrompt(t, customDomain.GetID(), prompt)
	})
}

func cleanupCustomPrompt(t *testing.T, customDomainID string, prompt CustomPromptType) {
	t.Helper()

	c := &CustomPrompt{Prompt: prompt}
	err := api.CustomPrompt.Delete(context.Background(), c)
	assertNoCustomPromptError(t, err)

	cleanupCustomDomain(t, customDomainID)
}

func assertNoCustomPromptError(t *testing.T, err error) {
	if err != nil {
		var mErr Error
		if errors.As(err, &mErr) && mErr.Status() == http.StatusForbidden {
			t.Skip(err)
			return
		}
		t.Error(err)
	}
}
