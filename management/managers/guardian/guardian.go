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

// DeleteEnrollmentsById Delete a multi-factor authentication enrollment
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/delete_enrollments_by_id
func (m *Manager) DeleteEnrollmentsById(ctx context.Context, id string, opts ...management.RequestOption) error {

	err := m.management.Request(ctx, "DELETE", m.management.URI("guardian", "enrollments", string(id)), nil, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetApns Retrieve APNS push notification configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_apns
func (m *Manager) GetApns(ctx context.Context, opts ...management.RequestOption) (*models.GetApns200Response, error) {
	var localVarReturnValue *models.GetApns200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "push-notification", "providers", "apns"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetEnrollmentsById Retrieve a multi-factor authentication enrollment
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_enrollments_by_id
func (m *Manager) GetEnrollmentsById(ctx context.Context, id string, opts ...management.RequestOption) (*models.Enrollment, error) {
	var localVarReturnValue *models.Enrollment
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "enrollments", string(id)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetFactorPhoneTemplates Retrieve Enrollment and Verification Phone Templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_factor_phone_templates
func (m *Manager) GetFactorPhoneTemplates(ctx context.Context, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "templates"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetFactorSmsTemplates Retrieve SMS Enrollment and Verification Templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_factor_sms_templates
func (m *Manager) GetFactorSmsTemplates(ctx context.Context, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "sms", "templates"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetFactors Retrieve Factors and their Status
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

// GetMessageTypes Retrieve the Enabled Phone Factors
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_message_types
func (m *Manager) GetMessageTypes(ctx context.Context, opts ...management.RequestOption) (*models.GetMessageTypes200Response, error) {
	var localVarReturnValue *models.GetMessageTypes200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "message-types"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPhoneProviders Retrieve phone configuration (one of auth0|twilio|phone-message-hook)
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_phone_providers
func (m *Manager) GetPhoneProviders(ctx context.Context, opts ...management.RequestOption) (*models.GetPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetPhoneProviders200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "selected-provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPhoneTwilioFactorProvider Retrieve Twilio phone configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_phone_twilio_factor_provider
func (m *Manager) GetPhoneTwilioFactorProvider(ctx context.Context, opts ...management.RequestOption) (*models.TwilioFactorProvider, error) {
	var localVarReturnValue *models.TwilioFactorProvider
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "phone", "providers", "twilio"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPnProviders Retrieve push notification provider
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_pn_providers
func (m *Manager) GetPnProviders(ctx context.Context, opts ...management.RequestOption) (*models.GetPnProviders200Response, error) {
	var localVarReturnValue *models.GetPnProviders200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "push-notification", "selected-provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetPolicies Get the Multi-factor Authentication policies
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

// GetSmsProviders Retrieve SMS configuration (one of auth0|twilio|phone-message-hook)
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_sms_providers
func (m *Manager) GetSmsProviders(ctx context.Context, opts ...management.RequestOption) (*models.GetPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetPhoneProviders200Response
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "sms", "selected-provider"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSmsTwilioFactorProvider Retrieve Twilio SMS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_sms_twilio_factor_provider
func (m *Manager) GetSmsTwilioFactorProvider(ctx context.Context, opts ...management.RequestOption) (*models.SmsTwilioFactorProvider, error) {
	var localVarReturnValue *models.SmsTwilioFactorProvider
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "sms", "providers", "twilio"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// GetSns Retrieve AWS SNS push notification configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/get_sns
func (m *Manager) GetSns(ctx context.Context, opts ...management.RequestOption) (*models.SnsFactorProvider, error) {
	var localVarReturnValue *models.SnsFactorProvider
	err := m.management.Request(ctx, "GET", m.management.URI("guardian", "factors", "push-notification", "providers", "sns"), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchApns Updates APNs provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/patch_apns
func (m *Manager) PatchApns(ctx context.Context, putApnsRequest *models.PutApnsRequest, opts ...management.RequestOption) (*models.PutApns200Response, error) {
	var localVarReturnValue *models.PutApns200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("guardian", "factors", "push-notification", "providers", "apns"), putApnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchFcm Updates FCM provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/patch_fcm
func (m *Manager) PatchFcm(ctx context.Context, putFcmRequest *models.PutFcmRequest, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "PATCH", m.management.URI("guardian", "factors", "push-notification", "providers", "fcm"), putFcmRequest, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// PatchSns Update SNS configuration for push notifications
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/patch_sns
func (m *Manager) PatchSns(ctx context.Context, putSnsRequest *models.PutSnsRequest, opts ...management.RequestOption) (*models.PutSnsRequest, error) {
	var localVarReturnValue *models.PutSnsRequest
	err := m.management.Request(ctx, "PATCH", m.management.URI("guardian", "factors", "push-notification", "providers", "sns"), putSnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostTicket Create a multi-factor authentication enrollment ticket
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/post_ticket
func (m *Manager) PostTicket(ctx context.Context, enrollmentCreate *models.EnrollmentCreate, opts ...management.RequestOption) (*models.PostTicket200Response, error) {
	var localVarReturnValue *models.PostTicket200Response
	err := m.management.Request(ctx, "POST", m.management.URI("guardian", "enrollments", "ticket"), enrollmentCreate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutApns Updates APNs provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_apns
func (m *Manager) PutApns(ctx context.Context, putApnsRequest *models.PutApnsRequest, opts ...management.RequestOption) (*models.PutApns200Response, error) {
	var localVarReturnValue *models.PutApns200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "providers", "apns"), putApnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutFactorPhoneTemplates Update Enrollment and Verification Phone Templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_factor_phone_templates
func (m *Manager) PutFactorPhoneTemplates(ctx context.Context, templateMessages *models.TemplateMessages, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "templates"), templateMessages, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutFactorSmsTemplates Update SMS Enrollment and Verification Templates
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_factor_sms_templates
func (m *Manager) PutFactorSmsTemplates(ctx context.Context, templateMessages *models.TemplateMessages, opts ...management.RequestOption) (*models.TemplateMessages, error) {
	var localVarReturnValue *models.TemplateMessages
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "sms", "templates"), templateMessages, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutFactorsByName Update a Multi-factor Authentication Factor
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_factors_by_name
func (m *Manager) PutFactorsByName(ctx context.Context, name models.PutFactorsByNameNameParameter, putFactorsByNameRequest *models.PutFactorsByNameRequest, opts ...management.RequestOption) (*models.PutFactorsByName200Response, error) {
	var localVarReturnValue *models.PutFactorsByName200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", string(name)), putFactorsByNameRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutFcm Updates FCM provider configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_fcm
func (m *Manager) PutFcm(ctx context.Context, putFcmRequest *models.PutFcmRequest, opts ...management.RequestOption) (*map[string]interface{}, error) {
	var localVarReturnValue *map[string]interface{}
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "providers", "fcm"), putFcmRequest, &localVarReturnValue, opts...)
	if err != nil {
		return localVarReturnValue, err
	}
	return localVarReturnValue, nil
}

// PutMessageTypes Update the Enabled Phone Factors
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_message_types
func (m *Manager) PutMessageTypes(ctx context.Context, putMessageTypesRequest *models.PutMessageTypesRequest, opts ...management.RequestOption) (*models.GetMessageTypes200Response, error) {
	var localVarReturnValue *models.GetMessageTypes200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "message-types"), putMessageTypesRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutPhoneProviders Update phone configuration (one of auth0|twilio|phone-message-hook)
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_phone_providers
func (m *Manager) PutPhoneProviders(ctx context.Context, putPhoneProvidersRequest *models.PutPhoneProvidersRequest, opts ...management.RequestOption) (*models.GetPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetPhoneProviders200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "selected-provider"), putPhoneProvidersRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutPnProviders Update Push Notification configuration (one of direct|sns|guardian)
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_pn_providers
func (m *Manager) PutPnProviders(ctx context.Context, putPnProvidersRequest *models.PutPnProvidersRequest, opts ...management.RequestOption) (*models.GetPnProviders200Response, error) {
	var localVarReturnValue *models.GetPnProviders200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "selected-provider"), putPnProvidersRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutPolicies Set the Multi-factor Authentication policies
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_policies
func (m *Manager) PutPolicies(ctx context.Context, getPolicies200ResponseInner []*models.GetPolicies200ResponseInner, opts ...management.RequestOption) ([]*models.GetPolicies200ResponseInner, error) {
	var localVarReturnValue []*models.GetPolicies200ResponseInner
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "policies"), getPolicies200ResponseInner, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutSmsProviders Update SMS configuration (one of auth0|twilio|phone-message-hook)
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_sms_providers
func (m *Manager) PutSmsProviders(ctx context.Context, putPhoneProvidersRequest *models.PutPhoneProvidersRequest, opts ...management.RequestOption) (*models.GetPhoneProviders200Response, error) {
	var localVarReturnValue *models.GetPhoneProviders200Response
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "sms", "selected-provider"), putPhoneProvidersRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutSmsTwilioFactorProvider Update Twilio SMS configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_sms_twilio_factor_provider
func (m *Manager) PutSmsTwilioFactorProvider(ctx context.Context, putTwilioRequest *models.PutTwilioRequest, opts ...management.RequestOption) (*models.SmsTwilioFactorProvider, error) {
	var localVarReturnValue *models.SmsTwilioFactorProvider
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "sms", "providers", "twilio"), putTwilioRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutSns Update AWS SNS push notification configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_sns
func (m *Manager) PutSns(ctx context.Context, putSnsRequest *models.PutSnsRequest, opts ...management.RequestOption) (*models.PutSnsRequest, error) {
	var localVarReturnValue *models.PutSnsRequest
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "push-notification", "providers", "sns"), putSnsRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutTwilio Update Twilio phone configuration
//
// https://auth0.com/docs/api/management/v2/#!/Guardian/put_twilio
func (m *Manager) PutTwilio(ctx context.Context, putTwilioRequest *models.PutTwilioRequest, opts ...management.RequestOption) (*models.TwilioFactorProvider, error) {
	var localVarReturnValue *models.TwilioFactorProvider
	err := m.management.Request(ctx, "PUT", m.management.URI("guardian", "factors", "phone", "providers", "twilio"), putTwilioRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
