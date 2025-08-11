package management

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestGuardian(t *testing.T) {
	t.Run("MultiFactor", func(t *testing.T) {
		t.Run("List", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			mfa, err := api.Guardian.MultiFactor.List(context.Background())
			assert.NoError(t, err)
			assert.Greater(t, len(mfa), 1)
		})

		t.Run("Policy", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			initialPolicy, err := api.Guardian.MultiFactor.Policy(context.Background())
			assert.NoError(t, err)

			t.Cleanup(func() {
				err = api.Guardian.MultiFactor.UpdatePolicy(context.Background(), initialPolicy)
				assert.NoError(t, err)
			})

			// Has to be one of "all-applications" or "confidence-score",
			// but not both. If omitted, it removes all policies.
			expectedPolicy := &MultiFactorPolicies{"all-applications"}
			err = api.Guardian.MultiFactor.UpdatePolicy(context.Background(), expectedPolicy)
			assert.NoError(t, err)

			actualPolicy, err := api.Guardian.MultiFactor.Policy(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, expectedPolicy, actualPolicy)
		})

		t.Run("Phone", func(t *testing.T) {
			t.Run("Provider", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialProvider, err := api.Guardian.MultiFactor.Phone.Provider(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err = api.Guardian.MultiFactor.Phone.UpdateProvider(context.Background(), initialProvider)
					assert.NoError(t, err)
				})

				expectedProvider := &MultiFactorProvider{Provider: auth0.String("phone-message-hook")}

				err = api.Guardian.MultiFactor.Phone.UpdateProvider(context.Background(), expectedProvider)
				assert.NoError(t, err)

				actualProvider, err := api.Guardian.MultiFactor.Phone.Provider(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedProvider, actualProvider)
			})

			t.Run("Enable", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialStatus, err := getInitialMFAStatus("sms")
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.Phone.Enable(context.Background(), initialStatus)
					require.NoError(t, err)
				})

				err = api.Guardian.MultiFactor.Phone.Enable(context.Background(), true)
				assert.NoError(t, err)
				assertMFAIsEnabled(t, "sms")
			})

			t.Run("Message-types", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialMessageTypes, err := api.Guardian.MultiFactor.Phone.MessageTypes(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err = api.Guardian.MultiFactor.Phone.UpdateMessageTypes(context.Background(), initialMessageTypes)
					assert.NoError(t, err)
				})

				messageTypes := []string{"voice"}
				expectedPhoneMessageTypes := &PhoneMessageTypes{
					MessageTypes: &messageTypes,
				}

				err = api.Guardian.MultiFactor.Phone.UpdateMessageTypes(context.Background(), expectedPhoneMessageTypes)
				assert.NoError(t, err)

				actualMessageTypes, err := api.Guardian.MultiFactor.Phone.MessageTypes(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedPhoneMessageTypes, actualMessageTypes)
			})
		})

		t.Run("SMS", func(t *testing.T) {
			t.Run("Enable", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialStatus, err := getInitialMFAStatus("sms")
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.SMS.Enable(context.Background(), initialStatus)
					require.NoError(t, err)
				})

				err = api.Guardian.MultiFactor.SMS.Enable(context.Background(), true)
				assert.NoError(t, err)
				assertMFAIsEnabled(t, "sms")
			})

			t.Run("Template", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialTemplate, err := api.Guardian.MultiFactor.SMS.Template(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.SMS.UpdateTemplate(context.Background(), initialTemplate)
					assert.NoError(t, err)
				})

				expectedTemplate := &MultiFactorSMSTemplate{
					EnrollmentMessage:   auth0.String("Test {{code}} for {{tenant.friendly_name}}"),
					VerificationMessage: auth0.String("Test {{code}} for {{tenant.friendly_name}}"),
				}
				err = api.Guardian.MultiFactor.SMS.UpdateTemplate(context.Background(), expectedTemplate)
				assert.NoError(t, err)

				actualTemplate, err := api.Guardian.MultiFactor.SMS.Template(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedTemplate, actualTemplate)
			})

			t.Run("Twilio", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialTwilio, err := api.Guardian.MultiFactor.SMS.Twilio(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.SMS.UpdateTwilio(context.Background(), initialTwilio)
					assert.NoError(t, err)
				})

				expectedTwilio := &MultiFactorProviderTwilio{
					From:      auth0.String("123456789"),
					AuthToken: auth0.String("test_token"),
					SID:       auth0.String("test_sid"),
				}
				err = api.Guardian.MultiFactor.SMS.UpdateTwilio(context.Background(), expectedTwilio)
				assert.NoError(t, err)

				actualTwilio, err := api.Guardian.MultiFactor.SMS.Twilio(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedTwilio, actualTwilio)
			})
		})

		t.Run("Push", func(t *testing.T) {
			t.Run("Provider", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialProvider, err := api.Guardian.MultiFactor.Push.Provider(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err = api.Guardian.MultiFactor.Push.UpdateProvider(context.Background(), initialProvider)
					assert.NoError(t, err)
				})

				expectedProvider := &MultiFactorProvider{Provider: auth0.String("sns")}

				err = api.Guardian.MultiFactor.Push.UpdateProvider(context.Background(), expectedProvider)
				assert.NoError(t, err)

				actualProvider, err := api.Guardian.MultiFactor.Push.Provider(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedProvider, actualProvider)
			})

			t.Run("Enable", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialStatus, err := getInitialMFAStatus("push-notification")
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.Push.Enable(context.Background(), initialStatus)
					require.NoError(t, err)
				})

				err = api.Guardian.MultiFactor.Push.Enable(context.Background(), true)
				assert.NoError(t, err)
				assertMFAIsEnabled(t, "push-notification")
			})

			t.Run("AmazonSNS", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialSNS, err := api.Guardian.MultiFactor.Push.AmazonSNS(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.Push.UpdateAmazonSNS(context.Background(), initialSNS)
					assert.NoError(t, err)
				})

				expectedSNS := &MultiFactorProviderAmazonSNS{
					AccessKeyID:                auth0.String("test"),
					SecretAccessKeyID:          auth0.String("test_secret"),
					Region:                     auth0.String("us-west-1"),
					APNSPlatformApplicationARN: auth0.String("test_arn"),
					GCMPlatformApplicationARN:  auth0.String("test_arn"),
				}
				err = api.Guardian.MultiFactor.Push.UpdateAmazonSNS(context.Background(), expectedSNS)
				assert.NoError(t, err)

				actualSNS, err := api.Guardian.MultiFactor.Push.AmazonSNS(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedSNS.GetAccessKeyID(), actualSNS.GetAccessKeyID())
				assert.Equal(t, expectedSNS.GetRegion(), actualSNS.GetRegion())
				assert.Equal(t, expectedSNS.GetAPNSPlatformApplicationARN(), actualSNS.GetAPNSPlatformApplicationARN())
				assert.Equal(t, expectedSNS.GetGCMPlatformApplicationARN(), actualSNS.GetGCMPlatformApplicationARN())
			})

			t.Run("CustomApp", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialCustomApp, err := api.Guardian.MultiFactor.Push.CustomApp(context.Background())
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.Push.UpdateCustomApp(context.Background(), initialCustomApp)
					assert.NoError(t, err)
				})

				expectedCustomApp := &MultiFactorPushCustomApp{
					AppName:       auth0.String("bestApp"),
					AppleAppLink:  auth0.String("https://itunes.apple.com/us/app/my-app/id123121"),
					GoogleAppLink: auth0.String("https://play.google.com/store/apps/details?id=com.my.app"),
				}
				err = api.Guardian.MultiFactor.Push.UpdateCustomApp(context.Background(), expectedCustomApp)
				assert.NoError(t, err)

				actualCustomApp, err := api.Guardian.MultiFactor.Push.CustomApp(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, expectedCustomApp.GetAppName(), actualCustomApp.GetAppName())
				assert.Equal(t, expectedCustomApp.GetAppleAppLink(), actualCustomApp.GetAppleAppLink())
				assert.Equal(t, expectedCustomApp.GetGoogleAppLink(), actualCustomApp.GetGoogleAppLink())
			})

			t.Run("DirectAPNS", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				_, err := api.Guardian.MultiFactor.Push.DirectAPNS(context.Background())
				assert.NoError(t, err)

				// Cannot reflect back the payload we initially received
				// as it is incompatible

				// Read the p12 from file
				expectedP12, err := os.ReadFile("../test/data/apns.p12")
				assert.NoError(t, err)

				expectedDirectAPNS := &MultiFactorPushDirectAPNS{
					Sandbox:  auth0.Bool(false),
					BundleID: auth0.String("com.my.app"),
					P12:      auth0.String(string(expectedP12)),
				}
				err = api.Guardian.MultiFactor.Push.UpdateDirectAPNS(context.Background(), expectedDirectAPNS)
				assert.NoError(t, err)

				actualDirectAPNS, err := api.Guardian.MultiFactor.Push.DirectAPNS(context.Background())
				// Cannot test `enabled` parameter as we cannot send it due to API limitations.
				assert.NoError(t, err)
				assert.Equal(t, expectedDirectAPNS.GetSandbox(), actualDirectAPNS.GetSandbox())
				assert.Equal(t, expectedDirectAPNS.GetBundleID(), actualDirectAPNS.GetBundleID())
			})

			t.Run("DirectFCM", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				// This is a write only property

				err := error(nil)

				expectedDirectFCM := &MultiFactorPushDirectFCM{
					ServerKey: auth0.String("abc123"),
				}
				err = api.Guardian.MultiFactor.Push.UpdateDirectFCM(context.Background(), expectedDirectFCM)
				assert.NoError(t, err)
			})

			t.Run("DirectFCMv1", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				// This is a write only property

				err := error(nil)

				expectedDirectFCMv1 := &MultiFactorPushDirectFCMv1{
					ServerCredentials: auth0.String("abc123"),
				}
				err = api.Guardian.MultiFactor.Push.UpdateDirectFCMv1(context.Background(), expectedDirectFCMv1)
				assert.NoError(t, err)
			})
		})

		t.Run("Email Enable", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			initialStatus, err := getInitialMFAStatus("email")
			assert.NoError(t, err)

			t.Cleanup(func() {
				err := api.Guardian.MultiFactor.Email.Enable(context.Background(), initialStatus)
				require.NoError(t, err)
			})

			err = api.Guardian.MultiFactor.Email.Enable(context.Background(), true)
			assert.NoError(t, err)
			assertMFAIsEnabled(t, "email")
		})

		t.Run("DUO", func(t *testing.T) {
			t.Run("Enable", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialStatus, err := getInitialMFAStatus("duo")
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.DUO.Enable(context.Background(), initialStatus)
					require.NoError(t, err)
				})

				err = api.Guardian.MultiFactor.DUO.Enable(context.Background(), true)
				assert.NoError(t, err)
				assertMFAIsEnabled(t, "duo")
			})
			t.Run("Settings", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialSettings, err := api.Guardian.MultiFactor.DUO.Read(context.Background())
				assert.NoError(t, err)
				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.DUO.Update(context.Background(), initialSettings)
					require.NoError(t, err)
				})

				updatedSettings := &MultiFactorDUOSettings{
					Hostname:       auth0.String("api-hostname"),
					IntegrationKey: auth0.String("someKey"),
					SecretKey:      auth0.String("someSecret"),
				}
				err = api.Guardian.MultiFactor.DUO.Update(context.Background(), updatedSettings)
				assert.NoError(t, err)

				actualSettings, err := api.Guardian.MultiFactor.DUO.Read(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, "api-hostname", actualSettings.GetHostname())
				assert.Equal(t, "someKey", actualSettings.GetIntegrationKey())
				assert.Equal(t, "someSecret", actualSettings.GetSecretKey())
			})
		})

		t.Run("OTP Enable", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			initialStatus, err := getInitialMFAStatus("otp")
			assert.NoError(t, err)

			t.Cleanup(func() {
				err := api.Guardian.MultiFactor.OTP.Enable(context.Background(), initialStatus)
				require.NoError(t, err)
			})

			err = api.Guardian.MultiFactor.OTP.Enable(context.Background(), true)
			assert.NoError(t, err)
			assertMFAIsEnabled(t, "otp")
		})

		t.Run("WebAuthn Roaming", func(t *testing.T) {
			t.Run("Enable", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialStatus, err := getInitialMFAStatus("webauthn-roaming")
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.WebAuthnRoaming.Enable(context.Background(), initialStatus)
					require.NoError(t, err)
				})

				err = api.Guardian.MultiFactor.WebAuthnRoaming.Enable(context.Background(), true)
				assert.NoError(t, err)
				assertMFAIsEnabled(t, "webauthn-roaming")
			})
			t.Run("Settings", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialSettings, err := api.Guardian.MultiFactor.WebAuthnRoaming.Read(context.Background())
				assert.NoError(t, err)
				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.WebAuthnRoaming.Update(context.Background(), initialSettings)
					require.NoError(t, err)
				})

				updatedSettings := &MultiFactorWebAuthnSettings{
					UserVerification: auth0.String("preferred"),
				}
				err = api.Guardian.MultiFactor.WebAuthnRoaming.Update(context.Background(), updatedSettings)
				assert.NoError(t, err)

				actualSettings, err := api.Guardian.MultiFactor.WebAuthnRoaming.Read(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, "preferred", actualSettings.GetUserVerification())
			})
		})

		t.Run("WebAuthn Platform", func(t *testing.T) {
			t.Run("Enable", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialStatus, err := getInitialMFAStatus("webauthn-platform")
				assert.NoError(t, err)

				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.WebAuthnPlatform.Enable(context.Background(), initialStatus)
					require.NoError(t, err)
				})

				err = api.Guardian.MultiFactor.WebAuthnPlatform.Enable(context.Background(), true)
				assert.NoError(t, err)
				assertMFAIsEnabled(t, "webauthn-platform")
			})
			t.Run("Settings", func(t *testing.T) {
				configureHTTPTestRecordings(t)

				initialSettings, err := api.Guardian.MultiFactor.WebAuthnPlatform.Read(context.Background())
				assert.NoError(t, err)
				t.Cleanup(func() {
					err := api.Guardian.MultiFactor.WebAuthnPlatform.Update(context.Background(), initialSettings)
					require.NoError(t, err)
				})

				updatedSettings := &MultiFactorWebAuthnSettings{
					OverrideRelyingParty: auth0.Bool(false),
				}
				err = api.Guardian.MultiFactor.WebAuthnPlatform.Update(context.Background(), updatedSettings)
				assert.NoError(t, err)

				actualSettings, err := api.Guardian.MultiFactor.WebAuthnPlatform.Read(context.Background())
				assert.NoError(t, err)
				assert.Equal(t, false, actualSettings.GetOverrideRelyingParty())
			})
		})

		t.Run("Recovery Code Enable", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			initialStatus, err := getInitialMFAStatus("recovery-code")
			assert.NoError(t, err)

			t.Cleanup(func() {
				err := api.Guardian.MultiFactor.RecoveryCode.Enable(context.Background(), initialStatus)
				require.NoError(t, err)
			})

			err = api.Guardian.MultiFactor.RecoveryCode.Enable(context.Background(), true)
			assert.NoError(t, err)
			assertMFAIsEnabled(t, "recovery-code")
		})
	})

	t.Run("Enrollment", func(t *testing.T) {
		t.Run("CreateTicket", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			user := givenAUser(t)

			ticket := &CreateEnrollmentTicket{
				UserID:                   user.GetID(),
				Email:                    "jon@example.com",
				SendMail:                 false,
				Factor:                   "push-notification",
				AllowMultipleEnrollments: true,
			}

			createdTicket, err := api.Guardian.Enrollment.CreateTicket(context.Background(), ticket)
			assert.NoError(t, err)
			assert.NotEmpty(t, createdTicket.TicketURL)
			assert.NotEmpty(t, createdTicket.TicketID)
		})

		t.Run("Get", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			_, err := api.Guardian.Enrollment.Get(context.Background(), "dev_0000000000000001")
			// Expect a 404 as we can't set this up through the API.
			assert.Error(t, err)
			assert.Implements(t, (*Error)(nil), err)
			assert.Equal(t, http.StatusNotFound, err.(Error).Status())
		})

		t.Run("Delete", func(t *testing.T) {
			configureHTTPTestRecordings(t)

			err := api.Guardian.Enrollment.Delete(context.Background(), "dev_0000000000000001")
			// Expect a 404 as we can't set this up through the API.
			assert.Error(t, err)
			assert.Implements(t, (*Error)(nil), err)
			assert.Equal(t, http.StatusNotFound, err.(Error).Status())
		})
	})
}

func getInitialMFAStatus(mfaName string) (bool, error) {
	mfaList, err := api.Guardian.MultiFactor.List(context.Background())
	if err != nil {
		return false, err
	}

	enabled := false

	for _, mfa := range mfaList {
		if mfa.GetName() == mfaName {
			enabled = mfa.GetEnabled()
		}
	}

	return enabled, nil
}

func assertMFAIsEnabled(t *testing.T, mfaName string) {
	t.Helper()

	mfaList, err := api.Guardian.MultiFactor.List(context.Background())
	assert.NoError(t, err)

	enabled := false

	for _, mfa := range mfaList {
		if mfa.GetName() == mfaName {
			enabled = mfa.GetEnabled()
		}
	}

	assert.True(t, enabled)
}
