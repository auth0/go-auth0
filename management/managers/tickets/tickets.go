package tickets

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

// VerifyEmail Create an email verification ticket
//
// https://auth0.com/docs/api/management/v2/#!/Tickets/post_email_verification
func (m *Manager) VerifyEmail(ctx context.Context, postEmailVerificationRequest *models.PostEmailVerificationRequest, opts ...management.RequestOption) (*models.PostEmailVerification201Response, error) {
	var localVarReturnValue *models.PostEmailVerification201Response
	err := m.management.Request(ctx, "POST", m.management.URI("tickets", "email-verification"), postEmailVerificationRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}

// ChangePassword Create a password change ticket
//
// https://auth0.com/docs/api/management/v2/#!/Tickets/post_password_change
func (m *Manager) ChangePassword(ctx context.Context, postPasswordChangeRequest *models.PostPasswordChangeRequest, opts ...management.RequestOption) (*models.PostPasswordChange201Response, error) {
	var localVarReturnValue *models.PostPasswordChange201Response
	err := m.management.Request(ctx, "POST", m.management.URI("tickets", "password-change"), postPasswordChangeRequest, &localVarReturnValue, opts...)
	if err != nil {
		return nil, err
	}
	return localVarReturnValue, nil
}
