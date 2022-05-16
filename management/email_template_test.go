package management

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestEmailTemplateManager_Create(t *testing.T) {
	template := &EmailTemplate{
		Template:  auth0.String("verify_email"),
		Body:      auth0.String("<html><body><h1>Verify your email</h1></body></html>"),
		From:      auth0.String("me@example.com"),
		ResultURL: auth0.String("https://www.example.com/verify-email"),
		Subject:   auth0.String("Verify your email"),
		Syntax:    auth0.String("liquid"),
		Enabled:   auth0.Bool(true),
	}

	err := m.EmailTemplate.Create(template)
	if err != nil {
		if err, ok := err.(Error); ok && err.Status() != http.StatusConflict {
			t.Error(err)
		}
	}

	defer cleanupEmailTemplate(t, template.GetTemplate())
}

func TestEmailTemplateManager_Read(t *testing.T) {
	expectedTemplate := givenAnEmailTemplate(t)
	defer cleanupEmailTemplate(t, expectedTemplate.GetTemplate())

	actualTemplate, err := m.EmailTemplate.Read(expectedTemplate.GetTemplate())

	assert.NoError(t, err)
	assert.ObjectsAreEqual(expectedTemplate, actualTemplate)
}

func TestEmailTemplateManager_Update(t *testing.T) {
	template := givenAnEmailTemplate(t)
	defer cleanupEmailTemplate(t, template.GetTemplate())

	expectedBody := "<html><body><h1>Let's get you verified!</h1></body></html>"
	err := m.EmailTemplate.Update(
		template.GetTemplate(),
		&EmailTemplate{
			Body: &expectedBody,
		},
	)
	assert.NoError(t, err)

	actualTemplate, err := m.EmailTemplate.Read(template.GetTemplate())
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, actualTemplate.GetBody())
}

func TestEmailTemplateManager_Replace(t *testing.T) {
	givenAnEmailProvider(t)
	defer cleanupEmailProvider(t)

	template := givenAnEmailTemplate(t)
	defer cleanupEmailTemplate(t, template.GetTemplate())

	template.Subject = auth0.String("Let's get you verified!")
	template.Body = auth0.String("<html><body><h1>Let's get you verified!</h1></body></html>")
	template.From = auth0.String("someone@example.com")

	err := m.EmailTemplate.Replace(template.GetTemplate(), template)
	assert.NoError(t, err)
}

func givenAnEmailTemplate(t *testing.T) *EmailTemplate {
	t.Helper()

	template := &EmailTemplate{
		Template:  auth0.String("verify_email"),
		Body:      auth0.String("<html><body><h1>Verify your email</h1></body></html>"),
		From:      auth0.String("me@example.com"),
		ResultURL: auth0.String("https://www.example.com/verify-email"),
		Subject:   auth0.String("Verify your email"),
		Syntax:    auth0.String("liquid"),
		Enabled:   auth0.Bool(true),
	}

	err := m.EmailTemplate.Create(template)
	if err != nil {
		if err, ok := err.(Error); ok && err.Status() != http.StatusConflict {
			t.Error(err)
		}
	}

	return template
}

func cleanupEmailTemplate(t *testing.T, templateName string) {
	t.Helper()

	err := m.EmailTemplate.Update(templateName, &EmailTemplate{Enabled: auth0.Bool(false)})
	require.NoError(t, err)
}
