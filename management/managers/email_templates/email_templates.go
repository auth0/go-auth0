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

// GetEmailTemplatesByTemplateName Get an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/get_email_templates_by_template_name
func (m *Manager) GetEmailTemplatesByTemplateName(ctx context.Context, templateName models.GetEmailTemplatesByTemplateNameTemplateNameParameter, opts ...management.RequestOption) (*models.GetEmailTemplatesByTemplateName200Response, error) {
	var localVarReturnValue *models.GetEmailTemplatesByTemplateName200Response
	err := m.management.Request(ctx, "GET", m.management.URI("email-templates", string(templateName)), nil, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PatchEmailTemplatesByTemplateName Patch an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/patch_email_templates_by_template_name
func (m *Manager) PatchEmailTemplatesByTemplateName(ctx context.Context, templateName models.GetEmailTemplatesByTemplateNameTemplateNameParameter, patchEmailTemplatesByTemplateNameRequest *models.PatchEmailTemplatesByTemplateNameRequest, opts ...management.RequestOption) (*models.GetEmailTemplatesByTemplateName200Response, error) {
	var localVarReturnValue *models.GetEmailTemplatesByTemplateName200Response
	err := m.management.Request(ctx, "PATCH", m.management.URI("email-templates", string(templateName)), patchEmailTemplatesByTemplateNameRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PostEmailTemplates Create an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/post_email_templates
func (m *Manager) PostEmailTemplates(ctx context.Context, postEmailTemplatesRequest *models.PostEmailTemplatesRequest, opts ...management.RequestOption) (*models.PostEmailTemplatesRequest, error) {
	var localVarReturnValue *models.PostEmailTemplatesRequest
	err := m.management.Request(ctx, "POST", m.management.URI("email-templates"), postEmailTemplatesRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// PutEmailTemplatesByTemplateName Update an email template
//
// https://auth0.com/docs/api/management/v2/#!/EmailTemplates/put_email_templates_by_template_name
func (m *Manager) PutEmailTemplatesByTemplateName(ctx context.Context, templateName models.GetEmailTemplatesByTemplateNameTemplateNameParameter, emailTemplateUpdate *models.EmailTemplateUpdate, opts ...management.RequestOption) (*models.PostEmailTemplatesRequest, error) {
	var localVarReturnValue *models.PostEmailTemplatesRequest
	err := m.management.Request(ctx, "PUT", m.management.URI("email-templates", string(templateName)), emailTemplateUpdate, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
