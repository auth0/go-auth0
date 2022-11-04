package management

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestEmailProviderJSON(t *testing.T) {
	var jsonTestCases = []struct {
		name          string
		emailProvider *EmailProvider
		json          string
	}{
		{
			name:          "it can %s an empty string",
			emailProvider: &EmailProvider{},
			json:          `{}`,
		},
		{
			name: "it can %s a mandrill email provider",
			emailProvider: &EmailProvider{
				Name:               auth0.String("mandrill"),
				Enabled:            auth0.Bool(true),
				DefaultFromAddress: auth0.String("accounts@example.com"),
				Credentials: &EmailProviderCredentialsMandrill{
					APIKey: auth0.String("secretApiKey"),
				},
				Settings: &EmailProviderSettingsMandrill{
					Message: &EmailProviderSettingsMandrillMessage{
						ViewContentLink: auth0.Bool(true),
					},
				},
			},
			json: `{"name":"mandrill","enabled":true,"default_from_address":"accounts@example.com","credentials":{"api_key":"secretApiKey"},"settings":{"message":{"view_content_link":true}}}`,
		},
		{
			name: "it can %s an ses email provider",
			emailProvider: &EmailProvider{
				Name:               auth0.String("ses"),
				Enabled:            auth0.Bool(true),
				DefaultFromAddress: auth0.String("accounts@example.com"),
				Credentials: &EmailProviderCredentialsSES{
					AccessKeyID:     auth0.String("accessKey"),
					SecretAccessKey: auth0.String("secret"),
					Region:          auth0.String("eu-west-2"),
				},
				Settings: &EmailProviderSettingsSES{
					Message: &EmailProviderSettingsSESMessage{
						ConfigurationSetName: auth0.String("example"),
					},
				},
			},
			json: `{"name":"ses","enabled":true,"default_from_address":"accounts@example.com","credentials":{"accessKeyId":"accessKey","secretAccessKey":"secret","region":"eu-west-2"},"settings":{"message":{"configuration_set_name":"example"}}}`,
		},
		{
			name: "it can %s a sendgrid email provider",
			emailProvider: &EmailProvider{
				Name:               auth0.String("sendgrid"),
				Enabled:            auth0.Bool(true),
				DefaultFromAddress: auth0.String("accounts@example.com"),
				Credentials: &EmailProviderCredentialsSendGrid{
					APIKey: auth0.String("apiKey"),
				},
			},
			json: `{"name":"sendgrid","enabled":true,"default_from_address":"accounts@example.com","credentials":{"api_key":"apiKey"}}`,
		},
		{
			name: "it can %s a sparkpost email provider",
			emailProvider: &EmailProvider{
				Name:               auth0.String("sparkpost"),
				Enabled:            auth0.Bool(true),
				DefaultFromAddress: auth0.String("accounts@example.com"),
				Credentials: &EmailProviderCredentialsSparkPost{
					APIKey: auth0.String("apiKey"),
					Region: auth0.String("eu"),
				},
			},
			json: `{"name":"sparkpost","enabled":true,"default_from_address":"accounts@example.com","credentials":{"api_key":"apiKey","region":"eu"}}`,
		},
		{
			name: "it can %s a mailgun email provider",
			emailProvider: &EmailProvider{
				Name:               auth0.String("mailgun"),
				Enabled:            auth0.Bool(true),
				DefaultFromAddress: auth0.String("accounts@example.com"),
				Credentials: &EmailProviderCredentialsMailgun{
					APIKey: auth0.String("apiKey"),
					Region: auth0.String("eu"),
					Domain: auth0.String("example.com"),
				},
			},
			json: `{"name":"mailgun","enabled":true,"default_from_address":"accounts@example.com","credentials":{"api_key":"apiKey","domain":"example.com","region":"eu"}}`,
		},
		{
			name: "it can %s an smtp email provider",
			emailProvider: &EmailProvider{
				Name:               auth0.String("smtp"),
				Enabled:            auth0.Bool(true),
				DefaultFromAddress: auth0.String("accounts@example.com"),
				Credentials: &EmailProviderCredentialsSmtp{
					SMTPHost: auth0.String("example.com"),
					SMTPPort: auth0.Int(3000),
					SMTPUser: auth0.String("user"),
					SMTPPass: auth0.String("pass"),
				},
				Settings: &EmailProviderSettingsSmtp{
					Headers: &EmailProviderSettingsSmtpHeaders{
						XMCViewContentLink:   auth0.String("true"),
						XSESConfigurationSet: auth0.String("example"),
					},
				},
			},
			json: `{"name":"smtp","enabled":true,"default_from_address":"accounts@example.com","credentials":{"smtp_host":"example.com","smtp_port":3000,"smtp_user":"user","smtp_pass":"pass"},"settings":{"headers":{"X-MC-ViewContentLink":"true","X-SES-Configuration-Set":"example"}}}`,
		},
	}

	for _, testCase := range jsonTestCases {
		t.Run(fmt.Sprintf(testCase.name, "marshal"), func(t *testing.T) {
			actualJSON, err := json.Marshal(testCase.emailProvider)
			assert.NoError(t, err)
			assert.Equal(t, testCase.json, string(actualJSON))
		})
	}

	for _, testCase := range jsonTestCases {
		t.Run(fmt.Sprintf(testCase.name, "unmarshal"), func(t *testing.T) {
			var actualEmailProvider EmailProvider
			err := json.Unmarshal([]byte(testCase.json), &actualEmailProvider)
			assert.NoError(t, err)
			assert.Equal(t, testCase.emailProvider, &actualEmailProvider)
		})
	}
}

func TestEmailProviderManager_Create(t *testing.T) {
	setupHTTPRecordings(t)

	emailProvider := &EmailProvider{
		Name:               auth0.String("smtp"),
		Enabled:            auth0.Bool(true),
		DefaultFromAddress: auth0.String("no-reply@example.com"),
		Credentials: &EmailProviderCredentialsSmtp{
			SMTPHost: auth0.String("smtp.example.com"),
			SMTPPort: auth0.Int(587),
			SMTPUser: auth0.String("user"),
			SMTPPass: auth0.String("pass"),
		},
	}

	err := m.EmailProvider.Create(emailProvider)
	assert.NoError(t, err)

	t.Cleanup(func() {
		cleanupEmailProvider(t)
	})
}

func TestEmailProviderManager_Read(t *testing.T) {
	setupHTTPRecordings(t)

	expectedEmailProvider := givenAnEmailProvider(t)

	actualEmailProvider, err := m.EmailProvider.Read()

	assert.NoError(t, err)
	assert.Equal(t, expectedEmailProvider.GetName(), actualEmailProvider.GetName())
	assert.Equal(t, expectedEmailProvider.GetEnabled(), actualEmailProvider.GetEnabled())
	assert.Equal(t, expectedEmailProvider.GetDefaultFromAddress(), actualEmailProvider.GetDefaultFromAddress())
	assert.IsType(t, &EmailProviderCredentialsSmtp{}, expectedEmailProvider.Credentials)
}

func TestEmailProviderManager_Update(t *testing.T) {
	setupHTTPRecordings(t)

	emailProvider := givenAnEmailProvider(t)

	emailProvider.Enabled = auth0.Bool(false)
	emailProvider.DefaultFromAddress = auth0.String("info@example.com")

	err := m.EmailProvider.Update(emailProvider)
	assert.NoError(t, err)

	actualEmailProvider, err := m.EmailProvider.Read()
	assert.NoError(t, err)

	assert.False(t, actualEmailProvider.GetEnabled())
	assert.Equal(t, "info@example.com", actualEmailProvider.GetDefaultFromAddress())
}

func TestEmailProviderManager_Delete(t *testing.T) {
	setupHTTPRecordings(t)

	givenAnEmailProvider(t)

	err := m.Email.Delete()
	assert.NoError(t, err)

	_, err = m.Email.Read()
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenAnEmailProvider(t *testing.T) *EmailProvider {
	t.Helper()

	emailProvider := &EmailProvider{
		Name:               auth0.String("smtp"),
		Enabled:            auth0.Bool(true),
		DefaultFromAddress: auth0.String("no-reply@example.com"),
		Credentials: &EmailProviderCredentialsSmtp{
			SMTPHost: auth0.String("smtp.example.com"),
			SMTPPort: auth0.Int(587),
			SMTPUser: auth0.String("user"),
			SMTPPass: auth0.String("pass"),
		},
	}

	err := m.EmailProvider.Create(emailProvider)
	if err != nil {
		if err.(Error).Status() != http.StatusConflict {
			t.Error(err)
		}
	}

	t.Cleanup(func() {
		cleanupEmailProvider(t)
	})

	return emailProvider
}

func cleanupEmailProvider(t *testing.T) {
	t.Helper()

	err := m.EmailProvider.Delete()
	require.NoError(t, err)
}
