package guardian

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/auth0/go-auth0/management/models"
)

// Manager defines the base manager interface
type Manager struct {
	management *management.Management
}

// NewManager creates a new manager for  operations
func NewManager(mgmt *management.Management) *Manager {
	return &Manager{
		management: mgmt,
	}
}

// DeleteGuardianEnrollment Delete a multi-factor authentication enrollment
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/delete_enrollments_by_id
func (m *Manager) DeleteGuardianEnrollment(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("guardian", "enrollments", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetPushNotificationProviderAPNS Get APNS push notification configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_apns
func (m *Manager) GetPushNotificationProviderAPNS(ctx context.Context, opts ...management.RequestOption) (*models.GetApns200Response, error) {
	var localVarReturnValue *models.GetApns200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "push-notification", "providers", "apns"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetGuardianEnrollment Get a multi-factor authentication enrollment
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_enrollments_by_id
func (m *Manager) GetGuardianEnrollment(ctx context.Context, id string, opts ...management.RequestOption) (*models.Enrollment, error) {
	var localVarReturnValue *models.Enrollment
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "enrollments", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPhoneFactorTemplates Get Enrollment and Verification Phone Templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_factor_phone_templates
func (m *Manager) GetPhoneFactorTemplates(ctx context.Context, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "templates"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSmsFactorTemplates Get SMS enrollment and verification templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_factor_sms_templates
func (m *Manager) GetSmsFactorTemplates(ctx context.Context, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "sms", "templates"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetFactors Get Factors and multi-factor authentication details
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_factors
func (m *Manager) GetFactors(ctx context.Context, opts ...management.RequestOption) ([]*models.Factor, error) {
	var localVarReturnValue []*models.Factor
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPhoneFactorSelectedProvider Get phone provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_guardian_phone_providers
func (m *Manager) GetPhoneFactorSelectedProvider(ctx context.Context, opts ...management.RequestOption) (*models.GetGuardianPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetGuardianPhoneProviders200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "selected-provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPhoneFactorMessageTypes Get Enabled Phone Factors
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_message_types
func (m *Manager) GetPhoneFactorMessageTypes(ctx context.Context, opts ...management.RequestOption) (*models.GetMessageTypes200Response, error) {
	var localVarReturnValue *models.GetMessageTypes200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "message-types"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPhoneFactorProviderTwilio Get Twilio configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_phone_twilio_factor_provider
func (m *Manager) GetPhoneFactorProviderTwilio(ctx context.Context, opts ...management.RequestOption) (*models.TwilioFactorProvider, error) {
	var localVarReturnValue *models.TwilioFactorProvider
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "providers", "twilio"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPushNotificationSelectedProvider Get push notification provider
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_pn_providers
func (m *Manager) GetPushNotificationSelectedProvider(ctx context.Context, opts ...management.RequestOption) (*models.GetPnProviders200Response, error) {
	var localVarReturnValue *models.GetPnProviders200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "push-notification", "selected-provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPolicies Get multi-factor authentication policies
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_policies
func (m *Manager) GetPolicies(ctx context.Context, opts ...management.RequestOption) ([]*models.GetPolicies200ResponseInner, error) {
	var localVarReturnValue []*models.GetPolicies200ResponseInner
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "policies"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSmsSelectedProvider Get SMS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_sms_providers
func (m *Manager) GetSmsSelectedProvider(ctx context.Context, opts ...management.RequestOption) (*models.GetGuardianPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetGuardianPhoneProviders200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "sms", "selected-provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSmsFactorProviderTwilio Get Twilio SMS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_sms_twilio_factor_provider
func (m *Manager) GetSmsFactorProviderTwilio(ctx context.Context, opts ...management.RequestOption) (*models.SmsTwilioFactorProvider, error) {
	var localVarReturnValue *models.SmsTwilioFactorProvider
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "sms", "providers", "twilio"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPushNotificationProviderSNS Get AWS SNS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_sns
func (m *Manager) GetPushNotificationProviderSNS(ctx context.Context, opts ...management.RequestOption) (*models.SnsFactorProvider, error) {
	var localVarReturnValue *models.SnsFactorProvider
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "push-notification", "providers", "sns"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdatePushNotificationProviderAPNS Update APNs provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/patch_apns
func (m *Manager) UpdatePushNotificationProviderAPNS(ctx context.Context, putApnsRequest *models.PutApnsRequest, opts ...management.RequestOption) (*models.PutApns200Response, error) {
	var localVarReturnValue *models.PutApns200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("guardian", "factors", "push-notification", "providers", "apns"), putApnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdatePushNotificationProviderFCM Updates FCM configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/patch_fcm
func (m *Manager) UpdatePushNotificationProviderFCM(ctx context.Context, putFcmRequest *models.PutFcmRequest, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "PATCH", m.management.URI("guardian", "factors", "push-notification", "providers", "fcm"), putFcmRequest, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// UpdatePushNotificationProviderSNS Update AWS SNS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/patch_sns
func (m *Manager) UpdatePushNotificationProviderSNS(ctx context.Context, putSnsRequest *models.PutSnsRequest, opts ...management.RequestOption) (*models.PutSnsRequest, error) {
	var localVarReturnValue *models.PutSnsRequest
	err := m.management.Request(ctx, "PATCH", m.management.URI("guardian", "factors", "push-notification", "providers", "sns"), putSnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// CreateEnrollmentTicket Create a multi-factor authentication enrollment ticket
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/post_ticket
func (m *Manager) CreateEnrollmentTicket(ctx context.Context, enrollmentCreate *models.EnrollmentCreate, opts ...management.RequestOption) (*models.PostTicket200Response, error) {
	var localVarReturnValue *models.PostTicket200Response
	err := m.management.Request(ctx, "POST", m.management.URI("guardian", "enrollments", "ticket"), enrollmentCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetPushNotificationProviderAPNS Update APNS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_apns
func (m *Manager) SetPushNotificationProviderAPNS(ctx context.Context, putApnsRequest *models.PutApnsRequest, opts ...management.RequestOption) (*models.PutApns200Response, error) {
	var localVarReturnValue *models.PutApns200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "providers", "apns"), putApnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetPhoneFactorTemplates Update Enrollment and Verification Phone Templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_factor_phone_templates
func (m *Manager) SetPhoneFactorTemplates(ctx context.Context, templateMessages *models.TemplateMessages, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "templates"), templateMessages, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetSmsFactorTemplates Update SMS enrollment and verification templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_factor_sms_templates
func (m *Manager) SetSmsFactorTemplates(ctx context.Context, templateMessages *models.TemplateMessages, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "sms", "templates"), templateMessages, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdateFactor Update multi-factor authentication type
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_factors_by_name
func (m *Manager) UpdateFactor(ctx context.Context, name models.PutFactorsByNameNameParameter, putFactorsByNameRequest *models.PutFactorsByNameRequest, opts ...management.RequestOption) (*models.PutFactorsByName200Response, error) {
	var localVarReturnValue *models.PutFactorsByName200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", string(name)), putFactorsByNameRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetPushNotificationProviderFCM Updates FCM configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_fcm
func (m *Manager) SetPushNotificationProviderFCM(ctx context.Context, putFcmRequest *models.PutFcmRequest, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "providers", "fcm"), putFcmRequest, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// UpdatePhoneFactorMessageTypes Update the Enabled Phone Factors
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_message_types
func (m *Manager) UpdatePhoneFactorMessageTypes(ctx context.Context, putMessageTypesRequest *models.PutMessageTypesRequest, opts ...management.RequestOption) (*models.GetMessageTypes200Response, error) {
	var localVarReturnValue *models.GetMessageTypes200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "message-types"), putMessageTypesRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdatePhoneFactorSelectedProvider Update phone provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_phone_providers
func (m *Manager) UpdatePhoneFactorSelectedProvider(ctx context.Context, putPhoneProvidersRequest *models.PutPhoneProvidersRequest, opts ...management.RequestOption) (*models.GetGuardianPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetGuardianPhoneProviders200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "selected-provider"), putPhoneProvidersRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetPushNotificationSelectedProvider Update Push Notification configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_pn_providers
func (m *Manager) SetPushNotificationSelectedProvider(ctx context.Context, putPnProvidersRequest *models.PutPnProvidersRequest, opts ...management.RequestOption) (*models.GetPnProviders200Response, error) {
	var localVarReturnValue *models.GetPnProviders200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "selected-provider"), putPnProvidersRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdatePolicies Update multi-factor authentication policies
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_policies
func (m *Manager) UpdatePolicies(ctx context.Context, getPolicies200ResponseInner []*models.GetPolicies200ResponseInner, opts ...management.RequestOption) ([]*models.GetPolicies200ResponseInner, error) {
	var localVarReturnValue []*models.GetPolicies200ResponseInner
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "policies"), getPolicies200ResponseInner, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetSmsSelectedProvider Update SMS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_sms_providers
func (m *Manager) SetSmsSelectedProvider(ctx context.Context, putPhoneProvidersRequest *models.PutPhoneProvidersRequest, opts ...management.RequestOption) (*models.GetGuardianPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetGuardianPhoneProviders200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "sms", "selected-provider"), putPhoneProvidersRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetSmsFactorProviderTwilio Update Twilio SMS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_sms_twilio_factor_provider
func (m *Manager) SetSmsFactorProviderTwilio(ctx context.Context, putTwilioRequest *models.PutTwilioRequest, opts ...management.RequestOption) (*models.SmsTwilioFactorProvider, error) {
	var localVarReturnValue *models.SmsTwilioFactorProvider
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "sms", "providers", "twilio"), putTwilioRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// SetPushNotificationProviderSNS Update AWS SNS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_sns
func (m *Manager) SetPushNotificationProviderSNS(ctx context.Context, putSnsRequest *models.PutSnsRequest, opts ...management.RequestOption) (*models.PutSnsRequest, error) {
	var localVarReturnValue *models.PutSnsRequest
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "providers", "sns"), putSnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// UpdatePhoneFactorProviderTwilio Update Twilio configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_twilio
func (m *Manager) UpdatePhoneFactorProviderTwilio(ctx context.Context, putTwilioRequest *models.PutTwilioRequest, opts ...management.RequestOption) (*models.TwilioFactorProvider, error) {
	var localVarReturnValue *models.TwilioFactorProvider
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "providers", "twilio"), putTwilioRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
