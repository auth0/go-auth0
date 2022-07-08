package management

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestEmailTemplateManager_Create(t *testing.T) {
	setupHTTPRecordings(t)

	template := &EmailTemplate{
		Template:               auth0.String("verify_email"),
		Body:                   auth0.String("<html><body><h1>Verify your email</h1></body></html>"),
		From:                   auth0.String("me@example.com"),
		ResultURL:              auth0.String("https://www.example.com/verify-email"),
		Subject:                auth0.String("Verify your email"),
		Syntax:                 auth0.String("liquid"),
		Enabled:                auth0.Bool(true),
		IncludeEmailInRedirect: auth0.Bool(true),
	}

	err := m.EmailTemplate.Create(template)
	if err != nil {
		if err, ok := err.(Error); ok && err.Status() != http.StatusConflict {
			t.Error(err)
		}
	}

	t.Cleanup(func() {
		cleanupEmailTemplate(t, template.GetTemplate())
	})
}

func TestEmailTemplateManager_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expectedTemplate := givenAnEmailTemplate(t)

	actualTemplate, err := m.EmailTemplate.Read(expectedTemplate.GetTemplate())

	assert.NoError(t, err)
	assert.ObjectsAreEqual(expectedTemplate, actualTemplate)
}

func TestEmailTemplateManager_Update(t *testing.T) {
	setupHTTPRecordings(t)

	template := givenAnEmailTemplate(t)

	expectedBody := "<html><body><h1>Let's get you verified!</h1></body></html>"
	expectedIncludeEmailInRedirect := false
	err := m.EmailTemplate.Update(
		template.GetTemplate(),
		&EmailTemplate{
			Body:                   &expectedBody,
			IncludeEmailInRedirect: &expectedIncludeEmailInRedirect,
		},
	)
	assert.NoError(t, err)

	actualTemplate, err := m.EmailTemplate.Read(template.GetTemplate())
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, actualTemplate.GetBody())
	assert.Equal(t, expectedIncludeEmailInRedirect, actualTemplate.GetIncludeEmailInRedirect())
}

func TestEmailTemplateManager_Replace(t *testing.T) {
	setupHTTPRecordings(t)

	givenAnEmailProvider(t)
	template := givenAnEmailTemplate(t)

	template.Template = auth0.String("verify_email")
	template.Subject = auth0.String("Let's get you verified!")
	template.Body = auth0.String("<html><body><h1>Let's get you verified!</h1></body></html>")
	template.From = auth0.String("someone@example.com")
	template.IncludeEmailInRedirect = auth0.Bool(true)

	err := m.EmailTemplate.Replace(template.GetTemplate(), template)
	assert.NoError(t, err)

	actualTemplate, err := m.EmailTemplate.Read(template.GetTemplate())
	assert.NoError(t, err)

	assert.Equal(t, actualTemplate.GetBody(), template.GetBody())
	assert.Equal(t, actualTemplate.GetSubject(), template.GetSubject())
	assert.Equal(t, actualTemplate.GetFrom(), template.GetFrom())
	assert.Equal(t, actualTemplate.GetTemplate(), template.GetTemplate())
	assert.Equal(t, actualTemplate.GetIncludeEmailInRedirect(), template.GetIncludeEmailInRedirect())
}

func givenAnEmailTemplate(t *testing.T) *EmailTemplate {
	t.Helper()

	template := &EmailTemplate{
		Template:               auth0.String("verify_email"),
		Body:                   auth0.String("<html><body><h1>Verify your email</h1></body></html>"),
		From:                   auth0.String("me@example.com"),
		ResultURL:              auth0.String("https://www.example.com/verify-email"),
		Subject:                auth0.String("Verify your email"),
		Syntax:                 auth0.String("liquid"),
		Enabled:                auth0.Bool(true),
		IncludeEmailInRedirect: auth0.Bool(true),
	}

	err := m.EmailTemplate.Create(template)
	if err != nil {
		if err, ok := err.(Error); ok && err.Status() != http.StatusConflict {
			t.Error(err)
		}
	}

	t.Cleanup(func() {
		cleanupEmailTemplate(t, template.GetTemplate())
	})

	return template
}

func cleanupEmailTemplate(t *testing.T, templateName string) {
	t.Helper()

	err := m.EmailTemplate.Update(templateName, &EmailTemplate{Enabled: auth0.Bool(false)})
	require.NoError(t, err)
}
