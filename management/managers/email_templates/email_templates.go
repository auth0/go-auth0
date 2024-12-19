package emailTemplates

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

// Get Get an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/get_email_templates_by_template_name
func (m *Manager) Get(ctx context.Context, templateName models.GetEmailTemplatesByTemplateNameTemplateNameParameter, opts ...management.RequestOption) (*models.GetEmailTemplatesByTemplateName200Response, error) {
	var localVarReturnValue *models.GetEmailTemplatesByTemplateName200Response
	err := m.management.Request(ctx, "GET", m.management.URI("email-templates", string(templateName)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Update Patch an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/patch_email_templates_by_template_name
func (m *Manager) Update(ctx context.Context, templateName models.GetEmailTemplatesByTemplateNameTemplateNameParameter, getEmailTemplatesByTemplateName200Response *models.GetEmailTemplatesByTemplateName200Response, opts ...management.RequestOption) (*models.GetEmailTemplatesByTemplateName200Response, error) {
	var localVarReturnValue *models.GetEmailTemplatesByTemplateName200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("email-templates", string(templateName)), getEmailTemplatesByTemplateName200Response, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Create Create an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/post_email_templates
func (m *Manager) Create(ctx context.Context, postEmailTemplatesRequest *models.PostEmailTemplatesRequest, opts ...management.RequestOption) (*models.PostEmailTemplatesRequest, error) {
	var localVarReturnValue *models.PostEmailTemplatesRequest
	err := m.management.Request(ctx, "POST", m.management.URI("email-templates"), postEmailTemplatesRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// Put Update an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/put_email_templates_by_template_name
func (m *Manager) Put(ctx context.Context, templateName models.GetEmailTemplatesByTemplateNameTemplateNameParameter, emailTemplateUpdate *models.EmailTemplateUpdate, opts ...management.RequestOption) (*models.PostEmailTemplatesRequest, error) {
	var localVarReturnValue *models.PostEmailTemplatesRequest
	err := m.management.Request(ctx, "PUT", m.management.URI("email-templates", string(templateName)), emailTemplateUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
